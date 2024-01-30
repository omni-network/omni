package resolvers

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
)

func (m XMsg) SourceMessageSender() string {
	return m.SourceMessageSenderRaw.String()
}
func (m XMsg) DestAddress() string {
	return m.DestAddressRaw.String()
}

func (m XMsg) TxHash() string {
	return m.TxHashRaw.String()
}

func (b *XBlock) BlockHash() string {
	return b.BlockHashRaw.String()
}

func (b *XBlock) SourceChainID() BigInt {
	return b.SourceChainIDRaw
}

func (b *XBlock) BlockHeight() BigInt {
	return b.BlockHeightRaw
}

type BlocksProvider interface {
	XBlock(SourceChainID uint64, Height uint64) (*XBlock, bool, error)
}

type BlocksResolver struct {
	BlocksProvider BlocksProvider
}

type BlockArgs struct {
	SourceChainID BigInt
	Height        BigInt
}

func (b *BlocksResolver) XBlock(_ context.Context, args BlockArgs) (*XBlock, error) {
	res, found, err := b.BlocksProvider.XBlock(args.SourceChainID.Int.Uint64(), args.Height.Int.Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch block")
	}
	if !found {
		return nil, errors.New("block not found")
	}

	return res, nil
}
