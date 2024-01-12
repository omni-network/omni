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
type DetectorCallback func(context.Context, []streamUpdate)

// Detector detects Stream updates that are approved and not yet submitted.
type Detector interface {
	// InsertBlock inserts a new block into the detector.
	InsertBlock(ctx context.Context, block xchain.Block)

	// InsertAggAttestation inserts an attestation into the detector.
	InsertAggAttestation(ctx context.Context, attestation xchain.AggAttestation)

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

// cursorFetcher fetches all supported portal cursors.
type cursorFetcher interface {
	Cursors(ctx context.Context) ([]xchain.StreamCursor, error)
}
