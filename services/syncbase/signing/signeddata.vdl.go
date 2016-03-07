// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: signeddata.vdl

package signing

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/security"
)

// A DataWithSignature represents a signed, and possibily validated, collection
// of Item structs.
//
// If IsValidated==false and the AuthorSigned signature is valid, it means:
//    The signer whose Blessings have hash BlessingsHash asserts Data.
//
// If IsValidated==true and both AuthorSigned and ValidatorSigned signatures are is valid,
// it means both:
// 1) The signer whose Blessings b have hash BlessingsHash asserts Data.
// 2) If vd is the ValidatorData with hash ValidatorDataHash, the owner of
//    vd.PublicKey asserts that it checked that at least the names vd.Names[] were
//    valid in b.
//
// The sender obtains:
// - BlessingsHash (and the wire form of the blessings) with ValidationCache.AddBlessings().
// - ValidatorDataHash (and the wire form of the ValidataData)  with ValidationCache.AddValidatorData().
//
// The receiver looks up:
// - BlessingsHash with ValidationCache.LookupBlessingsData()
// - ValidatorDataHash with ValidationCache.LookupValidatorData()
//
// If not yet there, the receiver inserts the valus into its ValidationCache with:
// - ValidationCache.AddWireBlessings()
// - ValidationCache.AddValidatorData()
type DataWithSignature struct {
	Data []Item
	// BlessingsHash is a key for the validation cache; the corresponding
	// cached value is a security.Blessings.
	BlessingsHash []byte
	// AuthorSigned is the signature of Data and BlessingsHash using the
	// private key associated with the blessings hashed in BlessingsHash.
	AuthorSigned security.Signature
	IsValidated  bool // Whether fields below are meaningful.
	// ValidatorDataHash is a key for the validation cache returned by
	// ValidatorData.Hash(); the corresponding cached value is the
	// ValidatorData.
	ValidatorDataHash []byte
	ValidatorSigned   security.Signature
}

func (DataWithSignature) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/signing.DataWithSignature"`
}) {
}

func (m *DataWithSignature) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_signeddata_v_io_x_ref_services_syncbase_signing_DataWithSignature == nil || __VDLTypesigneddata0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Data")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget4, err := fieldTarget3.StartList(__VDLTypesigneddata1, len(m.Data))
		if err != nil {
			return err
		}
		for i, elem6 := range m.Data {
			elemTarget5, err := listTarget4.StartElem(i)
			if err != nil {
				return err
			}

			unionValue7 := elem6
			if unionValue7 == nil {
				unionValue7 = ItemData{}
			}
			if err := unionValue7.FillVDLTarget(elemTarget5, __VDLType_signeddata_v_io_x_ref_services_syncbase_signing_Item); err != nil {
				return err
			}
			if err := listTarget4.FinishElem(elemTarget5); err != nil {
				return err
			}
		}
		if err := fieldTarget3.FinishList(listTarget4); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("BlessingsHash")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget9.FromBytes([]byte(m.BlessingsHash), __VDLTypesigneddata2); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
			return err
		}
	}
	keyTarget10, fieldTarget11, err := fieldsTarget1.StartField("AuthorSigned")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.AuthorSigned.FillVDLTarget(fieldTarget11, __VDLType_signeddata_v_io_v23_security_Signature); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget10, fieldTarget11); err != nil {
			return err
		}
	}
	keyTarget12, fieldTarget13, err := fieldsTarget1.StartField("IsValidated")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget13.FromBool(bool(m.IsValidated), vdl.BoolType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget12, fieldTarget13); err != nil {
			return err
		}
	}
	keyTarget14, fieldTarget15, err := fieldsTarget1.StartField("ValidatorDataHash")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget15.FromBytes([]byte(m.ValidatorDataHash), __VDLTypesigneddata2); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget14, fieldTarget15); err != nil {
			return err
		}
	}
	keyTarget16, fieldTarget17, err := fieldsTarget1.StartField("ValidatorSigned")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.ValidatorSigned.FillVDLTarget(fieldTarget17, __VDLType_signeddata_v_io_v23_security_Signature); err != nil {
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

func (m *DataWithSignature) MakeVDLTarget() vdl.Target {
	return nil
}

type (
	// Item represents any single field of the Item union type.
	//
	// An Item represents either a marshalled data item or its SHA-256 hash.
	// The Data field is a []byte, rather than an "any" to make signatures
	// determistic.  VOM encoding is not deterministic for two reasons:
	// - map elements may be marshalled in any order
	// - different versions of VOM may marshal in different ways.
	// Thus, the initial producer of a data item marshals the data once, and it is
	// this marshalled form that is transmitted from device to device.  If the
	// data were unmarshalled and then remarsahalled, the signatures might not
	// match.  The Hash field is used instead of the Data field when the recipient
	// of the DataWithSignature is not permitted to see certain Items' Data
	// fields.
	Item interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the Item union type.
		__VDLReflect(__ItemReflect)
		FillVDLTarget(vdl.Target, *vdl.Type) error
	}
	// ItemData represents field Data of the Item union type.
	ItemData struct{ Value []byte } // Marshalled form of data.
	// ItemHash represents field Hash of the Item union type.
	ItemHash struct{ Value []byte } // Hash of what would have been in Data, as returned by SumByteVectorWithLength(Data).
	// __ItemReflect describes the Item union type.
	__ItemReflect struct {
		Name  string `vdl:"v.io/x/ref/services/syncbase/signing.Item"`
		Type  Item
		Union struct {
			Data ItemData
			Hash ItemHash
		}
	}
)

func (x ItemData) Index() int                 { return 0 }
func (x ItemData) Interface() interface{}     { return x.Value }
func (x ItemData) Name() string               { return "Data" }
func (x ItemData) __VDLReflect(__ItemReflect) {}

func (m ItemData) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_signeddata_v_io_x_ref_services_syncbase_signing_Item)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Data")
	if err != nil {
		return err
	}

	if err := fieldTarget3.FromBytes([]byte(m.Value), __VDLTypesigneddata2); err != nil {
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

func (m ItemData) MakeVDLTarget() vdl.Target {
	return nil
}

func (x ItemHash) Index() int                 { return 1 }
func (x ItemHash) Interface() interface{}     { return x.Value }
func (x ItemHash) Name() string               { return "Hash" }
func (x ItemHash) __VDLReflect(__ItemReflect) {}

func (m ItemHash) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_signeddata_v_io_x_ref_services_syncbase_signing_Item)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Hash")
	if err != nil {
		return err
	}

	if err := fieldTarget3.FromBytes([]byte(m.Value), __VDLTypesigneddata2); err != nil {
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

func (m ItemHash) MakeVDLTarget() vdl.Target {
	return nil
}

// WireValidatorData is the wire form of ValidatorData.
// It excludes the unmarshalled form of the public key.
type WireValidatorData struct {
	Names               []string // Names of valid signing blessings in the Blessings referred to by BlessingsHash.
	MarshalledPublicKey []byte   // PublicKey, marshalled with MarshalBinary().
}

func (WireValidatorData) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/signing.WireValidatorData"`
}) {
}

func (m *WireValidatorData) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_signeddata_v_io_x_ref_services_syncbase_signing_WireValidatorData == nil || __VDLTypesigneddata3 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Names")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget4, err := fieldTarget3.StartList(__VDLTypesigneddata4, len(m.Names))
		if err != nil {
			return err
		}
		for i, elem6 := range m.Names {
			elemTarget5, err := listTarget4.StartElem(i)
			if err != nil {
				return err
			}
			if err := elemTarget5.FromString(string(elem6), vdl.StringType); err != nil {
				return err
			}
			if err := listTarget4.FinishElem(elemTarget5); err != nil {
				return err
			}
		}
		if err := fieldTarget3.FinishList(listTarget4); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget7, fieldTarget8, err := fieldsTarget1.StartField("MarshalledPublicKey")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget8.FromBytes([]byte(m.MarshalledPublicKey), __VDLTypesigneddata2); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget7, fieldTarget8); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *WireValidatorData) MakeVDLTarget() vdl.Target {
	return nil
}

func init() {
	vdl.Register((*DataWithSignature)(nil))
	vdl.Register((*Item)(nil))
	vdl.Register((*WireValidatorData)(nil))
}

var __VDLTypesigneddata0 *vdl.Type = vdl.TypeOf((*DataWithSignature)(nil))
var __VDLTypesigneddata3 *vdl.Type = vdl.TypeOf((*WireValidatorData)(nil))
var __VDLTypesigneddata2 *vdl.Type = vdl.TypeOf([]byte(nil))
var __VDLTypesigneddata4 *vdl.Type = vdl.TypeOf([]string(nil))
var __VDLTypesigneddata1 *vdl.Type = vdl.TypeOf([]Item(nil))
var __VDLType_signeddata_v_io_v23_security_Signature *vdl.Type = vdl.TypeOf(security.Signature{})
var __VDLType_signeddata_v_io_x_ref_services_syncbase_signing_DataWithSignature *vdl.Type = vdl.TypeOf(DataWithSignature{})
var __VDLType_signeddata_v_io_x_ref_services_syncbase_signing_Item *vdl.Type = vdl.TypeOf(Item(ItemData{[]byte(nil)}))
var __VDLType_signeddata_v_io_x_ref_services_syncbase_signing_WireValidatorData *vdl.Type = vdl.TypeOf(WireValidatorData{})

func __VDLEnsureNativeBuilt_signeddata() {
}
