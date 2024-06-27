package types

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

//nolint:gochecknoglobals // Static mappings
var (
	allShards = []xchain.ShardID{xchain.ShardFinalized0, xchain.ShardLatest0} // All EVM chains support both finalized and latest shards.

	chainEthereum = EVMChain{
		Metadata: mustMetadata(evmchain.IDEthereum),
		IsPublic: true,
		Shards:   allShards,
	}

	chainHolesky = EVMChain{
		Metadata: mustMetadata(evmchain.IDHolesky),
		IsPublic: true,
		Shards:   allShards,
	}

	chainArbSepolia = EVMChain{
		Metadata: mustMetadata(evmchain.IDArbSepolia),
		IsPublic: true,
		Shards:   allShards,
	}

	chainOpSepolia = EVMChain{
		Metadata: mustMetadata(evmchain.IDOpSepolia),
		IsPublic: true,
		Shards:   allShards,
	}

	chainBaseSepolia = EVMChain{
		Metadata: mustMetadata(evmchain.IDBaseSepolia),
		IsPublic: true,
		Shards:   allShards,
	}
)

// OmniEVMByNetwork returns the Omni evm chain definition by netconf network.
func OmniEVMByNetwork(network netconf.ID) EVMChain {
	shards := []xchain.ShardID{xchain.ShardFinalized0} // OmniEVM has instant finality, so Latest == Finalized.
	if network.IsEphemeral() {
		shards = allShards // Enable all shards for testing.
	}

	chainID := network.Static().OmniExecutionChainID

	return EVMChain{
		Metadata: mustMetadata(chainID),
		Shards:   shards,
	}
}

// AnvilChainsByNames returns the Anvil evm chain definitions by names.
func AnvilChainsByNames(names []string) ([]EVMChain, error) {
	var chains []EVMChain
	for _, name := range names {
		meta, ok := evmchain.MetadataByName(name)
		if !ok {
			return nil, errors.New("unknown anvil chain", "name", name)
		}
		chains = append(chains, EVMChain{
			Metadata: meta,
			Shards:   allShards,
		})
	}

	return chains, nil
}

// PublicChainByName returns the public chain definition by name.
func PublicChainByName(name string) (EVMChain, error) {
	switch name {
	case chainHolesky.Name:
		return chainHolesky, nil
	case chainArbSepolia.Name:
		return chainArbSepolia, nil
	case chainOpSepolia.Name:
		return chainOpSepolia, nil
	case chainBaseSepolia.Name:
		return chainBaseSepolia, nil
	case chainEthereum.Name:
		return chainEthereum, nil
	default:
		return EVMChain{}, errors.New("unknown chain name")
	}
}

// PublicRPCByName returns the public chain RPC address by name.
func PublicRPCByName(name string) string {
	switch name {
	case chainHolesky.Name:
		return "https://ethereum-holesky-rpc.publicnode.com"
	case chainArbSepolia.Name:
		return "https://sepolia-rollup.arbitrum.io/rpc"
	case chainOpSepolia.Name:
		return "https://sepolia.optimism.io"
	case chainBaseSepolia.Name:
		return "https://sepolia.base.org"
	case chainEthereum.Name:
		return "https://ethereum-rpc.publicnode.com"
	default:
		return ""
	}
}

func mustMetadata(chainID uint64) evmchain.Metadata {
	meta, ok := evmchain.MetadataByID(chainID)
	if !ok {
		panic("unknown chain ID")
	}

	return meta
}
