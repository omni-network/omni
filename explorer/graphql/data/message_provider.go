package data

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p Provider) XMsgCount(ctx context.Context) (*hexutil.Big, bool, error) {
	query, err := p.EntClient.Msg.Query().Count(ctx)
	if err != nil {
		log.Error(ctx, "Msg count query", err)
		return nil, false, err
	}

	hex, err := Uint2Hex(uint64(query))
	if err != nil {
		return nil, false, errors.Wrap(err, "decoding block count")
	}

	return &hex, true, nil
}

func (p Provider) XMsgRange(ctx context.Context, from uint64, to uint64) ([]*resolvers.XMsg, bool, error) {
	amount := to - from
	query, err := p.EntClient.Msg.Query().
		Order(ent.Desc(msg.FieldCreatedAt)).
		Offset(int(from)).
		Limit(int(amount)).
		All(ctx)
	if err != nil {
		log.Error(ctx, "Msg range query", err)
		return nil, false, err
	}

	var res []*resolvers.XMsg
	for _, m := range query {
		graphQL, err := EntMsgToGraphQLXMsg(ctx, m, nil)
		if err != nil {
			return nil, false, errors.Wrap(err, "decode message")
		}
		res = append(res, graphQL)
	}

	return res, true, nil
}

func (p Provider) XMsg(ctx context.Context, sourceChainID, destChainID, streamOffset uint64) (*resolvers.XMsg, bool, error) {
	query, err := p.EntClient.Msg.Query().
		Where(
			msg.SourceChainID(sourceChainID),
			msg.DestChainID(destChainID),
			msg.StreamOffset(streamOffset),
		).
		First(ctx)
	if err != nil {
		log.Error(ctx, "Msg query", err)
		return nil, false, err
	}

	block := query.QueryBlock().OnlyX(ctx)
	receipts := query.QueryReceipts().AllX(ctx)

	res, err := EntMsgToGraphQLXMsg(ctx, query, block)
	if err != nil {
		return nil, false, errors.Wrap(err, "decoding message")
	}

	var receiptsRes []resolvers.XReceipt
	for _, r := range receipts {
		receipt, err := EntReceiptToGraphQLXReceipt(r)
		if err != nil {
			return nil, false, errors.Wrap(err, "decoding receipt")
		}
		receiptsRes = append(receiptsRes, *receipt)
	}

	res.Receipts = receiptsRes

	return res, true, nil
}
