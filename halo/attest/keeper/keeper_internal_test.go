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
	"github.com/stretchr/testify/require"
)

// Hard-coded test data.
//
//nolint:gochecknoglobals // test data
var (
	blockHashes = []common.Hash{
		{
			0xb2, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
			0xa2, 0x18, 0xc6, 0xa9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x15,
			0xa2, 0x18, 0xc6, 0xa9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x16,
			0x10, 0x00,
		},
		{
			0xb1, 0x5f, 0x1b, 0x24, 0x4a, 0xab, 0x74, 0xac, 0xd6, 0x2a,
			0xb2, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
			0xb2, 0x28, 0xd6, 0xb9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x17,
			0x11, 0x01,
		},
		{
			0xb1, 0x5f, 0x1b, 0x24, 0x4a, 0xab, 0x74, 0xac, 0xd6, 0x2a,
			0xb2, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
			0xc2, 0x38, 0xe6, 0xc9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x18,
			0x12, 0x02,
		},
	}
	vals      = []k1.PrivKey{k1.GenPrivKey(), k1.GenPrivKey(), k1.GenPrivKey()}
	val1      = cmttypes.NewValidator(vals[0].PubKey(), 10)
	val2      = cmttypes.NewValidator(vals[1].PubKey(), 15)
	val3      = cmttypes.NewValidator(vals[2].PubKey(), 15)
	blockRoot = []byte{11, 5, 3, 67, 25, 34}
)

type MsgBuilder struct {
	msg *types.MsgAddVotes
}

func (b *MsgBuilder) WithAuthority(a string) *MsgBuilder {
	if b.msg == nil {
		b.msg = &types.MsgAddVotes{}
	}
	b.msg.Authority = a

	return b
}

func (b *MsgBuilder) WithVotes(votes ...*types.AggVote) *MsgBuilder {
	if b.msg == nil {
		b.msg = &types.MsgAddVotes{}
	}
	b.msg.Votes = votes

	return b
}

func (b *MsgBuilder) WithAppendVotes(votes ...*types.AggVote) *MsgBuilder {
	if b.msg == nil {
		b.msg = &types.MsgAddVotes{}
	}
	b.msg.Votes = append(b.msg.Votes, votes...)

	return b
}

func (b *MsgBuilder) Default() *MsgBuilder {
	b.msg = &types.MsgAddVotes{
		Authority: "test-authority",
		Votes: []*types.AggVote{
			new(AggVoteBuilder).Default().Vote(),
		},
	}

	return b
}

func (b *MsgBuilder) Msg() *types.MsgAddVotes {
	return b.msg
}

type AggVoteBuilder struct {
	vote *types.AggVote
}

func (b *AggVoteBuilder) Default() *AggVoteBuilder {
	b.vote = &types.AggVote{
		BlockHeader: &types.BlockHeader{
			ChainId: 1,
			Height:  500,
			Hash:    blockHashes[0].Bytes(),
		},
		BlockRoot:  blockRoot,
		Signatures: sigsTuples(val1, val2),
	}

	return b
}

func (b *AggVoteBuilder) WithChainID(id uint64) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{BlockHeader: &types.BlockHeader{}}
	} else if b.vote.BlockHeader == nil {
		b.vote.BlockHeader = &types.BlockHeader{}
	}
	b.vote.BlockHeader.ChainId = id

	return b
}

func (b *AggVoteBuilder) WithBlockHeight(h uint64) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{BlockHeader: &types.BlockHeader{}}
	} else if b.vote.BlockHeader == nil {
		b.vote.BlockHeader = &types.BlockHeader{}
	}
	b.vote.BlockHeader.Height = h

	return b
}

func (b *AggVoteBuilder) WithBlockHash(h common.Hash) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{BlockHeader: &types.BlockHeader{}}
	} else if b.vote.BlockHeader == nil {
		b.vote.BlockHeader = &types.BlockHeader{}
	}
	b.vote.BlockHeader.Hash = h.Bytes()

	return b
}

func (b *AggVoteBuilder) WithBlockHeader(chainID uint64, height uint64, hash common.Hash) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{BlockHeader: &types.BlockHeader{}}
	} else if b.vote.BlockHeader == nil {
		b.vote.BlockHeader = &types.BlockHeader{}
	}
	b.vote.BlockHeader.ChainId = chainID
	b.vote.BlockHeader.Height = height
	b.vote.BlockHeader.Hash = hash.Bytes()

	return b
}

func (b *AggVoteBuilder) WithBlockRoot(r []byte) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{}
	}
	b.vote.BlockRoot = r

	return b
}

func (b *AggVoteBuilder) WithSignatures(s ...*types.SigTuple) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{}
	}
	b.vote.Signatures = s

	return b
}

func (b *AggVoteBuilder) WithAppendSignatures(s ...*types.SigTuple) *AggVoteBuilder {
	if b.vote == nil {
		b.vote = &types.AggVote{}
	}
	b.vote.Signatures = append(b.vote.Signatures, s...)

	return b
}

func (b *AggVoteBuilder) Vote() *types.AggVote {
	return b.vote
}

func sigsTuples(vals ...*cmttypes.Validator) []*types.SigTuple {
	var sigs []*types.SigTuple
	for _, v := range vals {
		if v == nil {
			continue
		}
		sigs = append(sigs, &types.SigTuple{ValidatorAddress: v.Address, Signature: v.Bytes()})
	}

	return sigs
}

func TestKeeper_Add(t *testing.T) {
	t.Parallel()

	// cmp transformation options to ignore private fields of proto generated types.
	var (
		atteCmpOpts = cmp.Options{cmpopts.IgnoreUnexported(Attestation{})}
		sigsCmpOpts = cmp.Options{cmpopts.IgnoreUnexported(Signature{})}
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
				msg: new(MsgBuilder).Default().Msg(),
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
				msg: new(MsgBuilder).Default().
					WithAppendVotes(
						new(AggVoteBuilder).Default().WithBlockHeader(1, 501, blockHashes[1]).WithSignatures(sigsTuples(val1, val3)...).Vote(),
					).
					Msg(),
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
				msg: new(MsgBuilder).Default().
					WithVotes(
						new(AggVoteBuilder).Default().Vote(),
						new(AggVoteBuilder).Default().WithSignatures(sigsTuples(val2, val3)...).Vote(),
					).
					Msg(),
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
				msg: new(MsgBuilder).Default().Msg(),
			},
			prerequisites: []func(t *testing.T, k *Keeper, ctx sdk.Context){
				func(t *testing.T, k *Keeper, ctx sdk.Context) {
					t.Helper()
					// the same message as the one in the args
					msg := new(MsgBuilder).Default().Msg()
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
		{
			name: "mismatching_block_root",
			args: args{
				msg: new(MsgBuilder).Default().
					WithVotes(
						new(AggVoteBuilder).Default().
							WithBlockRoot([]byte("different root")). // the block root is intentionally different to cause an error
							Vote(),
					).Msg(),
			},
			prerequisites: []func(t *testing.T, k *Keeper, ctx sdk.Context){
				func(t *testing.T, k *Keeper, ctx sdk.Context) {
					t.Helper()
					// the same message as the one in the args
					msg := new(MsgBuilder).Default().Msg()
					err := k.Add(ctx, msg)
					require.NoError(t, err)
				},
			},
			wantErr: true,
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
