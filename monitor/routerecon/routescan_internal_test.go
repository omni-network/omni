package routerecon

import (
	"context"
	"encoding/json"
	"flag"
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain/connect"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -integration -v -run=TestQueryLatestXChain

var integration = flag.Bool("integration", false, "run routescan integration tests")

func TestReconLag(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()
	conn, err := connect.New(ctx, netconf.Omega)
	require.NoError(t, err)

	for _, stream := range conn.Network.EVMStreams() {
		if stream.DestChainID == evmchain.IDArbSepolia || stream.SourceChainID == evmchain.IDArbSepolia {
			// Skip arb since fetching cross txs frmo routescan fails
			continue // TODO(corver): Remove when routescan adds support for arb_sepolia.
		}

		streamName := conn.Network.StreamName(stream)

		cursor, ok, err := conn.XProvider.GetSubmittedCursor(ctx, stream)
		require.NoError(t, err, streamName)
		if !ok {
			// Skip streams without submissions
			continue
		}

		crossTx, err := paginateLatestCrossTx(ctx, queryFilter{Stream: stream})
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
	resp, err := paginateLatestCrossTx(ctx, queryFilter{})
	require.NoError(t, err)
	require.NotEmpty(t, resp.ID)

	bz, err := json.MarshalIndent(resp, "", "  ")
	require.NoError(t, err)
	t.Log(string(bz))
}

// TestIntegrationFalse ensures the integration flag defaults to false.
func TestIntegrationFalse(t *testing.T) {
	t.Parallel()
	require.False(t, *integration)
}
