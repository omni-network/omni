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

	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"

	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

//go:generate go test -run=TestKeeper_isNextProposer -count=100 -failfast

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
				incMoreTimes: 1,
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
			maxWithdrawalsPerBlock := uint64(32)
			keeper, err := NewKeeper(cdc, storeService, &mockEngine, txConfig, ap, frp, maxWithdrawalsPerBlock)
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

func TestWithdrawalsEqual(t *testing.T) {
	t.Parallel()

	addr1 := common.HexToAddress("0x1234567890123456789012345678901234567890")
	addr2 := common.HexToAddress("0x0987654321098765432109876543210987654321")

	tests := []struct {
		name string
		w1   []*etypes.Withdrawal
		w2   []*etypes.Withdrawal
		want bool
	}{
		{
			name: "both nil",
			w1:   nil,
			w2:   nil,
			want: true,
		},
		{
			name: "first nil",
			w1:   nil,
			w2:   []*etypes.Withdrawal{},
			want: true,
		},
		{
			name: "second nil",
			w1:   []*etypes.Withdrawal{},
			w2:   nil,
			want: true,
		},
		{
			name: "both empty",
			w1:   []*etypes.Withdrawal{},
			w2:   []*etypes.Withdrawal{},
			want: true,
		},
		{
			name: "different lengths",
			w1:   []*etypes.Withdrawal{{Index: 1}},
			w2:   []*etypes.Withdrawal{{Index: 1}, {Index: 2}},
			want: false,
		},
		{
			name: "one has nil element",
			w1:   []*etypes.Withdrawal{{Index: 1}, nil},
			w2:   []*etypes.Withdrawal{{Index: 1}, {Index: 2}},
			want: false,
		},
		{
			name: "other has nil element",
			w1:   []*etypes.Withdrawal{{Index: 1}, {Index: 2}},
			w2:   []*etypes.Withdrawal{{Index: 1}, nil},
			want: false,
		},
		{
			name: "both have nil elements",
			w1:   []*etypes.Withdrawal{{Index: 1}, nil},
			w2:   []*etypes.Withdrawal{{Index: 1}, nil},
			want: false,
		},
		{
			name: "identical",
			w1: []*etypes.Withdrawal{
				{Index: 1, Validator: 100, Address: addr1, Amount: 1000},
				{Index: 2, Validator: 200, Address: addr2, Amount: 2000},
			},
			w2: []*etypes.Withdrawal{
				{Index: 1, Validator: 100, Address: addr1, Amount: 1000},
				{Index: 2, Validator: 200, Address: addr2, Amount: 2000},
			},
			want: true,
		},
		{
			name: "different index",
			w1: []*etypes.Withdrawal{
				{Index: 1, Validator: 100, Address: addr1, Amount: 1000},
			},
			w2: []*etypes.Withdrawal{
				{Index: 2, Validator: 100, Address: addr1, Amount: 1000},
			},
			want: false,
		},
		{
			name: "different validator",
			w1: []*etypes.Withdrawal{
				{Index: 1, Validator: 100, Address: addr1, Amount: 1000},
			},
			w2: []*etypes.Withdrawal{
				{Index: 1, Validator: 200, Address: addr1, Amount: 1000},
			},
			want: false,
		},
		{
			name: "different address",
			w1: []*etypes.Withdrawal{
				{Index: 1, Validator: 100, Address: addr1, Amount: 1000},
			},
			w2: []*etypes.Withdrawal{
				{Index: 1, Validator: 100, Address: addr2, Amount: 1000},
			},
			want: false,
		},
		{
			name: "different amount",
			w1: []*etypes.Withdrawal{
				{Index: 1, Validator: 100, Address: addr1, Amount: 1000},
			},
			w2: []*etypes.Withdrawal{
				{Index: 1, Validator: 100, Address: addr1, Amount: 2000},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := withdrawalsEqual(tt.w1, tt.w2); got != tt.want {
				t.Errorf("withdrawalsEqual() = %v, want %v", got, tt.want)
			}
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
