// Package provider is the implementation of the Provider interface.
// it abstracts connecting to different rollup chains and collecting
// XMsgs and XReceipts to construct XBlock and deliver them to the
// calling module.
package provider

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/ethclient"
)

var _ xchain.Provider = (*Provider)(nil)

// Provider stores the source chain configuration and the global quit channel.
type Provider struct {
	network     netconf.Network
	rpcClients  map[uint64]*ethclient.Client // store config for every chain ID
	backoffFunc func(context.Context) (func(), func())
}

// New instantiates the provider instance which will be ready to accept
// subscriptions for respective destination XBlocks.
func New(network netconf.Network, rpcClients map[uint64]*ethclient.Client) *Provider {
	backoffFunc := func(ctx context.Context) (func(), func()) {
		return expbackoff.NewWithReset(ctx, expbackoff.WithFastConfig())
	}

	return &Provider{
		network:     network,
		rpcClients:  rpcClients,
		backoffFunc: backoffFunc,
	}
}

// Subscribe to the XBlock from a given destination chain.
func (p *Provider) Subscribe(
	ctx context.Context,
	chainID uint64,
	fromHeight uint64,
	callback xchain.ProviderCallback,
) error {
	// retrieve the respective config
	chain, _, err := p.getChain(chainID)
	if err != nil {
		return err
	}

	// Start streaming from chain's deploy height as per config.
	if fromHeight < chain.DeployHeight {
		fromHeight = chain.DeployHeight
	}

	ctx = log.WithCtx(ctx, "chain_name", chain.Name)
	log.Info(ctx, "Subscribing to provider", "from_height", fromHeight)

	// run the XBlock stream for this chain
	p.streamBlocks(ctx, chain.Name, chainID, fromHeight, callback)

	return nil
}

// getChain provides the configuration of the given chainID.
func (p *Provider) getChain(chainID uint64) (netconf.Chain, *ethclient.Client, error) {
	chain, ok := p.network.Chain(chainID)
	if !ok {
		return netconf.Chain{}, nil, errors.New("unknown chain ID for network")
	}

	client, ok := p.rpcClients[chainID]
	if !ok {
		return netconf.Chain{}, nil, errors.New("no rpc client for chain ID")
	}

	return chain, client, nil
}
