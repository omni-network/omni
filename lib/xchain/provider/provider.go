// Package provider is the implementation of the Provider interface.
// it abstracts connecting to different rollup chains and collecting
// XMsgs and XReceipts to construct XBlock and deliver them to the
// calling module.
package provider

import (
	"context"
	"path"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/stream"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/xchain"

	"go.opentelemetry.io/otel/trace"
)

const (
	streamTypeXBlock = "xblock"
	streamTypeEvent  = "event"
)

// fetchWorkerThresholds defines the number of concurrent workers
// to fetch xblocks by chain block period.
var fetchWorkerThresholds = []struct {
	Workers   uint64        // Number of workers
	MinPeriod time.Duration // Minimum threshold
}{
	{Workers: 1, MinPeriod: time.Second},     // 1 worker for normal chains.
	{Workers: 2, MinPeriod: time.Second / 2}, // 2 workers for fast chains (op_sepolia)
	{Workers: 4, MinPeriod: 0},               // 4 workers for fastest chains (arb_sepolia)
}

// getWorkers returns the workers from the first exceeded threshold that matches.
func getWorkers(chain netconf.Chain) (uint64, error) {
	for _, threshold := range fetchWorkerThresholds {
		if chain.BlockPeriod >= threshold.MinPeriod {
			return threshold.Workers, nil
		}
	}

	return 0, errors.New("no matching threshold [BUG]")
}

var _ xchain.Provider = (*Provider)(nil)

// Provider stores the source chain configuration and the global quit channel.
type Provider struct {
	network     netconf.Network
	ethClients  map[uint64]ethclient.Client // store config for every chain ID
	cChainID    uint64
	cProvider   cchain.Provider
	backoffFunc func(context.Context) func()

	mu sync.Mutex
	// confHeads caches the latest height by chain version.
	// It reduces HeaderByType queries if the stream is lagging
	// behind the chain version head.
	// Also, since many L2s finalize in batches, the stream
	// lags behind the chain version head every time a new batch is finalized.
	confHeads map[xchain.ChainVersion]uint64
}

// New instantiates the provider instance which will be ready to accept
// subscriptions for respective destination XBlocks.
func New(network netconf.Network, rpcClients map[uint64]ethclient.Client, cProvider cchain.Provider) *Provider {
	backoffFunc := func(ctx context.Context) func() {
		// Limit backoff to 10s for all EVM chains.
		const maxDelay = time.Second * 10
		cfg := expbackoff.DefaultConfig
		cfg.MaxDelay = maxDelay

		return expbackoff.New(ctx, expbackoff.With(cfg))
	}

	cChain, _ := network.OmniConsensusChain()

	return &Provider{
		network:     network,
		ethClients:  rpcClients,
		cChainID:    cChain.ID,
		cProvider:   cProvider,
		backoffFunc: backoffFunc,
		confHeads:   make(map[xchain.ChainVersion]uint64),
	}
}

// StreamAsync starts a goroutine that streams xblocks asynchronously forever.
// It returns immediately. It only returns an error if the chainID in invalid.
// This is the async version of StreamBlocks.
// It retries forever (with backoff) on all fetch and callback errors.
func (p *Provider) StreamAsync(
	ctx context.Context,
	req xchain.ProviderRequest,
	callback xchain.ProviderCallback,
) error {
	if _, ok := p.network.Chain(req.ChainID); !ok {
		return errors.New("unknown chain ID")
	}

	go func() {
		err := p.stream(ctx, req, callback, true)
		if err != nil { // RetryCallback==true so this should only ever return nil on ctx cancel.
			log.Error(ctx, "Streaming xprovider blocks failed unexpectedly [BUG]", err)
		}
	}()

	return nil
}

// StreamBlocks blocks, streaming all xblocks from the chain as they become available (finalized).
// It retries forever (with backoff) on all fetch errors. It however returns the first callback error.
// It returns nil when the context is canceled.
func (p *Provider) StreamBlocks(
	ctx context.Context,
	req xchain.ProviderRequest,
	callback xchain.ProviderCallback,
) error {
	return p.stream(ctx, req, callback, false)
}

func (p *Provider) stream(
	ctx context.Context,
	req xchain.ProviderRequest,
	callback xchain.ProviderCallback,
	retryCallback bool,
) error {
	chain, ok := p.network.Chain(req.ChainID)
	if !ok {
		return errors.New("unknown chain ID")
	} else if evmchain.IsSVM(chain.ID) {
		return errors.New("svm chains are not supported")
	}

	chainVersionName := p.network.ChainVersionName(xchain.ChainVersion{ID: req.ChainID, ConfLevel: req.ConfLevel})

	workers, err := getWorkers(chain)
	if err != nil {
		return err
	}

	// Start streaming from chain's deploy height as per config.
	fromHeight := req.Height
	if fromHeight < chain.DeployHeight {
		fromHeight = chain.DeployHeight
	}

	deps := stream.Deps[xchain.Block]{
		FetchWorkers: workers,
		FetchBatch: func(ctx context.Context, height uint64) ([]xchain.Block, error) {
			fetchReq := xchain.ProviderRequest{
				ChainID:   req.ChainID,
				Height:    height,
				ConfLevel: req.ConfLevel,
			}

			// Retry fetching blocks a few times, since RPC providers load balance requests and some servers may lag a bit.
			var block xchain.Block
			var exists bool
			err := expbackoff.Retry(ctx, func() (err error) { //nolint:nonamedreturns // Succinctness FTW
				block, exists, err = p.GetBlock(ctx, fetchReq)
				return err
			})
			if err != nil {
				return nil, err
			} else if !exists {
				return nil, nil
			}

			return []xchain.Block{block}, nil
		},
		Backoff:       p.backoffFunc,
		ElemLabel:     "block",
		HeightLabel:   "height",
		RetryCallback: retryCallback,
		Height: func(block xchain.Block) uint64 {
			return block.BlockHeight
		},
		Verify: func(_ context.Context, block xchain.Block, h uint64) error {
			if block.ChainID != req.ChainID {
				return errors.New("invalid block source chain id")
			} else if block.BlockHeight != h {
				return errors.New("invalid block height")
			}

			lag := time.Since(block.Timestamp)
			streamLag.WithLabelValues(chainVersionName, streamTypeXBlock).Set(lag.Seconds())

			return nil
		},
		IncFetchErr: func() {
			fetchErrTotal.WithLabelValues(chainVersionName, streamTypeXBlock).Inc()
		},
		IncCallbackErr: func() {
			callbackErrTotal.WithLabelValues(chainVersionName, streamTypeXBlock).Inc()
		},
		SetStreamHeight: func(h uint64) {
			streamHeight.WithLabelValues(chainVersionName, streamTypeXBlock).Set(float64(h))
		},
		SetCallbackLatency: func(d time.Duration) {
			callbackLatency.WithLabelValues(chainVersionName, streamTypeXBlock).Observe(d.Seconds())
		},
		StartTrace: func(ctx context.Context, height uint64, spanName string) (context.Context, trace.Span) {
			return tracer.StartChainHeight(ctx, p.network.ID.String(), chain.Name, height, path.Join("xblock", spanName))
		},
	}

	cb := (stream.Callback[xchain.Block])(callback)

	ctx = log.WithCtx(ctx, "chain_version", chainVersionName)
	log.Info(ctx, "Streaming xprovider blocks", "from_height", fromHeight)

	return stream.Stream(ctx, deps, fromHeight, cb)
}

// getEVMChain returns the configuration, eth client and header cache of the given EVM chainID.
func (p *Provider) getEVMChain(chainID uint64) (netconf.Chain, ethclient.Client, error) {
	if chainID == p.cChainID {
		return netconf.Chain{}, nil, errors.New("consensus chain not supported")
	} else if evmchain.IsSVM(chainID) {
		return netconf.Chain{}, nil, errors.New("svm chains are not supported")
	}

	chain, ok := p.network.Chain(chainID)
	if !ok {
		return netconf.Chain{}, nil, errors.New("unknown chain ID for network")
	}

	client, ok := p.ethClients[chainID]
	if !ok {
		return netconf.Chain{}, nil, errors.New("no rpc client for chain ID")
	}

	return chain, client, nil
}
