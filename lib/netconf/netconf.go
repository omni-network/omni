// Package netconf provides the configuration of an Omni network, an instance
// of the Omni cross chain protocol.
package netconf

import (
	"encoding/json"
	"os"
	"time"

	"github.com/omni-network/omni/lib/errors"
)

// Network defines a deployment of the Omni cross chain protocol.
// It spans an omni chain (both execution and consensus) and a set of
// supported rollup EVMs.
type Network struct {
	Name   string  `json:"name"`   // Name of the network. e.g. "simnet", "testnet", "staging", "mainnet"
	Chains []Chain `json:"chains"` // Chains that are part of the network
}

// Validate returns an error if the configuration is invalid.
func (n Network) Validate() error {
	if !supported[n.Name] {
		return errors.New("unsupported network", "name", n.Name)
	}

	// TODO(corver): Validate chains

	return nil
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

// OmniChain returns the Omni execution chain config or false if it does not exist.
func (n Network) OmniChain() (Chain, bool) {
	for _, chain := range n.Chains {
		if chain.IsOmni {
			return chain, true
		}
	}

	return Chain{}, false
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

// Chain defines the configuration of an execution chain that supports
// the Omni cross chain protocol. This is most supported Rollup EVM, but
// also the Omni EVM.
type Chain struct {
	ID              uint64        // Chain ID asa per https://chainlist.org
	Name            string        // Chain name as per https://chainlist.org
	RPCURL          string        // RPC URL of the chain
	AuthRPCURL      string        // RPC URL of the chain with JWT authentication enabled
	PortalAddress   string        // Address of the omni portal contract on the chain
	DeployHeight    uint64        // Height that the portal contracts were deployed
	IsOmni          bool          // Whether this is the Omni chain
	BlockPeriod     time.Duration // Block period of the chain
	CommitmentLevel string        // Commitment Level of the block ("latest", "finalized")
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
func Save(network Network, path string) error {
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
	ID              uint64 `json:"id"`
	Name            string `json:"name"`
	RPCURL          string `json:"rpcurl"`
	AuthRPCURL      string `json:"auth_rpcurl,omitempty"`
	PortalAddress   string `json:"portal_address"`
	DeployHeight    uint64 `json:"deploy_height"`
	IsOmni          bool   `json:"is_omni,omitempty"`
	BlockPeriod     string `json:"block_period"`
	CommitmentLevel string `json:"commitment_level"`
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

	*c = Chain{
		ID:              cj.ID,
		Name:            cj.Name,
		RPCURL:          cj.RPCURL,
		AuthRPCURL:      cj.AuthRPCURL,
		PortalAddress:   cj.PortalAddress,
		DeployHeight:    cj.DeployHeight,
		IsOmni:          cj.IsOmni,
		BlockPeriod:     blockPeriod,
		CommitmentLevel: cj.CommitmentLevel,
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (c Chain) MarshalJSON() ([]byte, error) {
	cj := chainJSON{
		ID:              c.ID,
		Name:            c.Name,
		RPCURL:          c.RPCURL,
		AuthRPCURL:      c.AuthRPCURL,
		PortalAddress:   c.PortalAddress,
		DeployHeight:    c.DeployHeight,
		IsOmni:          c.IsOmni,
		BlockPeriod:     c.BlockPeriod.String(),
		CommitmentLevel: c.CommitmentLevel,
	}

	bz, err := json.Marshal(cj)
	if err != nil {
		return nil, errors.Wrap(err, "marshal chain")
	}

	return bz, nil
}
