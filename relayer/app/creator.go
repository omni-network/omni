package relayer

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"
)

var _ Creator = (*CreatorService)(nil)

type CreatorService struct{}

// CreateSubmissions todo.
func (CreatorService) CreateSubmissions(ctx context.Context, streamUpdate StreamUpdate) ([]xchain.Submission, error) {
	_ = ctx // todo(lazar): use context?

	// todo(lazar): in future this will receive receipts as well
	tree, err := xchain.NewBlockTree(xchain.Block{
		BlockHeader: streamUpdate.AggAttestation.BlockHeader,
		Msgs:        streamUpdate.Msgs,
	})

	if err != nil {
		return nil, errors.Wrap(err, "error creating block tree")
	}

	multi, err := tree.Proof(streamUpdate.AggAttestation.BlockHeader, streamUpdate.Msgs)
	if err != nil {
		return nil, errors.Wrap(err, "error getting proofs")
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
