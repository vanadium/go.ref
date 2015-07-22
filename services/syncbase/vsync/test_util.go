// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vsync

// Utilities for testing sync.

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	wire "v.io/syncbase/v23/services/syncbase/nosql"
	"v.io/syncbase/x/ref/services/syncbase/localblobstore"
	"v.io/syncbase/x/ref/services/syncbase/localblobstore/fs_cablobstore"
	"v.io/syncbase/x/ref/services/syncbase/server/interfaces"
	"v.io/syncbase/x/ref/services/syncbase/server/util"
	"v.io/syncbase/x/ref/services/syncbase/server/watchable"
	"v.io/syncbase/x/ref/services/syncbase/store"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security/access"
	"v.io/v23/verror"
	"v.io/x/ref/test"
)

// mockService emulates a Syncbase service that includes store and sync.
// It is used to access a mock application.
type mockService struct {
	engine   string
	dir      string
	st       store.Store
	bst      localblobstore.BlobStore
	sync     *syncService
	shutdown func()
}

func (s *mockService) St() store.Store {
	return s.st
}

func (s *mockService) Sync() interfaces.SyncServerMethods {
	return s.sync
}

func (s *mockService) App(ctx *context.T, call rpc.ServerCall, appName string) (interfaces.App, error) {
	return &mockApp{st: s.st, bst: s.bst}, nil
}

func (s *mockService) AppNames(ctx *context.T, call rpc.ServerCall) ([]string, error) {
	return []string{"mockapp"}, nil
}

// mockApp emulates a Syncbase App.  It is used to access a mock database.
type mockApp struct {
	st  store.Store
	bst localblobstore.BlobStore
}

func (a *mockApp) NoSQLDatabase(ctx *context.T, call rpc.ServerCall, dbName string) (interfaces.Database, error) {
	return &mockDatabase{st: a.st, bst: a.bst}, nil
}

func (a *mockApp) NoSQLDatabaseNames(ctx *context.T, call rpc.ServerCall) ([]string, error) {
	return []string{"mockdb"}, nil
}

func (a *mockApp) CreateNoSQLDatabase(ctx *context.T, call rpc.ServerCall, dbName string, perms access.Permissions, metadata *wire.SchemaMetadata) error {
	return verror.NewErrNotImplemented(ctx)
}

func (a *mockApp) DeleteNoSQLDatabase(ctx *context.T, call rpc.ServerCall, dbName string) error {
	return verror.NewErrNotImplemented(ctx)
}

func (a *mockApp) SetDatabasePerms(ctx *context.T, call rpc.ServerCall, dbName string, perms access.Permissions, version string) error {
	return verror.NewErrNotImplemented(ctx)
}

func (a *mockApp) Service() interfaces.Service {
	return nil
}

func (a *mockApp) Name() string {
	return "mockapp"
}

func (a *mockApp) StKey() string {
	return ""
}

// mockDatabase emulates a Syncbase Database.  It is used to test sync functionality.
type mockDatabase struct {
	st  store.Store
	bst localblobstore.BlobStore
}

func (d *mockDatabase) St() store.Store {
	return d.st
}

func (d *mockDatabase) BlobSt() localblobstore.BlobStore {
	return d.bst
}

func (d *mockDatabase) CheckPermsInternal(ctx *context.T, call rpc.ServerCall, st store.StoreReader) error {
	return verror.NewErrNotImplemented(ctx)
}

func (d *mockDatabase) SetPermsInternal(ctx *context.T, call rpc.ServerCall, perms access.Permissions, version string) error {
	return verror.NewErrNotImplemented(ctx)
}

func (d *mockDatabase) Name() string {
	return "mockdb"
}

func (d *mockDatabase) StKey() string {
	return ""
}

func (d *mockDatabase) App() interfaces.App {
	return nil
}

// createService creates a mock Syncbase service used for testing sync functionality.
func createService(t *testing.T) *mockService {
	ctx, shutdown := test.V23Init()
	engine := "leveldb"
	opts := util.OpenOptions{CreateIfMissing: true, ErrorIfExists: false}
	dir := fmt.Sprintf("%s/vsync_test_%d_%d", os.TempDir(), os.Getpid(), time.Now().UnixNano())

	st, err := util.OpenStore(engine, path.Join(dir, engine), opts)
	if err != nil {
		t.Fatalf("cannot create store %s (%s): %v", engine, dir, err)
	}
	st, err = watchable.Wrap(st, &watchable.Options{
		ManagedPrefixes: []string{util.RowPrefix, util.PermsPrefix},
	})
	bst, err := fs_cablobstore.Create(ctx, path.Join(dir, "blobs"))
	if err != nil {
		t.Fatalf("cannot create blob store (%s): %v", dir, err)
	}

	s := &mockService{
		st:       st,
		bst:      bst,
		engine:   engine,
		dir:      dir,
		shutdown: shutdown,
	}
	if s.sync, err = New(ctx, nil, s, nil); err != nil {
		util.DestroyStore(engine, dir)
		t.Fatalf("cannot create sync service: %v", err)
	}
	return s
}

// destroyService cleans up the mock Syncbase service.
func destroyService(t *testing.T, s *mockService) {
	defer s.shutdown()
	defer s.sync.Close()
	if err := util.DestroyStore(s.engine, s.dir); err != nil {
		t.Fatalf("cannot destroy store %s (%s): %v", s.engine, s.dir, err)
	}
}

// makeResMark returns the resume marker for a given log entry position.
func makeResMark(pos int) string {
	return util.JoinKeyParts(util.LogPrefix, fmt.Sprintf("%016x", pos))
}

// makeRowKey returns the database row key for a given application key.
func makeRowKey(key string) string {
	return util.JoinKeyParts(util.RowPrefix, key)
}
