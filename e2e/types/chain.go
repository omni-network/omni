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

	// Mainnet.

	chainEthereum = EVMChain{
		Metadata: mustMetadata(evmchain.IDEthereum),
		IsPublic: true,
		Shards:   allShards,
	}

	chainOptimism = EVMChain{
		Metadata: mustMetadata(evmchain.IDOptimism),
		IsPublic: true,
		Shards:   allShards,
	}

	chainBSC = EVMChain{
		Metadata: mustMetadata(evmchain.IDBSC),
		IsPublic: true,
		Shards:   []xchain.ShardID{},
	}

	chainPolygon = EVMChain{
		Metadata: mustMetadata(evmchain.IDPolygon),
		IsPublic: true,
		Shards:   []xchain.ShardID{},
	}

	chainHyperEVM = EVMChain{
		Metadata: mustMetadata(evmchain.IDHyperEVM),
		IsPublic: true,
		Shards:   []xchain.ShardID{},
	}

	chainMantle = EVMChain{
		Metadata: mustMetadata(evmchain.IDMantle),
		IsPublic: true,
		Shards:   []xchain.ShardID{},
	}

	chainBase = EVMChain{
		Metadata: mustMetadata(evmchain.IDBase),
		IsPublic: true,
		Shards:   allShards,
	}

	chainArbitrum = EVMChain{
		Metadata: mustMetadata(evmchain.IDArbitrumOne),
		IsPublic: true,
		Shards:   allShards,
	}

	chainBerachain = EVMChain{
		Metadata: mustMetadata(evmchain.IDBerachain),
		IsPublic: true,
		Shards:   []xchain.ShardID{},
	}

	chainPlume = EVMChain{
		Metadata: mustMetadata(evmchain.IDPlume),
		IsPublic: true,
		Shards:   []xchain.ShardID{},
	}

	// Testnet.

	chainBSCTestnet = EVMChain{
		Metadata: mustMetadata(evmchain.IDBSCTestnet),
		IsPublic: true,
		Shards:   []xchain.ShardID{},
	}

	chainHyperEVMTestnet = EVMChain{
		Metadata: mustMetadata(evmchain.IDHyperEVMTestnet),
		IsPublic: true,
		Shards:   []xchain.ShardID{},
	}

	chainHolesky = EVMChain{
		Metadata: mustMetadata(evmchain.IDHolesky),
		IsPublic: true,
		Shards:   allShards,
	}

	chainPolygonAmoy = EVMChain{
		Metadata: mustMetadata(evmchain.IDPolygonAmoy),
		IsPublic: true,
		Shards:   []xchain.ShardID{},
	}

	chainBaseSepolia = EVMChain{
		Metadata: mustMetadata(evmchain.IDBaseSepolia),
		IsPublic: true,
		Shards:   allShards,
	}

	chainPlumeTestnet = EVMChain{
		Metadata: mustMetadata(evmchain.IDPlumeTestnet),
		IsPublic: true,
		Shards:   []xchain.ShardID{},
	}

	chainArbSepolia = EVMChain{
		Metadata: mustMetadata(evmchain.IDArbSepolia),
		IsPublic: true,
		Shards:   allShards,
	}

	chainSepolia = EVMChain{
		Metadata: mustMetadata(evmchain.IDSepolia),
		IsPublic: true,
		Shards:   allShards,
	}

	chainOpSepolia = EVMChain{
		Metadata: mustMetadata(evmchain.IDOpSepolia),
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
	// Mainnet
	case chainEthereum.Name:
		return chainEthereum, nil
	case chainOptimism.Name:
		return chainOptimism, nil
	case chainBSC.Name:
		return chainBSC, nil
	case chainPolygon.Name:
		return chainPolygon, nil
	case chainHyperEVM.Name:
		return chainHyperEVM, nil
	case chainMantle.Name:
		return chainMantle, nil
	case chainBase.Name:
		return chainBase, nil
	case chainArbitrum.Name:
		return chainArbitrum, nil
	case chainBerachain.Name:
		return chainBerachain, nil
	case chainPlume.Name:
		return chainPlume, nil
	// Testnet
	case chainBSCTestnet.Name:
		return chainBSCTestnet, nil
	case chainHyperEVMTestnet.Name:
		return chainHyperEVMTestnet, nil
	case chainHolesky.Name:
		return chainHolesky, nil
	case chainPolygonAmoy.Name:
		return chainPolygonAmoy, nil
	case chainBaseSepolia.Name:
		return chainBaseSepolia, nil
	case chainPlumeTestnet.Name:
		return chainPlumeTestnet, nil
	case chainArbSepolia.Name:
		return chainArbSepolia, nil
	case chainSepolia.Name:
		return chainSepolia, nil
	case chainOpSepolia.Name:
		return chainOpSepolia, nil
	default:
		return EVMChain{}, errors.New("unknown chain name")
	}
}

// PublicRPCByName returns the public chain RPC address by name.
func PublicRPCByName(name string) string {
	switch name {
	// Mainnet
	case chainEthereum.Name:
		return "https://eth.merkle.io"
	case chainOptimism.Name:
		return "https://optimism-rpc.publicnode.com"
	case chainBSC.Name:
		return "https://bsc-rpc.publicnode.com"
	case chainPolygon.Name:
		return "https://polygon-bor-rpc.publicnode.com"
	case chainHyperEVM.Name:
		return "https://rpc.hyperliquid.xyz/evm"
	case chainMantle.Name:
		return "https://mantle-rpc.publicnode.com"
	case chainBase.Name:
		return "https://base-rpc.publicnode.com"
	case chainArbitrum.Name:
		return "https://arbitrum-one-rpc.publicnode.com"
	case chainBerachain.Name:
		return "https://berachain-rpc.publicnode.com"
	case chainPlume.Name:
		return "https://rpc.plume.org"
	// Testnet
	case chainBSCTestnet.Name:
		return "https://bsc-testnet-rpc.publicnode.com"
	case chainHyperEVMTestnet.Name:
		return "https://rpc.hyperliquid-testnet.xyz/evm"
	case chainHolesky.Name:
		return "https://ethereum-holesky-rpc.publicnode.com"
	case chainPolygonAmoy.Name:
		return "https://polygon-amoy-bor-rpc.publicnode.com"
	case chainBaseSepolia.Name:
		return "https://sepolia.base.org"
	case chainPlumeTestnet.Name:
		return "https://testnet-rpc.plumenetwork.xyz"
	case chainArbSepolia.Name:
		return "https://sepolia-rollup.arbitrum.io/rpc"
	case chainSepolia.Name:
		return "https://ethereum-sepolia-rpc.publicnode.com"
	case chainOpSepolia.Name:
		return "https://sepolia.optimism.io"
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
