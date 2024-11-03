package relayer

import (
	"context"
	"maps"
	"sync"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

// TestLeaderStreamer tests the leader streamer.
func TestLeaderStreamer(t *testing.T) {
	t.Parallel()

	// upstream tracks upstream requests and returns test responses
	upstream := newUpstream()
	upstreamFunc := func(
		ctx context.Context,
		chainVer xchain.ChainVersion,
		attestOffset uint64,
		workerName string,
		callback cchain.ProviderCallback,
	) error {
		resps := upstream.Req(workerName, attestOffset)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case resp := <-resps:
				err := callback(ctx, xchain.Attestation{AttestHeader: xchain.AttestHeader{ChainVersion: chainVer, AttestOffset: resp}})
				if err != nil {
					return err
				}
			}
		}
	}

	streamer := newLeaderStreamer(upstreamFunc, netconf.Simnet)

	ctx := context.Background()
	var eg errgroup.Group
	errDone := errors.New("done")

	// startStream starts a worker stream from `from` to `to`, ensuring strictly sequential attestations.
	startStream := func(worker string, from uint64, to uint64) error {
		next := from
		return streamer(ctx, xchain.ChainVersion{}, from, worker, func(ctx context.Context, att xchain.Attestation) error {
			if att.AttestOffset != next {
				return errors.New("unexpected offset")
			}
			next++

			if att.AttestOffset == to {
				return errDone
			}

			return nil
		})
	}

	w1 := "worker1"
	w2 := "worker2"
	w3 := "worker3"

	eg.Go(func() error {
		// worker 1 streams from 3 to 5 as leader
		return startStream(w1, 3, 5)
	})
	upstream.Respond(w1, 3) // w1 starts leader streaming at 3
	eg.Go(func() error {
		// worker 2 streams from 1 to 7, 1-4 as leader, 4-5 from cache, 6-7 as leader
		return startStream(w2, 1, 7)
	})
	upstream.Respond(w2, 1) // w2 starts leader streaming at 1
	upstream.Respond(w1, 4) // w1 continues leader streaming
	upstream.Respond(w2, 2) // w2 continues leader streaming
	upstream.Respond(w2, 3) // w2 overlaps w1, switch to cache
	upstream.Respond(w1, 5) // w1 is done
	upstream.Respond(w2, 6) // w2 switches back to lead streaming
	upstream.Respond(w2, 7) // w2 is done

	eg.Go(func() error {
		// worker 3 streams from 8 to 8 as leader
		return startStream(w3, 1, 8)
	})
	upstream.Respond(w3, 8) // w3 starts leader streaming at 8 and is done

	require.ErrorIs(t, eg.Wait(), errDone)
	require.EqualValues(t, map[string][]uint64{
		w1: {3},
		w2: {1, 6},
		w3: {8},
	}, upstream.Reqs())
}

func newUpstream() *upstream {
	return &upstream{
		reqs:  make(map[string][]uint64),
		resps: make(map[string]chan uint64),
	}
}

type upstream struct {
	mu    sync.Mutex
	reqs  map[string][]uint64
	resps map[string]chan uint64
}

func (u *upstream) Respond(worker string, offset uint64) {
	u.mu.Lock()
	resp, ok := u.resps[worker]
	if !ok {
		resp = make(chan uint64)
		u.resps[worker] = resp
	}
	u.mu.Unlock()

	resp <- offset
}

func (u *upstream) Reqs() map[string][]uint64 {
	u.mu.Lock()
	defer u.mu.Unlock()

	return maps.Clone(u.reqs)
}

func (u *upstream) Req(worker string, from uint64) chan uint64 {
	u.mu.Lock()
	u.reqs[worker] = append(u.reqs[worker], from)

	resp, ok := u.resps[worker]
	if !ok {
		resp = make(chan uint64)
		u.resps[worker] = resp
	}
	u.mu.Unlock()

	return resp
}
