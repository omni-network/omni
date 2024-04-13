// Package provider implements the cchain.Provider interface.
package provider

import (
	"context"
	"path"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/stream"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/xchain"

	ctypes "github.com/cometbft/cometbft/rpc/core/types"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var _ cchain.Provider = Provider{}

type fetchFunc func(ctx context.Context, chainID uint64, fromHeight uint64) ([]xchain.Attestation, error)
type latestFunc func(ctx context.Context, chainID uint64) (xchain.Attestation, bool, error)
type windowFunc func(ctx context.Context, chainID uint64, height uint64) (int, error)
type valsetFunc func(ctx context.Context, valSetID uint64, latest bool) (valSetResponse, bool, error)
type headerFunc func(ctx context.Context, height *int64) (*ctypes.ResultHeader, error)
type chainIDFunc func(ctx context.Context) (uint64, error)

type valSetResponse struct {
	ValSetID      uint64
	Validators    []cchain.Validator
	CreatedHeight uint64
	activedHeight uint64
}

// Provider implements cchain.Provider.
type Provider struct {
	fetch       fetchFunc
	latest      latestFunc
	window      windowFunc
	valset      valsetFunc
	chainID     chainIDFunc
	header      headerFunc
	backoffFunc func(context.Context) func()
	chainNames  map[uint64]string
	network     netconf.ID
}

// NewProviderForT creates a new provider for testing.
func NewProviderForT(_ *testing.T, fetch fetchFunc, latest latestFunc, window windowFunc,
	backoffFunc func(context.Context) func(),
) Provider {
	return Provider{
		latest:      latest,
		fetch:       fetch,
		window:      window,
		backoffFunc: backoffFunc,
	}
}

func (p Provider) AttestationsFrom(ctx context.Context, sourceChainID uint64, sourceHeight uint64,
) ([]xchain.Attestation, error) {
	return p.fetch(ctx, sourceChainID, sourceHeight)
}

func (p Provider) LatestAttestation(ctx context.Context, sourceChainID uint64,
) (xchain.Attestation, bool, error) {
	return p.latest(ctx, sourceChainID)
}

func (p Provider) WindowCompare(ctx context.Context, sourceChainID uint64, height uint64) (int, error) {
	return p.window(ctx, sourceChainID, height)
}

func (p Provider) ValidatorSet(ctx context.Context, valSetID uint64) ([]cchain.Validator, bool, error) {
	resp, ok, err := p.valset(ctx, valSetID, false)
	return resp.Validators, ok, err
}

// Subscribe implements cchain.Provider.
func (p Provider) Subscribe(in context.Context, srcChainID uint64, height uint64, workerName string,
	callback cchain.ProviderCallback,
) {
	srcChain := p.chainNames[srcChainID]
	ctx := log.WithCtx(in, "src_chain", srcChain)

	deps := stream.Deps[xchain.Attestation]{
		FetchBatch:    p.fetch,
		Backoff:       p.backoffFunc,
		ElemLabel:     "attestation",
		RetryCallback: true,
		FetchWorkers:  1, // Only single worker supported since we fetch batches of unknown lengths so can't shard.
		Height: func(att xchain.Attestation) uint64 {
			return att.BlockHeight
		},
		Verify: func(ctx context.Context, att xchain.Attestation, h uint64) error {
			if att.SourceChainID != srcChainID {
				return errors.New("invalid attestation source chain ID")
			} else if att.BlockHeight != h {
				return errors.New("invalid attestation height",
					"actual", att.BlockHeight,
					"expected", h,
				)
			}

			return nil
		},
		IncFetchErr: func() {
			fetchErrTotal.WithLabelValues(workerName, srcChain).Inc()
		},
		IncCallbackErr: func() {
			callbackErrTotal.WithLabelValues(workerName, srcChain).Inc()
		},
		SetStreamHeight: func(h uint64) {
			streamHeight.WithLabelValues(workerName, srcChain).Set(float64(h))
		},
		SetCallbackLatency: func(d time.Duration) {
			callbackLatency.WithLabelValues(workerName, srcChain).Observe(d.Seconds())
		},
		StartTrace: func(ctx context.Context, height uint64, spanName string) (context.Context, trace.Span) {
			return tracer.StartChainHeight(ctx, p.network, srcChain, height,
				path.Join("cprovider", spanName),
				trace.WithAttributes(attribute.String("worker", workerName)),
			)
		},
	}

	cb := (stream.Callback[xchain.Attestation])(callback)

	go func() {
		err := stream.Stream(ctx, deps, srcChainID, height, cb)
		if err != nil { // RetryCallback==true, so this never return an error.
			log.Error(ctx, "Unexpected stream error [BUG]", err)
		}
	}()
}
