package keeper

import (
	"testing"

	"github.com/omni-network/omni/halo/attest/testutil"
	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"go.uber.org/mock/gomock"
)

type mocks struct {
	skeeper *testutil.MockStakingKeeper
	voter   *testutil.MockVoter
	namer   *testutil.MockChainNamer
}

func mockDefaultExpectations(_ sdk.Context, m mocks) {}

func setupKeeper(t *testing.T, expectations ...func(sdk.Context, mocks)) (*Keeper, sdk.Context) {
	t.Helper()

	key := storetypes.NewKVStoreKey(types.StoreKey)
	storeSvc := runtime.NewKVStoreService(key)
	ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))
	codec := moduletestutil.MakeTestEncodingConfig().Codec

	// gomock initialization
	ctrl := gomock.NewController(t)
	m := mocks{
		skeeper: testutil.NewMockStakingKeeper(ctrl),
		voter:   testutil.NewMockVoter(ctrl),
		namer:   testutil.NewMockChainNamer(ctrl),
	}
	if len(expectations) == 0 {
		mockDefaultExpectations(ctx, m)
	} else {
		for _, exp := range expectations {
			exp(ctx, m)
		}
	}

	const voteWindow = 1
	k, err := New(codec, storeSvc, m.skeeper, m.voter, m.namer.ChainName, voteWindow)
	if err != nil {
		t.Fatalf("error creating keeper: %v", err)
	}

	return k, ctx
}

// dumpTables returns all the items in the atestation and signature tables as slices.
func dumpTables(t *testing.T, k *Keeper, ctx sdk.Context) ([]*Attestation, []*Signature, error) {
	t.Helper()
	var atts []*Attestation
	lastAttID, err := k.attTable.LastInsertedSequence(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "last atestation id")
	}
	aitr, err := k.attTable.ListRange(ctx, AttestationIdIndexKey{}.WithId(0), AttestationIdIndexKey{}.WithId(lastAttID))
	if err != nil {
		return nil, nil, errors.Wrap(err, "iterate atestations")
	}
	defer aitr.Close()

	for aitr.Next() {
		a, err := aitr.Value()
		if err != nil {
			return nil, nil, errors.Wrap(err, "attestation iterator Value")
		}
		atts = append(atts, a)
	}

	var sigs []*Signature
	lastSigID, err := k.sigTable.LastInsertedSequence(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "last signature id")
	}
	sitr, err := k.sigTable.ListRange(ctx, SignatureIdIndexKey{}.WithId(0), SignatureIdIndexKey{}.WithId(lastSigID))
	if err != nil {
		return nil, nil, errors.Wrap(err, "iterate atestations")
	}
	defer sitr.Close()

	for sitr.Next() {
		s, err := sitr.Value()
		if err != nil {
			return nil, nil, errors.Wrap(err, "signature iterator Value")
		}
		sigs = append(sigs, s)
	}

	return atts, sigs, nil
}
