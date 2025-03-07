package provider

import (
	"context"
	"path"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/stream"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/core/types"

	"go.opentelemetry.io/otel/trace"
)

// events extends zero or more event logs with an explicit height.
type events struct {
	Height uint64
	Events []types.Log
}

func (p *Provider) StreamEventLogs(ctx context.Context, req xchain.EventLogsReq, callback xchain.EventLogsCallback) error {
	if req.Height == 0 {
		return errors.New("invalid zero height")
	}

	chain, _, err := p.getEVMChain(req.ChainID)
	if err != nil {
		return err
	}

	workers, err := getWorkers(chain)
	if err != nil {
		return err
	}

	chainVersionName := p.network.ChainVersionName(req.ChainVersion())
	ctx = log.WithCtx(ctx, "chain", chainVersionName)

	deps := stream.Deps[events]{
		FetchWorkers: workers,
		FetchBatch: func(ctx context.Context, height uint64) ([]events, error) {
			fetchReq := xchain.EventLogsReq{
				ChainID:       req.ChainID,
				Height:        height,
				ConfLevel:     req.ConfLevel,
				FilterAddress: req.FilterAddress,
				FilterTopics:  req.FilterTopics,
				Delay:         req.Delay,
			}

			// Retry fetching logs a few times, since RPC providers load balance requests and some servers may lag a bit.
			var logs []types.Log
			var exists bool
			err := expbackoff.Retry(ctx, func() (err error) { //nolint:nonamedreturns // Succinctness FTW
				logs, exists, err = p.GetEventLogs(ctx, fetchReq)
				return err
			})
			if err != nil {
				return nil, err
			} else if !exists {
				return nil, nil
			}

			return []events{{
				Height: height,
				Events: logs,
			}}, nil
		},
		Backoff:       p.backoffFunc,
		ElemLabel:     "events",
		HeightLabel:   "height",
		RetryCallback: false,
		Height: func(logs events) uint64 {
			return logs.Height
		},
		Verify: func(_ context.Context, events events, h uint64) error {
			if events.Height != h {
				return errors.New("invalid block height")
			}

			return nil
		},
		IncFetchErr: func() {
			fetchErrTotal.WithLabelValues(chainVersionName, streamTypeEvent).Inc()
		},
		IncCallbackErr: func() {
			callbackErrTotal.WithLabelValues(chainVersionName, streamTypeEvent).Inc()
		},
		SetStreamHeight: func(h uint64) {
			streamHeight.WithLabelValues(chainVersionName, streamTypeEvent).Set(float64(h))
		},
		SetCallbackLatency: func(d time.Duration) {
			callbackLatency.WithLabelValues(chainVersionName, streamTypeEvent).Observe(d.Seconds())
		},
		StartTrace: func(ctx context.Context, height uint64, spanName string) (context.Context, trace.Span) {
			return tracer.StartChainHeight(ctx, p.network.ID, chain.Name, height, path.Join("events", spanName))
		},
	}

	return stream.Stream[events](ctx, deps, req.Height, func(ctx context.Context, events events) error {
		return callback(ctx, events.Height, events.Events)
	})
}

// GetEventLogs returns the evm event logs for the provided request, or false if not available yet (not finalized),
// or an error.
func (p *Provider) GetEventLogs(ctx context.Context, req xchain.EventLogsReq) ([]types.Log, bool, error) {
	ctx, span := tracer.Start(ctx, spanName("get_events"))
	defer span.End()

	// Height too small, we can't start processing yet.
	if req.Height < req.Delay {
		return nil, false, nil
	}

	req.Height = umath.SubtractOrZero(req.Height, req.Delay)

	_, ethCl, err := p.getEVMChain(req.ChainID)
	if err != nil {
		return nil, false, err
	}

	// First check if height is confirmed.
	var header *types.Header
	if !p.confirmedCache(req.ChainVersion(), req.Height) {
		// No higher cached header available, so fetch the latest head
		latest, err := p.headerByChainVersion(ctx, req.ChainVersion())
		if err != nil {
			return nil, false, errors.Wrap(err, "header by strategy")
		}

		// If still lower, we reached the head of the chain, return false
		if latest.Number.Uint64() < req.Height {
			return nil, false, nil
		}

		// Use this header if it matches height
		if latest.Number.Uint64() == req.Height {
			header = latest
		}
	}

	// Fetch the header if we didn't find it in the cache
	if header == nil {
		header, err = ethCl.HeaderByNumber(ctx, umath.NewBigInt(req.Height))
		if err != nil {
			return nil, false, errors.Wrap(err, "header by number")
		}
	}

	events, err := getEventLogs(ctx, ethCl, req.FilterAddress, header.Hash(), req.FilterTopics)
	if err != nil {
		return nil, false, err
	}

	return events, true, nil
}
