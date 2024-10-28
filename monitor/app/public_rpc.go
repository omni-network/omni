package monitor

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

const omegaPublicRPC = "https://damp-wandering-gadget.omni-omega.quiknode.pro"

func monitorPublicRPCForever(
	ctx context.Context,
	network netconf.Network,
	ethClients map[uint64]ethclient.Client,
) {
	if network.ID != netconf.Omega {
		return // no public URL exists
	}

	omniChain, exist := network.OmniEVMChain()
	if !exist {
		return
	}

	log.Info(ctx, "Setting up monitoring of a public RPC for %v", network.ID)

	publicRPC, err := ethclient.Dial(omniChain.Name, omegaPublicRPC)
	if err != nil {
		log.Error(ctx, "Failed to dial into public RPC", err)
		return
	}

	omniNodeRPC := ethClients[omniChain.ID]

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

func monitorPublicRPCOnce(ctx context.Context, omniNodeRPC, publicRPC ethclient.Client) error {
	omniNodeProgress, err := omniNodeRPC.SyncProgress(ctx)
	if err != nil {
		return errors.Wrap(err, "omni node sync progress")
	}

	publicRPCProgress, err := publicRPC.SyncProgress(ctx)
	if err != nil {
		return errors.Wrap(err, "public RPC sync progress")
	}

	heightDiff := float64(omniNodeProgress.HighestBlock) - float64(publicRPCProgress.HighestBlock)
	publicRPCSyncDiff.Set(heightDiff)

	return nil
}
