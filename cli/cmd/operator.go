package cmd

import (
	"bytes"
	"context"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/omni-network/omni/e2e/app/geth"
	halocmd "github.com/omni-network/omni/halo/cmd"
	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/lib/buildinfo"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	cmtconfig "github.com/cometbft/cometbft/config"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/p2p/enode"

	"github.com/spf13/cobra"

	_ "embed"
)

type initConfig struct {
	Network netconf.ID
	Home    string
	Moniker string
	Clean   bool
}

func (c initConfig) Verify() error {
	return c.Network.Verify()
}

//go:embed compose.yml.tpl
var composeTpl []byte

func newInitCmd() *cobra.Command {
	var cfg initConfig

	cmd := &cobra.Command{
		Use:   "init-nodes",
		Short: "Initializes omni consensus and execution nodes",
		Long:  `Initializes omni consensus node (halo) and execution node (geth) files and configuration in order to join the Omni mainnet or testnet as a full node`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cfg.Verify(); err != nil {
				return errors.Wrap(err, "verify flags")
			}

			err := initNodes(cmd.Context(), cfg)
			if err != nil {
				return errors.Wrap(err, "init-node")
			}

			return nil
		},
	}

	bindInitConfig(cmd, &cfg)

	return cmd
}

func initNodes(ctx context.Context, cfg initConfig) error {
	if cfg.Network == "" {
		return errors.New("required flag --network not set")
	} else if cfg.Moniker == "" {
		return errors.New("required flag --moniker not set")
	}

	if cfg.Home == "" {
		var err error
		cfg.Home, err = homeDir(cfg.Network)
		if err != nil {
			return err
		}
	}

	if cfg.Clean {
		if err := os.RemoveAll(cfg.Home); err != nil {
			return errors.Wrap(err, "clean home")
		}
	}

	if files, err := filepath.Glob(cfg.Home + "/*"); err != nil {
		return errors.Wrap(err, "glob home")
	} else if len(files) > 0 {
		return &CliError{
			Msg:     "--home directory is not empty: " + cfg.Home,
			Suggest: "Use --clean flag to delete existing files (be careful!), or provide a different --home flag",
		}
	}

	if err := maybeDownloadGenesis(ctx, cfg.Network); err != nil {
		return errors.Wrap(err, "download genesis")
	}

	err := gethInit(ctx, cfg.Network, filepath.Join(cfg.Home, "geth"), cfg.Moniker)
	if err != nil {
		return errors.Wrap(err, "init geth")
	}

	err = halocmd.InitFiles(ctx, halocmd.InitConfig{
		HomeDir:     filepath.Join(cfg.Home, "halo"),
		Moniker:     cfg.Moniker,
		Network:     cfg.Network,
		TrustedSync: true,
		AddrBook:    true,
		HaloCfgFunc: func(cfg *halocfg.Config) {
			cfg.EngineEndpoint = "http://omni_evm:8551"
			cfg.EngineJWTFile = "/geth/jwtsecret"
			cfg.RPCEndpoints = xchain.RPCEndpoints{cfg.Network.Static().OmniExecutionChainName(): "http://omni_evm:8545"}
		},
		CometCfgFunc: func(cfg *cmtconfig.Config) {
			cfg.LogLevel = "info"
		},
	})
	if err != nil {
		return errors.Wrap(err, "init halo")
	}

	err = writeComposeFile(ctx, cfg.Home)
	if err != nil {
		return errors.Wrap(err, "write compose file")
	}

	return nil
}

// maybeDownloadGenesis downloads the genesis files via cprovider the network if they are not already set.
func maybeDownloadGenesis(ctx context.Context, network netconf.ID) error {
	if network.IsProtected() {
		return nil // No need to download genesis for protected networks
	}

	rpcServer := network.Static().ConsensusRPC()
	rpcCl, err := rpchttp.New(rpcServer, "/websocket")
	if err != nil {
		return errors.Wrap(err, "create rpc client")
	}
	stubNamer := func(xchain.ChainVersion) string { return "" }
	cprov := cprovider.NewABCIProvider(rpcCl, network, stubNamer)

	execution, consensus, err := cprov.GenesisFiles(ctx)
	if err != nil {
		return errors.Wrap(err, "fetching genesis files")
	} else if len(execution) == 0 {
		return errors.New("empty execution genesis file downloaded", "server", rpcServer)
	}

	log.Info(ctx, "Downloaded genesis files", "execution", len(execution), "consensus", len(consensus), "rpc", rpcServer)

	return netconf.SetEphemeralGenesis(network, execution, consensus)
}

func writeComposeFile(ctx context.Context, home string) error {
	tmpl, err := template.New("compose").Parse(string(composeTpl))
	if err != nil {
		return errors.Wrap(err, "parse template")
	}

	// TODO(corver): Replace git commit with buildinfo.Version once we release proper versions.
	commit, ok := buildinfo.GitCommit()
	if !ok {
		return errors.New("missing git commit (go install first?)")
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, struct {
		HaloTag string
		GethTag string
	}{
		HaloTag: commit,
		GethTag: geth.Version,
	})
	if err != nil {
		return errors.Wrap(err, "execute template")
	}

	if err := os.WriteFile(filepath.Join(home, "compose.yml"), buf.Bytes(), 0o644); err != nil {
		return errors.Wrap(err, "writing compose file")
	}

	log.Info(ctx, "Generated docker compose file", "path", filepath.Join(home, "compose.yml"), "geth_version", geth.Version, "halo_version", commit)

	return nil
}

func gethInit(ctx context.Context, network netconf.ID, dir string, moniker string) error {
	log.Info(ctx, "Initializing geth", "path", dir)

	// Create the dir, ensuring it doesn't already exist
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return errors.Wrap(err, "creating directory")
	}

	// Write genesis.json file
	{
		genesisJSON := network.Static().ExecutionGenesisJSON
		if len(genesisJSON) == 0 {
			return errors.New("genesis json is empty for network", "network", network)
		}
		if err := os.WriteFile(filepath.Join(dir, "genesis.json"), genesisJSON, 0o644); err != nil {
			return errors.Wrap(err, "writing genesis file", "network", network)
		}

		log.Info(ctx, "Generated geth genesis", "path", filepath.Join(dir, "genesis.json"))
	}

	// Write config.toml file
	{
		var bootnodes []*enode.Node
		for _, seed := range network.Static().ExecutionSeeds() {
			node, err := enode.ParseV4(seed)
			if err != nil {
				return errors.Wrap(err, "parsing seed", "seed", seed)
			}
			bootnodes = append(bootnodes, node)
		}
		cfg := geth.Config{
			Moniker:      moniker,
			ChainID:      network.Static().OmniExecutionChainID,
			IsArchive:    false,
			BootNodes:    bootnodes,
			TrustedNodes: nil,
		}
		if err := geth.WriteConfigTOML(cfg, filepath.Join(dir, "config.toml")); err != nil {
			return errors.Wrap(err, "writing config.toml", "network", network)
		}

		log.Info(ctx, "Generated geth config", "path", filepath.Join(dir, "config.toml"))
	}

	// Write jwtsecret file
	{
		secret := hex.EncodeToString(k1.GenPrivKey().Bytes())
		path := filepath.Join(dir, "geth", "jwtsecret")
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return errors.Wrap(err, "creating geth jwtsecret directory", "path", path)
		}
		if err := os.WriteFile(path, []byte(secret), 0o666); err != nil {
			return errors.Wrap(err, "writing geth jwtsecret", "path", path)
		}

		log.Info(ctx, "Generated geth jwtsecret", "path", path)
	}

	// Run geth init via docker
	{
		image := "ethereum/client-go:" + geth.Version
		dockerArgs := []string{"run",
			"-v", dir + ":/geth",
			image, "--",
			"init",
			"--datadir=/geth",
			"/geth/genesis.json",
		}

		cmd := exec.CommandContext(ctx, "docker", dockerArgs...)
		cmd.Dir = dir

		out, err := cmd.CombinedOutput()
		if err != nil {
			return errors.Wrap(err, "docker run geth init", "output", string(out))
		}

		log.Info(ctx, "Initialized geth chain data")
	}

	return nil
}
