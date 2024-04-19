package data

import (
	"context"
	"strconv"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/explorer/graphql/utils"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
)

// EntBlockToGraphQLBlock converts an ent.Block to a resolvers.XBlock.
func EntBlockToGraphQLBlock(ctx context.Context, block *ent.Block) (*resolvers.XBlock, error) {
	sourceChainIDBig, err := utils.Uint2Big(block.SourceChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	blockHeight, err := utils.Uint2Big(block.BlockHeight)
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
		receipt, err := EntReceiptToGraphQLXReceipt(ctx, receipt, block)
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

	sourceChainIDBig, err := utils.Uint2Big(msg.SourceChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destChainIDBig, err := utils.Uint2Big(msg.DestChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destGasLimit, err := utils.Uint2Big(msg.DestGasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "decoding dest gas limit")
	}

	streamOffset, err := utils.Uint2Big(msg.StreamOffset)
	if err != nil {
		return nil, errors.Wrap(err, "decoding stream offset")
	}

	blockHeight, err := utils.Uint2Big(block.BlockHeight)
	if err != nil {
		return nil, errors.Wrap(err, "decoding block height")
	}

	return &resolvers.XMsg{
		ID:                  graphql.ID(strconv.Itoa(msg.ID)),
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
func EntReceiptToGraphQLXReceipt(ctx context.Context, receipt *ent.Receipt, block *ent.Block) (*resolvers.XReceipt, error) {
	if block == nil {
		b, err := receipt.QueryBlock().Only(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "querying block for message")
		}
		block = b
	}

	gasUsed, err := utils.Uint2Big(receipt.GasUsed)
	if err != nil {
		return nil, errors.Wrap(err, "decoding gas used")
	}

	sourceChainIDBig, err := utils.Uint2Big(receipt.SourceChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destChainIDBig, err := utils.Uint2Big(receipt.DestChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	streamOffset, err := utils.Uint2Big(receipt.StreamOffset)
	if err != nil {
		return nil, errors.Wrap(err, "decoding stream offset")
	}

	blockHeight, err := utils.Uint2Big(block.BlockHeight)
	if err != nil {
		return nil, errors.Wrap(err, "decoding block height")
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
		BlockHeight:    hexutil.Big(blockHeight),
		BlockHash:      common.Hash(block.BlockHash),
	}, nil
}

func EntChainToGraphQLChain(chain *ent.Chain) resolvers.Chain {
	chainID, err := utils.Uint2Big(chain.ChainID)
	if err != nil {
		panic(errors.Wrap(err, "decoding chain id"))
	}

	return resolvers.Chain{
		Name:    chain.Name,
		ChainID: hexutil.Big(chainID),
	}
}

// EntMsgToGraphQLXMsg converts an ent.Msg to a resolvers.XMsg.
func EntMsgToGraphQLXMsgWithEdges(ctx context.Context, msg *ent.Msg) (*resolvers.XMsg, error) {
	block, err := msg.QueryBlock().Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "querying block for message")
	}

	sourceChainIDBig, err := utils.Uint2Big(msg.SourceChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destChainIDBig, err := utils.Uint2Big(msg.DestChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destGasLimit, err := utils.Uint2Big(msg.DestGasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "decoding dest gas limit")
	}

	streamOffset, err := utils.Uint2Big(msg.StreamOffset)
	if err != nil {
		return nil, errors.Wrap(err, "decoding stream offset")
	}

	blockHeight, err := utils.Uint2Big(block.BlockHeight)
	if err != nil {
		return nil, errors.Wrap(err, "decoding block height")
	}

	receipts, err := msg.QueryReceipts().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "querying receipts for message")
	}

	var xreceipts []resolvers.XReceipt
	for _, r := range receipts {
		r, err := EntReceiptToGraphQLXReceipt(ctx, r, block)
		if err != nil {
			return nil, errors.Wrap(err, "decoding receipt for message")
		}
		xreceipts = append(xreceipts, *r)
	}

	b := msg.QueryBlock().OnlyX(ctx)
	xblock := resolvers.XBlock{
		SourceChainID: hexutil.Big(sourceChainIDBig),
		BlockHeight:   hexutil.Big(blockHeight),
		BlockHash:     common.Hash(b.BlockHash),
		Timestamp:     graphql.Time{Time: block.CreatedAt},
	}

	return &resolvers.XMsg{
		ID:                  graphql.ID(strconv.Itoa(msg.ID)),
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
		Receipts:            xreceipts,
		Block:               xblock,
	}, nil
}
