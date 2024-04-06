package docker

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
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

var _ types.InfraProvider = (*Provider)(nil)

// Provider wraps the cometBFT docker provider, writing a different compose file.
type Provider struct {
	*cmtdocker.Provider
	servicesOnce   sync.Once
	testnet        types.Testnet
	omniTag        string
	explorerDBConn string
}

func (*Provider) Clean(ctx context.Context) error {
	log.Info(ctx, "Removing docker containers and networks")

	for _, cmd := range CleanCmds(false, runtime.GOOS == "linux" || runtime.GOOS == "windows") {
		err := exec.Command(ctx, "bash", "-c", cmd)
		if err != nil {
			return errors.Wrap(err, "remove docker containers")
		}
	}

	return nil
}

// NewProvider returns a new Provider.
func NewProvider(testnet types.Testnet, infd types.InfrastructureData, imgTag string, explorerDBConn string) *Provider {
	return &Provider{
		Provider: &cmtdocker.Provider{
			ProviderData: infra.ProviderData{
				Testnet:            testnet.Testnet,
				InfrastructureData: infd.InfrastructureData,
			},
		},
		testnet:        testnet,
		omniTag:        imgTag,
		explorerDBConn: explorerDBConn,
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
		Nodes:          p.testnet.Nodes,
		OmniEVMs:       p.testnet.OmniEVMs,
		Anvils:         p.testnet.AnvilChains,
		Relayer:        true,
		Prometheus:     p.testnet.Prometheus,
		Monitor:        true,
		Explorer:       p.testnet.Explorer,
		ExplorerMockDB: p.testnet.ExplorerMockDB,
		ExplorerDBConn: p.explorerDBConn,
		OmniTag:        p.omniTag,
	}

	bz, err := GenerateComposeFile(def)
	if err != nil {
		return errors.Wrap(err, "generate compose file")
	}

	err = os.WriteFile(filepath.Join(p.Testnet.Dir, "docker-compose.yml"), bz, 0o644)
	if err != nil {
		return errors.Wrap(err, "write compose file")
	}

	return nil
}

func (*Provider) Upgrade(_ context.Context) error {
	return errors.New("upgrade not supported for docker provider")
}

func (p *Provider) StartNodes(ctx context.Context, nodes ...*e2e.Node) error {
	var err error
	p.servicesOnce.Do(func() {
		svcs := additionalServices(p.testnet)
		log.Info(ctx, "Starting additional services", "names", svcs)

		err = cmtdocker.ExecCompose(ctx, p.Testnet.Dir, "create") // This fails if containers not available.
		if err != nil {
			err = errors.Wrap(err, "create containers")
			return
		}

		err = cmtdocker.ExecCompose(ctx, p.Testnet.Dir, append([]string{"up", "-d"}, svcs...)...)
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
	err = cmtdocker.ExecCompose(ctx, p.Testnet.Dir, append([]string{"up", "-d", "--no-deps"}, nodeNames...)...)
	if err != nil {
		return errors.Wrap(err, "start nodes")
	}

	return nil
}

type ComposeDef struct {
	Network     bool
	NetworkName string
	NetworkCIDR string
	BindAll     bool

	Nodes    []*e2e.Node
	OmniEVMs []types.OmniEVM
	Anvils   []types.AnvilChain

	Monitor    bool
	OmniTag    string
	Relayer    bool
	Prometheus bool

	ExplorerMockDB bool
	Explorer       bool
	ExplorerDBConn string
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
	for _, omniEVM := range testnet.OmniEVMs {
		resp = append(resp, omniEVM.InstanceName)
	}
	for _, anvil := range testnet.AnvilChains {
		resp = append(resp, anvil.Chain.Name)
	}

	if testnet.Prometheus {
		resp = append(resp, "prometheus")
	}

	resp = append(resp, "monitor")

	if testnet.Explorer {
		resp = append(resp, "explorer_indexer")
		resp = append(resp, "explorer_graphql")
	}

	if testnet.ExplorerMockDB {
		resp = append(resp, "explorer_db")
	}

	// In monitor only mode, we don't need to start the relayer (above omni and anvils will also be empty).
	if testnet.OnlyMonitor {
		return resp
	}

	resp = append(resp, "relayer")

	return resp
}
