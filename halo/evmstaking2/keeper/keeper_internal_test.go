package keeper

import (
	"testing"

	"github.com/omni-network/omni/lib/netconf"
	types "github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum/common"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	akeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest // Asserting insertion ids of sequential writes
func TestInsertAndDeleteEVMEvents(t *testing.T) {
	tests := []struct {
		name       string
		event      types.EVMEvent
		insertedID uint64
		height     int64
	}{
		{
			name: "Insert event with address [1,2,3]",
			event: types.EVMEvent{
				Address: []byte{1, 2, 3},
			},
			insertedID: 1,
			height:     0,
		},
		{
			name: "Insert event with address [2,3,4]",
			event: types.EVMEvent{
				Address: []byte{2, 3, 4},
			},
			insertedID: 2,
			height:     1,
		},
	}

	submissionDelay := int64(5)

	keeper, ctx := setupKeeper(t, submissionDelay, nil, nil, nil)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx = ctx.WithBlockHeight(test.height)

			err := keeper.Deliver(ctx, common.Hash{}, test.event)
			require.NoError(t, err)

			event, err := keeper.eventsTable.Get(ctx, test.insertedID)
			require.NoError(t, err)
			require.Equal(t, test.event.Address, event.GetEvent().Address)
		})
	}

	// Make sure no submission happens for heights in the range 2 to 4
	for h := int64(2); h < keeper.submissionDelay; h++ {
		ctx = ctx.WithBlockHeight(h)
		err := keeper.EndBlock(ctx)
		require.NoError(t, err)
	}

	// All events are present because we did not deliver them.
	for _, test := range tests {
		found, err := keeper.eventsTable.Has(ctx, test.insertedID)
		require.NoError(t, err)
		require.True(t, found)
	}

	// Now "execute" block number `submissionDelay`
	err := keeper.EndBlock(ctx.WithBlockHeight(submissionDelay))
	require.NoError(t, err)

	// All events are deleted now
	for _, test := range tests {
		found, err := keeper.eventsTable.Has(ctx, test.insertedID)
		require.NoError(t, err)
		require.False(t, found)
	}
}

func setupKeeper(
	t *testing.T,
	submissionDelay int64,
	aKeeper akeeper.AccountKeeperI,
	bKeeper bkeeper.Keeper,
	sKeeper *skeeper.Keeper,
) (*Keeper, sdk.Context) {
	t.Helper()

	key := storetypes.NewKVStoreKey(types.ModuleName)
	storeSvc := runtime.NewKVStoreService(key)
	ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))
	ctx = ctx.WithBlockHeight(1)
	ctx = ctx.WithChainID(netconf.Simnet.Static().OmniConsensusChainIDStr())

	k, err := NewKeeper(
		storeSvc,
		nil,
		aKeeper,
		bKeeper,
		sKeeper,
		submissionDelay,
	)
	require.NoError(t, err, "new keeper")

	return k, ctx
}
