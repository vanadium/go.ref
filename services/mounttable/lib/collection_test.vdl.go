// This file was auto-generated by the veyron vdl tool.
// Source: collection_test.vdl

package mounttable

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
)

// TODO(bprosnitz) Remove this line once signatures are updated to use typevals.
// It corrects a bug where _gen_wiretype is unused in VDL pacakges where only bootstrap types are used on interfaces.
const _ = _gen_wiretype.TypeIDInvalid

// Collection is the interface the client binds and uses.
// Collection_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Collection_ExcludingUniversal interface {
	// Export sets the value for a name.  Overwrite controls the behavior when
	// an entry exists, if Overwrite is true, then the binding is replaced,
	// otherwise the call fails with an error.  The Val must be no larger than
	// MaxSize bytes.
	Export(ctx _gen_context.T, Val string, Overwrite bool, opts ..._gen_ipc.CallOpt) (err error)
	// Lookup retrieves the value associated with a name.  Returns an error if
	// there is no such binding.
	Lookup(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []byte, err error)
}
type Collection interface {
	_gen_ipc.UniversalServiceMethods
	Collection_ExcludingUniversal
}

// CollectionService is the interface the server implements.
type CollectionService interface {

	// Export sets the value for a name.  Overwrite controls the behavior when
	// an entry exists, if Overwrite is true, then the binding is replaced,
	// otherwise the call fails with an error.  The Val must be no larger than
	// MaxSize bytes.
	Export(context _gen_ipc.ServerContext, Val string, Overwrite bool) (err error)
	// Lookup retrieves the value associated with a name.  Returns an error if
	// there is no such binding.
	Lookup(context _gen_ipc.ServerContext) (reply []byte, err error)
}

// BindCollection returns the client stub implementing the Collection
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindCollection(name string, opts ..._gen_ipc.BindOpt) (Collection, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		// Do nothing.
	case 1:
		if clientOpt, ok := opts[0].(_gen_ipc.Client); opts[0] == nil || ok {
			client = clientOpt
		} else {
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubCollection{defaultClient: client, name: name}

	return stub, nil
}

// NewServerCollection creates a new server stub.
//
// It takes a regular server implementing the CollectionService
// interface, and returns a new server stub.
func NewServerCollection(server CollectionService) interface{} {
	return &ServerStubCollection{
		service: server,
	}
}

// clientStubCollection implements Collection.
type clientStubCollection struct {
	defaultClient _gen_ipc.Client
	name          string
}

func (__gen_c *clientStubCollection) client(ctx _gen_context.T) _gen_ipc.Client {
	if __gen_c.defaultClient != nil {
		return __gen_c.defaultClient
	}
	return _gen_veyron2.RuntimeFromContext(ctx).Client()
}

func (__gen_c *clientStubCollection) Export(ctx _gen_context.T, Val string, Overwrite bool, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Export", []interface{}{Val, Overwrite}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubCollection) Lookup(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []byte, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Lookup", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubCollection) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubCollection) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubCollection) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubCollection wraps a server that implements
// CollectionService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubCollection struct {
	service CollectionService
}

func (__gen_s *ServerStubCollection) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Export":
		return []interface{}{}, nil
	case "Lookup":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubCollection) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Export"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Val", Type: 3},
			{Name: "Overwrite", Type: 2},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Lookup"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 67},
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x42, Name: "", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubCollection) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubCollection) Export(call _gen_ipc.ServerCall, Val string, Overwrite bool) (err error) {
	err = __gen_s.service.Export(call, Val, Overwrite)
	return
}

func (__gen_s *ServerStubCollection) Lookup(call _gen_ipc.ServerCall) (reply []byte, err error) {
	reply, err = __gen_s.service.Lookup(call)
	return
}
