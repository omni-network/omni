package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
)

func newRequestValidator() func(ctx context.Context, chainID uint64, req bindings.SolveRequest) (uint8, bool, error) {
	return func(context.Context, uint64, bindings.SolveRequest) (uint8, bool, error) {
		return 0, false, errors.New("not implemented")
	}
}
