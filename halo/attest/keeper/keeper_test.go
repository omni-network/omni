package keeper_test

import (
	"testing"

	"github.com/omni-network/omni/halo/attest/keeper"
	"github.com/omni-network/omni/halo/attest/types"

	cmttypes "github.com/cometbft/cometbft/types"

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
					{Id: 1, AttestationRoot: attRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: valAddr1[:]},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: valAddr2[:]},
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
					{Id: 1, AttestationRoot: attRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
					{Id: 2, AttestationRoot: attRoot, ChainId: 1, Hash: blockHashes[1].Bytes(), Height: 501, Status: int32(keeper.Status_Pending)},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: valAddr1[:]},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: valAddr2[:]},
					{Id: 3, AttId: 2, Signature: val1.Bytes(), ValidatorAddress: valAddr1[:]},
					{Id: 4, AttId: 2, Signature: val3.Bytes(), ValidatorAddress: valAddr3[:]},
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
					{Id: 1, AttestationRoot: attRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: valAddr1[:]},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: valAddr2[:]},
					{Id: 3, AttId: 1, Signature: val3.Bytes(), ValidatorAddress: valAddr3[:]},
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
					{Id: 1, AttestationRoot: attRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: valAddr1[:]},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: valAddr2[:]},
				},
			},
		},
		{
			name: "mismatching_block_root",
			args: args{
				msg: defaultMsg().
					WithVotes(
						defaultAggVote().
							WithAttestationRoot([]byte("different root")).
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
					{Id: 1, AttestationRoot: attRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
					{Id: 2, AttestationRoot: []byte("different root"), ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: valAddr1[:]},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: valAddr2[:]},
					{Id: 3, AttId: 2, Signature: val1.Bytes(), ValidatorAddress: valAddr1[:]},
					{Id: 4, AttId: 2, Signature: val2.Bytes(), ValidatorAddress: valAddr2[:]},
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

func TestKeeper_Approve(t *testing.T) {
	t.Parallel()

	// cmp transformation options to ignore private fields of proto generated types.
	var (
		atteCmpOpts = cmp.Options{cmpopts.IgnoreUnexported(keeper.Attestation{})}
		sigsCmpOpts = cmp.Options{cmpopts.IgnoreUnexported(keeper.Signature{})}
	)

	defaultExpectations := func(_ sdk.Context, m mocks) {
		m.voter.EXPECT().TrimBehind(gomock.Any()).Times(1).Return(0)
	}

	valset1_2_3 := cmttypes.NewValidatorSet([]*cmttypes.Validator{val1, val2, val3})
	valset1_2 := cmttypes.NewValidatorSet([]*cmttypes.Validator{val1, val2})
	valset2_3 := cmttypes.NewValidatorSet([]*cmttypes.Validator{val2, val3})

	_ = valset1_2_3
	_ = valset1_2
	type args struct {
		valset *cmttypes.ValidatorSet
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
			args: args{
				valset: nil,
			},
			wantErr: true,
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
					msg := defaultMsg().Msg()
					err := k.Add(ctx, msg)
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2,
			},
			want: want{
				atts: []*keeper.Attestation{
					{Id: 1, AttestationRoot: attRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Approved), ValidatorsHash: valset1_2.Hash()},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: valAddr1[:]},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: valAddr2[:]},
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
					msg := defaultMsg().Msg()
					err := k.Add(ctx, msg)
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2_3,
			},
			want: want{
				atts: []*keeper.Attestation{
					{Id: 1, AttestationRoot: attRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: valAddr1[:]},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: valAddr2[:]},
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
					sig := &keeper.Signature{AttId: 1, Signature: val3.Bytes(), ValidatorAddress: valAddr3[:]}
					err = k.SignatureTable().Insert(ctx, sig)
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset2_3,
			},
			want: want{
				atts: []*keeper.Attestation{
					{Id: 1, AttestationRoot: attRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Approved), ValidatorsHash: valset2_3.Hash()},
				},
				sigs: []*keeper.Signature{
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: valAddr2[:]},
					{Id: 3, AttId: 1, Signature: val3.Bytes(), ValidatorAddress: valAddr3[:]},
				},
			},
		},
		{
			name: "two_attestations_same_chain_skip_second",
			expectations: []expectation{
				defaultExpectations,
			},
			prerequisites: []prerequisite{
				func(t *testing.T, k *keeper.Keeper, ctx sdk.Context) {
					t.Helper()
					msg := defaultMsg().Msg() // has sigs of val1 and val2 - not approved
					err := k.Add(ctx, msg)
					require.NoError(t, err)

					vote := defaultAggVote().WithSignatures(sigsTuples(val2, val3)...).WithBlockHeight(501).Vote()
					msg2 := defaultMsg().Default().WithVotes(vote).Msg()
					err = k.Add(ctx, msg2)
					require.NoError(t, err)
				},
			},
			args: args{
				valset: valset1_2_3,
			},
			want: want{
				atts: []*keeper.Attestation{
					{Id: 1, AttestationRoot: attRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 500, Status: int32(keeper.Status_Pending)},
					{Id: 2, AttestationRoot: attRoot, ChainId: 1, Hash: blockHashes[0].Bytes(), Height: 501, Status: int32(keeper.Status_Pending)},
				},
				sigs: []*keeper.Signature{
					{Id: 1, AttId: 1, Signature: val1.Bytes(), ValidatorAddress: valAddr1[:]},
					{Id: 2, AttId: 1, Signature: val2.Bytes(), ValidatorAddress: valAddr2[:]},
					{Id: 3, AttId: 2, Signature: val2.Bytes(), ValidatorAddress: valAddr2[:]},
					{Id: 4, AttId: 2, Signature: val3.Bytes(), ValidatorAddress: valAddr3[:]},
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

			err := k.Approve(ctx, tt.args.valset)
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
