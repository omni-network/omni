package docker

import (
	"bytes"
	"context"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/types"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
	cmtdocker "github.com/cometbft/cometbft/test/e2e/pkg/infra/docker"

	"github.com/ethereum/go-ethereum/crypto"

	_ "embed"
)

// composeTmpl is our own custom docker compose template. This differs from cometBFT's.
//
//go:embed compose.yaml.tmpl
var composeTmpl []byte

var _ infra.Provider = &Provider{}

// Provider wraps the cometBFT docker provider, writing a different compose file.
type Provider struct {
	*cmtdocker.Provider
	servicesOnce sync.Once
	testnet      types.Testnet
}

// NewProvider returns a new Provider.
func NewProvider(testnet types.Testnet, infd types.InfrastructureData) *Provider {
	return &Provider{
		Provider: &cmtdocker.Provider{
			ProviderData: infra.ProviderData{
				Testnet:            testnet.Testnet,
				InfrastructureData: infd.InfrastructureData,
			},
		},
		testnet: testnet,
	}
}

// Setup generates the docker-compose file and write it to disk, erroring if
// any of these operations fail. It writes.
func (p *Provider) Setup() error {
	def := ComposeDef{
		NetworkName: p.testnet.Name,
		NetworkCIDR: p.testnet.IP.String(),
		Nodes:       p.testnet.Nodes,
		OmniEVMs:    p.testnet.OmniEVMs,
		Anvils:      p.testnet.AnvilChains,
		Relayer:     true,
	}

	bz, err := GenerateComposeFile(def)
	if err != nil {
		return errors.Wrap(err, "generate compose file")
	}

	err = os.WriteFile(filepath.Join(p.Testnet.Dir, "docker-compose.yml"), bz, 0o644)
	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) StartNodes(ctx context.Context, nodes ...*e2e.Node) error {
	var err error
	p.servicesOnce.Do(func() {
		svcs := additionalServices(p.testnet)
		log.Info(ctx, "Starting additional services", "names", svcs)
		err = cmtdocker.ExecCompose(ctx, p.Testnet.Dir, append([]string{"up", "-d"}, svcs...)...)
	})
	if err != nil {
		return errors.Wrap(err, "start additional services")
	}

	if err := p.Provider.StartNodes(ctx, nodes...); err != nil {
		return errors.Wrap(err, "start nodes")
	}

	return nil
}

type ComposeDef struct {
	NetworkName string
	NetworkCIDR string

	Nodes    []*e2e.Node
	OmniEVMs []types.OmniEVM
	Anvils   []types.AnvilChain
	Relayer  bool
}

func GenerateComposeFile(def ComposeDef) ([]byte, error) {
	tmpl, err := template.New("compose").Parse(string(composeTmpl))
	if err != nil {
		return nil, errors.Wrap(err, "parse template")
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, makeTemplateData(def))
	if err != nil {
		return nil, errors.Wrap(err, "execute template")
	}

	return buf.Bytes(), nil
}

type data struct {
	NetworkName  string
	NetworkCIDR  string
	Nodes        []*e2e.Node
	NodeOmniEVMs map[string]string // Maps node name to Omni EVM instance name.
	Anvils       []chain
	OmniEVMs     []chain
	Relayer      bool
}

type chain struct {
	Name       string
	ChainID    uint64
	InternalIP string
	ProxyPort  uint32

	// Only for Geth OmniEVM
	NodeKeyHex string // --nodekey: P2P node key file
	BootNodes  string // --bootnodes: Comma separated enode URLs for P2P discovery bootstrap
}

func makeTemplateData(def ComposeDef) data {
	var omniEVMs []chain
	for _, omniEVM := range def.OmniEVMs {
		var bootnodes []string
		for _, b := range omniEVM.BootNodes {
			bootnodes = append(bootnodes, b.String())
		}

		omniEVMs = append(omniEVMs, chain{
			Name:       omniEVM.InstanceName,
			ChainID:    omniEVM.Chain.ID,
			ProxyPort:  omniEVM.ProxyPort,
			InternalIP: omniEVM.InternalIP.String(),
			NodeKeyHex: hex.EncodeToString(crypto.FromECDSA(omniEVM.NodeKey)),
			BootNodes:  strings.Join(bootnodes, ","),
		})
	}

	nodeEVMs := make(map[string]string)
	for i, node := range def.Nodes {
		evm := omniEVMs[0].Name
		if len(omniEVMs) == len(def.Nodes) {
			evm = omniEVMs[i].Name
		}
		nodeEVMs[node.Name] = evm
	}

	var anvils []chain
	for _, anvil := range def.Anvils {
		anvils = append(anvils, chain{
			Name:       anvil.Chain.Name,
			ChainID:    anvil.Chain.ID,
			ProxyPort:  anvil.ProxyPort,
			InternalIP: anvil.InternalIP.String(),
		})
	}

	return data{
		NetworkName:  def.NetworkName,
		NetworkCIDR:  def.NetworkCIDR,
		Nodes:        def.Nodes,
		Anvils:       anvils,
		OmniEVMs:     omniEVMs,
		NodeOmniEVMs: nodeEVMs,
		Relayer:      def.Relayer,
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

	resp = append(resp, "relayer")

	return resp
}
