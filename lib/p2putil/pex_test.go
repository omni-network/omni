package p2putil_test

import (
	"context"
	"flag"
	"testing"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/p2putil"

	"github.com/cometbft/cometbft/p2p"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "Include integration tests")

//nolint:tparallel // Concurrent output is hard to read, so do it sequentially.
func TestSeedPeers(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()
	ctx, err := log.Init(ctx, log.Config{
		Level:  "debug",
		Color:  "force",
		Format: "console",
	})
	require.NoError(t, err)

	for _, network := range []netconf.ID{netconf.Staging, netconf.Omega} {
		t.Run(network.String(), func(t *testing.T) {
			for _, seedNode := range network.Static().ConsensusSeeds() {
				t.Logf("Fetching peers from %s", seedNode)

				seedAddr, err := p2p.NewNetAddressString(seedNode)
				require.NoError(t, err)

				peers, err := p2putil.FetchPexAddrs(ctx, network, seedAddr)
				require.NoError(t, err)

				t.Logf("Fetched %d peers", len(peers))
				for i, peer := range peers {
					t.Logf("Peer %d: %s", i, peer)
				}
			}
		})
	}
}
