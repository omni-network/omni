package types

import (
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

// InfrastructureData wraps e2e.InfrastructureData with additional omni-specific fields.
type InfrastructureData struct {
	e2e.InfrastructureData

	// OmniEVMs defines the infrastructure data for the deployed Omni EVMs (keyed by instance name).
	OmniEVMs map[string]e2e.InstanceData `json:"omni_evms"`

	// AnvilChains defines the instance data per deployed Anvil chains (keyed by chain name).
	AnvilChains map[string]e2e.InstanceData `json:"anvil_chains"`
}
