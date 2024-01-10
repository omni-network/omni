// Package provider is the implementation of the Provider interface.
// it abstracts connecting to different rollup chains and collecting
// XMsgs and XReceipts to construct XBlock and deliver them to the
// calling module.
package provider

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	prov "github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrChainAlreadySubscribed = errors.New("given destination chain id is already subscribed")
	ErrConfigNotPresent       = errors.New("config for chain id is not found")
)

// ChainConfig is the configuration parameters for all the chains
// that needs to be managed by the provider.
type ChainConfig struct {
	name          string
	id            uint64
	portalAddress common.Address
	rpcURL        string
	fromHeight    uint64
	callback      prov.ProviderCallback
}

// Provider stores the source chain configuration and
// other subscribed streams for the record.
type Provider struct {
	config        map[uint64]*ChainConfig // store config for every chain ID
	subscriptions map[uint64]*Streamer    // store the streamer for every chain Id
}

// New instantiates the provider instance which will be ready to accept
// subscriptions for respective destination XBlocks.
func New(chains []*ChainConfig) *Provider {
	chainConfigs := make(map[uint64]*ChainConfig)
	for _, chain := range chains {
		chainConfigs[chain.id] = chain
	}

	return &Provider{
		config:        chainConfigs,
		subscriptions: make(map[uint64]*Streamer),
	}
}

// Subscribe to the XBlock from a given destination chain.
func (p *Provider) Subscribe(
	ctx context.Context,
	chainID uint64,
	fromHeight uint64,
	callback prov.ProviderCallback,
) error {
	log.Info(ctx, "Subscribing to provider ", "id", chainID, "fromHeight", fromHeight)

	// check if configuration is present for this chain
	config, err := p.getConfig(chainID)
	if err != nil {
		return err
	}

	// check if the destination chain is already subscribed
	if p.isAlreadySubscribed(chainID) {
		return ErrChainAlreadySubscribed
	}

	// add few more configs obtained during subscription
	config.fromHeight = fromHeight
	config.callback = callback

	// start a XBlock stream for this chain
	go p.startStreamer(ctx, config)

	return nil
}

// Stop kills all the subscriptions.
func (p *Provider) Stop(ctx context.Context) {
	for _, streamer := range p.subscriptions {
		log.Info(ctx, "Stopping subscription to provider",
			"chain name", streamer.chainConfig.id,
			"chain id", streamer.chainConfig.name,
			"last xblock delivered", streamer.lastBlockHeight)
		streamer.stop(ctx)
	}
}

// startStreamer creates a new XBlock streamer for the given chain and kicks tarts its operation.
func (p *Provider) startStreamer(ctx context.Context, config *ChainConfig) {
	// instantiate a new streamer for this chain
	streamer, err := NewStreamer(ctx, config)
	if err != nil {
		log.Error(ctx, "Could not subscribe to chain", err,
			"chain name", config.name,
			"chain id", config.id,
			"rpcURL", config.rpcURL)

		return
	}

	// add the streamer to the active subscriptions
	p.subscriptions[streamer.chainConfig.id] = streamer

	// start the streaming process
	streamer.start(ctx)
}

func (p *Provider) isAlreadySubscribed(chainID uint64) bool {
	// check if this chain ID is already subscribed
	if _, ok := p.subscriptions[chainID]; ok {
		return false
	}

	return true
}

func (p *Provider) getConfig(chainID uint64) (*ChainConfig, error) {
	// check if the config for this chain ID is present
	if config, ok := p.config[chainID]; ok {
		return config, nil
	}

	return nil, ErrConfigNotPresent
}
