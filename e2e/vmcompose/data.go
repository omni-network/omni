package vmcompose

import (
	"context"
	"encoding/json"
	"net"
	"os"
	"regexp"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

var omniEvmRegx = regexp.MustCompile(".*_evm")

const (
	evmPort  = uint32(8545)
	haloPort = uint32(26657)
	relayer  = "relayer"
)

type vmJSON struct {
	Name       string `json:"name"`
	IP         string `json:"ip"`
	ExternalIP string `json:"external_ip,omitempty"`
}
type dataJSON struct {
	NetworkCIDR  string            `json:"network_cidr"`
	VMs          []vmJSON          `json:"vms"`
	ServicesByVM map[string]string `json:"services_by_vm"` // map[service_name]vm_name
}

// LoadData returns the vmcompose infrastructure data from the given path.
func LoadData(ctx context.Context, path string) (types.InfrastructureData, error) {
	bz, err := os.ReadFile(path)
	if err != nil {
		return types.InfrastructureData{}, errors.Wrap(err, "read file")
	}

	var data dataJSON
	err = json.Unmarshal(bz, &data)
	if err != nil {
		return types.InfrastructureData{}, errors.Wrap(err, "unmarshal json")
	}

	// hasService returns true if the given VM has a service mapped to it.
	hasService := func(vm string) bool {
		for _, name := range data.ServicesByVM {
			if vm == name {
				return true
			}
		}

		return false
	}

	vmsByName := make(map[string]e2e.InstanceData)
	for _, vm := range data.VMs {
		if !hasService(vm.Name) {
			log.Debug(ctx, "Ignoring VM in infra data without services", nil, "vm", vm.Name, "path", path)
			continue
		}

		ip := net.ParseIP(vm.IP)
		externalIP := net.ParseIP(vm.ExternalIP)

		vmsByName[vm.Name] = e2e.InstanceData{
			IPAddress:    ip,
			ExtIPAddress: externalIP,
		}
	}

	instances := make(map[string]e2e.InstanceData)
	for serviceName, vmName := range data.ServicesByVM {
		vm, ok := vmsByName[vmName]
		if !ok {
			return types.InfrastructureData{}, errors.New("vm not found", "name", vmName)
		}

		// Default ports, as VMs don't support overlapping ports.
		port := haloPort
		if omniEvmRegx.MatchString(serviceName) {
			port = evmPort
		} else if _, ok := evmchain.MetadataByName(serviceName); ok {
			port = evmPort
		} else if serviceName == relayer {
			port = 0 // No port for relayer
		}

		instances[serviceName] = e2e.InstanceData{
			IPAddress:    vm.IPAddress,
			ExtIPAddress: vm.ExtIPAddress,
			Port:         port,
		}
	}

	return types.InfrastructureData{
		InfrastructureData: e2e.InfrastructureData{
			Path:      path,
			Provider:  ProviderName,
			Instances: instances,
			Network:   data.NetworkCIDR,
		},
		VMs: vmsByName,
	}, nil
}
