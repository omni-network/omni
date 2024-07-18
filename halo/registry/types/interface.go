package types

import (
	"context"

	"github.com/omni-network/omni/lib/xchain"
)

type PortalRegistry interface {
	// SupportedChain returns true if the chain is included in the registry.
	SupportedChain(ctx context.Context, chainID uint64) (bool, error)

	// ConfLevels returns the all confirmation levels supported by all chains.
	ConfLevels(ctx context.Context) (map[uint64][]xchain.ConfLevel, error)
}
