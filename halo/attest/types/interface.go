package types

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"

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

	// TrimBehind deletes all available votes that are behind the vote window minimums (map[chainID]minimum) since
	// they will never be committed. It returns the number that was deleted.
	TrimBehind(minsByChain map[uint64]uint64) int

	// UpdateValidators sets the latest validator set when passed to cometBFT.
	// This is used to calculate whether the voter is in-or-out of the validator set.
	UpdateValidators(valset []abci.ValidatorUpdate)
}

// VoterDeps abstracts the Voter's internal cosmosSDK dependencies; basically the attest keeper.
// They have a circular dependency.
type VoterDeps interface {
	// LatestAttestationHeight returns the latest approved attestation height for the given chain.
	LatestAttestationHeight(ctx context.Context, chainID uint64) (uint64, bool, error)
}

// ChainNameFunc is a function that returns the name of a chain given its ID.
type ChainNameFunc func(chainID uint64) string
