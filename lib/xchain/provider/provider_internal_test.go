package provider

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

// NewForT returns a new provider for testing. Note that cprovider isn't supported yet.
func NewForT(
	t *testing.T,
	network netconf.Network,
	rpcClients map[uint64]ethclient.Client,
	backoffFunc func(ctx context.Context) func(),
	workers int) *Provider {
	t.Helper()

	for i := range fetchWorkerThresholds {
		fetchWorkerThresholds[i].Workers = uint64(workers)
	}

	return &Provider{
		network:     network,
		ethClients:  rpcClients,
		backoffFunc: backoffFunc,
		confHeads:   make(map[xchain.ChainVersion]uint64),
	}
}

//nolint:paralleltest // Access global thresholds not locked
func TestThresholds(t *testing.T) {
	for i, threshold := range fetchWorkerThresholds {
		if i == 0 {
			continue
		}

		// Ensure thresholds are in decreasing min period
		require.Greater(t, fetchWorkerThresholds[i-1].MinPeriod, threshold.MinPeriod)

		// Ensure workers are in increasing number
		require.Less(t, fetchWorkerThresholds[i-1].Workers, threshold.Workers)

		// Ensure last threshold is catch-all.
		if i == len(fetchWorkerThresholds)-1 {
			require.Empty(t, threshold.MinPeriod)
		}
	}
}
