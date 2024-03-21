package types

import "context"

// ValidatorProvider is the interface that provides the active set of validators at a given height.
type ValidatorProvider interface {
	ActiveSetByHeight(ctx context.Context, height uint64) (*ValidatorSetResponse, error)
}
