package types

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// Voter abstracts the validator duty of voting for all
// XBlocks for all source chains.
//
// It streams all finalized XBlocks from all source chains.
// It creates an Vote for each (a signature).
// It stores these to disk, setting their status as "available".
type Voter interface {
	// GetAvailable returns all "available" votes.
	// This basically queries all "available" votes.
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
}

// Windower abstract the logic to verify whether a vote's block header is
// in the voting window.
//
// The vote window is a number of configured blocks around the latest
// approved attestation for the provided chain.
type Windower interface {
	// Compare return whether the header is before (-1), in (0) or after (1) the vote window.
	Compare(ctx context.Context, header *BlockHeader) (int, error)
	// ResetCache resets the internal cache.
	ResetCache()
}

// ChainNameFunc is a function that returns the name of a chain given its ID.
type ChainNameFunc func(chainID uint64) string
