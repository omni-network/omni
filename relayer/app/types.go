package relayer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/xchain"
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

type XChainClient interface {
	GetBlock(ctx context.Context, chainID uint64, height uint64) (xchain.Block, bool, error)
	GetSubmittedCursors(ctx context.Context, chainID uint64) ([]xchain.StreamCursor, error)
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
