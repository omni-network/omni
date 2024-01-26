package resolvers

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
)

func (m Msg) SourceMessageSender() string {
	return m.SourceMessageSenderRaw.String()
}
func (m Msg) DestAddress() string {
	return m.DestAddressRaw.String()
}

func (m Msg) TxHash() string {
	return m.TxHashRaw.String()
}

func (b *Block) BlockHash() string {
	return b.BlockHashRaw.String()
}

type BlocksProvider interface {
	Block(SourceChainID uint64, Height uint64) (*Block, bool, error)
}

type BlocksResolver struct {
	BlocksProvider BlocksProvider
}

type BlockArgs struct {
	SourceChainID BigInt
	Height        BigInt
}

func (b *BlocksResolver) Block(_ context.Context, args BlockArgs) (*Block, error) {
	res, found, err := b.BlocksProvider.Block(args.SourceChainID.Int.Uint64(), args.Height.Int.Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch block")
	}
	if !found {
		return nil, errors.New("block not found")
	}

	return res, nil
}
