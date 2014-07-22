// This file was auto-generated by the veyron vdl tool.
// Source: tunnel.vdl

package tunnel

import (
	"veyron2/security"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_io "io"
	_gen_veyron2 "veyron2"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
)

type ShellOpts struct {
	UsePty      bool     // Whether to open a pseudo-terminal
	Environment []string // Environment variables to pass to the remote shell.
	Rows        uint32   // Window size.
	Cols        uint32
}

type ClientShellPacket struct {
	// Bytes going to the shell's stdin.
	Stdin []byte
	// A dynamic update of the window size. The default value of 0 means no-change.
	Rows uint32
	Cols uint32
}

type ServerShellPacket struct {
	// Bytes coming from the shell's stdout.
	Stdout []byte
	// Bytes coming from the shell's stderr.
	Stderr []byte
}

// TODO(bprosnitz) Remove this line once signatures are updated to use typevals.
// It corrects a bug where _gen_wiretype is unused in VDL pacakges where only bootstrap types are used on interfaces.
const _ = _gen_wiretype.TypeIDInvalid

// Tunnel is the interface the client binds and uses.
// Tunnel_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Tunnel_ExcludingUniversal interface {
	// The Forward method is used for network forwarding. All the data sent over
	// the byte stream is forwarded to the requested network address and all the
	// data received from that network connection is sent back in the reply
	// stream.
	Forward(ctx _gen_context.T, network string, address string, opts ..._gen_ipc.CallOpt) (reply TunnelForwardStream, err error)
	// The Shell method is used to either run shell commands remotely, or to open
	// an interactive shell. The data received over the byte stream is sent to the
	// shell's stdin, and the data received from the shell's stdout and stderr is
	// sent back in the reply stream. It returns the exit status of the shell
	// command.
	Shell(ctx _gen_context.T, command string, shellOpts ShellOpts, opts ..._gen_ipc.CallOpt) (reply TunnelShellStream, err error)
}
type Tunnel interface {
	_gen_ipc.UniversalServiceMethods
	Tunnel_ExcludingUniversal
}

// TunnelService is the interface the server implements.
type TunnelService interface {

	// The Forward method is used for network forwarding. All the data sent over
	// the byte stream is forwarded to the requested network address and all the
	// data received from that network connection is sent back in the reply
	// stream.
	Forward(context _gen_ipc.ServerContext, network string, address string, stream TunnelServiceForwardStream) (err error)
	// The Shell method is used to either run shell commands remotely, or to open
	// an interactive shell. The data received over the byte stream is sent to the
	// shell's stdin, and the data received from the shell's stdout and stderr is
	// sent back in the reply stream. It returns the exit status of the shell
	// command.
	Shell(context _gen_ipc.ServerContext, command string, shellOpts ShellOpts, stream TunnelServiceShellStream) (reply int32, err error)
}

// TunnelForwardStream is the interface for streaming responses of the method
// Forward in the service interface Tunnel.
type TunnelForwardStream interface {

	// Send places the item onto the output stream, blocking if there is no
	// buffer space available.  Calls to Send after having called CloseSend
	// or Cancel will fail.  Any blocked Send calls will be unblocked upon
	// calling Cancel.
	Send(item []byte) error

	// CloseSend indicates to the server that no more items will be sent;
	// server Recv calls will receive io.EOF after all sent items.  This is
	// an optional call - it's used by streaming clients that need the
	// server to receive the io.EOF terminator before the client calls
	// Finish (for example, if the client needs to continue receiving items
	// from the server after having finished sending).
	// Calls to CloseSend after having called Cancel will fail.
	// Like Send, CloseSend blocks when there's no buffer space available.
	CloseSend() error

	// Advance stages an element so the client can retrieve it
	// with Value.  Advance returns true iff there is an
	// element to retrieve.  The client must call Advance before
	// calling Value.  The client must call Cancel if it does
	// not iterate through all elements (i.e. until Advance
	// returns false).  Advance may block if an element is not
	// immediately available.
	Advance() bool

	// Value returns the element that was staged by Advance.
	// Value may panic if Advance returned false or was not
	// called at all.  Value does not block.
	Value() []byte

	// Err returns a non-nil error iff the stream encountered
	// any errors.  Err does not block.
	Err() error

	// Finish performs the equivalent of CloseSend, then blocks until the server
	// is done, and returns the positional return values for call.
	//
	// If Cancel has been called, Finish will return immediately; the output of
	// Finish could either be an error signalling cancelation, or the correct
	// positional return values from the server depending on the timing of the
	// call.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return a non-EOF error.
	// Finish should be called at most once.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.  It
	// is safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

// Implementation of the TunnelForwardStream interface that is not exported.
type implTunnelForwardStream struct {
	clientCall _gen_ipc.Call
	val        []byte
	err        error
}

func (c *implTunnelForwardStream) Send(item []byte) error {
	return c.clientCall.Send(item)
}

func (c *implTunnelForwardStream) CloseSend() error {
	return c.clientCall.CloseSend()
}

func (c *implTunnelForwardStream) Advance() bool {
	c.err = c.clientCall.Recv(&c.val)
	return c.err == nil
}

func (c *implTunnelForwardStream) Value() []byte {
	return c.val
}

func (c *implTunnelForwardStream) Err() error {
	if c.err == _gen_io.EOF {
		return nil
	}
	return c.err
}

func (c *implTunnelForwardStream) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implTunnelForwardStream) Cancel() {
	c.clientCall.Cancel()
}

// TunnelServiceForwardStream is the interface for streaming responses of the method
// Forward in the service interface Tunnel.
type TunnelServiceForwardStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.  If the client has canceled, an error is returned.
	Send(item []byte) error

	// Advance stages an element so the client can retrieve it
	// with Value.  Advance returns true iff there is an
	// element to retrieve.  The client must call Advance before
	// calling Value.  The client must call Cancel if it does
	// not iterate through all elements (i.e. until Advance
	// returns false).  Advance may block if an element is not
	// immediately available.
	Advance() bool

	// Value returns the element that was staged by Advance.
	// Value may panic if Advance returned false or was not
	// called at all.  Value does not block.
	//
	// In general, Value is undefined if the underlying collection
	// of elements changes while iteration is in progress.  If
	// <DataProvider> supports concurrent modification, it should
	// document its behavior.
	Value() []byte

	// Err returns a non-nil error iff the stream encountered
	// any errors.  Err does not block.
	Err() error
}

// Implementation of the TunnelServiceForwardStream interface that is not exported.
type implTunnelServiceForwardStream struct {
	serverCall _gen_ipc.ServerCall
	val        []byte
	err        error
}

func (s *implTunnelServiceForwardStream) Send(item []byte) error {
	return s.serverCall.Send(item)
}

func (s *implTunnelServiceForwardStream) Advance() bool {
	s.err = s.serverCall.Recv(&s.val)
	return s.err == nil
}

func (s *implTunnelServiceForwardStream) Value() []byte {
	return s.val
}

func (s *implTunnelServiceForwardStream) Err() error {
	if s.err == _gen_io.EOF {
		return nil
	}
	return s.err
}

// TunnelShellStream is the interface for streaming responses of the method
// Shell in the service interface Tunnel.
type TunnelShellStream interface {

	// Send places the item onto the output stream, blocking if there is no
	// buffer space available.  Calls to Send after having called CloseSend
	// or Cancel will fail.  Any blocked Send calls will be unblocked upon
	// calling Cancel.
	Send(item ClientShellPacket) error

	// CloseSend indicates to the server that no more items will be sent;
	// server Recv calls will receive io.EOF after all sent items.  This is
	// an optional call - it's used by streaming clients that need the
	// server to receive the io.EOF terminator before the client calls
	// Finish (for example, if the client needs to continue receiving items
	// from the server after having finished sending).
	// Calls to CloseSend after having called Cancel will fail.
	// Like Send, CloseSend blocks when there's no buffer space available.
	CloseSend() error

	// Advance stages an element so the client can retrieve it
	// with Value.  Advance returns true iff there is an
	// element to retrieve.  The client must call Advance before
	// calling Value.  The client must call Cancel if it does
	// not iterate through all elements (i.e. until Advance
	// returns false).  Advance may block if an element is not
	// immediately available.
	Advance() bool

	// Value returns the element that was staged by Advance.
	// Value may panic if Advance returned false or was not
	// called at all.  Value does not block.
	Value() ServerShellPacket

	// Err returns a non-nil error iff the stream encountered
	// any errors.  Err does not block.
	Err() error

	// Finish performs the equivalent of CloseSend, then blocks until the server
	// is done, and returns the positional return values for call.
	//
	// If Cancel has been called, Finish will return immediately; the output of
	// Finish could either be an error signalling cancelation, or the correct
	// positional return values from the server depending on the timing of the
	// call.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return a non-EOF error.
	// Finish should be called at most once.
	Finish() (reply int32, err error)

	// Cancel cancels the RPC, notifying the server to stop processing.  It
	// is safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

// Implementation of the TunnelShellStream interface that is not exported.
type implTunnelShellStream struct {
	clientCall _gen_ipc.Call
	val        ServerShellPacket
	err        error
}

func (c *implTunnelShellStream) Send(item ClientShellPacket) error {
	return c.clientCall.Send(item)
}

func (c *implTunnelShellStream) CloseSend() error {
	return c.clientCall.CloseSend()
}

func (c *implTunnelShellStream) Advance() bool {
	c.val = ServerShellPacket{}
	c.err = c.clientCall.Recv(&c.val)
	return c.err == nil
}

func (c *implTunnelShellStream) Value() ServerShellPacket {
	return c.val
}

func (c *implTunnelShellStream) Err() error {
	if c.err == _gen_io.EOF {
		return nil
	}
	return c.err
}

func (c *implTunnelShellStream) Finish() (reply int32, err error) {
	if ierr := c.clientCall.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implTunnelShellStream) Cancel() {
	c.clientCall.Cancel()
}

// TunnelServiceShellStream is the interface for streaming responses of the method
// Shell in the service interface Tunnel.
type TunnelServiceShellStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.  If the client has canceled, an error is returned.
	Send(item ServerShellPacket) error

	// Advance stages an element so the client can retrieve it
	// with Value.  Advance returns true iff there is an
	// element to retrieve.  The client must call Advance before
	// calling Value.  The client must call Cancel if it does
	// not iterate through all elements (i.e. until Advance
	// returns false).  Advance may block if an element is not
	// immediately available.
	Advance() bool

	// Value returns the element that was staged by Advance.
	// Value may panic if Advance returned false or was not
	// called at all.  Value does not block.
	//
	// In general, Value is undefined if the underlying collection
	// of elements changes while iteration is in progress.  If
	// <DataProvider> supports concurrent modification, it should
	// document its behavior.
	Value() ClientShellPacket

	// Err returns a non-nil error iff the stream encountered
	// any errors.  Err does not block.
	Err() error
}

// Implementation of the TunnelServiceShellStream interface that is not exported.
type implTunnelServiceShellStream struct {
	serverCall _gen_ipc.ServerCall
	val        ClientShellPacket
	err        error
}

func (s *implTunnelServiceShellStream) Send(item ServerShellPacket) error {
	return s.serverCall.Send(item)
}

func (s *implTunnelServiceShellStream) Advance() bool {
	s.val = ClientShellPacket{}
	s.err = s.serverCall.Recv(&s.val)
	return s.err == nil
}

func (s *implTunnelServiceShellStream) Value() ClientShellPacket {
	return s.val
}

func (s *implTunnelServiceShellStream) Err() error {
	if s.err == _gen_io.EOF {
		return nil
	}
	return s.err
}

// BindTunnel returns the client stub implementing the Tunnel
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindTunnel(name string, opts ..._gen_ipc.BindOpt) (Tunnel, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_veyron2.Runtime:
			client = o.Client()
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubTunnel{client: client, name: name}

	return stub, nil
}

// NewServerTunnel creates a new server stub.
//
// It takes a regular server implementing the TunnelService
// interface, and returns a new server stub.
func NewServerTunnel(server TunnelService) interface{} {
	return &ServerStubTunnel{
		service: server,
	}
}

// clientStubTunnel implements Tunnel.
type clientStubTunnel struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubTunnel) Forward(ctx _gen_context.T, network string, address string, opts ..._gen_ipc.CallOpt) (reply TunnelForwardStream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Forward", []interface{}{network, address}, opts...); err != nil {
		return
	}
	reply = &implTunnelForwardStream{clientCall: call}
	return
}

func (__gen_c *clientStubTunnel) Shell(ctx _gen_context.T, command string, shellOpts ShellOpts, opts ..._gen_ipc.CallOpt) (reply TunnelShellStream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Shell", []interface{}{command, shellOpts}, opts...); err != nil {
		return
	}
	reply = &implTunnelShellStream{clientCall: call}
	return
}

func (__gen_c *clientStubTunnel) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTunnel) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTunnel) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubTunnel wraps a server that implements
// TunnelService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubTunnel struct {
	service TunnelService
}

func (__gen_s *ServerStubTunnel) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Forward":
		return []interface{}{security.Label(4)}, nil
	case "Shell":
		return []interface{}{security.Label(4)}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubTunnel) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Forward"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "network", Type: 3},
			{Name: "address", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
		InStream:  67,
		OutStream: 67,
	}
	result.Methods["Shell"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "command", Type: 3},
			{Name: "shellOpts", Type: 68},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 36},
			{Name: "", Type: 65},
		},
		InStream:  69,
		OutStream: 70,
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x42, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x2, Name: "UsePty"},
				_gen_wiretype.FieldType{Type: 0x3d, Name: "Environment"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "Rows"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "Cols"},
			},
			"veyron/examples/tunnel.ShellOpts", []string(nil)},
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x43, Name: "Stdin"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "Rows"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "Cols"},
			},
			"veyron/examples/tunnel.ClientShellPacket", []string(nil)},
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x43, Name: "Stdout"},
				_gen_wiretype.FieldType{Type: 0x43, Name: "Stderr"},
			},
			"veyron/examples/tunnel.ServerShellPacket", []string(nil)},
	}

	return result, nil
}

func (__gen_s *ServerStubTunnel) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubTunnel) Forward(call _gen_ipc.ServerCall, network string, address string) (err error) {
	stream := &implTunnelServiceForwardStream{serverCall: call}
	err = __gen_s.service.Forward(call, network, address, stream)
	return
}

func (__gen_s *ServerStubTunnel) Shell(call _gen_ipc.ServerCall, command string, shellOpts ShellOpts) (reply int32, err error) {
	stream := &implTunnelServiceShellStream{serverCall: call}
	reply, err = __gen_s.service.Shell(call, command, shellOpts, stream)
	return
}
