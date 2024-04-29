package types

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
)

//nolint:gochecknoglobals // Static mappings
var (
	chainEthereum = EVMChain{
		Metadata:          mustMetadata(evmchain.IDEthereum),
		IsPublic:          true,
		FinalizationStrat: netconf.StratFinalized,
	}

	chainHolesky = EVMChain{
		Metadata:          mustMetadata(evmchain.IDHolesky),
		IsPublic:          true,
		FinalizationStrat: netconf.StratFinalized,
	}

	chainArbSepolia = EVMChain{
		Metadata:          mustMetadata(evmchain.IDArbSepolia),
		IsPublic:          true,
		FinalizationStrat: netconf.StratSafe,
	}

	chainOpSepolia = EVMChain{
		Metadata:          mustMetadata(evmchain.IDOpSepolia),
		IsPublic:          true,
		FinalizationStrat: netconf.StratSafe,
	}
)

// OmniEVMByNetwork returns the Omni evm chain definition by netconf network.
func OmniEVMByNetwork(network netconf.ID) EVMChain {
	return EVMChain{
		Metadata:          mustMetadata(network.Static().OmniExecutionChainID),
		FinalizationStrat: netconf.StratFinalized,
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
			Metadata:          meta,
			FinalizationStrat: netconf.StratLatest, // anvil doesn't support finalized
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
