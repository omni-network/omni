package keeper_test

import (
	"testing"

	"github.com/omni-network/omni/halo/attest/keeper"
	"github.com/omni-network/omni/halo/attest/types"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

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
		name           string
		expectations   []expectation   // These functions set expectations in the various mocked dependencies.
		prerequisites  []prerequisite  // These functions modify the keeper before calling its Add method.
		postrequisites []postrequisite // These functions run additional checks at the end of the test.
		args           args
		want           want
		wantErr        bool
	}{
		{
			name: "single_vote",
			args: args{
				msg: defaultMsg().Msg(),
			},
			want: want{
				atts: []*keeper.Attestation{
					expectPendingAtt(1, defaultOffset, 1),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1, defaultOffset),
					expectValSig(2, 1, val2, defaultOffset),
				},
			},
			postrequisites: []postrequisite{func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
				t.Helper()
				expectAtt := expectPendingAtt(1, defaultOffset, 1)
				allAtts, err := k.ListAllAttestations(ctx, &types.ListAllAttestationsRequest{
					ChainId:    expectAtt.GetChainId(),
					ConfLevel:  expectAtt.GetConfLevel(),
					Status:     uint32(keeper.Status_Pending),
					FromOffset: defaultOffset,
				})

				require.NoError(t, err)
				require.Len(t, allAtts.Attestations, 1)
				require.Equal(t, allAtts.Attestations[0].AttestHeader.AttestOffset, expectAtt.GetAttestOffset())
			}},
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
					expectPendingAtt(1, defaultOffset, 1),
					update(
						expectPendingAtt(2, defaultOffset+1, 1),
						func(att *keeper.Attestation) {
							att.BlockHash = blockHashes[1].Bytes()
						},
					),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1, defaultOffset),
					expectValSig(2, 1, val2, defaultOffset),
					expectValSig(3, 2, val1, defaultOffset+1),
					expectValSig(4, 2, val3, defaultOffset+1),
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
					expectPendingAtt(1, defaultOffset, 1),
					update(
						expectPendingAtt(2, defaultOffset+1, 1),
						func(att *keeper.Attestation) {
							att.BlockHeight = defaultHeight
						},
					),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1, defaultOffset),
					expectValSig(2, 1, val2, defaultOffset),
					expectValSig(3, 2, val1, defaultOffset+1),
					expectValSig(4, 2, val3, defaultOffset+1),
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
					expectPendingAtt(1, defaultOffset, 1),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1, defaultOffset),
					expectValSig(2, 1, val2, defaultOffset),
					expectValSig(3, 1, val3, defaultOffset),
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
					expectPendingAtt(1, defaultOffset, 1),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1, defaultOffset),
					expectValSig(2, 1, val2, defaultOffset),
				},
			},
		},
		{
			name: "skip_mismatching_att_root_same_block_and_vals",
			args: args{
				msg: defaultMsg().
					WithVotes(
						// Update agg vote to a different att root but identical block, signed by same vals (double sign)
						defaultAggVote().
							WithMsgRoot(common.BytesToHash([]byte("different root"))).
							Vote(),
					).Msg(),
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()
					// Insert default agg vote first, so above message is a double sign
					msg := defaultMsg().Msg()
					err := k.Add(ctx, msg)
					require.NoError(t, err)
				},
			},
			want: want{
				atts: []*keeper.Attestation{
					expectPendingAtt(1, defaultOffset, 1), // Default agg vote resulting in pending attestation.
					update( // Update agg vote resulting in second att with different root
						expectPendingAtt(2, defaultOffset, 1),
						func(att *keeper.Attestation) {
							att.MsgRoot = common.BytesToHash([]byte("different root")).Bytes()
						},
					),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1, defaultOffset),
					expectValSig(2, 1, val2, defaultOffset),
					// Update agg vote's signatures are not added, since they are double signs
				},
			},
		},
		{
			name: "mismatching_att_root_same_block_diff_vals",
			args: args{
				msg: defaultMsg().
					WithVotes(
						// Update agg vote to a different att root but identical block, signed by val 3 only
						defaultAggVote().
							WithMsgRoot(common.BytesToHash([]byte("different root"))).
							WithSignatures(sigsTuples(val3)...).
							Vote(),
					).Msg(),
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()
					// Insert default agg vote first, same block but signed by val 1 and 2
					msg := defaultMsg().Msg()
					err := k.Add(ctx, msg)
					require.NoError(t, err)
				},
			},
			want: want{
				atts: []*keeper.Attestation{
					expectPendingAtt(1, defaultOffset, 1), // Default agg vote resulting in pending attestation.
					update( // Update agg vote resulting in second att with different root
						expectPendingAtt(2, defaultOffset, 1),
						func(att *keeper.Attestation) {
							att.MsgRoot = common.BytesToHash([]byte("different root")).Bytes()
						},
					),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1, defaultOffset),
					expectValSig(2, 1, val2, defaultOffset),
					expectValSig(3, 2, val3, defaultOffset), // Signature of updated agg vote by val 3
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

			for _, p := range tt.postrequisites {
				p(t, k, ctx)
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
					expectValSig(1, 1, val1, defaultOffset),
					expectValSig(2, 1, val2, defaultOffset),
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
					expectPendingAtt(1, defaultOffset, 1),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1, defaultOffset),
					expectValSig(2, 1, val2, defaultOffset),
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
					sig := expectValSig(0, 1, val3, defaultOffset)
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
					expectValSig(2, 1, val2, defaultOffset),
					expectValSig(3, 1, val3, defaultOffset),
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
					expectValSig(1, 1, val1, defaultOffset),
					expectValSig(2, 1, val2, defaultOffset),
					expectValSig(3, 2, val1, defaultOffset+1),
					expectValSig(4, 2, val2, defaultOffset+1),
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
					expectPendingAtt(2, defaultOffset+2, 1),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1, defaultOffset),
					expectValSig(2, 1, val2, defaultOffset),
					expectValSig(3, 2, val1, defaultOffset+2),
					expectValSig(4, 2, val2, defaultOffset+2),
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
				trimBehindCalled(),
				noFuzzyDeps(),
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
					expectApprovedAtt(4, defaultOffset+3, valset1_2, 19),
				},
				sigs: []*keeper.Signature{
					expectValSig(7, 4, val1, defaultOffset+3),
					expectValSig(8, 4, val2, defaultOffset+3),
				},
			},
		},
		{
			name: "dont_delete_latest",
			expectations: []expectation{
				namerCalled(1),
				defaultExpectations,
				activeSetQueried(9),
				activeSetQueried(10),
				trimBehindCalled(),
				noFuzzyDeps(),
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

					// Approve all 2 attestations so they're no longer pending
					err = k.Approve(ctx, toValSet(valset1_2))
					require.NoError(t, err)

					// Begin the block at height 20, which should cause the first attestations to be deleted, but not the second (since it is latest)
					err = k.BeginBlock(ctx.WithBlockHeight(initHeight + 10))
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2,
			},
			want: want{
				atts: []*keeper.Attestation{
					expectApprovedAtt(2, defaultOffset+1, valset1_2, 11),
				},
				sigs: []*keeper.Signature{
					expectValSig(3, 2, val1, defaultOffset+1),
					expectValSig(4, 2, val2, defaultOffset+1),
				},
			},
		},
		{
			name: "only_delete_old_pending",
			expectations: []expectation{
				namerCalled(1),
				defaultExpectations,
				activeSetQueried(9),
				activeSetQueried(10),
				activeSetQueried(11),
				trimBehindCalled(),
				noFuzzyDeps(),
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()

					otherHash := common.BytesToHash([]byte("other hash"))
					initHeight := int64(10)

					defaultOffsetMin1 := umath.SubtractOrZero(defaultOffset, 1)

					// Att 1 at defaultOffset-1 (and other hash) signed by val3 (so not approved below so stays pending)
					vote1 := defaultAggVote().WithBlockHash(otherHash).WithBlockOffset(defaultOffsetMin1).WithSignatures(sigsTuples(val3)...).Vote()
					msg1 := defaultMsg().Default().WithVotes(vote1).Msg()
					err := k.Add(ctx.WithBlockHeight(initHeight), msg1)
					require.NoError(t, err)

					// Att 2 at defaultOffset (and defaultHeight) signed by val1 and val2 (and approved below and is latest approved att)
					vote2 := defaultAggVote().WithBlockOffset(defaultOffset).Vote()
					msg2 := defaultMsg().Default().WithVotes(vote2).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+1), msg2)
					require.NoError(t, err)

					// Att 3 at defaultOffset+1 signed by val3 (so not approved below so stays pending)
					vote3 := defaultAggVote().WithBlockOffset(defaultOffset + 1).WithSignatures(sigsTuples(val3)...).Vote()
					msg3 := defaultMsg().Default().WithVotes(vote3).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+2), msg3)
					require.NoError(t, err)

					// Approve all 2nd attestation
					err = k.Approve(ctx, toValSet(valset1_2))
					require.NoError(t, err)

					// Begin the block at height 20, which should cause the first pending attestations to be deleted,
					// but not the second (since it is latest),
					// and not the last pending attestation (since it is after latest approved)
					err = k.BeginBlock(ctx.WithBlockHeight(initHeight + 10))
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2,
			},
			want: want{
				atts: []*keeper.Attestation{
					expectApprovedAtt(2, defaultOffset, valset1_2, 11),
					expectPendingAtt(3, defaultOffset+1, 12),
				},
				sigs: []*keeper.Signature{
					expectValSig(2, 2, val1, defaultOffset),
					expectValSig(3, 2, val2, defaultOffset),
					expectValSig(4, 3, val3, defaultOffset+1),
				},
			},
		},
		{
			name: "delete_fuzzy_first",
			expectations: []expectation{
				namerCalled(1),
				defaultExpectations,
				activeSetQueried(9),
				activeSetQueried(10),
				activeSetQueried(11),
				activeSetQueried(12),
				trimBehindCalled(),
				fuzzyDeps(1),
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()

					initHeight := int64(10)

					// Finalized att 1
					vote1 := defaultAggVote().Vote()
					msg1 := defaultMsg().Default().WithVotes(vote1).Msg()
					err := k.Add(ctx.WithBlockHeight(initHeight), msg1)
					require.NoError(t, err)

					// Finalized att 2
					vote2 := defaultAggVote().WithBlockOffset(defaultOffset + 1).Vote()
					msg2 := defaultMsg().Default().WithVotes(vote2).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+1), msg2)
					require.NoError(t, err)

					// Fuzzy att 3
					vote3 := defaultAggVote().WithBlockOffset(defaultOffset).WithFuzzy().Vote()
					msg3 := defaultMsg().Default().WithVotes(vote3).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+2), msg3)
					require.NoError(t, err)

					// Fuzzy att 4
					vote4 := defaultAggVote().WithBlockOffset(defaultOffset + 1).WithFuzzy().Vote()
					msg4 := defaultMsg().Default().WithVotes(vote4).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+3), msg4)
					require.NoError(t, err)

					// Approve all 4 attestations so they're no longer pending
					err = k.Approve(ctx, toValSet(valset1_2))
					require.NoError(t, err)

					// Begin the block at height 20,
					// which should cause the Fuzzy att 3 to be deleted (fuzzy deleted first),
					// but not the first (since it is finalized and should be deleted after the fuzzy),
					// or the last 2 (since it is the latest fuzzy and finalized atts)
					err = k.BeginBlock(ctx.WithBlockHeight(initHeight + 10))
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2,
			},
			want: want{
				atts: []*keeper.Attestation{
					expectApprovedAtt(1, defaultOffset, valset1_2, 10),
					expectApprovedAtt(2, defaultOffset+1, valset1_2, 11),
					expectFuzzyAtt(expectApprovedAtt(4, defaultOffset+1, valset1_2, 13)),
				},
				sigs: []*keeper.Signature{
					expectValSig(1, 1, val1, defaultOffset),
					expectValSig(2, 1, val2, defaultOffset),
					expectValSig(3, 2, val1, defaultOffset+1),
					expectValSig(4, 2, val2, defaultOffset+1),
					expectFuzzySig(expectValSig(7, 4, val1, defaultOffset+1)),
					expectFuzzySig(expectValSig(8, 4, val2, defaultOffset+1)),
				},
			},
		},
		{
			name: "delete_final_after_fuzzy",
			expectations: []expectation{
				namerCalled(1),
				defaultExpectations,
				activeSetQueried(9),
				activeSetQueried(10),
				activeSetQueried(11),
				activeSetQueried(12),
				trimBehindCalled(),
				fuzzyDeps(2),
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()

					initHeight := int64(10)

					// Same setup as 'delete_fuzzy_first' test

					// Finalized att 1
					vote1 := defaultAggVote().Vote()
					msg1 := defaultMsg().Default().WithVotes(vote1).Msg()
					err := k.Add(ctx.WithBlockHeight(initHeight), msg1)
					require.NoError(t, err)

					// Finalized att 2
					vote2 := defaultAggVote().WithBlockOffset(defaultOffset + 1).Vote()
					msg2 := defaultMsg().Default().WithVotes(vote2).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+1), msg2)
					require.NoError(t, err)

					// Fuzzy att 3
					vote3 := defaultAggVote().WithBlockOffset(defaultOffset).WithFuzzy().Vote()
					msg3 := defaultMsg().Default().WithVotes(vote3).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+2), msg3)
					require.NoError(t, err)

					// Fuzzy att 4
					vote4 := defaultAggVote().WithBlockOffset(defaultOffset + 1).WithFuzzy().Vote()
					msg4 := defaultMsg().Default().WithVotes(vote4).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+3), msg4)
					require.NoError(t, err)

					// Approve all 4 attestations so they're no longer pending
					err = k.Approve(ctx, toValSet(valset1_2))
					require.NoError(t, err)

					// Begin the block at height 20,
					// which should cause the Fuzzy att 3 to be deleted (fuzzy deleted first),
					// but not the first (since it is finalized and should be deleted after the fuzzy),
					// or the last 2 (since it is the latest fuzzy and finalized atts)
					err = k.BeginBlock(ctx.WithBlockHeight(initHeight + 10))
					require.NoError(t, err)

					// Now delete again

					// This time the finalized att 1 should be deleted, since it doesn't have a fuzzy att anymore
					err = k.BeginBlock(ctx.WithBlockHeight(initHeight + 11))
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2,
			},
			want: want{
				atts: []*keeper.Attestation{
					expectApprovedAtt(2, defaultOffset+1, valset1_2, 11),
					expectFuzzyAtt(expectApprovedAtt(4, defaultOffset+1, valset1_2, 13)),
				},
				sigs: []*keeper.Signature{
					expectValSig(3, 2, val1, defaultOffset+1),
					expectValSig(4, 2, val2, defaultOffset+1),
					expectFuzzySig(expectValSig(7, 4, val1, defaultOffset+1)),
					expectFuzzySig(expectValSig(8, 4, val2, defaultOffset+1)),
				},
			},
		},
		{
			name: "dont_delete_consensus_yet",
			expectations: []expectation{
				func(_ sdk.Context, m mocks) {
					m.namer.EXPECT().ChainName(xchain.ChainVersion{ID: consensusID, ConfLevel: xchain.ConfFinalized}).Times(1).Return("test-chain")
				},
				defaultExpectations,
				activeSetQueried(9),
				activeSetQueried(10),
				trimBehindCalled(),
				noFuzzyDeps(),
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()

					initHeight := int64(10)

					// Consensus att 1
					vote1 := defaultAggVote().WithChainID(consensusID).WithBlockOffset(defaultOffset).Vote()
					msg1 := defaultMsg().Default().WithVotes(vote1).Msg()
					err := k.Add(ctx.WithBlockHeight(initHeight), msg1)
					require.NoError(t, err)

					// Consensus att 2
					vote2 := defaultAggVote().WithChainID(consensusID).WithBlockOffset(defaultOffset + 1).Vote()
					msg2 := defaultMsg().Default().WithVotes(vote2).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+1), msg2)
					require.NoError(t, err)

					// Approve all 2 attestations so they're no longer pending
					err = k.Approve(ctx, toValSet(valset1_2))
					require.NoError(t, err)

					// Begin the block at height 14,
					// which should not delete attestations since they require cTrimLag which is 5.
					err = k.BeginBlock(ctx.WithBlockHeight(initHeight + trimLag))
					require.NoError(t, err)

					// See next test for it being deleted
				},
			},
			args: args{
				valset: valset1_2,
			},
			want: want{
				atts: []*keeper.Attestation{
					consensusAtt(expectApprovedAtt(1, defaultOffset, valset1_2, 10)),
					consensusAtt(expectApprovedAtt(2, defaultOffset+1, valset1_2, 11)),
				},
				sigs: []*keeper.Signature{
					consensusSig(expectValSig(1, 1, val1, defaultOffset)),
					consensusSig(expectValSig(2, 1, val2, defaultOffset)),
					consensusSig(expectValSig(3, 2, val1, defaultOffset+1)),
					consensusSig(expectValSig(4, 2, val2, defaultOffset+1)),
				},
			},
		}, {
			name: "delete_consensus",
			expectations: []expectation{
				func(_ sdk.Context, m mocks) {
					m.namer.EXPECT().ChainName(xchain.ChainVersion{ID: consensusID, ConfLevel: xchain.ConfFinalized}).Times(1).Return("test-chain")
				},
				defaultExpectations,
				activeSetQueried(9),
				activeSetQueried(10),
				trimBehindCalled(),
				noFuzzyDeps(),
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()

					initHeight := int64(10)

					// Same setup as "dont_delete_consensus_yet"

					// Consensus att 1
					vote1 := defaultAggVote().WithChainID(consensusID).WithBlockOffset(defaultOffset).Vote()
					msg1 := defaultMsg().Default().WithVotes(vote1).Msg()
					err := k.Add(ctx.WithBlockHeight(initHeight), msg1)
					require.NoError(t, err)

					// Consensus att 2
					vote2 := defaultAggVote().WithChainID(consensusID).WithBlockOffset(defaultOffset + 1).Vote()
					msg2 := defaultMsg().Default().WithVotes(vote2).Msg()
					err = k.Add(ctx.WithBlockHeight(initHeight+1), msg2)
					require.NoError(t, err)

					// Approve all 2 attestations so they're no longer pending
					err = k.Approve(ctx, toValSet(valset1_2))
					require.NoError(t, err)

					// Begin the block at height 15,
					// which should delete att Consensus att 1.
					err = k.BeginBlock(ctx.WithBlockHeight(initHeight + cTrimLag))
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2,
			},
			want: want{
				atts: []*keeper.Attestation{
					consensusAtt(expectApprovedAtt(2, defaultOffset+1, valset1_2, 11)),
				},
				sigs: []*keeper.Signature{
					consensusSig(expectValSig(3, 2, val1, defaultOffset+1)),
					consensusSig(expectValSig(4, 2, val2, defaultOffset+1)),
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

func expectValSig(id uint64, attID uint64, val *vtypes.Validator, offset uint64) *keeper.Signature {
	ethAddr, _ := val.EthereumAddress()
	return &keeper.Signature{
		Id:               id,
		Signature:        ethAddr.Bytes(),
		ValidatorAddress: ethAddr.Bytes(),
		AttId:            attID,
		ChainId:          defaultChainID,
		AttestOffset:     offset,
		ConfLevel:        defaultConfLevel,
	}
}

func expectPendingAtt(id uint64, offset uint64, createdHeight uint64) *keeper.Attestation {
	return &keeper.Attestation{
		Id:            id,
		MsgRoot:       msgRoot.Bytes(),
		ChainId:       defaultChainID,
		BlockHash:     blockHashes[0].Bytes(),
		AttestOffset:  offset,
		BlockHeight:   defaultHeight,
		CreatedHeight: createdHeight,
		ConfLevel:     defaultConfLevel,
		Status:        uint32(keeper.Status_Pending),
	}
}

func expectApprovedAtt(id uint64, offset uint64, valset *vtypes.ValidatorSetResponse, createdHeight uint64) *keeper.Attestation {
	return &keeper.Attestation{
		Id:             id,
		MsgRoot:        msgRoot.Bytes(),
		ChainId:        defaultChainID,
		BlockHash:      blockHashes[0].Bytes(),
		AttestOffset:   offset,
		BlockHeight:    defaultHeight,
		CreatedHeight:  createdHeight,
		ConfLevel:      defaultConfLevel,
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
		a := keeper.AttestationFromDB(atts[i], consensusID, nil)
		attRoot, _ := a.AttestationRoot()
		atts[i].AttestationRoot = attRoot[:]
	}

	return atts
}

func expectFuzzyAtt(att *keeper.Attestation) *keeper.Attestation {
	att.ConfLevel = uint32(xchain.ConfLatest)
	return att
}
func expectFuzzySig(sig *keeper.Signature) *keeper.Signature {
	sig.ConfLevel = uint32(xchain.ConfLatest)
	return sig
}

func consensusAtt(att *keeper.Attestation) *keeper.Attestation {
	att.ChainId = consensusID
	return att
}

func consensusSig(sig *keeper.Signature) *keeper.Signature {
	sig.ChainId = consensusID
	return sig
}
