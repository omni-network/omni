package e2e_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/bridge"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

// TestBridge deploys the bridge and tests it.
func TestBridge(t *testing.T) {
	t.Parallel()
	skipFunc := func(manifest types.Manifest) bool {
		return !manifest.AllE2ETests
	}
	maybeTestNetwork(t, skipFunc, func(ctx context.Context, t *testing.T, deps NetworkDeps) {
		t.Helper()

		err := bridge.DeployBridge(ctx, deps.Testnet, deps.Backends)
		require.NoError(t, err)

		network := deps.Network

		if _, ok := network.EthereumChain(); !ok {
			t.Skip("no ethereum chain")
		}

		omniBackend, err := deps.OmniBackend()
		require.NoError(t, err)
		l1Backend, err := deps.L1Backend()
		require.NoError(t, err)
		addrs, err := contracts.GetAddresses(ctx, network.ID)
		require.NoError(t, err)
		l1Token, err := bindings.NewOmni(addrs.Token, l1Backend)
		require.NoError(t, err)
		nativeBridge, err := bindings.NewOmniBridgeNative(common.HexToAddress(predeploys.OmniBridgeNative), omniBackend)
		require.NoError(t, err)

		trackedL1BridgeBalanceBefore, err := nativeBridge.L1Deposits(nil)
		require.NoError(t, err)

		l1BridgeBalanceBefore, err := l1Token.BalanceOf(nil, addrs.L1Bridge)
		require.NoError(t, err)

		addr1, addr2 := tutil.RandomAddress(), tutil.RandomAddress()
		toNatives := []bridge.ToAmount{
			{To: addr1, Amount: bi.Ether(20)},
			{To: addr2, Amount: bi.Ether(30)},
		}
		toL1s := []bridge.ToAmount{
			{To: addr1, Amount: bi.Ether(2)},
			{To: addr2, Amount: bi.Ether(3)},
		}

		err = bridge.Test(ctx, deps.Testnet, deps.Backends, toNatives, toL1s)
		require.NoError(t, err)

		sumToNative := bi.Zero()
		sumToL1 := bi.Zero()

		for _, test := range toNatives {
			balance, err := omniBackend.BalanceAt(ctx, test.To, nil)
			require.NoError(t, err)
			require.Equal(t, balance, test.Amount)
			sumToNative = bi.Add(sumToNative, test.Amount)
		}

		for _, test := range toL1s {
			require.Eventually(t, func() bool {
				balance, err := l1Token.BalanceOf(nil, test.To)
				return err == nil && bi.EQ(balance, test.Amount)
			}, time.Second*30, time.Second)
			sumToL1 = bi.Add(sumToL1, test.Amount)
		}

		expectedL1BridgeBalanceDelta := bi.Sub(sumToNative, sumToL1)

		// assert l1 bridge balance tracked on native bridge is expected
		trackedL1BridgeBalanceAfter, err := nativeBridge.L1Deposits(nil)
		require.NoError(t, err)
		trackedL1BridgeBalanceDelta := bi.Sub(trackedL1BridgeBalanceAfter, trackedL1BridgeBalanceBefore)
		require.Equal(t, expectedL1BridgeBalanceDelta, trackedL1BridgeBalanceDelta)

		// assert actual token balance of l1 bridge is expected
		l1BridgeBalanceAfter, err := l1Token.BalanceOf(nil, addrs.L1Bridge)
		require.NoError(t, err)
		l1BridgeBalanceDelta := bi.Sub(l1BridgeBalanceAfter, l1BridgeBalanceBefore)
		// Ensure the L1Bridge balance delta is greater than or equal to the expected delta
		// Since Solver also does rebalancing which increases the L1Bridge balance
		require.Truef(t, bi.LTE(expectedL1BridgeBalanceDelta, l1BridgeBalanceDelta), "unexpected l1 bridge balance delta: min=%s, actual=%s", expectedL1BridgeBalanceDelta, l1BridgeBalanceDelta)
	})
}
