package p2putil_test

import (
	"context"
	"flag"
	"testing"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/p2putil"

	"github.com/cometbft/cometbft/p2p"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "Include integration tests")

// TestSeedP2PPeers fetches and prints P2P peers from the seed nodes of the specified networks.
//
//nolint:tparallel // Concurrent output is hard to read, so do it sequentially.
func TestSeedP2PPeers(t *testing.T) {
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
				t.Logf("Fetching P2P peers from %s", seedNode)

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

// TestRPCPeers fetches and prints RPC peers from the specified networks.
//
//nolint:tparallel // Concurrent output is hard to read, so do it sequentially.
func TestRPCPeers(t *testing.T) {
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

	for _, network := range []netconf.ID{netconf.Mainnet, netconf.Omega, netconf.Staging} {
		t.Run(network.String(), func(t *testing.T) {
			rpcServer := network.Static().ConsensusRPC()
			t.Logf("Fetching RPC peers from %s", rpcServer)

			rpcCl, err := rpchttp.New(rpcServer, "/websocket")
			require.NoError(t, err)

			info, err := rpcCl.NetInfo(ctx)
			require.NoError(t, err)

			t.Logf("Fetched %d peers", info.NPeers)
			for i, peer := range info.Peers {
				require.NoError(t, err)
				t.Logf("Peer %d: %s, %s, %s", i, peer.NodeInfo.Moniker, peer.RemoteIP, peer.NodeInfo.ID())
			}
		})
	}
}
