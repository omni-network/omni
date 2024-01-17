package relayer_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/xchain"
	relayer "github.com/omni-network/omni/relayer/app"

	"github.com/stretchr/testify/require"
)

func Test_StartRelayer(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())

	const (
		srcChain         = 1
		destChainA       = 2
		destChainB       = 3
		destChainACursor = 10 // ChainA is lagging
		destChainBCursor = 20 // ChainB is ahead
	)

	streamA := xchain.StreamID{
		SourceChainID: srcChain,
		DestChainID:   destChainA,
	}
	streamB := xchain.StreamID{
		SourceChainID: srcChain,
		DestChainID:   destChainB,
	}
	cursors := []xchain.StreamCursor{
		{StreamID: streamA, Offset: destChainACursor, SourceBlockHeight: destChainACursor},
		{StreamID: streamB, Offset: destChainBCursor, SourceBlockHeight: destChainBCursor},
	}

	// Return mock blocks (with a single msg per dest chain).
	mockXClient := &mockXChainClient{
		GetBlockFn: func(ctx context.Context, chainID uint64, height uint64) (xchain.Block, bool, error) {
			require.EqualValues(t, srcChain, chainID) // Only fetch blocks for source chains.

			// Each block has two messages, one for each stream.
			return xchain.Block{
				BlockHeader: xchain.BlockHeader{SourceChainID: chainID, BlockHeight: height},
				Msgs: []xchain.Msg{
					{MsgID: xchain.MsgID{StreamID: streamA, StreamOffset: height}},
					{MsgID: xchain.MsgID{StreamID: streamB, StreamOffset: height}},
				},
			}, true, nil
		},
		GetSubmittedCursorsFn: func(ctx context.Context, chainID uint64) ([]xchain.StreamCursor, error) {
			require.EqualValues(t, srcChain, chainID)

			return cursors, nil
		},
	}

	// Collect all stream updates via the creator, stop as soon as we get one msg from for streamB.
	var resp []relayer.StreamUpdate
	mockCreateFunc := func(streamUpdate relayer.StreamUpdate) ([]xchain.Submission, error) {
		resp = append(resp, streamUpdate)
		if streamUpdate.DestChainID == destChainB {
			cancel()
		}

		return nil, nil
	}

	// Sender should never be called, since we return empty slices from the creator.
	mockSender := &mockSender{
		SendTransactionFn: func(ctx context.Context, submission xchain.Submission) error {
			require.Fail(t, "should not be called")
			return nil
		},
	}

	// Provider mock attestations as requested until context canceled.
	mockProvider := &mockProvider{
		SubscribeFn: func(ctx context.Context, chainID uint64, fromHeight uint64, callback cchain.ProviderCallback) {
			require.EqualValues(t, srcChain, chainID)
			require.EqualValues(t, destChainACursor, fromHeight)

			height := fromHeight
			nextAtt := func() xchain.AggAttestation {
				defer func() { height++ }()
				return xchain.AggAttestation{
					BlockHeader: xchain.BlockHeader{SourceChainID: chainID, BlockHeight: height},
				}
			}

			for ctx.Err() == nil {
				err := callback(ctx, nextAtt())
				require.NoError(t, err)
			}
		},
	}

	relayer.StartRelayer(ctx, mockProvider, []uint64{srcChain}, mockXClient, mockCreateFunc, mockSender)

	// Verify responses
	expectChainA := destChainBCursor - destChainACursor + 1
	expectChainB := 1
	require.Len(t, resp, expectChainA+expectChainB)

	// Ensure msgs are delivered in sequence
	var actualChainA, actualChainB int
	prevChainA, prevChainB := destChainACursor, destChainBCursor
	for _, update := range resp {
		require.Len(t, update.Msgs, 1)
		next := update.Msgs[0].StreamOffset
		if update.DestChainID == destChainA {
			actualChainA++
			prevChainA++
			require.EqualValues(t, prevChainA, next)
		} else {
			actualChainB++
			prevChainB++
			require.EqualValues(t, prevChainB, next)
		}
	}

	// Ensure totals.
	require.EqualValues(t, expectChainA, actualChainA)
	require.EqualValues(t, expectChainB, actualChainB)
}

func Test_FromHeights(t *testing.T) {
	t.Parallel()
	type args struct {
		cursors  []xchain.StreamCursor
		chainIDs []uint64
	}
	tests := []struct {
		name string
		args args
		want map[uint64]uint64
	}{
		{
			name: "", args: args{
				cursors: []xchain.StreamCursor{
					{StreamID: xchain.StreamID{SourceChainID: 1, DestChainID: 2}, Offset: 100, SourceBlockHeight: 200},
					{StreamID: xchain.StreamID{SourceChainID: 2, DestChainID: 3}, Offset: 150, SourceBlockHeight: 250},
				},
				chainIDs: []uint64{1, 2, 3},
			}, want: map[uint64]uint64{
				1: 200,
				2: 250,
				3: 0,
			},
		},
		{
			name: "", args: args{
				cursors: []xchain.StreamCursor{
					{StreamID: xchain.StreamID{SourceChainID: 1, DestChainID: 2}, Offset: 100, SourceBlockHeight: 200},
					{StreamID: xchain.StreamID{SourceChainID: 1, DestChainID: 3}, Offset: 150, SourceBlockHeight: 0},
				},
				chainIDs: []uint64{1, 3},
			}, want: map[uint64]uint64{
				1: 0,
				3: 0,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := relayer.FromHeights(tt.args.cursors, tt.args.chainIDs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromHeights() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	_ cchain.Provider      = (*mockProvider)(nil)
	_ relayer.XChainClient = (*mockXChainClient)(nil)
	_ relayer.Sender       = (*mockSender)(nil)
)

type mockXChainClient struct {
	GetBlockFn            func(ctx context.Context, chainID uint64, height uint64) (xchain.Block, bool, error)
	GetSubmittedCursorsFn func(ctx context.Context, chainID uint64) ([]xchain.StreamCursor, error)
}

func (m *mockXChainClient) GetBlock(ctx context.Context, chainID uint64, height uint64) (xchain.Block, bool, error) {
	return m.GetBlockFn(ctx, chainID, height)
}

func (m *mockXChainClient) GetSubmittedCursors(ctx context.Context, chainID uint64) ([]xchain.StreamCursor, error) {
	return m.GetSubmittedCursorsFn(ctx, chainID)
}

type mockSender struct {
	SendTransactionFn func(ctx context.Context, submission xchain.Submission) error
}

func (m *mockSender) SendTransaction(ctx context.Context, submission xchain.Submission) error {
	return m.SendTransactionFn(ctx, submission)
}

type mockProvider struct {
	SubscribeFn func(ctx context.Context, sourceChainID uint64, sourceHeight uint64, callback cchain.ProviderCallback)
}

func (m *mockProvider) Subscribe(ctx context.Context, sourceChainID uint64, sourceHeight uint64,
	callback cchain.ProviderCallback,
) {
	m.SubscribeFn(ctx, sourceChainID, sourceHeight, callback)
}
