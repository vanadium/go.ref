// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"syscall"
	"testing"

	"v.io/v23"
	"v.io/v23/naming"
	"v.io/v23/security"
	"v.io/v23/security/access"
	"v.io/v23/services/application"
	"v.io/v23/verror"
	"v.io/x/lib/vlog"

	"v.io/x/ref/lib/signals"
	appd "v.io/x/ref/services/application/applicationd"
	"v.io/x/ref/services/internal/servicetest"
	"v.io/x/ref/services/repository"
	"v.io/x/ref/test"
	"v.io/x/ref/test/testutil"
)

//go:generate v23 test generate

const (
	repoCmd = "appRepository"
)

func appRepository(stdin io.Reader, stdout, stderr io.Writer, env map[string]string, args ...string) error {
	if len(args) < 2 {
		vlog.Fatalf("repository expected at least name and store arguments and optionally AccessList flags per PermissionsFromFlag")
	}
	publishName := args[0]
	storedir := args[1]

	ctx, shutdown := test.InitForTest()
	defer shutdown()

	v23.GetNamespace(ctx).CacheCtl(naming.DisableCache(true))

	defer fmt.Fprintf(stdout, "%v terminating\n", publishName)
	defer vlog.VI(1).Infof("%v terminating", publishName)
	server, endpoint := servicetest.NewServer(ctx)
	defer server.Stop()

	name := naming.JoinAddressName(endpoint, "")
	vlog.VI(1).Infof("applicationd name: %v", name)

	dispatcher, err := appd.NewDispatcher(storedir)
	if err != nil {
		vlog.Fatalf("Failed to create repository dispatcher: %v", err)
	}
	if err := server.ServeDispatcher(publishName, dispatcher); err != nil {
		vlog.Fatalf("Serve(%v) failed: %v", publishName, err)
	}

	fmt.Fprintf(stdout, "ready:%d\n", os.Getpid())
	<-signals.ShutdownOnSignals(ctx)

	return nil
}

func TestApplicationUpdateAccessList(t *testing.T) {
	ctx, shutdown := test.InitForTest()
	defer shutdown()
	v23.GetNamespace(ctx).CacheCtl(naming.DisableCache(true))

	// By default, all principals in this test will have blessings
	// generated based on the username/machine running this process. Give
	// them recognizable names ("root/self" etc.), so the AccessLists can be set
	// deterministically.
	idp := testutil.NewIDProvider("root")
	if err := idp.Bless(v23.GetPrincipal(ctx), "self"); err != nil {
		t.Fatal(err)
	}

	sh, deferFn := servicetest.CreateShellAndMountTable(t, ctx, v23.GetPrincipal(ctx))
	defer deferFn()

	// setup mock up directory to put state in
	storedir, cleanup := servicetest.SetupRootDir(t, "application")
	defer cleanup()

	nmh := servicetest.RunCommand(t, sh, nil, repoCmd, "repo", storedir)
	pid := servicetest.ReadPID(t, nmh)
	defer syscall.Kill(pid, syscall.SIGINT)

	otherCtx, err := v23.SetPrincipal(ctx, testutil.NewPrincipal())
	if err != nil {
		t.Fatal(err)
	}
	if err := idp.Bless(v23.GetPrincipal(otherCtx), "other"); err != nil {
		t.Fatal(err)
	}

	v1stub := repository.ApplicationClient("repo/search/v1")
	repostub := repository.ApplicationClient("repo")

	// Create example envelopes.
	envelopeV1 := application.Envelope{
		Args:   []string{"--help"},
		Env:    []string{"DEBUG=1"},
		Binary: application.SignedFile{File: "/v23/name/of/binary"},
	}

	// Envelope putting as other should fail.
	if err := v1stub.Put(otherCtx, []string{"base"}, envelopeV1); verror.ErrorID(err) != verror.ErrNoAccess.ID {
		t.Fatalf("Put() returned errorid=%v wanted errorid=%v [%v]", verror.ErrorID(err), verror.ErrNoAccess.ID, err)
	}

	// Envelope putting as global should succeed.
	if err := v1stub.Put(ctx, []string{"base"}, envelopeV1); err != nil {
		t.Fatalf("Put() failed: %v", err)
	}

	vlog.VI(2).Infof("Accessing the Permission Lists of the root returns a (simulated) list providing default authorization.")
	acl, version, err := repostub.GetPermissions(ctx)
	if err != nil {
		t.Fatalf("GetPermissions should not have failed: %v", err)
	}
	if got, want := version, ""; got != want {
		t.Fatalf("GetPermissions got %v, want %v", got, want)
	}
	expected := access.Permissions{
		"Admin": access.AccessList{
			In:    []security.BlessingPattern{"root/$", "root/self/$", "root/self/child"},
			NotIn: []string(nil)},
		"Read": access.AccessList{
			In:    []security.BlessingPattern{"root/$", "root/self/$", "root/self/child"},
			NotIn: []string(nil)},
		"Write": access.AccessList{
			In:    []security.BlessingPattern{"root/$", "root/self/$", "root/self/child"},
			NotIn: []string(nil)},
		"Debug": access.AccessList{
			In:    []security.BlessingPattern{"root/$", "root/self/$", "root/self/child"},
			NotIn: []string(nil)},
		"Resolve": access.AccessList{
			In:    []security.BlessingPattern{"root/$", "root/self/$", "root/self/child"},
			NotIn: []string(nil)}}
	if got := acl; !reflect.DeepEqual(expected.Normalize(), got.Normalize()) {
		t.Errorf("got %#v, exected %#v ", got, expected)
	}

	vlog.VI(2).Infof("self attempting to give other permission to update application")
	newAccessList := make(access.Permissions)
	for _, tag := range access.AllTypicalTags() {
		newAccessList.Add("root/self", string(tag))
		newAccessList.Add("root/other", string(tag))
	}
	if err := repostub.SetPermissions(ctx, newAccessList, ""); err != nil {
		t.Fatalf("SetPermissions failed: %v", err)
	}

	acl, version, err = repostub.GetPermissions(ctx)
	if err != nil {
		t.Fatalf("GetPermissions should not have failed: %v", err)
	}
	expected = newAccessList
	if got := acl; !reflect.DeepEqual(expected.Normalize(), got.Normalize()) {
		t.Errorf("got %#v, exected %#v ", got, expected)
	}

	// Envelope putting as other should now succeed.
	if err := v1stub.Put(otherCtx, []string{"base"}, envelopeV1); err != nil {
		t.Fatalf("Put() wrongly failed: %v", err)
	}

	// Other takes control.
	acl, version, err = repostub.GetPermissions(otherCtx)
	if err != nil {
		t.Fatalf("GetPermissions 2 should not have failed: %v", err)
	}
	acl["Admin"] = access.AccessList{
		In:    []security.BlessingPattern{"root/other"},
		NotIn: []string{}}
	if err = repostub.SetPermissions(otherCtx, acl, version); err != nil {
		t.Fatalf("SetPermissions failed: %v", err)
	}

	// Self is now locked out but other isn't.
	if _, _, err = repostub.GetPermissions(ctx); err == nil {
		t.Fatalf("GetPermissions should not have succeeded")
	}
	acl, _, err = repostub.GetPermissions(otherCtx)
	if err != nil {
		t.Fatalf("GetPermissions should not have failed: %v", err)
	}
	expected = access.Permissions{
		"Admin": access.AccessList{
			In:    []security.BlessingPattern{"root/other"},
			NotIn: []string{}},
		"Read": access.AccessList{In: []security.BlessingPattern{"root/other",
			"root/self"},
			NotIn: []string{}},
		"Write": access.AccessList{In: []security.BlessingPattern{"root/other",
			"root/self"},
			NotIn: []string{}},
		"Debug": access.AccessList{In: []security.BlessingPattern{"root/other",
			"root/self"},
			NotIn: []string{}},
		"Resolve": access.AccessList{In: []security.BlessingPattern{"root/other",
			"root/self"},
			NotIn: []string{}}}

	if got := acl; !reflect.DeepEqual(expected.Normalize(), got.Normalize()) {
		t.Errorf("got %#v, exected %#v ", got, expected)
	}
}

func TestPerAppAccessList(t *testing.T) {
	ctx, shutdown := test.InitForTest()
	defer shutdown()
	v23.GetNamespace(ctx).CacheCtl(naming.DisableCache(true))
	// By default, all principals in this test will have blessings
	// generated based on the username/machine running this process. Give
	// them recognizable names ("root/self" etc.), so the AccessLists can be set
	// deterministically.
	idp := testutil.NewIDProvider("root")
	if err := idp.Bless(v23.GetPrincipal(ctx), "self"); err != nil {
		t.Fatal(err)
	}

	sh, deferFn := servicetest.CreateShellAndMountTable(t, ctx, v23.GetPrincipal(ctx))
	defer deferFn()

	// setup mock up directory to put state in
	storedir, cleanup := servicetest.SetupRootDir(t, "application")
	defer cleanup()

	otherCtx, err := v23.SetPrincipal(ctx, testutil.NewPrincipal())
	if err != nil {
		t.Fatal(err)
	}
	if err := idp.Bless(v23.GetPrincipal(otherCtx), "other"); err != nil {
		t.Fatal(err)
	}

	nmh := servicetest.RunCommand(t, sh, nil, repoCmd, "repo", storedir)
	pid := servicetest.ReadPID(t, nmh)
	defer syscall.Kill(pid, syscall.SIGINT)

	// Create example envelope.
	envelopeV1 := application.Envelope{
		Args:   []string{"--help"},
		Env:    []string{"DEBUG=1"},
		Binary: application.SignedFile{File: "/v23/name/of/binary"},
	}

	vlog.VI(2).Info("Upload an envelope")
	v1stub := repository.ApplicationClient("repo/search/v1")
	if err := v1stub.Put(ctx, []string{"base"}, envelopeV1); err != nil {
		t.Fatalf("Put() failed: %v", err)
	}
	v2stub := repository.ApplicationClient("repo/search/v2")
	if err := v2stub.Put(ctx, []string{"base"}, envelopeV1); err != nil {
		t.Fatalf("Put() failed: %v", err)
	}
	v3stub := repository.ApplicationClient("repo/naps/v1")
	if err := v3stub.Put(ctx, []string{"base"}, envelopeV1); err != nil {
		t.Fatalf("Put() failed: %v", err)
	}

	vlog.VI(2).Info("Self can access.AccessLists but other can't.")
	expectedSelfPermissions := access.Permissions{
		"Admin": access.AccessList{
			In:    []security.BlessingPattern{"root/$", "root/self"},
			NotIn: []string{}},
		"Read": access.AccessList{In: []security.BlessingPattern{"root/$", "root/self"},
			NotIn: []string{}},
		"Write": access.AccessList{In: []security.BlessingPattern{"root/$", "root/self"},
			NotIn: []string{}},
		"Debug": access.AccessList{In: []security.BlessingPattern{"root/$", "root/self"},
			NotIn: []string{}},
		"Resolve": access.AccessList{In: []security.BlessingPattern{"root/$", "root/self"},
			NotIn: []string{}}}

	for _, path := range []string{"repo/search", "repo/search/v1", "repo/search/v2", "repo/naps", "repo/naps/v1"} {
		stub := repository.ApplicationClient(path)
		acl, _, err := stub.GetPermissions(ctx)
		if err != nil {
			t.Fatalf("Newly uploaded envelopes failed to receive permission lists: %v", err)
		}

		if got := acl; !reflect.DeepEqual(expectedSelfPermissions.Normalize(), got.Normalize()) {
			t.Errorf("got %#v, expected %#v ", got, expectedSelfPermissions)
		}

		// But otherCtx doesn't have admin permissions so has no access.
		if _, _, err := stub.GetPermissions(otherCtx); err == nil {
			t.Fatalf("GetPermissions didn't fail for other when it should have.")
		}
	}

	vlog.VI(2).Infof("Self sets root AccessLists.")
	repostub := repository.ApplicationClient("repo")
	newAccessList := make(access.Permissions)
	for _, tag := range access.AllTypicalTags() {
		newAccessList.Add("root/self", string(tag))
	}
	if err := repostub.SetPermissions(ctx, newAccessList, ""); err != nil {
		t.Fatalf("SetPermissions failed: %v", err)
	}

	vlog.VI(2).Infof("Other still can't access anything.")
	if _, _, err = repostub.GetPermissions(otherCtx); err == nil {
		t.Fatalf("GetPermissions should have failed")
	}

	vlog.VI(2).Infof("Self gives other full access to repo/search/...")
	newAccessList, version, err := v1stub.GetPermissions(ctx)
	if err != nil {
		t.Fatalf("GetPermissions should not have failed: %v", err)
	}
	for _, tag := range access.AllTypicalTags() {
		newAccessList.Add("root/other", string(tag))
	}
	if err := v1stub.SetPermissions(ctx, newAccessList, version); err != nil {
		t.Fatalf("SetPermissions failed: %v", err)
	}

	expected := access.Permissions{
		"Resolve": access.AccessList{In: []security.BlessingPattern{
			"root/$",
			"root/other",
			"root/self"},
			NotIn: []string(nil)},
		"Admin": access.AccessList{In: []security.BlessingPattern{
			"root/$",
			"root/other",
			"root/self"},
			NotIn: []string(nil)},
		"Read": access.AccessList{In: []security.BlessingPattern{
			"root/$",
			"root/other",
			"root/self"},
			NotIn: []string(nil)},
		"Write": access.AccessList{In: []security.BlessingPattern{
			"root/$",
			"root/other",
			"root/self"},
			NotIn: []string(nil)},
		"Debug": access.AccessList{In: []security.BlessingPattern{
			"root/$",
			"root/other", "root/self"},
			NotIn: []string(nil)},
	}

	for _, path := range []string{"repo/search", "repo/search/v1", "repo/search/v2"} {
		stub := repository.ApplicationClient(path)
		vlog.VI(2).Infof("Other can now access this app independent of version.")
		acl, _, err := stub.GetPermissions(otherCtx)
		if err != nil {
			t.Fatalf("GetPermissions should not have failed: %v", err)
		}

		if got := acl; !reflect.DeepEqual(expected.Normalize(), got.Normalize()) {
			t.Errorf("got %#v, expected %#v ", got, expected)
		}
		vlog.VI(2).Infof("Self can also access thanks to hierarchical auth.")
		if _, _, err = stub.GetPermissions(ctx); err != nil {
			t.Fatalf("GetPermissions should not have failed: %v", err)
		}
	}

	vlog.VI(2).Infof("But other locations are unaffected and other cannot access.")
	for _, path := range []string{"repo/naps", "repo/naps/v1"} {
		stub := repository.ApplicationClient(path)
		if _, _, err := stub.GetPermissions(otherCtx); err == nil {
			t.Fatalf("GetPermissions didn't fail when it should have.")
		}
	}

	// Self gives other write perms on base.
	newAccessList, version, err = repostub.GetPermissions(ctx)
	if err != nil {
		t.Fatalf("GetPermissions should not have failed: %v", err)
	}
	newAccessList["Write"] = access.AccessList{In: []security.BlessingPattern{"root/other", "root/self"}}
	if err := repostub.SetPermissions(ctx, newAccessList, version); err != nil {
		t.Fatalf("SetPermissions failed: %v", err)
	}

	// Other can now upload an envelope at both locations.
	for _, stub := range []repository.ApplicationClientStub{v1stub, v2stub} {
		if err := stub.Put(otherCtx, []string{"base"}, envelopeV1); err != nil {
			t.Fatalf("Put() failed: %v", err)
		}
	}

	// But because application search already exists, the ACLs do not change.
	for _, path := range []string{"repo/search", "repo/search/v1", "repo/search/v2"} {
		stub := repository.ApplicationClient(path)
		acl, _, err := stub.GetPermissions(otherCtx)
		if err != nil {
			t.Fatalf("GetPermissions should not have failed: %v", err)
		}
		if got := acl; !reflect.DeepEqual(expected.Normalize(), got.Normalize()) {
			t.Errorf("got %#v, expected %#v ", got, expected)
		}
	}

	// But self didn't give other AccessList modification permissions.
	for _, path := range []string{"repo/search", "repo/search/v2"} {
		stub := repository.ApplicationClient(path)
		if _, _, err := stub.GetPermissions(otherCtx); err != nil {
			t.Fatalf("GetPermissions failed when it should not have for same application: %v", err)
		}
	}
}