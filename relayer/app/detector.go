package relayer

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/xchain"
)

var _ Detector = (*detectorService)(nil)

type callBacker struct {
	callback DetectorCallback
	ctx      context.Context
}

type detectorService struct {
	mu               sync.Mutex
	submittedCursors map[xchain.StreamID]xchain.StreamCursor
	blocks           map[xchain.BlockHeader]xchain.Block
	aggAttestation   map[xchain.BlockHeader]xchain.AggAttestation
	callback         callBacker
}

func NewDetector(submittedCursors []xchain.StreamCursor) Detector {
	return &detectorService{
		submittedCursors: cursorsToMap(submittedCursors),
		blocks:           make(map[xchain.BlockHeader]xchain.Block),
		aggAttestation:   make(map[xchain.BlockHeader]xchain.AggAttestation),
	}
}

func (d *detectorService) InsertBlock(block xchain.Block) {
	d.mu.Lock()
	d.blocks[block.BlockHeader] = block
	d.mu.Unlock()
	d.process()
}

func (d *detectorService) InsertAggAttestation(attestation xchain.AggAttestation) {
	d.mu.Lock()
	d.aggAttestation[attestation.BlockHeader] = attestation
	d.mu.Unlock()
	d.process()
}

func (d *detectorService) RegisterOutput(ctx context.Context, cb DetectorCallback) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.callback = callBacker{
		callback: cb,
		ctx:      ctx,
	}
}

func (d *detectorService) process() {
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
	d.callback.callback(d.callback.ctx, stUpdates)
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
