// This file was auto-generated by the veyron vdl tool.
// Source: server.vdl

package server

import (
	// VDL system imports
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/vdl"
	"v.io/v23/verror"

	// VDL user imports
	"v.io/v23/security"
	"v.io/x/ref/services/wsprd/principal"
)

type SecurityCall struct {
	Method                string
	Suffix                string
	MethodTags            []*vdl.Value
	LocalBlessings        principal.BlessingsHandle
	LocalBlessingStrings  []string
	RemoteBlessings       principal.BlessingsHandle
	RemoteBlessingStrings []string
	LocalEndpoint         string
	RemoteEndpoint        string
}

func (SecurityCall) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wsprd/ipc/server.SecurityCall"
}) {
}

type CaveatValidationRequest struct {
	Call     SecurityCall
	CallSide security.CallSide
	Cavs     [][]security.Caveat
}

func (CaveatValidationRequest) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wsprd/ipc/server.CaveatValidationRequest"
}) {
}

type CaveatValidationResponse struct {
	Results []error
}

func (CaveatValidationResponse) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wsprd/ipc/server.CaveatValidationResponse"
}) {
}

func init() {
	vdl.Register((*SecurityCall)(nil))
	vdl.Register((*CaveatValidationRequest)(nil))
	vdl.Register((*CaveatValidationResponse)(nil))
}

var (
	ErrCaveatValidationTimeout                 = verror.Register("v.io/x/ref/services/wsprd/ipc/server.CaveatValidationTimeout", verror.NoRetry, "{1:}{2:} Caveat validation has timed out")
	ErrInvalidValidationResponseFromJavascript = verror.Register("v.io/x/ref/services/wsprd/ipc/server.InvalidValidationResponseFromJavascript", verror.NoRetry, "{1:}{2:} Invalid validation response from javascript")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrCaveatValidationTimeout.ID), "{1:}{2:} Caveat validation has timed out")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidValidationResponseFromJavascript.ID), "{1:}{2:} Invalid validation response from javascript")
}

// NewErrCaveatValidationTimeout returns an error with the ErrCaveatValidationTimeout ID.
func NewErrCaveatValidationTimeout(ctx *context.T) error {
	return verror.New(ErrCaveatValidationTimeout, ctx)
}

// NewErrInvalidValidationResponseFromJavascript returns an error with the ErrInvalidValidationResponseFromJavascript ID.
func NewErrInvalidValidationResponseFromJavascript(ctx *context.T) error {
	return verror.New(ErrInvalidValidationResponseFromJavascript, ctx)
}
