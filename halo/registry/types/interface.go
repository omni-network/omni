package types

import (
	"context"
)

type PortalRegistry interface {
	SupportedChain(ctx context.Context, chainID uint64) (bool, error)
}
