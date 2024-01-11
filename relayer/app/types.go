package relayer

import (
	"context"

	"github.com/omni-network/omni/lib/xchain"
)

type streamUpdate struct {
	xchain.StreamID
	AggAttestation xchain.AggAttestation // Attestation for the xmsgs
	Msgs           []xchain.Msg          // msgs that increment the cursor
}

// DetectorCallback is the callback function signature that will be called with stream updates.
type DetectorCallback func(context.Context, []streamUpdate) error

// Detector detects Stream updates that are approved and not yet submitted.
type Detector interface {
	// InsertBlock inserts a new block into the detector.
	InsertBlock(block xchain.Block) error

	// InsertAggAttestation inserts an attestation into the detector.
	InsertAggAttestation(attestation xchain.AggAttestation) error

	// RegisterOutput registers an output function that will be called with stream updates.
	RegisterOutput(ctx context.Context, cb DetectorCallback)
}

type Creator interface {
	// CreateSubmissions creates one or more submissions from the given stream update.
	CreateSubmissions(ctx context.Context, streamUpdate streamUpdate) ([]xchain.Submission, error)
}

type Sender interface {
	// SendTransaction sends a submission to the destination chain by invoking "xsubmit" on portal contract.
	SendTransaction(ctx context.Context, submission xchain.Submission) error
}
