package impl

// The node invoker is responsible for managing the state of the node manager
// itself.  The implementation expects that the node manager installations are
// all organized in the following directory structure:
//
// <config.Root>/
//   node-manager/
//     <version 1 timestamp>/  - timestamp of when the version was downloaded
//       noded                 - the node manager binary
//       noded.sh              - a shell script to start the binary
//     <version 2 timestamp>
//     ...
//     node-data/
//       acl.nodemanager
//	 acl.signature
//	 associated.accounts
//
// The node manager is always expected to be started through the symbolic link
// passed in as config.CurrentLink, which is monitored by an init daemon. This
// provides for simple and robust updates.
//
// To update the node manager to a newer version, a new workspace is created and
// the symlink is updated to the new noded.sh script. Similarly, to revert the
// node manager to a previous version, all that is required is to update the
// symlink to point to the previous noded.sh script.

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"

	"veyron.io/veyron/veyron2"
	"veyron.io/veyron/veyron2/context"
	"veyron.io/veyron/veyron2/ipc"
	"veyron.io/veyron/veyron2/mgmt"
	"veyron.io/veyron/veyron2/naming"
	"veyron.io/veyron/veyron2/services/mgmt/application"
	"veyron.io/veyron/veyron2/services/mgmt/binary"
	"veyron.io/veyron/veyron2/services/mgmt/node"
	"veyron.io/veyron/veyron2/services/security/access"
	"veyron.io/veyron/veyron2/verror2"
	"veyron.io/veyron/veyron2/vlog"

	vexec "veyron.io/veyron/veyron/lib/exec"
	"veyron.io/veyron/veyron/lib/netstate"
	"veyron.io/veyron/veyron/services/mgmt/node/config"
	"veyron.io/veyron/veyron/services/mgmt/profile"
)

type updatingState struct {
	// updating is a flag that records whether this instance of node
	// manager is being updated.
	updating bool
	// updatingMutex is a lock for coordinating concurrent access to
	// <updating>.
	updatingMutex sync.Mutex
}

func newUpdatingState() *updatingState {
	return new(updatingState)
}

func (u *updatingState) testAndSetUpdating() bool {
	u.updatingMutex.Lock()
	defer u.updatingMutex.Unlock()
	if u.updating {
		return true
	}
	u.updating = true
	return false
}

func (u *updatingState) unsetUpdating() {
	u.updatingMutex.Lock()
	u.updating = false
	u.updatingMutex.Unlock()
}

// nodeService implements the Node manager's Node interface.
type nodeService struct {
	updating *updatingState
	callback *callbackState
	config   *config.State
	disp     *dispatcher
	uat      BlessingSystemAssociationStore
}

func (i *nodeService) Claim(call ipc.ServerContext) error {
	// Get the blessing to be used by the claimant.
	blessings := call.Blessings()
	if blessings == nil {
		return verror2.Make(ErrInvalidBlessing, call)
	}
	return i.disp.claimNodeManager(call.LocalPrincipal(), blessings.ForContext(call), blessings)
}

func (*nodeService) Describe(ipc.ServerContext) (node.Description, error) {
	empty := node.Description{}
	nodeProfile, err := computeNodeProfile()
	if err != nil {
		return empty, err
	}
	knownProfiles, err := getKnownProfiles()
	if err != nil {
		return empty, err
	}
	result := matchProfiles(nodeProfile, knownProfiles)
	return result, nil
}

func (*nodeService) IsRunnable(_ ipc.ServerContext, description binary.Description) (bool, error) {
	nodeProfile, err := computeNodeProfile()
	if err != nil {
		return false, err
	}
	binaryProfiles := make([]profile.Specification, 0)
	for name, _ := range description.Profiles {
		profile, err := getProfile(name)
		if err != nil {
			return false, err
		}
		binaryProfiles = append(binaryProfiles, *profile)
	}
	result := matchProfiles(nodeProfile, binaryProfiles)
	return len(result.Profiles) > 0, nil
}

func (*nodeService) Reset(call ipc.ServerContext, deadline uint64) error {
	// TODO(jsimsa): Implement.
	return nil
}

// getCurrentFileInfo returns the os.FileInfo for both the symbolic link
// CurrentLink, and the node script in the workspace that this link points to.
func (i *nodeService) getCurrentFileInfo() (os.FileInfo, string, error) {
	path := i.config.CurrentLink
	link, err := os.Lstat(path)
	if err != nil {
		vlog.Errorf("Lstat(%v) failed: %v", path, err)
		return nil, "", err
	}
	scriptPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		vlog.Errorf("EvalSymlinks(%v) failed: %v", path, err)
		return nil, "", err
	}
	return link, scriptPath, nil
}

func (i *nodeService) revertNodeManager(ctx context.T) error {
	if err := updateLink(i.config.Previous, i.config.CurrentLink); err != nil {
		return err
	}
	runtime := veyron2.RuntimeFromContext(ctx)
	runtime.AppCycle().Stop()
	return nil
}

func (i *nodeService) newLogfile(prefix string) (*os.File, error) {
	d := filepath.Join(i.config.Root, "node_test_logs")
	if _, err := os.Stat(d); err != nil {
		if err := os.MkdirAll(d, 0700); err != nil {
			return nil, err
		}
	}
	f, err := ioutil.TempFile(d, "__node_impl_test__"+prefix)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// TODO(cnicolaou): would this be better implemented using the modules
// framework now that it exists?
func (i *nodeService) testNodeManager(ctx context.T, workspace string, envelope *application.Envelope) error {
	path := filepath.Join(workspace, "noded.sh")
	cmd := exec.Command(path)

	for k, v := range map[string]*io.Writer{
		"stdout": &cmd.Stdout,
		"stderr": &cmd.Stderr,
	} {
		// Using a log file makes it less likely that stdout and stderr
		// output will be lost if the child crashes.
		file, err := i.newLogfile(fmt.Sprintf("noded-test-%s", k))
		if err != nil {
			return err
		}
		fName := file.Name()
		defer os.Remove(fName)
		*v = file

		defer func() {
			if f, err := os.Open(fName); err == nil {
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					vlog.Infof("[testNodeManager %s] %s", k, scanner.Text())
				}
			}
		}()
	}

	// Setup up the child process callback.
	callbackState := i.callback
	listener := callbackState.listenFor(mgmt.ChildNameConfigKey)
	defer listener.cleanup()
	cfg := vexec.NewConfig()

	cfg.Set(mgmt.ParentNameConfigKey, listener.name())
	cfg.Set(mgmt.ProtocolConfigKey, "tcp")
	cfg.Set(mgmt.AddressConfigKey, "127.0.0.1:0")
	handle := vexec.NewParentHandle(cmd, vexec.ConfigOpt{cfg})
	// Start the child process.
	if err := handle.Start(); err != nil {
		vlog.Errorf("Start() failed: %v", err)
		return verror2.Make(ErrOperationFailed, ctx)
	}
	defer func() {
		if err := handle.Clean(); err != nil {
			vlog.Errorf("Clean() failed: %v", err)
		}
	}()
	// Wait for the child process to start.
	if err := handle.WaitForReady(childReadyTimeout); err != nil {
		vlog.Errorf("WaitForReady(%v) failed: %v", childReadyTimeout, err)
		return verror2.Make(ErrOperationFailed, ctx)
	}
	childName, err := listener.waitForValue(childReadyTimeout)
	if err != nil {
		return verror2.Make(ErrOperationFailed, ctx)
	}
	// Check that invoking Update() succeeds.
	childName = naming.Join(childName, "nm")
	nmClient := node.NodeClient(childName)
	linkOld, pathOld, err := i.getCurrentFileInfo()
	if err != nil {
		return verror2.Make(ErrOperationFailed, ctx)
	}
	// Since the resolution of mtime for files is seconds, the test sleeps
	// for a second to make sure it can check whether the current symlink is
	// updated.
	time.Sleep(time.Second)
	if err := nmClient.Revert(ctx); err != nil {
		return verror2.Make(ErrOperationFailed, ctx)
	}
	linkNew, pathNew, err := i.getCurrentFileInfo()
	if err != nil {
		return verror2.Make(ErrOperationFailed, ctx)
	}
	// Check that the new node manager updated the current symbolic link.
	if !linkOld.ModTime().Before(linkNew.ModTime()) {
		vlog.Errorf("new node manager test failed")
		return verror2.Make(ErrOperationFailed, ctx)
	}
	// Ensure that the current symbolic link points to the same script.
	if pathNew != pathOld {
		updateLink(pathOld, i.config.CurrentLink)
		vlog.Errorf("new node manager test failed")
		return verror2.Make(ErrOperationFailed, ctx)
	}
	return nil
}

// TODO(caprita): Move this to util.go since node_installer is also using it now.

func generateScript(workspace string, configSettings []string, envelope *application.Envelope) error {
	// TODO(caprita): Remove this snippet of code, it doesn't seem to serve
	// any purpose.
	path, err := filepath.EvalSymlinks(os.Args[0])
	if err != nil {
		vlog.Errorf("EvalSymlinks(%v) failed: %v", os.Args[0], err)
		return verror2.Make(ErrOperationFailed, nil)
	}

	output := "#!/bin/bash\n"
	output += strings.Join(config.QuoteEnv(append(envelope.Env, configSettings...)), " ") + " "
	// Escape the path to the binary; %q uses Go-syntax escaping, but it's
	// close enough to Bash that we're using it as an approximation.
	//
	// TODO(caprita/rthellend): expose and use shellEscape (from
	// veyron/tools/debug/impl.go) instead.
	output += fmt.Sprintf("%q", filepath.Join(workspace, "noded")) + " "
	output += strings.Join(envelope.Args, " ")
	output += "\n"
	path = filepath.Join(workspace, "noded.sh")
	if err := ioutil.WriteFile(path, []byte(output), 0700); err != nil {
		vlog.Errorf("WriteFile(%v) failed: %v", path, err)
		return verror2.Make(ErrOperationFailed, nil)
	}
	return nil
}

func (i *nodeService) updateNodeManager(ctx context.T) error {
	if len(i.config.Origin) == 0 {
		return verror2.Make(ErrUpdateNoOp, ctx)
	}
	envelope, err := fetchEnvelope(ctx, i.config.Origin)
	if err != nil {
		return err
	}
	if envelope.Title != application.NodeManagerTitle {
		return verror2.Make(ErrAppTitleMismatch, ctx)
	}
	if i.config.Envelope != nil && reflect.DeepEqual(envelope, i.config.Envelope) {
		return verror2.Make(ErrUpdateNoOp, ctx)
	}
	// Create new workspace.
	workspace := filepath.Join(i.config.Root, "node-manager", generateVersionDirName())
	perm := os.FileMode(0700)
	if err := os.MkdirAll(workspace, perm); err != nil {
		vlog.Errorf("MkdirAll(%v, %v) failed: %v", workspace, perm, err)
		return verror2.Make(ErrOperationFailed, ctx)
	}

	deferrer := func() {
		cleanupDir(workspace, "")
	}
	defer func() {
		if deferrer != nil {
			deferrer()
		}
	}()

	// Populate the new workspace with a node manager binary.
	// TODO(caprita): match identical binaries on binary metadata
	// rather than binary object name.
	sameBinary := i.config.Envelope != nil && envelope.Binary == i.config.Envelope.Binary
	if sameBinary {
		if err := linkSelf(workspace, "noded"); err != nil {
			return err
		}
	} else {
		if err := downloadBinary(ctx, workspace, "noded", envelope.Binary); err != nil {
			return err
		}
	}

	// Populate the new workspace with a node manager script.
	configSettings, err := i.config.Save(envelope)
	if err != nil {
		return verror2.Make(ErrOperationFailed, ctx)
	}

	if err := generateScript(workspace, configSettings, envelope); err != nil {
		return err
	}

	if err := i.testNodeManager(ctx, workspace, envelope); err != nil {
		return err
	}

	if err := updateLink(filepath.Join(workspace, "noded.sh"), i.config.CurrentLink); err != nil {
		return err
	}

	runtime := veyron2.RuntimeFromContext(ctx)
	runtime.AppCycle().Stop()
	deferrer = nil
	return nil
}

func (*nodeService) Install(ctx ipc.ServerContext, _ string) (string, error) {
	return "", verror2.Make(ErrInvalidSuffix, ctx)
}

func (*nodeService) Refresh(ipc.ServerContext) error {
	// TODO(jsimsa): Implement.
	return nil
}

func (*nodeService) Restart(ipc.ServerContext) error {
	// TODO(jsimsa): Implement.
	return nil
}

func (*nodeService) Resume(ctx ipc.ServerContext) error {
	return verror2.Make(ErrInvalidSuffix, ctx)
}

func (i *nodeService) Revert(call ipc.ServerContext) error {
	if i.config.Previous == "" {
		return verror2.Make(ErrUpdateNoOp, call)
	}
	updatingState := i.updating
	if updatingState.testAndSetUpdating() {
		return verror2.Make(ErrOperationInProgress, call)
	}
	err := i.revertNodeManager(call)
	if err != nil {
		updatingState.unsetUpdating()
	}
	return err
}

func (*nodeService) Start(ctx ipc.ServerContext) ([]string, error) {
	return nil, verror2.Make(ErrInvalidSuffix, ctx)
}

func (*nodeService) Stop(ctx ipc.ServerContext, _ uint32) error {
	return verror2.Make(ErrInvalidSuffix, ctx)
}

func (*nodeService) Suspend(ipc.ServerContext) error {
	// TODO(jsimsa): Implement.
	return nil
}

func (*nodeService) Uninstall(ctx ipc.ServerContext) error {
	return verror2.Make(ErrInvalidSuffix, ctx)
}

func (i *nodeService) Update(call ipc.ServerContext) error {
	ctx, cancel := call.WithTimeout(ipcContextTimeout)
	defer cancel()

	updatingState := i.updating
	if updatingState.testAndSetUpdating() {
		return verror2.Make(ErrOperationInProgress, call)
	}

	err := i.updateNodeManager(ctx)
	if err != nil {
		updatingState.unsetUpdating()
	}
	return err
}

func (*nodeService) UpdateTo(ipc.ServerContext, string) error {
	// TODO(jsimsa): Implement.
	return nil
}

func (i *nodeService) SetACL(ctx ipc.ServerContext, acl access.TaggedACLMap, etag string) error {
	return i.disp.setACL(ctx.LocalPrincipal(), acl, etag, true /* store ACL on disk */)
}

func (i *nodeService) GetACL(_ ipc.ServerContext) (acl access.TaggedACLMap, etag string, err error) {
	return i.disp.getACL()
}

func sameMachineCheck(ctx ipc.ServerContext) error {
	switch local, err := netstate.SameMachine(ctx.RemoteEndpoint().Addr()); {
	case err != nil:
		return err
	case local == false:
		vlog.Errorf("SameMachine() indicates that endpoint is not on the same node")
		return verror2.Make(ErrOperationFailed, ctx)
	}
	return nil
}

// TODO(rjkroege): Make it possible for users on the same system to also
// associate their accounts with their identities.
func (i *nodeService) AssociateAccount(call ipc.ServerContext, identityNames []string, accountName string) error {
	if err := sameMachineCheck(call); err != nil {
		return err
	}

	if accountName == "" {
		return i.uat.DisassociateSystemAccountForBlessings(identityNames)
	} else {
		// TODO(rjkroege): Optionally verify here that the required uname is a valid.
		return i.uat.AssociateSystemAccountForBlessings(identityNames, accountName)
	}
}

func (i *nodeService) ListAssociations(call ipc.ServerContext) (associations []node.Association, err error) {
	// Temporary code. Dump this.
	vlog.VI(2).Infof("ListAssociations given blessings: %v\n", call.RemoteBlessings().ForContext(call))

	if err := sameMachineCheck(call); err != nil {
		return nil, err
	}
	return i.uat.AllBlessingSystemAssociations()
}
