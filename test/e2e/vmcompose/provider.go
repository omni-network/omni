package vmcompose

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/test/e2e/docker"
	"github.com/omni-network/omni/test/e2e/types"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
)

const ProviderName = "vmcompose"

var _ infra.Provider = (*Provider)(nil)

type Provider struct {
	Testnet types.Testnet
	Data    types.InfrastructureData
}

func NewProvider(testnet types.Testnet, data types.InfrastructureData) *Provider {
	return &Provider{
		Testnet: testnet,
		Data:    data,
	}
}

// Setup generates the docker-compose file for each VM IP.
func (p *Provider) Setup() error {
	// Group infra services by VM IP
	for vmIP, services := range groupByVM(p.Data.Instances) {
		// Get all halo nodes in this VM
		var nodes []*e2e.Node
		for _, node := range p.Testnet.Nodes {
			if services[node.Name] {
				nodes = append(nodes, node)
			}
		}

		// Get all omniEVMs in this VM
		var omniEVMs []types.OmniEVM
		for _, omniEVM := range p.Testnet.OmniEVMs {
			if services[omniEVM.InstanceName] {
				omniEVMs = append(omniEVMs, omniEVM)
			}
		}

		// Get all anvil chains in this VM
		var anvilChains []types.AnvilChain
		for _, anvilChain := range p.Testnet.AnvilChains {
			if services[anvilChain.Chain.Name] {
				anvilChains = append(anvilChains, anvilChain)
			}
		}

		def := docker.ComposeDef{
			NetworkName: p.Testnet.Name,
			NetworkCIDR: p.Testnet.IP.String(),
			Nodes:       nodes,
			OmniEVMs:    omniEVMs,
			Anvils:      anvilChains,
			Relayer:     services["relayer"],
			Prometheus:  p.Testnet.Prometheus,
		}
		compose, err := docker.GenerateComposeFile(def)
		if err != nil {
			return errors.Wrap(err, "generate compose file")
		}

		filename := strings.ReplaceAll(vmIP, ".", "_") + "_compose.yaml"

		err = os.WriteFile(filepath.Join(p.Testnet.Dir, filename), compose, 0o644)
		if err != nil {
			return errors.Wrap(err, "write compose file")
		}
	}

	return nil
}

func (p *Provider) StartNodes(ctx context.Context, node ...*e2e.Node) error {
	// TODO(corver): Copy all locally generated compose files and folders and config to VMs (see ../docker sync.Once).

	return errors.New("not implemented")
}

func (p *Provider) StopTestnet(ctx context.Context) error {
	return errors.New("not implemented")
}

func (p *Provider) GetInfrastructureData() *e2e.InfrastructureData {
	return &p.Data.InfrastructureData
}

func groupByVM(instances map[string]e2e.InstanceData) map[string]map[string]bool {
	resp := make(map[string]map[string]bool) // map[vm_ip]map[service_name]true
	for serviceName, instance := range instances {
		ip := instance.IPAddress.String()
		m, ok := resp[ip]
		if !ok {
			m = make(map[string]bool)
		}
		m[serviceName] = true
		resp[ip] = m
	}

	return resp
}
