package netconf

const (
	// Simnet is a simulated network for very simple testing of individual binaries.
	// It is a single binary with mocked clients (no networking).
	Simnet = "simnet" // Single binary with mocked clients (no network)

	// Devnet is the most basic single-machine deployment of the Omni cross chain protocol.
	// It uses docker compose to setup a network with multi containers.
	// E.g. 2 geth nodes, 4 halo validators, a relayer, and 2 anvil rollups.
	Devnet = "devnet"

	// Staging is the Omni team's internal staging network, similar to a internal testnet.
	// It connects to real public rollup testnets (e.g. Arbitrum testnet).
	// It is deployed to GCP using terraform.
	// E.g. 1 Erigon, 1 Geth, 4 halo validators, 2 halo sentries, 1 relayer.
	Staging = "staging"

	// Testnet is the Omni public testnet.
	Testnet = "testnet"

	// Mainnet is the Omni public mainnet.
	Mainnet = "mainnet"
)

// supported is a map of supported networks.
//
//nolint:gochecknoglobals // Global state here is fine.
var supported = map[string]bool{
	Simnet:  true,
	Devnet:  true,
	Staging: true,
	Testnet: false,
	Mainnet: false,
}
