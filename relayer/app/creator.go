package relayer

import (
	"github.com/omni-network/omni/lib/xchain"
)

func CreateSubmissions(streamUpdate StreamUpdate) ([]xchain.Submission, error) {
	// todo(lazar): in future this will receive receipts as well
	tree, err := xchain.NewBlockTree(xchain.Block{
		BlockHeader: streamUpdate.AggAttestation.BlockHeader,
		Msgs:        streamUpdate.Msgs,
	})

	if err != nil {
		return nil, err
	}

	multi, err := tree.Proof(streamUpdate.AggAttestation.BlockHeader, streamUpdate.Msgs)
	if err != nil {
		return nil, err
	}

	// todo(lazar): in future add ability for this to be batched into multiple submissions
	submissions := []xchain.Submission{{
		AttestationRoot: tree.Root(),
		BlockHeader:     streamUpdate.AggAttestation.BlockHeader,
		Msgs:            streamUpdate.Msgs,
		Proof:           multi.Proof,
		ProofFlags:      multi.ProofFlags,
		Signatures:      streamUpdate.AggAttestation.Signatures,
	}}

	return submissions, nil
}
