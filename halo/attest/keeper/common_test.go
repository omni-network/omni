package keeper_test

import (
	"testing"

	"github.com/omni-network/omni/halo/attest/keeper"
	"github.com/omni-network/omni/halo/attest/testutil"
	"github.com/omni-network/omni/halo/attest/types"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type mocks struct {
	skeeper     *testutil.MockStakingKeeper
	voter       *testutil.MockVoter
	namer       *testutil.MockChainNamer
	valProvider *testutil.MockValProvider
}

type expectation func(sdk.Context, mocks)
type prerequisite func(t *testing.T, k *keeper.Keeper, ctx sdk.Context)

func mockDefaultExpectations(_ sdk.Context, m mocks) {
	m.namer.EXPECT().ChainName(uint64(1)).Return("test_chain").AnyTimes()
	m.valProvider.EXPECT().ActiveSetByHeight(gomock.Any(), uint64(0)).
		Return(newValSet(1, val1, val2, val3), nil).
		AnyTimes()
}

func namerCalled(times int) expectation {
	return func(_ sdk.Context, m mocks) {
		m.namer.EXPECT().ChainName(uint64(1)).Times(times).Return("test-chain")
	}
}

func setupKeeper(t *testing.T, expectations ...expectation) (*keeper.Keeper, sdk.Context) {
	t.Helper()

	key := storetypes.NewKVStoreKey(types.ModuleName)
	storeSvc := runtime.NewKVStoreService(key)
	ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))
	ctx = ctx.WithBlockHeight(1)
	codec := moduletestutil.MakeTestEncodingConfig().Codec

	// gomock initialization
	ctrl := gomock.NewController(t)
	m := mocks{
		skeeper:     testutil.NewMockStakingKeeper(ctrl),
		voter:       testutil.NewMockVoter(ctrl),
		namer:       testutil.NewMockChainNamer(ctrl),
		valProvider: testutil.NewMockValProvider(ctrl),
	}

	if len(expectations) == 0 {
		mockDefaultExpectations(ctx, m)
	} else {
		for _, exp := range expectations {
			exp(ctx, m)
		}
	}

	const voteWindow = 1
	const voteLimit = 4
	const trimLag = 4
	k, err := keeper.New(codec, storeSvc, m.skeeper, m.namer.ChainName, voteWindow, voteLimit, trimLag)
	require.NoError(t, err, "new keeper")
	k.SetVoter(m.voter)

	k.SetValidatorProvider(m.valProvider)

	return k, ctx
}

// dumpTables returns all the items in the atestation and signature tables as slices.
func dumpTables(t *testing.T, ctx sdk.Context, k *keeper.Keeper) ([]*keeper.Attestation, []*keeper.Signature) {
	t.Helper()
	var atts []*keeper.Attestation
	aitr, err := k.AttestTable().List(ctx, keeper.AttestationIdIndexKey{})
	require.NoError(t, err, "list attestations")
	defer aitr.Close()

	for aitr.Next() {
		a, err := aitr.Value()
		require.NoError(t, err, "signature iterator Value")
		atts = append(atts, a)
	}

	var sigs []*keeper.Signature
	sitr, err := k.SignatureTable().List(ctx, keeper.SignatureIdIndexKey{})
	require.NoError(t, err, "list signatures")
	defer sitr.Close()

	for sitr.Next() {
		s, err := sitr.Value()
		require.NoError(t, err, "signature iterator Value")
		sigs = append(sigs, s)
	}

	return atts, sigs
}
