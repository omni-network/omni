package data

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent/receipt"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/explorer/graphql/utils"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p Provider) XReceiptCount(ctx context.Context) (*hexutil.Big, bool, error) {
	query, err := p.EntClient.Receipt.Query().
		Count(ctx)
	if err != nil {
		log.Error(ctx, "Graphql provider err", err)
		return nil, false, err
	}

	hex, err := utils.Uint2Hex(uint64(query))
	if err != nil {
		return nil, false, errors.Wrap(err, "decoding block count")
	}

	return &hex, true, nil
}

//nolint:dupl // graphql library looks for the function name to match the resolver
func (p Provider) XReceipt(ctx context.Context, sourceChainID, destChainID, streamOffset uint64) (*resolvers.XReceipt, bool, error) {
	query, err := p.EntClient.Receipt.Query().
		Where(
			receipt.SourceChainID(sourceChainID),
			receipt.DestChainID(destChainID),
			receipt.StreamOffset(streamOffset),
		).
		First(ctx)
	if err != nil {
		log.Error(ctx, "Receipt query", err)
		return nil, false, err
	}

	block := query.QueryBlock().OnlyX(ctx)
	messages := query.QueryMsgs().AllX(ctx)

	res, err := EntReceiptToGraphQLXReceipt(ctx, query, block)
	if err != nil {
		return nil, false, errors.Wrap(err, "decoding message")
	}

	var msgs []resolvers.XMsg
	for _, m := range messages {
		msg, err := EntMsgToGraphQLXMsg(ctx, m, nil)
		if err != nil {
			return nil, false, errors.Wrap(err, "decoding message")
		}
		msgs = append(msgs, *msg)
	}

	res.Messages = msgs

	return res, true, nil
}
