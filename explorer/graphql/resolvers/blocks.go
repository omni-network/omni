//go:generate mockgen -destination=./tests/blocks_mock.go -package=resolvers_test -source=blocks.go
package resolvers

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type BlocksProvider interface {
	XBlock(SourceChainID uint64, Height uint64) (*XBlock, bool, error)
	XBlockRange(Amount uint64, Offset uint64) ([]*XBlock, bool, error)
	XBlockCount() (*hexutil.Big, bool, error)
	XMsgCount() (*hexutil.Big, bool, error)
	XReceiptCount() (*hexutil.Big, bool, error)
	// TODO: fill in the rest of the methods
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

func (b *BlocksResolver) XBlock(_ context.Context, args XBlockArgs) (*XBlock, error) {
	res, found, err := b.BlocksProvider.XBlock(args.SourceChainID.ToInt().Uint64(), args.Height.ToInt().Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch block")
	}
	if !found {
		return nil, errors.New("block not found")
	}

	return res, nil
}

func (b *BlocksResolver) XBlockRange(_ context.Context, args XBlockRangeArgs) ([]*XBlock, error) {
	res, found, err := b.BlocksProvider.XBlockRange(args.Amount.ToInt().Uint64(), args.Offset.ToInt().Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch block range")
	}
	if !found {
		return nil, errors.New("block range not found")
	}

	return res, nil
}

func (b *BlocksResolver) XBlockCount(_ context.Context) (*hexutil.Big, error) {
	res, found, err := b.BlocksProvider.XBlockCount()
	if err != nil {
		return nil, errors.New("failed to fetch block count")
	}
	if !found {
		return nil, errors.New("block count not found")
	}

	return res, nil
}

func (b *BlocksResolver) XMsgCount(_ context.Context) (*hexutil.Big, error) {
	res, found, err := b.BlocksProvider.XMsgCount()
	if err != nil {
		return nil, errors.New("failed to fetch message count")
	}
	if !found {
		return nil, errors.New("message count not found")
	}

	return res, nil
}

func (b *BlocksResolver) XReceiptCount(_ context.Context) (*hexutil.Big, error) {
	res, found, err := b.BlocksProvider.XReceiptCount()
	if err != nil {
		return nil, errors.New("failed to fetch receipt count")
	}
	if !found {
		return nil, errors.New("receipt count not found")
	}

	return res, nil
}
