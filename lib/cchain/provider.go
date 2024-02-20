package cchain

import (
	"context"

	"github.com/omni-network/omni/lib/xchain"
)

// ProviderCallback is the callback function signature that will be called with each approved attestation per
// source chain block in strictly sequential order.
type ProviderCallback func(ctx context.Context, approved xchain.AggAttestation) error

// Provider abstracts connecting to the omni consensus chain and streaming approved
// aggregate attestations for each source chain block from a specific height.
//
// It provides exactly once-delivery guarantees for the callback function.
// It will exponentially backoff and retry forever while the callback function returns an error.
type Provider interface {
	// Subscribe registers a callback function that will be called with all approved aggregate
	// attestations (as they become available per source chain block) on the consensus chain from
	// the provided source chain ID and height (inclusive).
	//
	// Worker name is only used for metrics.
	Subscribe(ctx context.Context, sourceChainID uint64, sourceHeight uint64,
		workerName string, callback ProviderCallback)

	// ApprovedFrom returns the subsequent approved aggregate attestations for the provided source chain
	// and height (inclusive). It will return max 100 aggregate attestations per call.
	ApprovedFrom(ctx context.Context, sourceChainID uint64, sourceHeight uint64) ([]xchain.AggAttestation, error)
}
