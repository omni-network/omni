package types

import (
	"time"

	"github.com/omni-network/omni/lib/errors"
)

//nolint:gochecknoglobals // Static mappings
var (
	ChainOmniEVM = EVMChain{
		Name:            "omni_evm",
		ID:              1,
		BlockPeriod:     time.Second,
		CommitmentLevel: "finalized",
	}

	chainArbGoerli = EVMChain{
		Name:            "arb_goerli",
		ID:              421613,
		IsPublic:        true,
		BlockPeriod:     6 * time.Second,
		CommitmentLevel: "finalized",
	}

	chainGoerli = EVMChain{
		Name:            "goerli",
		ID:              5,
		IsPublic:        true,
		BlockPeriod:     15 * time.Second,
		CommitmentLevel: "finalized",
	}
)

const anvilChainIDFactor = 100

// AnvilChainsByNames returns the Anvil evm chain definitions by names.
func AnvilChainsByNames(names []string) []EVMChain {
	var chains []EVMChain
	for i, name := range names {
		chains = append(chains, EVMChain{
			Name:            name,
			ID:              anvilChainIDFactor * uint64(i+1),
			BlockPeriod:     time.Second,
			CommitmentLevel: "latest",
		})
	}

	return chains
}

// PublicChainByName returns the public chain definition by name.
func PublicChainByName(name string) (EVMChain, error) {
	switch name {
	case chainArbGoerli.Name:
		return chainArbGoerli, nil
	case chainGoerli.Name:
		return chainGoerli, nil
	default:
		return EVMChain{}, errors.New("unknown chain name")
	}
}

// PublicRPCByName returns the public chain RPC address by name.
func PublicRPCByName(name string) string {
	switch name {
	case chainArbGoerli.Name:
		return "https://arbitrum-goerli.publicnode.com"
	case chainGoerli.Name:
		return "https://rpc.ankr.com/eth_goerli"
	default:
		return ""
	}
}
