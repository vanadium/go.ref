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

var (
	// ConcurrentTransaction means that the current transaction failed to commit
	// because its read set was invalidated by some other transaction.
	ErrConcurrentTransaction = verror.Register("v.io/x/ref/services/syncbase/store.ConcurrentTransaction", verror.NoRetry, "{1:}{2:} Concurrent transaction{:_}")
	// UnknownKey means the given key does not exist in the store.
	ErrUnknownKey = verror.Register("v.io/x/ref/services/syncbase/store.UnknownKey", verror.NoRetry, "{1:}{2:} Unknown key{:_}")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrConcurrentTransaction.ID), "{1:}{2:} Concurrent transaction{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnknownKey.ID), "{1:}{2:} Unknown key{:_}")
}

// NewErrConcurrentTransaction returns an error with the ErrConcurrentTransaction ID.
func NewErrConcurrentTransaction(ctx *context.T) error {
	return verror.New(ErrConcurrentTransaction, ctx)
}

// NewErrUnknownKey returns an error with the ErrUnknownKey ID.
func NewErrUnknownKey(ctx *context.T) error {
	return verror.New(ErrUnknownKey, ctx)
}
