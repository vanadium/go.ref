// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: watchable

package watchable

import (
	"fmt"
	"reflect"
	"v.io/v23/vdl"
	"v.io/v23/vom"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// GetOp represents a store get operation.
type GetOp struct {
	Key []byte
}

func (GetOp) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/store/watchable.GetOp"`
}) {
}

func (m *GetOp) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Key")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget3.FromBytes([]byte(m.Key), tt.NonOptional().Field(0).Type); err != nil {
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

func (m *GetOp) MakeVDLTarget() vdl.Target {
	return &GetOpTarget{Value: m}
}

type GetOpTarget struct {
	Value     *GetOp
	keyTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *GetOpTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*GetOp)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *GetOpTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Key":
		t.keyTarget.Value = &t.Value.Key
		target, err := &t.keyTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/store/watchable.GetOp", name)
	}
}
func (t *GetOpTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *GetOpTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// ScanOp represents a store scan operation.
type ScanOp struct {
	Start []byte
	Limit []byte
}

func (ScanOp) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/store/watchable.ScanOp"`
}) {
}

func (m *ScanOp) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Start")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget3.FromBytes([]byte(m.Start), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Limit")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget5.FromBytes([]byte(m.Limit), tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *ScanOp) MakeVDLTarget() vdl.Target {
	return &ScanOpTarget{Value: m}
}

type ScanOpTarget struct {
	Value       *ScanOp
	startTarget vdl.BytesTarget
	limitTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *ScanOpTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*ScanOp)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *ScanOpTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Start":
		t.startTarget.Value = &t.Value.Start
		target, err := &t.startTarget, error(nil)
		return nil, target, err
	case "Limit":
		t.limitTarget.Value = &t.Value.Limit
		target, err := &t.limitTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/store/watchable.ScanOp", name)
	}
}
func (t *ScanOpTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *ScanOpTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// PutOp represents a store put operation.  The new version is written instead
// of the value to avoid duplicating the user data in the store.  The version
// is used to access the user data of that specific mutation.  The key and the
// version of the permissions entry that was checked to allow this put operation
// are also tracked to secure the access to this history.
type PutOp struct {
	Key         []byte
	Version     []byte
	PermKey     []byte
	PermVersion []byte
}

func (PutOp) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/store/watchable.PutOp"`
}) {
}

func (m *PutOp) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Key")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget3.FromBytes([]byte(m.Key), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Version")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget5.FromBytes([]byte(m.Version), tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("PermKey")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget7.FromBytes([]byte(m.PermKey), tt.NonOptional().Field(2).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("PermVersion")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget9.FromBytes([]byte(m.PermVersion), tt.NonOptional().Field(3).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *PutOp) MakeVDLTarget() vdl.Target {
	return &PutOpTarget{Value: m}
}

type PutOpTarget struct {
	Value             *PutOp
	keyTarget         vdl.BytesTarget
	versionTarget     vdl.BytesTarget
	permKeyTarget     vdl.BytesTarget
	permVersionTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *PutOpTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*PutOp)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *PutOpTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Key":
		t.keyTarget.Value = &t.Value.Key
		target, err := &t.keyTarget, error(nil)
		return nil, target, err
	case "Version":
		t.versionTarget.Value = &t.Value.Version
		target, err := &t.versionTarget, error(nil)
		return nil, target, err
	case "PermKey":
		t.permKeyTarget.Value = &t.Value.PermKey
		target, err := &t.permKeyTarget, error(nil)
		return nil, target, err
	case "PermVersion":
		t.permVersionTarget.Value = &t.Value.PermVersion
		target, err := &t.permVersionTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/store/watchable.PutOp", name)
	}
}
func (t *PutOpTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *PutOpTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// DeleteOp represents a store delete operation.  The key and the version of the
// permissions entry that was checked to allow this delete operation are also
// tracked to secure the access to this history.
type DeleteOp struct {
	Key         []byte
	PermKey     []byte
	PermVersion []byte
}

func (DeleteOp) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/store/watchable.DeleteOp"`
}) {
}

func (m *DeleteOp) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Key")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget3.FromBytes([]byte(m.Key), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("PermKey")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget5.FromBytes([]byte(m.PermKey), tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("PermVersion")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget7.FromBytes([]byte(m.PermVersion), tt.NonOptional().Field(2).Type); err != nil {
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

func (m *DeleteOp) MakeVDLTarget() vdl.Target {
	return &DeleteOpTarget{Value: m}
}

type DeleteOpTarget struct {
	Value             *DeleteOp
	keyTarget         vdl.BytesTarget
	permKeyTarget     vdl.BytesTarget
	permVersionTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *DeleteOpTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*DeleteOp)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *DeleteOpTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Key":
		t.keyTarget.Value = &t.Value.Key
		target, err := &t.keyTarget, error(nil)
		return nil, target, err
	case "PermKey":
		t.permKeyTarget.Value = &t.Value.PermKey
		target, err := &t.permKeyTarget, error(nil)
		return nil, target, err
	case "PermVersion":
		t.permVersionTarget.Value = &t.Value.PermVersion
		target, err := &t.permVersionTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/store/watchable.DeleteOp", name)
	}
}
func (t *DeleteOpTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *DeleteOpTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// LogEntry represents a single store operation. This operation may have been
// part of a transaction, as signified by the Continued boolean. Read-only
// operations (and read-only transactions) are not logged.
type LogEntry struct {
	// The store operation that was performed.
	Op *vom.RawBytes
	// Time when the operation was committed in nanoseconds since the epoch.
	// Note: We don't use time.Time here because VDL's time.Time consists of
	// {Seconds int64, Nanos int32}, which is more expensive than a single int64.
	CommitTimestamp int64
	// Operation came from sync (used for echo suppression).
	// TODO(razvanm): this field is specific to syncbase. We should add a
	// generic way to add fields and use that instead.
	FromSync bool
	// If true, this entry is followed by more entries that belong to the same
	// commit as this entry.
	Continued bool
}

func (LogEntry) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/store/watchable.LogEntry"`
}) {
}

func (m *LogEntry) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Op")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if m.Op == nil {
			if err := fieldTarget3.FromNil(tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
		} else {
			if err := m.Op.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("CommitTimestamp")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromInt(int64(m.CommitTimestamp), tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("FromSync")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget7.FromBool(bool(m.FromSync), tt.NonOptional().Field(2).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("Continued")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget9.FromBool(bool(m.Continued), tt.NonOptional().Field(3).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *LogEntry) MakeVDLTarget() vdl.Target {
	return &LogEntryTarget{Value: m}
}

type LogEntryTarget struct {
	Value *LogEntry

	commitTimestampTarget vdl.Int64Target
	fromSyncTarget        vdl.BoolTarget
	continuedTarget       vdl.BoolTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *LogEntryTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*LogEntry)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *LogEntryTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Op":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.Op))
		return nil, target, err
	case "CommitTimestamp":
		t.commitTimestampTarget.Value = &t.Value.CommitTimestamp
		target, err := &t.commitTimestampTarget, error(nil)
		return nil, target, err
	case "FromSync":
		t.fromSyncTarget.Value = &t.Value.FromSync
		target, err := &t.fromSyncTarget, error(nil)
		return nil, target, err
	case "Continued":
		t.continuedTarget.Value = &t.Value.Continued
		target, err := &t.continuedTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/services/syncbase/store/watchable.LogEntry", name)
	}
}
func (t *LogEntryTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *LogEntryTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// Create zero values for each type.
var (
	__VDLZeroGetOp    = GetOp{}
	__VDLZeroScanOp   = ScanOp{}
	__VDLZeroPutOp    = PutOp{}
	__VDLZeroDeleteOp = DeleteOp{}
	__VDLZeroLogEntry = LogEntry{}
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
	vdl.Register((*GetOp)(nil))
	vdl.Register((*ScanOp)(nil))
	vdl.Register((*PutOp)(nil))
	vdl.Register((*DeleteOp)(nil))
	vdl.Register((*LogEntry)(nil))

	return struct{}{}
}
