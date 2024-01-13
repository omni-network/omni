package xchain

import "context"

// ProviderCallback is the callback function signature that will be called with every finalized.
type ProviderCallback func(context.Context, *Block) error

// Provider abstracts connecting to any supported source chain and streaming finalized
// XBlocks from a specific height. It provides exactly once-delivery guarantees for the callback function.
// It will exponentially backoff and retry forever while the callback function returns an error.
type Provider interface {
	// Subscribe registers a callback function that will be called with every finalized
	// XBlock (as they become available) on the source chain from the provided height (inclusive).
	Subscribe(
		ctx context.Context,
		chainID uint64,
		fromHeight uint64,
		callback ProviderCallback,
	) error
}
