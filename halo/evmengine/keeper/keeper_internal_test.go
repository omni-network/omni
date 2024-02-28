package keeper

import (
	"context"
	"testing"
	"time"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	cmttypes "github.com/cometbft/cometbft/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/ethereum/go-ethereum/common"
	fuzz "github.com/google/gofuzz"
	"github.com/omni-network/omni/halo/comet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/stretchr/testify/require"
)

func TestKeeper_isNextProposer(t *testing.T) {
	ctx, storeService := setupCtxStore(t)
	cdc := getCodec(t)
	txConfig := authtx.NewTxConfig(cdc, nil)

	mockEngine, err := newMockEngineAPI()
	require.NoError(t, err)

	ap := mockAddressProvider{
		address: common.BytesToAddress([]byte("test")), // todo valid address
	}
	keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig, ap)
	keeper.cmtAPI = newMockCometAPI()

	keeper.isNextProposer(ctx)
}

type mockCometAPI struct {
	comet.API
	fuzzer *fuzz.Fuzzer
}

func newMockCometAPI() mockCometAPI {
	return mockCometAPI{
		fuzzer: NewFuzzer(0),
	}
}

func (m mockCometAPI) Validators(context.Context, int64) (*cmttypes.ValidatorSet, bool, error) {
	var validators []*cmttypes.Validator

	m.fuzzer.NilChance(0).NumElements(3, 7).Fuzz(&validators)

	valSet := new(cmttypes.ValidatorSet)
	if err := valSet.UpdateWithChangeSet(validators); err != nil {
		return nil, false, errors.Wrap(err, "update with change set")
	}

	return valSet, true, nil
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
