package network

import "github.com/omni-network/omni/lib/netconf"

// NewE2E returns the default e2e network configuration.
// The RPC urls are for connecting from the host (outside docker).
// See writeNetworkConfig for the docker networking overrides.
func NewE2E() netconf.Network {
	return netconf.Network{
		Name: netconf.Devnet,
		Chains: []netconf.Chain{
			{
				ID:            1, // From static/geth_genesis.json
				Name:          "omni_evm",
				RPCURL:        "http://localhost:8545",
				AuthRPCURL:    "http://localhost:8551",
				PortalAddress: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
				IsOmni:        true,
			},
			{
				ID:            100, // From docker/compose.yaml.tmpl
				Name:          "chain_a",
				RPCURL:        "http://localhost:6545",
				PortalAddress: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
			},
		},
	}
}
