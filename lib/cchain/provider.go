package cchain

import (
	"context"

	"github.com/omni-network/omni/lib/xchain"
)

// ProviderCallback is the callback function signature that will be called with all approved attestation per
// source chain block.
type ProviderCallback func(ctx context.Context, att xchain.AggAttestation) error

// Provider abstracts connecting to the omni consensus chain and streaming approved
// aggregate attestations for each source chain block from a specific height.
//
// It provides exactly once-delivery guarantees for the callback function.
// It will exponentially backoff and retry forever while the callback function returns an error.
type Provider interface {
	// Subscribe registers a callback function that will be called with all approved aggregate
	// attestations (as they become available per source chain block) on the consensus chain from
	// the provided source chain ID and height (inclusive).
	Subscribe(ctx context.Context, sourceChainID uint64, sourceHeight uint64, callback ProviderCallback)
}
