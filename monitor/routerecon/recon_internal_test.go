package routerecon

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"
	xconnect "github.com/omni-network/omni/lib/xchain/connect"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -integration -v -run=TestReconOnce

func TestReconOnce(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()
	endpoints := xchain.RPCEndpoints{
		"omni_evm":     "https://omega.omni.network",
		"op_sepolia":   types.PublicRPCByName("op_sepolia"),
		"arb_sepolia":  types.PublicRPCByName("arb_sepolia"),
		"base_sepolia": types.PublicRPCByName("base_sepolia"),
		"holesky":      types.PublicRPCByName("holesky"),
	}
	conn, err := xconnect.New(ctx, netconf.Omega, endpoints)
	require.NoError(t, err)

	for _, stream := range conn.Network.EVMStreams() {
		err := reconStreamOnce(ctx, conn.Network, conn.XProvider, conn.EthClients, stream)
		if err != nil {
			log.Warn(ctx, "RouteRecon failed", err, "stream", conn.Network.StreamName(stream))
		} else {
			log.Info(ctx, "RouteRecon success", "stream", conn.Network.StreamName(stream))
		}
	}
}

//go:generate go test . -integration -v -run=TestBasicHistorical

// TestBasicHistorical does a basic recon of all completed/submitted historical routescan cross transactions.
// The analysis indicates gaps in attest offsets per stream.
func TestBasicHistorical(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()
	endpoints := xchain.RPCEndpoints{
		"omni_evm":     "https://omega.omni.network",
		"op_sepolia":   types.PublicRPCByName("op_sepolia"),
		"arb_sepolia":  types.PublicRPCByName("arb_sepolia"),
		"base_sepolia": types.PublicRPCByName("base_sepolia"),
		"holesky":      types.PublicRPCByName("holesky"),
	}
	conn, err := xconnect.New(ctx, netconf.Omega, endpoints)
	require.NoError(t, err)

	offsetsByStream := make(map[xchain.StreamID]map[uint64]crossTxJSON)
	callback := func(crossTx crossTxJSON) error {
		if crossTx.IsPending() {
			return nil // Skip pending (count as gaps)
		}

		msgID, err := crossTx.MsgID()
		require.NoError(t, err)
		streamName := conn.Network.StreamName(msgID.StreamID)

		offsets, ok := offsetsByStream[msgID.StreamID]
		if !ok {
			offsets = make(map[uint64]crossTxJSON)
			offsetsByStream[msgID.StreamID] = offsets
		}
		offsets[msgID.StreamOffset] = crossTx

		if len(offsets)%500 == 0 {
			log.Info(ctx, "Fetched offset", "stream", streamName, "count", len(offsets), "current", msgID.StreamOffset, "src_timestamp", crossTx.SrcTimestamp, "dst_timestamp", crossTx.DstTimestamp)
		}

		return nil
	}

	_, err = paginateLatestCrossTx(ctx, allCallback(callback))
	require.ErrorContains(t, err, "empty response") // Final pagination fails

	for _, streamID := range conn.Network.EVMStreams() {
		streamName := conn.Network.StreamName(streamID)
		offsets := offsetsByStream[streamID]

		var maxOffset uint64
		var maxCrossTx crossTxJSON
		for offset, crossTx := range offsets {
			if offset > maxOffset {
				maxOffset = offset
				maxCrossTx = crossTx
			}
		}
		var in bool
		var gaps int
		for i := uint64(1); i <= maxOffset; i++ {
			crossTx, ok := offsets[i]
			if in && ok {
				continue
			} else if !in && ok {
				log.Info(ctx, "Offsets started (incl)", "stream", streamName, "offset", i, "src_timestamp", fmtTime(crossTx.SrcTimestamp), "dst_timestamp", fmtTime(crossTx.DstTimestamp))
				in = true
			} else if !in && !ok {
				gaps++
				continue
			} else if in && !ok {
				gaps++
				log.Info(ctx, "Offsets ended (excl)", "stream", streamName, "offset", i)
				in = false
			}
		}

		log.Info(ctx, "Max routescan offset", "stream", streamName, "offset", maxOffset, "src_timestamp", fmtTime(maxCrossTx.SrcTimestamp), "dst_timestamp", fmtTime(maxCrossTx.DstTimestamp))

		sub, ok, err := conn.XProvider.GetSubmittedCursor(ctx, streamID)
		log.Info(ctx, "Max onchain offset", "stream", streamName, "offset", sub.MsgOffset, "ok", ok, "err", err)
		log.Info(ctx, "Missing routescan offsets", "stream", streamName, "indexed", len(offsets), "indexed_gaps", gaps, "unindexed", umath.SubtractOrZero(sub.MsgOffset, maxOffset))
	}
}

var _ filter = (allCallback)(nil)

type allCallback func(json crossTxJSON) error

func (allCallback) QueryParams(url.Values) {}

func (a allCallback) Match(crossTx crossTxJSON) (bool, error) {
	if err := a(crossTx); err != nil {
		return false, err
	}

	return false, nil
}

func fmtTime(t time.Time) string {
	return t.Format("01-02 15:04")
}
