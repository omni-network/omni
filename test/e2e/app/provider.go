package app

import (
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/test/e2e/docker"
)

var _ Provider = (*docker.Provider)(nil)

// Provider wraps infra.Provider with additional omni-specific methods.
type Provider interface {
	infra.Provider

	// InternalNetwork returns the network configuration from the perspective of
	// the the network nodes themselves, so from within the network, with internal
	// or private IPs.
	InternalNetwork() netconf.Network

	// ExternalNetwork returns the network configuration from the perspective of
	// an outsider. So from outside the network, with external or public IPs.
	ExternalNetwork() netconf.Network
}
