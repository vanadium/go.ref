// This file was auto-generated by the veyron vdl tool.
// Source: service.vdl

// package benchmark provides simple tools to measure the performance of the
// IPC system.
package benchmarks

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdl "veyron2/vdl"
	_gen_wiretype "veyron2/wiretype"
)

// Benchmark is the interface the client binds and uses.
// Benchmark_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Benchmark_ExcludingUniversal interface {
	// Echo returns the payload that it receives.
	Echo(ctx _gen_context.T, Payload []byte, opts ..._gen_ipc.CallOpt) (reply []byte, err error)
	// EchoStream returns the payload that it receives via the stream.
	EchoStream(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply BenchmarkEchoStreamStream, err error)
}
type Benchmark interface {
	_gen_ipc.UniversalServiceMethods
	Benchmark_ExcludingUniversal
}

// BenchmarkService is the interface the server implements.
type BenchmarkService interface {

	// Echo returns the payload that it receives.
	Echo(context _gen_ipc.ServerContext, Payload []byte) (reply []byte, err error)
	// EchoStream returns the payload that it receives via the stream.
	EchoStream(context _gen_ipc.ServerContext, stream BenchmarkServiceEchoStreamStream) (err error)
}

// BenchmarkEchoStreamStream is the interface for streaming responses of the method
// EchoStream in the service interface Benchmark.
type BenchmarkEchoStreamStream interface {

	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item []byte) error

	// CloseSend indicates to the server that no more items will be sent; server
	// Recv calls will receive io.EOF after all sent items.  Subsequent calls to
	// Send on the client will fail.  This is an optional call - it's used by
	// streaming clients that need the server to receive the io.EOF terminator.
	CloseSend() error

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item []byte, err error)

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the BenchmarkEchoStreamStream interface that is not exported.
type implBenchmarkEchoStreamStream struct {
	clientCall _gen_ipc.Call
}

func (c *implBenchmarkEchoStreamStream) Send(item []byte) error {
	return c.clientCall.Send(item)
}

func (c *implBenchmarkEchoStreamStream) CloseSend() error {
	return c.clientCall.CloseSend()
}

func (c *implBenchmarkEchoStreamStream) Recv() (item []byte, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implBenchmarkEchoStreamStream) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implBenchmarkEchoStreamStream) Cancel() {
	c.clientCall.Cancel()
}

// BenchmarkServiceEchoStreamStream is the interface for streaming responses of the method
// EchoStream in the service interface Benchmark.
type BenchmarkServiceEchoStreamStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item []byte) error

	// Recv fills itemptr with the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item []byte, err error)
}

// Implementation of the BenchmarkServiceEchoStreamStream interface that is not exported.
type implBenchmarkServiceEchoStreamStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implBenchmarkServiceEchoStreamStream) Send(item []byte) error {
	return s.serverCall.Send(item)
}

func (s *implBenchmarkServiceEchoStreamStream) Recv() (item []byte, err error) {
	err = s.serverCall.Recv(&item)
	return
}

// BindBenchmark returns the client stub implementing the Benchmark
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindBenchmark(name string, opts ..._gen_ipc.BindOpt) (Benchmark, error) {
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
			return nil, _gen_vdl.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdl.ErrTooManyOptionsToBind
	}
	stub := &clientStubBenchmark{client: client, name: name}

	return stub, nil
}

// NewServerBenchmark creates a new server stub.
//
// It takes a regular server implementing the BenchmarkService
// interface, and returns a new server stub.
func NewServerBenchmark(server BenchmarkService) interface{} {
	return &ServerStubBenchmark{
		service: server,
	}
}

// clientStubBenchmark implements Benchmark.
type clientStubBenchmark struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubBenchmark) Echo(ctx _gen_context.T, Payload []byte, opts ..._gen_ipc.CallOpt) (reply []byte, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Echo", []interface{}{Payload}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBenchmark) EchoStream(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply BenchmarkEchoStreamStream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "EchoStream", nil, opts...); err != nil {
		return
	}
	reply = &implBenchmarkEchoStreamStream{clientCall: call}
	return
}

func (__gen_c *clientStubBenchmark) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBenchmark) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBenchmark) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubBenchmark wraps a server that implements
// BenchmarkService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubBenchmark struct {
	service BenchmarkService
}

func (__gen_s *ServerStubBenchmark) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Echo":
		return []interface{}{}, nil
	case "EchoStream":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubBenchmark) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Echo"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Payload", Type: 66},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 66},
			{Name: "", Type: 67},
		},
	}
	result.Methods["EchoStream"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 67},
		},
		InStream:  66,
		OutStream: 66,
	}

	result.TypeDefs = []_gen_vdl.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x41, Name: "", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubBenchmark) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubBenchmark) Echo(call _gen_ipc.ServerCall, Payload []byte) (reply []byte, err error) {
	reply, err = __gen_s.service.Echo(call, Payload)
	return
}

func (__gen_s *ServerStubBenchmark) EchoStream(call _gen_ipc.ServerCall) (err error) {
	stream := &implBenchmarkServiceEchoStreamStream{serverCall: call}
	err = __gen_s.service.EchoStream(call, stream)
	return
}
