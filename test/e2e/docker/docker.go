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
	services     []string
}

// NewProvider returns a new Provider.
func NewProvider(testnet *e2e.Testnet, data e2e.InfrastructureData, services []string) *Provider {
	return &Provider{
		Provider: &cmtdocker.Provider{
			ProviderData: infra.ProviderData{
				Testnet:            testnet,
				InfrastructureData: data,
			},
		},
		services: services,
	}
}

// Setup generates the docker-compose file and write it to disk, erroring if
// any of these operations fail. It writes.
func (p *Provider) Setup() error {
	tmpl, err := template.New("").Parse(string(composeTmpl))
	if err != nil {
		return errors.Wrap(err, "parsing compose template")
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, p.Testnet); err != nil {
		return errors.Wrap(err, "executing compose template")
	}

	err = os.WriteFile(filepath.Join(p.Testnet.Dir, "docker-compose.yml"), buf.Bytes(), 0o644)
	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) StartNodes(ctx context.Context, nodes ...*e2e.Node) error {
	var err error
	p.servicesOnce.Do(func() {
		log.Info(ctx, "Starting additional services", "names", p.services)
		err = cmtdocker.ExecCompose(ctx, p.Testnet.Dir, append([]string{"up", "-d"}, p.services...)...)
	})
	if err != nil {
		return errors.Wrap(err, "start additional services")
	}

	return p.Provider.StartNodes(ctx, nodes...)
}
