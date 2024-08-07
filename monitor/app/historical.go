package monitor

import (
	"context"
	"strconv"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"
)

// runHistoricalBaselineForever blocks forever, periodically running historical baselines for all chains in the network.
// Historical baseline is the time it takes to stream a certain number of historical attestations from the cprovider.
// This is used to track the performance of the cprovider historical query logic.
func runHistoricalBaselineForever(ctx context.Context, network netconf.Network, cprov cchain.Provider) {
	chainVers := keys(network.ChainVersionNames())
	sizes := []uint64{100, 1000, 10000}

	const period = time.Minute
	timer := time.NewTimer(period)
	defer timer.Stop()

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			chainVer := pick(chainVers, i)
			chainVerName := network.ChainVersionName(chainVer)
			size := pick(sizes, i)
			t0 := time.Now()

			log.Info(ctx, "Running historical baseline", "chain", chainVerName, "size", size)

			measure, err := runHistoricalBaselineOnce(ctx, cprov, chainVer, size)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Warn(ctx, "Failed running historical baseline (will retry)", err,
					"chain", chainVerName, "size", size)
			} else if measure {
				duration := time.Since(t0)
				log.Info(ctx, "Ran historical baseline", "chain", chainVerName, "size", size, "duration", duration)
				histBaseline.WithLabelValues(chainVerName, strconv.FormatUint(size, 10)).Observe(duration.Seconds())
			}
			timer.Reset(period)
		}
	}
}

func runHistoricalBaselineOnce(
	ctx context.Context,
	cprov cchain.Provider,
	chainVer xchain.ChainVersion,
	size uint64,
) (bool, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	latest, ok, err := cprov.LatestAttestation(ctx, chainVer)
	if err != nil {
		return false, err
	} else if !ok {
		return false, nil
	}

	from := umath.SubtractOrZero(latest.AttestOffset, size)
	if from == 0 {
		return false, nil // Not enough attestations to run historical baseline
	}

	var last uint64
	err = cprov.StreamAttestations(ctx, chainVer, from, "hist_baseline",
		func(_ context.Context, att xchain.Attestation) error {
			last = att.AttestOffset
			if last >= latest.AttestOffset {
				cancel()
			}

			return nil
		})
	if ctx.Err() != nil {
		return true, nil
	} else if err != nil {
		return false, errors.Wrap(err, "stream attestations", "last", last, "from", from)
	}

	return false, errors.New("unexpected stream end [BUG]")
}

func keys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func pick[V any](arr []V, i int) V {
	if len(arr) == 0 {
		var zero V
		return zero
	}

	return arr[i%len(arr)]
}
