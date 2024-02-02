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

	// PublicChains defines the public chains to connect to; arb_goerli, etc.
	PublicChains []string `toml:"public_chains"`

	// MultiOmniEVMs defines whether to deploy one or multiple Omni EVMs.
	MultiOmniEVMs bool `toml:"multi_omni_evms"`
}

// OmniEVMs returns the names Omni EVMs to deploy.
// If only a single Omni EVM is to be deployed, the name is "omni_evm".
// Otherwise, the names are "<node>_evm".
func (m Manifest) OmniEVMs() []string {
	if !m.MultiOmniEVMs {
		return []string{"omni_evm"}
	}

	var resp []string
	for node := range m.Nodes {
		resp = append(resp, node+"_evm")
	}

	return resp
}
