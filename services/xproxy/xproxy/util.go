// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xproxy

import (
	"fmt"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/flow"
	"v.io/v23/flow/message"
	"v.io/v23/naming"
	"v.io/v23/security"
)

func setBidiProtocol(ep naming.Endpoint) (naming.Endpoint, error) {
	_, _, routes, rid, bnames, mountable := getEndpointParts(ep)
	opts := routes
	opts = append(opts, bnames...)
	opts = append(opts, rid)
	opts = append(opts, mountable)
	epString := naming.FormatEndpoint("bidi", "", opts...)
	return v23.NewEndpoint(epString)
}

// setEndpointRoutingID returns a copy of ep with RoutingId changed to rid.
func setEndpointRoutingID(ep naming.Endpoint, rid naming.RoutingID) (naming.Endpoint, error) {
	network, address, routes, _, bnames, mountable := getEndpointParts(ep)
	opts := routes
	opts = append(opts, bnames...)
	opts = append(opts, rid)
	opts = append(opts, mountable)
	epString := naming.FormatEndpoint(network, address, opts...)
	return v23.NewEndpoint(epString)
}

// setEndpointRoutes returns a copy of ep with Routes changed to routes.
func setEndpointRoutes(ep naming.Endpoint, routes []string) (naming.Endpoint, error) {
	network, address, _, rid, bnames, mountable := getEndpointParts(ep)
	opts := routeOpts(routes)
	opts = append(opts, bnames...)
	opts = append(opts, rid)
	opts = append(opts, mountable)
	epString := naming.FormatEndpoint(network, address, opts...)
	return v23.NewEndpoint(epString)
}

// getEndpointParts returns all the fields of ep.
func getEndpointParts(ep naming.Endpoint) (network string, address string,
	routes []naming.EndpointOpt, rid naming.RoutingID,
	blessingNames []naming.EndpointOpt, mountable naming.EndpointOpt) {
	network, address = ep.Addr().Network(), ep.Addr().String()
	routes = routeOpts(ep.Routes())
	rid = ep.RoutingID()
	blessingNames = blessingOpts(ep.BlessingNames())
	mountable = naming.ServesMountTable(ep.ServesMountTable())
	return
}

func routeOpts(routes []string) []naming.EndpointOpt {
	var routeOpts []naming.EndpointOpt
	for _, route := range routes {
		routeOpts = append(routeOpts, naming.RouteOpt(route))
	}
	return routeOpts
}

func blessingOpts(blessingNames []string) []naming.EndpointOpt {
	var blessingOpts []naming.EndpointOpt
	for _, b := range blessingNames {
		blessingOpts = append(blessingOpts, naming.BlessingOpt(b))
	}
	return blessingOpts
}

// TODO(suharshs): Figure out what blessings to present here. And present discharges.
type proxyAuthorizer struct{}

func (proxyAuthorizer) AuthorizePeer(
	ctx *context.T,
	localEndpoint, remoteEndpoint naming.Endpoint,
	remoteBlessings security.Blessings,
	remoteDischarges map[string]security.Discharge,
) ([]string, []security.RejectedBlessing, error) {
	return nil, nil, nil
}

func (a proxyAuthorizer) BlessingsForPeer(ctx *context.T, _ []string) (
	security.Blessings, map[string]security.Discharge, error) {
	return v23.GetPrincipal(ctx).BlessingStore().Default(), nil, nil
}

func writeMessage(ctx *context.T, m message.Message, f flow.Flow) error {
	w, err := message.Append(ctx, m, nil)
	if err != nil {
		return err
	}
	_, err = f.WriteMsg(w)
	return err
}

func readMessage(ctx *context.T, f flow.Flow) (message.Message, error) {
	b, err := f.ReadMsg()
	if err != nil {
		return nil, err
	}
	return message.Read(ctx, b)
}

func readProxyResponse(ctx *context.T, f flow.Flow) ([]naming.Endpoint, error) {
	msg, err := readMessage(ctx, f)
	if err != nil {
		return nil, err
	}
	res, ok := msg.(*message.ProxyResponse)
	if !ok {
		return nil, NewErrUnexpectedMessage(ctx, fmt.Sprintf("%t", msg))
	}
	return res.Endpoints, nil
}