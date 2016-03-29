// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: vsync

package vsync

import (
	"fmt"
	"v.io/v23/vdl"
	"v.io/x/ref/services/syncbase/server/interfaces"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// SyncData represents the persistent state of the sync module.
type SyncData struct {
	Id uint64
}

func (SyncData) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync.SyncData"`
}) {
}

func (m *SyncData) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Id")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromUint(uint64(m.Id), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *SyncData) MakeVDLTarget() vdl.Target {
	return &SyncDataTarget{Value: m}
}

type SyncDataTarget struct {
	Value    *SyncData
	idTarget vdl.Uint64Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *SyncDataTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*SyncData)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *SyncDataTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Id":
		t.idTarget.Value = &t.Value.Id
		target, err := &t.idTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vsync.SyncData", name)
	}
}
func (t *SyncDataTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *SyncDataTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// DbSyncState represents the persistent sync state of a Database.
type DbSyncState struct {
	GenVecs   interfaces.Knowledge // knowledge capturing the locally-known generations of remote peers for data in Database.
	SgGenVecs interfaces.Knowledge // knowledge capturing the locally-known generations of remote peers for syncgroups in Database.
	IsPaused  bool                 // tracks whether sync is paused by client.
}

func (DbSyncState) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync.DbSyncState"`
}) {
}

func (m *DbSyncState) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("GenVecs")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.GenVecs.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("SgGenVecs")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.SgGenVecs.FillVDLTarget(fieldTarget5, tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("IsPaused")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget7.FromBool(bool(m.IsPaused), tt.NonOptional().Field(2).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *DbSyncState) MakeVDLTarget() vdl.Target {
	return &DbSyncStateTarget{Value: m}
}

type DbSyncStateTarget struct {
	Value           *DbSyncState
	genVecsTarget   interfaces.KnowledgeTarget
	sgGenVecsTarget interfaces.KnowledgeTarget
	isPausedTarget  vdl.BoolTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *DbSyncStateTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*DbSyncState)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *DbSyncStateTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "GenVecs":
		t.genVecsTarget.Value = &t.Value.GenVecs
		target, err := &t.genVecsTarget, error(nil)
		return nil, target, err
	case "SgGenVecs":
		t.sgGenVecsTarget.Value = &t.Value.SgGenVecs
		target, err := &t.sgGenVecsTarget, error(nil)
		return nil, target, err
	case "IsPaused":
		t.isPausedTarget.Value = &t.Value.IsPaused
		target, err := &t.isPausedTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vsync.DbSyncState", name)
	}
}
func (t *DbSyncStateTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *DbSyncStateTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// LocalLogRec represents the persistent local state of a log record. Metadata
// is synced across peers, while pos is local-only.
type LocalLogRec struct {
	Metadata interfaces.LogRecMetadata
	Pos      uint64 // position in the Database log.
	Shell    bool   // indicates if the value was shelled by the sender.
}

func (LocalLogRec) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync.LocalLogRec"`
}) {
}

func (m *LocalLogRec) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Metadata")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Metadata.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Pos")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromUint(uint64(m.Pos), tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Shell")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget7.FromBool(bool(m.Shell), tt.NonOptional().Field(2).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *LocalLogRec) MakeVDLTarget() vdl.Target {
	return &LocalLogRecTarget{Value: m}
}

type LocalLogRecTarget struct {
	Value          *LocalLogRec
	metadataTarget interfaces.LogRecMetadataTarget
	posTarget      vdl.Uint64Target
	shellTarget    vdl.BoolTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *LocalLogRecTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*LocalLogRec)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *LocalLogRecTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Metadata":
		t.metadataTarget.Value = &t.Value.Metadata
		target, err := &t.metadataTarget, error(nil)
		return nil, target, err
	case "Pos":
		t.posTarget.Value = &t.Value.Pos
		target, err := &t.posTarget, error(nil)
		return nil, target, err
	case "Shell":
		t.shellTarget.Value = &t.Value.Shell
		target, err := &t.shellTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vsync.LocalLogRec", name)
	}
}
func (t *LocalLogRecTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *LocalLogRecTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// SgLocalState holds the syncgroup local state, only relevant to this member
// (i.e. the local Syncbase).  This is needed for crash recovery of the internal
// state transitions of the syncgroup.
type SgLocalState struct {
	// The count of local joiners to the same syncgroup.
	NumLocalJoiners uint32
	// The syncgroup is watched when the sync Watcher starts processing the
	// syncgroup data.  When a syncgroup is created or joined, an entry is
	// added to the Watcher queue (log) to inform it from which point to
	// start accepting store mutations, an asynchronous notification similar
	// to regular store mutations.  When the Watcher processes that queue
	// entry, it sets this bit to true.  When Syncbase restarts, the value
	// of this bit allows the new sync Watcher to recreate its in-memory
	// state by resuming to watch only the prefixes of syncgroups that were
	// previously being watched.
	Watched bool
	// The syncgroup was published here by this remote peer (if non-empty
	// string), typically the syncgroup creator.  In this case the syncgroup
	// cannot be GCed locally even if it has no local joiners.
	RemotePublisher string
	// The syncgroup is in pending state on a device that learns the current
	// state of the syncgroup from another device but has not yet received
	// through peer-to-peer sync the history of the changes (DAG and logs).
	// This happens in two cases:
	// 1- A joiner was accepted into a syncgroup by a syncgroup admin and
	//    only given the current syncgroup info synchronously and will
	//    receive the full history later via p2p sync.
	// 2- A remote server where the syncgroup is published was told by the
	//    syncgroup publisher the current syncgroup info synchronously and
	//    will receive the full history later via p2p sync.
	// The pending state is over when the device reaches or exceeds the
	// knowledge level indicated in the pending genvec.  While SyncPending
	// is true, no local syncgroup mutations are allowed (i.e. no join or
	// set-spec requests).
	SyncPending   bool
	PendingGenVec interfaces.GenVector
}

func (SgLocalState) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync.SgLocalState"`
}) {
}

func (m *SgLocalState) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("NumLocalJoiners")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromUint(uint64(m.NumLocalJoiners), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Watched")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromBool(bool(m.Watched), tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("RemotePublisher")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget7.FromString(string(m.RemotePublisher), tt.NonOptional().Field(2).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("SyncPending")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget9.FromBool(bool(m.SyncPending), tt.NonOptional().Field(3).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
			return err
		}
	}
	keyTarget10, fieldTarget11, err := fieldsTarget1.StartField("PendingGenVec")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.PendingGenVec.FillVDLTarget(fieldTarget11, tt.NonOptional().Field(4).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget10, fieldTarget11); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *SgLocalState) MakeVDLTarget() vdl.Target {
	return &SgLocalStateTarget{Value: m}
}

type SgLocalStateTarget struct {
	Value                 *SgLocalState
	numLocalJoinersTarget vdl.Uint32Target
	watchedTarget         vdl.BoolTarget
	remotePublisherTarget vdl.StringTarget
	syncPendingTarget     vdl.BoolTarget
	pendingGenVecTarget   interfaces.GenVectorTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *SgLocalStateTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*SgLocalState)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *SgLocalStateTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "NumLocalJoiners":
		t.numLocalJoinersTarget.Value = &t.Value.NumLocalJoiners
		target, err := &t.numLocalJoinersTarget, error(nil)
		return nil, target, err
	case "Watched":
		t.watchedTarget.Value = &t.Value.Watched
		target, err := &t.watchedTarget, error(nil)
		return nil, target, err
	case "RemotePublisher":
		t.remotePublisherTarget.Value = &t.Value.RemotePublisher
		target, err := &t.remotePublisherTarget, error(nil)
		return nil, target, err
	case "SyncPending":
		t.syncPendingTarget.Value = &t.Value.SyncPending
		target, err := &t.syncPendingTarget, error(nil)
		return nil, target, err
	case "PendingGenVec":
		t.pendingGenVecTarget.Value = &t.Value.PendingGenVec
		target, err := &t.pendingGenVecTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vsync.SgLocalState", name)
	}
}
func (t *SgLocalStateTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *SgLocalStateTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// DagNode holds the information on an object mutation in the DAG.  The node
// information is extracted from the log records exchanged between Syncbases.
// They are also stored in the DAG node to improve DAG traversal for conflict
// resolution and pruning without having to fetch the full log record.
type DagNode struct {
	Level    uint64   // node distance from root
	Parents  []string // references to parent versions
	Logrec   string   // reference to log record
	BatchId  uint64   // ID of a write batch
	Shell    bool     // true when the data is hidden due to permissions
	Deleted  bool     // true if the change was a delete
	PermId   string   // ID of the permissions controlling this version
	PermVers string   // current version of the permissions object
}

func (DagNode) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync.DagNode"`
}) {
}

func (m *DagNode) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Level")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromUint(uint64(m.Level), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Parents")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget6, err := fieldTarget5.StartList(tt.NonOptional().Field(1).Type, len(m.Parents))
		if err != nil {
			return err
		}
		for i, elem8 := range m.Parents {
			elemTarget7, err := listTarget6.StartElem(i)
			if err != nil {
				return err
			}
			if err := elemTarget7.FromString(string(elem8), tt.NonOptional().Field(1).Type.Elem()); err != nil {
				return err
			}
			if err := listTarget6.FinishElem(elemTarget7); err != nil {
				return err
			}
		}
		if err := fieldTarget5.FinishList(listTarget6); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget9, fieldTarget10, err := fieldsTarget1.StartField("Logrec")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget10.FromString(string(m.Logrec), tt.NonOptional().Field(2).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget9, fieldTarget10); err != nil {
			return err
		}
	}
	keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("BatchId")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget12.FromUint(uint64(m.BatchId), tt.NonOptional().Field(3).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
			return err
		}
	}
	keyTarget13, fieldTarget14, err := fieldsTarget1.StartField("Shell")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget14.FromBool(bool(m.Shell), tt.NonOptional().Field(4).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget13, fieldTarget14); err != nil {
			return err
		}
	}
	keyTarget15, fieldTarget16, err := fieldsTarget1.StartField("Deleted")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget16.FromBool(bool(m.Deleted), tt.NonOptional().Field(5).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget15, fieldTarget16); err != nil {
			return err
		}
	}
	keyTarget17, fieldTarget18, err := fieldsTarget1.StartField("PermId")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget18.FromString(string(m.PermId), tt.NonOptional().Field(6).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget17, fieldTarget18); err != nil {
			return err
		}
	}
	keyTarget19, fieldTarget20, err := fieldsTarget1.StartField("PermVers")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget20.FromString(string(m.PermVers), tt.NonOptional().Field(7).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget19, fieldTarget20); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *DagNode) MakeVDLTarget() vdl.Target {
	return &DagNodeTarget{Value: m}
}

type DagNodeTarget struct {
	Value          *DagNode
	levelTarget    vdl.Uint64Target
	parentsTarget  vdl.StringSliceTarget
	logrecTarget   vdl.StringTarget
	batchIdTarget  vdl.Uint64Target
	shellTarget    vdl.BoolTarget
	deletedTarget  vdl.BoolTarget
	permIdTarget   vdl.StringTarget
	permVersTarget vdl.StringTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *DagNodeTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*DagNode)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *DagNodeTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Level":
		t.levelTarget.Value = &t.Value.Level
		target, err := &t.levelTarget, error(nil)
		return nil, target, err
	case "Parents":
		t.parentsTarget.Value = &t.Value.Parents
		target, err := &t.parentsTarget, error(nil)
		return nil, target, err
	case "Logrec":
		t.logrecTarget.Value = &t.Value.Logrec
		target, err := &t.logrecTarget, error(nil)
		return nil, target, err
	case "BatchId":
		t.batchIdTarget.Value = &t.Value.BatchId
		target, err := &t.batchIdTarget, error(nil)
		return nil, target, err
	case "Shell":
		t.shellTarget.Value = &t.Value.Shell
		target, err := &t.shellTarget, error(nil)
		return nil, target, err
	case "Deleted":
		t.deletedTarget.Value = &t.Value.Deleted
		target, err := &t.deletedTarget, error(nil)
		return nil, target, err
	case "PermId":
		t.permIdTarget.Value = &t.Value.PermId
		target, err := &t.permIdTarget, error(nil)
		return nil, target, err
	case "PermVers":
		t.permVersTarget.Value = &t.Value.PermVers
		target, err := &t.permVersTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vsync.DagNode", name)
	}
}
func (t *DagNodeTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *DagNodeTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// BatchInfo holds the information on a write batch:
//  * The map of syncable (versioned) objects: {oid: version}
//  * The map of linked objects {oid: version} that were not explicitly written
//    as part of the batch but were reaffirmed during conflict resolution along
//    with other objects written in this batch by the app by choosing
//    "pickLocal" or "pickRemote". NOTE: this map is non empty only for batches
//    created during conflict resolution. Unlike the Objects map, the collection
//    of oid:version present in this map do not point back to this batch. They
//    point to the batches that there were originally created in.
//  * The total count of batch objects, including non-syncable ones.
// TODO(rdaoud): add support to track the read and scan sets.
type BatchInfo struct {
	Objects       map[string]string
	LinkedObjects map[string]string
	Count         uint64
}

func (BatchInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync.BatchInfo"`
}) {
}

func (m *BatchInfo) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Objects")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		mapTarget4, err := fieldTarget3.StartMap(tt.NonOptional().Field(0).Type, len(m.Objects))
		if err != nil {
			return err
		}
		for key6, value8 := range m.Objects {
			keyTarget5, err := mapTarget4.StartKey()
			if err != nil {
				return err
			}
			if err := keyTarget5.FromString(string(key6), tt.NonOptional().Field(0).Type.Key()); err != nil {
				return err
			}
			valueTarget7, err := mapTarget4.FinishKeyStartField(keyTarget5)
			if err != nil {
				return err
			}
			if err := valueTarget7.FromString(string(value8), tt.NonOptional().Field(0).Type.Elem()); err != nil {
				return err
			}
			if err := mapTarget4.FinishField(keyTarget5, valueTarget7); err != nil {
				return err
			}
		}
		if err := fieldTarget3.FinishMap(mapTarget4); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget9, fieldTarget10, err := fieldsTarget1.StartField("LinkedObjects")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		mapTarget11, err := fieldTarget10.StartMap(tt.NonOptional().Field(1).Type, len(m.LinkedObjects))
		if err != nil {
			return err
		}
		for key13, value15 := range m.LinkedObjects {
			keyTarget12, err := mapTarget11.StartKey()
			if err != nil {
				return err
			}
			if err := keyTarget12.FromString(string(key13), tt.NonOptional().Field(1).Type.Key()); err != nil {
				return err
			}
			valueTarget14, err := mapTarget11.FinishKeyStartField(keyTarget12)
			if err != nil {
				return err
			}
			if err := valueTarget14.FromString(string(value15), tt.NonOptional().Field(1).Type.Elem()); err != nil {
				return err
			}
			if err := mapTarget11.FinishField(keyTarget12, valueTarget14); err != nil {
				return err
			}
		}
		if err := fieldTarget10.FinishMap(mapTarget11); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget9, fieldTarget10); err != nil {
			return err
		}
	}
	keyTarget16, fieldTarget17, err := fieldsTarget1.StartField("Count")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget17.FromUint(uint64(m.Count), tt.NonOptional().Field(2).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget16, fieldTarget17); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *BatchInfo) MakeVDLTarget() vdl.Target {
	return &BatchInfoTarget{Value: m}
}

type BatchInfoTarget struct {
	Value               *BatchInfo
	objectsTarget       __VDLTarget1_map
	linkedObjectsTarget __VDLTarget1_map
	countTarget         vdl.Uint64Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BatchInfoTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*BatchInfo)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *BatchInfoTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Objects":
		t.objectsTarget.Value = &t.Value.Objects
		target, err := &t.objectsTarget, error(nil)
		return nil, target, err
	case "LinkedObjects":
		t.linkedObjectsTarget.Value = &t.Value.LinkedObjects
		target, err := &t.linkedObjectsTarget, error(nil)
		return nil, target, err
	case "Count":
		t.countTarget.Value = &t.Value.Count
		target, err := &t.countTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/vsync.BatchInfo", name)
	}
}
func (t *BatchInfoTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BatchInfoTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[string]string
type __VDLTarget1_map struct {
	Value      *map[string]string
	currKey    string
	currElem   string
	keyTarget  vdl.StringTarget
	elemTarget vdl.StringTarget
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *__VDLTarget1_map) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {

	if ttWant := vdl.TypeOf((*map[string]string)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(map[string]string)
	return t, nil
}
func (t *__VDLTarget1_map) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *__VDLTarget1_map) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = ""
	t.elemTarget.Value = &t.currElem
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *__VDLTarget1_map) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *__VDLTarget1_map) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

// Create zero values for each type.
var (
	__VDLZeroSyncData     = SyncData{}
	__VDLZeroDbSyncState  = DbSyncState{}
	__VDLZeroLocalLogRec  = LocalLogRec{}
	__VDLZeroSgLocalState = SgLocalState{}
	__VDLZeroDagNode      = DagNode{}
	__VDLZeroBatchInfo    = BatchInfo{}
)

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}

	// Register types.
	vdl.Register((*SyncData)(nil))
	vdl.Register((*DbSyncState)(nil))
	vdl.Register((*LocalLogRec)(nil))
	vdl.Register((*SgLocalState)(nil))
	vdl.Register((*DagNode)(nil))
	vdl.Register((*BatchInfo)(nil))

	return struct{}{}
}