package e2e_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

// TestGasPumps ensures that bridge tests cases defined in e2e/app/gaspump.go were successful.
func TestGasPumps(t *testing.T) {
	t.Parallel()
	testNetwork(t, func(ctx context.Context, t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()

		omniEVM, ok := network.OmniEVMChain()
		require.True(t, ok)

		omniRPC, err := endpoints.ByNameOrID(omniEVM.Name, omniEVM.ID)
		require.NoError(t, err)

		omniClient, err := ethclient.Dial(omniEVM.Name, omniRPC)
		require.NoError(t, err)

		// Sum targetOMNI for each chain / test case pair
		// Each test case is run on per chain, except for OmniEVM

		totalTargetOMNI := make(map[common.Address]*big.Int)
		for _, chain := range network.EVMChains() {
			// skip OmniEVM
			if chain.ID == omniEVM.ID {
				continue
			}

			for _, test := range app.GasPumpTests {
				current, ok := totalTargetOMNI[test.Recipient]
				if !ok {
					current = umath.Zero()
				}

				totalTargetOMNI[test.Recipient] = umath.Add(current, test.TargetOMNI)
			}
		}

		for _, test := range app.GasPumpTests {
			balance, err := omniClient.BalanceAt(ctx, test.Recipient, nil)
			require.NoError(t, err)
			require.Equalf(t, totalTargetOMNI[test.Recipient], balance, "recipient: %s", test.Recipient.Hex())
		}
	})
}
