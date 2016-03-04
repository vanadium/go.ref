// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: errors.vdl

package xwebsocket

import (
	// VDL system imports
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

func __VDLEnsureNativeBuilt_errors() {
}

var (
	ErrListenerClosed     = verror.Register("v.io/x/ref/runtime/internal/lib/xwebsocket.ListenerClosed", verror.NoRetry, "{1:}{2:} listener is already closed.")
	ErrListenCalledInNaCl = verror.Register("v.io/x/ref/runtime/internal/lib/xwebsocket.ListenCalledInNaCl", verror.NoRetry, "{1:}{2:} Listen cannot be called in NaCl code.")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrListenerClosed.ID), "{1:}{2:} listener is already closed.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrListenCalledInNaCl.ID), "{1:}{2:} Listen cannot be called in NaCl code.")
}

// NewErrListenerClosed returns an error with the ErrListenerClosed ID.
func NewErrListenerClosed(ctx *context.T) error {
	return verror.New(ErrListenerClosed, ctx)
}

// NewErrListenCalledInNaCl returns an error with the ErrListenCalledInNaCl ID.
func NewErrListenCalledInNaCl(ctx *context.T) error {
	return verror.New(ErrListenCalledInNaCl, ctx)
}
