// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: model.vdl

package store

import (
	// VDL system imports
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

func __VDLEnsureNativeBuilt_model() {
}

var (
	// KeyExists means the given key already exists in the store.
	ErrKeyExists = verror.Register("v.io/x/ref/services/groups/internal/store.KeyExists", verror.NoRetry, "{1:}{2:} Key exists{:_}")
	// UnknownKey means the given key does not exist in the store.
	ErrUnknownKey = verror.Register("v.io/x/ref/services/groups/internal/store.UnknownKey", verror.NoRetry, "{1:}{2:} Unknown key{:_}")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrKeyExists.ID), "{1:}{2:} Key exists{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnknownKey.ID), "{1:}{2:} Unknown key{:_}")
}

// NewErrKeyExists returns an error with the ErrKeyExists ID.
func NewErrKeyExists(ctx *context.T) error {
	return verror.New(ErrKeyExists, ctx)
}

// NewErrUnknownKey returns an error with the ErrUnknownKey ID.
func NewErrUnknownKey(ctx *context.T) error {
	return verror.New(ErrUnknownKey, ctx)
}
