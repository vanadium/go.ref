// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: tunnel.vdl

// Package tunnel describes a service that can be used to create a
// network tunnel from the client to the server.
package tunnel

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/services/security/access"
)

type ShellOpts struct {
	UsePty      bool       // Whether to open a pseudo-terminal.
	Environment []string   // Environment variables to pass to the remote shell.
	WinSize     WindowSize // The size of the window.
}

func (ShellOpts) __VDLReflect(struct {
	Name string "v.io/x/ref/examples/tunnel.ShellOpts"
}) {
}

type WindowSize struct {
	Rows uint16
	Cols uint16
}

func (WindowSize) __VDLReflect(struct {
	Name string "v.io/x/ref/examples/tunnel.WindowSize"
}) {
}

type (
	// ClientShellPacket represents any single field of the ClientShellPacket union type.
	ClientShellPacket interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the ClientShellPacket union type.
		__VDLReflect(__ClientShellPacketReflect)
	}
	// ClientShellPacketStdin represents field Stdin of the ClientShellPacket union type.
	//
	// Bytes going to the shell's stdin.
	ClientShellPacketStdin struct{ Value []byte }
	// ClientShellPacketEndOfFile represents field EndOfFile of the ClientShellPacket union type.
	//
	// Indicates that stdin should be closed. The presence of this field indicates
	// EOF. Its actual value is ignored.
	ClientShellPacketEndOfFile struct{ Value unused }
	// ClientShellPacketWinSize represents field WinSize of the ClientShellPacket union type.
	//
	// A dynamic update of the window size.
	ClientShellPacketWinSize struct{ Value WindowSize }
	// __ClientShellPacketReflect describes the ClientShellPacket union type.
	__ClientShellPacketReflect struct {
		Name  string "v.io/x/ref/examples/tunnel.ClientShellPacket"
		Type  ClientShellPacket
		Union struct {
			Stdin     ClientShellPacketStdin
			EndOfFile ClientShellPacketEndOfFile
			WinSize   ClientShellPacketWinSize
		}
	}
)

func (x ClientShellPacketStdin) Index() int                              { return 0 }
func (x ClientShellPacketStdin) Interface() interface{}                  { return x.Value }
func (x ClientShellPacketStdin) Name() string                            { return "Stdin" }
func (x ClientShellPacketStdin) __VDLReflect(__ClientShellPacketReflect) {}

func (x ClientShellPacketEndOfFile) Index() int                              { return 1 }
func (x ClientShellPacketEndOfFile) Interface() interface{}                  { return x.Value }
func (x ClientShellPacketEndOfFile) Name() string                            { return "EndOfFile" }
func (x ClientShellPacketEndOfFile) __VDLReflect(__ClientShellPacketReflect) {}

func (x ClientShellPacketWinSize) Index() int                              { return 2 }
func (x ClientShellPacketWinSize) Interface() interface{}                  { return x.Value }
func (x ClientShellPacketWinSize) Name() string                            { return "WinSize" }
func (x ClientShellPacketWinSize) __VDLReflect(__ClientShellPacketReflect) {}

type unused struct {
}

func (unused) __VDLReflect(struct {
	Name string "v.io/x/ref/examples/tunnel.unused"
}) {
}

type (
	// ServerShellPacket represents any single field of the ServerShellPacket union type.
	ServerShellPacket interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the ServerShellPacket union type.
		__VDLReflect(__ServerShellPacketReflect)
	}
	// ServerShellPacketStdout represents field Stdout of the ServerShellPacket union type.
	//
	// Bytes coming from the shell's stdout.
	ServerShellPacketStdout struct{ Value []byte }
	// ServerShellPacketStderr represents field Stderr of the ServerShellPacket union type.
	//
	// Bytes coming from the shell's stderr.
	ServerShellPacketStderr struct{ Value []byte }
	// __ServerShellPacketReflect describes the ServerShellPacket union type.
	__ServerShellPacketReflect struct {
		Name  string "v.io/x/ref/examples/tunnel.ServerShellPacket"
		Type  ServerShellPacket
		Union struct {
			Stdout ServerShellPacketStdout
			Stderr ServerShellPacketStderr
		}
	}
)

func (x ServerShellPacketStdout) Index() int                              { return 0 }
func (x ServerShellPacketStdout) Interface() interface{}                  { return x.Value }
func (x ServerShellPacketStdout) Name() string                            { return "Stdout" }
func (x ServerShellPacketStdout) __VDLReflect(__ServerShellPacketReflect) {}

func (x ServerShellPacketStderr) Index() int                              { return 1 }
func (x ServerShellPacketStderr) Interface() interface{}                  { return x.Value }
func (x ServerShellPacketStderr) Name() string                            { return "Stderr" }
func (x ServerShellPacketStderr) __VDLReflect(__ServerShellPacketReflect) {}

func init() {
	vdl.Register((*ShellOpts)(nil))
	vdl.Register((*WindowSize)(nil))
	vdl.Register((*ClientShellPacket)(nil))
	vdl.Register((*unused)(nil))
	vdl.Register((*ServerShellPacket)(nil))
}

// TunnelClientMethods is the client interface
// containing Tunnel methods.
type TunnelClientMethods interface {
	// The Forward method is used for network forwarding. All the data sent over
	// the byte stream is forwarded to the requested network address and all the
	// data received from that network connection is sent back in the reply
	// stream.
	Forward(ctx *context.T, network string, address string, opts ...rpc.CallOpt) (TunnelForwardClientCall, error)
	// The Shell method is used to either run shell commands remotely, or to open
	// an interactive shell. The data received over the byte stream is sent to the
	// shell's stdin, and the data received from the shell's stdout and stderr is
	// sent back in the reply stream. It returns the exit status of the shell
	// command.
	Shell(ctx *context.T, command string, shellOpts ShellOpts, opts ...rpc.CallOpt) (TunnelShellClientCall, error)
}

// TunnelClientStub adds universal methods to TunnelClientMethods.
type TunnelClientStub interface {
	TunnelClientMethods
	rpc.UniversalServiceMethods
}

// TunnelClient returns a client stub for Tunnel.
func TunnelClient(name string, opts ...rpc.BindOpt) TunnelClientStub {
	var client rpc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(rpc.Client); ok {
			client = clientOpt
		}
	}
	return implTunnelClientStub{name, client}
}

type implTunnelClientStub struct {
	name   string
	client rpc.Client
}

func (c implTunnelClientStub) c(ctx *context.T) rpc.Client {
	if c.client != nil {
		return c.client
	}
	return v23.GetClient(ctx)
}

func (c implTunnelClientStub) Forward(ctx *context.T, i0 string, i1 string, opts ...rpc.CallOpt) (ocall TunnelForwardClientCall, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Forward", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	ocall = &implTunnelForwardClientCall{ClientCall: call}
	return
}

func (c implTunnelClientStub) Shell(ctx *context.T, i0 string, i1 ShellOpts, opts ...rpc.CallOpt) (ocall TunnelShellClientCall, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Shell", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	ocall = &implTunnelShellClientCall{ClientCall: call}
	return
}

// TunnelForwardClientStream is the client stream for Tunnel.Forward.
type TunnelForwardClientStream interface {
	// RecvStream returns the receiver side of the Tunnel.Forward client stream.
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
	// SendStream returns the send side of the Tunnel.Forward client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item []byte) error
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

// TunnelForwardClientCall represents the call returned from Tunnel.Forward.
type TunnelForwardClientCall interface {
	TunnelForwardClientStream
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

type implTunnelForwardClientCall struct {
	rpc.ClientCall
	valRecv []byte
	errRecv error
}

func (c *implTunnelForwardClientCall) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implTunnelForwardClientCallRecv{c}
}

type implTunnelForwardClientCallRecv struct {
	c *implTunnelForwardClientCall
}

func (c implTunnelForwardClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implTunnelForwardClientCallRecv) Value() []byte {
	return c.c.valRecv
}
func (c implTunnelForwardClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implTunnelForwardClientCall) SendStream() interface {
	Send(item []byte) error
	Close() error
} {
	return implTunnelForwardClientCallSend{c}
}

type implTunnelForwardClientCallSend struct {
	c *implTunnelForwardClientCall
}

func (c implTunnelForwardClientCallSend) Send(item []byte) error {
	return c.c.Send(item)
}
func (c implTunnelForwardClientCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implTunnelForwardClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// TunnelShellClientStream is the client stream for Tunnel.Shell.
type TunnelShellClientStream interface {
	// RecvStream returns the receiver side of the Tunnel.Shell client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() ServerShellPacket
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Tunnel.Shell client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item ClientShellPacket) error
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

// TunnelShellClientCall represents the call returned from Tunnel.Shell.
type TunnelShellClientCall interface {
	TunnelShellClientStream
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
	Finish() (int32, error)
}

type implTunnelShellClientCall struct {
	rpc.ClientCall
	valRecv ServerShellPacket
	errRecv error
}

func (c *implTunnelShellClientCall) RecvStream() interface {
	Advance() bool
	Value() ServerShellPacket
	Err() error
} {
	return implTunnelShellClientCallRecv{c}
}

type implTunnelShellClientCallRecv struct {
	c *implTunnelShellClientCall
}

func (c implTunnelShellClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implTunnelShellClientCallRecv) Value() ServerShellPacket {
	return c.c.valRecv
}
func (c implTunnelShellClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implTunnelShellClientCall) SendStream() interface {
	Send(item ClientShellPacket) error
	Close() error
} {
	return implTunnelShellClientCallSend{c}
}

type implTunnelShellClientCallSend struct {
	c *implTunnelShellClientCall
}

func (c implTunnelShellClientCallSend) Send(item ClientShellPacket) error {
	return c.c.Send(item)
}
func (c implTunnelShellClientCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implTunnelShellClientCall) Finish() (o0 int32, err error) {
	err = c.ClientCall.Finish(&o0)
	return
}

// TunnelServerMethods is the interface a server writer
// implements for Tunnel.
type TunnelServerMethods interface {
	// The Forward method is used for network forwarding. All the data sent over
	// the byte stream is forwarded to the requested network address and all the
	// data received from that network connection is sent back in the reply
	// stream.
	Forward(call TunnelForwardServerCall, network string, address string) error
	// The Shell method is used to either run shell commands remotely, or to open
	// an interactive shell. The data received over the byte stream is sent to the
	// shell's stdin, and the data received from the shell's stdout and stderr is
	// sent back in the reply stream. It returns the exit status of the shell
	// command.
	Shell(call TunnelShellServerCall, command string, shellOpts ShellOpts) (int32, error)
}

// TunnelServerStubMethods is the server interface containing
// Tunnel methods, as expected by rpc.Server.
// The only difference between this interface and TunnelServerMethods
// is the streaming methods.
type TunnelServerStubMethods interface {
	// The Forward method is used for network forwarding. All the data sent over
	// the byte stream is forwarded to the requested network address and all the
	// data received from that network connection is sent back in the reply
	// stream.
	Forward(call *TunnelForwardServerCallStub, network string, address string) error
	// The Shell method is used to either run shell commands remotely, or to open
	// an interactive shell. The data received over the byte stream is sent to the
	// shell's stdin, and the data received from the shell's stdout and stderr is
	// sent back in the reply stream. It returns the exit status of the shell
	// command.
	Shell(call *TunnelShellServerCallStub, command string, shellOpts ShellOpts) (int32, error)
}

// TunnelServerStub adds universal methods to TunnelServerStubMethods.
type TunnelServerStub interface {
	TunnelServerStubMethods
	// Describe the Tunnel interfaces.
	Describe__() []rpc.InterfaceDesc
}

// TunnelServer returns a server stub for Tunnel.
// It converts an implementation of TunnelServerMethods into
// an object that may be used by rpc.Server.
func TunnelServer(impl TunnelServerMethods) TunnelServerStub {
	stub := implTunnelServerStub{
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

type implTunnelServerStub struct {
	impl TunnelServerMethods
	gs   *rpc.GlobState
}

func (s implTunnelServerStub) Forward(call *TunnelForwardServerCallStub, i0 string, i1 string) error {
	return s.impl.Forward(call, i0, i1)
}

func (s implTunnelServerStub) Shell(call *TunnelShellServerCallStub, i0 string, i1 ShellOpts) (int32, error) {
	return s.impl.Shell(call, i0, i1)
}

func (s implTunnelServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implTunnelServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{TunnelDesc}
}

// TunnelDesc describes the Tunnel interface.
var TunnelDesc rpc.InterfaceDesc = descTunnel

// descTunnel hides the desc to keep godoc clean.
var descTunnel = rpc.InterfaceDesc{
	Name:    "Tunnel",
	PkgPath: "v.io/x/ref/examples/tunnel",
	Methods: []rpc.MethodDesc{
		{
			Name: "Forward",
			Doc:  "// The Forward method is used for network forwarding. All the data sent over\n// the byte stream is forwarded to the requested network address and all the\n// data received from that network connection is sent back in the reply\n// stream.",
			InArgs: []rpc.ArgDesc{
				{"network", ``}, // string
				{"address", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Admin"))},
		},
		{
			Name: "Shell",
			Doc:  "// The Shell method is used to either run shell commands remotely, or to open\n// an interactive shell. The data received over the byte stream is sent to the\n// shell's stdin, and the data received from the shell's stdout and stderr is\n// sent back in the reply stream. It returns the exit status of the shell\n// command.",
			InArgs: []rpc.ArgDesc{
				{"command", ``},   // string
				{"shellOpts", ``}, // ShellOpts
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // int32
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Admin"))},
		},
	},
}

// TunnelForwardServerStream is the server stream for Tunnel.Forward.
type TunnelForwardServerStream interface {
	// RecvStream returns the receiver side of the Tunnel.Forward server stream.
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
	// SendStream returns the send side of the Tunnel.Forward server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item []byte) error
	}
}

// TunnelForwardServerCall represents the context passed to Tunnel.Forward.
type TunnelForwardServerCall interface {
	rpc.ServerCall
	TunnelForwardServerStream
}

// TunnelForwardServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements TunnelForwardServerCall.
type TunnelForwardServerCallStub struct {
	rpc.StreamServerCall
	valRecv []byte
	errRecv error
}

// Init initializes TunnelForwardServerCallStub from rpc.StreamServerCall.
func (s *TunnelForwardServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// RecvStream returns the receiver side of the Tunnel.Forward server stream.
func (s *TunnelForwardServerCallStub) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implTunnelForwardServerCallRecv{s}
}

type implTunnelForwardServerCallRecv struct {
	s *TunnelForwardServerCallStub
}

func (s implTunnelForwardServerCallRecv) Advance() bool {
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implTunnelForwardServerCallRecv) Value() []byte {
	return s.s.valRecv
}
func (s implTunnelForwardServerCallRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Tunnel.Forward server stream.
func (s *TunnelForwardServerCallStub) SendStream() interface {
	Send(item []byte) error
} {
	return implTunnelForwardServerCallSend{s}
}

type implTunnelForwardServerCallSend struct {
	s *TunnelForwardServerCallStub
}

func (s implTunnelForwardServerCallSend) Send(item []byte) error {
	return s.s.Send(item)
}

// TunnelShellServerStream is the server stream for Tunnel.Shell.
type TunnelShellServerStream interface {
	// RecvStream returns the receiver side of the Tunnel.Shell server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() ClientShellPacket
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Tunnel.Shell server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item ServerShellPacket) error
	}
}

// TunnelShellServerCall represents the context passed to Tunnel.Shell.
type TunnelShellServerCall interface {
	rpc.ServerCall
	TunnelShellServerStream
}

// TunnelShellServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements TunnelShellServerCall.
type TunnelShellServerCallStub struct {
	rpc.StreamServerCall
	valRecv ClientShellPacket
	errRecv error
}

// Init initializes TunnelShellServerCallStub from rpc.StreamServerCall.
func (s *TunnelShellServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// RecvStream returns the receiver side of the Tunnel.Shell server stream.
func (s *TunnelShellServerCallStub) RecvStream() interface {
	Advance() bool
	Value() ClientShellPacket
	Err() error
} {
	return implTunnelShellServerCallRecv{s}
}

type implTunnelShellServerCallRecv struct {
	s *TunnelShellServerCallStub
}

func (s implTunnelShellServerCallRecv) Advance() bool {
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implTunnelShellServerCallRecv) Value() ClientShellPacket {
	return s.s.valRecv
}
func (s implTunnelShellServerCallRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Tunnel.Shell server stream.
func (s *TunnelShellServerCallStub) SendStream() interface {
	Send(item ServerShellPacket) error
} {
	return implTunnelShellServerCallSend{s}
}

type implTunnelShellServerCallSend struct {
	s *TunnelShellServerCallStub
}

func (s implTunnelShellServerCallSend) Send(item ServerShellPacket) error {
	return s.s.Send(item)
}
