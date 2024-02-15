package app_test

import (
	"context"
	"testing"
	"time"

	halo1 "github.com/omni-network/omni/halo/app"
	halo1cmd "github.com/omni-network/omni/halo/cmd"
	"github.com/omni-network/omni/halo2/app"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/stretchr/testify/require"
)

func TestSmoke(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := setupSimnet(t)

	// Start the server async
	go func() {
		err := app.Run(ctx, cfg)
		require.NoError(t, err)
	}()

	// Wait until we get to block 3.
	require.Eventually(t, func() bool {
		cl, err := rpchttp.New("http://localhost:26657", "/websocket")
		if err != nil {
			t.Log("Failed to dail: ", err)
			return false
		}

		s, err := cl.Status(ctx)
		if err != nil {
			t.Log("Failed to get status: ", err)
			return false
		}

		return s.SyncInfo.LatestBlockHeight >= 3
	}, time.Second*5, time.Millisecond*100)
}

func setupSimnet(t *testing.T) halo1.Config {
	t.Helper()
	homeDir := t.TempDir()

	cfg := halo1.Config{
		HaloConfig: halo1.DefaultHaloConfig(),
		Comet:      halo1cmd.DefaultCometConfig(homeDir),
	}
	cfg.HomeDir = homeDir

	err := halo1cmd.InitFiles(log.WithNoopLogger(context.Background()), halo1cmd.InitConfig{
		HomeDir: homeDir,
		Network: netconf.Simnet,
		Cosmos:  true,
	})
	require.NoError(t, err)

	return cfg
}
