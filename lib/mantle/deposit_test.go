package mantle_test

import (
	"crypto/ecdsa"
	"flag"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/mantle"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "run integration tests")

// TestDepositUSDC tests the DepositERC20 function.
func TestDepositUSDC(t *testing.T) {
	t.Parallel()

	if !*integration {
		t.Skip("skipping integration test")
	}

	pk, user := newAccount(t)

	rpcURL := os.Getenv("ETH_RPC")
	require.NotEmpty(t, rpcURL, "ETH_RPC required")

	ctx := t.Context()

	ethCl, stop, err := anvil.Start(ctx, tutil.TempDir(t), evmchain.IDEthereum, anvil.WithFork(rpcURL))
	tutil.RequireNoError(t, err)
	defer stop()

	l1, ok := evmchain.MetadataByID(evmchain.IDEthereum)
	require.True(t, ok, "ethereum metadata not found")

	backend, err := ethbackend.NewBackend(l1.Name, l1.ChainID, time.Second, ethCl, pk)
	tutil.RequireNoError(t, err)

	l1USDC, ok := tokens.ByAsset(evmchain.IDEthereum, tokens.USDC)
	require.True(t, ok, "USDC token not found")

	l2USDC, ok := tokens.ByAsset(evmchain.IDMantle, tokens.USDC)
	require.True(t, ok, "USDC token not found")

	// Fund USDC for bridging
	err = anvil.FundUSDC(ctx, ethCl, l1USDC.Address, bi.Dec6(100), user)
	tutil.RequireNoError(t, err)

	// Fund  gas
	err = anvil.FundAccounts(ctx, ethCl, bi.Ether(1), user)
	tutil.RequireNoError(t, err)

	tutil.RequireEQ(t, balanceOf(t, backend, l1USDC, user), bi.Dec6(100))

	receipt, err := mantle.DepositUSDC(ctx, backend, user, bi.Dec6(100))
	tutil.RequireNoError(t, err)

	tutil.RequireEQ(t, balanceOf(t, backend, l1USDC, user), bi.Dec6(0))
	require.Equal(t, ethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Use bridge to parse receipt logs
	bridge, err := mantle.NewL1Bridge(
		common.HexToAddress("0x95fC37A27a2f68e3A647CDc081F0A89bb47c3012"), // L1StandardBridge address
		backend)
	tutil.RequireNoError(t, err)

	ev, err := bridge.ParseERC20DepositInitiated(
		*receipt.Logs[1], // will be second log
	)

	require.NoError(t, err)
	require.Equal(t, ev.From, user)
	require.Equal(t, ev.L1Token, l1USDC.Address)
	require.Equal(t, ev.L2Token, l2USDC.Address)
	require.Equal(t, []byte{}, ev.ExtraData)
	tutil.RequireEQ(t, ev.Amount, bi.Dec6(100))
}

func newAccount(t *testing.T) (*ecdsa.PrivateKey, common.Address) {
	t.Helper()

	pk, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr := crypto.PubkeyToAddress(pk.PublicKey)

	return pk, addr
}

func balanceOf(t *testing.T, backend *ethbackend.Backend, token tokens.Token, addr common.Address) *big.Int {
	t.Helper()

	balance, err := tokenutil.BalanceOf(t.Context(), backend, token, addr)
	require.NoError(t, err)

	return balance
}
