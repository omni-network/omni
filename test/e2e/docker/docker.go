package docker

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/types"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
	cmtdocker "github.com/cometbft/cometbft/test/e2e/pkg/infra/docker"

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
	bz, err := generateComposeFile(p.testnet)
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

func generateComposeFile(testnet types.Testnet) ([]byte, error) {
	tmpl, err := template.New("compose").Parse(string(composeTmpl))
	if err != nil {
		return nil, errors.Wrap(err, "parse template")
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, makeTemplateData(testnet))
	if err != nil {
		return nil, errors.Wrap(err, "execute template")
	}

	return buf.Bytes(), nil
}

type data struct {
	*e2e.Testnet
	Anvils   []chain
	OmniEVMs []chain
}

type chain struct {
	Name       string
	ChainID    uint64
	InternalIP string
	ProxyPort  uint32
}

func makeTemplateData(testnet types.Testnet) data {
	var omniEVMs []chain
	for _, omniEVM := range testnet.OmniEVMs {
		omniEVMs = append(omniEVMs, chain{
			Name:       omniEVM.InstanceName,
			ChainID:    omniEVM.Chain.ID,
			ProxyPort:  omniEVM.ProxyPort,
			InternalIP: omniEVM.InternalIP.String(),
		})
	}

	var anvils []chain
	for _, anvil := range testnet.AnvilChains {
		anvils = append(anvils, chain{
			Name:       anvil.Chain.Name,
			ChainID:    anvil.Chain.ID,
			ProxyPort:  anvil.ProxyPort,
			InternalIP: anvil.InternalIP.String(),
		})
	}

	return data{
		Testnet:  testnet.Testnet,
		Anvils:   anvils,
		OmniEVMs: omniEVMs,
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
