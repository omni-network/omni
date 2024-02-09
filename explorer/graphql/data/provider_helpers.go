package data

import (
	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
)

// EntBlockToGraphQLBlock converts an ent.Block to a resolvers.XBlock.
func EntBlockToGraphQLBlock(
	block *ent.Block,
) (*resolvers.XBlock, error) {
	sourceChainIDBig, err := hexutil.DecodeBig(hexutil.EncodeUint64(block.SourceChainID))
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	blockHeight, err := hexutil.DecodeBig(hexutil.EncodeUint64(block.BlockHeight))
	if err != nil {
		return nil, errors.Wrap(err, "decoding block height")
	}

	res := resolvers.XBlock{
		SourceChainID: hexutil.Big(*sourceChainIDBig),
		BlockHeight:   hexutil.Big(*blockHeight),
		BlockHash:     common.Hash(block.BlockHash),
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
		receipt, err := EntReceiptToGraphQLXReceipt(receipt)
		if err != nil {
			return nil, errors.Wrap(err, "decoding receipt for block")
		}
		res.Receipts = append(res.Receipts, *receipt)
	}

	return &res, nil
}

// EntMsgToGraphQLXMsg converts an ent.Msg to a resolvers.XMsg.
func EntMsgToGraphQLXMsg(
	msg *ent.Msg,
) (*resolvers.XMsg, error) {
	sourceChainIDBig, err := hexutil.DecodeBig(hexutil.EncodeUint64(msg.SourceChainID))
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destChainIDBig, err := hexutil.DecodeBig(hexutil.EncodeUint64(msg.DestChainID))
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destGasLimit, err := hexutil.DecodeBig(hexutil.EncodeUint64(msg.DestGasLimit))
	if err != nil {
		return nil, errors.Wrap(err, "decoding dest gas limit")
	}

	streamOffset, err := hexutil.DecodeBig(hexutil.EncodeUint64(msg.StreamOffset))
	if err != nil {
		return nil, errors.Wrap(err, "decoding stream offset")
	}

	return &resolvers.XMsg{
		SourceMessageSender: common.Address(msg.SourceMsgSender),
		SourceChainID:       hexutil.Big(*sourceChainIDBig),
		DestAddress:         common.Address(msg.DestAddress),
		DestGasLimit:        hexutil.Big(*destGasLimit),
		DestChainID:         hexutil.Big(*destChainIDBig),
		StreamOffset:        hexutil.Big(*streamOffset),
		TxHash:              common.Hash(msg.TxHash),
		Data:                msg.Data,
	}, nil
}

// EntReceiptToGraphQLXReceipt converts an ent.Receipt to a resolvers.XReceipt.
func EntReceiptToGraphQLXReceipt(
	receipt *ent.Receipt,
) (*resolvers.XReceipt, error) {
	gasUsed, err := hexutil.DecodeBig(hexutil.EncodeUint64(receipt.GasUsed))
	if err != nil {
		return nil, errors.Wrap(err, "decoding gas used")
	}

	sourceChainIDBig, err := hexutil.DecodeBig(hexutil.EncodeUint64(receipt.SourceChainID))
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	destChainIDBig, err := hexutil.DecodeBig(hexutil.EncodeUint64(receipt.DestChainID))
	if err != nil {
		return nil, errors.Wrap(err, "decoding source chain id")
	}

	streamOffset, err := hexutil.DecodeBig(hexutil.EncodeUint64(receipt.StreamOffset))
	if err != nil {
		return nil, errors.Wrap(err, "decoding stream offset")
	}

	return &resolvers.XReceipt{
		UUID:           graphql.ID(receipt.UUID.String()),
		Success:        graphql.NullBool{Value: &receipt.Success, Set: receipt.Success},
		GasUsed:        hexutil.Big(*gasUsed),
		RelayerAddress: common.Address(receipt.RelayerAddress),
		SourceChainID:  hexutil.Big(*sourceChainIDBig),
		DestChainID:    hexutil.Big(*destChainIDBig),
		StreamOffset:   hexutil.Big(*streamOffset),
		TxHash:         common.Hash(receipt.TxHash),
		Timestamp:      graphql.Time{Time: receipt.CreatedAt},
	}, nil
}
