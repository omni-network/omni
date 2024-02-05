package vmcompose

import (
	"encoding/json"
	"net"
	"os"
	"regexp"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/test/e2e/types"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

var evmRegx = regexp.MustCompile("(.*_evm|chain_.*)")

const (
	evmPort  = 8545
	haloPort = 26657
	relayer  = "relayer"
)

type vmJSON struct {
	Name string
	IP   string
}
type dataJSON struct {
	NetworkCIDR  string            `json:"network_cidr"`
	VMs          []vmJSON          `json:"vms"`
	ServicesByVM map[string]string `json:"services_by_vm"` // map[service_name]vm_name
}

// LoadData returns the vmcompose infrastructure data from the given path.
func LoadData(path string) (types.InfrastructureData, error) {
	bz, err := os.ReadFile(path)
	if err != nil {
		return types.InfrastructureData{}, errors.Wrap(err, "read file")
	}

	var data dataJSON
	err = json.Unmarshal(bz, &data)
	if err != nil {
		return types.InfrastructureData{}, errors.Wrap(err, "unmarshal json")
	}

	vmsByName := make(map[string]e2e.InstanceData)
	for _, vm := range data.VMs {
		ip := net.ParseIP(vm.IP)
		vmsByName[vm.Name] = e2e.InstanceData{
			IPAddress:    ip,
			ExtIPAddress: ip,
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
		if evmRegx.MatchString(serviceName) {
			port = evmPort
		} else if serviceName == relayer {
			port = 0 // No port for relayer
		}

		instances[serviceName] = e2e.InstanceData{
			IPAddress:    vm.IPAddress,
			ExtIPAddress: vm.IPAddress,
			Port:         uint32(port),
		}
	}

	return types.InfrastructureData{
		InfrastructureData: e2e.InfrastructureData{
			Path:      path,
			Provider:  ProviderName,
			Instances: instances,
			Network:   data.NetworkCIDR,
		},
	}, nil
}
