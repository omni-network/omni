// Package netconf provides the configuration of an Omni network, an instance
// of the Omni cross chain protocol.
package netconf

import (
	"encoding/json"
	"os"

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

// OmniChain returns the Omni execution chain config or false if it does not exist.
func (n Network) OmniChain() (Chain, bool) {
	for _, chain := range n.Chains {
		if chain.IsOmni {
			return chain, true
		}
	}

	return Chain{}, false
}

// Chain defines the configuration of an execution chain that supports
// the Omni cross chain protocol. This is most supported Rollup EVM, but
// also the Omni EVM.
type Chain struct {
	ID            uint64 `json:"id"`             // Chain ID asa per https://chainlist.org
	Name          string `json:"name"`           // Chain name as per https://chainlist.org
	RPCURL        string `json:"rpcurl"`         // RPC URL of the chain
	PortalAddress string `json:"portal_address"` // Address of the omni portal contract on the chain
	DeployHeight  uint64 `json:"deploy_height"`  // Height that the portal contracts were deployed
	IsOmni        bool   `json:"is_omni"`        // Whether this is the Omni chain
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
