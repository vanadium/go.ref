// This file was auto-generated by the vanadium vdl tool.
// Source: wire.vdl

package server

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"

	// VDL user imports
	"v.io/v23/security"
)

// AgentClientMethods is the client interface
// containing Agent methods.
type AgentClientMethods interface {
	Bless(ctx *context.T, key []byte, wit security.Blessings, extension string, caveat security.Caveat, additionalCaveats []security.Caveat, opts ...rpc.CallOpt) (security.Blessings, error)
	BlessSelf(ctx *context.T, name string, caveats []security.Caveat, opts ...rpc.CallOpt) (security.Blessings, error)
	Sign(ctx *context.T, message []byte, opts ...rpc.CallOpt) (security.Signature, error)
	MintDischarge(ctx *context.T, forCaveat security.Caveat, caveatOnDischarge security.Caveat, additionalCaveatsOnDischarge []security.Caveat, opts ...rpc.CallOpt) (security.Discharge, error)
	PublicKey(*context.T, ...rpc.CallOpt) ([]byte, error)
	BlessingsByName(ctx *context.T, name security.BlessingPattern, opts ...rpc.CallOpt) ([]security.Blessings, error)
	BlessingsInfo(ctx *context.T, blessings security.Blessings, opts ...rpc.CallOpt) (map[string][]security.Caveat, error)
	AddToRoots(ctx *context.T, blessing security.Blessings, opts ...rpc.CallOpt) error
	BlessingStoreSet(ctx *context.T, blessings security.Blessings, forPeers security.BlessingPattern, opts ...rpc.CallOpt) (security.Blessings, error)
	BlessingStoreForPeer(ctx *context.T, peerBlessings []string, opts ...rpc.CallOpt) (security.Blessings, error)
	BlessingStoreSetDefault(ctx *context.T, blessings security.Blessings, opts ...rpc.CallOpt) error
	BlessingStoreDefault(*context.T, ...rpc.CallOpt) (security.Blessings, error)
	BlessingStorePeerBlessings(*context.T, ...rpc.CallOpt) (map[security.BlessingPattern]security.Blessings, error)
	BlessingStoreDebugString(*context.T, ...rpc.CallOpt) (string, error)
	BlessingRootsAdd(ctx *context.T, root []byte, pattern security.BlessingPattern, opts ...rpc.CallOpt) error
	BlessingRootsRecognized(ctx *context.T, root []byte, blessing string, opts ...rpc.CallOpt) error
	BlessingRootsDebugString(*context.T, ...rpc.CallOpt) (string, error)
	// Clients using caching should call NotifyWhenChanged upon connecting to
	// the server. The server will stream back values whenever the client should
	// flush the cache. The streamed value is arbitrary, simply flush whenever
	// recieving a new item.
	NotifyWhenChanged(*context.T, ...rpc.CallOpt) (AgentNotifyWhenChangedClientCall, error)
}

// AgentClientStub adds universal methods to AgentClientMethods.
type AgentClientStub interface {
	AgentClientMethods
	rpc.UniversalServiceMethods
}

// AgentClient returns a client stub for Agent.
func AgentClient(name string, opts ...rpc.BindOpt) AgentClientStub {
	var client rpc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(rpc.Client); ok {
			client = clientOpt
		}
	}
	return implAgentClientStub{name, client}
}

type implAgentClientStub struct {
	name   string
	client rpc.Client
}

func (c implAgentClientStub) c(ctx *context.T) rpc.Client {
	if c.client != nil {
		return c.client
	}
	return v23.GetClient(ctx)
}

func (c implAgentClientStub) Bless(ctx *context.T, i0 []byte, i1 security.Blessings, i2 string, i3 security.Caveat, i4 []security.Caveat, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Bless", []interface{}{i0, i1, i2, i3, i4}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) BlessSelf(ctx *context.T, i0 string, i1 []security.Caveat, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessSelf", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) Sign(ctx *context.T, i0 []byte, opts ...rpc.CallOpt) (o0 security.Signature, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Sign", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) MintDischarge(ctx *context.T, i0 security.Caveat, i1 security.Caveat, i2 []security.Caveat, opts ...rpc.CallOpt) (o0 security.Discharge, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "MintDischarge", []interface{}{i0, i1, i2}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) PublicKey(ctx *context.T, opts ...rpc.CallOpt) (o0 []byte, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "PublicKey", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) BlessingsByName(ctx *context.T, i0 security.BlessingPattern, opts ...rpc.CallOpt) (o0 []security.Blessings, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessingsByName", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) BlessingsInfo(ctx *context.T, i0 security.Blessings, opts ...rpc.CallOpt) (o0 map[string][]security.Caveat, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessingsInfo", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) AddToRoots(ctx *context.T, i0 security.Blessings, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "AddToRoots", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implAgentClientStub) BlessingStoreSet(ctx *context.T, i0 security.Blessings, i1 security.BlessingPattern, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessingStoreSet", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) BlessingStoreForPeer(ctx *context.T, i0 []string, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessingStoreForPeer", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) BlessingStoreSetDefault(ctx *context.T, i0 security.Blessings, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessingStoreSetDefault", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implAgentClientStub) BlessingStoreDefault(ctx *context.T, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessingStoreDefault", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) BlessingStorePeerBlessings(ctx *context.T, opts ...rpc.CallOpt) (o0 map[security.BlessingPattern]security.Blessings, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessingStorePeerBlessings", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) BlessingStoreDebugString(ctx *context.T, opts ...rpc.CallOpt) (o0 string, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessingStoreDebugString", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) BlessingRootsAdd(ctx *context.T, i0 []byte, i1 security.BlessingPattern, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessingRootsAdd", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implAgentClientStub) BlessingRootsRecognized(ctx *context.T, i0 []byte, i1 string, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessingRootsRecognized", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implAgentClientStub) BlessingRootsDebugString(ctx *context.T, opts ...rpc.CallOpt) (o0 string, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "BlessingRootsDebugString", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implAgentClientStub) NotifyWhenChanged(ctx *context.T, opts ...rpc.CallOpt) (ocall AgentNotifyWhenChangedClientCall, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "NotifyWhenChanged", nil, opts...); err != nil {
		return
	}
	ocall = &implAgentNotifyWhenChangedClientCall{ClientCall: call}
	return
}

// AgentNotifyWhenChangedClientStream is the client stream for Agent.NotifyWhenChanged.
type AgentNotifyWhenChangedClientStream interface {
	// RecvStream returns the receiver side of the Agent.NotifyWhenChanged client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() bool
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// AgentNotifyWhenChangedClientCall represents the call returned from Agent.NotifyWhenChanged.
type AgentNotifyWhenChangedClientCall interface {
	AgentNotifyWhenChangedClientStream
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

type implAgentNotifyWhenChangedClientCall struct {
	rpc.ClientCall
	valRecv bool
	errRecv error
}

func (c *implAgentNotifyWhenChangedClientCall) RecvStream() interface {
	Advance() bool
	Value() bool
	Err() error
} {
	return implAgentNotifyWhenChangedClientCallRecv{c}
}

type implAgentNotifyWhenChangedClientCallRecv struct {
	c *implAgentNotifyWhenChangedClientCall
}

func (c implAgentNotifyWhenChangedClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implAgentNotifyWhenChangedClientCallRecv) Value() bool {
	return c.c.valRecv
}
func (c implAgentNotifyWhenChangedClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implAgentNotifyWhenChangedClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// AgentServerMethods is the interface a server writer
// implements for Agent.
type AgentServerMethods interface {
	Bless(call rpc.ServerCall, key []byte, wit security.Blessings, extension string, caveat security.Caveat, additionalCaveats []security.Caveat) (security.Blessings, error)
	BlessSelf(call rpc.ServerCall, name string, caveats []security.Caveat) (security.Blessings, error)
	Sign(call rpc.ServerCall, message []byte) (security.Signature, error)
	MintDischarge(call rpc.ServerCall, forCaveat security.Caveat, caveatOnDischarge security.Caveat, additionalCaveatsOnDischarge []security.Caveat) (security.Discharge, error)
	PublicKey(rpc.ServerCall) ([]byte, error)
	BlessingsByName(call rpc.ServerCall, name security.BlessingPattern) ([]security.Blessings, error)
	BlessingsInfo(call rpc.ServerCall, blessings security.Blessings) (map[string][]security.Caveat, error)
	AddToRoots(call rpc.ServerCall, blessing security.Blessings) error
	BlessingStoreSet(call rpc.ServerCall, blessings security.Blessings, forPeers security.BlessingPattern) (security.Blessings, error)
	BlessingStoreForPeer(call rpc.ServerCall, peerBlessings []string) (security.Blessings, error)
	BlessingStoreSetDefault(call rpc.ServerCall, blessings security.Blessings) error
	BlessingStoreDefault(rpc.ServerCall) (security.Blessings, error)
	BlessingStorePeerBlessings(rpc.ServerCall) (map[security.BlessingPattern]security.Blessings, error)
	BlessingStoreDebugString(rpc.ServerCall) (string, error)
	BlessingRootsAdd(call rpc.ServerCall, root []byte, pattern security.BlessingPattern) error
	BlessingRootsRecognized(call rpc.ServerCall, root []byte, blessing string) error
	BlessingRootsDebugString(rpc.ServerCall) (string, error)
	// Clients using caching should call NotifyWhenChanged upon connecting to
	// the server. The server will stream back values whenever the client should
	// flush the cache. The streamed value is arbitrary, simply flush whenever
	// recieving a new item.
	NotifyWhenChanged(AgentNotifyWhenChangedServerCall) error
}

// AgentServerStubMethods is the server interface containing
// Agent methods, as expected by rpc.Server.
// The only difference between this interface and AgentServerMethods
// is the streaming methods.
type AgentServerStubMethods interface {
	Bless(call rpc.ServerCall, key []byte, wit security.Blessings, extension string, caveat security.Caveat, additionalCaveats []security.Caveat) (security.Blessings, error)
	BlessSelf(call rpc.ServerCall, name string, caveats []security.Caveat) (security.Blessings, error)
	Sign(call rpc.ServerCall, message []byte) (security.Signature, error)
	MintDischarge(call rpc.ServerCall, forCaveat security.Caveat, caveatOnDischarge security.Caveat, additionalCaveatsOnDischarge []security.Caveat) (security.Discharge, error)
	PublicKey(rpc.ServerCall) ([]byte, error)
	BlessingsByName(call rpc.ServerCall, name security.BlessingPattern) ([]security.Blessings, error)
	BlessingsInfo(call rpc.ServerCall, blessings security.Blessings) (map[string][]security.Caveat, error)
	AddToRoots(call rpc.ServerCall, blessing security.Blessings) error
	BlessingStoreSet(call rpc.ServerCall, blessings security.Blessings, forPeers security.BlessingPattern) (security.Blessings, error)
	BlessingStoreForPeer(call rpc.ServerCall, peerBlessings []string) (security.Blessings, error)
	BlessingStoreSetDefault(call rpc.ServerCall, blessings security.Blessings) error
	BlessingStoreDefault(rpc.ServerCall) (security.Blessings, error)
	BlessingStorePeerBlessings(rpc.ServerCall) (map[security.BlessingPattern]security.Blessings, error)
	BlessingStoreDebugString(rpc.ServerCall) (string, error)
	BlessingRootsAdd(call rpc.ServerCall, root []byte, pattern security.BlessingPattern) error
	BlessingRootsRecognized(call rpc.ServerCall, root []byte, blessing string) error
	BlessingRootsDebugString(rpc.ServerCall) (string, error)
	// Clients using caching should call NotifyWhenChanged upon connecting to
	// the server. The server will stream back values whenever the client should
	// flush the cache. The streamed value is arbitrary, simply flush whenever
	// recieving a new item.
	NotifyWhenChanged(*AgentNotifyWhenChangedServerCallStub) error
}

// AgentServerStub adds universal methods to AgentServerStubMethods.
type AgentServerStub interface {
	AgentServerStubMethods
	// Describe the Agent interfaces.
	Describe__() []rpc.InterfaceDesc
}

// AgentServer returns a server stub for Agent.
// It converts an implementation of AgentServerMethods into
// an object that may be used by rpc.Server.
func AgentServer(impl AgentServerMethods) AgentServerStub {
	stub := implAgentServerStub{
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

type implAgentServerStub struct {
	impl AgentServerMethods
	gs   *rpc.GlobState
}

func (s implAgentServerStub) Bless(call rpc.ServerCall, i0 []byte, i1 security.Blessings, i2 string, i3 security.Caveat, i4 []security.Caveat) (security.Blessings, error) {
	return s.impl.Bless(call, i0, i1, i2, i3, i4)
}

func (s implAgentServerStub) BlessSelf(call rpc.ServerCall, i0 string, i1 []security.Caveat) (security.Blessings, error) {
	return s.impl.BlessSelf(call, i0, i1)
}

func (s implAgentServerStub) Sign(call rpc.ServerCall, i0 []byte) (security.Signature, error) {
	return s.impl.Sign(call, i0)
}

func (s implAgentServerStub) MintDischarge(call rpc.ServerCall, i0 security.Caveat, i1 security.Caveat, i2 []security.Caveat) (security.Discharge, error) {
	return s.impl.MintDischarge(call, i0, i1, i2)
}

func (s implAgentServerStub) PublicKey(call rpc.ServerCall) ([]byte, error) {
	return s.impl.PublicKey(call)
}

func (s implAgentServerStub) BlessingsByName(call rpc.ServerCall, i0 security.BlessingPattern) ([]security.Blessings, error) {
	return s.impl.BlessingsByName(call, i0)
}

func (s implAgentServerStub) BlessingsInfo(call rpc.ServerCall, i0 security.Blessings) (map[string][]security.Caveat, error) {
	return s.impl.BlessingsInfo(call, i0)
}

func (s implAgentServerStub) AddToRoots(call rpc.ServerCall, i0 security.Blessings) error {
	return s.impl.AddToRoots(call, i0)
}

func (s implAgentServerStub) BlessingStoreSet(call rpc.ServerCall, i0 security.Blessings, i1 security.BlessingPattern) (security.Blessings, error) {
	return s.impl.BlessingStoreSet(call, i0, i1)
}

func (s implAgentServerStub) BlessingStoreForPeer(call rpc.ServerCall, i0 []string) (security.Blessings, error) {
	return s.impl.BlessingStoreForPeer(call, i0)
}

func (s implAgentServerStub) BlessingStoreSetDefault(call rpc.ServerCall, i0 security.Blessings) error {
	return s.impl.BlessingStoreSetDefault(call, i0)
}

func (s implAgentServerStub) BlessingStoreDefault(call rpc.ServerCall) (security.Blessings, error) {
	return s.impl.BlessingStoreDefault(call)
}

func (s implAgentServerStub) BlessingStorePeerBlessings(call rpc.ServerCall) (map[security.BlessingPattern]security.Blessings, error) {
	return s.impl.BlessingStorePeerBlessings(call)
}

func (s implAgentServerStub) BlessingStoreDebugString(call rpc.ServerCall) (string, error) {
	return s.impl.BlessingStoreDebugString(call)
}

func (s implAgentServerStub) BlessingRootsAdd(call rpc.ServerCall, i0 []byte, i1 security.BlessingPattern) error {
	return s.impl.BlessingRootsAdd(call, i0, i1)
}

func (s implAgentServerStub) BlessingRootsRecognized(call rpc.ServerCall, i0 []byte, i1 string) error {
	return s.impl.BlessingRootsRecognized(call, i0, i1)
}

func (s implAgentServerStub) BlessingRootsDebugString(call rpc.ServerCall) (string, error) {
	return s.impl.BlessingRootsDebugString(call)
}

func (s implAgentServerStub) NotifyWhenChanged(call *AgentNotifyWhenChangedServerCallStub) error {
	return s.impl.NotifyWhenChanged(call)
}

func (s implAgentServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implAgentServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{AgentDesc}
}

// AgentDesc describes the Agent interface.
var AgentDesc rpc.InterfaceDesc = descAgent

// descAgent hides the desc to keep godoc clean.
var descAgent = rpc.InterfaceDesc{
	Name:    "Agent",
	PkgPath: "v.io/x/ref/security/agent/server",
	Methods: []rpc.MethodDesc{
		{
			Name: "Bless",
			InArgs: []rpc.ArgDesc{
				{"key", ``},               // []byte
				{"wit", ``},               // security.Blessings
				{"extension", ``},         // string
				{"caveat", ``},            // security.Caveat
				{"additionalCaveats", ``}, // []security.Caveat
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessSelf",
			InArgs: []rpc.ArgDesc{
				{"name", ``},    // string
				{"caveats", ``}, // []security.Caveat
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "Sign",
			InArgs: []rpc.ArgDesc{
				{"message", ``}, // []byte
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Signature
			},
		},
		{
			Name: "MintDischarge",
			InArgs: []rpc.ArgDesc{
				{"forCaveat", ``},                    // security.Caveat
				{"caveatOnDischarge", ``},            // security.Caveat
				{"additionalCaveatsOnDischarge", ``}, // []security.Caveat
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Discharge
			},
		},
		{
			Name: "PublicKey",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []byte
			},
		},
		{
			Name: "BlessingsByName",
			InArgs: []rpc.ArgDesc{
				{"name", ``}, // security.BlessingPattern
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []security.Blessings
			},
		},
		{
			Name: "BlessingsInfo",
			InArgs: []rpc.ArgDesc{
				{"blessings", ``}, // security.Blessings
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // map[string][]security.Caveat
			},
		},
		{
			Name: "AddToRoots",
			InArgs: []rpc.ArgDesc{
				{"blessing", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStoreSet",
			InArgs: []rpc.ArgDesc{
				{"blessings", ``}, // security.Blessings
				{"forPeers", ``},  // security.BlessingPattern
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStoreForPeer",
			InArgs: []rpc.ArgDesc{
				{"peerBlessings", ``}, // []string
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStoreSetDefault",
			InArgs: []rpc.ArgDesc{
				{"blessings", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStoreDefault",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStorePeerBlessings",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // map[security.BlessingPattern]security.Blessings
			},
		},
		{
			Name: "BlessingStoreDebugString",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // string
			},
		},
		{
			Name: "BlessingRootsAdd",
			InArgs: []rpc.ArgDesc{
				{"root", ``},    // []byte
				{"pattern", ``}, // security.BlessingPattern
			},
		},
		{
			Name: "BlessingRootsRecognized",
			InArgs: []rpc.ArgDesc{
				{"root", ``},     // []byte
				{"blessing", ``}, // string
			},
		},
		{
			Name: "BlessingRootsDebugString",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // string
			},
		},
		{
			Name: "NotifyWhenChanged",
			Doc:  "// Clients using caching should call NotifyWhenChanged upon connecting to\n// the server. The server will stream back values whenever the client should\n// flush the cache. The streamed value is arbitrary, simply flush whenever\n// recieving a new item.",
		},
	},
}

// AgentNotifyWhenChangedServerStream is the server stream for Agent.NotifyWhenChanged.
type AgentNotifyWhenChangedServerStream interface {
	// SendStream returns the send side of the Agent.NotifyWhenChanged server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item bool) error
	}
}

// AgentNotifyWhenChangedServerCall represents the context passed to Agent.NotifyWhenChanged.
type AgentNotifyWhenChangedServerCall interface {
	rpc.ServerCall
	AgentNotifyWhenChangedServerStream
}

// AgentNotifyWhenChangedServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements AgentNotifyWhenChangedServerCall.
type AgentNotifyWhenChangedServerCallStub struct {
	rpc.StreamServerCall
}

// Init initializes AgentNotifyWhenChangedServerCallStub from rpc.StreamServerCall.
func (s *AgentNotifyWhenChangedServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the Agent.NotifyWhenChanged server stream.
func (s *AgentNotifyWhenChangedServerCallStub) SendStream() interface {
	Send(item bool) error
} {
	return implAgentNotifyWhenChangedServerCallSend{s}
}

type implAgentNotifyWhenChangedServerCallSend struct {
	s *AgentNotifyWhenChangedServerCallStub
}

func (s implAgentNotifyWhenChangedServerCallSend) Send(item bool) error {
	return s.s.Send(item)
}
