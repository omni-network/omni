// Package provider implements the cchain.Provider interface.
package provider

import (
	"context"
	"path"
	"testing"
	"time"

	ptypes "github.com/omni-network/omni/halo/portal/types"
	rtypes "github.com/omni-network/omni/halo/registry/types"
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

type fetchFunc func(ctx context.Context, chainVer xchain.ChainVersion, fromOffset uint64) ([]xchain.Attestation, error)
type latestFunc func(ctx context.Context, chainVer xchain.ChainVersion) (xchain.Attestation, bool, error)
type windowFunc func(ctx context.Context, chainVer xchain.ChainVersion, xBlockOffset uint64) (int, error)
type portalBlockFunc func(ctx context.Context, blockOffset uint64, latest bool) (*ptypes.BlockResponse, bool, error)
type networkFunc func(ctx context.Context, networkID uint64, latest bool) (*rtypes.NetworkResponse, bool, error)
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
	portalBlock portalBlockFunc
	networkFunc networkFunc
	backoffFunc func(context.Context) func()
	chainNamer  func(xchain.ChainVersion) string
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
		chainNamer:  func(xchain.ChainVersion) string { return "" },
	}
}

func (p Provider) AttestationsFrom(ctx context.Context, chainVer xchain.ChainVersion, xBlockOffset uint64,
) ([]xchain.Attestation, error) {
	return p.fetch(ctx, chainVer, xBlockOffset)
}

func (p Provider) LatestAttestation(ctx context.Context, chainVer xchain.ChainVersion,
) (xchain.Attestation, bool, error) {
	return p.latest(ctx, chainVer)
}

func (p Provider) WindowCompare(ctx context.Context, chainVer xchain.ChainVersion, xBlockOffset uint64) (int, error) {
	return p.window(ctx, chainVer, xBlockOffset)
}

func (p Provider) ValidatorSet(ctx context.Context, valSetID uint64) ([]cchain.Validator, bool, error) {
	resp, ok, err := p.valset(ctx, valSetID, false)
	return resp.Validators, ok, err
}

// Subscribe implements cchain.Provider.
func (p Provider) Subscribe(in context.Context, chainVer xchain.ChainVersion, xBlockOffset uint64, workerName string,
	callback cchain.ProviderCallback,
) {
	srcChain := p.chainNamer(chainVer)
	ctx := log.WithCtx(in, "src_chain", srcChain)

	deps := stream.Deps[xchain.Attestation]{
		FetchBatch: func(ctx context.Context, _ uint64, offset uint64) ([]xchain.Attestation, error) {
			return p.fetch(ctx, chainVer, offset)
		},
		Backoff:       p.backoffFunc,
		ElemLabel:     "attestation",
		RetryCallback: true,
		FetchWorkers:  1, // Only single worker supported since we fetch batches of unknown lengths so can't shard.
		Height: func(att xchain.Attestation) uint64 {
			return att.BlockOffset
		},
		Verify: func(_ context.Context, att xchain.Attestation, h uint64) error {
			if !chainVer.ConfLevel.IsFuzzy() && att.ConfLevel.IsFuzzy() {
				return errors.New("fuzzy attestation while streaming finalized [BUG]")
			} else if att.BlockOffset != h {
				return errors.New("invalid attestation height",
					"actual", att.BlockOffset,
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
		err := stream.Stream(ctx, deps, chainVer.ID, xBlockOffset, cb)
		if err != nil { // RetryCallback==true, so this never return an error.
			log.Error(ctx, "Unexpected stream error [BUG]", err)
		}
	}()
}
