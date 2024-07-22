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

	// Prune everything > 10 blocks old
	cfg.PruningOption = pruningtypes.PruningOptionEverything
	cfg.MinRetainBlocks = 10

	// Start the server async
	async, stopfunc, err := haloapp.Start(ctx, cfg)
	require.NoError(t, err)
	go func() {
		tutil.RequireNoError(t, <-async)
	}()

	// Connect to the server.
	cl, err := rpchttp.New("http://localhost:26657", "/websocket")
	require.NoError(t, err)

	cprov := cprovider.NewABCIProvider(cl, netconf.Simnet, netconf.ChainVersionNamer(netconf.Simnet))

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

	srcChain := evmchain.IDMockL1Fast
	chainVer := xchain.ChainVersion{ID: srcChain, ConfLevel: xchain.ConfFinalized}

	var attOffset uint64
	// Wait until we have an attestation for srcChain
	require.Eventually(t, func() bool {
		att, ok, err := cprov.LatestAttestation(ctx, chainVer)
		tutil.RequireNoError(t, err)
		if !ok {
			t.Log("still waiting for an attestation")
			return false
		}

		require.Equal(t, srcChain, att.SourceChainID)

		attOffset = att.BlockOffset

		status, err := cl.Status(ctx)
		tutil.RequireNoError(t, err)

		// Wait until after the height at which the state containing this attestation should be pruned.
		waitUntilHeight = uint64(status.SyncInfo.LatestBlockHeight) + cfg.MinRetainBlocks + 2

		return true
	}, time.Minute, time.Second)

	// Now wait until we pass the height at which the attestation should be pruned.
	require.Eventually(t, func() bool {
		s, err := cl.Status(ctx)
		if err != nil {
			t.Log("Failed to get status: ", err)
			return false
		}

		return s.SyncInfo.LatestBlockHeight >= int64(waitUntilHeight)
	}, time.Second*time.Duration(waitUntilHeight*2), time.Millisecond*100)

	_, err = cprov.AttestationsFrom(ctx, chainVer, attOffset)
	require.True(t, cprovider.IsErrHistoryPruned(err))

	cancel()

	// Stop the server.
	require.NoError(t, stopfunc(ctx))
}
