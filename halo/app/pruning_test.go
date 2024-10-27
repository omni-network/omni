// Only run this if -race=false, since CosmosSDK has known data races when doing gRPC queries.
//go:build !race

//nolint:paralleltest // CosmosSDK dependency injection prevents parallel execution
package app_test

import (
	"context"
	"testing"
	"time"

	haloapp "github.com/omni-network/omni/halo/app"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	pruningtypes "cosmossdk.io/store/pruning/types"
	"github.com/stretchr/testify/require"
)

func TestPruningHistory(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx, err := log.Init(ctx, log.Config{Color: log.ColorForce, Level: "debug", Format: log.FormatConsole})
	require.NoError(t, err)

	cfg := setupSimnet(t)

	// Prune everything > 2 blocks old (every interval=10 blocks)
	cfg.PruningOption = pruningtypes.PruningOptionEverything
	cfg.MinRetainBlocks = 2
	// Note the pruneInterval is 10, so we only prune at height 20

	// Start the server async
	async, stopfunc, err := haloapp.Start(ctx, cfg)
	require.NoError(t, err)
	go func() {
		tutil.RequireNoError(t, <-async)
	}()

	// Connect to the server.
	cl, err := rpchttp.New(cfg.Comet.RPC.ListenAddress, "/websocket")
	require.NoError(t, err)

	cprov := cprovider.NewABCI(cl, netconf.Simnet)

	// Wait until we get to block 1.
	waitUntilHeight := uint64(1)
	require.Eventually(t, func() bool {
		s, err := cl.Status(ctx)
		if err != nil {
			t.Log("Failed to get status: ", err)
			return false
		}

		return s.SyncInfo.LatestBlockHeight >= int64(waitUntilHeight)
	}, time.Second*time.Duration(waitUntilHeight*2), time.Millisecond*100)

	srcChain := evmchain.IDOmniDevnet // Pick chain without fuzzy conf levels
	chainVer := xchain.ChainVersion{ID: srcChain, ConfLevel: xchain.ConfFinalized}

	// Wait until we have an attestation with offset 2-or-more for srcChain
	// That means that offset=1 is eligible for deletion.
	var eligibleHeight uint64
	require.Eventually(t, func() bool {
		status, err := cl.Status(ctx)
		tutil.RequireNoError(t, err)

		att, ok, err := cprov.LatestAttestation(ctx, chainVer)
		tutil.RequireNoError(t, err)
		if !ok {
			t.Logf("still waiting for an attestation: height=%d", status.SyncInfo.LatestBlockHeight)
			return false
		} else if att.AttestOffset == 1 {
			t.Logf("still waiting for 2nd attestation: height=%d", status.SyncInfo.LatestBlockHeight)
		}

		require.NotEmpty(t, att.AttestOffset)
		require.Equal(t, srcChain, att.ChainID)

		eligibleHeight = uint64(status.SyncInfo.LatestBlockHeight)

		return true
	}, time.Minute, time.Millisecond*300)

	// Now wait until we pass the height at which the offset=1 attestation
	// has been deleted and the eligibleHeight state pruned from the DB.
	var prunedHeight uint64
	require.Eventually(t, func() bool {
		status, err := cl.Status(ctx)
		tutil.RequireNoError(t, err)

		_, err = cprov.AttestationsFrom(ctx, chainVer, 1)
		if err == nil {
			t.Logf("att still available: created=%d, current=%d", eligibleHeight, status.SyncInfo.LatestBlockHeight)
			return false
		}

		require.True(t, cprovider.IsErrHistoryPruned(err))
		prunedHeight = uint64(status.SyncInfo.LatestBlockHeight)

		return true
	}, time.Minute, time.Millisecond*300)

	const pruneInterval = uint64(10)
	minPrunedHeight := eligibleHeight + cfg.MinRetainBlocks
	maxPrunedHeight := eligibleHeight/pruneInterval + (2 * pruneInterval) // Allow pruning up to the next interval
	t.Logf("prunedHeight=%d, createdHeight=%d, minRetainBlocks=%d", prunedHeight, eligibleHeight, cfg.MinRetainBlocks)

	require.GreaterOrEqual(t, prunedHeight, minPrunedHeight, "prunedHeight=%d too low (eligibleHeight=%d)", prunedHeight, eligibleHeight)
	require.LessOrEqual(t, prunedHeight, maxPrunedHeight, "prunedHeight=%d too height (eligibleHeight=%d)", prunedHeight, eligibleHeight)

	cancel()

	// Stop the server.
	require.NoError(t, stopfunc(context.Background()))
}
