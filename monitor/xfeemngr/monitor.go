package xfeemngr

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
)

// startMonitoring starts the monitoring goroutines.
func startMonitoring(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints) error {
	clients, err := makeClients(network, endpoints)
	if err != nil {
		return errors.Wrap(err, "make clients")
	}

	tick := ticker.New(ticker.WithInterval(30 * time.Second))

	for _, chain := range network.EVMChains() {
		once := func(ctx context.Context) {
			if err := monitorPortalBalanceOnce(ctx, chain, clients[chain.ID]); err != nil {
				log.Error(ctx, "Error monitoring portal balance, will retry", err, "chain", chain.Name)
			}
		}

		tick.Go(ctx, once)
	}

	return nil
}

// makeClients creates a map of ethereum clients for each chain.
func makeClients(network netconf.Network, endpoints xchain.RPCEndpoints) (map[uint64]ethclient.Client, error) {
	clients := make(map[uint64]ethclient.Client)

	for _, chain := range network.EVMChains() {
		rpc, err := endpoints.ByNameOrID(chain.Name, chain.ID)
		if err != nil {
			return nil, errors.Wrap(err, "rpc endpoint")
		}

		client, err := ethclient.Dial(chain.Name, rpc)
		if err != nil {
			return nil, errors.Wrap(err, "dial client")
		}

		clients[chain.ID] = client
	}

	return clients, nil
}

// monitorPortalBalanceOnce updates the portal balance metric once.
func monitorPortalBalanceOnce(ctx context.Context, chain netconf.Chain, client ethclient.Client) error {
	balance, err := client.EtherBalanceAt(ctx, chain.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "balance at")
	}

	portalBalance.WithLabelValues(chain.Name, chain.PortalAddress.Hex()).Set(balance)

	return nil
}
