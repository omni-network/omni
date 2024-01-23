package relayer_test

import (
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/xchain"
	relayer "github.com/omni-network/omni/relayer/app"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func Test_translateSubmission(t *testing.T) {
	t.Parallel()
	type args struct {
		submission xchain.Submission
	}
	var sub xchain.Submission
	fuzz.New().NilChance(0).Fuzz(&sub)

	msgs := make([]bindings.XChainMsg, 0, len(sub.Msgs))
	for _, msg := range sub.Msgs {
		msgs = append(msgs, bindings.XChainMsg{
			SourceChainId: msg.SourceChainID,
			DestChainId:   msg.DestChainID,
			StreamOffset:  msg.StreamOffset,
			Sender:        msg.SourceMsgSender,
			To:            msg.DestAddress,
			Data:          msg.Data,
			GasLimit:      msg.DestGasLimit,
		})
	}
	signatures := make([]bindings.XChainSigTuple, 0, len(sub.Signatures))
	for _, sig := range sub.Signatures {
		signatures = append(signatures, bindings.XChainSigTuple{
			ValidatorPubKey: sig.ValidatorPubKey[:],
			Signature:       sig.Signature[:],
		})
	}

	var xSub = bindings.XChainSubmission{
		AttestationRoot: sub.AttestationRoot,
		BlockHeader: bindings.XChainBlockHeader{
			SourceChainId: sub.BlockHeader.SourceChainID,
			BlockHeight:   sub.BlockHeader.BlockHeight,
			BlockHash:     sub.BlockHeader.BlockHash,
		},
		Msgs:       msgs,
		Proof:      sub.Proof,
		ProofFlags: sub.ProofFlags,
		Signatures: signatures,
	}

	tests := []struct {
		name string
		args args
		want bindings.XChainSubmission
	}{
		{
			name: "",
			args: args{
				submission: sub,
			},
			want: xSub,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := relayer.TranslateSubmission(tt.args.submission)
			require.Equal(t, got, tt.want)
		})
	}
}
