package app_test

import (
	"context"
	"testing"
	"time"

	halo1 "github.com/omni-network/omni/halo/app"
	halo1cmd "github.com/omni-network/omni/halo/cmd"
	"github.com/omni-network/omni/halo2/app"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	db "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

func TestSmoke(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := setupSimnet(t)

	spy := newSpyCtx(ctx)

	// Start the server async
	done := make(chan struct{})
	go func() {
		err := app.Run(spy, cfg)
		require.NoError(t, err)
		close(done)
	}()

	// Wait until the server is ready.
	select {
	case <-spy.SpyDone():
	case <-done:
		require.Fail(t, "server stopped before it was ready")
	}

	// Connect to the server.
	cl, err := rpchttp.New("http://localhost:26657", "/websocket")
	require.NoError(t, err)

	// Wait until we get to block 3.
	require.Eventually(t, func() bool {
		s, err := cl.Status(ctx)
		if err != nil {
			t.Log("Failed to get status: ", err)
			return false
		}

		return s.SyncInfo.LatestBlockHeight >= 3
	}, time.Second*5, time.Millisecond*100)

	cprov := cprovider.NewABCIProvider2(cl)

	aggs, err := cprov.ApprovedFrom(ctx, 999, 1)
	require.NoError(t, err)
	require.Empty(t, aggs)

	// Stop the server.
	cancel()
	<-done
}

//nolint:containedctx // This wrap a context explicitly.
type spyCtx struct {
	context.Context
	doneCalled chan struct{}
}

func newSpyCtx(ctx context.Context) *spyCtx {
	return &spyCtx{
		Context:    ctx,
		doneCalled: make(chan struct{}),
	}
}

func (s *spyCtx) Done() <-chan struct{} {
	close(s.doneCalled)
	return s.Context.Done()
}

func (s *spyCtx) SpyDone() <-chan struct{} {
	return s.doneCalled
}

func setupSimnet(t *testing.T) halo1.Config {
	t.Helper()
	homeDir := t.TempDir()

	cfg := halo1.Config{
		HaloConfig: halo1.DefaultHaloConfig(),
		Comet:      halo1cmd.DefaultCometConfig(homeDir),
	}
	cfg.HomeDir = homeDir
	cfg.BackendType = db.MemDBBackend

	err := halo1cmd.InitFiles(log.WithNoopLogger(context.Background()), halo1cmd.InitConfig{
		HomeDir: homeDir,
		Network: netconf.Simnet,
		Cosmos:  true,
	})
	require.NoError(t, err)

	return cfg
}
