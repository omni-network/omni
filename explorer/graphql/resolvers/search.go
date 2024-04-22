package resolvers

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
)

type SearchArgs struct {
	Query string
}

func (b *BlocksResolver) Search(ctx context.Context, args SearchArgs) (*SearchResult, error) {
	res, found, err := b.BlocksProvider.Search(ctx, args.Query)
	if err != nil {
		return nil, errors.New("failed to resolve search")
	}
	if !found {
		return nil, errors.New("search failed")
	}

	return res, nil
}
