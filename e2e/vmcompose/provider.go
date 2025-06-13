package vmcompose

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/omni-network/omni/e2e/app/agent"
	"github.com/omni-network/omni/e2e/docker"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"golang.org/x/sync/errgroup"
)

const ProviderName = "vmcompose"

var _ types.InfraProvider = (*Provider)(nil)

type Provider struct {
	Testnet types.Testnet
	Data    types.InfrastructureData
	once    sync.Once
	omniTag string
}

func NewProvider(testnet types.Testnet, data types.InfrastructureData, imgTag string) *Provider {
	return &Provider{
		Testnet: testnet,
		Data:    data,
		omniTag: imgTag,
	}
}

func (*Provider) AllLogs(context.Context) ([]byte, error) {
	return nil, errors.New("vmcompose doesn't support all logs")
}

// Setup generates the docker-compose file for each VM IP.
func (p *Provider) Setup() error {
	if len(p.Testnet.SVMChains) > 0 {
		return errors.New("vmcompose doesn't support SVM chains")
	}

	// Group infra services by VM IP
	for vmIP, services := range groupByVM(p.Data.Instances) {
		// Get all halo nodes in this VM
		var nodes []*e2e.Node
		var halos []string
		for _, node := range p.Testnet.Nodes {
			if node.Version != p.Testnet.UpgradeVersion {
				return errors.New("node upgrades not supported for vmcompose",
					"node_version", node.Version,
					"upgrade_version", p.Testnet.UpgradeVersion,
				)
			}

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
				omniEVM.ProxyPort = 8544 // Halo EVM-proxy uses 8454 (via BindAll=true)
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

		gethVerbosity := 4 // Debug level

		def := docker.ComposeDef{
			UpgradeVersion: p.Testnet.UpgradeVersion,
			Network:        false,
			BindAll:        true,
			NetworkName:    p.Testnet.Name,
			NetworkCIDR:    p.Testnet.IP.String(),
			Nodes:          nodes,
			OmniEVMs:       omniEVMs,
			Anvils:         anvilChains,
			Relayer:        services["relayer"],
			Monitor:        services["monitor"],
			Solver:         services["solver"],
			Prometheus:     p.Testnet.Prometheus,
			GethVerbosity:  gethVerbosity,
			AnvilAMD:       docker.AnvilAMD(),
		}
		def = docker.SetImageTags(def, p.Testnet.Manifest, p.omniTag)

		compose, err := docker.GenerateComposeFile(def)
		if err != nil {
			return errors.Wrap(err, "generate compose file")
		}

		err = os.WriteFile(filepath.Join(p.Testnet.Dir, vmComposeFile(vmIP)), compose, 0o644)
		if err != nil {
			return errors.Wrap(err, "write compose file")
		}

		evmInit, err := docker.GenerateOmniEVMInitFile(def)
		if err != nil {
			return errors.Wrap(err, "generate evm init file")
		}

		err = os.WriteFile(filepath.Join(p.Testnet.Dir, vmInitFile(vmIP)), evmInit, 0o755)
		if err != nil {
			return errors.Wrap(err, "write compose file")
		}

		if !p.Testnet.Prometheus {
			continue // No need to generate prometheus config
		}

		// Update custom prometheus.yaml config for this VM
		promCfgFile := filepath.Join(p.Testnet.Dir, "prometheus", "prometheus.yaml")
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

func (p *Provider) Restart(ctx context.Context, cfg types.ServiceConfig) error {
	log.Info(ctx, "Restarting docker-compose on VMs")

	for vmName, instance := range p.Data.VMs {
		if !matchAny(cfg, p.Data.ServicesByInstance(instance)) {
			log.Debug(ctx, "Skipping vm restart, no matching services", "vm", vmName, "regexp", cfg.Regexp)
			continue
		}

		startCmd := fmt.Sprintf("cd /omni/%s && "+
			"sudo docker compose down && "+
			"sudo docker compose up -d", // Don't pull on restart
			p.Testnet.Name)

		err := execOnVM(ctx, vmName, startCmd)
		if err != nil {
			return errors.Wrap(err, "compose down up ", "vm", vmName)
		}
	}

	return nil
}

// Upgrade copies some of the local e2e generated artifacts to the VMs and starts the docker-compose services.
func (p *Provider) Upgrade(ctx context.Context, cfg types.ServiceConfig) error {
	log.Info(ctx, "Upgrading docker-compose artifacts on VMs")

	filesByService := make(map[string][]string)
	addFile := func(service string, paths ...string) {
		filesByService[service] = append(filesByService[service], filepath.Join(paths...))
	}

	// Include halo config
	for _, node := range p.Testnet.Nodes {
		addFile(node.Name, "config", "halo.toml")
		addFile(node.Name, "config", "config.toml")
		addFile(node.Name, "config", "priv_validator_key.json")
		addFile(node.Name, "config", "node_key.json")
	}

	// Include geth config
	for _, omniEVM := range p.Testnet.OmniEVMs {
		addFile(omniEVM.InstanceName, "config.toml")
		addFile(omniEVM.InstanceName, "geth", "nodekey")
	}

	// Also relayer and monitor and solver
	addFile("relayer", "relayer.toml")
	addFile("relayer", "privatekey")
	addFile("monitor", "monitor.toml")
	addFile("monitor", "privatekey")
	addFile("monitor", "xcaller_privatekey")
	addFile("monitor", "flowgen_privatekey")
	addFile("solver", "solver.toml")
	addFile("solver", "privatekey")

	// Do initial sequential ssh to each VM, ensure we can connect.
	for vmName, instance := range p.Data.VMs {
		if !matchAny(cfg, p.Data.ServicesByInstance(instance)) {
			log.Debug(ctx, "Skipping vm upgrade, no matching services", "vm", vmName, "regexp", cfg.Regexp)
			continue
		}

		log.Debug(ctx, "Ensuring VM SSH connection and permissions", "vm", vmName) // Also fix permissions
		if err := execOnVM(ctx, vmName, "sudo chmod -R o+w /omni/"+p.Testnet.Name); err != nil {
			return errors.Wrap(err, "chown perms exec on vm", "vm", vmName)
		}
	}

	// Backup files to GCP, e.g.: e2e-configs/omega/UPGRADE2024-09-19T16:27:55Z/...
	gcpBucket := fmt.Sprintf("e2e-configs/%s/UPGRADE%s/", p.Testnet.Name, time.Now().Format(time.RFC3339))

	// Then upgrade VMs in parallel
	eg, ctx := errgroup.WithContext(ctx)

	for vmName, instance := range p.Data.VMs {
		services := p.Data.ServicesByInstance(instance)
		if !matchAny(cfg, services) {
			continue
		}

		eg.Go(func() error {
			filesToCopy := map[string]string{ // Local to remote files to copy.
				vmInitFile(instance.IPAddress.String()):    "evm-init.sh",
				vmComposeFile(instance.IPAddress.String()): "docker-compose.yaml",
				vmAgentFile(instance.IPAddress.String()):   "prometheus/prometheus.yaml",
			}
			for service, filePaths := range filesByService {
				if !services[service] {
					continue
				}
				for _, filePath := range filePaths {
					serviceFile := filepath.Join(service, filePath)
					filesToCopy[serviceFile] = serviceFile
				}
			}

			log.Debug(ctx, "Copying artifacts", "vm", vmName, "count", len(filesToCopy))

			for local, remote := range filesToCopy {
				localPath := filepath.Join(p.Testnet.Dir, local)
				remotePath := filepath.Join("/omni", p.Testnet.Name, remote)
				if err := copyFileToVM(ctx, vmName, localPath, remotePath); err != nil {
					return errors.Wrap(err, "copy file to VM", "vm", vmName, "local", local, "remote", remotePath)
				}
				go backupToGCP(ctx, localPath, filepath.Join(gcpBucket, local)) // Best-effort backup
			}

			// TODO(corver): Once evm-init.sh is idempotent, call it here.

			startCmd := fmt.Sprintf("cd /omni/%s && "+
				"sudo docker compose pull && "+
				"sudo docker compose down && "+
				"sudo docker compose up -d &&"+
				"sudo docker system prune -a -f", // Prune old images
				p.Testnet.Name)

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
func matchAny(cfg types.ServiceConfig, services map[string]bool) bool {
	for service := range services {
		matched := cfg.MatchService(service)
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
		timestampDir := time.Now()
		identifier := fmt.Sprintf("DEPLOY%s", timestampDir.Format(time.RFC3339))
		for vmName := range p.Data.VMs {
			err := copyToVM(ctx, vmName, p.Testnet.Dir)
			if err != nil {
				onceErr = errors.Wrap(err, "copy files to VM", "vm", vmName)
				return
			}
			go func(vmName string) {
				err = copyToGCP(ctx, vmName, p.Testnet.Dir, p.Testnet.Name, identifier)
				if err != nil {
					log.Warn(ctx, "Failed to copy to GCP", err, "vm", vmName)
				}
			}(vmName)
		}

		log.Info(ctx, "Starting VM deployments")
		// TODO(corver): Only start additional services and then start halo as per above StartNodes.
		for vmName, instance := range p.Data.VMs {
			composeFile := vmComposeFile(instance.IPAddress.String())
			initFile := vmInitFile(instance.IPAddress.String())
			agentFile := vmAgentFile(instance.IPAddress.String())
			startCmd := fmt.Sprintf("cd /omni/%s && "+
				"sudo mv %s docker-compose.yaml && "+
				"sudo mv %s evm-init.sh && "+
				"sudo mv %s prometheus/prometheus.yaml && "+
				"sudo docker compose pull &&"+
				"sudo ./evm-init.sh && "+
				"sudo docker compose up -d &&"+
				"sudo docker system prune -a -f", // Prune old images
				p.Testnet.Name, composeFile, initFile, agentFile)

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

func copyToGCP(ctx context.Context, vmName string, path string, network string, identifier string) error {
	cpgcp := fmt.Sprintf("tar czf %s.tar.gz %s && gcloud storage cp %s.tar.gz gs://e2e-configs/%s/%s/", vmName, filepath.Base(path), vmName, network, identifier)

	cmd := exec.CommandContext(ctx, "bash", "-c", cpgcp)
	cmd.Dir = filepath.Dir(path)
	if out, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrap(err, "copy to GCP error", "output", string(out), "cmd", cpgcp)
	}

	return nil
}

func copyFileToVM(ctx context.Context, vmName string, localPath string, remotePath string) error {
	scp := fmt.Sprintf("gcloud compute scp --zone=us-east1-c --quiet %s %s:%s", localPath, vmName, remotePath)

	// Scp sporadically fails with "Permission denied" on GCP, retry a few times.
	errPermission := errors.New("scp permission denied")
	check := func(err error) bool {
		return errors.Is(err, errPermission)
	}

	do := func() error {
		cmd := exec.CommandContext(ctx, "bash", "-c", scp)
		out, err := cmd.CombinedOutput()
		if err != nil && strings.Contains(string(out), "Permission denied") {
			// First try best-effort deleting remote file, then retry.
			if err = execOnVM(ctx, vmName, "sudo rm "+remotePath); err != nil {
				log.Warn(ctx, "Failed to delete remote file as perm issue workaround", err, "path", remotePath)
			}

			return errors.Wrap(errPermission, "copy to VM", "output", string(out), "cmd", scp)
		} else if err != nil {
			return errors.Wrap(err, "copy to VM", "output", string(out), "cmd", scp)
		}

		return nil
	}

	return expbackoff.Retry(ctx, do, expbackoff.WithRetryCheck(check))
}

func backupToGCP(ctx context.Context, localPath string, bucket string) {
	cmd := fmt.Sprintf("gcloud storage cp %s gs://%s", localPath, bucket)
	out, err := exec.CommandContext(ctx, "bash", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Warn(ctx, "Backup to GCP failed", err,
			"cmd", cmd,
			"path", localPath,
			"bucket", bucket,
			"output", string(out),
		)
	}
}

func vmAgentFile(internalIP string) string {
	return "prometheus/" + strings.ReplaceAll(internalIP, ".", "-") + "-prometheus.yaml"
}

func vmComposeFile(internalIP string) string {
	return strings.ReplaceAll(internalIP, ".", "-") + "-compose.yaml"
}

func vmInitFile(internalIP string) string {
	return strings.ReplaceAll(internalIP, ".", "-") + "-evm-init.sh"
}
