package keeper

import (
	"context"
	"sort"
	"sync"

	"github.com/omni-network/omni/halo/registry/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.PortalRegistry = Keeper{}

// cache caches the latest network portals to mitigate DB lookups.
type cache struct {
	sync.RWMutex
	network *Network
}

func (c *cache) Set(network *Network) {
	c.Lock()
	c.network = network
	c.Unlock()
}

func (c *cache) Get() (*Network, bool) {
	c.RLock()
	defer c.RUnlock()

	return c.network, c.network != nil
}

func (k Keeper) updateNetwork(ctx context.Context, network *Network) error {
	if err := k.networkTable.Update(ctx, network); err != nil {
		return errors.Wrap(err, "update network")
	}

	k.latestCache.Set(network)

	return nil
}

func (k Keeper) ConfLevels(ctx context.Context) (map[uint64][]xchain.ConfLevel, error) {
	portals, err := k.getLatestPortals(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get latest portals")
	}

	resp := make(map[uint64][]xchain.ConfLevel, len(portals)+1)
	for _, portal := range portals {
		uniqLevels := make(map[xchain.ConfLevel]struct{})
		for _, shardID := range portal.GetShardIds() {
			uniqLevels[xchain.ShardID(shardID).ConfLevel()] = struct{}{}
		}

		var confLevels []xchain.ConfLevel
		for confLevel := range uniqLevels {
			confLevels = append(confLevels, confLevel)
		}

		sort.Slice(confLevels, func(i, j int) bool {
			return confLevels[i] < confLevels[j]
		})

		resp[portal.GetChainId()] = confLevels
	}

	// Add the consensus chain shard/conf level; only ConfFinalized
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	consensusID, err := netconf.ConsensusChainIDStr2Uint64(sdkCtx.ChainID())
	if err != nil {
		return nil, errors.Wrap(err, "parse chain id")
	}

	resp[consensusID] = []xchain.ConfLevel{xchain.ConfFinalized}

	return resp, nil
}

func (k Keeper) SupportedChain(ctx context.Context, chainID uint64) (bool, error) {
	portals, err := k.getLatestPortals(ctx)
	if err != nil {
		return false, errors.Wrap(err, "get latest portals")
	}

	for _, p := range portals {
		if p.GetChainId() == chainID {
			return true, nil
		}
	}

	// Always allow the consensus chain
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	consensusID, err := netconf.ConsensusChainIDStr2Uint64(sdkCtx.ChainID())
	if err != nil {
		return false, errors.Wrap(err, "parse chain id")
	} else if consensusID == chainID {
		return true, nil
	}

	return false, nil
}

func (k Keeper) getLatestPortals(ctx context.Context) ([]*Portal, error) {
	if network, ok := k.latestCache.Get(); ok {
		return network.GetPortals(), nil
	}

	latestNetworkID, err := k.networkTable.LastInsertedSequence(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get last network ID")
	} else if latestNetworkID == 0 {
		// No network exists yet, return empty list
		return nil, nil
	}

	lastNetwork, err := k.networkTable.Get(ctx, latestNetworkID)
	if err != nil {
		return nil, errors.Wrap(err, "get network")
	}

	k.latestCache.Set(lastNetwork)

	return lastNetwork.GetPortals(), nil
}
