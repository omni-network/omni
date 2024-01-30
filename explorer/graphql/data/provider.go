package data

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/block"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
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

	sourceChainIDString := hexutil.EncodeUint64(query.SourceChainID)
	heightString := hexutil.EncodeUint64(query.BlockHeight)

	res := resolvers.XBlock{
		SourceChainID: hexutil.Big(*hexutil.MustDecodeBig(sourceChainIDString)),
		BlockHeight:   hexutil.Big(*hexutil.MustDecodeBig(heightString)),
		BlockHash:     common.Hash(query.BlockHash),
		Timestamp:     graphql.Time{Time: query.Timestamp},
	}

	return &res, true, nil
}
