package relayer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/xchain"
)

var _ Sender = (*SenderService)(nil)

type SenderService struct {
}

func (s SenderService) SendTransaction(ctx context.Context, submission xchain.Submission) error {
	//TODO implement me
	panic("implement me")
}

func translateSubmission(submission xchain.Submission) bindings.XChainSubmission {
	chainSubmission := bindings.XChainSubmission{
		AttestationRoot: submission.AttestationRoot,
		BlockHeader: bindings.XChainBlockHeader{
			SourceChainId: submission.BlockHeader.SourceChainID,
			BlockHeight:   submission.BlockHeader.BlockHeight,
			BlockHash:     submission.BlockHeader.BlockHash,
		},
		Proof:      submission.Proof,
		ProofFlags: submission.ProofFlags,
	}

	signatures := make([]bindings.XChainSigTuple, len(submission.Signatures))
	for i, sig := range submission.Signatures {
		signatures[i] = bindings.XChainSigTuple{
			ValidatorPubKey: sig.ValidatorPubKey[:],
			Signature:       sig.Signature[:],
		}
	}

	chainSubmission.Signatures = signatures

	msgs := make([]bindings.XChainMsg, len(submission.Msgs))
	for i, msg := range submission.Msgs {
		msgs[i] = bindings.XChainMsg{
			SourceChainId: msg.SourceChainID,
			DestChainId:   msg.DestChainID,
			StreamOffset:  msg.StreamOffset,
			Sender:        common.BytesToAddress(msg.SourceMsgSender[:]),
			To:            common.BytesToAddress(msg.DestAddress[:]),
			Data:          msg.Data,
			GasLimit:      msg.DestGasLimit,
			TxHash:        msg.TxHash,
		}
	}

	chainSubmission.Msgs = msgs

	return chainSubmission
}
