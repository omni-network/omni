package relayer

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func Test_fromOffsets(t *testing.T) {
	t.Parallel()

	const dstChain = uint64(4)
	const srcChain1 = uint64(1)
	const srcChain2 = uint64(2)
	const srcChain3 = uint64(3)
	stream1 := xchain.StreamID{SourceChainID: srcChain1, DestChainID: dstChain}
	stream2 := xchain.StreamID{SourceChainID: srcChain2, DestChainID: dstChain}
	stream3 := xchain.StreamID{SourceChainID: srcChain3, DestChainID: dstChain}
	allStreams := []xchain.StreamID{stream1, stream2, stream3}

	makeState := func(t *testing.T, offset1, offset2, offset3 uint64) *State {
		t.Helper()
		state := NewEmptyState(filepath.Join(t.TempDir(), "state.json"))

		if offset1 != 0 {
			err := state.Persist(dstChain, srcChain1, offset1)
			require.NoError(t, err)
		}

		if offset2 != 0 {
			err := state.Persist(dstChain, srcChain2, offset2)
			require.NoError(t, err)
		}

		if offset3 != 0 {
			err := state.Persist(dstChain, srcChain3, offset3)
			require.NoError(t, err)
		}

		return state
	}

	makeCursors := func(offset1, offset2, offset3 uint64) []xchain.StreamCursor {
		var resp []xchain.StreamCursor
		if offset1 != 0 {
			resp = append(resp, xchain.StreamCursor{StreamID: stream1, BlockOffset: offset1})
		}
		if offset2 != 0 {
			resp = append(resp, xchain.StreamCursor{StreamID: stream2, BlockOffset: offset2})
		}
		if offset3 != 0 {
			resp = append(resp, xchain.StreamCursor{StreamID: stream3, BlockOffset: offset3})
		}

		return resp
	}

	makeResult := func(offset1, offset2, offset3 uint64) map[xchain.StreamID]uint64 {
		return map[xchain.StreamID]uint64{
			stream1: offset1,
			stream2: offset2,
			stream3: offset3,
		}
	}

	tests := []struct {
		name    string
		onChain []xchain.StreamCursor
		onDisk  *State
		want    map[xchain.StreamID]uint64
	}{
		{
			name:    "on-disk empty, on-chain empty",
			onChain: makeCursors(0, 0, 0),
			onDisk:  makeState(t, 0, 0, 0),
			want:    makeResult(1, 1, 1), // All streams initialize at 1
		},
		{
			name:    "on-disk only",
			onChain: makeCursors(0, 0, 0),
			onDisk:  makeState(t, 5, 6, 7),
			want:    makeResult(5, 6, 7),
		},
		{
			name:    "on-disk higher",
			onChain: makeCursors(2, 2, 2),
			onDisk:  makeState(t, 5, 6, 7),
			want:    makeResult(5, 6, 7),
		},
		{
			name:    "on-chain only",
			onChain: makeCursors(5, 6, 7),
			onDisk:  makeState(t, 0, 0, 0),
			want:    makeResult(5, 6, 7),
		},
		{
			name:    "on-chain higher",
			onChain: makeCursors(5, 6, 7),
			onDisk:  makeState(t, 2, 2, 2),
			want:    makeResult(5, 6, 7),
		},
		{
			name:    "mix",
			onChain: makeCursors(1, 6, 0),
			onDisk:  makeState(t, 2, 5, 0),
			want:    makeResult(2, 6, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := fromOffsets(tt.onChain, allStreams, tt.onDisk)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

var (
	_ cchain.Provider = (*mockProvider)(nil)
	_ xchain.Provider = (*mockXChainClient)(nil)
)

type mockXChainClient struct {
	GetBlockFn           func(context.Context, uint64, uint64, uint64) (xchain.Block, bool, error)
	GetSubmittedCursorFn func(context.Context, xchain.StreamID) (xchain.StreamCursor, bool, error)
	GetEmittedCursorFn   func(context.Context, xchain.EmitRef, xchain.StreamID) (xchain.StreamCursor, bool, error)
}

func (m *mockXChainClient) StreamAsync(context.Context, uint64, uint64, uint64, xchain.ProviderCallback) error {
	panic("unexpected")
}

func (m *mockXChainClient) StreamAsyncNoOffset(context.Context, uint64, uint64, xchain.ProviderCallback) error {
	panic("unexpected")
}

func (m *mockXChainClient) StreamBlocks(context.Context, uint64, uint64, uint64, xchain.ProviderCallback) error {
	panic("unexpected")
}

func (m *mockXChainClient) StreamBlocksNoOffset(context.Context, uint64, uint64, xchain.ProviderCallback) error {
	panic("unexpected")
}

func (m *mockXChainClient) GetBlock(ctx context.Context, chainID uint64, height uint64, xOffset uint64) (xchain.Block, bool, error) {
	return m.GetBlockFn(ctx, chainID, height, xOffset)
}

func (m *mockXChainClient) GetSubmittedCursor(ctx context.Context, stream xchain.StreamID,
) (xchain.StreamCursor, bool, error) {
	return m.GetSubmittedCursorFn(ctx, stream)
}

func (m *mockXChainClient) GetEmittedCursor(ctx context.Context, ref xchain.EmitRef, stream xchain.StreamID,
) (xchain.StreamCursor, bool, error) {
	return m.GetEmittedCursorFn(ctx, ref, stream)
}

type mockSender struct {
	SendTransactionFn func(ctx context.Context, submission xchain.Submission) error
}

func (m *mockSender) SendTransaction(ctx context.Context, submission xchain.Submission) error {
	return m.SendTransactionFn(ctx, submission)
}

type mockProvider struct {
	cchain.Provider
	SubscribeFn func(ctx context.Context, sourceChainID uint64, conf xchain.ConfLevel, xBlockOffset uint64, callback cchain.ProviderCallback)
}

func (m *mockProvider) Subscribe(ctx context.Context, sourceChainID uint64, conf xchain.ConfLevel, xBlockOffset uint64,
	_ string, callback cchain.ProviderCallback,
) {
	m.SubscribeFn(ctx, sourceChainID, conf, xBlockOffset, callback)
}
