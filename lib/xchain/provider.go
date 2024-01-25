package xchain

import "context"

// ProviderCallback is the callback function signature that will be called with every finalized.
type ProviderCallback func(context.Context, Block) error

// Provider abstracts fetching cross chain data from any supported chain.
// This is basically a cross-chain data client for all supported chains.
type Provider interface {
	// Subscribe registers a callback function that will be called with each XBlock
	// (as they become finalized per source chain) for the provided source chain ID and height (inclusive).
	Subscribe(ctx context.Context, chainID uint64, fromHeight uint64, callback ProviderCallback) error

	// GetBlock returns the block for the given chain and height, or false if not available (not finalized yet),
	// or an error.
	GetBlock(ctx context.Context, chainID uint64, height uint64) (Block, bool, error)

	// GetSubmittedCursor returns the submitted cursor for the provided chain and source chain,
	// or false if not available, or an error.
	GetSubmittedCursor(ctx context.Context, chainID uint64, sourceChainID uint64) (StreamCursor, bool, error)
}
