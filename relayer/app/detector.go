package relayer

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/xchain"
)

var _ Detector = (*detectorService)(nil)

type detectorService struct {
	mu               sync.Mutex
	submittedCursors map[xchain.StreamID]xchain.StreamCursor
	blocks           map[xchain.BlockHeader]xchain.Block
	aggAttestation   map[xchain.BlockHeader]xchain.AggAttestation
	callback         DetectorCallback
}

func NewDetector(submittedCursors []xchain.StreamCursor) Detector {
	return &detectorService{
		submittedCursors: cursorsToMap(submittedCursors),
		blocks:           make(map[xchain.BlockHeader]xchain.Block),
		aggAttestation:   make(map[xchain.BlockHeader]xchain.AggAttestation),
	}
}

func (d *detectorService) InsertBlock(ctx context.Context, block xchain.Block) {
	d.mu.Lock()
	d.blocks[block.BlockHeader] = block
	d.mu.Unlock()
	d.process(ctx)
}

func (d *detectorService) InsertAggAttestation(ctx context.Context, attestation xchain.AggAttestation) {
	d.mu.Lock()
	d.aggAttestation[attestation.BlockHeader] = attestation
	d.mu.Unlock()
	d.process(ctx)
}

func (d *detectorService) RegisterOutput(cb DetectorCallback) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.callback = cb
}

func (d *detectorService) process(ctx context.Context) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// iterate over attestations check if any exist in block map
	var streamUpdates map[xchain.StreamID]streamUpdate
	for _, attestation := range d.aggAttestation {
		// if exists, check if any msgs exist in block map
		block, ok := d.blocks[attestation.BlockHeader]
		if !ok {
			continue
		}

		// todo(lazar): check if xmsgs should be sorted by offset in the block or do they come sorted?

		for _, msg := range block.Msgs {
			cursor, ok := d.submittedCursors[msg.StreamID]
			if !ok {
				continue
			}
			if cursor.SourceBlockHeight > block.BlockHeader.BlockHeight {
				continue
			}
			// todo(lazar): handle offset

			stUp, ok := streamUpdates[msg.StreamID]
			if !ok {
				streamUpdates[msg.StreamID] = streamUpdate{
					StreamID:       msg.StreamID,
					AggAttestation: attestation,
				}
			}
			stUp.Msgs = append(stUp.Msgs, msg)
		}
	}
	stUpdates := streamUpdateToSlice(streamUpdates)
	d.callback(ctx, stUpdates)
	// todo(Lazar): append new streamcursor to submittedCursors
}

func streamUpdateToSlice(streamUpdates map[xchain.StreamID]streamUpdate) []streamUpdate {
	res := make([]streamUpdate, 0, len(streamUpdates))
	for _, v := range streamUpdates {
		res = append(res, v)
	}

	return res
}

func cursorsToMap(cursors []xchain.StreamCursor) map[xchain.StreamID]xchain.StreamCursor {
	res := make(map[xchain.StreamID]xchain.StreamCursor)
	for _, cursor := range cursors {
		res[cursor.StreamID] = cursor
	}

	return res
}
