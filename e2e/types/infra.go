package types

import (
	"context"
	"regexp"
	"strings"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
)

const (
	// regexpCanary is a convenient way to specify canary upgrades.
	regexpCanary = "canary"
	// regexpNonCanary is a convenient way to specify non-canary upgrades.
	regexpNonCanary = "non-canary"
	// regexpHalo is a convenient way to specify all halo nodes (excluding relayer/monitor/solver etc).
	regexpHalo = "halo"
)

// canaries define services included in canary upgrades
// and excluded from non-canary upgrades.
var canaries = map[string]bool{
	"validator01": true,
	"fullnode01":  true,
	"archive01":   true,
	"seed01":      true,
	"relayer":     true,
	"monitor":     true,
	"solver":      true,
}

func DefaultServiceConfig() ServiceConfig {
	return ServiceConfig{
		Regexp: ".*",
	}
}

type ServiceConfig struct {
	// Regexp to match the service names.
	Regexp string
}

// MatchService returns true if the service matches the regexp config.
func (c ServiceConfig) MatchService(service string) bool {
	if c.Regexp == "" {
		return true
	}

	// isCanary returns true if the service is a canary node.
	isCanary := func() bool {
		for canary := range canaries {
			if strings.HasPrefix(service, canary) {
				return true
			}
		}

		return false
	}

	// isHalo returns true if the service is a halo node.
	isHalo := func() bool {
		if strings.HasSuffix(service, "_evm") {
			return false
		}

		for _, prefix := range []string{"validator", "full", "seed", "archive"} {
			if strings.HasPrefix(service, prefix) {
				return true
			}
		}

		return false
	}

	if c.Regexp == regexpCanary {
		return isCanary()
	} else if c.Regexp == regexpNonCanary {
		return !isCanary()
	} else if c.Regexp == regexpHalo {
		return isHalo()
	}

	ok, _ := regexp.MatchString(c.Regexp, service) // Nothing matches invalid regex

	return ok
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
