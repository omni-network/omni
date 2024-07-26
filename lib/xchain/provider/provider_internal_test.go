package provider

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
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
