package e2e_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/omni-network/omni/e2e/gasstation"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

// TestGasPumps deploys the gasstation app and pumps and tests them.
func TestGasPumps(t *testing.T) {
	t.Parallel()
	skipFunc := func(manifest types.Manifest) bool {
		return !manifest.AllE2ETests
	}
	maybeTestNetwork(t, skipFunc, func(ctx context.Context, t *testing.T, deps NetworkDeps) {
		t.Helper()

		err := gasstation.DeployEphemeralGasApp(ctx, deps.Testnet, deps.Backends)
		require.NoError(t, err)

		tests := []gasstation.GasPumpTest{
			{
				Recipient:  tutil.RandomAddress(),
				TargetOMNI: bi.Ether(0.1),
			},
			{
				Recipient:  tutil.RandomAddress(),
				TargetOMNI: bi.Ether(0.2),
			},
		}
		err = gasstation.TestGasPumps(ctx, deps.Testnet, deps.Backends, tests)
		require.NoError(t, err)

		network := deps.Network
		omniBackend, err := deps.OmniBackend()
		require.NoError(t, err)

		// Sum targetOMNI for each chain / test case pair
		// Each test case is run on per chain, except for OmniEVM

		totalTargetOMNI := make(map[common.Address]*big.Int)
		for _, chain := range network.EVMChains() {
			// skip OmniEVM
			if chain.ID == network.ID.Static().OmniExecutionChainID {
				continue
			}

			for _, test := range tests {
				current, ok := totalTargetOMNI[test.Recipient]
				if !ok {
					current = bi.Zero()
				}

				totalTargetOMNI[test.Recipient] = bi.Add(current, test.TargetOMNI)
			}
		}

		for _, test := range tests {
			require.Eventually(t, func() bool {
				balance, err := omniBackend.BalanceAt(ctx, test.Recipient, nil)
				return err == nil && bi.EQ(balance, totalTargetOMNI[test.Recipient])
			}, time.Second*30, time.Second)
		}
	})
}
