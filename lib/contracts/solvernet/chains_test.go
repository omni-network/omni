package solvernet_test

import (
	"testing"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func TestChains(t *testing.T) {
	t.Parallel()

	require.NotPanics(t, func() {
		for _, network := range netconf.All() {
			_ = solvernet.Chains(network)
		}
	})
}

func TestFilterByContracts(t *testing.T) {
	t.Parallel()
	network := netconf.Network{ID: netconf.Mainnet}
	endpoints := xchain.RPCEndpoints{"bsc": "https://foo.bar"}

	network = solvernet.AddNetwork(t.Context(), network, solvernet.FilterByContracts(t.Context(), endpoints))
	require.Empty(t, network.Chains)
}
