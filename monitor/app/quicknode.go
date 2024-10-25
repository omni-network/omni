package monitor

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

const (
	quicknodePublicRPC = "https://damp-wandering-gadget.omni-omega.quiknode.pro"
)

func monitorQuicknodePublicRPCForever(
	ctx context.Context,
	network netconf.Network,
	ethClients map[uint64]ethclient.Client,
) {
	if network.ID != netconf.Omega {
		return // No other network has a quicknode RPC yet
	}

	omniChain, exist := network.OmniEVMChain()
	if !exist {
		return
	}

	quicknodeRPC, err := ethclient.Dial(omniChain.Name, quicknodePublicRPC)
	if err != nil {
		log.Error(ctx, "Failed to dial into quicknode public rpc", err)
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
			err := monitorQuicknodePublicRPCOnce(ctx, omniNodeRPC, quicknodeRPC)
			if err != nil {
				log.Warn(ctx, "Failed monitoring quicknode RPC endpoint (will retry)", err)
			}
		}
	}
}

func monitorQuicknodePublicRPCOnce(ctx context.Context, omniNodeRPC, quicknodeRPC ethclient.Client) error {
	omniNodeProgress, err := omniNodeRPC.SyncProgress(ctx)
	if err != nil {
		return errors.Wrap(err, "omni node sync progress")
	}

	quicknodeProgress, err := quicknodeRPC.SyncProgress(ctx)
	if err != nil {
		return errors.Wrap(err, "quicknode sync progress")
	}

	heightDiff := float64(omniNodeProgress.HighestBlock) - float64(quicknodeProgress.HighestBlock)
	quicknodeRPCSyncDiff.Set(heightDiff)

	return nil
}
