package keeper_test

import (
	"testing"

	"github.com/omni-network/omni/halo/attest/keeper"
	"github.com/omni-network/omni/halo/attest/types"
	vtypes "github.com/omni-network/omni/halo/valsync/types"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
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
		expectations  []expectation  // These functions set expectations in the various mocked dependencies.
		prerequisites []prerequisite // These functions modify the keeper before calling its Add method.
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
					expectPendingAtt(1, defaultOffset),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1),
					expectValSig(2, 1, val2),
				},
			},
		},
		{
			name: "two_votes_diff_block_hashes",
			args: args{
				msg: defaultMsg().
					WithAppendVotes(
						defaultAggVote().WithBlockHeader(1, defaultOffset+1, defaultHeight, blockHashes[1]).WithSignatures(sigsTuples(val1, val3)...).Vote(),
					).
					Msg(),
			},
			want: want{
				atts: []*keeper.Attestation{
					expectPendingAtt(1, defaultOffset),
					update(
						expectPendingAtt(2, defaultOffset+1),
						func(att *keeper.Attestation) {
							att.BlockHash = blockHashes[1].Bytes()
						},
					),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1),
					expectValSig(2, 1, val2),
					expectValSig(3, 2, val1),
					expectValSig(4, 2, val3),
				},
			},
		},
		{
			name: "two_votes_diff_block_numbers",
			args: args{
				msg: defaultMsg().
					WithAppendVotes(
						defaultAggVote().WithBlockHeader(1, defaultOffset+1, defaultHeight, blockHashes[0]).WithSignatures(sigsTuples(val1, val3)...).Vote(),
					).
					Msg(),
			},
			want: want{
				atts: []*keeper.Attestation{
					expectPendingAtt(1, defaultOffset),
					update(
						expectPendingAtt(2, defaultOffset+1),
						func(att *keeper.Attestation) {
							att.BlockHeight = defaultHeight
						},
					),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1),
					expectValSig(2, 1, val2),
					expectValSig(3, 2, val1),
					expectValSig(4, 2, val3),
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
					expectPendingAtt(1, defaultOffset),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1),
					expectValSig(2, 1, val2),
					expectValSig(3, 1, val3),
				},
			},
		},
		{
			name: "add_same_vote_msg_twice",
			args: args{
				msg: defaultMsg().Msg(),
			},
			prerequisites: []prerequisite{
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
					expectPendingAtt(1, defaultOffset),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1),
					expectValSig(2, 1, val2),
				},
			},
		},
		{
			name: "mismatching_block_root",
			args: args{
				msg: defaultMsg().
					WithVotes(
						defaultAggVote().
							WithMsgRoot(common.BytesToHash([]byte("different root"))).
							Vote(),
					).Msg(),
			},
			prerequisites: []prerequisite{
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
					expectPendingAtt(1, defaultOffset),
					update(
						expectPendingAtt(2, defaultOffset),
						func(att *keeper.Attestation) {
							att.MsgRoot = common.BytesToHash([]byte("different root")).Bytes()
						},
					),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1),
					expectValSig(2, 1, val2),
					expectValSig(3, 2, val1),
					expectValSig(4, 2, val2),
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

			wantAtts := populateKeyHashes(tt.want.atts)
			if !cmp.Equal(wantAtts, gotAtts, atteCmpOpts) {
				t.Error(cmp.Diff(wantAtts, gotAtts, atteCmpOpts))
			}

			if !cmp.Equal(tt.want.sigs, gotSigs, sigsCmpOpts) {
				t.Error(cmp.Diff(tt.want.sigs, gotSigs, sigsCmpOpts))
			}
		})
	}
}

func TestKeeper_Approve(t *testing.T) {
	t.Parallel()

	// cmp transformation options to ignore private fields of proto generated types.
	var (
		atteCmpOpts = cmp.Options{cmpopts.IgnoreUnexported(keeper.Attestation{})}
		sigsCmpOpts = cmp.Options{cmpopts.IgnoreUnexported(keeper.Signature{})}
	)

	valset1_2_3 := newValSet(7, val1, val2, val3)
	valset1_2 := newValSet(8, val1, val2)
	valset2_3 := newValSet(9, val2, val3)

	defaultExpectations := func(_ sdk.Context, m mocks) {
		m.namer.EXPECT().ChainName(gomock.Any()).AnyTimes().Return("")
		m.voter.EXPECT().TrimBehind(gomock.Any()).Times(1).Return(0)
		m.valProvider.EXPECT().ActiveSetByHeight(gomock.Any(), uint64(0)).
			Return(valset1_2_3, nil).
			AnyTimes()
	}

	_ = valset1_2_3
	_ = valset1_2
	type args struct {
		valset *vtypes.ValidatorSetResponse
	}
	type want struct {
		atts []*keeper.Attestation
		sigs []*keeper.Signature
	}

	tests := []struct {
		name          string
		expectations  []expectation  // These functions set expectations in the various mocked dependencies.
		prerequisites []prerequisite // These functions modify the keeper before calling its methods.
		args          args
		want          want
		wantErr       bool
	}{
		{
			name: "nil_validator_set",
			expectations: []expectation{
				defaultExpectations,
			},
			args: args{
				valset: nil,
			},
		},
		{
			name: "single_attestation_two_validators_approve",
			expectations: []expectation{
				namerCalled(1),
				defaultExpectations,
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()
					msg := defaultMsg().Msg() // Signed by 1 and 2, but also approved by 1 and 2
					err := k.Add(ctx, msg)
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2,
			},
			want: want{
				atts: []*keeper.Attestation{
					expectApprovedAtt(1, defaultOffset, valset1_2, 1),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1),
					expectValSig(2, 1, val2),
				},
			},
		},
		{
			name: "single_attestation_no_quorum_not_approved",
			expectations: []expectation{
				defaultExpectations,
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()
					msg := defaultMsg().Msg() // Only signed by 1 and 2 (25), approved by 1,2,3 (40)
					err := k.Add(ctx, msg)
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2_3,
			},
			want: want{
				atts: []*keeper.Attestation{
					expectPendingAtt(1, defaultOffset),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1),
					expectValSig(2, 1, val2),
				},
			},
		},
		{
			name: "single_attestation_quorum_approved_1_sig_deleted",
			expectations: []expectation{
				namerCalled(1),
				defaultExpectations,
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()
					msg := defaultMsg().Msg() // has sigs from val1 and val2
					err := k.Add(ctx, msg)
					require.NoError(t, err)

					// add sig from val3
					ethAddr, err := val3.EthereumAddress()
					require.NoError(t, err)
					sig := &keeper.Signature{AttId: 1, Signature: ethAddr.Bytes(), ValidatorAddress: ethAddr.Bytes()}
					err = k.SignatureTable().Insert(ctx, sig)
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset2_3, // Approve from 2_3 only (so sig 1 is deleted)
			},
			want: want{
				atts: []*keeper.Attestation{
					expectApprovedAtt(1, defaultOffset, valset2_3, 1),
				},
				sigs: []*keeper.Signature{
					expectValSig(2, 1, val2),
					expectValSig(3, 1, val3),
				},
			},
		},
		{
			name: "sequential_attestations",
			expectations: []expectation{
				namerCalled(2),
				defaultExpectations,
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()
					vote1 := defaultAggVote().WithBlockOffset(defaultOffset).Vote()
					vote2 := defaultAggVote().WithBlockOffset(defaultOffset + 1).Vote()

					msg1 := defaultMsg().Default().WithVotes(vote1).Msg()
					err := k.Add(ctx, msg1)
					require.NoError(t, err)
					msg := defaultMsg().Default().WithVotes(vote2).Msg()
					err = k.Add(ctx, msg)
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2,
			},
			want: want{
				atts: []*keeper.Attestation{
					expectApprovedAtt(1, defaultOffset, valset1_2, 1),
					expectApprovedAtt(2, defaultOffset+1, valset1_2, 1),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1),
					expectValSig(2, 1, val2),
					expectValSig(3, 2, val1),
					expectValSig(4, 2, val2),
				},
			},
		},
		{
			name: "non_sequential_attestations",
			expectations: []expectation{
				namerCalled(1),
				defaultExpectations,
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()
					vote1 := defaultAggVote().WithBlockOffset(defaultOffset).Vote()
					vote3 := defaultAggVote().WithBlockOffset(defaultOffset + 2).Vote()

					msg1 := defaultMsg().Default().WithVotes(vote1).Msg()
					err := k.Add(ctx, msg1)
					require.NoError(t, err)
					msg3 := defaultMsg().Default().WithVotes(vote3).Msg()
					err = k.Add(ctx, msg3)
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2,
			},
			want: want{
				atts: []*keeper.Attestation{
					expectApprovedAtt(1, defaultOffset, valset1_2, 1),
					expectPendingAtt(2, defaultOffset+2),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1),
					expectValSig(2, 1, val2),
					expectValSig(3, 2, val1),
					expectValSig(4, 2, val2),
				},
			},
		},
		{
			name: "delete_old_attestations",
			expectations: []expectation{
				namerCalled(1),
				defaultExpectations,
				activeSetQueried(9),
				activeSetQueried(10),
				activeSetQueried(17),
				activeSetQueried(18),
				trimBehindCalled(1),
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()

					initHeight := int64(10)
					vote1 := defaultAggVote().WithBlockOffset(defaultOffset).Vote()
					msg1 := defaultMsg().Default().WithVotes(vote1).Msg()
					err := k.Add(ctx.WithBlockHeight(initHeight), msg1)
					require.NoError(t, err)

					vote2 := defaultAggVote().WithBlockOffset(defaultOffset + 1).Vote()
					msg2 := defaultMsg().Default().WithVotes(vote2).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+1), msg2)
					require.NoError(t, err)

					vote3 := defaultAggVote().WithBlockOffset(defaultOffset + 2).Vote()
					msg3 := defaultMsg().Default().WithVotes(vote3).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+8), msg3)
					require.NoError(t, err)

					vote4 := defaultAggVote().WithBlockOffset(defaultOffset + 3).Vote()
					msg4 := defaultMsg().Default().WithVotes(vote4).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+9), msg4)
					require.NoError(t, err)

					// Approve all four attestations so they're no longer pending
					err = k.Approve(ctx, toValSet(valset1_2))
					require.NoError(t, err)

					// Begin the block at height 20, which should cause the first 2 attestations to be deleted, but not the third and fourth
					err = k.BeginBlock(ctx.WithBlockHeight(initHeight + 10))
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2,
			},
			want: want{
				atts: []*keeper.Attestation{
					expectApprovedAtt(3, defaultOffset+2, valset1_2, 18),
					expectApprovedAtt(4, defaultOffset+3, valset1_2, 19),
				},
				sigs: []*keeper.Signature{
					expectValSig(5, 3, val1),
					expectValSig(6, 3, val2),
					expectValSig(7, 4, val1),
					expectValSig(8, 4, val2),
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

			err := k.Approve(ctx, toValSet(tt.args.valset))
			if (err != nil) != tt.wantErr {
				t.Errorf("keeper.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotAtts, gotSigs := dumpTables(t, ctx, k)

			wantAtts := populateKeyHashes(tt.want.atts)
			if !cmp.Equal(wantAtts, gotAtts, atteCmpOpts) {
				t.Error(cmp.Diff(wantAtts, gotAtts, atteCmpOpts))
			}

			if !cmp.Equal(tt.want.sigs, gotSigs, sigsCmpOpts) {
				t.Error(cmp.Diff(tt.want.sigs, gotSigs, sigsCmpOpts))
			}
		})
	}
}

func toValSet(valset *vtypes.ValidatorSetResponse) keeper.ValSet {
	if valset == nil {
		return keeper.ValSet{}
	}

	vals := make(map[common.Address]int64)
	for _, v := range valset.Validators {
		ethAddr, _ := v.EthereumAddress()
		vals[ethAddr] = v.Power
	}

	return keeper.ValSet{
		ID:   valset.Id,
		Vals: vals,
	}
}

func expectValSig(id uint64, attID uint64, val *vtypes.Validator) *keeper.Signature {
	ethAddr, _ := val.EthereumAddress()
	return &keeper.Signature{Id: id, AttId: attID, Signature: ethAddr.Bytes(), ValidatorAddress: ethAddr.Bytes()}
}

func expectPendingAtt(id uint64, offset uint64) *keeper.Attestation {
	return &keeper.Attestation{
		Id:            id,
		MsgRoot:       msgRoot.Bytes(),
		ChainId:       1,
		BlockHash:     blockHashes[0].Bytes(),
		BlockOffset:   offset,
		BlockHeight:   defaultHeight,
		CreatedHeight: 1,
		Status:        uint32(keeper.Status_Pending),
	}
}

func expectApprovedAtt(id uint64, offset uint64, valset *vtypes.ValidatorSetResponse, createdHeight uint64) *keeper.Attestation {
	return &keeper.Attestation{
		Id:             id,
		MsgRoot:        msgRoot.Bytes(),
		ChainId:        1,
		BlockHash:      blockHashes[0].Bytes(),
		BlockOffset:    offset,
		BlockHeight:    defaultHeight,
		CreatedHeight:  createdHeight,
		Status:         uint32(keeper.Status_Approved),
		ValidatorSetId: valset.Id,
	}
}

func update[T any](t T, fn func(T)) T {
	fn(t)

	return t
}

func populateKeyHashes(atts []*keeper.Attestation) []*keeper.Attestation {
	for i := range atts {
		a := keeper.AttestationFromDB(atts[i], nil)
		attRoot, _ := a.AttestationRoot()
		atts[i].AttestationRoot = attRoot[:]
	}

	return atts
}
