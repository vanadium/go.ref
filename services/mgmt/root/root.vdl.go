// This file was auto-generated by the veyron vdl tool.
// Source: root.vdl

// Package root contains root process implementation and internal
// interfaces and types used by the implementation.
package root

import (
	// The non-user imports are prefixed with "__" to prevent collisions.
	__veyron2 "veyron.io/veyron/veyron2"
	__context "veyron.io/veyron/veyron2/context"
	__ipc "veyron.io/veyron/veyron2/ipc"
	__vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
	__wiretype "veyron.io/veyron/veyron2/wiretype"
)

// TODO(toddw): Remove this line once the new signature support is done.
// It corrects a bug where __wiretype is unused in VDL pacakges where only
// bootstrap types are used on interfaces.
const _ = __wiretype.TypeIDInvalid

// RootClientMethods is the client interface
// containing Root methods.
//
// Root is an interface to be implemented by a process with root level
// privileges.
type RootClientMethods interface {
	// Reset waits for the given deadline (in milliseconds) and then
	// restars the host node machine.
	Reset(ctx __context.T, Deadline uint64, opts ...__ipc.CallOpt) error
}

// RootClientStub adds universal methods to RootClientMethods.
type RootClientStub interface {
	RootClientMethods
	__ipc.UniversalServiceMethods
}

// RootClient returns a client stub for Root.
func RootClient(name string, opts ...__ipc.BindOpt) RootClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implRootClientStub{name, client}
}

type implRootClientStub struct {
	name   string
	client __ipc.Client
}

func (c implRootClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implRootClientStub) Reset(ctx __context.T, i0 uint64, opts ...__ipc.CallOpt) (err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Reset", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c implRootClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// RootServerMethods is the interface a server writer
// implements for Root.
//
// Root is an interface to be implemented by a process with root level
// privileges.
type RootServerMethods interface {
	// Reset waits for the given deadline (in milliseconds) and then
	// restars the host node machine.
	Reset(ctx __ipc.ServerContext, Deadline uint64) error
}

// RootServerStubMethods is the server interface containing
// Root methods, as expected by ipc.Server.
// There is no difference between this interface and RootServerMethods
// since there are no streaming methods.
type RootServerStubMethods RootServerMethods

// RootServerStub adds universal methods to RootServerStubMethods.
type RootServerStub interface {
	RootServerStubMethods
	// Describe the Root interfaces.
	Describe__() []__ipc.InterfaceDesc
	// Signature will be replaced with Describe__.
	Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error)
}

// RootServer returns a server stub for Root.
// It converts an implementation of RootServerMethods into
// an object that may be used by ipc.Server.
func RootServer(impl RootServerMethods) RootServerStub {
	stub := implRootServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := __ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := __ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implRootServerStub struct {
	impl RootServerMethods
	gs   *__ipc.GlobState
}

func (s implRootServerStub) Reset(ctx __ipc.ServerContext, i0 uint64) error {
	return s.impl.Reset(ctx, i0)
}

func (s implRootServerStub) Globber() *__ipc.GlobState {
	return s.gs
}

func (s implRootServerStub) Describe__() []__ipc.InterfaceDesc {
	return []__ipc.InterfaceDesc{RootDesc}
}

// RootDesc describes the Root interface.
var RootDesc __ipc.InterfaceDesc = descRoot

// descRoot hides the desc to keep godoc clean.
var descRoot = __ipc.InterfaceDesc{
	Name:    "Root",
	PkgPath: "veyron.io/veyron/veyron/services/mgmt/root",
	Doc:     "// Root is an interface to be implemented by a process with root level\n// privileges.",
	Methods: []__ipc.MethodDesc{
		{
			Name: "Reset",
			Doc:  "// Reset waits for the given deadline (in milliseconds) and then\n// restars the host node machine.",
			InArgs: []__ipc.ArgDesc{
				{"Deadline", ``}, // uint64
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
		},
	},
}

func (s implRootServerStub) Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error) {
	// TODO(toddw): Replace with new Describe__ implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}
	result.Methods["Reset"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "Deadline", Type: 53},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []__vdlutil.Any{
		__wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}
