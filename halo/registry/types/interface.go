package types

import (
	"context"

	"github.com/omni-network/omni/lib/xchain"
)

type PortalRegistry interface {
	// SupportedChain returns true if the chain is included in the registry.
	SupportedChain(ctx context.Context, chainID uint64) (bool, error)

	// ConfLevels returns all confirmation levels supported by all chains.
	ConfLevels(ctx context.Context) (map[uint64][]xchain.ConfLevel, error)
}

// ChainNameFunc returns the name of the chain.
type ChainNameFunc func(chainID uint64) string
