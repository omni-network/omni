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

type Creator interface {
	// CreateSubmissions creates one or more submissions from the given stream update.
	// creates submissions by splitting xmsgs into batches if required and generating merkle proofs for each submission
	CreateSubmissions(ctx context.Context, streamUpdate streamUpdate) ([]xchain.Submission, error)
}

type Sender interface {
	// SendTransaction sends a submission to the destination chain by invoking "xsubmit" on portal contract.
	SendTransaction(ctx context.Context, submission xchain.Submission) error
}

//// cursorFetcher fetches all supported portal cursors.
//type cursorFetcher interface {
//	Cursors(ctx context.Context) ([]xchain.StreamCursor, error)
//}

type xChainClient interface {
	GetBlock(ctx context.Context, chainID uint64, height uint64) (xchain.Block, bool, error)
	GetSubmittedCursors(ctx context.Context, chainID uint64) ([]xchain.StreamCursor, error)
}
