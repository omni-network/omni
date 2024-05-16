package xchain

import (
	"context"

	"github.com/omni-network/omni/lib/ethclient"
)

// ProviderCallback is the callback function signature that will be called with every finalized.
type ProviderCallback func(context.Context, Block) error

// Provider abstracts fetching cross chain data from any supported chain.
// This is basically a cross-chain data client for all supported chains.
type Provider interface {
	// StreamAsync starts a goroutine that streams xblocks forever from the provided source chain and height (inclusive).
	// It returns immediately. It only returns an error if the chainID in invalid.
	// This is the async version of StreamBlocks.
	// It retries forever (with backoff) on all fetch and callback errors.
	StreamAsync(ctx context.Context, chainID uint64, fromHeight uint64, fromOffset uint64, callback ProviderCallback) error

	// StreamAsyncNoOffset is the same as StreamAsync except that XBlockOffset is not populated in the returned XBlocks.
	StreamAsyncNoOffset(ctx context.Context, chainID uint64, fromHeight uint64, callback ProviderCallback) error

	// StreamBlocks is the synchronous fail-fast version of Subscribe. It streams
	// xblocks as they become available but returns on the first callback error.
	// This is useful for workers that need to reset on application errors.
	StreamBlocks(ctx context.Context, chainID uint64, fromHeight uint64, fromOffset uint64, callback ProviderCallback) error

	// StreamBlocksNoOffset is the same as StreamBlocks except that XBlockOffset is not populated in the returned XBlocks.
	StreamBlocksNoOffset(ctx context.Context, chainID uint64, fromHeight uint64, callback ProviderCallback) error

	// GetBlock returns the block for the given chain and height, or false if not available (not finalized yet),
	// or an error. The XBlockOffset field is populated with the provided offset (if required).
	GetBlock(ctx context.Context, chainID uint64, height uint64, xOffset uint64) (Block, bool, error)

	// GetSubmittedCursor returns the submitted cursor for the source chain on the destination chain,
	// or false if not available, or an error. Calls the destination chain portal InXStreamOffset method.
	// Note this is only supported for EVM chains, no the consensus chain.
	GetSubmittedCursor(ctx context.Context, destChainID uint64, sourceChainID uint64) (StreamCursor, bool, error)

	// GetEmittedCursor returns the emitted cursor for the destination chain on the source chain,
	// or false if not available, or an error. Calls the source chain portal OutXStreamOffset method.
	//
	// Note that the BlockOffset field is not populated for emit cursors, since it isn't stored on-chain
	// but tracked off-chain.
	GetEmittedCursor(ctx context.Context, ref EmitRef, srcChainID uint64, destChainID uint64) (StreamCursor, bool, error)
}

type EmitRef struct {
	Height   *uint64
	HeadType *ethclient.HeadType
}
