package xmonitor

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

const (
	// cacheTrimLag is the number of blocks after which cursors are evicted from the cache.
	cacheTrimLag = 10_000
	// cacheStartLag is the number of blocks behind latest to start streaming and populating the cache.
	// 128 is the default number of historical block state that geth stores in non-archive mode.
	cacheStartLag = 128
)

// startEmitCursorCache subscribes the xprovider iot populate the emit cursor cache.
// It returns a cache that will be populated and trimmed asynchronously.
func startEmitCursorCache(
	ctx context.Context,
	network netconf.Network,
	xprov xchain.Provider,
) (*emitCursorCache, error) {
	cache := newEmitCursorCache()

	for _, chain := range network.Chains {
		callback := func(ctx context.Context, block xchain.Block) error {
			// Ignore blocks that are not attested.
			if !block.ShouldAttest(chain.AttestInterval) {
				return nil
			}

			latest, err := xprov.ChainVersionHeight(ctx, xchain.ChainVersion{ID: chain.ID, ConfLevel: xchain.ConfLatest})
			if err != nil {
				return err
			}

			// Update the emit cursor cache for each stream for this height.
			for _, stream := range network.StreamsFrom(chain.ID) {
				ref := xchain.EmitRef{Height: &block.BlockHeight}
				emit, _, err := xprov.GetEmittedCursor(ctx, ref, stream)
				if err != nil {
					log.Warn(ctx, "Skipping populating emit cursor cache", err,
						"stream", network.StreamName(stream),
						"lagging", subtract(latest, block.BlockHeight),
					)

					continue
				}

				cache.set(block.BlockHeight, stream, emit)
				if block.BlockHeight > cacheTrimLag { // Only trim after cacheTrimLag blocks.
					cache.trim(block.BlockHeight-cacheTrimLag, stream)
				}
			}

			return nil
		}

		// Figure out where to start streaming from.
		latest, err := xprov.ChainVersionHeight(ctx, xchain.ChainVersion{ID: chain.ID, ConfLevel: xchain.ConfLatest})
		if err != nil {
			return nil, errors.Wrap(err, "latest height", "chain", chain.Name)
		}

		var fromHeight uint64
		if latest > cacheStartLag {
			fromHeight = latest - cacheStartLag
		}
		req := xchain.ProviderRequest{
			ChainID:   chain.ID,
			Height:    fromHeight,
			ConfLevel: xchain.ConfLatest, // Stream latest height to ensure state is available for querying.
			Offset:    0,                 // No offset required for emit cursors.
		}

		log.Info(ctx, "Subscribing to xblocks to populate emit cursor cache", "chain", chain.Name, "from_height", fromHeight)

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

// AtOrBefore returns the stream emit cursor at-or-before the given height.
// Only attested heights are populated, so the first cursor at-or-before any
// height will return the correct cursor.
// TODO(corver): Maybe rather do binary search of heights.
func (c *emitCursorCache) AtOrBefore(height uint64, stream xchain.StreamID) (xchain.EmitCursor, bool) {
	for {
		c.mu.RLock()
		cursor, ok := c.cursors[height][stream]
		c.mu.RUnlock()

		if ok {
			// Return if we have a cursor for the current height
			return cursor, true
		}

		// Get earliest and latest heights for stream.
		c.mu.RLock()
		var earliest, latest uint64
		if heights := c.heights[stream]; len(heights) > 0 {
			earliest = heights[0]
			latest = heights[len(heights)-1]
		}
		c.mu.RUnlock()

		if height <= earliest {
			// We reached earliest height
			return xchain.EmitCursor{}, false
		} else if height > latest {
			height = latest // Jump directly to latest
		} else {
			height-- // Try one less
		}
	}
}

// trim removes all cursors for the stream that are older or equaled to the provided height.
func (c *emitCursorCache) trim(height uint64, stream xchain.StreamID) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Remove all heights that are older than the cache trim lag.
	trimAfter := -1
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
	if trimAfter >= 0 {
		c.heights[stream] = c.heights[stream][trimAfter+1:]
	}
}

func subtract(a uint64, b uint64) int64 {
	return int64(a) - int64(b)
}
