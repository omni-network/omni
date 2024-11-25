package e2e_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/e2e/solve/devapp"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

// TestSolver submits deposits to the solve inbox and waits for them to be processed.
func TestSolver(t *testing.T) {
	t.Parallel()
	skipFunc := func(manifest types.Manifest) bool {
		return !manifest.DeploySolve
	}
	maybeTestNetwork(t, skipFunc, func(t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()
		ctx := context.Background()

		backends, err := ethbackend.BackendsFromNetwork(network, endpoints)
		require.NoError(t, err)

		deposits, err := devapp.RequestDeposits(ctx, backends)
		require.NoError(t, err)

		timeout, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()

		toCheck := toSet(deposits)
		for {
			if timeout.Err() != nil {
				require.Fail(t, "timeout waiting for deposits")
			}

			for deposit := range toCheck {
				ok, err := devapp.IsDeposited(ctx, backends, deposit)
				require.NoError(t, err)
				if ok {
					log.Info(ctx, "Deposit complete", "remaining", len(toCheck)-1)
					delete(toCheck, deposit)
				}
			}

			if len(toCheck) == 0 {
				return
			}

			time.Sleep(time.Second)
		}
	})
}

func toSet[T comparable](slice []T) map[T]bool {
	set := make(map[T]bool)
	for _, v := range slice {
		set[v] = true
	}

	return set
}
