package docker

import (
	"bytes"
	"context"
	"fmt"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/test/e2e/netman"
	"os"
	"path/filepath"
	"strings"
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
	mngr         netman.Manager
}

// NewProvider returns a new Provider.
func NewProvider(testnet *e2e.Testnet, data e2e.InfrastructureData, mngr netman.Manager) *Provider {
	return &Provider{
		Provider: &cmtdocker.Provider{
			ProviderData: infra.ProviderData{
				Testnet:            testnet,
				InfrastructureData: data,
			},
		},
		mngr: mngr,
	}
}

func (p *Provider) InternalNetwork() netconf.Network {
	network := p.mngr.Network()
	for i, c := range network.Chains {
		if c.IsPublic {
			continue // Public chains are always accessed the same as per original RPCAddress.
		}

		// Docker internal networking uses app name as hostname. Ports are always the same.
		network.Chains[i].RPCURL = fmt.Sprintf("http://%s:8584", c.Name)
		network.Chains[i].AuthRPCURL = fmt.Sprintf("http://%s:8551", c.Name)
	}

	return network
}

const startRPCPort = 8000
const startAuthRPCPort = 9000

func (p *Provider) ExternalNetwork() netconf.Network {
	network := p.mngr.Network()
	for i, c := range network.Chains {
		if c.IsPublic {
			continue // Public chains are always accessed the same as per original RPCAddress.
		}

		// Docker external networking uses localhost and a unique port.
		network.Chains[i].RPCURL = fmt.Sprintf("http://localhost:%d", startRPCPort+i)
		network.Chains[i].AuthRPCURL = fmt.Sprintf("http://localhost:%d", startAuthRPCPort+i)
	}

	return network
}

// Setup generates the docker-compose file and write it to disk, erroring if
// any of these operations fail. It writes.
func (p *Provider) Setup() error {
	tmpl, err := template.New("").Parse(string(composeTmpl))
	if err != nil {
		return errors.Wrap(err, "parsing compose template")
	}

	data := makeTemplateData(p.Testnet, p.ExternalNetwork())

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
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
		log.Info(ctx, "Starting additional services", "names", p.mngr.AdditionalService())
		err = cmtdocker.ExecCompose(ctx, p.Testnet.Dir, append([]string{"up", "-d"}, p.mngr.AdditionalService()...)...)
	})
	if err != nil {
		return errors.Wrap(err, "start additional services")
	}

	return p.Provider.StartNodes(ctx, nodes...)
}

type data struct {
	*e2e.Testnet
	Anvils   []chain
	OmniEVMs []chain
}

type chain struct {
	Name            string
	ChainID         uint64
	InternalIP      string
	HostRPCPort     string
	HostAuthRPCPort string // Only applicable to OmniEVM
}

func makeTemplateData(testnet *e2e.Testnet, extNetwork netconf.Network) data {
	var anvils []chain
	var omniEVMs []chain

	for _, c := range extNetwork.Chains {
		chain := chain{
			Name:            c.Name,
			ChainID:         c.ID,
			HostRPCPort:     port(c.RPCURL),
			HostAuthRPCPort: port(c.AuthRPCURL),
		}
		if c.IsPublic {
			continue
		} else if c.IsOmni {
			omniEVMs = append(omniEVMs, chain)
		} else {
			anvils = append(anvils, chain)
		}
	}

	return data{
		Testnet:  testnet,
		Anvils:   anvils,
		OmniEVMs: omniEVMs,
	}
}

// port returns the port from an address string if it exists.
func port(addr string) string {
	split := strings.Split(addr, ":")
	if len(split) != 3 {
		return ""
	}

	return split[2]
}
