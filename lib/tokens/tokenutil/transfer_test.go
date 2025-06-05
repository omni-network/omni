package tokenutil_test

import (
	"context"
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
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "run integration tests")

// TestTransferIntegration tests tokenutil.Transfer on an eth l1 fork.
func TestTransferIntegration(t *testing.T) {
	t.Parallel()

	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()

	rpcURL := os.Getenv("ETH_RPC")
	if rpcURL == "" {
		t.Skip("skipping test, ETH_RPC environment variable is not set")
	}

	ethCl, stop, err := anvil.Start(ctx, tutil.TempDir(t), evmchain.IDEthereum, anvil.WithFork(rpcURL))
	require.NoError(t, err)
	defer stop()

	pk, sender := newAccount(t)
	_, recipient := newAccount(t)
	chain := mustChain(t, evmchain.IDEthereum)
	usdc := mustToken(t, evmchain.IDEthereum, tokens.USDC)
	eth := mustToken(t, evmchain.IDEthereum, tokens.ETH)

	backend, err := ethbackend.NewBackend(chain.Name, chain.ChainID, time.Second, ethCl, pk)
	require.NoError(t, err)

	err = anvil.FundAccounts(ctx, ethCl, bi.Ether(10), sender)
	require.NoError(t, err)

	err = anvil.FundUSDC(ctx, ethCl, usdc.Address, bi.Dec6(10000), sender)
	require.NoError(t, err)

	// Native transfer
	receipt, err := tokenutil.Transfer(ctx, backend, eth, sender, recipient, bi.Ether(1))
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)
	require.Equal(t, bi.Ether(1), mustBalance(t, ctx, backend, eth, recipient))

	// USDC transfer
	receipt, err = tokenutil.Transfer(ctx, backend, usdc, sender, recipient, bi.Dec6(100))
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)
	require.Equal(t, bi.Dec6(100), mustBalance(t, ctx, backend, usdc, recipient))
}

func mustToken(t *testing.T, chainID uint64, asset tokens.Asset) tokens.Token {
	t.Helper()
	token, ok := tokens.ByAsset(chainID, asset)
	require.True(t, ok)

	return token
}

func newAccount(t *testing.T) (*ecdsa.PrivateKey, common.Address) {
	t.Helper()
	pk, err := crypto.GenerateKey()
	require.NoError(t, err)

	return pk, crypto.PubkeyToAddress(pk.PublicKey)
}

func mustBalance(t *testing.T, ctx context.Context, backend *ethbackend.Backend, token tokens.Token, account common.Address) *big.Int {
	t.Helper()
	balance, err := tokenutil.BalanceOf(ctx, backend, token, account)
	require.NoError(t, err)

	return balance
}

func mustChain(t *testing.T, chainID uint64) evmchain.Metadata {
	t.Helper()
	meta, ok := evmchain.MetadataByID(chainID)
	require.True(t, ok)

	return meta
}
