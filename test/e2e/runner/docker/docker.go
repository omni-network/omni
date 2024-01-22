package docker

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/omni-network/omni/lib/errors"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
	cmtdocker "github.com/cometbft/cometbft/test/e2e/pkg/infra/docker"

	_ "embed"
)

// composeTmpl is our own custom docker compose template. This differs from cometBFT's.
//
//go:embed compose.yaml.tmpl
var composeTmpl []byte

var _ infra.Provider = Provider{}

// Provider wraps the cometBFT docker provider, writing a different compose file.
type Provider struct {
	*cmtdocker.Provider
}

// NewProvider returns a new Provider.
func NewProvider(testnet *e2e.Testnet, data e2e.InfrastructureData) Provider {
	return Provider{
		Provider: &cmtdocker.Provider{
			ProviderData: infra.ProviderData{
				Testnet:            testnet,
				InfrastructureData: data,
			},
		},
	}
}

// Setup generates the docker-compose file and write it to disk, erroring if
// any of these operations fail. It writes.
func (p Provider) Setup() error {
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
