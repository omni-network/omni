package types

import (
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

// Manifest wraps e2e.Manifest with additional omni-specific fields.
type Manifest struct {
	e2e.Manifest

	Network string `toml:"network"`

	// AnvilChains defines the anvil chains to deploy; chain_a, chain_b, etc.
	AnvilChains []string `toml:"anvil_chains"`

	// PublicChains defines the public chains to connect to; arb_sepolia, etc.
	PublicChains []string `toml:"public_chains"`

	// AVSTarget identifies the chain to deploy the AVS contracts to.
	// It must be one of the anvil or public chains.
	AVSTarget string `toml:"avs_target"`

	// MultiOmniEVMs defines whether to deploy one or multiple Omni EVMs.
	MultiOmniEVMs bool `toml:"multi_omni_evms"`

	// SlowTests defines whether to run slow tests (e.g. tests/eigen_tests.go)
	SlowTests bool `toml:"slow_tests"`

	// OnlyMonitor indicates that the monitor is the only thing that we deploy in this network.
	OnlyMonitor bool `toml:"only_monitor"`

	// PingPongN defines the number of ping pong messages to send. Defaults 3 if 0.
	PingPongN uint64 `toml:"pingpong_n"`
}

// OmniEVMs returns the map names and GcMode of Omni EVMs to deploy.
// If only a single Omni EVM is to be deployed, the name is "omni_evm".
// Otherwise, the names are "<node>_evm".
func (m Manifest) OmniEVMs() map[string]GcMode {
	if !m.MultiOmniEVMs {
		return map[string]GcMode{
			"omni_evm": GcModeFull,
		}
	}

	resp := make(map[string]GcMode)
	for name, node := range m.Nodes {
		var gcmode GcMode
		switch node.Mode {
		case "full":
			gcmode = GcModeArchive
		case "seed":
			gcmode = GcModeFull
		default:
			gcmode = GcModeFull
		}

		resp[name+"_evm"] = gcmode
	}

	return resp
}
