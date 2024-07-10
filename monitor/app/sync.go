package monitor

import (
	"context"
	"math"
	"time"

	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

// startMonitoringSyncDiff starts a goroutine per chain that periodically calculates the rpc-sync-diff
// which indicates the difference in latest heights (sync) of possible HA upstream RPC servers.
// A rpc-sync-diff of 2 or more can indicate the upstreams are not 100% consistently synced.
func startMonitoringSyncDiff(ctx context.Context, network netconf.Network, ethClients map[uint64]ethclient.Client) {
	for chainID, ethCl := range ethClients {
		go func(chainID uint64, ethCl ethclient.Client) {
			ticker := time.NewTicker(time.Second * 20)
			defer ticker.Stop()

			const parallel = 3
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					diff := calcMaxDiff(concurrentLatestHeights(ctx, ethCl, parallel))
					syncDiff.WithLabelValues(network.ChainName(chainID)).Set(float64(diff))
				}
			}
		}(chainID, ethCl)
	}
}

// calcMaxDiff returns the maximum difference between the provided heights.
func calcMaxDiff(heights []int64) int64 {
	if len(heights) == 0 {
		return 0
	}

	minimum, maximum := int64(math.MaxInt64), int64(math.MinInt64)
	for _, height := range heights {
		if minimum > height {
			minimum = height
		}
		if maximum < height {
			maximum = height
		}
	}

	return maximum - minimum
}

// concurrentLatestHeights returns the non-zero latest heights returned from N concurrent queries.
func concurrentLatestHeights(ctx context.Context, ethCl ethclient.Client, n int) []int64 {
	collect := make(chan uint64, n)
	for i := 0; i < n; i++ {
		go func() {
			number, err := ethCl.BlockNumber(ctx)
			if err != nil {
				log.Warn(ctx, "Failed to get block number", err)
			}
			collect <- number
		}()
	}

	var resp []int64
	for i := 0; i < n; i++ {
		height := <-collect
		if height > 0 {
			resp = append(resp, int64(height))
		}
	}

	return resp
}
