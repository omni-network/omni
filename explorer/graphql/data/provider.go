package data

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/block"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
)

type Provider struct {
	EntClient *ent.Client
}

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

	sourceChainIDBig, err := hexutil.DecodeBig(hexutil.EncodeUint64(query.SourceChainID))
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to decode source chain id")
	}

	blockHeight, err := hexutil.DecodeBig(hexutil.EncodeUint64(query.BlockHeight))
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to decode block height")
	}

	res := resolvers.XBlock{
		SourceChainID: hexutil.Big(*sourceChainIDBig),
		BlockHeight:   hexutil.Big(*blockHeight),
		BlockHash:     common.Hash(query.BlockHash),
		Timestamp:     graphql.Time{Time: query.Timestamp},
	}

	return &res, true, nil
}
