package docker

import (
	"fmt"
	"net"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

const (
	ipPrefix      = "10.186.73." // See github.com/cometbft/cometbft/test/e2e/pkg for reference
	startIPSuffix = 100
	startPort     = uint32(8000)
)

var localhost = net.ParseIP("127.0.0.1") //nolint:gochecknoglobals // Static IP

// NewInfraData returns a new InfrastructureData for the given manifest.
// In addition to normal.
func NewInfraData(manifest types.Manifest) (types.InfrastructureData, error) {
	infd, err := e2e.NewDockerInfrastructureData(manifest.Manifest)
	if err != nil {
		return types.InfrastructureData{}, errors.Wrap(err, "creating docker infrastructure data")
	}

	// IP generator
	ipSuffix := startIPSuffix
	nextInternalIP := func() net.IP {
		defer func() { ipSuffix++ }()
		return net.ParseIP(fmt.Sprintf(ipPrefix+"%d", ipSuffix))
	}

	// Port generator
	port := startPort
	nextPort := func() uint32 {
		defer func() { port++ }()
		return port
	}

	for name := range manifest.OmniEVMs() {
		infd.Instances[name] = e2e.InstanceData{
			IPAddress:    nextInternalIP(),
			ExtIPAddress: localhost,
			Port:         nextPort(),
		}
	}

	for _, name := range manifest.AnvilChains {
		infd.Instances[name] = e2e.InstanceData{
			IPAddress:    nextInternalIP(),
			ExtIPAddress: localhost,
			Port:         nextPort(),
		}
	}

	// No IP for relayer or monitor required since they doesn't serve an API.

	// Solver IP and port hardcoded in docker/compose.yaml.tmpl
	infd.Instances["solver"] = e2e.InstanceData{
		IPAddress:    net.ParseIP("10.186.73.203"),
		ExtIPAddress: localhost,
		Port:         26661,
	}

	return types.InfrastructureData{
		InfrastructureData: infd,
	}, nil
}
