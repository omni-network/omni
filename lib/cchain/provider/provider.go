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

	"github.com/ethereum/go-ethereum/common"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/client/grpc/node"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	_ "github.com/omni-network/omni/halo/sdk" // To init SDK config.
)

var _ cchain.Provider = Provider{}

// fetchFunc returns a slice of strictly-sequential attestations for the
// provided fromOffset (inclusive) and chain version if present.
//
// It also accepts (optional) and returns an opaque cursor (when error is nil).
// When streaming attestation, the returned cursor can be provided in the subsequent call for improved performance.
type fetchFunc func(ctx context.Context, chainVer xchain.ChainVersion, fromOffset uint64, cursor uint64) ([]xchain.Attestation, uint64, error)
type allAttsFunc func(ctx context.Context, chainVer xchain.ChainVersion, fromOffset uint64) ([]xchain.Attestation, error)
type latestFunc func(ctx context.Context, chainVer xchain.ChainVersion) (xchain.Attestation, bool, error)
type windowFunc func(ctx context.Context, chainVer xchain.ChainVersion, attestOffset uint64) (int, error)
type portalBlockFunc func(ctx context.Context, attestOffset uint64, latest bool) (*ptypes.BlockResponse, bool, error)
type networkFunc func(ctx context.Context, networkID uint64, latest bool) (*rtypes.NetworkResponse, bool, error)
type valFunc func(ctx context.Context, operator common.Address) (cchain.SDKValidator, bool, error)
type valsFunc func(ctx context.Context) ([]cchain.SDKValidator, error)
type rewardsFunc func(ctx context.Context, operator common.Address) (float64, bool, error)
type valsetFunc func(ctx context.Context, valSetID uint64, latest bool) (valSetResponse, bool, error)
type chainIDFunc func(ctx context.Context) (uint64, error)
type genesisFunc func(ctx context.Context) (execution []byte, consensus []byte, err error)
type planedUpgradeFunc func(ctx context.Context) (upgradetypes.Plan, bool, error)
type appliedUpgradeFunc func(ctx context.Context, name string) (upgradetypes.Plan, bool, error)
type signingFunc func(ctx context.Context) ([]cchain.SDKSigningInfo, error)
type executionHeadFunc func(ctx context.Context) (cchain.ExecutionHead, error)

type valSetResponse struct {
	ValSetID        uint64
	Validators      []cchain.PortalValidator
	CreatedHeight   uint64
	activatedHeight uint64
}

// Provider implements cchain.Provider.
type Provider struct {
	fetch         fetchFunc
	allAtts       allAttsFunc
	latest        latestFunc
	window        windowFunc
	valset        valsetFunc
	val           valFunc
	signing       signingFunc
	vals          valsFunc
	rewards       rewardsFunc
	chainID       chainIDFunc
	portalBlock   portalBlockFunc
	networkFunc   networkFunc
	genesisFunc   genesisFunc
	plannedFunc   planedUpgradeFunc
	appliedFunc   appliedUpgradeFunc
	executionHead executionHeadFunc
	backoffFunc   func(context.Context) func()
	chainNamer    func(xchain.ChainVersion) string
	network       netconf.ID
	queryClients  cchain.QueryClients
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

func (p Provider) QueryClients() cchain.QueryClients {
	return p.queryClients
}

func (p Provider) CurrentPlannedPlan(ctx context.Context) (upgradetypes.Plan, bool, error) {
	return p.plannedFunc(ctx)
}

func (p Provider) AppliedPlan(ctx context.Context, name string) (upgradetypes.Plan, bool, error) {
	return p.appliedFunc(ctx, name)
}

// NodeStatus returns the current consensus node status.
func (p Provider) NodeStatus(ctx context.Context) (*node.StatusResponse, error) {
	status, err := p.QueryClients().Node.Status(ctx, &node.StatusRequest{})
	if err != nil {
		return &node.StatusResponse{}, errors.Wrap(err, "node status query")
	}

	return status, nil
}

func (p Provider) AttestationsFrom(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	attestOffset uint64,
) ([]xchain.Attestation, error) {
	atts, _, err := p.fetch(ctx, chainVer, attestOffset, 0)
	return atts, err
}

func (p Provider) AllAttestationsFrom(ctx context.Context, chainVer xchain.ChainVersion, attestOffset uint64,
) ([]xchain.Attestation, error) {
	return p.allAtts(ctx, chainVer, attestOffset)
}

func (p Provider) LatestAttestation(ctx context.Context, chainVer xchain.ChainVersion,
) (xchain.Attestation, bool, error) {
	return p.latest(ctx, chainVer)
}

func (p Provider) WindowCompare(ctx context.Context, chainVer xchain.ChainVersion, attestOffset uint64) (int, error) {
	return p.window(ctx, chainVer, attestOffset)
}

func (p Provider) PortalValidatorSet(ctx context.Context, valSetID uint64) ([]cchain.PortalValidator, bool, error) {
	resp, ok, err := p.valset(ctx, valSetID, false)
	return resp.Validators, ok, err
}

func (p Provider) SDKValidator(ctx context.Context, operator common.Address) (cchain.SDKValidator, bool, error) {
	return p.val(ctx, operator)
}

func (p Provider) SDKValidators(ctx context.Context) ([]cchain.SDKValidator, error) {
	return p.vals(ctx)
}

func (p Provider) SDKSigningInfos(ctx context.Context) ([]cchain.SDKSigningInfo, error) {
	return p.signing(ctx)
}

func (p Provider) SDKRewards(ctx context.Context, operator common.Address) (float64, bool, error) {
	return p.rewards(ctx, operator)
}

func (p Provider) GenesisFiles(ctx context.Context) (execution []byte, consensus []byte, err error) { //nolint:nonamedreturns // Disambiguate identical return types
	return p.genesisFunc(ctx)
}

func (p Provider) StreamAttestations(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	attestOffset uint64,
	workerName string,
	callback cchain.ProviderCallback,
) error {
	return p.stream(ctx, chainVer, attestOffset, workerName, callback, false)
}

func (p Provider) StreamAsync(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	attestOffset uint64,
	workerName string,
	callback cchain.ProviderCallback,
) {
	go func() {
		err := p.stream(ctx, chainVer, attestOffset, workerName, callback, true)
		if err != nil { // RetryCallback==true, so this never return an error.
			log.Error(ctx, "Unexpected stream error [BUG]", err)
		}
	}()
}

func (p Provider) stream(
	in context.Context,
	chainVer xchain.ChainVersion,
	attestOffset uint64,
	workerName string,
	callback cchain.ProviderCallback,
	retryCallback bool,
) error {
	if attestOffset == 0 {
		return errors.New("invalid zero attest offset [BUG]", "worker", workerName)
	}

	srcChain := p.chainNamer(chainVer)
	ctx := log.WithCtx(in, "src_chain", srcChain, "worker", workerName)

	// Cache the previous cursor (consensus height) at which we found the attestation
	// to be used in the next fetch call as the search start height
	var fetchCursor uint64

	deps := stream.Deps[xchain.Attestation]{
		FetchBatch: func(ctx context.Context, offset uint64) ([]xchain.Attestation, error) {
			atts, _, err := p.fetch(ctx, chainVer, offset, fetchCursor)
			if err != nil {
				return nil, err
			}
			// fetchCursor = cursor

			return atts, nil
		},
		Backoff:       p.backoffFunc,
		ElemLabel:     "attestation",
		HeightLabel:   "offset",
		RetryCallback: retryCallback,
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
			return tracer.StartChainHeight(ctx, p.network.String(), srcChain, height,
				path.Join("cprovider", spanName),
				trace.WithAttributes(attribute.String("worker", workerName)),
			)
		},
	}

	cb := (stream.Callback[xchain.Attestation])(callback)
	err := stream.Stream(ctx, deps, attestOffset, cb)
	if err != nil {
		return errors.Wrap(err, "stream attestations", "worker", workerName, "chain", srcChain)
	}

	return nil
}

func (p Provider) Portals(ctx context.Context) ([]rtypes.Portal, bool, error) {
	// networkID is ignored when latest is true.
	netResp, ok, err := p.networkFunc(ctx, 0, true)
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to fetch network info")
	}

	if !ok {
		return nil, ok, nil
	}

	return netResp.Portals, true, nil
}

func (p Provider) ExecutionHead(ctx context.Context) (cchain.ExecutionHead, error) {
	return p.executionHead(ctx)
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
