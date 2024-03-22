package relayer

import (
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

	var resp []xchain.Submission //nolint:prealloc // Cannot predetermine size
	for _, msgs := range groupMsgsByCost(up.Msgs) {
		multi, err := up.Tree.Proof(att.BlockHeader, msgs)
		if err != nil {
			return nil, err
		}

		resp = append(resp, xchain.Submission{
			AttestationRoot: att.AttestationRoot,
			ValidatorSetID:  att.ValidatorSetID,
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
//
// TODO(corver): For now, we only create a single group per update.
func groupMsgsByCost(msgs []xchain.Msg) [][]xchain.Msg {
	return [][]xchain.Msg{msgs}
}
