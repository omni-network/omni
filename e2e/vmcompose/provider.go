package vmcompose

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/omni-network/omni/e2e/app/agent"
	"github.com/omni-network/omni/e2e/docker"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"golang.org/x/sync/errgroup"
)

const ProviderName = "vmcompose"

var _ types.InfraProvider = (*Provider)(nil)

type Provider struct {
	Testnet    types.Testnet
	Data       types.InfrastructureData
	once       sync.Once
	omniTag    string
	graphQLURL string
}

func NewProvider(testnet types.Testnet, data types.InfrastructureData, imgTag, graphQLURL string) *Provider {
	return &Provider{
		Testnet:    testnet,
		Data:       data,
		omniTag:    imgTag,
		graphQLURL: graphQLURL,
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

		var geths []string
		for _, omniEVM := range p.Testnet.OmniEVMs {
			if services[omniEVM.InstanceName] {
				geths = append(geths, omniEVM.InstanceName)
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

		// All explorer containers are run on the same VM
		explorer := p.Testnet.Explorer && (services["explorer_ui"] || services["explorer_graphql"] || services["explorer_indexer"])

		def := docker.ComposeDef{
			Network:     false,
			BindAll:     true,
			NetworkName: p.Testnet.Name,
			NetworkCIDR: p.Testnet.IP.String(),
			Nodes:       nodes,
			OmniEVMs:    omniEVMs,
			Anvils:      anvilChains,
			Relayer:     services["relayer"],
			Monitor:     services["monitor"],
			Prometheus:  p.Testnet.Prometheus,
			OmniTag:     p.omniTag,
			Explorer:    explorer,
			ExplorerDB:  explorer && p.Testnet.Network.IsEphemeral(),
			GraphQLURL:  p.graphQLURL,
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
		agentCfg = agent.ConfigForHost(agentCfg, hostname, halos, geths, services)
		err = os.WriteFile(filepath.Join(p.Testnet.Dir, vmAgentFile(vmIP)), agentCfg, 0o644)
		if err != nil {
			return errors.Wrap(err, "write compose file")
		}
	}

	return nil
}

// Upgrade copies some of the local e2e generated artifacts to the VMs and starts the docker-compose services.
func (p *Provider) Upgrade(ctx context.Context, cfg types.UpgradeConfig) error {
	log.Info(ctx, "Upgrading docker-compose on VMs", "image", p.Testnet.UpgradeVersion)

	filesByService := make(map[string][]string)
	addFile := func(service string, paths ...string) {
		filesByService[service] = append(filesByService[service], filepath.Join(paths...))
	}

	// Include halo config
	for _, node := range p.Testnet.Nodes {
		addFile(node.Name, "config", "halo.toml")
		addFile(node.Name, "config", "config.toml")
		addFile(node.Name, "config", "network.json")
		addFile(node.Name, "config", "jwtsecret")
		addFile(node.Name, "config", "priv_validator_key.json")
		addFile(node.Name, "config", "node_key.json")
	}

	// Include geth config
	for _, omniEVM := range p.Testnet.OmniEVMs {
		addFile(omniEVM.InstanceName, "config.toml")
		addFile(omniEVM.InstanceName, "geth", "jwtsecret")
		addFile(omniEVM.InstanceName, "geth", "nodekey")
	}

	// Also relayer and monitor
	addFile("relayer", "relayer.toml")
	addFile("relayer", "network.json")
	addFile("relayer", "privatekey")
	addFile("monitor", "monitor.toml")
	addFile("monitor", "network.json")
	addFile("monitor", "privatekey")

	// TODO(corver): Add explorer stuff

	addFile("prometheus", "prometheus.yml") // Prometheus isn't a "service", so not actually copied

	// Do initial sequential ssh to each VM, ensure we can connect.
	for vmName, instance := range p.Data.VMs {
		if !matchAny(cfg, p.Data.ServicesByInstance(instance)) {
			log.Debug(ctx, "Skipping vm upgrade, no matching services", "vm", vmName, "regexp", cfg.ServiceRegexp)
			continue
		}

		log.Debug(ctx, "Ensuring VM SSH connection", "vm", vmName)
		if err := execOnVM(ctx, vmName, "ls"); err != nil {
			return errors.Wrap(err, "test exec on vm", "vm", vmName)
		}
	}

	// Then upgrade VMs in parallel
	eg, ctx := errgroup.WithContext(ctx)
	for vmName, instance := range p.Data.VMs {
		services := p.Data.ServicesByInstance(instance)
		if !matchAny(cfg, services) {
			continue
		}

		eg.Go(func() error {
			log.Debug(ctx, "Copying artifacts", "vm", vmName, "count", len(filesByService))
			for service, filePaths := range filesByService {
				if !services[service] {
					continue
				}
				for _, filePath := range filePaths {
					localPath := filepath.Join(p.Testnet.Dir, service, filePath)
					remotePath := filepath.Join("/omni", p.Testnet.Name, service, filePath)
					if err := copyFileToVM(ctx, vmName, localPath, remotePath); err != nil {
						return errors.Wrap(err, "copy file", "vm", vmName, "service", service, "file", filePath)
					}
				}
			}

			log.Debug(ctx, "Copying docker-compose.yml", "vm", vmName)
			composeFile := vmComposeFile(instance.IPAddress.String())
			localComposePath := filepath.Join(p.Testnet.Dir, composeFile)
			remoteComposePath := filepath.Join("/omni", p.Testnet.Name, composeFile)
			if err := copyFileToVM(ctx, vmName, localComposePath, remoteComposePath); err != nil {
				return errors.Wrap(err, "copy docker compose", "vm", vmName)
			}

			// Figure out whether we need to call "docker compose down" before "docker compose up"
			var maybeDown string
			if services := p.Data.ServicesByInstance(instance); containsEVM(services) {
				if containsAnvil(services) {
					return errors.New("cannot upgrade VM with both omni_evm and anvil containers since omni_evm needs downing and anvil cannot be restarted", "vm", vmName)
				}
				maybeDown = "sudo docker compose down && "
			}

			startCmd := fmt.Sprintf("cd /omni/%s && "+
				"sudo mv %s docker-compose.yaml && "+
				maybeDown+
				"sudo docker compose up -d",
				p.Testnet.Name, composeFile)

			log.Debug(ctx, "Executing docker-compose up", "vm", vmName)
			if err := execOnVM(ctx, vmName, startCmd); err != nil {
				return errors.Wrap(err, "compose up", "vm", vmName)
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "wait errgroup")
	}

	return nil
}

// matchAny returns true if the pattern matches any of the services in the services map.
// An empty pattern returns true, matching anything.
func matchAny(cfg types.UpgradeConfig, services map[string]bool) bool {
	if cfg.ServiceRegexp == "" {
		return true
	}

	for service := range services {
		matched, _ := regexp.MatchString(cfg.ServiceRegexp, service)
		if matched {
			return true
		}
	}

	return false
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

func (*Provider) StopTestnet(context.Context) error {
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

func copyFileToVM(ctx context.Context, vmName string, localPath string, remotePath string) error {
	scp := fmt.Sprintf("gcloud compute scp --zone=us-east1-c --quiet %s %s:%s", localPath, vmName, remotePath)
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

// containsEVM returns true if the services map contains an omni evm.
// TODO(corver): This isn't very robust.
func containsEVM(services map[string]bool) bool {
	for service, ok := range services {
		if !ok {
			continue
		}
		if strings.Contains(service, "evm") {
			return true
		}
	}

	return false
}

// containsAnvil returns true if the services map contains an anvil chain.
// TODO(corver): This isn't very robust.
func containsAnvil(services map[string]bool) bool {
	for service, ok := range services {
		if !ok {
			continue
		}
		if strings.Contains(service, "mock") || strings.Contains(service, "chain") {
			return true
		}
	}

	return false
}
