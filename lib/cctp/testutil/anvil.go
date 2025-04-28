package testutil

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

// StartAnvilForks starts anvil forks fork each chain, returning a clients map and a "stop all" function.
func StartAnvilForks(t *testing.T, ctx context.Context, rpcs map[uint64]string, chains []evmchain.Metadata) (map[uint64]ethclient.Client, func()) {
	t.Helper()

	clients := make(map[uint64]ethclient.Client)

	var stops []func()
	for _, chain := range chains {
		ethCl, stop, err := anvil.Start(ctx, tutil.TempDir(t), chain.ChainID,
			anvil.WithFork(rpcs[chain.ChainID]),
			anvil.WithAutoImpersonate(),
			// quick finalization for testing
			anvil.WithBlockTime(1),
			anvil.WithSlotsInEpoch(2),
		)
		require.NoError(t, err)

		log.Info(ctx, "Stated anvil fork", "chain", chain.Name)

		clients[chain.ChainID] = ethCl
		stops = append(stops, stop)
	}

	stopAll := func() {
		for _, stop := range stops {
			stop()
		}
	}

	return clients, stopAll
}
