package types

import (
	"context"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
)

func DefaultServiceConfig() ServiceConfig {
	return ServiceConfig{
		Regexp: ".*",
	}
}

type ServiceConfig struct {
	// Regexp to match the service names.
	Regexp string
}

type InfraProvider interface {
	infra.Provider

	// Upgrade copies dynamic config and files to VMs and restarts services.
	// This assumes that important files are long-lived/deterministic (e.g. private keys).
	// It notably doesn't copy newly generated genesis files.
	// Note that all services on matching VMs are upgraded.
	Upgrade(ctx context.Context, cfg ServiceConfig) error

	// Clean deletes all containers, networks, and data on disk.
	Clean(ctx context.Context) error

	// Restart restarts the services that match the given config.
	// I.e., docker-compose up/down.
	// Note that all services on matching VMs are restarted.
	Restart(ctx context.Context, cfg ServiceConfig) error
}

// InfrastructureData wraps e2e.InfrastructureData with additional omni-specific fields.
type InfrastructureData struct {
	e2e.InfrastructureData

	// VMs maps the VM name to its instance data.
	// Note this differs from e2e.InfrastructureData.Instances, which maps the service names to its instance data.
	VMs map[string]e2e.InstanceData
}

// ServicesByInstance returns the set of services associated to the instance.
func (d InfrastructureData) ServicesByInstance(data e2e.InstanceData) map[string]bool {
	resp := make(map[string]bool)
	for serviceName, instance := range d.Instances {
		if instancesEqual(data, instance) {
			resp[serviceName] = true
		}
	}

	return resp
}

// instancesEqual returns true if the two instances are equal, as identified by IPs.
func instancesEqual(a, b e2e.InstanceData) bool {
	return a.IPAddress.Equal(b.IPAddress) && a.ExtIPAddress.Equal(b.ExtIPAddress)
}
