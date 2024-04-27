package types

import (
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
)

//nolint:gochecknoglobals // Static mappings
var (
	chainOmniEVM = EVMChain{
		Name: "omni_evm",
		// ID:  // Depends on netconf.Static.
		BlockPeriod:       2 * time.Second, // TODO(corver): Make this more robust.
		FinalizationStrat: netconf.StratFinalized,
	}

	chainEthereum = EVMChain{
		Name:              "ethereum",
		ID:                evmchain.IDEthereum,
		IsPublic:          true,
		BlockPeriod:       12 * time.Second,
		FinalizationStrat: netconf.StratFinalized,
	}

	chainHolesky = EVMChain{
		Name:              "holesky",
		ID:                evmchain.IDHolesky,
		IsPublic:          true,
		BlockPeriod:       12 * time.Second,
		FinalizationStrat: netconf.StratFinalized,
	}

	chainArbSepolia = EVMChain{
		Name:              "arb_sepolia",
		ID:                evmchain.IDArbSepolia,
		IsPublic:          true,
		BlockPeriod:       300 * time.Millisecond,
		FinalizationStrat: netconf.StratSafe,
	}

	chainOpSepolia = EVMChain{
		Name:              "op_sepolia",
		ID:                evmchain.IDOpSepolia,
		IsPublic:          true,
		BlockPeriod:       2 * time.Second,
		FinalizationStrat: netconf.StratSafe,
	}
)

// OmniEVMByNetwork returns the Omni evm chain definition by netconf network.
func OmniEVMByNetwork(network netconf.ID) EVMChain {
	resp := chainOmniEVM
	resp.ID = network.Static().OmniExecutionChainID

	return resp
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
			Name:              name,
			ID:                meta.ChainID,
			BlockPeriod:       meta.BlockPeriod,
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
