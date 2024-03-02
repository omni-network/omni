// Package provider implements the cchain.Provider interface.
package provider

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/stream"
	"github.com/omni-network/omni/lib/xchain"
)

var _ cchain.Provider = Provider{}

// FetchFunc abstracts fetching attestations from the consensus chain.
type FetchFunc func(ctx context.Context, chainID uint64, fromHeight uint64,
) ([]xchain.Attestation, error)

// LatestFunc abstracts fetching the latest attestation from the consensus chain.
type LatestFunc func(ctx context.Context, chainID uint64) (xchain.Attestation, bool, error)
type WindowFunc func(ctx context.Context, chainID uint64, height uint64) (int, error)

// Provider implements cchain.Provider.
type Provider struct {
	fetch       FetchFunc
	latest      LatestFunc
	window      WindowFunc
	backoffFunc func(context.Context) (func(), func())
	chainNames  map[uint64]string
}

// NewProviderForT creates a new provider for testing.
func NewProviderForT(_ *testing.T, fetch FetchFunc, latest LatestFunc, window WindowFunc,
	backoffFunc func(context.Context) (func(), func()),
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
		Verify: func(ctx context.Context, att xchain.Attestation, h uint64) error {
			if att.SourceChainID != srcChainID {
				return errors.New("invalid attestation source chain ID")
			} else if att.BlockHeight != h {
				return errors.New("invalid attestation height")
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
	}

	cb := (stream.Callback[xchain.Attestation])(callback)

	go func() {
		err := stream.Stream(ctx, deps, srcChainID, height, cb)
		if err != nil { // RetryCallback==true, so this never return an error.
			log.Error(ctx, "Unexpected stream error [BUG]", err)
		}
	}()
}
