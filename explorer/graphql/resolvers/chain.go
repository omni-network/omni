package resolvers

import "github.com/omni-network/omni/explorer/graphql/data"

type ChainsProvider interface {
	ChainList() []data.Chain
}

type Chains struct {
	list []data.Chain
}

func NewChainsProvider(p ChainsProvider) *Chains {
	return &Chains{
		list: p.ChainList(), // with the current implementation chains are not added dynamically, so load the chains only once
	}
}

func (c *Chains) SupportedChains() []data.Chain {
	return c.list
}
