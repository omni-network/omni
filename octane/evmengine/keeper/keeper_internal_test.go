package keeper

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/comet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttypes "github.com/cometbft/cometbft/types"

	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestKeeper_isNextProposer(t *testing.T) {
	t.Parallel()
	type args struct {
		height         int64
		validatorsFunc func(context.Context, int64) (*cmttypes.ValidatorSet, error)
		incMoreTimes   int32
		header         func(height int64) cmtproto.Header
	}
	height := int64(1)
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "not proposer",
			args: args{
				height:       height,
				incMoreTimes: 9,
				header: func(height int64) cmtproto.Header {
					return cmtproto.Header{Height: height}
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "next proposer",
			args: args{
				height: height,
				header: func(height int64) cmtproto.Header {
					return cmtproto.Header{Height: height}
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "validatorsFunc error",
			args: args{
				height: height,
				validatorsFunc: func(ctx context.Context, i int64) (*cmttypes.ValidatorSet, error) {
					return nil, errors.New("error")
				},
				header: func(height int64) cmtproto.Header {
					return cmtproto.Header{Height: height}
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cdc := getCodec(t)
			txConfig := authtx.NewTxConfig(cdc, nil)
			mockEngine, err := newMockEngineAPI(0)
			require.NoError(t, err)

			cmtAPI := newMockCometAPI(t, tt.args.validatorsFunc)
			header := tt.args.header(height)

			nxtAddr, err := k1util.PubKeyToAddress(cmtAPI.validatorSet.CopyIncrementProposerPriority(1 + tt.args.incMoreTimes).Proposer.PubKey)
			require.NoError(t, err)

			ctx, storeService := setupCtxStore(t, &header)

			ap := mockAddressProvider{
				address: nxtAddr,
			}
			frp := newRandomFeeRecipientProvider()
			keeper, err := NewKeeper(cdc, storeService, &mockEngine, txConfig, ap, frp)
			require.NoError(t, err)
			keeper.SetCometAPI(cmtAPI)
			populateGenesisHead(ctx, t, keeper)

			got, err := keeper.isNextProposer(ctx, ctx.BlockHeader().Height)
			if (err != nil) != tt.wantErr {
				t.Errorf("isNextProposer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isNextProposer() got = %v, want %v", got, tt.want)
			}
			// make sure that height passed into Validators is correct
			require.Equal(t, tt.args.height, cmtAPI.height)
		})
	}
}

var _ comet.API = (*mockCometAPI)(nil)

type mockCometAPI struct {
	comet.API
	fuzzer         *fuzz.Fuzzer
	validatorSet   *cmttypes.ValidatorSet
	validatorsFunc func(context.Context, int64) (*cmttypes.ValidatorSet, error)
	height         int64
}

func newMockCometAPI(t *testing.T, valFun func(context.Context, int64) (*cmttypes.ValidatorSet, error)) *mockCometAPI {
	t.Helper()
	fuzzer := newFuzzer(0)
	valSet := fuzzValidators(t, fuzzer)

	return &mockCometAPI{
		fuzzer:         fuzzer,
		validatorSet:   valSet,
		validatorsFunc: valFun,
	}
}

func fuzzValidators(t *testing.T, fuzzer *fuzz.Fuzzer) *cmttypes.ValidatorSet {
	t.Helper()
	var validators []*cmttypes.Validator

	fuzzer.NilChance(0).NumElements(3, 7).Fuzz(&validators)

	valSet, err := cmttypes.ValidatorSetFromExistingValidators(validators)
	require.NoError(t, err)

	return valSet
}

func (m *mockCometAPI) Validators(ctx context.Context, height int64) (*cmttypes.ValidatorSet, error) {
	m.height = height
	if m.validatorsFunc != nil {
		return m.validatorsFunc(ctx, height)
	}

	return m.validatorSet, nil
}

// newFuzzer - create a new custom cmttypes.Validator fuzzer.
func newFuzzer(seed int64) *fuzz.Fuzzer {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	f := fuzz.NewWithSeed(seed).NilChance(0)
	f.Funcs(
		func(v *cmttypes.Validator, c fuzz.Continue) {
			privKey := k1.GenPrivKey()
			v.PubKey = privKey.PubKey()
			v.VotingPower = 200
			val := cmttypes.NewValidator(v.PubKey, v.VotingPower)

			*v = *val
		},
	)

	return f
}
