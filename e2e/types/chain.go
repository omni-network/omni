package types

import (
	"time"

	"github.com/omni-network/omni/lib/chainids"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
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

	chainHolesky = EVMChain{
		Name:               "holesky",
		ID:                 chainids.Holesky,
		IsPublic:           true,
		BlockPeriod:        15 * time.Second,
		FinalizationStrat:  netconf.StratLatest,
		AVSContractAddress: contracts.TestnetAVS(),
	}

	chainArbSepolia = EVMChain{
		Name:              "arb_sepolia",
		ID:                chainids.ArbSepolia,
		IsPublic:          true,
		BlockPeriod:       6 * time.Second,
		FinalizationStrat: netconf.StratFinalized,
	}

	chainOpSepolia = EVMChain{
		Name:              "op_sepolia",
		ID:                chainids.OpSepolia,
		IsPublic:          true,
		BlockPeriod:       6 * time.Second,
		FinalizationStrat: netconf.StratFinalized,
	}
)

const anvilChainIDFactor = 100

// OmniEVMByNetwork returns the Omni evm chain definition by netconf network.
func OmniEVMByNetwork(network netconf.ID) EVMChain {
	resp := chainOmniEVM
	resp.ID = network.Static().OmniExecutionChainID

	return resp
}

// AnvilChainsByNames returns the Anvil evm chain definitions by names.
func AnvilChainsByNames(names []string) []EVMChain {
	var chains []EVMChain
	for i, name := range names {
		chains = append(chains, EVMChain{
			Name:              name,
			ID:                anvilChainIDFactor * uint64(i+1),
			BlockPeriod:       time.Second,
			FinalizationStrat: netconf.StratLatest, // anvil doesn't support finalized
		})
	}

	return chains
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
	default:
		return ""
	}
}
