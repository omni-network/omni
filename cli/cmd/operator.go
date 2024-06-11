package cmd

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/omni-network/omni/e2e/app/geth"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/p2p/enode"

	"github.com/spf13/cobra"
)

type initConfig struct {
	Network netconf.ID
	Home    string
	Moniker string
}

func newInitCmd() *cobra.Command {
	var cfg initConfig

	cmd := &cobra.Command{
		Use:   "init-nodes",
		Short: "Initializes omni consensus and execution nodes",
		Long:  `Initializes omni consensus node (halo) and execution node (geth) files and configuration in order to join the Omni mainnet or testnet as a full node`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
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

	// TODO(corver): Init halo

	return gethInit(ctx, cfg.Network, filepath.Join(cfg.Home, "geth"), cfg.Moniker)
}

func gethInit(ctx context.Context, network netconf.ID, dir string, moniker string) error {
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
	}

	// Run geth init via docker
	{
		image := "ethereum/client-go:" + geth.Version
		//nolint:gosec // Command not "tainted"
		cmd := exec.CommandContext(ctx,
			"docker", "run",
			"-v", dir+":/geth",
			image, "--",
			"--datadir=/geth",
			"init",
			"/geth/genesis.json")
		cmd.Dir = dir

		out, err := cmd.CombinedOutput()
		if err != nil {
			return errors.Wrap(err, "docker run geth init", "output", string(out))
		}
	}

	return nil
}
