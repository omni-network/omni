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
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/stream"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/xchain"

	"go.opentelemetry.io/otel/trace"
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

var _ xchain.Provider = (*Provider)(nil)

// Provider stores the source chain configuration and the global quit channel.
type Provider struct {
	network     netconf.Network
	ethClients  map[uint64]ethclient.Client // store config for every chain ID
	cChainID    uint64
	cProvider   cchain.Provider
	backoffFunc func(context.Context) func()

	mu sync.Mutex
	// stratHeads caches highest finalized height by chain.
	// It reduces HeaderByType queries if the stream is lagging
	// behind finalized head.
	// Also, since many L2s finalize in batches, the stream
	// lags behind finalized head every time a new batch is finalized.
	stratHeads map[uint64]uint64
}

// New instantiates the provider instance which will be ready to accept
// subscriptions for respective destination XBlocks.
func New(network netconf.Network, rpcClients map[uint64]ethclient.Client, cProvider cchain.Provider) *Provider {
	backoffFunc := func(ctx context.Context) func() {
		return expbackoff.New(ctx)
	}

	cChain, _ := network.OmniConsensusChain()

	return &Provider{
		network:     network,
		ethClients:  rpcClients,
		cChainID:    cChain.ID,
		cProvider:   cProvider,
		backoffFunc: backoffFunc,
		stratHeads:  make(map[uint64]uint64),
	}
}

// StreamAsync starts a goroutine that streams xblocks asynchronously forever.
// It returns immediately. It only returns an error if the chainID in invalid.
// This is the async version of StreamBlocks.
// It retries forever (with backoff) on all fetch and callback errors.
func (p *Provider) StreamAsync(
	ctx context.Context,
	chainID uint64,
	fromHeight uint64,
	callback xchain.ProviderCallback,
) error {
	if _, ok := p.network.Chain(chainID); !ok {
		return errors.New("unknown chain ID")
	}

	go func() {
		err := p.stream(ctx, chainID, fromHeight, callback, true)
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
	chainID uint64,
	fromHeight uint64,
	callback xchain.ProviderCallback,
) error {
	return p.stream(ctx, chainID, fromHeight, callback, false)
}

func (p *Provider) stream(
	ctx context.Context,
	chainID uint64,
	fromHeight uint64,
	callback xchain.ProviderCallback,
	retryCallback bool,
) error {
	chain, ok := p.network.Chain(chainID)
	if !ok {
		return errors.New("unknown chain ID")
	}

	var workers uint64 // Pick the first threshold that matches (or the last one)
	for _, threshold := range fetchWorkerThresholds {
		workers = threshold.Workers
		if chain.BlockPeriod >= threshold.MinPeriod {
			break
		}
	}

	deps := stream.Deps[xchain.Block]{
		FetchWorkers: workers,
		FetchBatch: func(ctx context.Context, chainID uint64, height uint64) ([]xchain.Block, error) {
			xBlock, exists, err := p.GetBlock(ctx, chainID, height)
			if err != nil {
				return nil, err
			} else if !exists {
				return nil, nil
			}

			return []xchain.Block{xBlock}, nil
		},
		Backoff:       p.backoffFunc,
		ElemLabel:     "attestation",
		RetryCallback: retryCallback,
		Height: func(block xchain.Block) uint64 {
			return block.BlockHeight
		},
		Verify: func(ctx context.Context, block xchain.Block, h uint64) error {
			if block.SourceChainID != chainID {
				return errors.New("invalid block source chain id")
			} else if block.BlockHeight != h {
				return errors.New("invalid block height")
			}

			return nil
		},
		IncFetchErr: func() {
			fetchErrTotal.WithLabelValues(chain.Name).Inc()
		},
		IncCallbackErr: func() {
			callbackErrTotal.WithLabelValues(chain.Name).Inc()
		},
		SetStreamHeight: func(h uint64) {
			streamHeight.WithLabelValues(chain.Name).Set(float64(h))
		},
		SetCallbackLatency: func(d time.Duration) {
			callbackLatency.WithLabelValues(chain.Name).Observe(d.Seconds())
		},
		StartTrace: func(ctx context.Context, height uint64, spanName string) (context.Context, trace.Span) {
			return tracer.StartChainHeight(ctx, p.network.ID, chain.Name, height,
				path.Join("xprovider", spanName),
			)
		},
	}

	cb := (stream.Callback[xchain.Block])(callback)

	// Start streaming from chain's deploy height as per config.
	if fromHeight < chain.DeployHeight {
		fromHeight = chain.DeployHeight
	}

	ctx = log.WithCtx(ctx, "chain", chain.Name)
	log.Info(ctx, "Streaming xprovider blocks", "from_height", fromHeight)

	return stream.Stream(ctx, deps, chainID, fromHeight, cb)
}

// getEVMChain provides the configuration of the given chainID.
func (p *Provider) getEVMChain(chainID uint64) (netconf.Chain, ethclient.Client, error) {
	if chainID == p.cChainID {
		return netconf.Chain{}, nil, errors.New("consensus chain not supported")
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
