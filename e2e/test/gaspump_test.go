package e2e_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

// TestGasPumps ensures that bridge tests cases defined in e2e/app/gaspump.go were successful.
func TestGasPumps(t *testing.T) {
	t.Parallel()
	testNetwork(t, func(t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()
		ctx := context.Background()

		omniEVM, ok := network.OmniEVMChain()
		require.True(t, ok)

		omniRPC, err := endpoints.ByNameOrID(omniEVM.Name, omniEVM.ID)
		require.NoError(t, err)

		omniClient, err := ethclient.Dial(omniEVM.Name, omniRPC)
		require.NoError(t, err)

		for _, test := range app.GasPumpTests {
			balance, err := omniClient.BalanceAt(ctx, test.Recipient, nil)
			require.NoError(t, err)

			// Just test that balance > 0 for now
			// TODO: assert that amount is equal to sum of AmountETH spent converted to OMNI
			// Should account for the xcall fee, gas pump toll, and fee oracle conversion rates
			require.Positive(t, balance.Uint64(), "recipient: %s", test.Recipient)
		}
	})
}
