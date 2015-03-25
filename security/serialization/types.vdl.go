// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package serialization

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/security"
)

type SignedHeader struct {
	ChunkSizeBytes int64
}

func (SignedHeader) __VDLReflect(struct {
	Name string "v.io/x/ref/security/serialization.SignedHeader"
}) {
}

type HashCode [32]byte

func (HashCode) __VDLReflect(struct {
	Name string "v.io/x/ref/security/serialization.HashCode"
}) {
}

type (
	// SignedData represents any single field of the SignedData union type.
	//
	// SignedData describes the information sent by a SigningWriter and read by VerifiyingReader.
	SignedData interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the SignedData union type.
		__VDLReflect(__SignedDataReflect)
	}
	// SignedDataSignature represents field Signature of the SignedData union type.
	SignedDataSignature struct{ Value security.Signature }
	// SignedDataHash represents field Hash of the SignedData union type.
	SignedDataHash struct{ Value HashCode }
	// __SignedDataReflect describes the SignedData union type.
	__SignedDataReflect struct {
		Name  string "v.io/x/ref/security/serialization.SignedData"
		Type  SignedData
		Union struct {
			Signature SignedDataSignature
			Hash      SignedDataHash
		}
	}
)

func (x SignedDataSignature) Index() int                       { return 0 }
func (x SignedDataSignature) Interface() interface{}           { return x.Value }
func (x SignedDataSignature) Name() string                     { return "Signature" }
func (x SignedDataSignature) __VDLReflect(__SignedDataReflect) {}

func (x SignedDataHash) Index() int                       { return 1 }
func (x SignedDataHash) Interface() interface{}           { return x.Value }
func (x SignedDataHash) Name() string                     { return "Hash" }
func (x SignedDataHash) __VDLReflect(__SignedDataReflect) {}

func init() {
	vdl.Register((*SignedHeader)(nil))
	vdl.Register((*HashCode)(nil))
	vdl.Register((*SignedData)(nil))
}
