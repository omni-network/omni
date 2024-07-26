package voter

import "github.com/omni-network/omni/lib/errors"

// offsetTracker tracks and assigns the AttestOffset.
type offsetTracker struct {
	nextAttestOffset uint64
	prevBlockHeight  uint64
}

// newOffsetTracker returns a new offset tracker, setting the next state to the provided values.
func newOffsetTracker(nextAttestOffset uint64) *offsetTracker {
	return &offsetTracker{
		nextAttestOffset: nextAttestOffset,
	}
}

// NextAttestOffset returns the next attestation offset ensuring the block height is increasing.
func (c *offsetTracker) NextAttestOffset(blockHeight uint64) (uint64, error) {
	if c.prevBlockHeight != 0 && c.prevBlockHeight >= blockHeight {
		return 0, errors.New("unexpected block height for attest offset [BUG]", "prev", c.prevBlockHeight, "new", blockHeight)
	}

	resp := c.nextAttestOffset
	c.nextAttestOffset++
	c.prevBlockHeight = blockHeight

	return resp, nil
}
