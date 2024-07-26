package relayer

import (
	"slices"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"
)

// CreateSubmissions splits the update into multiple submissions that are each small enough (wrt calldata and gas)
// to be submitted on-chain.
func CreateSubmissions(up StreamUpdate) ([]xchain.Submission, error) {
	// Sanity check on input, should only be for a single stream.
	for i, msg := range up.Msgs {
		if msg.SourceChainID != up.SourceChainID {
			return nil, errors.New("invalid msgs [BUG]")
		} else if i > 0 && msg.StreamOffset != up.Msgs[i-1].StreamOffset+1 {
			return nil, errors.New("msgs not sequential [BUG]")
		}
	}

	att := up.Attestation

	attRoot, err := xchain.AttestationRoot(att.AttestHeader, att.BlockHeader, att.MsgRoot)
	if err != nil {
		return nil, err
	}

	var resp []xchain.Submission //nolint:prealloc // Cannot predetermine size
	for _, msgs := range groupMsgsByCost(up.Msgs) {
		multi, err := up.MsgTree.Proof(msgs)
		if err != nil {
			return nil, err
		}

		resp = append(resp, xchain.Submission{
			AttestationRoot: attRoot,
			ValidatorSetID:  att.ValidatorSetID,
			AttHeader:       att.AttestHeader,
			BlockHeader:     att.BlockHeader,
			Msgs:            msgs,
			Proof:           multi.Proof,
			ProofFlags:      multi.ProofFlags,
			Signatures:      att.Signatures,
			DestChainID:     up.DestChainID,
		})
	}

	return resp, nil
}

// groupMsgsByCost split the messages into groups that are each small enough (wrt calldata and gas)
// to be submitted on-chain.
func groupMsgsByCost(msgs []xchain.Msg) [][]xchain.Msg {
	var resp [][]xchain.Msg

	var current []xchain.Msg
	for _, msg := range msgs {
		// If adding the msg to the current batch will cross max limit, then
		// complete the batch by adding it to the response and starting a new empty current batch.

		// Note that even though the naive gas model doesn't work for all chains,
		// it is good enough for this use-case; i.e., splitting xmsgs.
		if naiveSubmissionGas(append(slices.Clone(current), msg)) > subGasMax {
			resp = append(resp, current)
			current = nil
		}

		// Add the message to the current batch
		current = append(current, msg)
	}

	// Add current batch to response and return
	return append(resp, current)
}
