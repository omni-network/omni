package monitor

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

func monitorPublicRPCForever(
	ctx context.Context,
	network netconf.Network,
	ethClients map[uint64]ethclient.Client,
) {
	omniChain, exist := network.OmniEVMChain()
	if !exist {
		return
	}

	if ethClients == nil {
		return
	}

	publicRPC, err := publicRPCEndpoint(ctx, network, omniChain, ethClients)
	if err != nil {
		log.Error(ctx, "Failed to dial into public RPC", err)
		return
	}
	go publicRPC.CloseIdleConnectionsForever(ctx)

	omniNodeRPC := ethClients[omniChain.ID]

	log.Info(ctx, "Monitoring public RPC", "public", publicRPC.Address(), "local", omniNodeRPC.Address())

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorPublicRPCOnce(ctx, omniNodeRPC, publicRPC)
			if err != nil {
				log.Warn(ctx, "Failed monitoring public RPC endpoint (will retry)", err)
			}
		}
	}
}

// publicRPCEndpoint returns the public RPC endpoint for the network and chain specified.
// If no public RPC is known, return a node of the chain directly.
func publicRPCEndpoint(ctx context.Context, network netconf.Network, chain netconf.Chain, ethClients map[uint64]ethclient.Client) (ethclient.Client, error) {
	urls := map[netconf.ID]string{
		netconf.Staging: "https://staging.omni.network",
		netconf.Omega:   "https://omega.omni.network",
		netconf.Mainnet: "https://mainnet.omni.network",
	}

	if url, exists := urls[network.ID]; exists {
		return ethclient.DialContext(ctx, chain.Name, url)
	}

	return ethClients[chain.ID], nil
}

func monitorPublicRPCOnce(ctx context.Context, omniNodeRPC, publicRPC ethclient.Client) error {
	omniNodeHeight, err := omniNodeRPC.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "omni node height")
	}

	publicRPCHeight, err := publicRPC.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "public RPC height")
	}

	heightDiff := float64(omniNodeHeight) - float64(publicRPCHeight)
	publicRPCSyncDiff.Set(heightDiff)

	return nil
}
