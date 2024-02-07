package relayer

import (
	"context"
	"sort"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/xchain"
)

type StreamUpdate struct {
	xchain.StreamID
	Tree           xchain.BlockTree
	AggAttestation xchain.AggAttestation // Attestation for the xmsgs
	Msgs           []xchain.Msg          // msgs that increment the cursor
}

// CreateFunc is a function that creates one or more submissions from the given stream update.
type CreateFunc func(streamUpdate StreamUpdate) ([]xchain.Submission, error)

// SendFunc sends a submission to the destination chain by invoking "xsubmit" on portal contract.
type SendFunc func(ctx context.Context, submission xchain.Submission) error

// SubmissionToBinding converts a go xchain submission to a solidity binding submission.
func SubmissionToBinding(sub xchain.Submission) bindings.XTypesSubmission {
	// Sort the signatures by validator address to ensure deterministic ordering.
	sort.Slice(sub.Signatures, func(i, j int) bool {
		return sub.Signatures[i].ValidatorAddress.Cmp(sub.Signatures[j].ValidatorAddress) < 0
	})

	sigs := make([]bindings.ValidatorSigTuple, 0, len(sub.Signatures))
	for _, sig := range sub.Signatures {
		sigs = append(sigs, bindings.ValidatorSigTuple{
			ValidatorAddr: sig.ValidatorAddress,
			Signature:     sig.Signature[:],
		})
	}

	msgs := make([]bindings.XTypesMsg, 0, len(sub.Msgs))
	for _, msg := range sub.Msgs {
		msgs = append(msgs, bindings.XTypesMsg{
			SourceChainId: msg.SourceChainID,
			DestChainId:   msg.DestChainID,
			StreamOffset:  msg.StreamOffset,
			Sender:        msg.SourceMsgSender,
			To:            msg.DestAddress,
			Data:          msg.Data,
			GasLimit:      msg.DestGasLimit,
		})
	}

	return bindings.XTypesSubmission{
		AttestationRoot: sub.AttestationRoot,
		// ValidatorSetID:  sub.ValidatorSetID, // TODO(corver): Uncomment when bindings are updated.
		BlockHeader: bindings.XTypesBlockHeader{
			SourceChainId: sub.BlockHeader.SourceChainID,
			BlockHeight:   sub.BlockHeader.BlockHeight,
			BlockHash:     sub.BlockHeader.BlockHash,
		},
		Proof:      sub.Proof,
		ProofFlags: sub.ProofFlags,
		Signatures: sigs,
		Msgs:       msgs,
	}
}
