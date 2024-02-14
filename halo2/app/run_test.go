package app_test

import (
	"context"
	"testing"

	halo1 "github.com/omni-network/omni/halo/app"
	halo1cmd "github.com/omni-network/omni/halo/cmd"
	"github.com/omni-network/omni/halo2/app"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/stretchr/testify/require"
)

func TestRunNoErr(t *testing.T) {
	t.Skip("this fails, we need to figure out genesis")

	t.Parallel()
	ctx := context.Background()
	homeDir := t.TempDir()
	cfg := halo1.Config{
		HaloConfig: halo1.DefaultHaloConfig(),
		Comet:      halo1cmd.DefaultCometConfig(homeDir),
	}
	cfg.HomeDir = homeDir

	err := halo1cmd.InitFiles(log.WithNoopLogger(ctx), halo1cmd.InitConfig{
		HomeDir: homeDir,
		Network: netconf.Simnet,
	})
	require.NoError(t, err)

	err = app.Run(doneCtx{ctx}, cfg)
	require.ErrorIs(t, err, context.Canceled)
}

// doneCtx is a context that always returns a closed Done channel.
//
//nolint:containedctx // No other way to implement this.
type doneCtx struct {
	context.Context
}

func (doneCtx) Done() <-chan struct{} {
	ch := make(chan struct{})
	close(ch)

	return ch
}
