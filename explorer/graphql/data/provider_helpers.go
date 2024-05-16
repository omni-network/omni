package data

import (
	"context"
	"math/big"
	"strconv"
	"time"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/graphql/uintconv"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// EntBlockToGraphQLBlock converts an ent.Block to a XBlock.
func EntBlockToGraphQLBlock(_ context.Context, block *ent.Block) (*XBlock, error) {
	sourceChainIDBig, err := uintconv.ToBig(block.ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	blockHeight, err := uintconv.ToBig(block.Height)
	if err != nil {
		return nil, errors.Wrap(err, "decoding block height")
	}

	res := XBlock{
		ChainID:   hexutil.Big(sourceChainIDBig),
		Height:    hexutil.Big(blockHeight),
		Hash:      common.Hash(block.Hash),
		Timestamp: graphql.Time{Time: block.Timestamp},
	}

	// Decode messages
	for _, msg := range block.Edges.Msgs {
		msg, err := EntMsgToGraphQLXMsg(msg)
		if err != nil {
			return nil, errors.Wrap(err, "decoding msg for block")
		}
		res.Messages = append(res.Messages, *msg)
	}

	return &res, nil
}

// EntMsgToGraphQLXMsg converts an ent.Msg to a XMsg.
func EntMsgToGraphQLXMsg(msg *ent.Msg) (*XMsg, error) {
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

	// blockHeight, err := uintconv.ToBig(msg.BlockHeight)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "decoding block height")
	// }

	xmsg := &XMsg{
		ID:            graphql.ID(strconv.Itoa(msg.ID)),
		Sender:        common.Address(msg.Sender),
		SourceChainID: hexutil.Big(sourceChainIDBig),
		To:            common.Address(msg.To),
		GasLimit:      hexutil.Big(destGasLimit),
		DestChainID:   hexutil.Big(destChainIDBig),
		Offset:        hexutil.Big(offset),
		TxHash:        common.Hash(msg.TxHash),
		Data:          msg.Data,
		Status:        Status(msg.Status),
	}

	// if len(msg.ReceiptHash) == 32 {
	// 	hash := common.Hash(msg.ReceiptHash)
	// 	xmsg.ReceiptTxHash = &hash
	// }

	return xmsg, nil
}

// EntReceiptToGraphQLXReceipt converts an ent.Receipt to a XReceipt.
func EntReceiptToGraphQLXReceipt(ctx context.Context, receipt *ent.Receipt, block *ent.Block) (*XReceipt, error) {
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

	// blockHeight, err := uintconv.ToBig(block.Height)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "decoding block height")
	// }

	res := &XReceipt{
		ID:            relay.MarshalID("Receipt", receipt.ID),
		Success:       receipt.Success,
		GasUsed:       hexutil.Big(gasUsed),
		Relayer:       common.Address(receipt.RelayerAddress),
		SourceChainID: hexutil.Big(sourceChainIDBig),
		DestChainID:   hexutil.Big(destChainIDBig),
		Offset:        hexutil.Big(streamOffset),
		TxHash:        common.Hash(receipt.TxHash),
		Timestamp:     graphql.Time{Time: receipt.CreatedAt},
	}

	return res, nil
}

func EntChainToGraphQLChain(chain *ent.Chain) Chain {
	chainID, err := uintconv.ToBig(chain.ChainID)
	if err != nil {
		panic(errors.Wrap(err, "decoding chain id"))
	}

	return Chain{
		Name:      chain.Name,
		ChainID:   hexutil.Big(chainID),
		DisplayID: Long(chain.ChainID),
	}
}

// EntMsgToGraphQLXMsg converts an ent.Msg to a XMsg.
func EntMsgToGraphQLXMsgWithEdges(ctx context.Context, msg *ent.Msg) (*XMsg, error) {
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

	receipts, err := msg.QueryReceipts().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "querying receipts for message")
	}

	var xreceipt *XReceipt
	var rts time.Time
	for _, r := range receipts {
		rec, err := EntReceiptToGraphQLXReceipt(ctx, r, block)
		if err != nil {
			return nil, errors.Wrap(err, "decoding receipt for message")
		}
		// only use the latest receipt if more than one is available
		if r.CreatedAt.After(rts) {
			rts = r.CreatedAt
			xreceipt = rec
		}
	}

	b := msg.QueryBlock().OnlyX(ctx)
	xblock := XBlock{
		ChainID:   hexutil.Big(sourceChainIDBig),
		Height:    hexutil.Big(*big.NewInt(int64(b.Height))),
		Hash:      common.Hash(b.Hash),
		Timestamp: graphql.Time{Time: block.Timestamp},
	}
	xmsg := &XMsg{
		ID:            graphql.ID(strconv.Itoa(msg.ID)),
		Sender:        common.Address(msg.Sender),
		SourceChainID: hexutil.Big(sourceChainIDBig),
		To:            common.Address(msg.To),
		GasLimit:      hexutil.Big(destGasLimit),
		DestChainID:   hexutil.Big(destChainIDBig),
		Offset:        hexutil.Big(streamOffset),
		TxHash:        common.Hash(msg.TxHash),
		Data:          msg.Data,
		Receipt:       xreceipt,
		Block:         xblock,
		Status:        MustParseStatus(msg.Status),
	}

	return xmsg, nil
}
