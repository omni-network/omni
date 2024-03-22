package relayer

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func Test_fromHeights(t *testing.T) {
	t.Parallel()

	state := NewEmptyState(filepath.Join(t.TempDir(), "state.json"))
	err := state.Persist(4, 1, 300)
	require.NoError(t, err)
	err = state.Persist(4, 2, 53)
	require.NoError(t, err)
	err = state.Persist(4, 3, 20)
	require.NoError(t, err)

	type args struct {
		cursors []xchain.StreamCursor
		chains  []netconf.Chain
		state   *State
	}
	tests := []struct {
		name string
		args args
		want map[uint64]uint64
	}{
		{
			name: "1", args: args{
				cursors: []xchain.StreamCursor{
					{StreamID: xchain.StreamID{SourceChainID: 1, DestChainID: 2}, SourceBlockHeight: 200},
					{StreamID: xchain.StreamID{SourceChainID: 2, DestChainID: 3}, SourceBlockHeight: 250},
				},
				chains: []netconf.Chain{{ID: 1}, {ID: 2}, {ID: 3}},
				state:  NewEmptyState(filepath.Join(t.TempDir(), "state.json")),
			}, want: map[uint64]uint64{
				1: 200,
				2: 250,
				3: 0,
			},
		},
		{
			name: "2", args: args{
				cursors: []xchain.StreamCursor{
					{StreamID: xchain.StreamID{SourceChainID: 1, DestChainID: 3}, SourceBlockHeight: 200},
					{StreamID: xchain.StreamID{SourceChainID: 2, DestChainID: 3}, SourceBlockHeight: 100},
				},
				chains: []netconf.Chain{{ID: 1}, {ID: 2, DeployHeight: 55}, {ID: 3}},
				state:  NewEmptyState(filepath.Join(t.TempDir(), "state.json")),
			}, want: map[uint64]uint64{
				1: 200,
				2: 100,
				3: 0,
			},
		},
		{
			name: "3", args: args{
				cursors: []xchain.StreamCursor{
					{StreamID: xchain.StreamID{SourceChainID: 1, DestChainID: 2}, SourceBlockHeight: 200},
				},
				chains: []netconf.Chain{{ID: 1}, {ID: 2, DeployHeight: 55}, {ID: 3}},
				state:  NewEmptyState(filepath.Join(t.TempDir(), "state.json")),
			}, want: map[uint64]uint64{
				1: 200,
				2: 55,
				3: 0,
			},
		},
		{
			name: "4",
			args: args{
				cursors: []xchain.StreamCursor{
					{StreamID: xchain.StreamID{SourceChainID: 1, DestChainID: 2}, SourceBlockHeight: 200},
				},
				chains: []netconf.Chain{{ID: 1}, {ID: 2, DeployHeight: 55}, {ID: 3}},
				state:  state,
			}, want: map[uint64]uint64{
				1: 300,
				2: 55,
				3: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := fromHeights(tt.args.cursors, netconf.Chain{ID: 4}, tt.args.chains, tt.args.state)
			require.Equal(t, tt.want, got)
		})
	}
}

var (
	_ cchain.Provider = (*mockProvider)(nil)
	_ xchain.Provider = (*mockXChainClient)(nil)
)

type mockXChainClient struct {
	GetBlockFn           func(context.Context, uint64, uint64) (xchain.Block, bool, error)
	GetSubmittedCursorFn func(context.Context, uint64, uint64) (xchain.StreamCursor, bool, error)
	GetEmittedCursorFn   func(context.Context, uint64, uint64) (xchain.StreamCursor, bool, error)
}

func (m *mockXChainClient) StreamAsync(context.Context, uint64, uint64, xchain.ProviderCallback) error {
	panic("unexpected")
}

func (m *mockXChainClient) StreamBlocks(context.Context, uint64, uint64, xchain.ProviderCallback) error {
	panic("unexpected")
}

func (m *mockXChainClient) GetBlock(ctx context.Context, chainID uint64, height uint64) (xchain.Block, bool, error) {
	return m.GetBlockFn(ctx, chainID, height)
}

func (m *mockXChainClient) GetSubmittedCursor(ctx context.Context, chainID uint64, sourceChain uint64,
) (xchain.StreamCursor, bool, error) {
	return m.GetSubmittedCursorFn(ctx, chainID, sourceChain)
}

func (m *mockXChainClient) GetEmittedCursor(ctx context.Context, srcChainID uint64, destChainID uint64,
) (xchain.StreamCursor, bool, error) {
	return m.GetEmittedCursorFn(ctx, srcChainID, destChainID)
}

type mockSender struct {
	SendTransactionFn func(ctx context.Context, submission xchain.Submission) error
}

func (m *mockSender) SendTransaction(ctx context.Context, submission xchain.Submission) error {
	return m.SendTransactionFn(ctx, submission)
}

type mockProvider struct {
	cchain.Provider
	SubscribeFn func(ctx context.Context, sourceChainID uint64, sourceHeight uint64, callback cchain.ProviderCallback)
}

func (m *mockProvider) Subscribe(ctx context.Context, sourceChainID uint64, sourceHeight uint64,
	_ string, callback cchain.ProviderCallback,
) {
	m.SubscribeFn(ctx, sourceChainID, sourceHeight, callback)
}
