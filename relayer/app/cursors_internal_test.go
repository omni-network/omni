package relayer

import (
	"context"
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
	chainVer1 := xchain.ChainVersion{ID: srcChain1, ConfLevel: xchain.ConfFinalized}
	chainVer2 := xchain.ChainVersion{ID: srcChain2, ConfLevel: xchain.ConfFinalized}
	chainVer3 := xchain.ChainVersion{ID: srcChain2, ConfLevel: xchain.ConfLatest}
	allChainVers := []xchain.ChainVersion{chainVer1, chainVer2, chainVer3}

	streamID := func(chainVer xchain.ChainVersion) xchain.StreamID {
		return xchain.StreamID{
			SourceChainID: chainVer.ID,
			DestChainID:   dstChain,
			ShardID:       xchain.ShardID(chainVer.ConfLevel),
		}
	}
	stream1 := streamID(chainVer1)
	stream2 := streamID(chainVer2)
	stream3 := streamID(chainVer3)

	makeCursors := func(offset1, offset2, offset3 uint64) []xchain.SubmitCursor {
		var resp []xchain.SubmitCursor
		if offset1 != 0 {
			resp = append(resp, xchain.SubmitCursor{StreamID: stream1, BlockOffset: offset1})
		}
		if offset2 != 0 {
			resp = append(resp, xchain.SubmitCursor{StreamID: stream2, BlockOffset: offset2})
		}
		if offset3 != 0 {
			resp = append(resp, xchain.SubmitCursor{StreamID: stream3, BlockOffset: offset3})
		}

		return resp
	}

	makeResult := func(offset1, offset2, offset3 uint64) map[xchain.ChainVersion]uint64 {
		return map[xchain.ChainVersion]uint64{
			chainVer1: offset1,
			chainVer2: offset2,
			chainVer3: offset3,
		}
	}

	tests := []struct {
		name    string
		onChain []xchain.SubmitCursor
		want    map[xchain.ChainVersion]uint64
	}{
		{
			name:    "on-chain empty",
			onChain: makeCursors(0, 0, 0),
			want:    makeResult(1, 1, 1), // All streams initialize at 1
		},
		{
			name:    "on-chain only",
			onChain: makeCursors(5, 6, 7),
			want:    makeResult(5, 6, 7),
		},
		{
			name:    "mix",
			onChain: makeCursors(1, 6, 0),
			want:    makeResult(1, 6, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := fromChainVersionOffsets(tt.onChain, allChainVers)
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
	GetBlockFn           func(context.Context, xchain.ProviderRequest) (xchain.Block, bool, error)
	GetSubmittedCursorFn func(context.Context, xchain.StreamID) (xchain.SubmitCursor, bool, error)
	GetEmittedCursorFn   func(context.Context, xchain.EmitRef, xchain.StreamID) (xchain.EmitCursor, bool, error)
}

func (m *mockXChainClient) StreamAsync(context.Context, xchain.ProviderRequest, xchain.ProviderCallback) error {
	panic("unexpected")
}

func (m *mockXChainClient) StreamBlocks(context.Context, xchain.ProviderRequest, xchain.ProviderCallback) error {
	panic("unexpected")
}

func (*mockXChainClient) ChainVersionHeight(context.Context, xchain.ChainVersion) (uint64, error) {
	panic("unexpected")
}

func (m *mockXChainClient) GetBlock(ctx context.Context, req xchain.ProviderRequest) (xchain.Block, bool, error) {
	return m.GetBlockFn(ctx, req)
}

func (m *mockXChainClient) GetSubmittedCursor(ctx context.Context, stream xchain.StreamID,
) (xchain.SubmitCursor, bool, error) {
	return m.GetSubmittedCursorFn(ctx, stream)
}

func (m *mockXChainClient) GetEmittedCursor(ctx context.Context, ref xchain.EmitRef, stream xchain.StreamID,
) (xchain.EmitCursor, bool, error) {
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
	SubscribeFn func(ctx context.Context, chainVer xchain.ChainVersion, xBlockOffset uint64, callback cchain.ProviderCallback)
}

func (m *mockProvider) Subscribe(ctx context.Context, chainVer xchain.ChainVersion, xBlockOffset uint64,
	_ string, callback cchain.ProviderCallback,
) {
	m.SubscribeFn(ctx, chainVer, xBlockOffset, callback)
}
