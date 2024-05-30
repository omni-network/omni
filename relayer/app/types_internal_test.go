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

	// Zero BlockHeight as we only submit BlockOffset
	sub.BlockHeader.BlockHeight = 0

	// TODO(corver): Add support for conf level to contracts and bindings
	sub.BlockHeader.ConfLevel = 0
	for i := range sub.Msgs {
		sub.Msgs[i].StreamID.ShardID = 0
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
				StreamOffset: msg.Offset,
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
			BlockOffset:   sub.BlockHeader.Offset,
			BlockHash:     sub.BlockHeader.SourceBlockHash,
		},
		Proof:       sub.Proof,
		ProofFlags:  sub.ProofFlags,
		DestChainID: destChainID,
		Signatures:  sigs,
		Msgs:        msgs,
	}
}
