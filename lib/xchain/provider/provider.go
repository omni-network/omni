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

type chainVersion struct {
	ID        uint64
	ConfLevel xchain.ConfLevel
}

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
	confHeads map[chainVersion]uint64
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
		confHeads:   make(map[chainVersion]uint64),
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
	}

	chainVersionName := p.network.ChainVersionName(xchain.ChainVersion{ID: req.ChainID, ConfLevel: req.ConfLevel})

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
	fromHeight := req.Height
	if fromHeight < chain.DeployHeight {
		fromHeight = chain.DeployHeight
	}

	// XBlockOffset is 1-indexed (starts at 1)
	fromOffset := req.Offset
	if fromOffset == 0 {
		fromOffset = initialXOffset
	}

	tracker := newOffsetTracker(fromOffset, fromHeight, chain.AttestInterval)

	deps := stream.Deps[xchain.Block]{
		FetchWorkers: workers,
		FetchBatch: func(ctx context.Context, chainID uint64, height uint64) ([]xchain.Block, error) {
			fetchReq := xchain.ProviderRequest{
				ChainID:   chainID,
				Height:    height,
				ConfLevel: req.ConfLevel,
				Offset:    0, // Don't assign block offset yet
			}

			xBlock, exists, err := p.GetBlock(ctx, fetchReq)
			if err != nil {
				return nil, err
			} else if !exists {
				return nil, nil
			}

			// Get block offset from tracker.
			// Possibly waiting for another fetch worker to fetch "next height";
			// since offsets need to be assigned in sequential order.
			xBlock.BlockOffset, err = tracker.awaitOffset(ctx, xBlock)
			if err != nil {
				return nil, err
			}

			return []xchain.Block{xBlock}, nil
		},
		Backoff:       p.backoffFunc,
		ElemLabel:     "block",
		RetryCallback: retryCallback,
		Height: func(block xchain.Block) uint64 {
			return block.BlockHeight
		},
		Verify: func(_ context.Context, block xchain.Block, h uint64) error {
			if block.SourceChainID != req.ChainID {
				return errors.New("invalid block source chain id")
			} else if block.BlockHeight != h {
				return errors.New("invalid block height")
			} else if block.ShouldAttest(chain.AttestInterval) && block.BlockOffset == 0 {
				return errors.New("invalid block offset")
			}

			return nil
		},
		IncFetchErr: func() {
			fetchErrTotal.WithLabelValues(chainVersionName).Inc()
		},
		IncCallbackErr: func() {
			callbackErrTotal.WithLabelValues(chainVersionName).Inc()
		},
		SetStreamHeight: func(h uint64) {
			streamHeight.WithLabelValues(chainVersionName).Set(float64(h))
		},
		SetCallbackLatency: func(d time.Duration) {
			callbackLatency.WithLabelValues(chainVersionName).Observe(d.Seconds())
		},
		StartTrace: func(ctx context.Context, height uint64, spanName string) (context.Context, trace.Span) {
			return tracer.StartChainHeight(ctx, p.network.ID, chain.Name, height, path.Join("xprovider", spanName))
		},
	}

	cb := (stream.Callback[xchain.Block])(callback)

	ctx = log.WithCtx(ctx, "chain", chainVersionName)
	log.Info(ctx, "Streaming xprovider blocks", "from_height", fromHeight)

	return stream.Stream(ctx, deps, req.ChainID, fromHeight, cb)
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

// offsetTracker tracks and assigns the XBlockOffset.
type offsetTracker struct {
	attestInterval uint64

	// mutable defines all mutable fields protected by a RWMutex.
	mutable struct {
		sync.RWMutex
		nextOffset uint64
		nextHeight uint64
	}
}

// newOffsetTracker returns a new offset tracker, setting the next state to the provided values.
func newOffsetTracker(nextOffset uint64, nextHeight uint64, attestInterval uint64) *offsetTracker {
	tracker := offsetTracker{
		attestInterval: attestInterval,
	}

	// Take lock, this is only require due to `go test -race`
	tracker.mutable.Lock()
	defer tracker.mutable.Unlock()

	tracker.mutable.nextOffset = nextOffset
	tracker.mutable.nextHeight = nextHeight

	return &tracker
}

// awaitOffset block and waits for the provided block to be the next expected height,
// it then maybe assigns-and-returns a block offset if block.ShouldAttest,
// it also updates internal state..
func (c *offsetTracker) awaitOffset(ctx context.Context, block xchain.Block) (uint64, error) {
	// Wait for this block to be the next expected height.
	ticker := time.NewTicker(time.Millisecond) // Avoid spinning
	defer ticker.Stop()
	ctx, cancel := context.WithTimeout(ctx, time.Minute) // Avoid blocking forever
	defer cancel()
	for c.nextHeightRLock() != block.BlockHeight {
		select {
		case <-ctx.Done():
			return 0, errors.Wrap(ctx.Err(), "timeout waiting for next height", "next_height", c.nextHeightRLock(), "height", block.BlockHeight)
		case <-ticker.C:
		}
	}

	c.mutable.Lock()
	defer c.mutable.Unlock()

	// Sanity check that next height is still as expected.
	if c.mutable.nextHeight != block.BlockHeight {
		return 0, errors.New("unexpected next height [BUG]")
	}

	// Get block offset to return, maybe increment nextOffset.
	var blockOffset uint64
	if block.ShouldAttest(c.attestInterval) {
		blockOffset = c.mutable.nextOffset
		c.mutable.nextOffset++
	}

	// Update next height.
	c.mutable.nextHeight++

	return blockOffset, nil
}

// nextHeightRLock returns the next height.
// It takes a read lock, so the write lock MUST not be held.
func (c *offsetTracker) nextHeightRLock() uint64 {
	c.mutable.RLock()
	defer c.mutable.RUnlock()

	return c.mutable.nextHeight
}
