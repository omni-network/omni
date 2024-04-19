package data

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/block"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/explorer/graphql/utils"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p Provider) XBlock(ctx context.Context, sourceChainID uint64, height uint64) (*resolvers.XBlock, bool, error) {
	query, err := p.EntClient.Block.Query().
		Where(block.SourceChainID(sourceChainID)).
		Where(block.BlockHeight(height)).
		First(ctx)
	if err != nil {
		log.Error(ctx, "Graphql provider err", err)
		return nil, false, err
	}

	msgs, err := query.QueryMsgs().All(ctx)
	if err != nil {
		log.Error(ctx, "Block msg edges query", err)
		return nil, false, err
	}

	receipts, err := query.QueryReceipts().All(ctx)
	if err != nil {
		log.Error(ctx, "Block receipt edges query", err)
		return nil, false, err
	}

	query.Edges.Msgs = msgs
	query.Edges.Receipts = receipts

	b, err := EntBlockToGraphQLBlock(ctx, query)
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to decode block")
	}

	return b, true, nil
}

func (p Provider) XBlockRange(ctx context.Context, from uint64, to uint64) ([]*resolvers.XBlock, bool, error) {
	amount := to - from
	query, err := p.EntClient.Block.Query().
		Order(ent.Desc(block.FieldTimestamp)).
		Limit(int(amount)).
		Offset(int(from)).
		All(ctx)
	if err != nil {
		log.Error(ctx, "Ent query", err)
		return nil, false, err
	}

	var res []*resolvers.XBlock
	for _, b := range query {
		graphQL, err := EntBlockToGraphQLBlock(ctx, b)
		if err != nil {
			return nil, false, errors.Wrap(err, "failed to decode block")
		}

		res = append(res, graphQL)
	}

	return res, true, nil
}

func (p Provider) XBlockCount(ctx context.Context) (*hexutil.Big, bool, error) {
	query, err := p.EntClient.Block.Query().Count(ctx)
	if err != nil {
		log.Error(ctx, "Graphql provider err", err)
		return nil, false, err
	}

	res := query

	big, err := utils.Uint2Hex(uint64(res))
	if err != nil {
		return nil, false, errors.Wrap(err, "decoding block count")
	}

	return &big, true, nil
}
