// Package provider implements the cchain.Provider interface.
package provider

import (
	"context"
	"path"
	"strings"
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

	rpcclient "github.com/cometbft/cometbft/rpc/client"
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
type genesisFunc func(ctx context.Context) (execution []byte, consensus []byte, err error)

type valSetResponse struct {
	ValSetID      uint64
	Validators    []cchain.Validator
	CreatedHeight uint64
	activedHeight uint64
}

// Provider implements cchain.Provider.
type Provider struct {
	cometCl     rpcclient.Client
	fetch       fetchFunc
	latest      latestFunc
	window      windowFunc
	valset      valsetFunc
	chainID     chainIDFunc
	header      headerFunc
	portalBlock portalBlockFunc
	networkFunc networkFunc
	genesisFunc genesisFunc
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

func (p Provider) CometClient() rpcclient.Client {
	return p.cometCl
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

func (p Provider) GenesisFiles(ctx context.Context) (execution []byte, consensus []byte, err error) { //nolint:nonamedreturns // Disambiguate identical return types
	return p.genesisFunc(ctx)
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
			return att.AttestOffset
		},
		Verify: func(_ context.Context, att xchain.Attestation, offset uint64) error {
			if !chainVer.ConfLevel.IsFuzzy() && att.ChainVersion.ConfLevel.IsFuzzy() {
				return errors.New("fuzzy attestation while streaming finalized [BUG]")
			} else if att.AttestOffset != offset {
				return errors.New("invalid attestation offset",
					"actual", att.AttestOffset,
					"expected", offset,
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

// ErrHistoryPruned indicates that the necessary state for the requested height isn't found in the store.
var ErrHistoryPruned = errors.New("no commit info found (history pruned)")

// IsErrHistoryPruned reports whether the input error matches the CosmosSDK errors returned when
// the state for the requested height isn't found in the store.
func IsErrHistoryPruned(err error) bool {
	if err == nil {
		return false
	}

	// There are two possible errors CosmosSDK returns when the state for the requested height isn't found in the store.
	// First: https://github.com/cosmos/cosmos-sdk/blob/1bbb499cdf32dbf2bed3607860c30693c3f5674a/baseapp/abci.go#L1244
	// Second: https://github.com/cosmos/cosmos-sdk/blob/7edd86813f4b17bed6f603bc5b3629a1a5aa41e8/store/rootmulti/store.go#L1134
	return strings.Contains(err.Error(), "failed to load state at height") || strings.Contains(err.Error(), "no commit info found")
}
