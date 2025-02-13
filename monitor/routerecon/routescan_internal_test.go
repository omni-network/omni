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
	integration = flag.Bool("integration", false, "run RouteScan integration tests")
	apiKey      = flag.String("routeScanAPIKey", "", "RouteScan API key for enabling increased rate limiting")
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

		crossTx, err := paginateLatestCrossTx(ctx, network, *apiKey, queryFilter{Stream: stream})
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
	resp, err := paginateLatestCrossTx(ctx, netconf.Mainnet, *apiKey, queryFilter{})
	require.NoError(t, err)
	require.NotEmpty(t, resp.ID)

	bz, err := json.MarshalIndent(resp, "", "  ")
	require.NoError(t, err)
	t.Log(string(bz))
}

func TestHasLargeRateLimit(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()

	_, _, hasLargeRateLimitOmega, err := queryLatestCrossTx(ctx, netconf.Omega, *apiKey, queryFilter{}, "")
	require.NoError(t, err)

	if *apiKey != "" {
		require.True(t, hasLargeRateLimitOmega)
	} else {
		require.False(t, hasLargeRateLimitOmega)
	}

	_, _, hasLargeRateLimitMainnet, err := queryLatestCrossTx(ctx, netconf.Mainnet, *apiKey, queryFilter{}, "")
	require.NoError(t, err)

	if *apiKey != "" {
		require.True(t, hasLargeRateLimitMainnet)
	} else {
		require.False(t, hasLargeRateLimitMainnet)
	}
}

// TestIntegrationFalse ensures the integration flag defaults to false.
func TestIntegrationFalse(t *testing.T) {
	t.Parallel()
	require.False(t, *integration)
}

// TestRouteScanAPIKeyEmpty ensures the apiKey flag defaults to empty string.
func TestRouteScanAPIKeyEmpty(t *testing.T) {
	t.Parallel()
	require.Empty(t, *apiKey)
}
