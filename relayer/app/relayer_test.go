package relayer_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/xchain"
	relayer "github.com/omni-network/omni/relayer/app"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_StartRelayer(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())

	chainIDs := []uint64{1, 2}
	cursors := []xchain.StreamCursor{{StreamID: xchain.StreamID{
		SourceChainID: 1,
		DestChainID:   2,
	}, Offset: 0, SourceBlockHeight: 10}, {StreamID: xchain.StreamID{
		SourceChainID: 2,
		DestChainID:   1,
	}, Offset: 0, SourceBlockHeight: 20}}
	xBlock := xchain.Block{Msgs: []xchain.Msg{{
		MsgID: xchain.MsgID{
			StreamID: xchain.StreamID{
				SourceChainID: 1,
				DestChainID:   2,
			},
			StreamOffset: 1,
		},
	}}}
	aggAttestation := xchain.AggAttestation{
		BlockHeader: xchain.BlockHeader{
			SourceChainID: 1,
			BlockHeight:   2,
		},
	}

	// Mock client, creator, and sender
	mockXClient := &mockXChainClient{
		GetBlockFn: func(ctx context.Context, chainID uint64, height uint64) (xchain.Block, bool, error) {
			assert.Equal(t, aggAttestation.SourceChainID, chainID)
			assert.Equal(t, aggAttestation.BlockHeight, height)

			return xBlock, true, nil
		},
		GetSubmittedCursorsFn: func(ctx context.Context, chainID uint64) ([]xchain.StreamCursor, error) {
			assert.Contains(t, chainIDs, chainID)

			return cursors, nil
		},
	}

	mockCreator := &mockCreator{
		CreateSubmissionsFn: func(ctx context.Context, streamUpdate relayer.StreamUpdate) ([]xchain.Submission, error) {
			assert.Equal(t, aggAttestation, streamUpdate.AggAttestation)

			return []xchain.Submission{{}}, nil
		},
	}

	mockSender := &mockSender{
		SendTransactionFn: func(ctx context.Context, submission xchain.Submission) error {
			return nil
		},
	}

	mockProvider := &mockProvider{
		SubscribeFn: func(ctx context.Context, chainID uint64, fromHeight uint64, callback cchain.ProviderCallback) {
			assert.Contains(t, chainIDs, chainID)
			assert.NotNil(t, callback)

			// Simulate a callback with mock data
			err := callback(ctx, aggAttestation)
			if err == nil {
				cancel()
			}
			require.NoError(t, err)
		},
	}

	relayer.StartRelayer(ctx, mockProvider, chainIDs, mockXClient, mockCreator, mockSender)
	<-ctx.Done()
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
	_ relayer.Creator      = (*mockCreator)(nil)
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

type mockCreator struct {
	CreateSubmissionsFn func(ctx context.Context, streamUpdate relayer.StreamUpdate) ([]xchain.Submission, error)
}

func (m *mockCreator) CreateSubmissions(ctx context.Context,
	streamUpdate relayer.StreamUpdate,
) ([]xchain.Submission, error) {
	return m.CreateSubmissionsFn(ctx, streamUpdate)
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
