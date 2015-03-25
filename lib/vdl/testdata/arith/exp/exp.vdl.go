// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: exp.vdl

// Package exp is used to test that embedding interfaces works across packages.
// The arith.Calculator vdl interface embeds the Exp interface.
package exp

import (
	// VDL system imports
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
)

// ExpClientMethods is the client interface
// containing Exp methods.
type ExpClientMethods interface {
	Exp(ctx *context.T, x float64, opts ...rpc.CallOpt) (float64, error)
}

// ExpClientStub adds universal methods to ExpClientMethods.
type ExpClientStub interface {
	ExpClientMethods
	rpc.UniversalServiceMethods
}

// ExpClient returns a client stub for Exp.
func ExpClient(name string, opts ...rpc.BindOpt) ExpClientStub {
	var client rpc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(rpc.Client); ok {
			client = clientOpt
		}
	}
	return implExpClientStub{name, client}
}

type implExpClientStub struct {
	name   string
	client rpc.Client
}

func (c implExpClientStub) c(ctx *context.T) rpc.Client {
	if c.client != nil {
		return c.client
	}
	return v23.GetClient(ctx)
}

func (c implExpClientStub) Exp(ctx *context.T, i0 float64, opts ...rpc.CallOpt) (o0 float64, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Exp", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

// ExpServerMethods is the interface a server writer
// implements for Exp.
type ExpServerMethods interface {
	Exp(call rpc.ServerCall, x float64) (float64, error)
}

// ExpServerStubMethods is the server interface containing
// Exp methods, as expected by rpc.Server.
// There is no difference between this interface and ExpServerMethods
// since there are no streaming methods.
type ExpServerStubMethods ExpServerMethods

// ExpServerStub adds universal methods to ExpServerStubMethods.
type ExpServerStub interface {
	ExpServerStubMethods
	// Describe the Exp interfaces.
	Describe__() []rpc.InterfaceDesc
}

// ExpServer returns a server stub for Exp.
// It converts an implementation of ExpServerMethods into
// an object that may be used by rpc.Server.
func ExpServer(impl ExpServerMethods) ExpServerStub {
	stub := implExpServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implExpServerStub struct {
	impl ExpServerMethods
	gs   *rpc.GlobState
}

func (s implExpServerStub) Exp(call rpc.ServerCall, i0 float64) (float64, error) {
	return s.impl.Exp(call, i0)
}

func (s implExpServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implExpServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{ExpDesc}
}

// ExpDesc describes the Exp interface.
var ExpDesc rpc.InterfaceDesc = descExp

// descExp hides the desc to keep godoc clean.
var descExp = rpc.InterfaceDesc{
	Name:    "Exp",
	PkgPath: "v.io/x/ref/lib/vdl/testdata/arith/exp",
	Methods: []rpc.MethodDesc{
		{
			Name: "Exp",
			InArgs: []rpc.ArgDesc{
				{"x", ``}, // float64
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // float64
			},
		},
	},
}
