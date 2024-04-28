// Package netconf provides the configuration of an Omni network, an instance
// of the Omni cross chain protocol.
package netconf

import (
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/common"
)

// Network defines a deployment of the Omni cross chain protocol.
// It spans an omni chain (both execution and consensus) and a set of
// supported rollup EVMs.
type Network struct {
	ID     ID      `json:"name"`   // ID of the network. e.g. "simnet", "testnet", "staging", "mainnet"
	Chains []Chain `json:"chains"` // Chains that are part of the network
}

// Validate returns an error if the configuration is invalid.
func (n Network) Validate() error {
	if err := n.ID.Verify(); err != nil {
		return err
	}

	// TODO(corver): Validate chains

	return nil
}

// EVMChains returns all evm chains in the network. It excludes the omni consensus chain.
func (n Network) EVMChains() []Chain {
	resp := make([]Chain, 0, len(n.Chains))
	for _, chain := range n.Chains {
		if IsOmniConsensus(n.ID, chain.ID) {
			continue
		}

		resp = append(resp, chain)
	}

	return resp
}

// ChainIDs returns the all chain IDs in the network.
// This is a convenience method.
func (n Network) ChainIDs() []uint64 {
	resp := make([]uint64, 0, len(n.Chains))
	for _, chain := range n.Chains {
		resp = append(resp, chain.ID)
	}

	return resp
}

// ChainNamesByIDs returns the all chain IDs and names in the network.
// This is a convenience method.
func (n Network) ChainNamesByIDs() map[uint64]string {
	resp := make(map[uint64]string)
	for _, chain := range n.Chains {
		resp[chain.ID] = chain.Name
	}

	return resp
}

// OmniEVMChain returns the Omni execution chain config or false if it does not exist.
func (n Network) OmniEVMChain() (Chain, bool) {
	for _, chain := range n.Chains {
		if IsOmniExecution(n.ID, chain.ID) {
			return chain, true
		}
	}

	return Chain{}, false
}

// OmniConsensusChain returns the Omni consensus chain config or false if it does not exist.
func (n Network) OmniConsensusChain() (Chain, bool) {
	for _, chain := range n.Chains {
		if IsOmniConsensus(n.ID, chain.ID) {
			return chain, true
		}
	}

	return Chain{}, false
}

// EthereumChain returns the ethereum Layer1 chain config or false if it does not exist.
func (n Network) EthereumChain() (Chain, bool) {
	for _, chain := range n.Chains {
		switch n.ID {
		case Mainnet:
			if chain.ID == evmchain.IDEthereum {
				return chain, true
			}
		case Testnet:
			if chain.ID == evmchain.IDHolesky {
				return chain, true
			}
		default:
			if chain.ID == evmchain.IDMockL1Fast || chain.ID == evmchain.IDMockL1Slow {
				return chain, true
			}
		}
	}

	return Chain{}, false
}

// ChainName returns the chain name for the given ID or an empty string if it does not exist.
func (n Network) ChainName(id uint64) string {
	chain, _ := n.Chain(id)
	return chain.Name
}

// Chain returns the chain config for the given ID or false if it does not exist.
func (n Network) Chain(id uint64) (Chain, bool) {
	for _, chain := range n.Chains {
		if chain.ID == id {
			return chain, true
		}
	}

	return Chain{}, false
}

// FinalizationStrat defines the finalization strategy of a chain.
// This is mostly ethclient.HeadFinalized, but some chains may not support
// it, like zkEVM chains which would need a much more involved strategy.
type FinalizationStrat string

func (h FinalizationStrat) Verify() error {
	if !allStrats[h] {
		return errors.New("invalid finalization strategy", "start", h)
	}

	return nil
}

func (h FinalizationStrat) String() string {
	return string(h)
}

//nolint:gochecknoglobals // Static mappings
var allStrats = map[FinalizationStrat]bool{
	StratFinalized: true,
	StratLatest:    true,
	StratSafe:      true,
}

const (
	StratFinalized = FinalizationStrat("finalized")
	StratLatest    = FinalizationStrat("latest")
	StratSafe      = FinalizationStrat("safe")
)

// Chain defines the configuration of an execution chain that supports
// the Omni cross chain protocol. This is most supported Rollup EVMs, but
// also the Omni EVM, and the Omni Consensus chain.
type Chain struct {
	ID   uint64 // Chain ID asa per https://chainlist.org
	Name string // Chain name as per https://chainlist.org
	// RPCURL            string            // RPC URL of the chain
	PortalAddress     common.Address    // Address of the omni portal contract on the chain
	DeployHeight      uint64            // Height that the portal contracts were deployed
	BlockPeriod       time.Duration     // Block period of the chain
	FinalizationStrat FinalizationStrat // Finalization strategy of the chain
}
