package app_test

import (
	"bytes"
	"context"
	"os"
	"testing"
	"time"

	haloapp "github.com/omni-network/omni/halo/app"
	halocmd "github.com/omni-network/omni/halo/cmd"
	halocfg "github.com/omni-network/omni/halo/config"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cometbft/cometbft/types"

	db "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

func TestSmoke(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx, err := log.Init(ctx, log.Config{Color: log.ColorForce, Level: "debug", Format: log.FormatConsole})
	require.NoError(t, err)

	cfg := setupSimnet(t)

	// Start the server async
	stopfunc, err := haloapp.Start(ctx, cfg)
	require.NoError(t, err)

	// Connect to the server.
	cl, err := rpchttp.New("http://localhost:26657", "/websocket")
	require.NoError(t, err)

	cprov := cprovider.NewABCIProvider(cl, netconf.Simnet, nil)

	// Wait until we get to block 3.
	const target = uint64(3)
	require.Eventually(t, func() bool {
		s, err := cl.Status(ctx)
		if err != nil {
			t.Log("Failed to get status: ", err)
			return false
		}

		return s.SyncInfo.LatestBlockHeight >= int64(target)
	}, time.Second*time.Duration(target*2), time.Millisecond*100)

	_, ok, err := cprov.LatestAttestation(ctx, 0) // Ensure it doesn't error for unknown chains.
	require.NoError(t, err)
	require.False(t, ok)

	xblock, ok, err := cprov.XBlock(ctx, 0, true) // Ensure getting latest xblock.
	require.NoError(t, err)
	require.True(t, ok)
	require.GreaterOrEqual(t, xblock.BlockHeight, uint64(1))
	require.Len(t, xblock.Msgs, 1)

	_, ok, err = cprov.ValidatorSet(ctx, 33) // Ensure it doesn't error for unknown validator sets.
	require.NoError(t, err)
	require.False(t, ok)

	genSet, err := cl.Validators(ctx, int64Ptr(1), nil, nil)
	require.NoError(t, err)
	getSetHash := types.NewValidatorSet(genSet.Validators).Hash()

	// Wait for cometBFT validator set to change
	require.Eventually(t, func() bool {
		set, err := cl.Validators(ctx, nil, nil, nil)
		require.NoError(t, err)
		setHash := types.NewValidatorSet(set.Validators).Hash()

		return !bytes.Equal(getSetHash, setHash)
	}, time.Second*time.Duration(target*2), time.Millisecond*100)

	srcChain := netconf.Simnet.Static().OmniExecutionChainID
	// Ensure all blocks are attested and approved.
	cprov.Subscribe(ctx, srcChain, 0, "test", func(ctx context.Context, approved xchain.Attestation) error {
		// Sanity check we can fetch latest directly as well.
		att, ok, err := cprov.LatestAttestation(ctx, srcChain)
		require.NoError(t, err)
		require.True(t, ok)
		require.Equal(t, srcChain, att.SourceChainID)

		require.Equal(t, srcChain, approved.SourceChainID)
		t.Logf("cprovider streamed approved block: %d", approved.BlockHeight)
		if approved.BlockHeight >= target {
			cancel()
		}

		return nil
	})

	<-ctx.Done()

	// Stop the server.
	require.NoError(t, stopfunc(ctx))
}

func setupSimnet(t *testing.T) haloapp.Config {
	t.Helper()
	homeDir := t.TempDir()

	cmtCfg := halocmd.DefaultCometConfig(homeDir)
	cmtCfg.BaseConfig.DBBackend = string(db.MemDBBackend)

	haloCfg := halocfg.DefaultConfig()
	haloCfg.HomeDir = homeDir
	haloCfg.BackendType = string(db.MemDBBackend)
	haloCfg.EVMBuildDelay = time.Millisecond

	cfg := haloapp.Config{
		Config: haloCfg,
		Comet:  cmtCfg,
	}

	err := halocmd.InitFiles(log.WithNoopLogger(context.Background()), halocmd.InitConfig{
		HomeDir: homeDir,
		Network: netconf.Simnet,
		Cosmos:  true,
	})
	tutil.RequireNoError(t, err)

	// CometBFT doesn't shutdown cleanly. It leaves goroutines running that write to disk.
	// The test sometimes fails with: TempDir RemoveAll cleanup: unlinkat ... directory not empty
	// Manually retry deleting everything a few times. This should prevent to test from flapping.
	t.Cleanup(func() {
		for i := 0; i < 5; i++ {
			err := os.RemoveAll(homeDir)
			if err == nil {
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
	})

	return cfg
}

func int64Ptr(i int64) *int64 {
	return &i
}
