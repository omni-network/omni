package e2e_test

import (
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestPortalRegistry(t *testing.T) {
	t.Parallel()
	testNetwork(t, func(t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()

		omniEVM, ok := network.OmniEVMChain()
		require.True(t, ok)

		rpc, err := endpoints.ByNameOrID(omniEVM.Name, omniEVM.ID)
		require.NoError(t, err)

		omniClient, err := ethclient.Dial(omniEVM.Name, rpc)
		require.NoError(t, err)

		// test that all portals are registered
		preg, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), omniClient)
		require.NoError(t, err)

		for _, chain := range network.EVMChains() {
			registration, err := preg.Get(nil, chain.ID)
			require.NoError(t, err)

			require.Equal(t, chain.PortalAddress, registration.Addr, "chain %v portal", chain.ID)
			require.Equal(t, chain.ID, registration.ChainId, "chain %v id", chain.ID)
			// TODO(kevin): fix test when XRegistry supports conf level shards
			// require.Equal(t, netconf.MustShardToStrat(chain.Shards[0]), registration.FinalizationStrat, "chain %v finalization strategy", chain.ID)
			require.Equal(t, chain.DeployHeight, registration.DeployHeight, "chain %v deploy height", chain.ID)
		}
	})
}
