package e2e_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

// TestBridge ensures that bridge tests cases defined in e2e/app/tokenbridge.go were successful.
func TestBridge(t *testing.T) {
	t.Parallel()
	testNetwork(t, func(t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()
		ctx := context.Background()

		if _, ok := network.EthereumChain(); !ok {
			t.Skip("no ethereum chain")
		}

		omniEVM, ok := network.OmniEVMChain()
		require.True(t, ok)

		omniRPC, err := endpoints.ByNameOrID(omniEVM.Name, omniEVM.ID)
		require.NoError(t, err)

		omniClient, err := ethclient.Dial(omniEVM.Name, omniRPC)
		require.NoError(t, err)

		l1, ok := network.EthereumChain()
		require.True(t, ok)

		l1RPC, err := endpoints.ByNameOrID(l1.Name, l1.ID)
		require.NoError(t, err)

		l1Client, err := ethclient.Dial(l1.Name, l1RPC)
		require.NoError(t, err)

		l1TokenAddr := contracts.Token(network.ID)
		l1BridgeAddr := contracts.L1Bridge(network.ID)

		l1Token, err := bindings.NewOmni(l1TokenAddr, l1Client)
		require.NoError(t, err)

		nativeBridge, err := bindings.NewOmniBridgeNative(common.HexToAddress(predeploys.OmniBridgeNative), omniClient)
		require.NoError(t, err)

		sumToNative := big.NewInt(0)
		sumToL1 := big.NewInt(0)

		for _, test := range app.ToNativeBridgeTests {
			balance, err := omniClient.BalanceAt(ctx, test.To, nil)
			require.NoError(t, err)
			require.Equal(t, balance, test.Amount)
			sumToNative = new(big.Int).Add(sumToNative, test.Amount)
		}

		for _, test := range app.ToL1BridgeTests {
			balance, err := l1Token.BalanceOf(nil, test.To)
			require.NoError(t, err)
			require.Equal(t, balance, test.Amount)
			sumToL1 = new(big.Int).Add(sumToL1, test.Amount)
		}

		expectedL1BridgeBalance := new(big.Int).Sub(sumToNative, sumToL1)

		// assert l1 bridge balance tracked on native bridge is expected
		trackedL1BridgeBalance, err := nativeBridge.L1BridgeBalance(nil)
		require.NoError(t, err)
		require.Equal(t, expectedL1BridgeBalance, trackedL1BridgeBalance)

		// assert actual token balance of l1 bridge is expected
		l1BridgeBalance, err := l1Token.BalanceOf(nil, l1BridgeAddr)
		require.NoError(t, err)
		require.Equal(t, expectedL1BridgeBalance, l1BridgeBalance)
	})
}
