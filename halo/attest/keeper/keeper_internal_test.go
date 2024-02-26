package keeper

import (
	"testing"

	"github.com/omni-network/omni/halo/attest/types"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	cmttypes "github.com/cometbft/cometbft/types"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestKeeper_Add(t *testing.T) {
	t.Parallel()

	// hard-coded test data
	var (
		blockHashes []common.Hash
		vals        = []k1.PrivKey{k1.GenPrivKey(), k1.GenPrivKey(), k1.GenPrivKey()}
		val1        = cmttypes.NewValidator(vals[0].PubKey(), 10)
		val2        = cmttypes.NewValidator(vals[1].PubKey(), 15)
		val3        = cmttypes.NewValidator(vals[2].PubKey(), 15)
		blockRoot   = []byte{11, 5, 3, 67, 25, 34}
	)
	fuzz.New().NilChance(0).NumElements(3, 3).Fuzz(&blockHashes)

	// cmp transformation options to ignore private fields of proto generated types.
	var (
		atteTrans = cmp.Options{cmpopts.IgnoreUnexported(Attestation{})}
		sigsTrans = cmp.Options{cmpopts.IgnoreUnexported(Signature{})}
	)

	type args struct {
		msg *types.MsgAddVotes
	}
	type want struct {
		atts []*Attestation
		sigs []*Signature
	}
	tests := []struct {
		name          string
		expectations  []func(sdk.Context, mocks)                       // These functions set expectations in the various mocked dependencies.
		prerequisites []func(t *testing.T, k *Keeper, ctx sdk.Context) // These functions modify the keeper before calling its Add method.
		args          args
		want          want
		wantErr       bool
	}{
		{
			name: "single_vote",
			args: args{
				msg: &types.MsgAddVotes{
					Authority: "test-authority",
					Votes: []*types.AggVote{
						{
							BlockHeader: &types.BlockHeader{
								ChainId: 1,
								Height:  500,
								Hash:    blockHashes[0].Bytes(),
							},
							BlockRoot: blockRoot,
							Signatures: []*types.SigTuple{
								{ValidatorAddress: val1.Address, Signature: val1.Bytes()},
								{ValidatorAddress: val2.Address, Signature: val2.Bytes()},
							},
						},
					},
				},
			},
			want: want{
				atts: []*Attestation{
					{Id: 1, BlockRoot: blockRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(Status_Pending)},
				},
				sigs: []*Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: val1.Address},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: val2.Address},
				},
			},
		},
		{
			name: "two_votes_diff_blocks",
			args: args{
				msg: &types.MsgAddVotes{
					Authority: "test-authority",
					Votes: []*types.AggVote{
						{
							BlockHeader: &types.BlockHeader{
								ChainId: 1,
								Height:  500,
								Hash:    blockHashes[0].Bytes(),
							},
							BlockRoot: blockRoot,
							Signatures: []*types.SigTuple{
								{ValidatorAddress: val1.Address, Signature: val1.Bytes()},
								{ValidatorAddress: val2.Address, Signature: val2.Bytes()},
							},
						},
						{
							BlockHeader: &types.BlockHeader{
								ChainId: 1,
								Height:  501,
								Hash:    blockHashes[1].Bytes(),
							},
							BlockRoot: blockRoot,
							Signatures: []*types.SigTuple{
								{ValidatorAddress: val1.Address, Signature: val1.Bytes()},
								{ValidatorAddress: val3.Address, Signature: val3.Bytes()},
							},
						},
					},
				},
			},
			want: want{
				atts: []*Attestation{
					{Id: 1, BlockRoot: blockRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(Status_Pending)},
					{Id: 2, BlockRoot: blockRoot, ChainId: 1, Hash: blockHashes[1].Bytes(), Height: 501, Status: int32(Status_Pending)},
				},
				sigs: []*Signature{
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
				msg: &types.MsgAddVotes{
					Authority: "test-authority",
					Votes: []*types.AggVote{
						{
							BlockHeader: &types.BlockHeader{
								ChainId: 1,
								Height:  500,
								Hash:    blockHashes[0].Bytes(),
							},
							BlockRoot: blockRoot,
							Signatures: []*types.SigTuple{
								{ValidatorAddress: val1.Address, Signature: val1.Bytes()},
								{ValidatorAddress: val2.Address, Signature: val2.Bytes()},
							},
						},
						{
							BlockHeader: &types.BlockHeader{
								ChainId: 1,
								Height:  500,
								Hash:    blockHashes[0].Bytes(),
							},
							BlockRoot: blockRoot,
							Signatures: []*types.SigTuple{
								{ValidatorAddress: val2.Address, Signature: val2.Bytes()},
								{ValidatorAddress: val3.Address, Signature: val3.Bytes()},
							},
						},
					},
				},
			},
			want: want{
				atts: []*Attestation{
					{Id: 1, BlockRoot: blockRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(Status_Pending)},
				},
				sigs: []*Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: val1.Address},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: val2.Address},
					{Id: 3, AttId: 1, Signature: val3.Bytes(), ValidatorAddress: val3.Address},
				},
			},
		},
		{
			name: "add_same_vote_msg_twice",
			args: args{
				msg: &types.MsgAddVotes{
					Authority: "test-authority",
					Votes: []*types.AggVote{
						{
							BlockHeader: &types.BlockHeader{
								ChainId: 1,
								Height:  500,
								Hash:    blockHashes[0].Bytes(),
							},
							BlockRoot: blockRoot,
							Signatures: []*types.SigTuple{
								{ValidatorAddress: val1.Address, Signature: val1.Bytes()},
								{ValidatorAddress: val2.Address, Signature: val2.Bytes()},
							},
						},
					},
				},
			},
			prerequisites: []func(t *testing.T, k *Keeper, ctx sdk.Context){
				func(t *testing.T, k *Keeper, ctx sdk.Context) {
					t.Helper()
					// the same message as the one in the args
					msg := &types.MsgAddVotes{
						Authority: "test-authority",
						Votes: []*types.AggVote{
							{
								BlockHeader: &types.BlockHeader{
									ChainId: 1,
									Height:  500,
									Hash:    blockHashes[0].Bytes(),
								},
								BlockRoot: blockRoot,
								Signatures: []*types.SigTuple{
									{ValidatorAddress: val1.Address, Signature: val1.Bytes()},
									{ValidatorAddress: val2.Address, Signature: val2.Bytes()},
								},
							},
						},
					}
					err := k.Add(ctx, msg)
					require.NoError(t, err)
				},
			},
			want: want{
				atts: []*Attestation{
					{Id: 1, BlockRoot: blockRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(Status_Pending)},
				},
				sigs: []*Signature{
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
			if err != nil {
				return
			}

			gotAtts, gotSigs, err := dumpTables(t, k, ctx)
			require.NoError(t, err, "dump orm tables")

			if !cmp.Equal(gotAtts, tt.want.atts, atteTrans) {
				t.Error(cmp.Diff(gotAtts, tt.want.atts, atteTrans))
			}

			if !cmp.Equal(gotSigs, tt.want.sigs, sigsTrans) {
				t.Error(cmp.Diff(gotSigs, tt.want.sigs, sigsTrans))
			}
		})
	}
}
