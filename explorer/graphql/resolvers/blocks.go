//go:generate mockgen -destination=./tests/blocks_mock.go -package=resolvers_test -source=blocks.go
package resolvers

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type XBlockArgs struct {
	SourceChainID hexutil.Big
	Height        hexutil.Big
}

type XBlockRangeArgs struct {
	From hexutil.Big
	To   hexutil.Big
}

func (b *BlocksResolver) XBlock(ctx context.Context, args XBlockArgs) (*XBlock, error) {
	res, found, err := b.BlocksProvider.XBlock(ctx, args.SourceChainID.ToInt().Uint64(), args.Height.ToInt().Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch block")
	}
	if !found {
		return nil, errors.New("block not found")
	}

	return res, nil
}

func (b *BlocksResolver) XBlockRange(ctx context.Context, args XBlockRangeArgs) ([]*XBlock, error) {
	if args.From.ToInt().Cmp(args.To.ToInt()) >= 0 {
		return nil, errors.New("invalid range")
	}
	res, found, err := b.BlocksProvider.XBlockRange(ctx, args.From.ToInt().Uint64(), args.To.ToInt().Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch block range")
	}
	if !found {
		return nil, errors.New("block range not found")
	}

	return res, nil
}

func (b *BlocksResolver) XBlockCount(ctx context.Context) (*hexutil.Big, error) {
	res, found, err := b.BlocksProvider.XBlockCount(ctx)
	if err != nil {
		return nil, errors.New("failed to fetch block count")
	}
	if !found {
		return nil, errors.New("block count not found")
	}

	return res, nil
}
