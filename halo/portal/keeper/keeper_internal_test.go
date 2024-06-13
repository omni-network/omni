package keeper_test

import (
	"testing"

	"github.com/omni-network/omni/halo/portal/keeper"
	ptypes "github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func TestKeeper(t *testing.T) {
	t.Parallel()
	keeper, sdkCtx := keeper.SetupKeeper(t)

	addValSet := func(id uint64) {
		_, err := keeper.EmitMsg(sdkCtx, ptypes.MsgTypeValSet, id, xchain.BroadcastChainID, xchain.ShardBroadcast0)
		require.NoError(t, err)
	}

	const (
		typWithdrawal = ptypes.MsgType(99)
		omniEVM       = uint64(999)
	)
	addWithDrawal := func(id uint64) {
		_, err := keeper.EmitMsg(sdkCtx, typWithdrawal, id, omniEVM, xchain.ShardFinalized0)
		require.NoError(t, err)
	}

	offsets := make(map[xchain.StreamID]map[uint64]uint64)
	assert := func(t *testing.T, offset uint64, height uint64, valsets []uint64, withdrawals []uint64) {
		t.Helper()

		resp, err := keeper.Block(sdkCtx, &ptypes.BlockRequest{Id: offset})
		require.NoError(t, err)

		require.Equal(t, offset, resp.Id)
		require.Equal(t, height, resp.CreatedHeight)

		require.Len(t, resp.Msgs, len(valsets)+len(withdrawals))

		for _, msg := range resp.Msgs {
			streamID := xchain.StreamID{DestChainID: msg.DestChainId, ShardID: xchain.ShardID(msg.ShardId)}
			inner, ok := offsets[streamID]
			if !ok {
				inner = make(map[uint64]uint64)
				offsets[streamID] = inner
			}
			inner[msg.Id] = msg.StreamOffset

			switch ptypes.MsgType(msg.Type) {
			case ptypes.MsgTypeValSet:
				require.Contains(t, valsets, msg.MsgTypeId)
				require.Equal(t, xchain.BroadcastChainID, msg.DestChainId)
				require.EqualValues(t, xchain.ShardBroadcast0, msg.ShardId)
			case typWithdrawal:
				require.Contains(t, withdrawals, msg.MsgTypeId)
				require.Equal(t, omniEVM, msg.DestChainId)
				require.EqualValues(t, xchain.ShardFinalized0, msg.ShardId)
			default:
				t.Fatalf("unexpected message type: %v", msg.Type)
			}
		}
	}

	sdkCtx = sdkCtx.WithBlockHeight(11)
	addValSet(1)
	assert(t, 1, 11, []uint64{1}, nil)
	addValSet(2)
	assert(t, 1, 11, []uint64{1, 2}, nil)

	sdkCtx = sdkCtx.WithBlockHeight(22)
	addWithDrawal(1)
	addWithDrawal(2)
	addValSet(3)

	assert(t, 2, 22, []uint64{3}, []uint64{1, 2})

	require.EqualValues(t, map[xchain.StreamID]map[uint64]uint64{
		{DestChainID: xchain.BroadcastChainID, ShardID: xchain.ShardBroadcast0}: {
			1: 1,
			2: 2,
			5: 3,
		},
		{DestChainID: omniEVM, ShardID: xchain.ShardFinalized0}: {
			3: 1,
			4: 2,
		},
	}, offsets)
}
