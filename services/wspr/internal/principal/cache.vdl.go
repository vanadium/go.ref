// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: cache.vdl

package principal

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/security"
)

// Identifier of a blessings cache entry.
type BlessingsId uint32

func (BlessingsId) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/principal.BlessingsId"`
}) {
}

func (m BlessingsId) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromUint(uint64(m), __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsId); err != nil {
		return err
	}
	return nil
}

func (m BlessingsId) MakeVDLTarget() vdl.Target {
	return nil
}

type BlessingsCacheAddMessage struct {
	CacheId   BlessingsId
	Blessings security.Blessings
}

func (BlessingsCacheAddMessage) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/principal.BlessingsCacheAddMessage"`
}) {
}

func (m *BlessingsCacheAddMessage) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	__VDLEnsureNativeBuilt_cache()
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("CacheId")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.CacheId.FillVDLTarget(fieldTarget3, __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsId); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	var wireValue4 security.WireBlessings
	if err := security.WireBlessingsFromNative(&wireValue4, m.Blessings); err != nil {
		return err
	}

	keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Blessings")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue4.FillVDLTarget(fieldTarget6, __VDLType_cache_v_io_v23_security_WireBlessings); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *BlessingsCacheAddMessage) MakeVDLTarget() vdl.Target {
	return nil
}

// Message from Blessings Cache GC to delete a cache entry in Javascript.
type BlessingsCacheDeleteMessage struct {
	CacheId BlessingsId
	// Number of references expected. Javascript should wait until this number
	// has been received before deleting the entry because up until that point
	// messages with further references are expected.
	DeleteAfter uint32
}

func (BlessingsCacheDeleteMessage) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/principal.BlessingsCacheDeleteMessage"`
}) {
}

func (m *BlessingsCacheDeleteMessage) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheDeleteMessage == nil || __VDLTypecache1 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("CacheId")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.CacheId.FillVDLTarget(fieldTarget3, __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsId); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("DeleteAfter")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromUint(uint64(m.DeleteAfter), vdl.Uint32Type); err != nil {
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

func (m *BlessingsCacheDeleteMessage) MakeVDLTarget() vdl.Target {
	return nil
}

type (
	// BlessingsCacheMessage represents any single field of the BlessingsCacheMessage union type.
	BlessingsCacheMessage interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the BlessingsCacheMessage union type.
		__VDLReflect(__BlessingsCacheMessageReflect)
		FillVDLTarget(vdl.Target, *vdl.Type) error
	}
	// BlessingsCacheMessageAdd represents field Add of the BlessingsCacheMessage union type.
	BlessingsCacheMessageAdd struct{ Value BlessingsCacheAddMessage }
	// BlessingsCacheMessageDelete represents field Delete of the BlessingsCacheMessage union type.
	BlessingsCacheMessageDelete struct{ Value BlessingsCacheDeleteMessage }
	// __BlessingsCacheMessageReflect describes the BlessingsCacheMessage union type.
	__BlessingsCacheMessageReflect struct {
		Name  string `vdl:"v.io/x/ref/services/wspr/internal/principal.BlessingsCacheMessage"`
		Type  BlessingsCacheMessage
		Union struct {
			Add    BlessingsCacheMessageAdd
			Delete BlessingsCacheMessageDelete
		}
	}
)

func (x BlessingsCacheMessageAdd) Index() int                                  { return 0 }
func (x BlessingsCacheMessageAdd) Interface() interface{}                      { return x.Value }
func (x BlessingsCacheMessageAdd) Name() string                                { return "Add" }
func (x BlessingsCacheMessageAdd) __VDLReflect(__BlessingsCacheMessageReflect) {}

func (m BlessingsCacheMessageAdd) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Add")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m BlessingsCacheMessageAdd) MakeVDLTarget() vdl.Target {
	return nil
}

func (x BlessingsCacheMessageDelete) Index() int                                  { return 1 }
func (x BlessingsCacheMessageDelete) Interface() interface{}                      { return x.Value }
func (x BlessingsCacheMessageDelete) Name() string                                { return "Delete" }
func (x BlessingsCacheMessageDelete) __VDLReflect(__BlessingsCacheMessageReflect) {}

func (m BlessingsCacheMessageDelete) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Delete")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheDeleteMessage); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m BlessingsCacheMessageDelete) MakeVDLTarget() vdl.Target {
	return nil
}

func init() {
	vdl.Register((*BlessingsId)(nil))
	vdl.Register((*BlessingsCacheAddMessage)(nil))
	vdl.Register((*BlessingsCacheDeleteMessage)(nil))
	vdl.Register((*BlessingsCacheMessage)(nil))
}

var __VDLTypecache0 *vdl.Type

func __VDLTypecache0_gen() *vdl.Type {
	__VDLTypecache0Builder := vdl.TypeBuilder{}

	__VDLTypecache01 := __VDLTypecache0Builder.Optional()
	__VDLTypecache02 := __VDLTypecache0Builder.Struct()
	__VDLTypecache03 := __VDLTypecache0Builder.Named("v.io/x/ref/services/wspr/internal/principal.BlessingsCacheAddMessage").AssignBase(__VDLTypecache02)
	__VDLTypecache04 := vdl.Uint32Type
	__VDLTypecache05 := __VDLTypecache0Builder.Named("v.io/x/ref/services/wspr/internal/principal.BlessingsId").AssignBase(__VDLTypecache04)
	__VDLTypecache02.AppendField("CacheId", __VDLTypecache05)
	__VDLTypecache06 := __VDLTypecache0Builder.Struct()
	__VDLTypecache07 := __VDLTypecache0Builder.Named("v.io/v23/security.WireBlessings").AssignBase(__VDLTypecache06)
	__VDLTypecache08 := __VDLTypecache0Builder.List()
	__VDLTypecache09 := __VDLTypecache0Builder.List()
	__VDLTypecache010 := __VDLTypecache0Builder.Struct()
	__VDLTypecache011 := __VDLTypecache0Builder.Named("v.io/v23/security.Certificate").AssignBase(__VDLTypecache010)
	__VDLTypecache012 := vdl.StringType
	__VDLTypecache010.AppendField("Extension", __VDLTypecache012)
	__VDLTypecache013 := __VDLTypecache0Builder.List()
	__VDLTypecache014 := vdl.ByteType
	__VDLTypecache013.AssignElem(__VDLTypecache014)
	__VDLTypecache010.AppendField("PublicKey", __VDLTypecache013)
	__VDLTypecache015 := __VDLTypecache0Builder.List()
	__VDLTypecache016 := __VDLTypecache0Builder.Struct()
	__VDLTypecache017 := __VDLTypecache0Builder.Named("v.io/v23/security.Caveat").AssignBase(__VDLTypecache016)
	__VDLTypecache018 := __VDLTypecache0Builder.Array()
	__VDLTypecache019 := __VDLTypecache0Builder.Named("v.io/v23/uniqueid.Id").AssignBase(__VDLTypecache018)
	__VDLTypecache018.AssignElem(__VDLTypecache014)
	__VDLTypecache018.AssignLen(16)
	__VDLTypecache016.AppendField("Id", __VDLTypecache019)
	__VDLTypecache016.AppendField("ParamVom", __VDLTypecache013)
	__VDLTypecache015.AssignElem(__VDLTypecache017)
	__VDLTypecache010.AppendField("Caveats", __VDLTypecache015)
	__VDLTypecache020 := __VDLTypecache0Builder.Struct()
	__VDLTypecache021 := __VDLTypecache0Builder.Named("v.io/v23/security.Signature").AssignBase(__VDLTypecache020)
	__VDLTypecache020.AppendField("Purpose", __VDLTypecache013)
	__VDLTypecache022 := vdl.StringType
	__VDLTypecache023 := __VDLTypecache0Builder.Named("v.io/v23/security.Hash").AssignBase(__VDLTypecache022)
	__VDLTypecache020.AppendField("Hash", __VDLTypecache023)
	__VDLTypecache020.AppendField("R", __VDLTypecache013)
	__VDLTypecache020.AppendField("S", __VDLTypecache013)
	__VDLTypecache010.AppendField("Signature", __VDLTypecache021)
	__VDLTypecache09.AssignElem(__VDLTypecache011)
	__VDLTypecache08.AssignElem(__VDLTypecache09)
	__VDLTypecache06.AppendField("CertificateChains", __VDLTypecache08)
	__VDLTypecache02.AppendField("Blessings", __VDLTypecache07)
	__VDLTypecache01.AssignElem(__VDLTypecache03)
	__VDLTypecache0Builder.Build()
	__VDLTypecache0v, err := __VDLTypecache01.Built()
	if err != nil {
		panic(err)
	}
	return __VDLTypecache0v
}
func init() {
	__VDLTypecache0 = __VDLTypecache0_gen()
}

var __VDLTypecache1 *vdl.Type = vdl.TypeOf((*BlessingsCacheDeleteMessage)(nil))
var __VDLType_cache_v_io_v23_security_WireBlessings *vdl.Type

func __VDLType_cache_v_io_v23_security_WireBlessings_gen() *vdl.Type {
	__VDLType_cache_v_io_v23_security_WireBlessingsBuilder := vdl.TypeBuilder{}

	__VDLType_cache_v_io_v23_security_WireBlessings1 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.Struct()
	__VDLType_cache_v_io_v23_security_WireBlessings2 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.Named("v.io/v23/security.WireBlessings").AssignBase(__VDLType_cache_v_io_v23_security_WireBlessings1)
	__VDLType_cache_v_io_v23_security_WireBlessings3 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.List()
	__VDLType_cache_v_io_v23_security_WireBlessings4 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.List()
	__VDLType_cache_v_io_v23_security_WireBlessings5 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.Struct()
	__VDLType_cache_v_io_v23_security_WireBlessings6 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.Named("v.io/v23/security.Certificate").AssignBase(__VDLType_cache_v_io_v23_security_WireBlessings5)
	__VDLType_cache_v_io_v23_security_WireBlessings7 := vdl.StringType
	__VDLType_cache_v_io_v23_security_WireBlessings5.AppendField("Extension", __VDLType_cache_v_io_v23_security_WireBlessings7)
	__VDLType_cache_v_io_v23_security_WireBlessings8 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.List()
	__VDLType_cache_v_io_v23_security_WireBlessings9 := vdl.ByteType
	__VDLType_cache_v_io_v23_security_WireBlessings8.AssignElem(__VDLType_cache_v_io_v23_security_WireBlessings9)
	__VDLType_cache_v_io_v23_security_WireBlessings5.AppendField("PublicKey", __VDLType_cache_v_io_v23_security_WireBlessings8)
	__VDLType_cache_v_io_v23_security_WireBlessings10 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.List()
	__VDLType_cache_v_io_v23_security_WireBlessings11 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.Struct()
	__VDLType_cache_v_io_v23_security_WireBlessings12 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.Named("v.io/v23/security.Caveat").AssignBase(__VDLType_cache_v_io_v23_security_WireBlessings11)
	__VDLType_cache_v_io_v23_security_WireBlessings13 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.Array()
	__VDLType_cache_v_io_v23_security_WireBlessings14 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.Named("v.io/v23/uniqueid.Id").AssignBase(__VDLType_cache_v_io_v23_security_WireBlessings13)
	__VDLType_cache_v_io_v23_security_WireBlessings13.AssignElem(__VDLType_cache_v_io_v23_security_WireBlessings9)
	__VDLType_cache_v_io_v23_security_WireBlessings13.AssignLen(16)
	__VDLType_cache_v_io_v23_security_WireBlessings11.AppendField("Id", __VDLType_cache_v_io_v23_security_WireBlessings14)
	__VDLType_cache_v_io_v23_security_WireBlessings11.AppendField("ParamVom", __VDLType_cache_v_io_v23_security_WireBlessings8)
	__VDLType_cache_v_io_v23_security_WireBlessings10.AssignElem(__VDLType_cache_v_io_v23_security_WireBlessings12)
	__VDLType_cache_v_io_v23_security_WireBlessings5.AppendField("Caveats", __VDLType_cache_v_io_v23_security_WireBlessings10)
	__VDLType_cache_v_io_v23_security_WireBlessings15 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.Struct()
	__VDLType_cache_v_io_v23_security_WireBlessings16 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.Named("v.io/v23/security.Signature").AssignBase(__VDLType_cache_v_io_v23_security_WireBlessings15)
	__VDLType_cache_v_io_v23_security_WireBlessings15.AppendField("Purpose", __VDLType_cache_v_io_v23_security_WireBlessings8)
	__VDLType_cache_v_io_v23_security_WireBlessings17 := vdl.StringType
	__VDLType_cache_v_io_v23_security_WireBlessings18 := __VDLType_cache_v_io_v23_security_WireBlessingsBuilder.Named("v.io/v23/security.Hash").AssignBase(__VDLType_cache_v_io_v23_security_WireBlessings17)
	__VDLType_cache_v_io_v23_security_WireBlessings15.AppendField("Hash", __VDLType_cache_v_io_v23_security_WireBlessings18)
	__VDLType_cache_v_io_v23_security_WireBlessings15.AppendField("R", __VDLType_cache_v_io_v23_security_WireBlessings8)
	__VDLType_cache_v_io_v23_security_WireBlessings15.AppendField("S", __VDLType_cache_v_io_v23_security_WireBlessings8)
	__VDLType_cache_v_io_v23_security_WireBlessings5.AppendField("Signature", __VDLType_cache_v_io_v23_security_WireBlessings16)
	__VDLType_cache_v_io_v23_security_WireBlessings4.AssignElem(__VDLType_cache_v_io_v23_security_WireBlessings6)
	__VDLType_cache_v_io_v23_security_WireBlessings3.AssignElem(__VDLType_cache_v_io_v23_security_WireBlessings4)
	__VDLType_cache_v_io_v23_security_WireBlessings1.AppendField("CertificateChains", __VDLType_cache_v_io_v23_security_WireBlessings3)
	__VDLType_cache_v_io_v23_security_WireBlessingsBuilder.Build()
	__VDLType_cache_v_io_v23_security_WireBlessingsv, err := __VDLType_cache_v_io_v23_security_WireBlessings2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_cache_v_io_v23_security_WireBlessingsv
}
func init() {
	__VDLType_cache_v_io_v23_security_WireBlessings = __VDLType_cache_v_io_v23_security_WireBlessings_gen()
}

var __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage *vdl.Type

func __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage_gen() *vdl.Type {
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder := vdl.TypeBuilder{}

	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage1 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Struct()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage2 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Named("v.io/x/ref/services/wspr/internal/principal.BlessingsCacheAddMessage").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage1)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage3 := vdl.Uint32Type
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage4 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Named("v.io/x/ref/services/wspr/internal/principal.BlessingsId").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage3)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage1.AppendField("CacheId", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage4)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage5 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Struct()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage6 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Named("v.io/v23/security.WireBlessings").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage5)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage7 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.List()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage8 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.List()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage9 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Struct()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage10 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Named("v.io/v23/security.Certificate").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage9)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage11 := vdl.StringType
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage9.AppendField("Extension", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage11)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage12 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.List()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage13 := vdl.ByteType
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage12.AssignElem(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage13)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage9.AppendField("PublicKey", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage12)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage14 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.List()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage15 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Struct()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage16 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Named("v.io/v23/security.Caveat").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage15)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage17 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Array()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage18 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Named("v.io/v23/uniqueid.Id").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage17)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage17.AssignElem(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage13)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage17.AssignLen(16)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage15.AppendField("Id", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage18)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage15.AppendField("ParamVom", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage12)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage14.AssignElem(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage16)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage9.AppendField("Caveats", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage14)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage19 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Struct()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage20 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Named("v.io/v23/security.Signature").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage19)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage19.AppendField("Purpose", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage12)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage21 := vdl.StringType
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage22 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Named("v.io/v23/security.Hash").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage21)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage19.AppendField("Hash", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage22)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage19.AppendField("R", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage12)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage19.AppendField("S", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage12)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage9.AppendField("Signature", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage20)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage8.AssignElem(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage10)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage7.AssignElem(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage8)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage5.AppendField("CertificateChains", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage7)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage1.AppendField("Blessings", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage6)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessageBuilder.Build()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessagev, err := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessagev
}
func init() {
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage = __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage_gen()
}

var __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheDeleteMessage *vdl.Type = vdl.TypeOf(BlessingsCacheDeleteMessage{})
var __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage *vdl.Type

func __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage_gen() *vdl.Type {
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder := vdl.TypeBuilder{}

	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage1 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Union()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage2 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Named("v.io/x/ref/services/wspr/internal/principal.BlessingsCacheMessage").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage1)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage3 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Struct()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage4 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Named("v.io/x/ref/services/wspr/internal/principal.BlessingsCacheAddMessage").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage3)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage5 := vdl.Uint32Type
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage6 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Named("v.io/x/ref/services/wspr/internal/principal.BlessingsId").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage5)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage3.AppendField("CacheId", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage6)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage7 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Struct()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage8 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Named("v.io/v23/security.WireBlessings").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage7)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage9 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.List()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage10 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.List()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage11 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Struct()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage12 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Named("v.io/v23/security.Certificate").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage11)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage13 := vdl.StringType
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage11.AppendField("Extension", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage13)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage14 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.List()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage15 := vdl.ByteType
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage14.AssignElem(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage15)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage11.AppendField("PublicKey", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage14)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage16 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.List()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage17 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Struct()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage18 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Named("v.io/v23/security.Caveat").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage17)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage19 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Array()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage20 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Named("v.io/v23/uniqueid.Id").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage19)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage19.AssignElem(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage15)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage19.AssignLen(16)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage17.AppendField("Id", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage20)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage17.AppendField("ParamVom", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage14)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage16.AssignElem(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage18)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage11.AppendField("Caveats", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage16)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage21 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Struct()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage22 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Named("v.io/v23/security.Signature").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage21)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage21.AppendField("Purpose", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage14)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage23 := vdl.StringType
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage24 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Named("v.io/v23/security.Hash").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage23)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage21.AppendField("Hash", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage24)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage21.AppendField("R", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage14)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage21.AppendField("S", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage14)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage11.AppendField("Signature", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage22)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage10.AssignElem(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage12)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage9.AssignElem(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage10)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage7.AppendField("CertificateChains", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage9)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage3.AppendField("Blessings", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage8)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage1.AppendField("Add", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage4)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage25 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Struct()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage26 := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Named("v.io/x/ref/services/wspr/internal/principal.BlessingsCacheDeleteMessage").AssignBase(__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage25)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage25.AppendField("CacheId", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage6)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage27 := vdl.Uint32Type
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage25.AppendField("DeleteAfter", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage27)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage1.AppendField("Delete", __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage26)
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessageBuilder.Build()
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessagev, err := __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessagev
}
func init() {
	__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage = __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage_gen()
}

var __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsId *vdl.Type = vdl.TypeOf(BlessingsId(0))

func __VDLEnsureNativeBuilt_cache() {
	if __VDLTypecache0 == nil {
		__VDLTypecache0 = __VDLTypecache0_gen()
	}
	if __VDLType_cache_v_io_v23_security_WireBlessings == nil {
		__VDLType_cache_v_io_v23_security_WireBlessings = __VDLType_cache_v_io_v23_security_WireBlessings_gen()
	}
	if __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage == nil {
		__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage = __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheAddMessage_gen()
	}
	if __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage == nil {
		__VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage = __VDLType_cache_v_io_x_ref_services_wspr_internal_principal_BlessingsCacheMessage_gen()
	}
}
