package types

import (
	"context"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
)

type InfraProvider interface {
	infra.Provider

	Upgrade(ctx context.Context) error

	// Clean deletes all containers, networks, and data on disk.
	Clean(ctx context.Context) error
}

// InfrastructureData wraps e2e.InfrastructureData with additional omni-specific fields.
type InfrastructureData struct {
	e2e.InfrastructureData

	// VMs maps the VM name to its instance data.
	// Note this differs from e2e.InfrastructureData.Instances, which maps the service names to its instance data.
	VMs map[string]e2e.InstanceData
}
