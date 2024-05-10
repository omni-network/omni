package data

import (
	"context"
	"strconv"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/explorer/graphql/uintconv"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// EntBlockToGraphQLBlock converts an ent.Block to a resolvers.XBlock.
func EntBlockToGraphQLBlock(ctx context.Context, block *ent.Block) (*resolvers.XBlock, error) {
	sourceChainIDBig, err := uintconv.ToBig(block.ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	blockHeight, err := uintconv.ToBig(block.Height)
	if err != nil {
		return nil, errors.Wrap(err, "decoding block height")
	}

	res := resolvers.XBlock{
		SourceChainID: hexutil.Big(sourceChainIDBig),
		BlockHeight:   hexutil.Big(blockHeight),
		BlockHash:     common.Hash(block.Hash),
		Timestamp:     graphql.Time{Time: block.Timestamp},
	}

	// Decode messages
	for _, msg := range block.Edges.Msgs {
		msg, err := EntMsgToGraphQLXMsg(msg)
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
func EntMsgToGraphQLXMsg(msg *ent.Msg) (*resolvers.XMsg, error) {
	sourceChainIDBig, err := uintconv.ToBig(msg.SourceChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destChainIDBig, err := uintconv.ToBig(msg.DestChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destGasLimit, err := uintconv.ToBig(msg.GasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "decoding dest gas limit")
	}

	offset, err := uintconv.ToBig(msg.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "decoding stream offset")
	}

	blockHeight, err := uintconv.ToBig(msg.BlockHeight)
	if err != nil {
		return nil, errors.Wrap(err, "decoding block height")
	}

	xmsg := &resolvers.XMsg{
		ID:                  graphql.ID(strconv.Itoa(msg.ID)),
		SourceMessageSender: common.Address(msg.Sender),
		SourceChainID:       hexutil.Big(sourceChainIDBig),
		DestAddress:         common.Address(msg.To),
		DestGasLimit:        hexutil.Big(destGasLimit),
		DestChainID:         hexutil.Big(destChainIDBig),
		StreamOffset:        hexutil.Big(offset),
		TxHash:              common.Hash(msg.TxHash),
		Data:                msg.Data,
		BlockHeight:         hexutil.Big(blockHeight),
		BlockHash:           common.Hash(msg.BlockHash),
		SourceBlockTime:     graphql.Time{Time: msg.BlockTime},
		Status:              msg.Status,
	}

	if len(msg.ReceiptHash) == 32 {
		hash := common.Hash(msg.ReceiptHash)
		xmsg.ReceiptTxHash = &hash
	}

	return xmsg, nil
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

	gasUsed, err := uintconv.ToBig(receipt.GasUsed)
	if err != nil {
		return nil, errors.Wrap(err, "decoding gas used")
	}

	sourceChainIDBig, err := uintconv.ToBig(receipt.SourceChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destChainIDBig, err := uintconv.ToBig(receipt.DestChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	streamOffset, err := uintconv.ToBig(receipt.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "decoding stream offset")
	}

	blockHeight, err := uintconv.ToBig(block.Height)
	if err != nil {
		return nil, errors.Wrap(err, "decoding block height")
	}

	return &resolvers.XReceipt{
		UUID:           relay.MarshalID("Receipt", receipt.ID),
		Success:        graphql.NullBool{Value: &receipt.Success, Set: receipt.Success},
		GasUsed:        hexutil.Big(gasUsed),
		RelayerAddress: common.Address(receipt.RelayerAddress),
		SourceChainID:  hexutil.Big(sourceChainIDBig),
		DestChainID:    hexutil.Big(destChainIDBig),
		StreamOffset:   hexutil.Big(streamOffset),
		TxHash:         common.Hash(receipt.TxHash),
		Timestamp:      graphql.Time{Time: receipt.CreatedAt},
		BlockHeight:    hexutil.Big(blockHeight),
		BlockHash:      common.Hash(block.Hash),
	}, nil
}

func EntChainToGraphQLChain(chain *ent.Chain) resolvers.Chain {
	chainID, err := uintconv.ToBig(chain.ChainID)
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

	sourceChainIDBig, err := uintconv.ToBig(msg.SourceChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destChainIDBig, err := uintconv.ToBig(msg.DestChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destGasLimit, err := uintconv.ToBig(msg.GasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "decoding dest gas limit")
	}

	streamOffset, err := uintconv.ToBig(msg.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "decoding stream offset")
	}

	blockHeight, err := uintconv.ToBig(msg.BlockHeight)
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
		BlockHash:     common.Hash(b.Hash),
		Timestamp:     graphql.Time{Time: block.Timestamp},
	}
	xmsg := &resolvers.XMsg{
		ID:                  graphql.ID(strconv.Itoa(msg.ID)),
		SourceMessageSender: common.Address(msg.Sender),
		SourceChainID:       hexutil.Big(sourceChainIDBig),
		DestAddress:         common.Address(msg.To),
		DestGasLimit:        hexutil.Big(destGasLimit),
		DestChainID:         hexutil.Big(destChainIDBig),
		StreamOffset:        hexutil.Big(streamOffset),
		TxHash:              common.Hash(msg.TxHash),
		Data:                msg.Data,
		BlockHeight:         hexutil.Big(blockHeight),
		BlockHash:           common.Hash(msg.BlockHash),
		Receipts:            xreceipts,
		Block:               xblock,
		SourceBlockTime:     graphql.Time{Time: msg.BlockTime},
		Status:              msg.Status,
	}

	if len(msg.ReceiptHash) == 32 {
		hash := common.Hash(msg.ReceiptHash)
		xmsg.ReceiptTxHash = &hash
	}

	return xmsg, nil
}
