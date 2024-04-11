package data

import (
	"context"

	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/lib/log"
)

func (p Provider) SupportedChains(ctx context.Context) ([]*resolvers.Chain, bool, error) {
	query, err := p.EntClient.Chain.Query().All(ctx)
	if err != nil {
		log.Error(ctx, "Msg count query", err)
		return nil, false, err
	}

	var res []*resolvers.Chain
	for _, chain := range query {
		c := EntChainToGraphQLChain(chain)
		res = append(res, &c)
	}

	return res, true, nil
}
