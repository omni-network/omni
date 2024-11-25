// Package chaos provides a simple API to inject errors into applications to test error handling in-the-wild.
package chaos

import (
	"context"
	"math/rand/v2"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
)

var (
	networkErrorProb = map[netconf.ID]float64{
		netconf.Devnet:  0.001,  // 0.1% error rate in devnet (1 in 1_000)
		netconf.Staging: 0.001,  // 0.1% error rate in staging (1 in 1_000)
		netconf.Omega:   0.0001, // 0.01% error rate in omega (1 in 10_000)
		netconf.Mainnet: 0.00,   // No chaos errors in mainnet
	}

	// ErrChaos is the error returned my MaybeError. It supports checking for chaos errors via: errors.Is(err, chaos.ErrChaos).
	ErrChaos = errors.NewSentinel("chaos error")
)

type key struct{}

// WithErrProbability sets the chaos error probability for the given network in the context.
func WithErrProbability(ctx context.Context, network netconf.ID) context.Context {
	return context.WithValue(ctx, key{}, networkErrorProb[network])
}

// MaybeError sometimes returns an error based on the probability set in the context.
func MaybeError(ctx context.Context) error {
	if prob, ok := ctx.Value(key{}).(float64); ok && prob > 0 {
		if rand.Float64() < prob { //nolint:gosec // Weak randomness isn't a problem here.
			chaosErrorCount.Inc()
			return errors.Wrap(ErrChaos, "maybe error") // Wrap error for proper stack traces.
		}
	}

	return nil
}
