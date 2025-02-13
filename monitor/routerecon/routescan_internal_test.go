package routerecon

import (
	"context"
	"encoding/json"
	"flag"
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/lib/xchain/connect"

	"github.com/stretchr/testify/require"
)

var (
	integration     = flag.Bool("integration", false, "run routescan integration tests")
	routeScanAPIKey = flag.String("routeScanAPIKey", "", "RouteScan API key for enabling increased rate limiting")
)

//go:generate go test . -integration -v -run=TestQueryLatestXChain

func TestReconLag(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	network := netconf.Omega
	ctx := context.Background()
	conn, err := connect.New(ctx, network)
	require.NoError(t, err)

	for _, stream := range conn.Network.EVMStreams() {
		if stream.DestChainID == evmchain.IDArbSepolia || stream.SourceChainID == evmchain.IDArbSepolia {
			// Skip arb since fetching cross txs from routescan fails
			continue // TODO(corver): Remove when routescan adds support for arb_sepolia.
		}

		streamName := conn.Network.StreamName(stream)

		cursor, ok, err := conn.XProvider.GetSubmittedCursor(ctx, xchain.LatestRef, stream)
		require.NoError(t, err, streamName)
		if !ok {
			// Skip streams without submissions
			continue
		}

		crossTx, err := paginateLatestCrossTx(ctx, network, *routeScanAPIKey, queryFilter{Stream: stream})
		require.NoError(t, err, streamName)

		lag := float64(cursor.MsgOffset) - float64(crossTx.Data.Offset)
		log.Info(ctx, "Routescan lag", "stream", streamName, "submitted_offset", cursor.MsgOffset, "lag", lag)
	}
}

func TestQueryLatestXChain(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()
	resp, err := paginateLatestCrossTx(ctx, netconf.Mainnet, *routeScanAPIKey, queryFilter{})
	require.NoError(t, err)
	require.NotEmpty(t, resp.ID)

	bz, err := json.MarshalIndent(resp, "", "  ")
	require.NoError(t, err)
	t.Log(string(bz))
}

func TestIsWithIncreasedRateLimit(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()

	_, _, isWithIncreasedRateLimitOmega, err := queryLatestCrossTx(ctx, netconf.Omega, *routeScanAPIKey, queryFilter{}, "")
	require.NoError(t, err)

	if *routeScanAPIKey != "" {
		require.True(t, isWithIncreasedRateLimitOmega)
	} else {
		require.False(t, isWithIncreasedRateLimitOmega)
	}

	_, _, isWithIncreasedRateLimitMainnet, err := queryLatestCrossTx(ctx, netconf.Mainnet, *routeScanAPIKey, queryFilter{}, "")
	require.NoError(t, err)

	if *routeScanAPIKey != "" {
		require.True(t, isWithIncreasedRateLimitMainnet)
	} else {
		require.False(t, isWithIncreasedRateLimitMainnet)
	}
}

// TestIntegrationFalse ensures the integration flag defaults to false.
func TestIntegrationFalse(t *testing.T) {
	t.Parallel()
	require.False(t, *integration)
}

// TestRouteScanAPIKeyEmpty ensures the routeScanAPIKey flag defaults to empty string.
func TestRouteScanAPIKeyEmpty(t *testing.T) {
	t.Parallel()
	require.Empty(t, *routeScanAPIKey)
}
