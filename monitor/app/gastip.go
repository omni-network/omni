package monitor

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
)

// monitorOmniEVMGasTipForever monitors the suggested gas tip cap for the Omni EVM chain.
func monitorOmniEVMGasTipForever(ctx context.Context,
	network netconf.Network,
	ethClients map[uint64]ethclient.Client,
) {
	ethCl, ok := ethClients[network.ID.Static().OmniExecutionChainID]
	if !ok {
		log.Error(ctx, "No eth client for omni chain", nil)
		return
	}

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			tip, err := ethCl.SuggestGasTipCap(ctx)
			if err != nil {
				log.Warn(ctx, "Failed to get gas tip (will retry)", err)
				continue
			}

			gasTipCap.Set(umath.ToGweiF64(tip))
		}
	}
}
