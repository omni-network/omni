package cchain

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

// ProviderCallback is the callback function signature that will be called with each approved attestation per
// source chain block in strictly sequential order.
type ProviderCallback func(ctx context.Context, approved xchain.Attestation) error

// Provider abstracts connecting to the omni consensus chain and streaming approved
// attestations for each source chain block from a specific height.
//
// It provides exactly once-delivery guarantees for the callback function.
// It will exponentially backoff and retry forever while the callback function returns an error.
type Provider interface {
	// Subscribe registers a callback function that will be called with all approved
	// attestations (as they become available per source chain block) on the consensus chain from
	// the provided source chain ID and height (inclusive).
	//
	// Worker name is only used for metrics.
	Subscribe(ctx context.Context, sourceChainID uint64, sourceHeight uint64,
		workerName string, callback ProviderCallback)

	// AttestationsFrom returns the subsequent approved attestations for the provided source chain
	// and height (inclusive). It will return max 100 attestations per call.
	AttestationsFrom(ctx context.Context, sourceChainID uint64, sourceHeight uint64) ([]xchain.Attestation, error)

	// LatestAttestation returns the latest approved attestation for the provided source chain or false
	// if none exist.
	LatestAttestation(ctx context.Context, sourceChainID uint64) (xchain.Attestation, bool, error)

	// WindowCompare returns whether the given attestation block header is behind (-1), or in (0), or ahead (1)
	// of the current vote window. The vote window is a configured number of blocks around the latest approved
	// attestation for the provided chain.
	WindowCompare(ctx context.Context, sourceChainID uint64, height uint64) (int, error)

	// ValidatorSet returns the validators for the given validator set ID or false if none exist or an error.
	// Note the genesis validator set has ID 1.
	ValidatorSet(ctx context.Context, valSetID uint64) ([]Validator, bool, error)

	// XBlock returns the validator sync xblock for the given height (or latest) or false if none exist or an error.
	XBlock(ctx context.Context, height uint64, latest bool) (xchain.Block, bool, error)
}

// Validator is a consensus chain validator in a validator set.
type Validator struct {
	Address common.Address
	Power   int64
}

// Verify returns an error if the validator is invalid.
func (v Validator) Verify() error {
	if v.Address == (common.Address{}) {
		return errors.New("empty validator address")
	}
	if v.Power <= 0 {
		return errors.New("invalid validator power")
	}

	return nil
}
