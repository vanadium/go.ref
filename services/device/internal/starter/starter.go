// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package starter provides a single function that starts up servers for a
// mounttable and a device manager that is mounted on it.
package starter

import (
	"encoding/base64"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"v.io/x/ref/runtime/factories/roaming"
	"v.io/x/ref/services/debug/debuglib"
	"v.io/x/ref/services/device/internal/config"
	"v.io/x/ref/services/device/internal/impl"
	"v.io/x/ref/services/internal/pathperms"
	"v.io/x/ref/services/mounttable/mounttablelib"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/naming"
	"v.io/v23/rpc"
	"v.io/v23/security"
	"v.io/v23/verror"
	"v.io/x/lib/vlog"
	"v.io/x/ref/lib/xrpc"
)

const pkgPath = "v.io/x/ref/services/device/internal/starter"

var (
	errCantSaveInfo      = verror.Register(pkgPath+".errCantSaveInfo", verror.NoRetry, "{1:}{2:} failed to save info{:_}")
	errBadPort           = verror.Register(pkgPath+".errBadPort", verror.NoRetry, "{1:}{2:} invalid port{:_}")
	errCantCreateProxy   = verror.Register(pkgPath+".errCantCreateProxy", verror.NoRetry, "{1:}{2:} Failed to create proxy{:_}")
	errNoEndpointToClaim = verror.Register(pkgPath+".errNoEndpointToClaim", verror.NoRetry, "{1:}{2:} failed to find an endpoint for claiming{:_}")
)

type NamespaceArgs struct {
	Name            string         // Name to publish the mounttable service under (after claiming).
	ListenSpec      rpc.ListenSpec // ListenSpec for the server.
	PermissionsFile string         // Path to the Permissions file used by the mounttable.
	PersistenceDir  string         // Path to the directory holding persistent acls.
	// Name in the local neighborhood on which to make the mounttable
	// visible. If empty, the mounttable will not be visible in the local
	// neighborhood.
	Neighborhood string
}

type DeviceArgs struct {
	Name            string         // Name to publish the device service under (after claiming).
	ListenSpec      rpc.ListenSpec // ListenSpec for the device server.
	ConfigState     *config.State  // Configuration for the device.
	TestMode        bool           // Whether the device is running in test mode or not.
	RestartCallback func()         // Callback invoked when the device service is restarted.
	PairingToken    string         // PairingToken that a claimer needs to provide.
}

func (d *DeviceArgs) name(mt string) string {
	if d.Name != "" {
		return d.Name
	}
	return naming.Join(mt, "devmgr")
}

type ProxyArgs struct {
	Port int
}

type Args struct {
	Namespace NamespaceArgs
	Device    DeviceArgs
	Proxy     ProxyArgs

	// If true, the global namespace will be made available on the
	// mounttable server under "global/".
	MountGlobalNamespaceInLocalNamespace bool
}

// Start creates servers for the mounttable and device services and links them together.
//
// Returns the object name for the claimable service (empty if already claimed),
// a callback to be invoked to shutdown the services on success, or an error on
// failure.
func Start(ctx *context.T, args Args) (string, func(), error) {
	// Is this binary compatible with the state on disk?
	if err := impl.CheckCompatibility(args.Device.ConfigState.Root); err != nil {
		return "", nil, err
	}
	// In test mode, we skip writing the info file to disk, and we skip
	// attempting to start the claimable service: the device must have been
	// claimed already to enable updates anyway, and checking for perms in
	// NewClaimableDispatcher needlessly prints a perms signature
	// verification error to the logs.
	if args.Device.TestMode {
		cleanup, err := startClaimedDevice(ctx, args)
		return "", cleanup, err
	}

	// TODO(caprita): use some mechanism (a file lock or presence of entry
	// in mounttable) to ensure only one device manager is running in an
	// installation?
	mi := &impl.ManagerInfo{
		Pid: os.Getpid(),
	}
	if err := impl.SaveManagerInfo(filepath.Join(args.Device.ConfigState.Root, "device-manager"), mi); err != nil {
		return "", nil, verror.New(errCantSaveInfo, ctx, err)
	}

	// If the device has not yet been claimed, start the mounttable and
	// claimable service and wait for it to be claimed.
	// Once a device is claimed, close any previously running servers and
	// start a new mounttable and device service.
	claimable, claimed := impl.NewClaimableDispatcher(ctx, args.Device.ConfigState, args.Device.PairingToken)
	if claimable == nil {
		// Device has already been claimed, bypass claimable service
		// stage.
		cleanup, err := startClaimedDevice(ctx, args)
		return "", cleanup, err
	}
	epName, stopClaimable, err := startClaimableDevice(ctx, claimable, args)
	if err != nil {
		return "", nil, err
	}
	stop := make(chan struct{})
	stopped := make(chan struct{})
	go waitToBeClaimedAndStartClaimedDevice(ctx, stopClaimable, claimed, stop, stopped, args)
	return epName, func() {
		close(stop)
		<-stopped
	}, nil
}

func startClaimableDevice(ctx *context.T, dispatcher rpc.Dispatcher, args Args) (string, func(), error) {
	// TODO(caprita,ashankar): We create a context with a new stream manager
	// that we can cancel once the device has been claimed. This gets around
	// the following issue: if we publish the claimable server to the local
	// mounttable, and then (following claim) we restart the mounttable
	// server on the same port, we fail to publish the device service to the
	// (new) mounttable server (Mount fails with "VC handshake failed:
	// remote end closed VC(VCs not accepted)".  Presumably, something to do
	// with caching connections (following the claim, the mounttable comes
	// back on the same port as before, and the client-side of the mount
	// gets confused trying to reuse the old connection and doesn't attempt
	// to create a new connection).  We should get to the bottom of it.
	ctx, cancel := context.WithCancel(ctx)
	var err error
	if ctx, err = v23.WithNewStreamManager(ctx); err != nil {
		cancel()
		return "", nil, err
	}
	ctx = v23.WithListenSpec(ctx, args.Device.ListenSpec)
	server, err := xrpc.NewDispatchingServer(ctx, "", dispatcher)
	if err != nil {
		cancel()
		return "", nil, err
	}
	shutdown := func() {
		vlog.Infof("Stopping claimable server...")
		server.Stop()
		vlog.Infof("Stopped claimable server.")
		cancel()
	}
	endpoints := server.Status().Endpoints
	publicKey, err := v23.GetPrincipal(ctx).PublicKey().MarshalBinary()
	if err != nil {
		shutdown()
		return "", nil, err
	}
	var epName string
	if args.Device.ListenSpec.Proxy != "" {
		for {
			p := server.Status().Proxies
			if len(p) == 0 {
				vlog.Infof("Waiting for proxy address to appear...")
				time.Sleep(time.Second)
				continue
			}
			epName = p[0].Endpoint.Name()
			vlog.Infof("Proxied address: %s", epName)
			break
		}
	} else {
		if len(endpoints) == 0 {
			return "", nil, verror.New(errNoEndpointToClaim, ctx, err)
		}
		epName = endpoints[0].Name()
	}
	vlog.Infof("Unclaimed device manager (%v) with public_key: %s", epName, base64.URLEncoding.EncodeToString(publicKey))
	vlog.FlushLog()
	return epName, shutdown, nil
}

func waitToBeClaimedAndStartClaimedDevice(ctx *context.T, stopClaimable func(), claimed, stop <-chan struct{}, stopped chan<- struct{}, args Args) {
	// Wait for either the claimable service to complete, or be stopped
	defer close(stopped)
	select {
	case <-claimed:
		stopClaimable()
	case <-stop:
		stopClaimable()
		return
	}
	shutdown, err := startClaimedDevice(ctx, args)
	if err != nil {
		vlog.Errorf("Failed to start device service after it was claimed: %v", err)
		v23.GetAppCycle(ctx).Stop()
		return
	}
	defer shutdown()
	<-stop // Wait to be stopped
}

func startClaimedDevice(ctx *context.T, args Args) (func(), error) {
	permStore := pathperms.NewPathStore(ctx)
	permsdir := impl.PermsDir(args.Device.ConfigState)
	debugAuth, err := pathperms.NewHierarchicalAuthorizer(permsdir, permsdir, permStore)
	if err != nil {
		return nil, err
	}

	debugDisp := debuglib.NewDispatcher(vlog.Log.LogDir, debugAuth)

	ctx = v23.WithReservedNameDispatcher(ctx, debugDisp)

	mtName, stopMT, err := startMounttable(ctx, args.Namespace)
	if err != nil {
		vlog.Errorf("Failed to start mounttable service: %v", err)
		return nil, err
	}
	// TODO(caprita): We link in a proxy server into the device manager so that we
	// can bootstrap with install-local before we can install an actual proxy app.
	// Once support is added to the RPC layer to allow install-local to serve on
	// the same connection it established to the device manager (see TODO in
	// v.io/x/ref/services/device/device/local_install.go), we can get rid of this
	// local proxy altogether.
	stopProxy, err := startProxyServer(ctx, args.Proxy, mtName)
	if err != nil {
		vlog.Errorf("Failed to start proxy service: %v", err)
		stopMT()
		return nil, err
	}
	stopDevice, err := startDeviceServer(ctx, args.Device, mtName, permStore)
	if err != nil {
		vlog.Errorf("Failed to start device service: %v", err)
		stopProxy()
		stopMT()
		return nil, err
	}
	if args.MountGlobalNamespaceInLocalNamespace {
		mountGlobalNamespaceInLocalNamespace(ctx, mtName)
	}

	impl.InvokeCallback(ctx, args.Device.ConfigState.Name)

	return func() {
		stopDevice()
		stopProxy()
		stopMT()
	}, nil
}

func startProxyServer(ctx *context.T, p ProxyArgs, localMT string) (func(), error) {
	switch port := p.Port; {
	case port == 0:
		return func() {}, nil
	case port < 0:
		return nil, verror.New(errBadPort, ctx, port)
	}
	port := strconv.Itoa(p.Port)
	protocol, addr := "tcp", net.JoinHostPort("", port)
	// Attempt to get a publicly accessible address for the proxy to publish
	// under.
	ls := v23.GetListenSpec(ctx)
	ls.Addrs = rpc.ListenAddrs{{protocol, addr}}
	// TODO(ashankar): Revisit this choice of security.AllowEveryone
	// See: https://v.io/i/387
	shutdown, ep, err := roaming.NewProxy(ctx, ls, security.AllowEveryone())
	if err != nil {
		return nil, verror.New(errCantCreateProxy, ctx, err)
	}
	vlog.Infof("Local proxy (%v)", ep.Name())
	return func() {
		vlog.Infof("Stopping proxy...")
		shutdown()
		vlog.Infof("Stopped proxy.")
	}, nil
}

func startMounttable(ctx *context.T, n NamespaceArgs) (string, func(), error) {
	mtName, stopMT, err := mounttablelib.StartServers(ctx, n.ListenSpec, n.Name, n.Neighborhood, n.PermissionsFile, n.PersistenceDir, "mounttable")
	if err != nil {
		vlog.Errorf("mounttablelib.StartServers(%#v) failed: %v", n, err)
	} else {
		vlog.Infof("Local mounttable (%v) published as %q", mtName, n.Name)
	}
	return mtName, func() {
		vlog.Infof("Stopping mounttable...")
		stopMT()
		vlog.Infof("Stopped mounttable.")
	}, err
}

// startDeviceServer creates an rpc.Server and sets it up to server the Device service.
//
// ls: ListenSpec for the server
// configState: configuration for the Device service dispatcher
// mt: Object address of the mounttable
// dm: Name to publish the device service under
// testMode: whether the service is to be run in test mode
// restarted: callback invoked when the device manager is restarted.
//
// Returns:
// (1) Function to be called to force the service to shutdown
// (2) Any errors in starting the service (in which case, (1) will be nil)
func startDeviceServer(ctx *context.T, args DeviceArgs, mt string, permStore *pathperms.PathStore) (shutdown func(), err error) {
	server, err := v23.NewServer(ctx)
	if err != nil {
		return nil, err
	}
	shutdown = func() { server.Stop() }
	endpoints, err := server.Listen(args.ListenSpec)
	if err != nil {
		shutdown()
		return nil, err
	}
	args.ConfigState.Name = endpoints[0].Name()

	dispatcher, err := impl.NewDispatcher(ctx, args.ConfigState, mt, args.TestMode, args.RestartCallback, permStore)
	if err != nil {
		shutdown()
		return nil, err
	}

	shutdown = func() {
		// TODO(caprita): Capture the Dying state by feeding it back to
		// the dispatcher and exposing it in Status.
		vlog.Infof("Stopping device server...")
		server.Stop()
		impl.Shutdown(dispatcher)
		vlog.Infof("Stopped device.")
	}
	if err := server.ServeDispatcher(args.name(mt), dispatcher); err != nil {
		shutdown()
		return nil, err
	}
	vlog.Infof("Device manager (%v) published as %v", args.ConfigState.Name, args.name(mt))
	return shutdown, nil
}

func mountGlobalNamespaceInLocalNamespace(ctx *context.T, localMT string) {
	ns := v23.GetNamespace(ctx)
	for _, root := range ns.Roots() {
		go func(r string) {
			for {
				err := ns.Mount(ctx, naming.Join(localMT, "global"), r, 0 /* forever */, naming.ServesMountTable(true))
				if err == nil {
					break
				}
				vlog.Infof("Failed to Mount global namespace: %v", err)
				time.Sleep(time.Second)
			}
		}(root)
	}
}
