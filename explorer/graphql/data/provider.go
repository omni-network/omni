package data

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent"
)

type Provider struct {
	*ChainsProvider
	client *ent.Client
	*StatsProvider
}

func NewProvider(ctx context.Context, cl *ent.Client, network string) *Provider {
	cp := NewChainsProvider(network)
	return &Provider{
		ChainsProvider: cp,
		client:         cl,
		StatsProvider:  NewStatsProvider(ctx, cl, cp),
	}
}
