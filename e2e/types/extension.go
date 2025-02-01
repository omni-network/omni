package types

import (
	"context"

	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

type Extension interface {
	// Name returns the name of the extension.
	Name() string

	// Deploy deploys the extension.
	Deploy(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints, backends ethbackend.Backends) error

	// Test tests the extension.
	Test(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints, backends ethbackend.Backends) error
}
