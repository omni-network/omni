package data

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/block"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p Provider) XBlock(sourceChainID uint64, height uint64) (*resolvers.XBlock, bool, error) {
	ctx := context.Background()
	query, err := p.EntClient.Block.Query().
		Where(block.SourceChainID(sourceChainID)).
		Where(block.BlockHeight(height)).
		First(ctx)

	if err != nil {
		log.Error(ctx, "Graphql provider err", err)
		return nil, false, err
	}

	graphQLBlock, err := EntBlockToGraphQLBlock(query)
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to decode block")
	}

	return graphQLBlock, true, nil
}

func (p Provider) XBlockRange(amount uint64, offset uint64) ([]*resolvers.XBlock, bool, error) {
	ctx := context.Background()
	query, err := p.EntClient.Block.Query().
		Order(ent.Asc(block.FieldTimestamp)).
		Limit(int(amount)).
		Offset(int(offset)).
		All(ctx)

	if err != nil {
		log.Error(ctx, "Graphql provider err", err)
		return nil, false, err
	}

	res := []*resolvers.XBlock{}
	for _, b := range query {
		graphQL, err := EntBlockToGraphQLBlock(b)
		if err != nil {
			return nil, false, errors.Wrap(err, "failed to decode block")
		}

		res = append(res, graphQL)
	}

	return res, true, nil
}

func (p Provider) XBlockCount() (*hexutil.Big, bool, error) {
	ctx := context.Background()
	query, err := p.EntClient.Block.Query().
		Count(ctx)

	if err != nil {
		log.Error(ctx, "Graphql provider err", err)
		return nil, false, err
	}

	res := query

	big, err := hexutil.DecodeBig(hexutil.EncodeUint64(uint64(res)))
	if err != nil {
		return nil, false, errors.Wrap(err, "decoding block count")
	}

	b := hexutil.Big(*big)

	return &b, true, nil
}
