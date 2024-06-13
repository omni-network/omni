package netconf

import (
	"github.com/omni-network/omni/lib/errors"
)

// ID is a network identifier.
type ID string

// IsEphemeral returns true if the network is short-lived, therefore ephemeral.
func (i ID) IsEphemeral() bool {
	return i == Simnet || i == Devnet || i == Staging
}

// IsProtected returns true if the network is long-lived, therefore protected.
func (i ID) IsProtected() bool {
	return !i.IsEphemeral()
}

// Static returns the static config and data for the network.
func (i ID) Static() Static {
	return statics[i]
}

func (i ID) Verify() error {
	if !supported[i] {
		return errors.New("unsupported network", "network", i)
	}

	return nil
}

func (i ID) String() string {
	return string(i)
}

func (i ID) Version() string {
	return i.Static().Version
}

const (
	// Simnet is a simulated network for very simple testing of individual binaries.
	// It is a single binary with mocked clients (no networking).
	Simnet ID = "simnet" // Single binary with mocked clients (no network)

	// Devnet is the most basic single-machine deployment of the Omni cross chain protocol.
	// It uses docker compose to setup a network with multi containers.
	// E.g. 2 geth nodes, 4 halo validators, a relayer, and 2 anvil rollups.
	Devnet ID = "devnet"

	// Staging is the Omni team's internal staging network, similar to a internal testnet.
	// It connects to real public rollup testnets (e.g. Arbitrum testnet).
	// It is deployed to GCP using terraform.
	// E.g. 1 Erigon, 1 Geth, 4 halo validators, 2 halo sentries, 1 relayer.
	Staging ID = "staging"

	// Omega is a Omni public testnet.
	Omega ID = "omega"

	// Mainnet is the Omni public mainnet.
	Mainnet ID = "mainnet"
)

// supported is a map of supported networks.
//
//nolint:gochecknoglobals // Global state here is fine.
var supported = map[ID]bool{
	Simnet:  true,
	Devnet:  true,
	Staging: true,
	Omega:   true,
	Mainnet: true,
}

// IsAny returns true if the `ID` matches any of the provided targets.
func IsAny(id ID, targets ...ID) bool {
	for _, target := range targets {
		if id == target {
			return true
		}
	}

	return false
}

// All returns all the supported network IDs.
func All() []ID {
	var resp []ID
	for id, ok := range supported {
		if ok {
			resp = append(resp, id)
		}
	}

	return resp
}
