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
		if msg.DestChainID != up.DestChainID || msg.SourceChainID != up.SourceChainID {
			return nil, errors.New("invalid msgs [BUG]")
		} else if i > 0 && msg.StreamOffset != up.Msgs[i-1].StreamOffset+1 {
			return nil, errors.New("msgs not sequential [BUG]")
		}
	}

	agg := up.AggAttestation

	var resp []xchain.Submission //nolint:prealloc // Cannot predetermine size
	for _, msgs := range groupMsgsByCost(up.Msgs) {
		multi, err := up.Tree.Proof(agg.BlockHeader, msgs)
		if err != nil {
			return nil, err
		}

		resp = append(resp, xchain.Submission{
			AttestationRoot: agg.BlockRoot,
			ValidatorSetID:  1, // TODO(corver): Use agg.ValidatorSetHash once bindings supports it.
			BlockHeader:     agg.BlockHeader,
			Msgs:            msgs,
			Proof:           multi.Proof,
			ProofFlags:      multi.ProofFlags,
			Signatures:      agg.Signatures,
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
