package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

// Definition defines the solver rules and addresses and config per network.
type Definition struct {
	InboxAddress       common.Address
	OutboxAddress      common.Address
	InboxDeployHeights map[uint64]uint64
}

//nolint:unused // False positive.
func newRequestValidator(_ Definition) func(ctx context.Context, chainID uint64, req bindings.SolveRequest) (uint8, bool, error) {
	return func(context.Context, uint64, bindings.SolveRequest) (uint8, bool, error) {
		return 0, false, errors.New("not implemented")
	}
}
