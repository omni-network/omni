package types

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// ValidatorProvider is the interface that provides the active set of validators at a given height.
type ValidatorProvider interface {
	ActiveSetByHeight(ctx context.Context, height uint64) (*ValidatorSetResponse, error)
}

type RegisterPortalRequest struct {
	ChainID      uint64
	Address      common.Address
	DeployHeight uint64
	ShardIDs     []uint64
}
