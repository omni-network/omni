package resolvers

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/graph-gophers/graphql-go"
)

type StreamID struct {
	SourceChainID BigInt
	DestChainID   BigInt
}

type MsgID struct {
	StreamID     StreamID
	StreamOffset BigInt
}

type Msg struct {
	MsgID
	SenderRaw      common.Address
	DestAddressRaw common.Address
	DestGasLimit   BigInt
	TxHashRaw      common.Hash
}

func (m Msg) Sender() string {
	return m.SenderRaw.String()
}
func (m Msg) DestAddress() string {
	return m.DestAddressRaw.String()
}

func (m Msg) TxHash() string {
	return m.TxHashRaw.String()
}

type Block struct {
	SourceChainID BigInt
	BlockHeight   BigInt
	Hash          common.Hash
	Timestamp     graphql.Time

	// TODO(Pavel): add paging for the messages.
	Messages []Msg
}

func (b *Block) BlockHash() string {
	return b.Hash.String()
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

func (b *BlocksResolver) Block(ctx context.Context, args BlockArgs) (*Block, error) {
	res, found, err := b.BlocksProvider.Block(args.SourceChainID.Int.Uint64(), args.Height.Int.Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch block")
	}
	if !found {
		return nil, errors.New("block not found")
	}
	return res, nil
}
