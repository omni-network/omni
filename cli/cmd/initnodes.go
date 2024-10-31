package cmd

import (
	"bytes"
	"context"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/omni-network/omni/e2e/app/geth"
	halocmd "github.com/omni-network/omni/halo/cmd"
	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/lib/buildinfo"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	cmtconfig "github.com/cometbft/cometbft/config"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	cmtos "github.com/cometbft/cometbft/libs/os"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/p2p/enode"

	"github.com/spf13/cobra"

	_ "embed"
)

const (
	gethVerbosityInfo  = 3 // Geth log level (1=error,2=warn,3=info,4=debug,5=trace)
	gethVerbosityDebug = 4
)

type InitConfig struct {
	Network netconf.ID
	Home    string
	Moniker string
	Clean   bool
	Archive bool
	Debug   bool
	HaloTag string
}

func (c InitConfig) Verify() error {
	return c.Network.Verify()
}

//go:embed compose.yaml.tpl
var composeTpl []byte

func newInitCmd() *cobra.Command {
	var cfg InitConfig

	cmd := &cobra.Command{
		Use:   "init-nodes",
		Short: "Initializes omni consensus and execution nodes",
		Long:  `Initializes omni consensus node (halo) and execution node (geth) files and configuration in order to join the Omni mainnet or testnet as a full node`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cfg.Verify(); err != nil {
				return errors.Wrap(err, "verify flags")
			}

			err := InitNodes(cmd.Context(), cfg)
			if err != nil {
				return errors.Wrap(err, "init-node")
			}

			return nil
		},
	}

	bindInitConfig(cmd, &cfg)

	return cmd
}

func InitNodes(ctx context.Context, cfg InitConfig) error {
	if cfg.Network == "" {
		return errors.New("required flag --network not set")
	} else if cfg.Moniker == "" {
		return errors.New("required flag --moniker not set")
	}

	if !filepath.IsAbs(cfg.Home) {
		absPath, err := filepath.Abs(cfg.Home)
		if err != nil {
			return errors.Wrap(err, "convert path")
		}
		cfg.Home = absPath
	}

	if cfg.Home == "" {
		var err error
		cfg.Home, err = homeDir(cfg.Network)
		if err != nil {
			return err
		}
	}

	if cfg.Clean {
		log.Info(ctx, "Deleting home since --clean=true", "path", cfg.Home)
		if err := os.RemoveAll(cfg.Home); err != nil {
			return errors.Wrap(err, "clean home")
		}
	}

	if err := maybeDownloadGenesis(ctx, cfg.Network); err != nil {
		return errors.Wrap(err, "download genesis")
	}

	err := gethInit(ctx, cfg, filepath.Join(cfg.Home, "geth"))
	if err != nil {
		return errors.Wrap(err, "init geth")
	}

	logLevel := log.LevelInfo
	if cfg.Debug {
		logLevel = log.LevelDebug
	}

	err = halocmd.InitFiles(ctx, halocmd.InitConfig{
		HomeDir:     filepath.Join(cfg.Home, "halo"),
		Moniker:     cfg.Moniker,
		Network:     cfg.Network,
		TrustedSync: !cfg.Archive, // Don't state sync if archive
		AddrBook:    true,
		HaloCfgFunc: func(haloCfg *halocfg.Config) {
			haloCfg.EngineEndpoint = "http://omni_evm:8551"
			haloCfg.EngineJWTFile = "/geth/jwtsecret"
			if cfg.Archive {
				haloCfg.PruningOption = "nothing"
				// Setting this to 0 retains all blocks
				haloCfg.MinRetainBlocks = 0
			}
		},
		CometCfgFunc: func(cmtCfg *cmtconfig.Config) {
			cmtCfg.LogLevel = logLevel
			cmtCfg.Instrumentation.Prometheus = true
			if cfg.Archive {
				if cmtCfg.P2P.PersistentPeers != "" {
					cmtCfg.P2P.PersistentPeers += ","
				}
				cmtCfg.P2P.PersistentPeers += strings.Join(cfg.Network.Static().ConsensusArchives(), ",")
			}
		},
		LogCfgFunc: func(logCfg *log.Config) {
			logCfg.Color = log.ColorForce
			logCfg.Level = logLevel
		},
	})
	if err != nil {
		return errors.Wrap(err, "init halo")
	}

	err = writeComposeFile(ctx, cfg)
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
	cprov := cprovider.NewABCI(rpcCl, network)

	execution, consensus, err := cprov.GenesisFiles(ctx)
	if err != nil {
		return errors.Wrap(err, "fetching genesis files")
	} else if len(execution) == 0 {
		return errors.New("empty execution genesis file downloaded", "server", rpcServer)
	}

	log.Info(ctx, "Downloaded genesis files", "execution", len(execution), "consensus", len(consensus), "rpc", rpcServer)

	return netconf.SetEphemeralGenesis(network, execution, consensus)
}

func writeComposeFile(ctx context.Context, cfg InitConfig) error {
	composeFile := filepath.Join(cfg.Home, "compose.yaml")

	if cmtos.FileExists(composeFile) {
		log.Info(ctx, "Found existing compose file", "path", composeFile)
		return nil
	}

	tmpl, err := template.New("compose").Parse(string(composeTpl))
	if err != nil {
		return errors.Wrap(err, "parse template")
	}

	if cfg.HaloTag == "" {
		gitCommit, ok := buildinfo.GitCommit()
		if !ok {
			return errors.New("missing git commit (go install first?)")
		}
		cfg.HaloTag = gitCommit
	}

	verbosity := gethVerbosityInfo
	if cfg.Debug {
		verbosity = gethVerbosityDebug
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, struct {
		HaloTag       string
		GethTag       string
		GethVerbosity int
		GethArchive   bool
	}{
		HaloTag:       cfg.HaloTag,
		GethTag:       geth.Version,
		GethVerbosity: verbosity,
		GethArchive:   cfg.Archive,
	})
	if err != nil {
		return errors.Wrap(err, "execute template")
	}

	if err := os.WriteFile(composeFile, buf.Bytes(), 0o644); err != nil {
		return errors.Wrap(err, "writing compose file")
	}

	log.Info(ctx, "Generated docker compose file", "path", filepath.Join(cfg.Home, "compose.yaml"), "geth_version", geth.Version, "halo_version", cfg.HaloTag)

	return nil
}

func gethInit(ctx context.Context, cfg InitConfig, dir string) error {
	log.Info(ctx, "Initializing geth", "path", dir)

	// Create the dir, ensuring it doesn't already exist
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return errors.Wrap(err, "creating directory")
	}

	// Write genesis.json file
	{
		genesisFile := filepath.Join(dir, "genesis.json")
		if cmtos.FileExists(genesisFile) {
			log.Info(ctx, "Found existing execution genesis file", "path", genesisFile)
		} else {
			genesisJSON := cfg.Network.Static().ExecutionGenesisJSON
			if len(genesisJSON) == 0 {
				return errors.New("genesis json is empty for network", "network", cfg.Network)
			}
			if err := os.WriteFile(genesisFile, genesisJSON, 0o644); err != nil {
				return errors.Wrap(err, "writing genesis file", "network", cfg.Network)
			}

			log.Info(ctx, "Generated geth genesis", "path", genesisFile)
		}
	}

	// Write config.toml file
	{
		configFile := filepath.Join(dir, "config.toml")
		if cmtos.FileExists(configFile) {
			log.Info(ctx, "Found existing geth config file", "path", configFile)
		} else {
			var bootnodes []*enode.Node
			for _, seed := range cfg.Network.Static().ExecutionSeeds() {
				node, err := enode.ParseV4(seed)
				if err != nil {
					return errors.Wrap(err, "parsing seed", "seed", seed)
				}
				bootnodes = append(bootnodes, node)
			}
			gethCfg := geth.Config{
				Moniker:      cfg.Moniker,
				ChainID:      cfg.Network.Static().OmniExecutionChainID,
				IsArchive:    cfg.Archive,
				BootNodes:    bootnodes,
				TrustedNodes: nil,
			}
			if err := geth.WriteConfigTOML(gethCfg, configFile); err != nil {
				return errors.Wrap(err, "writing config.toml", "network", cfg.Network)
			}

			log.Info(ctx, "Generated geth config", "path", configFile)
		}
	}

	// Write jwtsecret file
	{
		secretFile := filepath.Join(dir, "geth", "jwtsecret")
		if cmtos.FileExists(secretFile) {
			log.Info(ctx, "Found existing geth jwtsecret file", "path", secretFile)
		} else {
			secret := hex.EncodeToString(k1.GenPrivKey().Bytes())
			if err := os.MkdirAll(filepath.Dir(secretFile), 0o755); err != nil {
				return errors.Wrap(err, "creating geth jwtsecret directory", "path", secretFile)
			}
			if err := os.WriteFile(secretFile, []byte(secret), 0o666); err != nil {
				return errors.Wrap(err, "writing geth jwtsecret", "path", secretFile)
			}

			log.Info(ctx, "Generated geth jwtsecret", "path", secretFile)
		}
	}

	// Run geth init via docker
	{
		image := "ethereum/client-go:" + geth.Version
		stateScheme := "path"
		if cfg.Archive {
			stateScheme = "hash"
		}
		dockerArgs := []string{"run",
			"-v", dir + ":/geth",
			image, "--",
			"init",
			"--datadir=/geth",
			"--state.scheme=" + stateScheme,
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
