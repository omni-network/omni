package vmcompose

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/omni-network/omni/e2e/app/agent"
	"github.com/omni-network/omni/e2e/docker"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

const ProviderName = "vmcompose"

var _ types.InfraProvider = (*Provider)(nil)

type Provider struct {
	Testnet    types.Testnet
	Data       types.InfrastructureData
	once       sync.Once
	relayerTag string
}

func NewProvider(testnet types.Testnet, data types.InfrastructureData, relayerTag string) *Provider {
	return &Provider{
		Testnet:    testnet,
		Data:       data,
		relayerTag: relayerTag,
	}
}

// Setup generates the docker-compose file for each VM IP.
func (p *Provider) Setup() error {
	// Group infra services by VM IP
	for vmIP, services := range groupByVM(p.Data.Instances) {
		// Get all halo nodes in this VM
		var nodes []*e2e.Node
		var halos []string
		for _, node := range p.Testnet.Nodes {
			if services[node.Name] {
				nodes = append(nodes, node)
				halos = append(halos, node.Name)
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
			Network:       false,
			BindAll:       true,
			NetworkName:   p.Testnet.Name,
			NetworkCIDR:   p.Testnet.IP.String(),
			Nodes:         nodes,
			OmniEVMs:      omniEVMs,
			Anvils:        anvilChains,
			Relayer:       services["relayer"],
			Prometheus:    p.Testnet.Prometheus,
			RelayerTag:    p.relayerTag,
			OmniLogFormat: log.FormatLogfmt, // VM compose always use logfmt log format.
		}
		compose, err := docker.GenerateComposeFile(def)
		if err != nil {
			return errors.Wrap(err, "generate compose file")
		}

		err = os.WriteFile(filepath.Join(p.Testnet.Dir, vmComposeFile(vmIP)), compose, 0o644)
		if err != nil {
			return errors.Wrap(err, "write compose file")
		}

		if !p.Testnet.Prometheus {
			continue // No need to generate prometheus config
		}

		// Update custom prometheus.yml config for this VM
		promCfgFile := filepath.Join(p.Testnet.Dir, "prometheus", "prometheus.yml")
		agentCfg, err := os.ReadFile(promCfgFile)
		if err != nil {
			return errors.Wrap(err, "read prometheus config")
		}

		hostname := vmIP // TODO(corver): Add hostnames to infra instances.
		agentCfg = agent.ConfigForHost(agentCfg, hostname, halos, services["relayer"])
		err = os.WriteFile(filepath.Join(p.Testnet.Dir, vmAgentFile(vmIP)), agentCfg, 0o644)
		if err != nil {
			return errors.Wrap(err, "write compose file")
		}
	}

	return nil
}

func (p *Provider) Upgrade(ctx context.Context) error {
	log.Info(ctx, "Upgrading docker-compose on VMs", "image", p.Testnet.UpgradeVersion)
	for vmName, instance := range p.Data.VMs {
		log.Debug(ctx, "Upgrading docker-compose", "vm", vmName)

		composeFile := vmComposeFile(instance.IPAddress.String())
		err := copyFileToVM(ctx, vmName, filepath.Join(p.Testnet.Dir, composeFile))
		if err != nil {
			return errors.Wrap(err, "copy compose", "vm", vmName)
		}

		startCmd := fmt.Sprintf("cd /omni && "+
			"sudo mv %s %s/docker-compose.yaml && "+
			"cd %s && "+
			"sudo docker compose up -d",
			composeFile, p.Testnet.Name, p.Testnet.Name)

		err = execOnVM(ctx, vmName, startCmd)
		if err != nil {
			return errors.Wrap(err, "compose up", "vm", vmName)
		}
	}

	return nil
}

func (p *Provider) StartNodes(ctx context.Context, _ ...*e2e.Node) error {
	var onceErr error
	p.once.Do(func() {
		log.Info(ctx, "Copying artifacts to VMs")
		for vmName := range p.Data.VMs {
			err := copyToVM(ctx, vmName, p.Testnet.Dir)
			if err != nil {
				onceErr = errors.Wrap(err, "copy files", "vm", vmName)
				return
			}
		}

		log.Info(ctx, "Starting VM deployments")
		// TODO(corver): Only start additional services and then start halo as per above StartNodes.
		for vmName, instance := range p.Data.VMs {
			composeFile := vmComposeFile(instance.IPAddress.String())
			agentFile := vmAgentFile(instance.IPAddress.String())
			startCmd := fmt.Sprintf("cd /omni/%s && "+
				"sudo mv %s docker-compose.yaml && "+
				"sudo mv %s prometheus/prometheus.yml && "+
				"sudo docker compose up -d",
				p.Testnet.Name, composeFile, agentFile)

			err := execOnVM(ctx, vmName, startCmd)
			if err != nil {
				onceErr = errors.Wrap(err, "compose up", "vm", vmName)
				return
			}
		}
	})

	return onceErr
}

func (p *Provider) Clean(ctx context.Context) error {
	log.Info(ctx, "Deleting existing VM deployments including data")
	for vmName := range p.Data.VMs {
		for _, cmd := range docker.CleanCmds(true, true) {
			err := execOnVM(ctx, vmName, cmd)
			if err != nil {
				return errors.Wrap(err, "clean docker containers", "vm", vmName)
			}

			err = execOnVM(ctx, vmName, "sudo rm -rf /omni/*")
			if err != nil {
				return errors.Wrap(err, "clean docker containers", "vm", vmName)
			}
		}
	}

	return nil
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

func execOnVM(ctx context.Context, vmName string, cmd string) error {
	ssh := fmt.Sprintf("gcloud compute ssh --zone=us-east1-c %s --quiet -- \"%s\"", vmName, cmd)

	out, err := exec.CommandContext(ctx, "bash", "-c", ssh).CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "exec on VM", "output", string(out), "cmd", ssh)
	}

	return nil
}

func copyToVM(ctx context.Context, vmName string, path string) error {
	tarscp := fmt.Sprintf("tar czf - %s | gcloud compute ssh --zone=us-east1-c %s --quiet -- \"cd /omni && tar xzf -\"", filepath.Base(path), vmName)

	cmd := exec.CommandContext(ctx, "bash", "-c", tarscp)
	cmd.Dir = filepath.Dir(path)
	if out, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrap(err, "copy to VM", "output", string(out))
	}

	return nil
}

func copyFileToVM(ctx context.Context, vmName string, path string) error {
	scp := fmt.Sprintf("gcloud compute scp --zone=us-east1-c --quiet %s %s:/omni/", path, vmName)

	cmd := exec.CommandContext(ctx, "bash", "-c", scp)
	if out, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrap(err, "copy to VM", "output", string(out), "cmd", scp)
	}

	return nil
}

func vmAgentFile(internalIP string) string {
	return "prometheus/" + strings.ReplaceAll(internalIP, ".", "_") + "_prometheus.yml"
}

func vmComposeFile(internalIP string) string {
	return strings.ReplaceAll(internalIP, ".", "_") + "_compose.yaml"
}
