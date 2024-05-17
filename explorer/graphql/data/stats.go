package data

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/lib/log"
)

const (
	statsRefreshInterval = 30 * time.Second
	statusQueryTimeout   = 15 * time.Second
)

// StatsProvider provides stats data.
type StatsProvider struct {
	ch    Chainer
	cl    *ent.Client
	stats StatsResult
}

func NewStatsProvider(ctx context.Context, cl *ent.Client, ch Chainer) *StatsProvider {
	res := &StatsProvider{
		ch: ch,
		cl: cl,
	}

	// populate the stats for the first time
	res.updateStats(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(statsRefreshInterval):
				// populates stats
				res.updateStats(ctx)
			}
		}
	}(ctx)

	return res
}

func (s *StatsProvider) Stats(ctx context.Context) StatsResult {
	return s.stats
}

func (s *StatsProvider) TotalMsgs() uint64 {
	return uint64(s.stats.TotalMsgs)
}

func (s *StatsProvider) updateStats(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, statusQueryTimeout)
	defer cancel()
	stats := StatsResult{}

	total, err := s.cl.Msg.Query().Count(ctx)
	if err != nil {
		log.Warn(ctx, "Calling updateStats(): Msg count query", err)
	}
	stats.TotalMsgs = Long(total)

	// Aggregate multiple fields.
	// var streams []StreamStats

	var v []struct {
		SourceChainID uint64 `json:"source_chain_id"`
		DestChainID   uint64 `json:"dest_chain_id"`
		Count         uint64 `json:"count"`
	}
	// ent doesn't support ordering by aggregated column :(
	err = s.cl.Msg.Query().
		Select(msg.FieldSourceChainID, msg.FieldDestChainID).
		GroupBy(msg.FieldSourceChainID, msg.FieldDestChainID).
		Aggregate(
			ent.Count(),
		).
		Scan(ctx, &v)
	if err != nil {
		log.Warn(ctx, "Calling updateStats(): top streams query", err)
	}

	sort.Slice(v, func(i, j int) bool {
		return v[i].Count > v[j].Count
	})

	for i, stream := range v {
		if i == 3 {
			break
		}
		sc, _ := s.ch.Chain(fmt.Sprintf("0x%x", stream.SourceChainID))
		dc, _ := s.ch.Chain(fmt.Sprintf("0x%x", stream.DestChainID))
		stats.TopStreams = append(stats.TopStreams, StreamStats{
			SourceChain: sc,
			DestChain:   dc,
			MsgCount:    Long(stream.Count),
		})
	}

	s.stats = stats
}
