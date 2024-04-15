package keeper

import (
	"testing"

	"github.com/omni-network/omni/halo/valsync/testutil"
	"github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/netconf"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type mocks struct {
	sKeeper    *testutil.MockStakingKeeper
	aKeeper    *testutil.MockAttestKeeper
	subscriber *testutil.MockSubscriber
}

type expectation func(sdk.Context, mocks)

func setupKeeper(t *testing.T, expectations ...expectation) (*Keeper, sdk.Context) {
	t.Helper()

	key := storetypes.NewKVStoreKey(types.ModuleName)
	storeSvc := runtime.NewKVStoreService(key)
	ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))
	ctx = ctx.WithBlockHeight(1)
	ctx = ctx.WithChainID(netconf.Simnet.Static().OmniConsensusChainIDStr())
	codec := moduletestutil.MakeTestEncodingConfig().Codec

	// gomock initialization
	ctrl := gomock.NewController(t)
	m := mocks{
		sKeeper:    testutil.NewMockStakingKeeper(ctrl),
		aKeeper:    testutil.NewMockAttestKeeper(ctrl),
		subscriber: testutil.NewMockSubscriber(ctrl),
	}

	for _, exp := range expectations {
		exp(ctx, m)
	}

	k, err := NewKeeper(codec, storeSvc, m.sKeeper, m.aKeeper)
	require.NoError(t, err, "new keeper")
	k.SetSubscriber(m.subscriber)

	return k, ctx
}
