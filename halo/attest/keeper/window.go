package keeper

import (
	"context"
	"sync"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
)

type latestAttFunc func(context.Context, uint64) (*types.Attestation, bool, error)

type windower struct {
	// Immutable fields
	mu            sync.RWMutex
	latestAttFunc latestAttFunc
	voteWindow    uint64

	// Mutable field
	cache map[uint64]*types.Attestation
}

func newWindower(voteWindow uint64, latestAttFunc latestAttFunc) *windower {
	return &windower{
		latestAttFunc: latestAttFunc,
		voteWindow:    voteWindow,
		cache:         make(map[uint64]*types.Attestation),
	}
}

func (w *windower) ResetCache() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.cache = make(map[uint64]*types.Attestation)
}

// Compare returns whether the header is before (-1), or in (0), or after (1) the vote window.
func (w *windower) Compare(ctx context.Context, header *types.BlockHeader) (int, error) {
	w.mu.RLock()
	defer w.mu.RUnlock()

	chainID := header.ChainId

	latest, ok := w.cache[chainID]
	if !ok {
		att, exists, err := w.latestAttFunc(ctx, chainID)
		if err != nil {
			return 0, errors.Wrap(err, "latest attestation")
		} else if exists {
			latest = att
			w.cache[chainID] = att
		}
	}
	if latest == nil {
		// TODO(corver): Pass in netconf deploy height to use as initial window.
		return 0, nil // Allow any height while no approved attestation exists.
	}

	x := header.Height
	mid := latest.BlockHeader.Height
	delta := w.voteWindow

	if x < uintSub(mid, delta) {
		return -1, nil
	} else if x > mid+delta {
		return 1, nil
	}

	return 0, nil
}

// uintSub returns a - b if a > b, else 0.
// Subtracting uints can result in underflow, so we need to check for that.
func uintSub(a, b uint64) uint64 {
	if a <= b {
		return 0
	}

	return a - b
}
