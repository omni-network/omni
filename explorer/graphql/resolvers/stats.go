package resolvers

import (
	"context"

	"github.com/omni-network/omni/explorer/graphql/data"
)

type StatsProvider interface {
	Stats(ctx context.Context) data.StatsResult
}

type StatsResolver struct {
	Provider StatsProvider
}

func (s *StatsResolver) Stats(ctx context.Context) data.StatsResult {
	return s.Provider.Stats(ctx)
}
