package data

import (
	"context"

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
		graphQL, err := EntMsgToGraphQLXMsg(m)
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

	res, err := EntMsgToGraphQLXMsgWithEdges(ctx, query)
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
		// Most recent messages first
		Order(ent.Desc(msg.FieldBlockTime), ent.Desc(msg.FieldStreamOffset)).
		// limit will always set, defaulting to 1
		Limit(int(limit))

	// If cursor is not 0, we want to query the message with the cursor ID.
	if cursor != nil {
		val := int(*cursor)
		// We query by less than or equal to ensure that we are going down the stream of messages
		query = query.Where(msg.IDLTE(val))
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
		graphQL, err := EntMsgToGraphQLXMsg(m)
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
		return nil, false, errors.Wrap(err, "fetch message count")
	}

	// Get the total count in hex
	totalCountHex, err := utils.Uint2Hex(uint64(totalCount))
	if err != nil {
		return nil, false, errors.Wrap(err, "decoding message count")
	}

	// Get the start cursor
	startCursor := res[0].Cursor.ToInt().Uint64()

	// Calculate the page info
	pageInfo, err := calculatePageInfo(startCursor, limit, totalCount)
	if err != nil {
		return nil, false, errors.Wrap(err, "calculating page info")
	}

	// Create the result
	result := resolvers.XMsgResult{
		TotalCount: totalCountHex,
		Edges:      res,
		PageInfo:   pageInfo,
	}

	return &result, true, nil
}

// calculatePageInfo calculates the next and previous cursors for a given start cursor and limit
// The next cursor is the start cursor - the limit meaning we are moving down the stream of messages, towards the first/oldest
// The previous cursor is the start cursor + the limit meaning we are moving up the stream of messages, towards the most recent.
func calculatePageInfo(startCursor, limit uint64, totalCount int) (resolvers.PageInfo, error) {
	prevCursor := startCursor + limit

	nextCursor := startCursor - limit
	if int64(startCursor)-int64(limit) < 0 {
		nextCursor = uint64(0)
	}

	// convert the cursors to hex
	prevCursorHex, err := utils.Uint2Hex(prevCursor)
	if err != nil {
		return resolvers.PageInfo{}, errors.Wrap(err, "decoding message cursor")
	}

	nextCursorHex, err := utils.Uint2Hex(nextCursor)
	if err != nil {
		return resolvers.PageInfo{}, errors.Wrap(err, "decoding message cursor")
	}

	return resolvers.PageInfo{
		NextCursor:  nextCursorHex,
		PrevCursor:  prevCursorHex,
		HasNextPage: nextCursor > 0,
		HasPrevPage: startCursor < uint64(totalCount),
	}, nil
}
