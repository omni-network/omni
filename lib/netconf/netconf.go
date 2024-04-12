// Package netconf provides the configuration of an Omni network, an instance
// of the Omni cross chain protocol.
package netconf

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

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
		if chain.IsOmniConsensus {
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
		if chain.IsOmniEVM {
			return chain, true
		}
	}

	return Chain{}, false
}

// OmniConsensusChain returns the Omni consensus chain config or false if it does not exist.
func (n Network) OmniConsensusChain() (Chain, bool) {
	for _, chain := range n.Chains {
		if chain.IsOmniConsensus {
			return chain, true
		}
	}

	return Chain{}, false
}

// EthereumChain returns the Eth Layer1 chain config or false if it does not exist.
func (n Network) EthereumChain() (Chain, bool) {
	for _, chain := range n.Chains {
		if chain.IsEthereum {
			return chain, true
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
	ID                uint64            // Chain ID asa per https://chainlist.org
	Name              string            // Chain name as per https://chainlist.org
	RPCURL            string            // RPC URL of the chain
	AuthRPCURL        string            // RPC URL of the chain with JWT authentication enabled
	PortalAddress     common.Address    // Address of the omni portal contract on the chain
	DeployHeight      uint64            // Height that the portal contracts were deployed
	IsOmniEVM         bool              // Whether this is the Omni EVM chain
	IsOmniConsensus   bool              // Whether this is the Omni consensus chain
	IsEthereum        bool              // Whether this is the ethereum layer1 chain
	BlockPeriod       time.Duration     // Block period of the chain
	FinalizationStrat FinalizationStrat // Finalization strategy of the chain
	AVSContractAddr   common.Address    // Address of Omni AVS contracts for the chain
}

// Load loads the network configuration from the given path.
func Load(path string) (Network, error) {
	bz, err := os.ReadFile(path)
	if err != nil {
		return Network{}, errors.Wrap(err, "read network config file")
	}

	var net Network
	if err := json.Unmarshal(bz, &net); err != nil {
		return Network{}, errors.Wrap(err, "unmarshal network config file")
	}

	return net, nil
}

// Save saves the network configuration to the given path.
func Save(ctx context.Context, network Network, path string) error {
	for _, chain := range network.Chains {
		if chain.IsOmniConsensus {
			continue
		}
		if chain.PortalAddress == (common.Address{}) {
			log.Warn(ctx, "Netconf network.json portal address empty", nil, "chain", chain.Name, "path", path)
		}
	}

	bz, err := json.MarshalIndent(network, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal network config file")
	}

	if err := os.WriteFile(path, bz, 0o600); err != nil {
		return errors.Wrap(err, "write network config file")
	}

	return nil
}

type chainJSON struct {
	ID                uint64            `json:"id"`
	Name              string            `json:"name"`
	RPCURL            string            `json:"rpcurl"`
	AuthRPCURL        string            `json:"auth_rpcurl,omitempty"`
	PortalAddress     string            `json:"portal_address"`
	DeployHeight      uint64            `json:"deploy_height"`
	IsOmniEVM         bool              `json:"is_omni_evm,omitempty"`
	IsOmniConsensus   bool              `json:"is_omni_consensus,omitempty"`
	IsEthereum        bool              `json:"is_ethereum,omitempty"`
	BlockPeriod       string            `json:"block_period"`
	FinalizationStrat FinalizationStrat `json:"finalization_start"`
	AVSContractAddr   string            `json:"avs_contract_address,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (c *Chain) UnmarshalJSON(bz []byte) error {
	var cj chainJSON
	if err := json.Unmarshal(bz, &cj); err != nil {
		return errors.Wrap(err, "unmarshal chain")
	}

	blockPeriod, err := time.ParseDuration(cj.BlockPeriod)
	if err != nil {
		return errors.Wrap(err, "parse block period")
	}

	var avsAddr common.Address
	if cj.AVSContractAddr != "" {
		avsAddr = common.HexToAddress(cj.AVSContractAddr)
	}

	var portalAddr common.Address
	if cj.PortalAddress != "" {
		portalAddr = common.HexToAddress(cj.PortalAddress)
	}

	*c = Chain{
		ID:                cj.ID,
		Name:              cj.Name,
		RPCURL:            cj.RPCURL,
		AuthRPCURL:        cj.AuthRPCURL,
		PortalAddress:     portalAddr,
		DeployHeight:      cj.DeployHeight,
		IsOmniEVM:         cj.IsOmniEVM,
		IsOmniConsensus:   cj.IsOmniConsensus,
		IsEthereum:        cj.IsEthereum,
		BlockPeriod:       blockPeriod,
		FinalizationStrat: cj.FinalizationStrat,
		AVSContractAddr:   avsAddr,
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (c Chain) MarshalJSON() ([]byte, error) {
	portalAddr := c.PortalAddress.Hex()
	if c.PortalAddress == (common.Address{}) {
		portalAddr = ""
	}
	avsAddr := c.AVSContractAddr.Hex()
	if c.AVSContractAddr == (common.Address{}) {
		avsAddr = ""
	}

	cj := chainJSON{
		ID:                c.ID,
		Name:              c.Name,
		RPCURL:            c.RPCURL,
		AuthRPCURL:        c.AuthRPCURL,
		PortalAddress:     portalAddr,
		DeployHeight:      c.DeployHeight,
		IsOmniEVM:         c.IsOmniEVM,
		IsOmniConsensus:   c.IsOmniConsensus,
		IsEthereum:        c.IsEthereum,
		BlockPeriod:       c.BlockPeriod.String(),
		FinalizationStrat: c.FinalizationStrat,
		AVSContractAddr:   avsAddr,
	}

	bz, err := json.Marshal(cj)
	if err != nil {
		return nil, errors.Wrap(err, "marshal chain")
	}

	return bz, nil
}
