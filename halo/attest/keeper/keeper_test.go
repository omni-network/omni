package keeper_test

import (
	"testing"

	"github.com/omni-network/omni/halo/attest/keeper"
	"github.com/omni-network/omni/halo/attest/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

func TestKeeper_Add(t *testing.T) {
	t.Parallel()

	// cmp transformation options to ignore private fields of proto generated types.
	var (
		atteCmpOpts = cmp.Options{cmpopts.IgnoreUnexported(keeper.Attestation{})}
		sigsCmpOpts = cmp.Options{cmpopts.IgnoreUnexported(keeper.Signature{})}
	)

	type args struct {
		msg *types.MsgAddVotes
	}
	type want struct {
		atts []*keeper.Attestation
		sigs []*keeper.Signature
	}

	tests := []struct {
		name          string
		expectations  []func(sdk.Context, mocks)                              // These functions set expectations in the various mocked dependencies.
		prerequisites []func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) // These functions modify the keeper before calling its Add method.
		args          args
		want          want
		wantErr       bool
	}{
		{
			name: "single_vote",
			args: args{
				msg: defaultMsg().Msg(),
			},
			want: want{
				atts: []*keeper.Attestation{
					{Id: 1, BlockRoot: blockRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: val1.Address},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: val2.Address},
				},
			},
		},
		{
			name: "two_votes_diff_blocks",
			args: args{
				msg: defaultMsg().
					WithAppendVotes(
						defaultAggVote().WithBlockHeader(1, 501, blockHashes[1]).WithSignatures(sigsTuples(val1, val3)...).Vote(),
					).
					Msg(),
			},
			want: want{
				atts: []*keeper.Attestation{
					{Id: 1, BlockRoot: blockRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
					{Id: 2, BlockRoot: blockRoot, ChainId: 1, Hash: blockHashes[1].Bytes(), Height: 501, Status: int32(keeper.Status_Pending)},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: val1.Address},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: val2.Address},
					{Id: 3, AttId: 2, Signature: val1.Bytes(), ValidatorAddress: val1.Address},
					{Id: 4, AttId: 2, Signature: val3.Bytes(), ValidatorAddress: val3.Address},
				},
			},
		},
		{
			name: "two_votes_same_block_with_different_signatures",
			args: args{
				msg: defaultMsg().
					WithVotes(
						defaultAggVote().Vote(),
						defaultAggVote().WithSignatures(sigsTuples(val2, val3)...).Vote(),
					).
					Msg(),
			},
			want: want{
				atts: []*keeper.Attestation{
					{Id: 1, BlockRoot: blockRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: val1.Address},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: val2.Address},
					{Id: 3, AttId: 1, Signature: val3.Bytes(), ValidatorAddress: val3.Address},
				},
			},
		},
		{
			name: "add_same_vote_msg_twice",
			args: args{
				msg: defaultMsg().Msg(),
			},
			prerequisites: []func(t *testing.T, k *keeper.Keeper, ctx sdk.Context){
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()
					// the same message as the one in the args
					msg := defaultMsg().Msg()
					err := k.Add(ctx, msg)
					require.NoError(t, err)
				},
			},
			want: want{
				atts: []*keeper.Attestation{
					{Id: 1, BlockRoot: blockRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: val1.Address},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: val2.Address},
				},
			},
		},
		{
			name: "mismatching_block_root",
			args: args{
				msg: defaultMsg().
					WithVotes(
						defaultAggVote().
							WithBlockRoot([]byte("different root")). // the block root is intentionally different to cause an error
							Vote(),
					).Msg(),
			},
			prerequisites: []func(t *testing.T, k *keeper.Keeper, ctx sdk.Context){
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()
					// the same message as the one in the args
					msg := defaultMsg().Msg()
					err := k.Add(ctx, msg)
					require.NoError(t, err)
				},
			},
			wantErr: true,
			want: want{
				atts: []*keeper.Attestation{
					{Id: 1, BlockRoot: blockRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: val1.Address},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: val2.Address},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			k, ctx := setupKeeper(t, tt.expectations...)

			for _, p := range tt.prerequisites {
				p(t, k, ctx)
			}

			err := k.Add(ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("keeper.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotAtts, gotSigs := dumpTables(t, ctx, k)

			if !cmp.Equal(gotAtts, tt.want.atts, atteCmpOpts) {
				t.Error(cmp.Diff(gotAtts, tt.want.atts, atteCmpOpts))
			}

			if !cmp.Equal(gotSigs, tt.want.sigs, sigsCmpOpts) {
				t.Error(cmp.Diff(gotSigs, tt.want.sigs, sigsCmpOpts))
			}
		})
	}
}
