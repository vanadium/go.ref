// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blesser

import (
	"fmt"
	"strings"
	"time"

	"v.io/x/ref/services/identity"
	"v.io/x/ref/services/identity/internal/oauth"
	"v.io/x/ref/services/identity/internal/util"

	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security"
	"v.io/v23/vom"
)

type macaroonBlesser struct {
	key []byte
}

// NewMacaroonBlesserServer provides an identity.MacaroonBlesser Service that generates blessings
// after unpacking a BlessingMacaroon.
func NewMacaroonBlesserServer(key []byte) identity.MacaroonBlesserServerStub {
	return identity.MacaroonBlesserServer(&macaroonBlesser{key})
}

func (b *macaroonBlesser) Bless(ctx *context.T, call rpc.ServerCall, macaroon string) (security.Blessings, error) {
	secCall := call.Security()
	var empty security.Blessings
	inputs, err := util.Macaroon(macaroon).Decode(b.key)
	if err != nil {
		return empty, err
	}
	var m oauth.BlessingMacaroon
	if err := vom.Decode(inputs, &m); err != nil {
		return empty, err
	}
	if time.Now().After(m.Creation.Add(time.Minute * 5)) {
		return empty, fmt.Errorf("macaroon has expired")
	}
	if secCall.LocalPrincipal() == nil {
		return empty, fmt.Errorf("server misconfiguration: no authentication happened")
	}
	if len(m.Caveats) == 0 {
		m.Caveats = []security.Caveat{security.UnconstrainedUse()}
	}
	var extension string
	// TODO(ashankar): Remove after the transition from running identityd as "dev.v.io/root" to "dev.v.io/u" is complete.
	if local := security.LocalBlessingNames(ctx, secCall); len(local) == 1 && strings.HasSuffix(local[0], security.ChainSeparator+"u") {
		extension = m.Name
	} else {
		extension = strings.Join([]string{"users", m.Name}, security.ChainSeparator)
	}
	return secCall.LocalPrincipal().Bless(secCall.RemoteBlessings().PublicKey(),
		secCall.LocalBlessings(), extension, m.Caveats[0], m.Caveats[1:]...)
}
