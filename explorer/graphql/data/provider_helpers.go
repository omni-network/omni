package data

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
)

// EntBlockToGraphQLBlock converts an ent.Block to a resolvers.XBlock.
func EntBlockToGraphQLBlock(ctx context.Context, block *ent.Block) (*resolvers.XBlock, error) {
	sourceChainIDBig, err := Uint2Big(block.SourceChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	blockHeight, err := Uint2Big(block.BlockHeight)
	if err != nil {
		return nil, errors.Wrap(err, "decoding block height")
	}

	res := resolvers.XBlock{
		SourceChainID: hexutil.Big(sourceChainIDBig),
		BlockHeight:   hexutil.Big(blockHeight),
		BlockHash:     common.Hash(block.BlockHash),
		Timestamp:     graphql.Time{Time: block.CreatedAt},
	}

	// Decode messages
	for _, msg := range block.Edges.Msgs {
		msg, err := EntMsgToGraphQLXMsg(ctx, msg, block)
		if err != nil {
			return nil, errors.Wrap(err, "decoding msg for block")
		}
		res.Messages = append(res.Messages, *msg)
	}

	// Decode receipts
	for _, receipt := range block.Edges.Receipts {
		receipt, err := EntReceiptToGraphQLXReceipt(receipt)
		if err != nil {
			return nil, errors.Wrap(err, "decoding receipt for block")
		}
		res.Receipts = append(res.Receipts, *receipt)
	}

	return &res, nil
}

// EntMsgToGraphQLXMsg converts an ent.Msg to a resolvers.XMsg.
func EntMsgToGraphQLXMsg(ctx context.Context, msg *ent.Msg, block *ent.Block) (*resolvers.XMsg, error) {
	if block == nil {
		b, err := msg.QueryBlock().Only(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "querying block for message")
		}
		block = b
	}

	sourceChainIDBig, err := Uint2Big(msg.SourceChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destChainIDBig, err := Uint2Big(msg.DestChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destGasLimit, err := Uint2Big(msg.DestGasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "decoding dest gas limit")
	}

	streamOffset, err := Uint2Big(msg.StreamOffset)
	if err != nil {
		return nil, errors.Wrap(err, "decoding stream offset")
	}

	blockHeight, err := Uint2Big(block.BlockHeight)
	if err != nil {
		return nil, errors.Wrap(err, "decoding block height")
	}

	return &resolvers.XMsg{
		SourceMessageSender: common.Address(msg.SourceMsgSender),
		SourceChainID:       hexutil.Big(sourceChainIDBig),
		DestAddress:         common.Address(msg.DestAddress),
		DestGasLimit:        hexutil.Big(destGasLimit),
		DestChainID:         hexutil.Big(destChainIDBig),
		StreamOffset:        hexutil.Big(streamOffset),
		TxHash:              common.Hash(msg.TxHash),
		Data:                msg.Data,
		BlockHeight:         hexutil.Big(blockHeight),
		BlockHash:           common.Hash(block.BlockHash),
	}, nil
}

// EntReceiptToGraphQLXReceipt converts an ent.Receipt to a resolvers.XReceipt.
func EntReceiptToGraphQLXReceipt(receipt *ent.Receipt) (*resolvers.XReceipt, error) {
	gasUsed, err := Uint2Big(receipt.GasUsed)
	if err != nil {
		return nil, errors.Wrap(err, "decoding gas used")
	}

	sourceChainIDBig, err := Uint2Big(receipt.SourceChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destChainIDBig, err := Uint2Big(receipt.DestChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	streamOffset, err := Uint2Big(receipt.StreamOffset)
	if err != nil {
		return nil, errors.Wrap(err, "decoding stream offset")
	}

	return &resolvers.XReceipt{
		UUID:           graphql.ID(receipt.UUID.String()),
		Success:        graphql.NullBool{Value: &receipt.Success, Set: receipt.Success},
		GasUsed:        hexutil.Big(gasUsed),
		RelayerAddress: common.Address(receipt.RelayerAddress),
		SourceChainID:  hexutil.Big(sourceChainIDBig),
		DestChainID:    hexutil.Big(destChainIDBig),
		StreamOffset:   hexutil.Big(streamOffset),
		TxHash:         common.Hash(receipt.TxHash),
		Timestamp:      graphql.Time{Time: receipt.CreatedAt},
	}, nil
}
