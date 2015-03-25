// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: arith.vdl

// Package arith is an example of an IdL definition in vanadium.  The syntax for
// IdL files is similar to, but not identical to, Go.  Here are the main
// concepts:
//   * PACKAGES - Just like in Go you must define the package at the beginning
//     of an IdL file, and everything defined in the file is part of this
//     package.  By convention all files in the same dir should be in the same
//     package.
//   * IMPORTS - Just like in Go you can import other idl packages, and you may
//     assign a local package name, or if unspecified the basename of the import
//     path is used as the import package name.
//   * DATA TYPES - Just like in Go you can define data types.  You get most of
//     the primitives (int32, float64, string, etc), the "error" built-in, and a
//     special "any" built-in described below.  In addition you can create
//     composite types like arrays, structs, etc.
//   * CONSTS - Just like in Go you can define constants, and numerics are
//     "infinite precision" within expressions.  Unlike Go numerics must be
//     typed to be used as const definitions or tags.
//   * INTERFACES - Just like in Go you can define interface types, which are
//     just a set of methods.  Interfaces can embed other interfaces.  Unlike
//     Go, you cannot use an interface as a data type; interfaces are purely
//     method sets.
//   * ERRORS - Errors may be defined in IdL files, and unlike Go they work
//     across separate address spaces.
package arith

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/x/ref/lib/vdl/testdata/arith/exp"
	"v.io/x/ref/lib/vdl/testdata/base"
)

// Yes shows that bools may be untyped.
const Yes = true // yes trailing doc

// No shows explicit boolean typing.
const No = false

const Hello = "hello"

// Int32Const shows explicit integer typing.
const Int32Const = int32(123)

// Int64Const shows explicit integer conversion from another type, and referencing
// a constant from another package.
const Int64Const = int64(128)

// FloatConst shows arithmetic expressions may be used.
const FloatConst = float64(2)

// Mask shows bitwise operations.
const Mask = uint64(256)

// ArithClientMethods is the client interface
// containing Arith methods.
//
// Arith is an example of an interface definition for an arithmetic service.
// Things to note:
//   * There must be at least 1 out-arg, and the last out-arg must be error.
type ArithClientMethods interface {
	// Add is a typical method with multiple input and output arguments.
	Add(ctx *context.T, a int32, b int32, opts ...rpc.CallOpt) (int32, error)
	// DivMod shows that runs of args with the same type can use the short form,
	// just like Go.
	DivMod(ctx *context.T, a int32, b int32, opts ...rpc.CallOpt) (quot int32, rem int32, err error)
	// Sub shows that you can use data types defined in other packages.
	Sub(ctx *context.T, args base.Args, opts ...rpc.CallOpt) (int32, error)
	// Mul tries another data type defined in another package.
	Mul(ctx *context.T, nested base.NestedArgs, opts ...rpc.CallOpt) (int32, error)
	// GenError shows that it's fine to have no in args, and no out args other
	// than "error".  In addition GenError shows the usage of tags.  Tags are a
	// sequence of constants.  There's no requirement on uniqueness of types or
	// values, and regular const expressions may also be used.
	GenError(*context.T, ...rpc.CallOpt) error
	// Count shows using only an int32 out-stream type, with no in-stream type.
	Count(ctx *context.T, start int32, opts ...rpc.CallOpt) (ArithCountClientCall, error)
	// StreamingAdd shows a bidirectional stream.
	StreamingAdd(*context.T, ...rpc.CallOpt) (ArithStreamingAddClientCall, error)
	// QuoteAny shows the any built-in type, representing a value of any type.
	QuoteAny(ctx *context.T, a *vdl.Value, opts ...rpc.CallOpt) (*vdl.Value, error)
}

// ArithClientStub adds universal methods to ArithClientMethods.
type ArithClientStub interface {
	ArithClientMethods
	rpc.UniversalServiceMethods
}

// ArithClient returns a client stub for Arith.
func ArithClient(name string, opts ...rpc.BindOpt) ArithClientStub {
	var client rpc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(rpc.Client); ok {
			client = clientOpt
		}
	}
	return implArithClientStub{name, client}
}

type implArithClientStub struct {
	name   string
	client rpc.Client
}

func (c implArithClientStub) c(ctx *context.T) rpc.Client {
	if c.client != nil {
		return c.client
	}
	return v23.GetClient(ctx)
}

func (c implArithClientStub) Add(ctx *context.T, i0 int32, i1 int32, opts ...rpc.CallOpt) (o0 int32, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Add", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implArithClientStub) DivMod(ctx *context.T, i0 int32, i1 int32, opts ...rpc.CallOpt) (o0 int32, o1 int32, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "DivMod", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0, &o1)
	return
}

func (c implArithClientStub) Sub(ctx *context.T, i0 base.Args, opts ...rpc.CallOpt) (o0 int32, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Sub", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implArithClientStub) Mul(ctx *context.T, i0 base.NestedArgs, opts ...rpc.CallOpt) (o0 int32, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Mul", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implArithClientStub) GenError(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "GenError", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implArithClientStub) Count(ctx *context.T, i0 int32, opts ...rpc.CallOpt) (ocall ArithCountClientCall, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Count", []interface{}{i0}, opts...); err != nil {
		return
	}
	ocall = &implArithCountClientCall{ClientCall: call}
	return
}

func (c implArithClientStub) StreamingAdd(ctx *context.T, opts ...rpc.CallOpt) (ocall ArithStreamingAddClientCall, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "StreamingAdd", nil, opts...); err != nil {
		return
	}
	ocall = &implArithStreamingAddClientCall{ClientCall: call}
	return
}

func (c implArithClientStub) QuoteAny(ctx *context.T, i0 *vdl.Value, opts ...rpc.CallOpt) (o0 *vdl.Value, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "QuoteAny", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

// ArithCountClientStream is the client stream for Arith.Count.
type ArithCountClientStream interface {
	// RecvStream returns the receiver side of the Arith.Count client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() int32
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// ArithCountClientCall represents the call returned from Arith.Count.
type ArithCountClientCall interface {
	ArithCountClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implArithCountClientCall struct {
	rpc.ClientCall
	valRecv int32
	errRecv error
}

func (c *implArithCountClientCall) RecvStream() interface {
	Advance() bool
	Value() int32
	Err() error
} {
	return implArithCountClientCallRecv{c}
}

type implArithCountClientCallRecv struct {
	c *implArithCountClientCall
}

func (c implArithCountClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implArithCountClientCallRecv) Value() int32 {
	return c.c.valRecv
}
func (c implArithCountClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implArithCountClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// ArithStreamingAddClientStream is the client stream for Arith.StreamingAdd.
type ArithStreamingAddClientStream interface {
	// RecvStream returns the receiver side of the Arith.StreamingAdd client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() int32
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Arith.StreamingAdd client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item int32) error
		// Close indicates to the server that no more items will be sent;
		// server Recv calls will receive io.EOF after all sent items.
		// This is an optional call - e.g. a client might call Close if it
		// needs to continue receiving items from the server after it's
		// done sending.  Returns errors encountered while closing, or if
		// Close is called after the stream has been canceled.  Like Send,
		// blocks if there is no buffer space available.
		Close() error
	}
}

// ArithStreamingAddClientCall represents the call returned from Arith.StreamingAdd.
type ArithStreamingAddClientCall interface {
	ArithStreamingAddClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() (total int32, err error)
}

type implArithStreamingAddClientCall struct {
	rpc.ClientCall
	valRecv int32
	errRecv error
}

func (c *implArithStreamingAddClientCall) RecvStream() interface {
	Advance() bool
	Value() int32
	Err() error
} {
	return implArithStreamingAddClientCallRecv{c}
}

type implArithStreamingAddClientCallRecv struct {
	c *implArithStreamingAddClientCall
}

func (c implArithStreamingAddClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implArithStreamingAddClientCallRecv) Value() int32 {
	return c.c.valRecv
}
func (c implArithStreamingAddClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implArithStreamingAddClientCall) SendStream() interface {
	Send(item int32) error
	Close() error
} {
	return implArithStreamingAddClientCallSend{c}
}

type implArithStreamingAddClientCallSend struct {
	c *implArithStreamingAddClientCall
}

func (c implArithStreamingAddClientCallSend) Send(item int32) error {
	return c.c.Send(item)
}
func (c implArithStreamingAddClientCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implArithStreamingAddClientCall) Finish() (o0 int32, err error) {
	err = c.ClientCall.Finish(&o0)
	return
}

// ArithServerMethods is the interface a server writer
// implements for Arith.
//
// Arith is an example of an interface definition for an arithmetic service.
// Things to note:
//   * There must be at least 1 out-arg, and the last out-arg must be error.
type ArithServerMethods interface {
	// Add is a typical method with multiple input and output arguments.
	Add(call rpc.ServerCall, a int32, b int32) (int32, error)
	// DivMod shows that runs of args with the same type can use the short form,
	// just like Go.
	DivMod(call rpc.ServerCall, a int32, b int32) (quot int32, rem int32, err error)
	// Sub shows that you can use data types defined in other packages.
	Sub(call rpc.ServerCall, args base.Args) (int32, error)
	// Mul tries another data type defined in another package.
	Mul(call rpc.ServerCall, nested base.NestedArgs) (int32, error)
	// GenError shows that it's fine to have no in args, and no out args other
	// than "error".  In addition GenError shows the usage of tags.  Tags are a
	// sequence of constants.  There's no requirement on uniqueness of types or
	// values, and regular const expressions may also be used.
	GenError(rpc.ServerCall) error
	// Count shows using only an int32 out-stream type, with no in-stream type.
	Count(call ArithCountServerCall, start int32) error
	// StreamingAdd shows a bidirectional stream.
	StreamingAdd(ArithStreamingAddServerCall) (total int32, err error)
	// QuoteAny shows the any built-in type, representing a value of any type.
	QuoteAny(call rpc.ServerCall, a *vdl.Value) (*vdl.Value, error)
}

// ArithServerStubMethods is the server interface containing
// Arith methods, as expected by rpc.Server.
// The only difference between this interface and ArithServerMethods
// is the streaming methods.
type ArithServerStubMethods interface {
	// Add is a typical method with multiple input and output arguments.
	Add(call rpc.ServerCall, a int32, b int32) (int32, error)
	// DivMod shows that runs of args with the same type can use the short form,
	// just like Go.
	DivMod(call rpc.ServerCall, a int32, b int32) (quot int32, rem int32, err error)
	// Sub shows that you can use data types defined in other packages.
	Sub(call rpc.ServerCall, args base.Args) (int32, error)
	// Mul tries another data type defined in another package.
	Mul(call rpc.ServerCall, nested base.NestedArgs) (int32, error)
	// GenError shows that it's fine to have no in args, and no out args other
	// than "error".  In addition GenError shows the usage of tags.  Tags are a
	// sequence of constants.  There's no requirement on uniqueness of types or
	// values, and regular const expressions may also be used.
	GenError(rpc.ServerCall) error
	// Count shows using only an int32 out-stream type, with no in-stream type.
	Count(call *ArithCountServerCallStub, start int32) error
	// StreamingAdd shows a bidirectional stream.
	StreamingAdd(*ArithStreamingAddServerCallStub) (total int32, err error)
	// QuoteAny shows the any built-in type, representing a value of any type.
	QuoteAny(call rpc.ServerCall, a *vdl.Value) (*vdl.Value, error)
}

// ArithServerStub adds universal methods to ArithServerStubMethods.
type ArithServerStub interface {
	ArithServerStubMethods
	// Describe the Arith interfaces.
	Describe__() []rpc.InterfaceDesc
}

// ArithServer returns a server stub for Arith.
// It converts an implementation of ArithServerMethods into
// an object that may be used by rpc.Server.
func ArithServer(impl ArithServerMethods) ArithServerStub {
	stub := implArithServerStub{
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

type implArithServerStub struct {
	impl ArithServerMethods
	gs   *rpc.GlobState
}

func (s implArithServerStub) Add(call rpc.ServerCall, i0 int32, i1 int32) (int32, error) {
	return s.impl.Add(call, i0, i1)
}

func (s implArithServerStub) DivMod(call rpc.ServerCall, i0 int32, i1 int32) (int32, int32, error) {
	return s.impl.DivMod(call, i0, i1)
}

func (s implArithServerStub) Sub(call rpc.ServerCall, i0 base.Args) (int32, error) {
	return s.impl.Sub(call, i0)
}

func (s implArithServerStub) Mul(call rpc.ServerCall, i0 base.NestedArgs) (int32, error) {
	return s.impl.Mul(call, i0)
}

func (s implArithServerStub) GenError(call rpc.ServerCall) error {
	return s.impl.GenError(call)
}

func (s implArithServerStub) Count(call *ArithCountServerCallStub, i0 int32) error {
	return s.impl.Count(call, i0)
}

func (s implArithServerStub) StreamingAdd(call *ArithStreamingAddServerCallStub) (int32, error) {
	return s.impl.StreamingAdd(call)
}

func (s implArithServerStub) QuoteAny(call rpc.ServerCall, i0 *vdl.Value) (*vdl.Value, error) {
	return s.impl.QuoteAny(call, i0)
}

func (s implArithServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implArithServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{ArithDesc}
}

// ArithDesc describes the Arith interface.
var ArithDesc rpc.InterfaceDesc = descArith

// descArith hides the desc to keep godoc clean.
var descArith = rpc.InterfaceDesc{
	Name:    "Arith",
	PkgPath: "v.io/x/ref/lib/vdl/testdata/arith",
	Doc:     "// Arith is an example of an interface definition for an arithmetic service.\n// Things to note:\n//   * There must be at least 1 out-arg, and the last out-arg must be error.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Add",
			Doc:  "// Add is a typical method with multiple input and output arguments.",
			InArgs: []rpc.ArgDesc{
				{"a", ``}, // int32
				{"b", ``}, // int32
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // int32
			},
		},
		{
			Name: "DivMod",
			Doc:  "// DivMod shows that runs of args with the same type can use the short form,\n// just like Go.",
			InArgs: []rpc.ArgDesc{
				{"a", ``}, // int32
				{"b", ``}, // int32
			},
			OutArgs: []rpc.ArgDesc{
				{"quot", ``}, // int32
				{"rem", ``},  // int32
			},
		},
		{
			Name: "Sub",
			Doc:  "// Sub shows that you can use data types defined in other packages.",
			InArgs: []rpc.ArgDesc{
				{"args", ``}, // base.Args
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // int32
			},
		},
		{
			Name: "Mul",
			Doc:  "// Mul tries another data type defined in another package.",
			InArgs: []rpc.ArgDesc{
				{"nested", ``}, // base.NestedArgs
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // int32
			},
		},
		{
			Name: "GenError",
			Doc:  "// GenError shows that it's fine to have no in args, and no out args other\n// than \"error\".  In addition GenError shows the usage of tags.  Tags are a\n// sequence of constants.  There's no requirement on uniqueness of types or\n// values, and regular const expressions may also be used.",
			Tags: []*vdl.Value{vdl.ValueOf("foo"), vdl.ValueOf("barz"), vdl.ValueOf("hello"), vdl.ValueOf(int32(129)), vdl.ValueOf(uint64(36))},
		},
		{
			Name: "Count",
			Doc:  "// Count shows using only an int32 out-stream type, with no in-stream type.",
			InArgs: []rpc.ArgDesc{
				{"start", ``}, // int32
			},
		},
		{
			Name: "StreamingAdd",
			Doc:  "// StreamingAdd shows a bidirectional stream.",
			OutArgs: []rpc.ArgDesc{
				{"total", ``}, // int32
			},
		},
		{
			Name: "QuoteAny",
			Doc:  "// QuoteAny shows the any built-in type, representing a value of any type.",
			InArgs: []rpc.ArgDesc{
				{"a", ``}, // *vdl.Value
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // *vdl.Value
			},
		},
	},
}

// ArithCountServerStream is the server stream for Arith.Count.
type ArithCountServerStream interface {
	// SendStream returns the send side of the Arith.Count server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item int32) error
	}
}

// ArithCountServerCall represents the context passed to Arith.Count.
type ArithCountServerCall interface {
	rpc.ServerCall
	ArithCountServerStream
}

// ArithCountServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements ArithCountServerCall.
type ArithCountServerCallStub struct {
	rpc.StreamServerCall
}

// Init initializes ArithCountServerCallStub from rpc.StreamServerCall.
func (s *ArithCountServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the Arith.Count server stream.
func (s *ArithCountServerCallStub) SendStream() interface {
	Send(item int32) error
} {
	return implArithCountServerCallSend{s}
}

type implArithCountServerCallSend struct {
	s *ArithCountServerCallStub
}

func (s implArithCountServerCallSend) Send(item int32) error {
	return s.s.Send(item)
}

// ArithStreamingAddServerStream is the server stream for Arith.StreamingAdd.
type ArithStreamingAddServerStream interface {
	// RecvStream returns the receiver side of the Arith.StreamingAdd server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() int32
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Arith.StreamingAdd server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item int32) error
	}
}

// ArithStreamingAddServerCall represents the context passed to Arith.StreamingAdd.
type ArithStreamingAddServerCall interface {
	rpc.ServerCall
	ArithStreamingAddServerStream
}

// ArithStreamingAddServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements ArithStreamingAddServerCall.
type ArithStreamingAddServerCallStub struct {
	rpc.StreamServerCall
	valRecv int32
	errRecv error
}

// Init initializes ArithStreamingAddServerCallStub from rpc.StreamServerCall.
func (s *ArithStreamingAddServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// RecvStream returns the receiver side of the Arith.StreamingAdd server stream.
func (s *ArithStreamingAddServerCallStub) RecvStream() interface {
	Advance() bool
	Value() int32
	Err() error
} {
	return implArithStreamingAddServerCallRecv{s}
}

type implArithStreamingAddServerCallRecv struct {
	s *ArithStreamingAddServerCallStub
}

func (s implArithStreamingAddServerCallRecv) Advance() bool {
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implArithStreamingAddServerCallRecv) Value() int32 {
	return s.s.valRecv
}
func (s implArithStreamingAddServerCallRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Arith.StreamingAdd server stream.
func (s *ArithStreamingAddServerCallStub) SendStream() interface {
	Send(item int32) error
} {
	return implArithStreamingAddServerCallSend{s}
}

type implArithStreamingAddServerCallSend struct {
	s *ArithStreamingAddServerCallStub
}

func (s implArithStreamingAddServerCallSend) Send(item int32) error {
	return s.s.Send(item)
}

// CalculatorClientMethods is the client interface
// containing Calculator methods.
type CalculatorClientMethods interface {
	// Arith is an example of an interface definition for an arithmetic service.
	// Things to note:
	//   * There must be at least 1 out-arg, and the last out-arg must be error.
	ArithClientMethods
	// AdvancedMath is an interface for more advanced math than arith.  It embeds
	// interfaces defined both in the same file and in an external package; and in
	// turn it is embedded by arith.Calculator (which is in the same package but
	// different file) to verify that embedding works in all these scenarios.
	AdvancedMathClientMethods
	On(*context.T, ...rpc.CallOpt) error  // On turns the calculator on.
	Off(*context.T, ...rpc.CallOpt) error // Off turns the calculator off.
}

// CalculatorClientStub adds universal methods to CalculatorClientMethods.
type CalculatorClientStub interface {
	CalculatorClientMethods
	rpc.UniversalServiceMethods
}

// CalculatorClient returns a client stub for Calculator.
func CalculatorClient(name string, opts ...rpc.BindOpt) CalculatorClientStub {
	var client rpc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(rpc.Client); ok {
			client = clientOpt
		}
	}
	return implCalculatorClientStub{name, client, ArithClient(name, client), AdvancedMathClient(name, client)}
}

type implCalculatorClientStub struct {
	name   string
	client rpc.Client

	ArithClientStub
	AdvancedMathClientStub
}

func (c implCalculatorClientStub) c(ctx *context.T) rpc.Client {
	if c.client != nil {
		return c.client
	}
	return v23.GetClient(ctx)
}

func (c implCalculatorClientStub) On(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "On", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implCalculatorClientStub) Off(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Off", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

// CalculatorServerMethods is the interface a server writer
// implements for Calculator.
type CalculatorServerMethods interface {
	// Arith is an example of an interface definition for an arithmetic service.
	// Things to note:
	//   * There must be at least 1 out-arg, and the last out-arg must be error.
	ArithServerMethods
	// AdvancedMath is an interface for more advanced math than arith.  It embeds
	// interfaces defined both in the same file and in an external package; and in
	// turn it is embedded by arith.Calculator (which is in the same package but
	// different file) to verify that embedding works in all these scenarios.
	AdvancedMathServerMethods
	On(rpc.ServerCall) error  // On turns the calculator on.
	Off(rpc.ServerCall) error // Off turns the calculator off.
}

// CalculatorServerStubMethods is the server interface containing
// Calculator methods, as expected by rpc.Server.
// The only difference between this interface and CalculatorServerMethods
// is the streaming methods.
type CalculatorServerStubMethods interface {
	// Arith is an example of an interface definition for an arithmetic service.
	// Things to note:
	//   * There must be at least 1 out-arg, and the last out-arg must be error.
	ArithServerStubMethods
	// AdvancedMath is an interface for more advanced math than arith.  It embeds
	// interfaces defined both in the same file and in an external package; and in
	// turn it is embedded by arith.Calculator (which is in the same package but
	// different file) to verify that embedding works in all these scenarios.
	AdvancedMathServerStubMethods
	On(rpc.ServerCall) error  // On turns the calculator on.
	Off(rpc.ServerCall) error // Off turns the calculator off.
}

// CalculatorServerStub adds universal methods to CalculatorServerStubMethods.
type CalculatorServerStub interface {
	CalculatorServerStubMethods
	// Describe the Calculator interfaces.
	Describe__() []rpc.InterfaceDesc
}

// CalculatorServer returns a server stub for Calculator.
// It converts an implementation of CalculatorServerMethods into
// an object that may be used by rpc.Server.
func CalculatorServer(impl CalculatorServerMethods) CalculatorServerStub {
	stub := implCalculatorServerStub{
		impl:                   impl,
		ArithServerStub:        ArithServer(impl),
		AdvancedMathServerStub: AdvancedMathServer(impl),
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

type implCalculatorServerStub struct {
	impl CalculatorServerMethods
	ArithServerStub
	AdvancedMathServerStub
	gs *rpc.GlobState
}

func (s implCalculatorServerStub) On(call rpc.ServerCall) error {
	return s.impl.On(call)
}

func (s implCalculatorServerStub) Off(call rpc.ServerCall) error {
	return s.impl.Off(call)
}

func (s implCalculatorServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implCalculatorServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{CalculatorDesc, ArithDesc, AdvancedMathDesc, TrigonometryDesc, exp.ExpDesc}
}

// CalculatorDesc describes the Calculator interface.
var CalculatorDesc rpc.InterfaceDesc = descCalculator

// descCalculator hides the desc to keep godoc clean.
var descCalculator = rpc.InterfaceDesc{
	Name:    "Calculator",
	PkgPath: "v.io/x/ref/lib/vdl/testdata/arith",
	Embeds: []rpc.EmbedDesc{
		{"Arith", "v.io/x/ref/lib/vdl/testdata/arith", "// Arith is an example of an interface definition for an arithmetic service.\n// Things to note:\n//   * There must be at least 1 out-arg, and the last out-arg must be error."},
		{"AdvancedMath", "v.io/x/ref/lib/vdl/testdata/arith", "// AdvancedMath is an interface for more advanced math than arith.  It embeds\n// interfaces defined both in the same file and in an external package; and in\n// turn it is embedded by arith.Calculator (which is in the same package but\n// different file) to verify that embedding works in all these scenarios."},
	},
	Methods: []rpc.MethodDesc{
		{
			Name: "On",
		},
		{
			Name: "Off",
			Tags: []*vdl.Value{vdl.ValueOf("offtag")},
		},
	},
}
