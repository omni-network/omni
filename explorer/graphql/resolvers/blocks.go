//go:generate mockgen -destination=./tests/blocks_mock.go -package=resolvers_test -source=blocks.go
package resolvers

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type BlocksProvider interface {
	XBlock(ctx context.Context, SourceChainID uint64, Height uint64) (*XBlock, bool, error)
	XBlockRange(ctx context.Context, Amount uint64, Offset uint64) ([]*XBlock, bool, error)
	XBlockCount(ctx context.Context) (*hexutil.Big, bool, error)
	XMsgCount(ctx context.Context) (*hexutil.Big, bool, error)
	XReceiptCount(ctx context.Context) (*hexutil.Big, bool, error)
}

type BlocksResolver struct {
	BlocksProvider BlocksProvider
}

type XBlockArgs struct {
	SourceChainID hexutil.Big
	Height        hexutil.Big
}

type XBlockRangeArgs struct {
	Amount hexutil.Big
	Offset hexutil.Big
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
	res, found, err := b.BlocksProvider.XBlockRange(ctx, args.Amount.ToInt().Uint64(), args.Offset.ToInt().Uint64())
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

func (b *BlocksResolver) XMsgCount(ctx context.Context) (*hexutil.Big, error) {
	res, found, err := b.BlocksProvider.XMsgCount(ctx)
	if err != nil {
		return nil, errors.New("failed to fetch message count")
	}
	if !found {
		return nil, errors.New("message count not found")
	}

	return res, nil
}

func (b *BlocksResolver) XReceiptCount(ctx context.Context) (*hexutil.Big, error) {
	res, found, err := b.BlocksProvider.XReceiptCount(ctx)
	if err != nil {
		return nil, errors.New("failed to fetch receipt count")
	}
	if !found {
		return nil, errors.New("receipt count not found")
	}

	return res, nil
}
