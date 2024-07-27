package types

import (
	"context"

	"github.com/omni-network/omni/lib/xchain"
)

type PortalRegistry interface {
	// ConfLevels returns all confirmation levels supported by all chains.
	// TODO(corver): Rename this to SupportedChainVersions
	ConfLevels(ctx context.Context) (map[uint64][]xchain.ConfLevel, error)
}

// ChainNameFunc returns the name of the chain.
type ChainNameFunc func(chainID uint64) string
