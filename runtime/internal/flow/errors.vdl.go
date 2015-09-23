// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: errors.vdl

package flow

import (
	// VDL system imports
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var (
	ErrWrongObjectInContext = verror.Register("v.io/x/ref/runtime/internal/flow.WrongObjectInContext", verror.NoRetry, "{1:}{2:} context passed to method of {3} object, but that object is not attached to the context.")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrWrongObjectInContext.ID), "{1:}{2:} context passed to method of {3} object, but that object is not attached to the context.")
}

// NewErrWrongObjectInContext returns an error with the ErrWrongObjectInContext ID.
func NewErrWrongObjectInContext(ctx *context.T, typ string) error {
	return verror.New(ErrWrongObjectInContext, ctx, typ)
}