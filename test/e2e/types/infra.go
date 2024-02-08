package types

import (
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

// InfrastructureData wraps e2e.InfrastructureData with additional omni-specific fields.
type InfrastructureData struct {
	e2e.InfrastructureData

	// VMs maps the VM name to its instance data.
	// Note this differs from e2e.InfrastructureData.Instances, which maps the service names to its instance data.
	VMs map[string]e2e.InstanceData
}
