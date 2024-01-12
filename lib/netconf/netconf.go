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
	Name   string  `json:"name"`   // Name of the network. e.g. "testnet", "staging", "mainnet"
	Chains []Chain `json:"chains"` // Chains that are part of the network
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
