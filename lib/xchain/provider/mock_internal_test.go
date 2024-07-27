package provider

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func TestMock(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())

	const (
		chainID    = 123
		fromHeight = 0
		total      = 5
	)

	mock, err := NewMock(time.Millisecond, 0, nil)
	require.NoError(t, err)

	req := xchain.ProviderRequest{
		ChainID:   chainID,
		Height:    fromHeight,
		ConfLevel: xchain.ConfLatest,
	}
	var blocks []xchain.Block
	err = mock.StreamAsync(ctx, req, func(ctx context.Context, block xchain.Block) error {
		blocks = append(blocks, block)
		if len(blocks) == total {
			cancel()
		}

		return nil
	})
	require.NoError(t, err)

	<-ctx.Done()
	require.Len(t, blocks, total)

	// Just some very basic sanity checks
	assertMsgs(t, mock, blocks[0].Msgs, 0, 0)
	assertMsgs(t, mock, blocks[1].Msgs, 1, 0)
	assertMsgs(t, mock, blocks[2].Msgs, 1, 1)
	assertMsgs(t, mock, blocks[3].Msgs, 2, 1)
	assertMsgs(t, mock, blocks[4].Msgs, 0, 0)

	assertOffsets(t, blocks)
}

func assertOffsets(t *testing.T, blocks []xchain.Block) {
	t.Helper()
	sOffsets := make(map[xchain.StreamID]uint64)

	for _, block := range blocks {
		for _, msg := range block.Msgs {
			require.Equal(t, sOffsets[msg.StreamID], msg.StreamOffset)
			sOffsets[msg.StreamID]++
		}
	}
}

func assertMsgs(t *testing.T, mock *Mock, msgs []xchain.Msg, a, b int) {
	t.Helper()
	count := func(msgs []xchain.Msg, chainID uint64) int {
		var resp int
		for _, msg := range msgs {
			if msg.DestChainID == chainID {
				resp++
			}
		}

		return resp
	}

	require.Equal(t, a, count(msgs, mock.destChains[0]))
	require.Equal(t, b, count(msgs, mock.destChains[1]))
}
