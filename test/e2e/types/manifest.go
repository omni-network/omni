package types

import (
	"sort"
	"strings"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

// Manifest wraps e2e.Manifest with additional omni-specific fields.
type Manifest struct {
	e2e.Manifest

	Network string `toml:"network"`

	// AnvilChains defines the anvil chains to deploy; chain_a, chain_b, etc.
	AnvilChains []string `toml:"anvil_chains"`

	// PublicChains defines the public chains to connect to; arb_goerli, etc.
	PublicChains []string `toml:"public_chains"`

	// AVSTarget identifies the chain to deploy the AVS contracts to.
	// It must be one of the anvil or public chains.
	AVSTarget string `toml:"avs_target"`

	// MultiOmniEVMs defines whether to deploy one or multiple Omni EVMs.
	MultiOmniEVMs bool `toml:"multi_omni_evms"`

	// SlowTests defines whether to run slow tests (e.g. tests/eigen_tests.go)
	SlowTests bool `toml:"slow_tests"`
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

	// todo(lazar): tmp: figure out why deploying to seed and full node omnievms resulsts in tx not mined errors
	// Extract keys
	keys := make([]string, 0, len(resp))
	for key := range resp {
		keys = append(keys, key)
	}

	// Sort keys with "validator" nodes first
	sort.Slice(keys, func(i, j int) bool {
		if strings.Contains(keys[i], "validator") && !strings.Contains(keys[j], "validator") {
			return true
		} else if !strings.Contains(keys[i], "validator") && strings.Contains(keys[j], "validator") {
			return false
		}
		return keys[i] < keys[j]
	})

	// Create sorted map
	sortedResp := make(map[string]GcMode)
	for _, key := range keys {
		sortedResp[key] = resp[key]
	}

	return sortedResp
}
