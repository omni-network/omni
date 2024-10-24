package provider

import (
	"sync"

	"github.com/omni-network/omni/lib/xchain"
)

// chainVersionCache contains approved attestations of a chain version with all heights above the watermark.
type chainVersionCache struct {
	cache     map[uint64]xchain.Attestation
	watermark uint64
}

// attestationCache if a thread-safe collection of caches containing attestations per chain version.
// All cache sizes will be kept under `cacheSize` elements.
type attestationCache struct {
	cacheSizes int
	caches     map[xchain.ChainVersion]*chainVersionCache
	mu         sync.RWMutex
}

// newAttestationCache constructs a default cache collection with size 100 per cache.
func newAttestationCache() *attestationCache {
	return &attestationCache{
		cacheSizes: 100,
		caches:     make(map[xchain.ChainVersion]*chainVersionCache),
	}
}

// get returns all cached consequent attestations, starting from the specified height.
func (ac *attestationCache) get(chainVer xchain.ChainVersion, height uint64) []xchain.Attestation {
	var attestations []xchain.Attestation

	ac.mu.RLock()
	defer ac.mu.RUnlock()
	chainVerCache := ac.caches[chainVer]

	if chainVerCache == nil || height < chainVerCache.watermark {
		return attestations
	}

	for {
		value, exists := chainVerCache.cache[height]
		if !exists {
			break
		}
		attestations = append(attestations, value)
		height++
	}

	return attestations
}

// update extends the cache of chain version with all new attestations that are above the stored watermark, and then
// removes the oldest ones to keep the total cache size at `cachesSize`. Note we do not cache attestations below the
// watermark because they are only expected to be occasionally fetched by retrying workers.
func (ac *attestationCache) update(chainVer xchain.ChainVersion, attestastions []xchain.Attestation) {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	// If no cache exist yet, initialize it.
	if _, present := ac.caches[chainVer]; !present {
		ac.caches[chainVer] = &chainVersionCache{map[uint64]xchain.Attestation{}, 0}
	}

	chainVerCache := ac.caches[chainVer]

	// Insert each attestation above the watermark into the cache.
	// Note, it could overwrite previous values, but this is idempotent.
	for _, att := range attestastions {
		if att.BlockHeader.BlockHeight < chainVerCache.watermark {
			continue
		}
		attHeight := att.BlockHeader.BlockHeight
		chainVerCache.cache[attHeight] = att
		chainVerCache.watermark = max(chainVerCache.watermark, attHeight)
	}

	// Raise the watermark until we hit the configured cache size and delete all
	// attestations below the watermark.
	for len(chainVerCache.cache) > ac.cacheSizes {
		delete(chainVerCache.cache, chainVerCache.watermark)
		chainVerCache.watermark++
	}
}
