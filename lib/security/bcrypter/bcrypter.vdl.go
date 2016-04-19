// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: bcrypter

package bcrypter

import (
	"fmt"
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/security"
	"v.io/v23/vdl"
	"v.io/v23/verror"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// WireCiphertext represents the wire format of the ciphertext
// generated by a Crypter.
type WireCiphertext struct {
	// PatternId is an identifier of the blessing pattern that this
	// ciphertext is for. It is represented by a 16 byte truncated
	// SHA256 hash of the pattern.
	PatternId string
	// Bytes is a map from an identifier of the public IBE params to
	// the ciphertext bytes that were generated using those params.
	//
	// The params identifier is a 16 byte truncated SHA256 hash
	// of the marshaled form of the IBE params.
	Bytes map[string][]byte
}

func (WireCiphertext) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security/bcrypter.WireCiphertext"`
}) {
}

func (m *WireCiphertext) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.PatternId == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("PatternId"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("PatternId")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.PatternId), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var var7 bool
	if len(m.Bytes) == 0 {
		var7 = true
	}
	if var7 {
		if err := fieldsTarget1.ZeroField("Bytes"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Bytes")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			mapTarget8, err := fieldTarget6.StartMap(tt.NonOptional().Field(1).Type, len(m.Bytes))
			if err != nil {
				return err
			}
			for key10, value12 := range m.Bytes {
				keyTarget9, err := mapTarget8.StartKey()
				if err != nil {
					return err
				}
				if err := keyTarget9.FromString(string(key10), tt.NonOptional().Field(1).Type.Key()); err != nil {
					return err
				}
				valueTarget11, err := mapTarget8.FinishKeyStartField(keyTarget9)
				if err != nil {
					return err
				}

				if err := valueTarget11.FromBytes([]byte(value12), tt.NonOptional().Field(1).Type.Elem()); err != nil {
					return err
				}
				if err := mapTarget8.FinishField(keyTarget9, valueTarget11); err != nil {
					return err
				}
			}
			if err := fieldTarget6.FinishMap(mapTarget8); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *WireCiphertext) MakeVDLTarget() vdl.Target {
	return &WireCiphertextTarget{Value: m}
}

type WireCiphertextTarget struct {
	Value           *WireCiphertext
	patternIdTarget vdl.StringTarget
	bytesTarget     __VDLTarget1_map
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *WireCiphertextTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*WireCiphertext)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *WireCiphertextTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "PatternId":
		t.patternIdTarget.Value = &t.Value.PatternId
		target, err := &t.patternIdTarget, error(nil)
		return nil, target, err
	case "Bytes":
		t.bytesTarget.Value = &t.Value.Bytes
		target, err := &t.bytesTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/lib/security/bcrypter.WireCiphertext", name)
	}
}
func (t *WireCiphertextTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *WireCiphertextTarget) ZeroField(name string) error {
	switch name {
	case "PatternId":
		t.Value.PatternId = ""
		return nil
	case "Bytes":
		t.Value.Bytes = map[string][]byte(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/x/ref/lib/security/bcrypter.WireCiphertext", name)
	}
}
func (t *WireCiphertextTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[string][]byte
type __VDLTarget1_map struct {
	Value      *map[string][]byte
	currKey    string
	currElem   []byte
	keyTarget  vdl.StringTarget
	elemTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *__VDLTarget1_map) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {

	if ttWant := vdl.TypeOf((*map[string][]byte)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(map[string][]byte)
	return t, nil
}
func (t *__VDLTarget1_map) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *__VDLTarget1_map) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = []byte(nil)
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

func (x *WireCiphertext) VDLRead(dec vdl.Decoder) error {
	*x = WireCiphertext{}
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "PatternId":
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.PatternId, err = dec.DecodeString(); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "Bytes":
			if err = __VDLRead1_map(dec, &x.Bytes); err != nil {
				return err
			}
		default:
			if err = dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLRead1_map(dec vdl.Decoder, x *map[string][]byte) error {
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible map %T, from %v", *x, dec.Type())
	}
	var tmpMap map[string][]byte
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(map[string][]byte, len)
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		}
		var key string
		{
			if err = dec.StartValue(); err != nil {
				return err
			}
			if key, err = dec.DecodeString(); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		}
		var elem []byte
		{
			if err = dec.StartValue(); err != nil {
				return err
			}
			if err = dec.DecodeBytes(-1, &elem); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		}
		if tmpMap == nil {
			tmpMap = make(map[string][]byte)
		}
		tmpMap[key] = elem
	}
}

func (x WireCiphertext) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*WireCiphertext)(nil)).Elem()); err != nil {
		return err
	}
	var1 := (x.PatternId == "")
	if !(var1) {
		if err := enc.NextField("PatternId"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*string)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeString(x.PatternId); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	var var2 bool
	if len(x.Bytes) == 0 {
		var2 = true
	}
	if !(var2) {
		if err := enc.NextField("Bytes"); err != nil {
			return err
		}
		if err := __VDLWrite1_map(enc, &x.Bytes); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWrite1_map(enc vdl.Encoder, x *map[string][]byte) error {
	if err := enc.StartValue(vdl.TypeOf((*map[string][]byte)(nil))); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(*x)); err != nil {
		return err
	}
	for key, elem := range *x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*string)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeString(key); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*[]byte)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeBytes(elem); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

// WireParams represents the wire format of the public parameters
// of an identity provider (aka Root).
type WireParams struct {
	// Blessing is the blessing name of the identity provider. The identity
	// provider  can extract private keys for blessings that are extensions
	// of this blessing name.
	Blessing string
	// Params is the marshaled form of the public IBE params of the
	// the identity provider.
	Params []byte
}

func (WireParams) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security/bcrypter.WireParams"`
}) {
}

func (m *WireParams) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Blessing == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Blessing"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Blessing")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Blessing), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var var7 bool
	if len(m.Params) == 0 {
		var7 = true
	}
	if var7 {
		if err := fieldsTarget1.ZeroField("Params"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Params")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := fieldTarget6.FromBytes([]byte(m.Params), tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *WireParams) MakeVDLTarget() vdl.Target {
	return &WireParamsTarget{Value: m}
}

type WireParamsTarget struct {
	Value          *WireParams
	blessingTarget vdl.StringTarget
	paramsTarget   vdl.BytesTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *WireParamsTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*WireParams)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *WireParamsTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Blessing":
		t.blessingTarget.Value = &t.Value.Blessing
		target, err := &t.blessingTarget, error(nil)
		return nil, target, err
	case "Params":
		t.paramsTarget.Value = &t.Value.Params
		target, err := &t.paramsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/lib/security/bcrypter.WireParams", name)
	}
}
func (t *WireParamsTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *WireParamsTarget) ZeroField(name string) error {
	switch name {
	case "Blessing":
		t.Value.Blessing = ""
		return nil
	case "Params":
		t.Value.Params = []byte(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/x/ref/lib/security/bcrypter.WireParams", name)
	}
}
func (t *WireParamsTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x *WireParams) VDLRead(dec vdl.Decoder) error {
	*x = WireParams{}
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Blessing":
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.Blessing, err = dec.DecodeString(); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "Params":
			if err = dec.StartValue(); err != nil {
				return err
			}
			if err = dec.DecodeBytes(-1, &x.Params); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		default:
			if err = dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func (x WireParams) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*WireParams)(nil)).Elem()); err != nil {
		return err
	}
	var1 := (x.Blessing == "")
	if !(var1) {
		if err := enc.NextField("Blessing"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*string)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Blessing); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	var var2 bool
	if len(x.Params) == 0 {
		var2 = true
	}
	if !(var2) {
		if err := enc.NextField("Params"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*[]byte)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeBytes(x.Params); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

// WirePrivateKey represents the wire format of the private key corresponding
// to a blessing.
type WirePrivateKey struct {
	// Blessing is the blessing for which this private key was extracted for.
	Blessing string
	// Params are the public parameters of the identity provider that extracted
	// this private key.
	Params WireParams
	// Keys contain the extracted IBE private keys for each pattern that is
	// matched by the blessing and is an extension of the identity provider's
	// name. The keys are enumerated in increasing order of the lengths of the
	// corresponding patterns.
	//
	// For example, if the blessing is "google:u:alice:phone" and the identity
	// provider's name is "google:u" then the keys are extracted for the patterns
	// - "google:u"
	// - "google:u:alice"
	// - "google:u:alice:phone"
	// - "google:u:alice:phone:$"
	//
	// The private keys are listed in increasing order of the lengths of the
	// corresponding patterns.
	Keys [][]byte
}

func (WirePrivateKey) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security/bcrypter.WirePrivateKey"`
}) {
}

func (m *WirePrivateKey) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Blessing == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Blessing"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Blessing")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Blessing), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := true
	var8 := (m.Params.Blessing == "")
	var7 = var7 && var8
	var var9 bool
	if len(m.Params.Params) == 0 {
		var9 = true
	}
	var7 = var7 && var9
	if var7 {
		if err := fieldsTarget1.ZeroField("Params"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Params")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Params.FillVDLTarget(fieldTarget6, tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	var var12 bool
	if len(m.Keys) == 0 {
		var12 = true
	}
	if var12 {
		if err := fieldsTarget1.ZeroField("Keys"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget10, fieldTarget11, err := fieldsTarget1.StartField("Keys")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			listTarget13, err := fieldTarget11.StartList(tt.NonOptional().Field(2).Type, len(m.Keys))
			if err != nil {
				return err
			}
			for i, elem15 := range m.Keys {
				elemTarget14, err := listTarget13.StartElem(i)
				if err != nil {
					return err
				}

				if err := elemTarget14.FromBytes([]byte(elem15), tt.NonOptional().Field(2).Type.Elem()); err != nil {
					return err
				}
				if err := listTarget13.FinishElem(elemTarget14); err != nil {
					return err
				}
			}
			if err := fieldTarget11.FinishList(listTarget13); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget10, fieldTarget11); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *WirePrivateKey) MakeVDLTarget() vdl.Target {
	return &WirePrivateKeyTarget{Value: m}
}

type WirePrivateKeyTarget struct {
	Value          *WirePrivateKey
	blessingTarget vdl.StringTarget
	paramsTarget   WireParamsTarget
	keysTarget     __VDLTarget2_list
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *WirePrivateKeyTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*WirePrivateKey)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *WirePrivateKeyTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Blessing":
		t.blessingTarget.Value = &t.Value.Blessing
		target, err := &t.blessingTarget, error(nil)
		return nil, target, err
	case "Params":
		t.paramsTarget.Value = &t.Value.Params
		target, err := &t.paramsTarget, error(nil)
		return nil, target, err
	case "Keys":
		t.keysTarget.Value = &t.Value.Keys
		target, err := &t.keysTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/x/ref/lib/security/bcrypter.WirePrivateKey", name)
	}
}
func (t *WirePrivateKeyTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *WirePrivateKeyTarget) ZeroField(name string) error {
	switch name {
	case "Blessing":
		t.Value.Blessing = ""
		return nil
	case "Params":
		t.Value.Params = WireParams{}
		return nil
	case "Keys":
		t.Value.Keys = [][]byte(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/x/ref/lib/security/bcrypter.WirePrivateKey", name)
	}
}
func (t *WirePrivateKeyTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// [][]byte
type __VDLTarget2_list struct {
	Value      *[][]byte
	elemTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.ListTargetBase
}

func (t *__VDLTarget2_list) StartList(tt *vdl.Type, len int) (vdl.ListTarget, error) {

	if ttWant := vdl.TypeOf((*[][]byte)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	if cap(*t.Value) < len {
		*t.Value = make([][]byte, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *__VDLTarget2_list) StartElem(index int) (elem vdl.Target, _ error) {
	t.elemTarget.Value = &(*t.Value)[index]
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *__VDLTarget2_list) FinishElem(elem vdl.Target) error {
	return nil
}
func (t *__VDLTarget2_list) FinishList(elem vdl.ListTarget) error {

	return nil
}

func (x *WirePrivateKey) VDLRead(dec vdl.Decoder) error {
	*x = WirePrivateKey{}
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Blessing":
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.Blessing, err = dec.DecodeString(); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "Params":
			if err = x.Params.VDLRead(dec); err != nil {
				return err
			}
		case "Keys":
			if err = __VDLRead2_list(dec, &x.Keys); err != nil {
				return err
			}
		default:
			if err = dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLRead2_list(dec vdl.Decoder, x *[][]byte) error {
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible list %T, from %v", *x, dec.Type())
	}
	switch len := dec.LenHint(); {
	case len > 0:
		*x = make([][]byte, 0, len)
	default:
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		}
		var elem []byte
		if err = dec.StartValue(); err != nil {
			return err
		}
		if err = dec.DecodeBytes(-1, &elem); err != nil {
			return err
		}
		if err = dec.FinishValue(); err != nil {
			return err
		}
		*x = append(*x, elem)
	}
}

func (x WirePrivateKey) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*WirePrivateKey)(nil)).Elem()); err != nil {
		return err
	}
	var1 := (x.Blessing == "")
	if !(var1) {
		if err := enc.NextField("Blessing"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*string)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Blessing); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	var2 := true
	var3 := (x.Params.Blessing == "")
	var2 = var2 && var3
	var var4 bool
	if len(x.Params.Params) == 0 {
		var4 = true
	}
	var2 = var2 && var4
	if !(var2) {
		if err := enc.NextField("Params"); err != nil {
			return err
		}
		if err := x.Params.VDLWrite(enc); err != nil {
			return err
		}
	}
	var var5 bool
	if len(x.Keys) == 0 {
		var5 = true
	}
	if !(var5) {
		if err := enc.NextField("Keys"); err != nil {
			return err
		}
		if err := __VDLWrite2_list(enc, &x.Keys); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWrite2_list(enc vdl.Encoder, x *[][]byte) error {
	if err := enc.StartValue(vdl.TypeOf((*[][]byte)(nil))); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(*x)); err != nil {
		return err
	}
	for i := 0; i < len(*x); i++ {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*[]byte)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeBytes((*x)[i]); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

//////////////////////////////////////////////////
// Error definitions

var (
	ErrInternal           = verror.Register("v.io/x/ref/lib/security/bcrypter.Internal", verror.NoRetry, "{1:}{2:} internal error: {3}")
	ErrNoParams           = verror.Register("v.io/x/ref/lib/security/bcrypter.NoParams", verror.NoRetry, "{1:}{2:} no public parameters available for encrypting for pattern: {3}")
	ErrPrivateKeyNotFound = verror.Register("v.io/x/ref/lib/security/bcrypter.PrivateKeyNotFound", verror.NoRetry, "{1:}{2:} no private key found for decrypting ciphertext")
	ErrInvalidPrivateKey  = verror.Register("v.io/x/ref/lib/security/bcrypter.InvalidPrivateKey", verror.NoRetry, "{1:}{2:} private key is invalid: {3}")
)

// NewErrInternal returns an error with the ErrInternal ID.
func NewErrInternal(ctx *context.T, err error) error {
	return verror.New(ErrInternal, ctx, err)
}

// NewErrNoParams returns an error with the ErrNoParams ID.
func NewErrNoParams(ctx *context.T, pattern security.BlessingPattern) error {
	return verror.New(ErrNoParams, ctx, pattern)
}

// NewErrPrivateKeyNotFound returns an error with the ErrPrivateKeyNotFound ID.
func NewErrPrivateKeyNotFound(ctx *context.T) error {
	return verror.New(ErrPrivateKeyNotFound, ctx)
}

// NewErrInvalidPrivateKey returns an error with the ErrInvalidPrivateKey ID.
func NewErrInvalidPrivateKey(ctx *context.T, err error) error {
	return verror.New(ErrInvalidPrivateKey, ctx, err)
}

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
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*WireCiphertext)(nil))
	vdl.Register((*WireParams)(nil))
	vdl.Register((*WirePrivateKey)(nil))

	// Set error format strings.
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInternal.ID), "{1:}{2:} internal error: {3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoParams.ID), "{1:}{2:} no public parameters available for encrypting for pattern: {3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrPrivateKeyNotFound.ID), "{1:}{2:} no private key found for decrypting ciphertext")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidPrivateKey.ID), "{1:}{2:} private key is invalid: {3}")

	return struct{}{}
}
