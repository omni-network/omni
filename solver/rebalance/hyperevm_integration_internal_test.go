package rebalance

import (
	"context"
	"flag"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cctp/testutil"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/layerzero"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/usdt0"
	"github.com/omni-network/omni/solver/fundthresh"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestRebalanceHyperEVMOnce(t *testing.T) {
	t.Parallel()

	// Currently timeing out waiting for OFTSent event. Contracts likely upgraded, and bindings / tests need updating.
	// USDT0 sends still functioning. Solver hyper evm rebalance paused.
	t.Skip()

	if f := flag.Lookup("integration"); f != nil && !f.Value.(flag.Getter).Get().(bool) {
		t.Skip("Skipping integration test")
	}

	rpcs := map[uint64]string{
		evmchain.IDEthereum: mustEnv(t, "ETH_RPC"),
		evmchain.IDHyperEVM: mustEnv(t, "HYPER_EVM_RPC"),
	}

	chains := []evmchain.Metadata{
		mustMeta(t, evmchain.IDEthereum),
		mustMeta(t, evmchain.IDHyperEVM),
	}

	ctx := t.Context()

	clients, stop := testutil.StartAnvilForks(t, ctx, rpcs, chains)
	defer stop()

	solverPk, solver := testutil.NewAccount(t)

	backends, err := ethbackend.BackendsFromClients(clients, solverPk)
	tutil.RequireNoError(t, err)

	// Fund gas
	err = anvil.FundAccounts(ctx, clients[evmchain.IDEthereum], bi.Ether(1), solver)
	tutil.RequireNoError(t, err)

	// Fund at surplus USDC + 1000
	usdc := mustToken(evmchain.IDEthereum, tokens.USDC)
	fundToken(t, ctx, clients[evmchain.IDEthereum], usdc, solver, bi.Add(fundthresh.Get(usdc).Surplus(), bi.Dec6(1000)))

	// Fund at surplus USDT, so anything more can be sent to HyperEVM
	usdt := mustToken(evmchain.IDEthereum, tokens.USDT)
	fundToken(t, ctx, clients[evmchain.IDEthereum], usdt, solver, fundthresh.Get(usdt).Surplus())

	// ~1000 USDC should be swapped to USDT, and sent to HyperEVM
	oftAddr := usdt0.OFTByChain(evmchain.IDEthereum)
	oft, err := usdt0.NewIOFT(oftAddr, clients[evmchain.IDEthereum])
	tutil.RequireNoError(t, err)

	// Get latest block, used as startBlock for OFTSent event filtering
	startBlock, err := clients[evmchain.IDEthereum].BlockNumber(ctx)
	tutil.RequireNoError(t, err)

	// Run single rebalance
	err = rebalanceHyperEVMOnce(ctx, backends, solver, nil)
	tutil.RequireNoError(t, err)

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Wait for OFTSent event to be emitted
	for {
		select {
		case <-timeout:
			t.Fatal("timeout waiting for OFTSent event")
		case <-ticker.C:
			// Query OFTSent logs from start block to latest
			endBlock, err := clients[evmchain.IDEthereum].BlockNumber(ctx)
			tutil.RequireNoError(t, err)
			logs, err := oft.FilterOFTSent(&bind.FilterOpts{
				Start: startBlock,
				End:   &endBlock,
			}, nil, []common.Address{solver})
			tutil.RequireNoError(t, err)

			// Any event must match (we are only eoa transacting)
			for logs.Next() {
				event := logs.Event
				require.Equal(t, mustEID(t, evmchain.IDHyperEVM), event.DstEid)
				require.Equal(t, solver, event.FromAddress)

				// Verify amount send is at least 999 (unless to account for swap fees)
				tutil.RequireGT(t, event.AmountSentLD, bi.Dec6(999))
				tutil.RequireGT(t, event.AmountReceivedLD, bi.Dec6(999))

				return
			}

			tutil.RequireNoError(t, logs.Error())
		}
	}
}

func mustEnv(t *testing.T, key string) string {
	t.Helper()
	value := os.Getenv(key)
	require.NotEmpty(t, value, "Environment variable %s is not set", key)

	return value
}

func mustEID(t *testing.T, chainID uint64) uint32 {
	t.Helper()

	eid, ok := layerzero.EIDByChain(chainID)
	require.True(t, ok, "EID not found for chain ID %d", chainID)

	return eid
}

func mustMeta(t *testing.T, chainID uint64) evmchain.Metadata {
	t.Helper()

	meta, ok := evmchain.MetadataByID(chainID)
	require.True(t, ok)

	return meta
}

// fundToken funds the solver with a specific token.
func fundToken(t *testing.T, ctx context.Context, client ethclient.Client, token tokens.Token, account common.Address, amt *big.Int) {
	t.Helper()

	if token.Is(tokens.USDC) {
		err := anvil.FundUSDC(ctx, client, token.Address, amt, account)
		tutil.RequireNoError(t, err)

		return
	}

	if token.Is(tokens.USDT) && token.ChainID == evmchain.IDEthereum {
		err := anvil.FundL1USDT(ctx, client, token.Address, amt, account)
		tutil.RequireNoError(t, err)

		return
	}

	t.Fatalf("unsupported token %s", token)
}
