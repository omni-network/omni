package app_test

import (
	"context"
	"testing"
	"time"

	haloapp "github.com/omni-network/omni/halo/app"
	halocmd "github.com/omni-network/omni/halo/cmd"
	halocfg "github.com/omni-network/omni/halo/config"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

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

	cprov := cprovider.NewABCIProvider(cl, nil)

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

	// Ensure all blocks are attested and approved.
	cprov.Subscribe(ctx, 999, 0, func(ctx context.Context, approved xchain.AggAttestation) error {
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
	cfg := haloapp.Config{
		Config: halocfg.DefaultConfig(),
		Comet:  cmtCfg,
	}
	cfg.HomeDir = homeDir

	err := halocmd.InitFiles(log.WithNoopLogger(context.Background()), halocmd.InitConfig{
		HomeDir: homeDir,
		Network: netconf.Simnet,
		Cosmos:  true,
	})
	require.NoError(t, err)

	return cfg
}
