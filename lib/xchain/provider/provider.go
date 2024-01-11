// Package provider is the implementation of the Provider interface.
// it abstracts connecting to different rollup chains and collecting
// XMsgs and XReceipts to construct XBlock and deliver them to the
// calling module.
package provider

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
)

var ErrConfigNotPresent = errors.New("config for chain id is not found")

// ChainConfig is the configuration parameters for all the chains
// that needs to be managed by the provider.
type ChainConfig struct {
	name      string // name of the rollup chain
	id        uint64 // network id of the chain
	minHeight uint64 // minimum configured height from which blocks should be fetched
	// TODO(jmozah): use portal address when fetching blocks, commenting for now
	// portalAddress common.Address // the portal smart contract address
	rpcURL string // the rpc url to connect to the chain
}

// Provider stores the source chain configuration and the global quit channel.
type Provider struct {
	config []*ChainConfig // store config for every chain ID
	quitC  chan bool      // to stop all operations of the provider
}

// New instantiates the provider instance which will be ready to accept
// subscriptions for respective destination XBlocks.
func New(chains []*ChainConfig) *Provider {
	return &Provider{
		config: chains,
	}
}

// Subscribe to the XBlock from a given destination chain.
func (p *Provider) Subscribe(
	ctx context.Context,
	chainID uint64,
	minHeight uint64,
	callback xchain.ProviderCallback,
) error {
	log.Info(ctx, "Subscribing to provider ", "id", chainID, "minHeight", minHeight)

	// retrieve the respective config
	config, err := p.getConfig(chainID)
	if err != nil {
		return err
	}

	// Start streaming from chain's minimum height as per config.
	if minHeight < config.minHeight {
		minHeight = config.minHeight
	}

	// run the XBlock stream for this chain
	go p.runStreamer(ctx, config, minHeight, callback)

	return nil
}

// startStreamer creates a new XBlock streamer for the given chain and kicks tarts its operation.
func (p *Provider) runStreamer(ctx context.Context,
	config *ChainConfig,
	minHeight uint64,
	callback xchain.ProviderCallback,
) {
	// instantiate a new streamer for this chain
	streamer, err := NewStreamer(ctx, config, minHeight, callback, p.quitC)
	if err != nil {
		log.Error(ctx, "Could not subscribe to chain", err,
			"chain name", config.name,
			"chain id", config.id,
			"rpcURL", config.rpcURL)

		return
	}

	// start the streaming process
	streamer.start(ctx)
}

func (p *Provider) getConfig(chainID uint64) (*ChainConfig, error) {
	// check if the config for this chain ID is present
	for _, config := range p.config {
		if config.id == chainID {
			return config, nil
		}
	}

	return nil, ErrConfigNotPresent
}
