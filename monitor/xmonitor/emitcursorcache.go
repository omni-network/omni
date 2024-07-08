package xmonitor

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

// cacheTrimLag is the number of blocks after which cursors are evicted from the cache.
const cacheTrimLag = 10_000

// startEmitCursorCache subscribes the xprovider iot populate the emit cursor cache.
// It returns a cache that will be populated and trimmed asynchronously.
func startEmitCursorCache(
	ctx context.Context,
	network netconf.Network,
	xprov xchain.Provider,
	cprov cchain.Provider,
) (*emitCursorCache, error) {
	cache := newEmitCursorCache()

	for _, chain := range network.Chains {
		callback := func(ctx context.Context, block xchain.Block) error {
			// Ignore blocks that are not attested.
			if !block.ShouldAttest(chain.AttestInterval) {
				return nil
			}

			// Update the emit cursor cache for each stream for this height.
			for _, stream := range network.StreamsFrom(chain.ID) {
				ref := xchain.EmitRef{Height: &block.BlockHeight}
				emit, _, err := xprov.GetEmittedCursor(ctx, ref, stream)
				if err != nil {
					log.Warn(ctx, "Skipping populating emit cursor cache", err, "stream", network.StreamName(stream))
					continue
				}

				cache.set(block.BlockHeight, stream, emit)
				if block.BlockHeight > cacheTrimLag { // Only trim after cacheTrimLag blocks.
					cache.trim(block.BlockHeight-cacheTrimLag, stream)
				}
			}

			return nil
		}

		// Stream from latest finalized attestation height
		fromHeight := chain.DeployHeight
		att, ok, err := cprov.LatestAttestation(ctx, xchain.ChainVersion{ID: chain.ID, ConfLevel: xchain.ConfFinalized})
		if err != nil {
			return nil, errors.Wrap(err, "latest attestation", "chain", chain.Name)
		} else if ok {
			fromHeight = att.BlockHeight
		}

		req := xchain.ProviderRequest{
			ChainID:   chain.ID,
			Height:    fromHeight,
			ConfLevel: xchain.ConfLatest, // Stream latest height to ensure state is available for querying.
			Offset:    0,                 // No offset required for emit cursors.
		}

		log.Info(ctx, "Subscribing to xblock to populate emit cursor cache", "chain", chain.Name, "from_height", fromHeight)

		if err := xprov.StreamAsync(ctx, req, callback); err != nil {
			return nil, err
		}
	}

	return cache, nil
}

// newEmitCursorCache creates a new emit cursor cache.
func newEmitCursorCache() *emitCursorCache {
	return &emitCursorCache{
		cursors: make(map[uint64]map[xchain.StreamID]xchain.EmitCursor),
		heights: make(map[xchain.StreamID][]uint64),
	}
}

// emitCursorCache is a cache of the last 10k emit cursors for each stream.
// This is used to avoid querying chain state (emit cursor) for historical blocks
// as this requires archive nodes.
// Instead we cache the emit cursor of latest blocks, and query the cache for historical blocks
// while monitoring attested stream offsets.
type emitCursorCache struct {
	mu      sync.RWMutex
	cursors map[uint64]map[xchain.StreamID]xchain.EmitCursor
	heights map[xchain.StreamID][]uint64
}

// set adds a cursor to the cache for the given height and stream.
func (c *emitCursorCache) set(height uint64, stream xchain.StreamID, cursor xchain.EmitCursor) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Add the cursor to the cache.
	if _, ok := c.cursors[height]; !ok {
		c.cursors[height] = make(map[xchain.StreamID]xchain.EmitCursor)
	}
	c.cursors[height][stream] = cursor

	// Add the height to the list of heights for this stream.
	c.heights[stream] = append(c.heights[stream], height)
}

// Get returns the emit cursor for the given height and stream.
func (c *emitCursorCache) Get(height uint64, stream xchain.StreamID) (xchain.EmitCursor, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	streams, ok := c.cursors[height]
	if !ok {
		return xchain.EmitCursor{}, false
	}

	cursor, ok := streams[stream]

	return cursor, ok
}

// trim removes all cursors for the stream that are older or equaled to the provided height.
func (c *emitCursorCache) trim(height uint64, stream xchain.StreamID) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Remove all heights that are older than the cache trim lag.
	var trimAfter int
	for i, h := range c.heights[stream] {
		if h > height {
			break // All remaining heights are within the trim lag.
		}

		trimAfter = i
		delete(c.cursors[h], stream)
		if len(c.cursors[h]) == 0 {
			delete(c.cursors, h)
		}
	}

	// Remove the trimmed heights from the list.
	c.heights[stream] = c.heights[stream][trimAfter+1:]
}
