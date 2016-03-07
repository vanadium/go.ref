// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: stress.vdl

package stress

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/security/access"
)

type SumArg struct {
	ABool        bool
	AInt64       int64
	AListOfBytes []byte
}

func (SumArg) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/runtime/internal/rpc/stress.SumArg"`
}) {
}

func (m *SumArg) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_stress_v_io_x_ref_runtime_internal_rpc_stress_SumArg == nil || __VDLTypestress0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("ABool")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromBool(bool(m.ABool), vdl.BoolType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("AInt64")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromInt(int64(m.AInt64), vdl.Int64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("AListOfBytes")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget7.FromBytes([]byte(m.AListOfBytes), __VDLTypestress1); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *SumArg) MakeVDLTarget() vdl.Target {
	return nil
}

type SumStats struct {
	SumCount       uint64
	SumStreamCount uint64
	BytesRecv      uint64
	BytesSent      uint64
}

func (SumStats) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/runtime/internal/rpc/stress.SumStats"`
}) {
}

func (m *SumStats) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_stress_v_io_x_ref_runtime_internal_rpc_stress_SumStats == nil || __VDLTypestress2 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("SumCount")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromUint(uint64(m.SumCount), vdl.Uint64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("SumStreamCount")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromUint(uint64(m.SumStreamCount), vdl.Uint64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("BytesRecv")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget7.FromUint(uint64(m.BytesRecv), vdl.Uint64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("BytesSent")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget9.FromUint(uint64(m.BytesSent), vdl.Uint64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *SumStats) MakeVDLTarget() vdl.Target {
	return nil
}

func init() {
	vdl.Register((*SumArg)(nil))
	vdl.Register((*SumStats)(nil))
}

var __VDLTypestress0 *vdl.Type = vdl.TypeOf((*SumArg)(nil))
var __VDLTypestress2 *vdl.Type = vdl.TypeOf((*SumStats)(nil))
var __VDLTypestress1 *vdl.Type = vdl.TypeOf([]byte(nil))
var __VDLType_stress_v_io_x_ref_runtime_internal_rpc_stress_SumArg *vdl.Type = vdl.TypeOf(SumArg{})
var __VDLType_stress_v_io_x_ref_runtime_internal_rpc_stress_SumStats *vdl.Type = vdl.TypeOf(SumStats{})

func __VDLEnsureNativeBuilt_stress() {
}

// StressClientMethods is the client interface
// containing Stress methods.
type StressClientMethods interface {
	// Echo returns the payload that it receives.
	Echo(_ *context.T, Payload []byte, _ ...rpc.CallOpt) ([]byte, error)
	// Do returns the checksum of the payload that it receives.
	Sum(_ *context.T, arg SumArg, _ ...rpc.CallOpt) ([]byte, error)
	// DoStream returns the checksum of the payload that it receives via the stream.
	SumStream(*context.T, ...rpc.CallOpt) (StressSumStreamClientCall, error)
	// GetSumStats returns the stats on the Sum calls that the server received.
	GetSumStats(*context.T, ...rpc.CallOpt) (SumStats, error)
	// Stop stops the server.
	Stop(*context.T, ...rpc.CallOpt) error
}

// StressClientStub adds universal methods to StressClientMethods.
type StressClientStub interface {
	StressClientMethods
	rpc.UniversalServiceMethods
}

// StressClient returns a client stub for Stress.
func StressClient(name string) StressClientStub {
	return implStressClientStub{name}
}

type implStressClientStub struct {
	name string
}

func (c implStressClientStub) Echo(ctx *context.T, i0 []byte, opts ...rpc.CallOpt) (o0 []byte, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Echo", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implStressClientStub) Sum(ctx *context.T, i0 SumArg, opts ...rpc.CallOpt) (o0 []byte, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Sum", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implStressClientStub) SumStream(ctx *context.T, opts ...rpc.CallOpt) (ocall StressSumStreamClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "SumStream", nil, opts...); err != nil {
		return
	}
	ocall = &implStressSumStreamClientCall{ClientCall: call}
	return
}

func (c implStressClientStub) GetSumStats(ctx *context.T, opts ...rpc.CallOpt) (o0 SumStats, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "GetSumStats", nil, []interface{}{&o0}, opts...)
	return
}

func (c implStressClientStub) Stop(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Stop", nil, nil, opts...)
	return
}

// StressSumStreamClientStream is the client stream for Stress.SumStream.
type StressSumStreamClientStream interface {
	// RecvStream returns the receiver side of the Stress.SumStream client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() []byte
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Stress.SumStream client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item SumArg) error
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

// StressSumStreamClientCall represents the call returned from Stress.SumStream.
type StressSumStreamClientCall interface {
	StressSumStreamClientStream
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
	Finish() error
}

type implStressSumStreamClientCall struct {
	rpc.ClientCall
	valRecv []byte
	errRecv error
}

func (c *implStressSumStreamClientCall) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implStressSumStreamClientCallRecv{c}
}

type implStressSumStreamClientCallRecv struct {
	c *implStressSumStreamClientCall
}

func (c implStressSumStreamClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implStressSumStreamClientCallRecv) Value() []byte {
	return c.c.valRecv
}
func (c implStressSumStreamClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implStressSumStreamClientCall) SendStream() interface {
	Send(item SumArg) error
	Close() error
} {
	return implStressSumStreamClientCallSend{c}
}

type implStressSumStreamClientCallSend struct {
	c *implStressSumStreamClientCall
}

func (c implStressSumStreamClientCallSend) Send(item SumArg) error {
	return c.c.Send(item)
}
func (c implStressSumStreamClientCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implStressSumStreamClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// StressServerMethods is the interface a server writer
// implements for Stress.
type StressServerMethods interface {
	// Echo returns the payload that it receives.
	Echo(_ *context.T, _ rpc.ServerCall, Payload []byte) ([]byte, error)
	// Do returns the checksum of the payload that it receives.
	Sum(_ *context.T, _ rpc.ServerCall, arg SumArg) ([]byte, error)
	// DoStream returns the checksum of the payload that it receives via the stream.
	SumStream(*context.T, StressSumStreamServerCall) error
	// GetSumStats returns the stats on the Sum calls that the server received.
	GetSumStats(*context.T, rpc.ServerCall) (SumStats, error)
	// Stop stops the server.
	Stop(*context.T, rpc.ServerCall) error
}

// StressServerStubMethods is the server interface containing
// Stress methods, as expected by rpc.Server.
// The only difference between this interface and StressServerMethods
// is the streaming methods.
type StressServerStubMethods interface {
	// Echo returns the payload that it receives.
	Echo(_ *context.T, _ rpc.ServerCall, Payload []byte) ([]byte, error)
	// Do returns the checksum of the payload that it receives.
	Sum(_ *context.T, _ rpc.ServerCall, arg SumArg) ([]byte, error)
	// DoStream returns the checksum of the payload that it receives via the stream.
	SumStream(*context.T, *StressSumStreamServerCallStub) error
	// GetSumStats returns the stats on the Sum calls that the server received.
	GetSumStats(*context.T, rpc.ServerCall) (SumStats, error)
	// Stop stops the server.
	Stop(*context.T, rpc.ServerCall) error
}

// StressServerStub adds universal methods to StressServerStubMethods.
type StressServerStub interface {
	StressServerStubMethods
	// Describe the Stress interfaces.
	Describe__() []rpc.InterfaceDesc
}

// StressServer returns a server stub for Stress.
// It converts an implementation of StressServerMethods into
// an object that may be used by rpc.Server.
func StressServer(impl StressServerMethods) StressServerStub {
	stub := implStressServerStub{
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

type implStressServerStub struct {
	impl StressServerMethods
	gs   *rpc.GlobState
}

func (s implStressServerStub) Echo(ctx *context.T, call rpc.ServerCall, i0 []byte) ([]byte, error) {
	return s.impl.Echo(ctx, call, i0)
}

func (s implStressServerStub) Sum(ctx *context.T, call rpc.ServerCall, i0 SumArg) ([]byte, error) {
	return s.impl.Sum(ctx, call, i0)
}

func (s implStressServerStub) SumStream(ctx *context.T, call *StressSumStreamServerCallStub) error {
	return s.impl.SumStream(ctx, call)
}

func (s implStressServerStub) GetSumStats(ctx *context.T, call rpc.ServerCall) (SumStats, error) {
	return s.impl.GetSumStats(ctx, call)
}

func (s implStressServerStub) Stop(ctx *context.T, call rpc.ServerCall) error {
	return s.impl.Stop(ctx, call)
}

func (s implStressServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implStressServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{StressDesc}
}

// StressDesc describes the Stress interface.
var StressDesc rpc.InterfaceDesc = descStress

// descStress hides the desc to keep godoc clean.
var descStress = rpc.InterfaceDesc{
	Name:    "Stress",
	PkgPath: "v.io/x/ref/runtime/internal/rpc/stress",
	Methods: []rpc.MethodDesc{
		{
			Name: "Echo",
			Doc:  "// Echo returns the payload that it receives.",
			InArgs: []rpc.ArgDesc{
				{"Payload", ``}, // []byte
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []byte
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "Sum",
			Doc:  "// Do returns the checksum of the payload that it receives.",
			InArgs: []rpc.ArgDesc{
				{"arg", ``}, // SumArg
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []byte
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "SumStream",
			Doc:  "// DoStream returns the checksum of the payload that it receives via the stream.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "GetSumStats",
			Doc:  "// GetSumStats returns the stats on the Sum calls that the server received.",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // SumStats
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "Stop",
			Doc:  "// Stop stops the server.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Admin"))},
		},
	},
}

// StressSumStreamServerStream is the server stream for Stress.SumStream.
type StressSumStreamServerStream interface {
	// RecvStream returns the receiver side of the Stress.SumStream server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() SumArg
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Stress.SumStream server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item []byte) error
	}
}

// StressSumStreamServerCall represents the context passed to Stress.SumStream.
type StressSumStreamServerCall interface {
	rpc.ServerCall
	StressSumStreamServerStream
}

// StressSumStreamServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements StressSumStreamServerCall.
type StressSumStreamServerCallStub struct {
	rpc.StreamServerCall
	valRecv SumArg
	errRecv error
}

// Init initializes StressSumStreamServerCallStub from rpc.StreamServerCall.
func (s *StressSumStreamServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// RecvStream returns the receiver side of the Stress.SumStream server stream.
func (s *StressSumStreamServerCallStub) RecvStream() interface {
	Advance() bool
	Value() SumArg
	Err() error
} {
	return implStressSumStreamServerCallRecv{s}
}

type implStressSumStreamServerCallRecv struct {
	s *StressSumStreamServerCallStub
}

func (s implStressSumStreamServerCallRecv) Advance() bool {
	s.s.valRecv = SumArg{}
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implStressSumStreamServerCallRecv) Value() SumArg {
	return s.s.valRecv
}
func (s implStressSumStreamServerCallRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Stress.SumStream server stream.
func (s *StressSumStreamServerCallStub) SendStream() interface {
	Send(item []byte) error
} {
	return implStressSumStreamServerCallSend{s}
}

type implStressSumStreamServerCallSend struct {
	s *StressSumStreamServerCallStub
}

func (s implStressSumStreamServerCallSend) Send(item []byte) error {
	return s.s.Send(item)
}
