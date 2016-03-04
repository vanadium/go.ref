// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: errors.vdl

package manager

import (
	// VDL system imports
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"

	// VDL user imports
	"v.io/v23/security"
)

func __VDLEnsureNativeBuilt_errors() {
}

var (
	ErrUnknownProtocol           = verror.Register("v.io/x/ref/runtime/internal/flow/manager.UnknownProtocol", verror.NoRetry, "{1:}{2:} unknown protocol{:3}")
	ErrManagerClosed             = verror.Register("v.io/x/ref/runtime/internal/flow/manager.ManagerClosed", verror.NoRetry, "{1:}{2:} manager is already closed")
	ErrAcceptFailed              = verror.Register("v.io/x/ref/runtime/internal/flow/manager.AcceptFailed", verror.NoRetry, "{1:}{2:} accept failed{:3}")
	ErrCacheClosed               = verror.Register("v.io/x/ref/runtime/internal/flow/manager.CacheClosed", verror.NoRetry, "{1:}{2:} cache is closed")
	ErrConnKilledToFreeResources = verror.Register("v.io/x/ref/runtime/internal/flow/manager.ConnKilledToFreeResources", verror.NoRetry, "{1:}{2:} Connection killed to free resources.")
	ErrInvalidProxyResponse      = verror.Register("v.io/x/ref/runtime/internal/flow/manager.InvalidProxyResponse", verror.NoRetry, "{1:}{2:} Invalid proxy response{:3}")
	ErrManagerDialingSelf        = verror.Register("v.io/x/ref/runtime/internal/flow/manager.ManagerDialingSelf", verror.NoRetry, "{1:}{2:} manager cannot be used to dial itself")
	ErrListeningWithNullRid      = verror.Register("v.io/x/ref/runtime/internal/flow/manager.ListeningWithNullRid", verror.NoRetry, "{1:}{2:} manager cannot listen when created with NullRoutingID")
	ErrProxyResponse             = verror.Register("v.io/x/ref/runtime/internal/flow/manager.ProxyResponse", verror.NoRetry, "{1:}{2:} proxy returned{:3}")
	ErrNoBlessingsForPeer        = verror.Register("v.io/x/ref/runtime/internal/flow/manager.NoBlessingsForPeer", verror.NoRetry, "{1:}{2:} no blessings tagged for peer {3}, rejected:{4}{:5}")
	ErrConnNotInCache            = verror.Register("v.io/x/ref/runtime/internal/flow/manager.ConnNotInCache", verror.NoRetry, "{1:}{2:} connection to {3} not in cache")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnknownProtocol.ID), "{1:}{2:} unknown protocol{:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrManagerClosed.ID), "{1:}{2:} manager is already closed")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrAcceptFailed.ID), "{1:}{2:} accept failed{:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrCacheClosed.ID), "{1:}{2:} cache is closed")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrConnKilledToFreeResources.ID), "{1:}{2:} Connection killed to free resources.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidProxyResponse.ID), "{1:}{2:} Invalid proxy response{:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrManagerDialingSelf.ID), "{1:}{2:} manager cannot be used to dial itself")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrListeningWithNullRid.ID), "{1:}{2:} manager cannot listen when created with NullRoutingID")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrProxyResponse.ID), "{1:}{2:} proxy returned{:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoBlessingsForPeer.ID), "{1:}{2:} no blessings tagged for peer {3}, rejected:{4}{:5}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrConnNotInCache.ID), "{1:}{2:} connection to {3} not in cache")
}

// NewErrUnknownProtocol returns an error with the ErrUnknownProtocol ID.
func NewErrUnknownProtocol(ctx *context.T, protocol string) error {
	return verror.New(ErrUnknownProtocol, ctx, protocol)
}

// NewErrManagerClosed returns an error with the ErrManagerClosed ID.
func NewErrManagerClosed(ctx *context.T) error {
	return verror.New(ErrManagerClosed, ctx)
}

// NewErrAcceptFailed returns an error with the ErrAcceptFailed ID.
func NewErrAcceptFailed(ctx *context.T, err error) error {
	return verror.New(ErrAcceptFailed, ctx, err)
}

// NewErrCacheClosed returns an error with the ErrCacheClosed ID.
func NewErrCacheClosed(ctx *context.T) error {
	return verror.New(ErrCacheClosed, ctx)
}

// NewErrConnKilledToFreeResources returns an error with the ErrConnKilledToFreeResources ID.
func NewErrConnKilledToFreeResources(ctx *context.T) error {
	return verror.New(ErrConnKilledToFreeResources, ctx)
}

// NewErrInvalidProxyResponse returns an error with the ErrInvalidProxyResponse ID.
func NewErrInvalidProxyResponse(ctx *context.T, typ string) error {
	return verror.New(ErrInvalidProxyResponse, ctx, typ)
}

// NewErrManagerDialingSelf returns an error with the ErrManagerDialingSelf ID.
func NewErrManagerDialingSelf(ctx *context.T) error {
	return verror.New(ErrManagerDialingSelf, ctx)
}

// NewErrListeningWithNullRid returns an error with the ErrListeningWithNullRid ID.
func NewErrListeningWithNullRid(ctx *context.T) error {
	return verror.New(ErrListeningWithNullRid, ctx)
}

// NewErrProxyResponse returns an error with the ErrProxyResponse ID.
func NewErrProxyResponse(ctx *context.T, msg string) error {
	return verror.New(ErrProxyResponse, ctx, msg)
}

// NewErrNoBlessingsForPeer returns an error with the ErrNoBlessingsForPeer ID.
func NewErrNoBlessingsForPeer(ctx *context.T, peerNames []string, rejected []security.RejectedBlessing, err error) error {
	return verror.New(ErrNoBlessingsForPeer, ctx, peerNames, rejected, err)
}

// NewErrConnNotInCache returns an error with the ErrConnNotInCache ID.
func NewErrConnNotInCache(ctx *context.T, remote string) error {
	return verror.New(ErrConnNotInCache, ctx, remote)
}
