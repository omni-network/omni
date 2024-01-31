package app

import (
	"path/filepath"
	"strings"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/test/e2e/docker"
	"github.com/omni-network/omni/test/e2e/netman"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
)

const infraDocker = "docker"

// DefinitionConfig is the configuration required to create a full Definition.
type DefinitionConfig struct {
	ManifestFile  string
	InfraProvider string

	// Secrets (not required for devnet)
	DeployKeyFile  string
	RelayerKeyFile string
}

// DefaultDefinitionConfig returns a default configuration for a Definition.
func DefaultDefinitionConfig() DefinitionConfig {
	return DefinitionConfig{
		InfraProvider: infraDocker,
	}
}

// Definition defines a e2e network. All (sub)commands of the e2e cli requires a definition operate.
// Armed with a definition, a e2e network can be deployed, started, tested, stopped, etc.
type Definition struct {
	Testnet *e2e.Testnet // Note that testnet is the cometBFT term.
	Infra   infra.Provider
	Netman  netman.Manager
}

func MakeDefinition(cfg DefinitionConfig) (Definition, error) {
	manifest, err := LoadManifest(cfg.ManifestFile)
	if err != nil {
		return Definition{}, errors.Wrap(err, "loading manifest")
	}

	mngr, err := netman.NewManager(manifest.Network, cfg.DeployKeyFile, cfg.RelayerKeyFile)
	if err != nil {
		return Definition{}, errors.Wrap(err, "get network")
	}

	ifd, err := e2e.NewDockerInfrastructureData(manifest.Manifest)
	if err != nil {
		return Definition{}, errors.Wrap(err, "creating docker infrastructure data")
	}

	testnet, err := e2e.NewTestnetFromManifest(manifest.Manifest, cfg.ManifestFile, ifd)
	if err != nil {
		return Definition{}, errors.Wrap(err, "loading testnet")
	}

	return Definition{
		Testnet: adaptTestnet(testnet),
		Infra:   docker.NewProvider(testnet, ifd, mngr.AdditionalService()),
		Netman:  mngr,
	}, nil
}

func adaptTestnet(testnet *e2e.Testnet) *e2e.Testnet {
	testnet.Dir = runsDir(testnet.File)
	testnet.VoteExtensionsEnableHeight = 1
	testnet.UpgradeVersion = "omniops/halo:main"
	for i := range testnet.Nodes {
		testnet.Nodes[i] = adaptNode(testnet.Nodes[i])
	}

	return testnet
}

func adaptNode(node *e2e.Node) *e2e.Node {
	node.Version = "omniops/halo:main"
	node.PrivvalKey = k1.GenPrivKey()

	return node
}

// runsDir returns the runs directory for a given manifest file.
// E.g. /path/to/manifests/manifest.toml > /path/to/runs/manifest.
func runsDir(manifestFile string) string {
	resp := strings.TrimSuffix(manifestFile, filepath.Ext(manifestFile))
	return strings.Replace(resp, "manifests", "runs", 1)
}
