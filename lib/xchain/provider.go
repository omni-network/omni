package xchain

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ProviderCallback is the callback function signature that will be called with every finalized.
type ProviderCallback func(context.Context, Block) error

// ProviderRequest is the request struct for fetching cross-chain blocks.
// When used in streaming context, the Height defines the starting point (inclusive).
type ProviderRequest struct {
	ChainID   uint64    // Source chain ID to query for xblocks.
	Height    uint64    // Height to query (from inclusive).
	ConfLevel ConfLevel // Confirmation level to ensure
}

func (r ProviderRequest) ChainVersion() ChainVersion {
	return ChainVersion{
		ID:        r.ChainID,
		ConfLevel: r.ConfLevel,
	}
}

// EventLogsReq is the request to fetch EVM event logs.
type EventLogsReq struct {
	ChainID         uint64           // Source chain ID to query for xblocks.
	Height          uint64           // Height to query (from inclusive).
	ConfLevel       ConfLevel        // Confirmation level to ensure
	FilterAddresses []common.Address // Filter by zero or more addresses.
	FilterTopics    []common.Hash    // Filters zero or more topics (in the first position only).
}

func (r EventLogsReq) ChainVersion() ChainVersion {
	return ChainVersion{
		ID:        r.ChainID,
		ConfLevel: r.ConfLevel,
	}
}

type EventLogsCallback func(ctx context.Context, header *types.Header, events []types.Log) error

// Provider abstracts fetching cross chain data from any supported chain.
// This is basically a cross-chain data client for all supported chains.
type Provider interface {
	// StreamAsync starts a goroutine that streams xblocks forever from the provided source chain and height (inclusive).
	//
	// It returns immediately. It only returns an error if the chainID in invalid.
	// This is the async version of StreamBlocks.
	// It retries forever (with backoff) on all fetch and callback errors.
	StreamAsync(ctx context.Context, req ProviderRequest, callback ProviderCallback) error

	// StreamBlocks is the synchronous fail-fast version of StreamBlocks. It streams
	// xblocks as they become available but returns on the first callback error.
	// This is useful for workers that need to reset on application errors.
	StreamBlocks(ctx context.Context, req ProviderRequest, callback ProviderCallback) error

	// StreamEventLogs streams EVM event logs as they become available.
	//
	// The callback will be called with strictly-sequential heights with logs matching the provided filter (which may be none).
	// It returns any error encountered.
	StreamEventLogs(ctx context.Context, req EventLogsReq, callback EventLogsCallback) error

	// GetBlock returns the block for the given chain and height, or false if not available (not finalized yet),
	// or an error. The AttestOffset field is populated with the provided offset (if required).
	GetBlock(ctx context.Context, req ProviderRequest) (Block, bool, error)

	// GetSubmittedCursor returns the submitted cursor for the provided stream,
	// or false if not available, or an error.
	// Calls the destination chain portal InXStreamOffset method.
	// Note this is only supported for EVM chains, no the consensus chain.
	GetSubmittedCursor(ctx context.Context, ref Ref, stream StreamID) (SubmitCursor, bool, error)

	// GetEmittedCursor returns the emitted cursor for the provided stream,
	// or false if not available, or an error.
	// Calls the source chain portal OutXStreamOffset method.
	//
	// Note that the AttestOffset field is not populated for emit cursors, since it isn't stored on-chain
	// but tracked off-chain.
	GetEmittedCursor(ctx context.Context, ref Ref, stream StreamID) (EmitCursor, bool, error)

	// ChainVersionHeight returns the height for the provided chain version.
	ChainVersionHeight(ctx context.Context, chainVer ChainVersion) (uint64, error)

	// GetSubmission returns the submission for the provided chain and tx hash, or an error.
	GetSubmission(ctx context.Context, chainID uint64, txHash common.Hash) (Submission, error)
}

// Ref specifies which block to query for cursors.
type Ref struct {
	// Height specifies an absolute height to query; if non-nil.
	Height *uint64
	// ConfLevel specifies a relative-to-head block to query; if non-nil.
	ConfLevel *ConfLevel
}

func (r Ref) Valid() bool {
	return r.Height != nil || r.ConfLevel != nil
}

// ConfRef returns a Ref with the provided confirmation level.
func ConfRef(level ConfLevel) Ref {
	return Ref{
		ConfLevel: &level,
	}
}

// HeightRef returns a Ref with the provided confirmation level.
func HeightRef(height uint64) Ref {
	return Ref{
		Height: &height,
	}
}

var (
	// LatestRef references the latest confirmation level.
	LatestRef = ConfRef(ConfLatest)
	// FinalizedRef references the latest confirmation level.
	FinalizedRef = ConfRef(ConfFinalized)
)
