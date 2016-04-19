// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: xproxy

package xproxy

import (
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Error definitions

var (
	ErrNotListening              = verror.Register("v.io/x/ref/services/xproxy/xproxy.NotListening", verror.NoRetry, "{1:}{2:} Proxy is not listening on any endpoints.")
	ErrUnexpectedMessage         = verror.Register("v.io/x/ref/services/xproxy/xproxy.UnexpectedMessage", verror.NoRetry, "{1:}{2:} Unexpected message of type{:3}")
	ErrFailedToResolveToEndpoint = verror.Register("v.io/x/ref/services/xproxy/xproxy.FailedToResolveToEndpoint", verror.NoRetry, "{1:}{2:} Failed to resolve '{3}' to endpoint")
	ErrProxyAlreadyClosed        = verror.Register("v.io/x/ref/services/xproxy/xproxy.ProxyAlreadyClosed", verror.NoRetry, "{1:}{2:} Proxy has already been closed")
	ErrProxyResponse             = verror.Register("v.io/x/ref/services/xproxy/xproxy.ProxyResponse", verror.NoRetry, "{1:}{2:} proxy returned{:3}")
)

// NewErrNotListening returns an error with the ErrNotListening ID.
func NewErrNotListening(ctx *context.T) error {
	return verror.New(ErrNotListening, ctx)
}

// NewErrUnexpectedMessage returns an error with the ErrUnexpectedMessage ID.
func NewErrUnexpectedMessage(ctx *context.T, msgType string) error {
	return verror.New(ErrUnexpectedMessage, ctx, msgType)
}

// NewErrFailedToResolveToEndpoint returns an error with the ErrFailedToResolveToEndpoint ID.
func NewErrFailedToResolveToEndpoint(ctx *context.T, name string) error {
	return verror.New(ErrFailedToResolveToEndpoint, ctx, name)
}

// NewErrProxyAlreadyClosed returns an error with the ErrProxyAlreadyClosed ID.
func NewErrProxyAlreadyClosed(ctx *context.T) error {
	return verror.New(ErrProxyAlreadyClosed, ctx)
}

// NewErrProxyResponse returns an error with the ErrProxyResponse ID.
func NewErrProxyResponse(ctx *context.T, msg string) error {
	return verror.New(ErrProxyResponse, ctx, msg)
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

	// Set error format strings.
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNotListening.ID), "{1:}{2:} Proxy is not listening on any endpoints.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnexpectedMessage.ID), "{1:}{2:} Unexpected message of type{:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrFailedToResolveToEndpoint.ID), "{1:}{2:} Failed to resolve '{3}' to endpoint")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrProxyAlreadyClosed.ID), "{1:}{2:} Proxy has already been closed")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrProxyResponse.ID), "{1:}{2:} proxy returned{:3}")

	return struct{}{}
}
