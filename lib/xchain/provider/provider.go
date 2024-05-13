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

const initialXOffset uint64 = 1

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
	fromOffset uint64,
	callback xchain.ProviderCallback,
) error {
	if _, ok := p.network.Chain(chainID); !ok {
		return errors.New("unknown chain ID")
	}

	go func() {
		err := p.stream(ctx, chainID, fromHeight, &fromOffset, callback, true)
		if err != nil { // RetryCallback==true so this should only ever return nil on ctx cancel.
			log.Error(ctx, "Streaming xprovider blocks failed unexpectedly [BUG]", err)
		}
	}()

	return nil
}

// StreamAsyncNoOffset is the same as StreamAsync except that XBlockOffset is not populated in the returned XBlocks.
func (p *Provider) StreamAsyncNoOffset(
	ctx context.Context,
	chainID uint64,
	fromHeight uint64,
	callback xchain.ProviderCallback,
) error {
	if _, ok := p.network.Chain(chainID); !ok {
		return errors.New("unknown chain ID")
	}

	go func() {
		err := p.stream(ctx, chainID, fromHeight, nil, callback, true)
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
	fromOffset uint64,
	callback xchain.ProviderCallback,
) error {
	return p.stream(ctx, chainID, fromHeight, &fromOffset, callback, false)
}

// StreamBlocksNoOffset is the same as StreamBlocks except that XBlockOffset is not populated in the returned XBlocks.
func (p *Provider) StreamBlocksNoOffset(
	ctx context.Context,
	chainID uint64,
	fromHeight uint64,
	callback xchain.ProviderCallback,
) error {
	return p.stream(ctx, chainID, fromHeight, nil, callback, false)
}

func (p *Provider) stream(
	ctx context.Context,
	chainID uint64,
	fromHeight uint64,
	fromOffset *uint64,
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
	if workers == 0 {
		return errors.New("zero workers [BUG]")
	}

	// Start streaming from chain's deploy height as per config.
	if fromHeight < chain.DeployHeight {
		fromHeight = chain.DeployHeight
	}
	// XBlockOffset is 1-indexed (starts at 1)
	if fromOffset != nil && *fromOffset == 0 {
		fromOffset = ptr(initialXOffset)
	}

	tracker := newOffsetTracker(fromOffset)

	deps := stream.Deps[xchain.Block]{
		FetchWorkers: workers,
		FetchBatch: func(ctx context.Context, chainID uint64, height uint64) ([]xchain.Block, error) {
			offset, err := tracker.getOffset(height)
			if err != nil {
				return nil, err
			}

			xBlock, exists, err := p.GetBlock(ctx, chainID, height, offset)
			if err != nil {
				return nil, err
			} else if !exists {
				return nil, nil
			}

			if xBlock.ShouldAttest() {
				if err := tracker.assignOffset(height); err != nil {
					return nil, err
				}
			}

			return []xchain.Block{xBlock}, nil
		},
		Backoff:       p.backoffFunc,
		ElemLabel:     "block",
		RetryCallback: retryCallback,
		Height: func(block xchain.Block) uint64 {
			return block.BlockHeight
		},
		Verify: func(ctx context.Context, block xchain.Block, h uint64) error {
			if block.SourceChainID != chainID {
				return errors.New("invalid block source chain id")
			} else if block.BlockHeight != h {
				return errors.New("invalid block height")
			} else if block.ShouldAttest() && block.BlockOffset == 0 {
				return errors.New("invalid block offset")
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

// offsetTracker tracks the XBlockOffset for non-empty blocks.
//
// It supports at-least-once semantics `assignOffset` and `getOffset`
// for monotonically incrementing heights; i.e., get and assign may
// be called for the same or later (higher) heights multiple times,
// but it may not be called for a previous (lower) height.
type offsetTracker struct {
	mu            sync.RWMutex
	enabled       bool
	currentOffset uint64
	prevHeight    *uint64
}

func newOffsetTracker(from *uint64) *offsetTracker {
	if from == nil {
		// Return a disabled tracker that is basically noop, resulting in always populating 0 XBlockOffsets.
		return &offsetTracker{}
	}

	return &offsetTracker{
		enabled:       true,
		currentOffset: *from,
	}
}

// assignOffset assigns the current offset to the provided height
// incrementing the current offset by one.
func (c *offsetTracker) assignOffset(height uint64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.enabled {
		return nil
	}

	if c.prevHeight != nil && *c.prevHeight > height {
		return errors.New("assign unexpected old height [BUG]")
	} else if c.prevHeight != nil && *c.prevHeight == height {
		return nil
	}

	c.prevHeight = &height
	c.currentOffset++

	return nil
}

// getOffset returns the XBlockOffset for the provided height,
// only the previous offset (current-1) or the current offset is supported.
func (c *offsetTracker) getOffset(height uint64) (uint64, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.enabled {
		return 0, nil
	}

	if c.prevHeight == nil {
		return c.currentOffset, nil
	}

	if *c.prevHeight > height {
		return 0, errors.New("get unexpected old height [BUG]")
	}
	if *c.prevHeight == height {
		return c.currentOffset - 1, nil
	}

	return c.currentOffset, nil
}

func ptr[T any](t T) *T {
	return &t
}
