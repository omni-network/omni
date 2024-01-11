package relayer

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/xchain"
)

type detectorService struct {
	mu sync.Mutex
}

func (d *detectorService) InsertBlock(block xchain.Block) error {
	// Implement logic to insert a new block into the detector.
	return nil
}

func (d *detectorService) InsertAggAttestation(attestation xchain.AggAttestation) error {
	// Implement logic to insert an attestation into the detector.
	return nil
}

func (d *detectorService) RegisterOutput(ctx context.Context, cb detectorCallback) {

}
