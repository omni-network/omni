package gaspumps

import (
	"context"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

type ext struct{}

func Ext() types.Extension { return ext{} }

var _ types.Extension = ext{}

func (ext) Name() string {
	return "gaspumps"
}

func (ext) Deploy(_ context.Context, _ netconf.Network, _ xchain.RPCEndpoints, _ ethbackend.Backends) error {
	// TODO
	return nil
}

func (ext) Test(_ context.Context, _ netconf.Network, _ xchain.RPCEndpoints, _ ethbackend.Backends) error {
	// TODO
	return nil
}
