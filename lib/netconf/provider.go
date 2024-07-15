package netconf

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// AwaitOnChain blocks and returns network configuration as soon as it can be loaded from the on-chain registry.
// It only returns an error if the context is canceled.
func AwaitOnChain(ctx context.Context, netID ID, portalRegistry *bindings.PortalRegistry, expected []string) (Network, error) {
	if netID == Simnet {
		return SimnetNetwork(), nil
	}

	cfg := expbackoff.DefaultConfig
	cfg.MaxDelay = 5 * time.Second
	backoff := expbackoff.New(ctx, expbackoff.With(cfg))

	for {
		if ctx.Err() != nil {
			return Network{}, errors.Wrap(ctx.Err(), "provider timeout")
		}

		portals, err := portalRegistry.List(&bind.CallOpts{Context: ctx})
		if err != nil {
			log.Warn(ctx, "Failed fetching xchain registry from omni_evm (will retry)", err)
			backoff()

			continue
		}

		network, err := networkFromPortals(ctx, netID, portals)
		if err != nil {
			return Network{}, err
		}

		if !containsAll(network, expected) {
			log.Info(ctx, "XChain registry doesn't contain all expected chains (will retry)", ""+
				"expected", expected, "actual", network.ChainNamesByIDs())
			backoff()

			continue
		}

		if err := network.Verify(); err != nil {
			return Network{}, errors.Wrap(err, "invalid network configuration")
		}

		log.Info(ctx, "XChain network configuration initialized from on-chain registry", "chains", network.ChainNamesByIDs())

		return network, nil
	}
}

// containsAll returns true if the network contains the all expected chains (by name or ID).
func containsAll(network Network, expected []string) bool {
	want := make(map[string]struct{}, len(expected))
	for _, name := range expected {
		want[name] = struct{}{}
	}

	for _, chain := range network.Chains {
		delete(want, chain.Name)
		delete(want, strconv.FormatUint(chain.ID, 10))
	}

	return len(want) == 0
}

func networkFromPortals(ctx context.Context, network ID, portals []bindings.PortalRegistryDeployment) (Network, error) {
	var chains []Chain
	for _, portal := range portals {
		// Ephemeral networks may contain mock portals for testing purposes, just ignore them.
		if _, ok := evmchain.MetadataByID(portal.ChainId); !ok && network.IsEphemeral() {
			log.Warn(ctx, "Ignoring epheral network mock portal", nil, "chain_id", portal.ChainId)
			continue
		}

		// PortalRegistry garuntees BlockPeriod <= int64 max, but we check here to be safe.
		blockPeriod := time.Duration(int64(portal.BlockPeriod))
		if blockPeriod < 0 {
			return Network{}, errors.New("block period overflow", "period", portal.BlockPeriod)
		}

		chains = append(chains, Chain{
			ID:             portal.ChainId,
			Name:           portal.Name,
			PortalAddress:  portal.Addr,
			DeployHeight:   portal.DeployHeight,
			BlockPeriod:    blockPeriod,
			Shards:         toShardIDs(portal.Shards),
			AttestInterval: portal.AttestInterval,
		})
	}

	// Add omni consensus chain
	chains = append(chains, network.Static().OmniConsensusChain())

	return Network{
		ID:     network,
		Chains: chains,
	}, nil
}

func MetadataByID(network ID, chainID uint64) evmchain.Metadata {
	if IsOmniConsensus(network, chainID) {
		chain := network.Static().OmniConsensusChain()
		return evmchain.Metadata{
			ChainID:     chain.ID,
			Name:        chain.Name,
			BlockPeriod: chain.BlockPeriod,
		}
	}

	if meta, ok := evmchain.MetadataByID(chainID); ok { // If well-known chain, use it.
		return meta
	}

	return evmchain.Metadata{
		ChainID:     chainID,
		Name:        fmt.Sprintf("unknown_%d", chainID),
		BlockPeriod: time.Second * 2,
	}
}

func MetadataByName(network ID, name string) (evmchain.Metadata, bool) {
	// there are multiple omni EVM metadatas with the same name,
	// so we need go get by chainID of this network's omniEVM chain
	if evmchain.IsOmniEVM(name) {
		return evmchain.MetadataByID(network.Static().OmniExecutionChainID)
	}

	return evmchain.MetadataByName(name)
}

func ChainNamer(network ID) func(uint64) string {
	return func(chainID uint64) string {
		return MetadataByID(network, chainID).Name
	}
}

func ChainVersionNamer(network ID) func(version xchain.ChainVersion) string {
	return func(chainVer xchain.ChainVersion) string {
		return fmt.Sprintf("%s|%s", MetadataByID(network, chainVer.ID).Name, chainVer.ConfLevel.Label())
	}
}

func toShardIDs(shards []uint64) []xchain.ShardID {
	var resp []xchain.ShardID
	for _, shard := range shards {
		resp = append(resp, xchain.ShardID(shard))
	}

	return resp
}
