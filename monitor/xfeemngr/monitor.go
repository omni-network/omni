package xfeemngr

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
)

// startMonitoring starts the monitoring goroutines.
func startMonitoring(ctx context.Context, network netconf.Network, ethClients map[uint64]ethclient.Client) {
	tick := ticker.New(ticker.WithInterval(30 * time.Second))

	for _, chain := range network.EVMChains() {
		once := func(ctx context.Context) {
			if err := monitorPortalBalanceOnce(ctx, chain, ethClients[chain.ID]); err != nil {
				log.Error(ctx, "Error monitoring portal balance, will retry", err, "chain", chain.Name)
			}
		}

		tick.Go(ctx, once)
	}
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
