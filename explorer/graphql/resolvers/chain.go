package resolvers

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
)

func (b *BlocksResolver) SupportedChains(ctx context.Context) ([]*Chain, error) {
	res, found, err := b.BlocksProvider.SupportedChains(ctx)
	if err != nil {
		return nil, errors.New("failed to fetch message count")
	}
	if !found {
		return nil, errors.New("message count not found")
	}

	return res, nil
}
