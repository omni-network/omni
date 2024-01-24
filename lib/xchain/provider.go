package xchain

import "context"

// ProviderCallback is the callback function signature that will be called with every finalized.
type ProviderCallback func(context.Context, *Block) error

// Provider abstracts fetching cross chain data from any supported chain.
// This is basically a cross-chain data client for all supported chains.
type Provider interface {
	Subscribe(ctx context.Context, chainID uint64, fromHeight uint64, callback ProviderCallback) error
	GetBlock(ctx context.Context, chainID uint64, height uint64) (Block, bool, error)
	GetSubmittedCursor(ctx context.Context, chainID uint64, sourceChainID uint64) (StreamCursor, error)
}
