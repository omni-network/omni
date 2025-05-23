package provider

import (
	"context"
	"path"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/stream"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/core/types"

	"go.opentelemetry.io/otel/trace"
)

// events extends zero or more event logs with an explicit block header.
type events struct {
	Header *types.Header
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

	deps := stream.Deps[events]{
		FetchWorkers: workers,
		FetchBatch: func(ctx context.Context, height uint64) ([]events, error) {
			fetchReq := xchain.EventLogsReq{
				ChainID:         req.ChainID,
				Height:          height,
				ConfLevel:       req.ConfLevel,
				FilterAddresses: req.FilterAddresses,
				FilterTopics:    req.FilterTopics,
			}

			// Retry fetching logs a few times, since RPC providers load balance requests and some servers may lag a bit.
			var logs []types.Log
			var header *types.Header
			var exists bool
			err := expbackoff.Retry(ctx, func() (err error) { //nolint:nonamedreturns // Succinctness FTW
				logs, header, exists, err = p.GetEventLogs(ctx, fetchReq)
				return err
			})
			if err != nil {
				return nil, err
			} else if !exists {
				return nil, nil
			}

			return []events{{
				Header: header,
				Events: logs,
			}}, nil
		},
		Backoff:       p.backoffFunc,
		ElemLabel:     "events",
		HeightLabel:   "height",
		RetryCallback: false,
		Height: func(e events) uint64 {
			return e.Header.Number.Uint64()
		},
		Verify: func(_ context.Context, e events, h uint64) error {
			if e.Header.Number.Uint64() != h {
				return errors.New("invalid block height")
			}

			timestamp, err := umath.ToInt64(e.Header.Time)
			if err != nil {
				return err
			}
			lag := time.Since(time.Unix(timestamp, 0))
			streamLag.WithLabelValues(chainVersionName, streamTypeEvent).Set(lag.Seconds())

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
			return tracer.StartChainHeight(ctx, p.network.ID.String(), chain.Name, height, path.Join("events", spanName))
		},
	}

	return stream.Stream[events](ctx, deps, req.Height, func(ctx context.Context, events events) error {
		return callback(ctx, events.Header, events.Events)
	})
}

// GetEventLogs returns the evm event logs for the provided request, or false if not available yet (not finalized),
// or an error.
func (p *Provider) GetEventLogs(ctx context.Context, req xchain.EventLogsReq) ([]types.Log, *types.Header, bool, error) {
	ctx, span := tracer.Start(ctx, spanName("get_events"))
	defer span.End()

	_, ethCl, err := p.getEVMChain(req.ChainID)
	if err != nil {
		return nil, nil, false, err
	}

	// First check if height is confirmed.
	var header *types.Header
	if !p.confirmedCache(req.ChainVersion(), req.Height) {
		// No higher cached header available, so fetch the latest head
		latest, err := p.headerByChainVersion(ctx, req.ChainVersion())
		if err != nil {
			return nil, nil, false, errors.Wrap(err, "header by strategy")
		}

		// If still lower, we reached the head of the chain, return false
		if latest.Number.Uint64() < req.Height {
			return nil, nil, false, nil
		}

		// Use this header if it matches height
		if latest.Number.Uint64() == req.Height {
			header = latest
		}
	}

	// Fetch the header if we didn't find it in the cache
	if header == nil {
		header, err = ethCl.HeaderByNumber(ctx, bi.N(req.Height))
		if err != nil {
			return nil, nil, false, errors.Wrap(err, "header by number")
		}
	}

	events, err := getEventLogs(ctx, ethCl, req.FilterAddresses, header.Hash(), req.FilterTopics)
	if err != nil {
		return nil, nil, false, err
	}

	return events, header, true, nil
}
