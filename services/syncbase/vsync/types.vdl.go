// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package vsync

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/syncbase/x/ref/services/syncbase/server/interfaces"
)

// syncData represents the persistent state of the sync module.
type syncData struct {
	Id uint64
}

func (syncData) __VDLReflect(struct {
	Name string `vdl:"v.io/syncbase/x/ref/services/syncbase/vsync.syncData"`
}) {
}

// dbSyncState represents the persistent sync state of a Database.
type dbSyncState struct {
	Gen        uint64               // local generation number incremented on every local update.
	CheckptGen uint64               // local generation number advertised to remote peers (used by the responder).
	GenVec     interfaces.GenVector // generation vector capturing the locally-known generations of remote peers.
}

func (dbSyncState) __VDLReflect(struct {
	Name string `vdl:"v.io/syncbase/x/ref/services/syncbase/vsync.dbSyncState"`
}) {
}

// localLogRec represents the persistent local state of a log record. Metadata
// is synced across peers, while pos is local-only.
type localLogRec struct {
	Metadata interfaces.LogRecMetadata
	Pos      uint64 // position in the Database log.
}

func (localLogRec) __VDLReflect(struct {
	Name string `vdl:"v.io/syncbase/x/ref/services/syncbase/vsync.localLogRec"`
}) {
}

func init() {
	vdl.Register((*syncData)(nil))
	vdl.Register((*dbSyncState)(nil))
	vdl.Register((*localLogRec)(nil))
}

const logPrefix = "log"

const dbssPrefix = "dbss"

const dagPrefix = "dag"

const sgPrefix = "sg"
