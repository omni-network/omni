package routerecon

import (
	"encoding/json"
	"flag"
	"net/http"
	"strconv"
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/lib/xchain/connect"

	"github.com/stretchr/testify/require"
)

const (
	rpdLimitHeader    = "X-Ratelimit-Rpd-Limit"
	rpmLimitHeader    = "X-Ratelimit-Rpm-Limit"
	increasedRPDLimit = 200000
	increasedRPMLimit = 600
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
	ctx := t.Context()
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

	ctx := t.Context()
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

	ctx := t.Context()

	responseHook = func(resp *http.Response) {
		rpdLimit, err := strconv.Atoi(resp.Header.Get(rpdLimitHeader))
		require.NoError(t, err)

		rpmLimit, err := strconv.Atoi(resp.Header.Get(rpmLimitHeader))
		require.NoError(t, err)

		hasLargeRateLimit := rpdLimit >= increasedRPDLimit && rpmLimit >= increasedRPMLimit
		if *apiKey != "" {
			require.True(t, hasLargeRateLimit)
		} else {
			require.False(t, hasLargeRateLimit)
		}
	}
	defer func() { responseHook = func(*http.Response) {} }()

	_, _, err := queryLatestCrossTx(ctx, netconf.Omega, *apiKey, queryFilter{}, "")
	require.NoError(t, err)

	_, _, err = queryLatestCrossTx(ctx, netconf.Mainnet, *apiKey, queryFilter{}, "")
	require.NoError(t, err)
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
