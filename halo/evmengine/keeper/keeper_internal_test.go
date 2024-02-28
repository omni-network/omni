package keeper

import (
	"context"
	"testing"
	"time"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttypes "github.com/cometbft/cometbft/types"
	cmttime "github.com/cometbft/cometbft/types/time"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	fuzz "github.com/google/gofuzz"
	"github.com/omni-network/omni/halo/comet"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/stretchr/testify/require"
)

func TestKeeper_isNextProposer(t *testing.T) {
	cdc := getCodec(t)
	txConfig := authtx.NewTxConfig(cdc, nil)

	mockEngine, err := newMockEngineAPI()
	require.NoError(t, err)

	cometAPI := newMockCometAPI()
	cometAPI.validatorSet = cometAPI.fuzzValidators(t)

	curr := cometAPI.validatorSet.Validators[0].Address

	nxtAddr, err := k1util.PubKeyToAddress(cometAPI.validatorSet.Validators[1].PubKey)
	require.NoError(t, err)
	ap := mockAddressProvider{
		address: nxtAddr,
	}
	ctx, storeService := setupCtxStore(t, cmtproto.Header{Time: cmttime.Now(), Height: 1, ProposerAddress: curr})

	keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig, ap)
	keeper.cmtAPI = cometAPI

	next, _, err := keeper.isNextProposer(ctx)
	require.NoError(t, err)
	require.True(t, next)
}

type mockCometAPI struct {
	comet.API
	fuzzer       *fuzz.Fuzzer
	validatorSet *cmttypes.ValidatorSet
}

func newMockCometAPI() mockCometAPI {
	return mockCometAPI{
		fuzzer: NewFuzzer(0),
	}
}

func (m mockCometAPI) fuzzValidators(t *testing.T) *cmttypes.ValidatorSet {
	t.Helper()
	var validators []*cmttypes.Validator

	m.fuzzer.NilChance(0).NumElements(3, 7).Fuzz(&validators)

	valSet := new(cmttypes.ValidatorSet)
	err := valSet.UpdateWithChangeSet(validators)
	require.NoError(t, err)

	return valSet

}

func (m mockCometAPI) Validators(context.Context, int64) (*cmttypes.ValidatorSet, bool, error) {
	return m.validatorSet, true, nil
}

// NewFuzzer returns a new fuzzer for valid ethereum types.
// If seed is zero, it uses current nano time as the seed.
func NewFuzzer(seed int64) *fuzz.Fuzzer {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	f := fuzz.NewWithSeed(seed).NilChance(0)
	f.Funcs(
		func(v *cmttypes.Validator, c fuzz.Continue) {
			privKey := k1.GenPrivKey()
			v.PubKey = privKey.PubKey()
			v.VotingPower = int64(c.Intn(200))
			val := cmttypes.NewValidator(v.PubKey, v.VotingPower)

			*v = *val
		},
	)

	return f
}
