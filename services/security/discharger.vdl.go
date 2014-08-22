// This file was auto-generated by the veyron vdl tool.
// Source: discharger.vdl

package security

import (
	"veyron2/security"

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

// Discharger is the interface for obtaining discharges for ThirdPartyCaveats.
// Discharger is the interface the client binds and uses.
// Discharger_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Discharger_ExcludingUniversal interface {
	// Discharge is called by a principal that holds a blessing with a third
	// party caveat and seeks to get a discharge that proves the fulfillment of
	// this caveat.
	//
	// Caveat and Discharge are of type ThirdPartyCaveat and ThirdPartyDischarge
	// respectively. (not enforced here because vdl does not know these types)
	// TODO(ataly,ashankar): Figure out a VDL representation for ThirdPartyCaveat
	// and Discharge and use those here?
	Discharge(ctx _gen_context.T, Caveat _gen_vdlutil.Any, Impetus security.DischargeImpetus, opts ..._gen_ipc.CallOpt) (reply _gen_vdlutil.Any, err error)
}
type Discharger interface {
	_gen_ipc.UniversalServiceMethods
	Discharger_ExcludingUniversal
}

// DischargerService is the interface the server implements.
type DischargerService interface {

	// Discharge is called by a principal that holds a blessing with a third
	// party caveat and seeks to get a discharge that proves the fulfillment of
	// this caveat.
	//
	// Caveat and Discharge are of type ThirdPartyCaveat and ThirdPartyDischarge
	// respectively. (not enforced here because vdl does not know these types)
	// TODO(ataly,ashankar): Figure out a VDL representation for ThirdPartyCaveat
	// and Discharge and use those here?
	Discharge(context _gen_ipc.ServerContext, Caveat _gen_vdlutil.Any, Impetus security.DischargeImpetus) (reply _gen_vdlutil.Any, err error)
}

// BindDischarger returns the client stub implementing the Discharger
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindDischarger(name string, opts ..._gen_ipc.BindOpt) (Discharger, error) {
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
	stub := &clientStubDischarger{defaultClient: client, name: name}

	return stub, nil
}

// NewServerDischarger creates a new server stub.
//
// It takes a regular server implementing the DischargerService
// interface, and returns a new server stub.
func NewServerDischarger(server DischargerService) interface{} {
	return &ServerStubDischarger{
		service: server,
	}
}

// clientStubDischarger implements Discharger.
type clientStubDischarger struct {
	defaultClient _gen_ipc.Client
	name          string
}

func (__gen_c *clientStubDischarger) client(ctx _gen_context.T) _gen_ipc.Client {
	if __gen_c.defaultClient != nil {
		return __gen_c.defaultClient
	}
	return _gen_veyron2.RuntimeFromContext(ctx).Client()
}

func (__gen_c *clientStubDischarger) Discharge(ctx _gen_context.T, Caveat _gen_vdlutil.Any, Impetus security.DischargeImpetus, opts ..._gen_ipc.CallOpt) (reply _gen_vdlutil.Any, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Discharge", []interface{}{Caveat, Impetus}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubDischarger) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubDischarger) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubDischarger) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubDischarger wraps a server that implements
// DischargerService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubDischarger struct {
	service DischargerService
}

func (__gen_s *ServerStubDischarger) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Discharge":
		return []interface{}{security.Label(2)}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubDischarger) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Discharge"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Caveat", Type: 65},
			{Name: "Impetus", Type: 67},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "Discharge", Type: 65},
			{Name: "err", Type: 68},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "anydata", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x41, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x41, Name: "Server"},
				_gen_wiretype.FieldType{Type: 0x3, Name: "Method"},
				_gen_wiretype.FieldType{Type: 0x42, Name: "Arguments"},
			},
			"veyron2/security.DischargeImpetus", []string(nil)},
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubDischarger) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubDischarger) Discharge(call _gen_ipc.ServerCall, Caveat _gen_vdlutil.Any, Impetus security.DischargeImpetus) (reply _gen_vdlutil.Any, err error) {
	reply, err = __gen_s.service.Discharge(call, Caveat, Impetus)
	return
}
