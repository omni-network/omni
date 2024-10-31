package docker

import (
	"bytes"
	"context"
	"fmt"
	"os"
	osexec "os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"
	"text/template"

	"github.com/omni-network/omni/e2e/app/geth"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/exec"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
	cmtdocker "github.com/cometbft/cometbft/test/e2e/pkg/infra/docker"

	_ "embed"
)

const ProviderName = "docker"

// composeTmpl is our own custom docker compose template. This differs from cometBFT's.
//
//go:embed compose.yaml.tmpl
var composeTmpl []byte

// evmInitTmpl is a bash script that initializes all omni-evm geth nodes.
//
//go:embed init-omni-evms.sh.tmpl
var evmInitTmpl []byte

var _ types.InfraProvider = (*Provider)(nil)

// Provider wraps the cometBFT docker provider, writing a different compose file.
type Provider struct {
	*cmtdocker.Provider
	servicesOnce sync.Once
	testnet      types.Testnet
	omniTag      string
}

func (*Provider) Clean(ctx context.Context) error {
	log.Info(ctx, "Removing docker containers and networks")

	for _, cmd := range CleanCmds(false, runtime.GOOS == "linux") {
		err := exec.Command(ctx, "bash", "-c", cmd)
		if err != nil {
			return errors.Wrap(err, "remove docker containers")
		}
	}

	return nil
}

// NewProvider returns a new Provider.
func NewProvider(testnet types.Testnet, infd types.InfrastructureData, imgTag string) *Provider {
	return &Provider{
		Provider: &cmtdocker.Provider{
			ProviderData: infra.ProviderData{
				Testnet:            testnet.Testnet,
				InfrastructureData: infd.InfrastructureData,
			},
		},
		testnet: testnet,
		omniTag: imgTag,
	}
}

// Setup generates the docker-compose file and write it to disk, erroring if
// any of these operations fail.
func (p *Provider) Setup() error {
	def := ComposeDef{
		Network:        true,
		NetworkName:    p.testnet.Name,
		NetworkCIDR:    p.testnet.IP.String(),
		BindAll:        false,
		UpgradeVersion: p.testnet.UpgradeVersion,
		Nodes:          p.testnet.Nodes,
		OmniEVMs:       p.testnet.OmniEVMs,
		Anvils:         p.testnet.AnvilChains,
		Relayer:        true,
		Prometheus:     p.testnet.Prometheus,
		Monitor:        true,
		GethVerbosity:  3, // Info
	}
	def = SetImageTags(def, p.testnet.Manifest, p.omniTag)

	bz, err := GenerateComposeFile(def)
	if err != nil {
		return errors.Wrap(err, "generate compose file")
	}

	err = os.WriteFile(filepath.Join(p.Testnet.Dir, "docker-compose.yaml"), bz, 0o644)
	if err != nil {
		return errors.Wrap(err, "write compose file")
	}

	bz, err = GenerateOmniEVMInitFile(def)
	if err != nil {
		return errors.Wrap(err, "generate evm init file")
	}

	err = os.WriteFile(filepath.Join(p.Testnet.Dir, "evm-init.sh"), bz, 0o755)
	if err != nil {
		return errors.Wrap(err, "write init file")
	}

	return nil
}

func (*Provider) Upgrade(context.Context, types.ServiceConfig) error {
	return errors.New("upgrade not supported for docker provider")
}

func (*Provider) Restart(context.Context, types.ServiceConfig) error {
	return errors.New("restart not supported for docker provider")
}

func (p *Provider) StartNodes(ctx context.Context, nodes ...*e2e.Node) error {
	var err error
	p.servicesOnce.Do(func() {
		svcs := additionalServices(p.testnet)
		log.Info(ctx, "Starting additional services", "names", svcs)

		err = ExecCompose(ctx, p.Testnet.Dir, "create") // This fails if containers not available.
		if err != nil {
			err = errors.Wrap(err, "create containers")
			return
		}

		if err = ExecEVMInit(ctx, p.Testnet.Dir, "evm-init.sh"); err != nil {
			return
		}

		err = ExecCompose(ctx, p.Testnet.Dir, append([]string{"up", "-d"}, svcs...)...)
		if err != nil {
			err = errors.Wrap(err, "start additional services")
			return
		}
	})
	if err != nil {
		return err
	}

	// when we run only a (monitor) service there are no halo nodes available therefore exit early to prevent panics
	if len(nodes) == 0 {
		return nil
	}

	// Start all requested nodes (use --no-deps to avoid starting the additional services again).
	nodeNames := make([]string, len(nodes))
	for i, n := range nodes {
		nodeNames[i] = n.Name
	}
	err = ExecCompose(ctx, p.Testnet.Dir, append([]string{"up", "-d", "--no-deps"}, nodeNames...)...)
	if err != nil {
		return errors.Wrap(err, "start nodes")
	}

	return nil
}

type ComposeDef struct {
	Network        bool
	NetworkName    string
	NetworkCIDR    string
	BindAll        bool
	UpgradeVersion string
	GethVerbosity  int // # Geth log level (1=error,2=warn,3=info(default),4=debug,5=trace)

	Nodes    []*e2e.Node
	OmniEVMs []types.OmniEVM
	Anvils   []types.AnvilChain

	Monitor       bool
	AnvilProxyTag string
	RelayerTag    string
	MonitorTag    string
	Relayer       bool
	Prometheus    bool
}

func (ComposeDef) GethTag() string {
	return geth.Version
}

// NodeOmniEVMs returns a map of node name to OmniEVM instance name; map[node_name]omni_evm.
func (c ComposeDef) NodeOmniEVMs() map[string]string {
	resp := make(map[string]string)
	for i, node := range c.Nodes {
		evm := c.OmniEVMs[0].InstanceName
		if len(c.OmniEVMs) == len(c.Nodes) {
			evm = c.OmniEVMs[i].InstanceName
		}
		resp[node.Name] = evm
	}

	return resp
}

// SetImageTags returns a new ComposeDef with the image tags set.
// This is a convenience function to avoid setting the tags manually.
func SetImageTags(def ComposeDef, manifest types.Manifest, omniImgTag string) ComposeDef {
	anvilProxyTag := omniImgTag

	monitorTag := omniImgTag
	if manifest.PinnedMonitorTag != "" {
		monitorTag = manifest.PinnedMonitorTag
	}

	relayerTag := omniImgTag
	if manifest.PinnedRelayerTag != "" {
		relayerTag = manifest.PinnedRelayerTag
	}

	def.AnvilProxyTag = anvilProxyTag
	def.MonitorTag = monitorTag
	def.RelayerTag = relayerTag

	return def
}

func GenerateComposeFile(def ComposeDef) ([]byte, error) {
	tmpl, err := template.New("compose").Parse(string(composeTmpl))
	if err != nil {
		return nil, errors.Wrap(err, "parse template")
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, def)
	if err != nil {
		return nil, errors.Wrap(err, "execute template")
	}

	return buf.Bytes(), nil
}

func GenerateOmniEVMInitFile(def ComposeDef) ([]byte, error) {
	tmpl, err := template.New("init").Parse(string(evmInitTmpl))
	if err != nil {
		return nil, errors.Wrap(err, "parse template")
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, def)
	if err != nil {
		return nil, errors.Wrap(err, "execute template")
	}

	return buf.Bytes(), nil
}

func ExecEVMInit(ctx context.Context, dir string, evmInitFilename string) error {
	cmd := osexec.CommandContext(ctx, "bash", evmInitFilename)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "exec init command", "output", string(out))
	}

	return nil
}

// CleanCmds returns generic docker commands to clean up docker containers and networks.
// This bypasses the need to a specific docker-compose context.
func CleanCmds(sudo bool, isLinux bool) []string {
	// GNU xargs requires the -r flag to not run when input is empty, macOS
	// does this by default. Ugly, but works.
	xargsR := ""
	if isLinux {
		xargsR = "-r"
	}

	// Some environments need sudo to run docker commands.
	perm := ""
	if sudo {
		perm = "sudo"
	}

	return []string{
		fmt.Sprintf("%s docker container ls -qa --filter label=e2e | xargs %v %s docker container rm -f",
			perm, xargsR, perm),
		fmt.Sprintf("%s docker network ls -q --filter label=e2e | xargs %v %s docker network rm",
			perm, xargsR, perm),
	}
}

// additionalServices returns additional (to halo) docker-compose services to start.
func additionalServices(testnet types.Testnet) []string {
	var resp []string
	if testnet.Prometheus {
		resp = append(resp, "prometheus")
	}

	resp = append(resp, "monitor")

	for _, omniEVM := range testnet.OmniEVMs {
		resp = append(resp, omniEVM.InstanceName)
	}
	for _, anvil := range testnet.AnvilChains {
		resp = append(resp, anvil.Chain.Name)
	}
	resp = append(resp, "relayer")

	return resp
}

// ExecCompose runs a Docker Compose command for a testnet.
func ExecCompose(ctx context.Context, dir string, args ...string) error {
	err := exec.Command(ctx, append(
		[]string{"docker", "compose", "-f", filepath.Join(dir, "docker-compose.yaml")},
		args...)...)
	if err != nil {
		return errors.Wrap(err, "exec docker-compose")
	}

	return nil
}

// ExecComposeVerbose runs a Docker Compose command for a testnet and displays its output.
func ExecComposeVerbose(ctx context.Context, dir string, args ...string) error {
	err := exec.CommandVerbose(ctx, append(
		[]string{"docker", "compose", "-f", filepath.Join(dir, "docker-compose.yaml")},
		args...)...)
	if err != nil {
		return errors.Wrap(err, "exec docker-compose verbose")
	}

	return nil
}

// Exec runs a Docker command.
func Exec(ctx context.Context, args ...string) error {
	err := exec.Command(ctx, append([]string{"docker"}, args...)...)
	if err != nil {
		return errors.Wrap(err, "exec docker")
	}

	return nil
}

// ReplaceUpgradeImage replaces the docker image of the provided service with the
// version specified in comments. Expected format below upgrades node0 from main v1.0 to v2.0:
//
//	  services:
//		 node0:
//		   labels:
//		     e2e: true
//		   container_name: node0
//		   image: omniops/halo:main # Upgrade node0:omniops/halo:v1.0
//		   restart: unless-stopped
func ReplaceUpgradeImage(dir, service string) error {
	before, err := os.ReadFile(filepath.Join(dir, "docker-compose.yaml"))
	if err != nil {
		return errors.Wrap(err, "read compose file")
	}

	regex := regexp.MustCompile(`(\s+image: ).+\s#\sUpgrade ` + service + `:(.*)`)

	after := regex.ReplaceAll(before, []byte("$1$2"))
	if bytes.Equal(before, after) {
		return errors.New("no upgrade image found")
	}

	err = os.WriteFile(filepath.Join(dir, "docker-compose.yaml"), after, 0o644)
	if err != nil {
		return errors.Wrap(err, "write compose file")
	}

	return nil
}
