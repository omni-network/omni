package provider

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

// NewForT returns a new provider for testing. Note that cprovider isn't supported yet.
func NewForT(
	t *testing.T,
	network netconf.Network,
	rpcClients map[uint64]ethclient.Client,
	backoffFunc func(ctx context.Context) func(),
	workers int) *Provider {
	t.Helper()

	for i := range fetchWorkerThresholds {
		fetchWorkerThresholds[i].Workers = uint64(workers)
	}

	return &Provider{
		network:     network,
		ethClients:  rpcClients,
		backoffFunc: backoffFunc,
		confHeads:   make(map[chainVersion]uint64),
	}
}

func TestOffsetTracker(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	expectedOffsetsByHeight := map[uint64]uint64{ // map[height]offset
		1: 1,
		2: 0,
		3: 0,
		4: 2,
		5: 3,
		6: 0,
		7: 4,
	}

	// Start at 1,1
	tracker := newOffsetTracker(1, 1, 0)

	errChan := make(chan error)

	// Concurrently (and in random order), call awaitOffset.
	for height, offset := range expectedOffsetsByHeight {
		go func() {
			var msgs []xchain.Msg
			if offset > 0 {
				// If an offset is expected, include an xmsg in the block.
				msgs = append(msgs, xchain.Msg{})
			}

			actualOffset, err := tracker.awaitOffset(ctx, xchain.Block{
				BlockHeader: xchain.BlockHeader{BlockHeight: height},
				Msgs:        msgs,
			})
			if err != nil {
				errChan <- err
				return
			}
			if offset != actualOffset {
				errChan <- errors.New("unexpected offset", "actual", actualOffset, "expected", offset)
				return
			}
			errChan <- nil // All good
		}()
	}

	// Collect the results, ensuring no errors.
	for range expectedOffsetsByHeight {
		require.NoError(t, <-errChan)
	}
}
