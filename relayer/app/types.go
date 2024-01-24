package relayer

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

type StreamUpdate struct {
	xchain.StreamID
	AggAttestation xchain.AggAttestation // Attestation for the xmsgs
	Msgs           []xchain.Msg          // msgs that increment the cursor
}

// CreateFunc is a function that creates one or more submissions from the given stream update.
type CreateFunc func(streamUpdate StreamUpdate) ([]xchain.Submission, error)

type Sender interface {
	// SendTransaction sends a submission to the destination chain by invoking "xsubmit" on portal contract.
	SendTransaction(ctx context.Context, submission xchain.Submission) error
}

func TranslateSubmission(submission xchain.Submission) bindings.XChainSubmission {
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

	chainSubmission.Signatures = make([]bindings.XChainSigTuple, 0, len(submission.Signatures))
	for _, sig := range submission.Signatures {
		validatorPubKey := make([]byte, len(sig.ValidatorPubKey))
		copy(validatorPubKey, sig.ValidatorPubKey[:])
		signature := make([]byte, len(sig.Signature))
		copy(signature, sig.Signature[:])
		chainSubmission.Signatures = append(chainSubmission.Signatures, bindings.XChainSigTuple{
			ValidatorPubKey: validatorPubKey,
			Signature:       signature,
		})
	}

	msgs := make([]bindings.XChainMsg, 0, len(submission.Msgs))
	for _, msg := range submission.Msgs {
		msgs = append(msgs, bindings.XChainMsg{
			SourceChainId: msg.SourceChainID,
			DestChainId:   msg.DestChainID,
			StreamOffset:  msg.StreamOffset,
			Sender:        common.BytesToAddress(msg.SourceMsgSender[:]),
			To:            common.BytesToAddress(msg.DestAddress[:]),
			Data:          msg.Data,
			GasLimit:      msg.DestGasLimit,
		})
	}

	chainSubmission.Msgs = msgs

	return chainSubmission
}
