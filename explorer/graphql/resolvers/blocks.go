//go:generate mockgen -destination=./tests/blocks_mock.go -package=resolvers_tests -source=blocks.go
package resolvers

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type BlocksProvider interface {
	XBlock(SourceChainID uint64, Height uint64) (*XBlock, bool, error)
}

type BlocksResolver struct {
	BlocksProvider BlocksProvider
}

type BlockArgs struct {
	SourceChainID hexutil.Big
	Height        hexutil.Big
}

func (b *BlocksResolver) XBlock(_ context.Context, args BlockArgs) (*XBlock, error) {
	res, found, err := b.BlocksProvider.XBlock(args.SourceChainID.ToInt().Uint64(), args.Height.ToInt().Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch block")
	}
	if !found {
		return nil, errors.New("block not found")
	}

	return res, nil
}
