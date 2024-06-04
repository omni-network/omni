package xchain

import (
	"context"
)

// ProviderCallback is the callback function signature that will be called with every finalized.
type ProviderCallback func(context.Context, Block) error

// ProviderRequest is the request struct for fetching cross-chain blocks.
// When used in streaming context, the Height and Offset fields define the starting point (inclusive).
type ProviderRequest struct {
	ChainID   uint64    // Source chain ID to query for xblocks.
	Height    uint64    // Height to query (from inclusive).
	ConfLevel ConfLevel // Confirmation level to ensure (and populate in returned BlockHeader)
	Offset    uint64    // Cross-chain block offset to populate (from inclusive) in BlockHeader (if required).
}

// Provider abstracts fetching cross chain data from any supported chain.
// This is basically a cross-chain data client for all supported chains.
type Provider interface {
	// StreamAsync starts a goroutine that streams xblocks forever from the provided source chain and height (inclusive).
	//
	// It returns immediately. It only returns an error if the chainID in invalid.
	// This is the async version of StreamBlocks.
	// It retries forever (with backoff) on all fetch and callback errors.
	StreamAsync(ctx context.Context, req ProviderRequest, callback ProviderCallback) error

	// StreamBlocks is the synchronous fail-fast version of Subscribe. It streams
	// xblocks as they become available but returns on the first callback error.
	// This is useful for workers that need to reset on application errors.
	StreamBlocks(ctx context.Context, req ProviderRequest, callback ProviderCallback) error

	// GetBlock returns the block for the given chain and height, or false if not available (not finalized yet),
	// or an error. The XBlockOffset field is populated with the provided offset (if required).
	GetBlock(ctx context.Context, req ProviderRequest) (Block, bool, error)

	// GetSubmittedCursor returns the submitted cursor for the provided stream,
	// or false if not available, or an error.
	// Calls the destination chain portal InXStreamOffset method.
	// Note this is only supported for EVM chains, no the consensus chain.
	GetSubmittedCursor(ctx context.Context, stream StreamID) (SubmitCursor, bool, error)

	// GetEmittedCursor returns the emitted cursor for the provided stream,
	// or false if not available, or an error.
	// Calls the source chain portal OutXStreamOffset method.
	//
	// Note that the BlockOffset field is not populated for emit cursors, since it isn't stored on-chain
	// but tracked off-chain.
	GetEmittedCursor(ctx context.Context, ref EmitRef, stream StreamID) (EmitCursor, bool, error)
}

type EmitRef struct {
	Height    *uint64
	ConfLevel *ConfLevel
}

func (r EmitRef) Valid() bool {
	return r.Height != nil || r.ConfLevel != nil
}
