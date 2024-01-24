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
	var sub xchain.Submission
	fuzz.New().NilChance(0).Fuzz(&sub)

	xsub := relayer.TranslateSubmission(sub)
	reversedSub := translateXSubmission(xsub, sub.DestChainID)

	// Zero TxHash for comparison since it isn't translated.
	for i := range sub.Msgs {
		sub.Msgs[i].TxHash = [32]byte{}
	}

	require.Equal(t, sub, reversedSub)
}

func translateXSubmission(submission bindings.XChainSubmission, destChainID uint64) xchain.Submission {
	chainSubmission := xchain.Submission{
		AttestationRoot: submission.AttestationRoot,
		BlockHeader: xchain.BlockHeader{
			SourceChainID: submission.BlockHeader.SourceChainId,
			BlockHeight:   submission.BlockHeader.BlockHeight,
			BlockHash:     submission.BlockHeader.BlockHash,
		},
		Proof:       submission.Proof,
		ProofFlags:  submission.ProofFlags,
		DestChainID: destChainID,
	}

	signatures := make([]xchain.SigTuple, len(submission.Signatures))
	for i, sig := range submission.Signatures {
		signatures[i] = xchain.SigTuple{
			ValidatorPubKey: [33]byte(sig.ValidatorPubKey),
			Signature:       [65]byte(sig.Signature),
		}
	}

	chainSubmission.Signatures = signatures

	msgs := make([]xchain.Msg, len(submission.Msgs))
	for i, msg := range submission.Msgs {
		msgs[i] = xchain.Msg{
			MsgID: xchain.MsgID{
				StreamID: xchain.StreamID{
					SourceChainID: msg.SourceChainId,
					DestChainID:   msg.DestChainId,
				},
				StreamOffset: msg.StreamOffset,
			},
			SourceMsgSender: msg.Sender,
			DestAddress:     msg.To,
			Data:            msg.Data,
			DestGasLimit:    msg.GasLimit,
		}
	}

	chainSubmission.Msgs = msgs

	return chainSubmission
}
