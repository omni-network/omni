package data

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/lib/netconf"
)

type Provider struct {
	*ChainsProvider
	client *ent.Client
	*StatsProvider
}

func NewProvider(ctx context.Context, cl *ent.Client, network netconf.ID) *Provider {
	cp := NewChainsProvider(network)
	return &Provider{
		ChainsProvider: cp,
		client:         cl,
		StatsProvider:  NewStatsProvider(ctx, cl, cp),
	}
}
