package types

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// Voter abstracts the validator duty of v∂oting for all
// XBlocks for all source chains.
//
// It streams all finalized XBlocks from all source chains.
// It creates a Vote for each (a signature).
// It stores these to disk, setting their status as "available".
type Voter interface {
	// GetAvailable returns all "available" votes in the vote window.
	// This basically queries all "available" votes and filters by current vote window.
	GetAvailable() []*Vote

	// SetProposed updates the status of the provided votes to "proposed",
	// i.e., they were included by a proposer in a new proposed block.
	// All other existing "proposed" votes are reset to "available", i.e. they were
	// proposed previously by another proposer, but that block was never finalized/committed.
	SetProposed(headers []*BlockHeader) error

	// SetCommitted updates the status of the provided votes to "committed",
	// i.e., they were included in a finalized consensus block and is now part of the consensus chain.
	// All other existing "proposed" votes are reset to "available", i.e. we probably
	// missed the proposal step and only learnt of the finalized block post-fact.
	// All but the latest "confirmed" attestation for each source chain can be safely deleted from disk.
	SetCommitted(headers []*BlockHeader) error

	// LocalAddress returns the local validator's ethereum address.
	LocalAddress() common.Address

	// TrimBehind deletes all available votes that are behind the vote window edges (map[chainID]edge) since
	// they will never be committed. It returns the number that was deleted.
	TrimBehind(edgesByChain map[uint64]uint64) int
}

// VoterDeps abstracts the Voter's internal cosmosSDK dependencies; basically the attest keeper.
// They have a circular dependency.
type VoterDeps interface {
	// IsValidator returns true if the given address is a validator on the given chain at the time of calling.
	IsValidator(ctx context.Context, address common.Address) bool
	// WindowCompare returns wether the given attestation block header is behind (-1), or in (0), or ahead (1)
	// of the current vote window. The vote window is a configured number of blocks around the latest approved
	// attestation for the provided chain.
	WindowCompare(ctx context.Context, chainID uint64, height uint64) (int, error)
	// LatestAttestationHeight returns the latest approved attestation height for the given chain.
	LatestAttestationHeight(ctx context.Context, chainID uint64) (uint64, bool, error)
}

// ChainNameFunc is a function that returns the name of a chain given its ID.
type ChainNameFunc func(chainID uint64) string
