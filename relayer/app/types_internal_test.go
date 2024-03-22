package relayer

import (
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func Test_translateSubmission(t *testing.T) {
	t.Parallel()
	var sub xchain.Submission
	fuzz.New().NilChance(0).Fuzz(&sub)

	xsub := submissionToBinding(sub)
	reversedSub := submissionFromBinding(xsub, sub.DestChainID)

	// Zero TxHash for comparison since it isn't translated.
	for i := range sub.Msgs {
		sub.Msgs[i].TxHash = common.Hash{}
	}

	require.Equal(t, sub, reversedSub)
}

func submissionFromBinding(sub bindings.XTypesSubmission, destChainID uint64) xchain.Submission {
	sigs := make([]xchain.SigTuple, 0, len(sub.Signatures))
	for _, sig := range sub.Signatures {
		sigs = append(sigs, xchain.SigTuple{
			ValidatorAddress: sig.ValidatorAddr,
			Signature:        xchain.Signature65(sig.Signature),
		})
	}

	msgs := make([]xchain.Msg, 0, len(sub.Msgs))
	for _, msg := range sub.Msgs {
		msgs = append(msgs, xchain.Msg{
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
		})
	}

	return xchain.Submission{
		AttestationRoot: sub.AttestationRoot,
		ValidatorSetID:  sub.ValidatorSetId,
		BlockHeader: xchain.BlockHeader{
			SourceChainID: sub.BlockHeader.SourceChainId,
			BlockHeight:   sub.BlockHeader.BlockHeight,
			BlockHash:     sub.BlockHeader.BlockHash,
		},
		Proof:       sub.Proof,
		ProofFlags:  sub.ProofFlags,
		DestChainID: destChainID,
		Signatures:  sigs,
		Msgs:        msgs,
	}
}
