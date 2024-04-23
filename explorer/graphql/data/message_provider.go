package data

import (
	"context"
	"strconv"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/explorer/graphql/utils"
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

	hex, err := utils.Uint2Hex(uint64(query))
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

//nolint:dupl // graphql library looks for the function name to match the resolver
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
		receipt, err := EntReceiptToGraphQLXReceipt(ctx, r, block)
		if err != nil {
			return nil, false, errors.Wrap(err, "decoding receipt")
		}
		receiptsRes = append(receiptsRes, *receipt)
	}

	res.Receipts = receiptsRes

	return res, true, nil
}

func (p Provider) XMsgs(ctx context.Context, limit uint64, cursor *uint64) (*resolvers.XMsgResult, bool, error) {
	query := p.EntClient.Msg.Query().
		Order(ent.Desc(msg.FieldCreatedAt)).
		Limit(int(limit)) // limit will always set, defaulting to 1

	// If cursor is not 0, we want to query the message with the cursor ID.
	if cursor != nil {
		query = query.Where(msg.IDLTE(int(*cursor)))
	}

	// Execute the query.
	msgs, err := query.All(ctx)
	if err != nil {
		log.Error(ctx, "Msgs query", err)
		return nil, false, err
	}

	// Create the xmsg array
	var res []resolvers.XMsgEdge
	for _, m := range msgs {
		graphQL, err := EntMsgToGraphQLXMsgWithEdges(ctx, m)
		if err != nil {
			return nil, false, errors.Wrap(err, "decoding message")
		}
		cursor, err := utils.Uint2Hex(uint64(m.ID))
		if err != nil {
			return nil, false, errors.Wrap(err, "decoding message cursor")
		}
		res = append(res, resolvers.XMsgEdge{
			Cursor: cursor,
			Node:   *graphQL,
		})
	}

	// Get the total count of messages
	totalCount, err := p.EntClient.Msg.Query().Count(ctx)
	if err != nil {
		return nil, false, errors.New("failed to fetch message count")
	}

	// Get the total count in hex
	totalCountHex, err := utils.Uint2Hex(uint64(totalCount))
	if err != nil {
		return nil, false, errors.Wrap(err, "decoding message count")
	}

	// Get the start cursor
	startCursor, err := strconv.ParseUint(string(res[0].Node.ID), 10, 64)
	if err != nil {
		return nil, false, errors.New("failed to parse start cursor")
	}

	endCursor, err := strconv.ParseUint(string(res[len(res)-1].Node.ID), 10, 64)
	if err != nil {
		return nil, false, errors.New("failed to parse end cursor")
	}

	// Get the start cursor in hex
	c, err := utils.Uint2Hex(endCursor + 1)
	if err != nil {
		return nil, false, errors.Wrap(err, "decoding message cursor")
	}

	// Create the result
	result := resolvers.XMsgResult{
		TotalCount: totalCountHex,
		Edges:      res,
		PageInfo: resolvers.PageInfo{
			StartCursor: c,
			HasNextPage: endCursor-uint64(1) > 0,
			HasPrevPage: startCursor > 0,
		},
	}

	return &result, true, nil
}
