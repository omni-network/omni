package netman

import (
	"context"
	"crypto/ecdsa"

	"github.com/omni-network/omni/lib/netconf"
)

// defaultDevnet returns the default e2e devnetManager network configuration.
// The RPC urls are for connecting from the host (outside docker).
// See writeNetworkConfig for the docker networking overrides.
func defaultDevnet() netconf.Network {
	return netconf.Network{
		Name: netconf.Devnet,
		Chains: []netconf.Chain{
			{
				ID:            1, // From static/geth_genesis.json
				Name:          "omni_evm",
				RPCURL:        "", // Populated by infra provider
				AuthRPCURL:    "", // Populated by infra provider
				PortalAddress: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
				IsOmni:        true,
			},
			{
				ID:            100, // From docker/compose.yaml.tmpl
				Name:          "chain_a",
				RPCURL:        "", // Populated by infra provider
				PortalAddress: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
			},
		},
	}
}

type devnetManager struct {
	network netconf.Network
	portals map[uint64]Portal
}

func (m *devnetManager) DeployPublicPortals(context.Context) error {
	// No public chains in devnetManager.
	return nil
}

func (m *devnetManager) DeployPrivatePortals(ctx context.Context) error {
	portals, err := deployPrivateContracts(ctx, m.network, privKey0)
	if err != nil {
		return err
	}

	m.portals = portals

	return nil
}

func (m *devnetManager) Network() netconf.Network {
	return m.network
}

func (m *devnetManager) RelayerKey() (*ecdsa.PrivateKey, error) {
	return privKey1, nil
}

func (m *devnetManager) Portals() map[uint64]Portal {
	return m.portals
}

func (m *devnetManager) AdditionalService() []string {
	return additionalServices(m.network)
}
