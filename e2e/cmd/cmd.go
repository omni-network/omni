package cmd

import (
	"context"
	"log/slog"
	"regexp"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/e2e/types"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	cmtdocker "github.com/cometbft/cometbft/test/e2e/pkg/infra/docker"

	"github.com/ethereum/go-ethereum/common"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	// E2E app is aimed at devs and CI, so debug level and force colors by default.
	logCfg := log.DefaultConfig()
	logCfg.Level = slog.LevelDebug.String()
	logCfg.Color = log.ColorForce

	defCfg := app.DefaultDefinitionConfig(context.Background())

	var def app.Definition

	cmd := libcmd.NewRootCmd("e2e", "e2e network generator and test runner")
	cmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		ctx := cmd.Context()
		if _, err := log.Init(ctx, logCfg); err != nil {
			return err
		}

		if err := libcmd.LogFlags(ctx, cmd.Flags()); err != nil {
			return err
		}

		var err error
		def, err = app.MakeDefinition(ctx, defCfg, cmd.Use)
		if err != nil {
			return errors.Wrap(err, "make definition")
		}

		// Some commands require networking, this ensures proper errors instead of panics.
		if matchAny(cmd.Use, ".*deploy.*", ".*update.*", "e2e") {
			if err := def.InitLazyNetwork(); err != nil {
				return errors.Wrap(err, "init network")
			}
		}

		return err
	}

	bindDefFlags(cmd.PersistentFlags(), &defCfg)
	log.BindFlags(cmd.PersistentFlags(), &logCfg)

	// Root command runs the full E2E test.
	e2eTestCfg := app.DefaultE2ETestConfig()
	bindE2EFlags(cmd.Flags(), &e2eTestCfg)
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		return app.E2ETest(cmd.Context(), def, e2eTestCfg)
	}

	// Add subcommands
	cmd.AddCommand(
		newCreate3DeployCmd(&def),
		newDeployCmd(&def),
		newLogsCmd(&def),
		newCleanCmd(&def),
		newTestCmd(&def),
		newUpgradeCmd(&def),
		newRestartCmd(&def),
		newKeyCreate(&def),
		newAdminCmd(&def),
		fundAccounts(&def),
	)

	return cmd
}

func matchAny(str string, patterns ...string) bool {
	for _, pattern := range patterns {
		if ok, _ := regexp.MatchString(pattern, str); ok {
			return true
		}
	}

	return false
}

func newDeployCmd(def *app.Definition) *cobra.Command {
	cfg := app.DefaultDeployConfig()

	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploys the e2e network",
		RunE: func(cmd *cobra.Command, _ []string) error {
			_, err := app.Deploy(cmd.Context(), *def, cfg)
			return err
		},
	}

	bindDeployFlags(cmd.Flags(), &cfg)

	return cmd
}

func newLogsCmd(def *app.Definition) *cobra.Command {
	return &cobra.Command{
		Use:   "logs",
		Short: "Prints the infrastructure logs (of a previously preserved network)",
		RunE: func(cmd *cobra.Command, _ []string) error {
			err := cmtdocker.ExecComposeVerbose(cmd.Context(), def.Testnet.Dir, "logs")
			if err != nil {
				return errors.Wrap(err, "executing docker-compose logs")
			}

			return nil
		},
	}
}

func newCleanCmd(def *app.Definition) *cobra.Command {
	return &cobra.Command{
		Use:   "clean",
		Short: "Cleans (deletes) previously preserved network infrastructure",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := app.CleanInfra(cmd.Context(), *def); err != nil {
				return err
			}

			return app.CleanupDir(cmd.Context(), def.Testnet.Dir)
		},
	}
}

func newTestCmd(def *app.Definition) *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Runs go tests against the a previously preserved network",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.Test(cmd.Context(), *def, true)
		},
	}
}

func newUpgradeCmd(def *app.Definition) *cobra.Command {
	cfg := app.DefaultDeployConfig()
	svcCfg := types.DefaultServiceConfig()

	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrades docker containers of a vmcompose network",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.Upgrade(cmd.Context(), *def, cfg, svcCfg)
		},
	}

	bindDeployFlags(cmd.Flags(), &cfg)
	bindServiceFlags(cmd.Flags(), &svcCfg)

	return cmd
}

func newRestartCmd(def *app.Definition) *cobra.Command {
	cfg := app.DefaultDeployConfig()
	svcCfg := types.DefaultServiceConfig()

	cmd := &cobra.Command{
		Use:   "restart",
		Short: "Restarts docker containers of a vmcompose network",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.Restart(cmd.Context(), *def, cfg, svcCfg)
		},
	}

	bindDeployFlags(cmd.Flags(), &cfg)
	bindServiceFlags(cmd.Flags(), &svcCfg)

	return cmd
}

func newCreate3DeployCmd(def *app.Definition) *cobra.Command {
	cfg := app.Create3DeployConfig{}

	cmd := &cobra.Command{
		Use:   "create3-deploy",
		Short: "Deploys the Create3 factory",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.Create3Deploy(cmd.Context(), *def, cfg)
		},
	}

	bindCreate3DeployFlags(cmd.Flags(), &cfg)

	return cmd
}

func newKeyCreate(def *app.Definition) *cobra.Command {
	cfg := key.UploadConfig{}

	cmd := &cobra.Command{
		Use:   "key-create",
		Short: "Creates a private key in GCP secret manager for a node in a manifest",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if def.Testnet.Network == netconf.Simnet || def.Testnet.Network == netconf.Devnet {
				return errors.New("cannot create keys for simnet or devnet")
			}

			cfg.Network = def.Testnet.Network

			if err := verifyKeyNodeType(*def, cfg); err != nil {
				return err
			}

			_, err := key.UploadNew(cmd.Context(), cfg)

			return err
		},
	}

	bindKeyCreateFlags(cmd, &cfg)

	return cmd
}

func fundAccounts(def *app.Definition) *cobra.Command {
	var dryRun bool

	cmd := &cobra.Command{
		Use:   "fund",
		Short: "Funds accounts to their target balance, network based on the manifest",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if def.Testnet.Network == netconf.Simnet || def.Testnet.Network == netconf.Devnet {
				return errors.New("cannot fund accounts on simnet or devnet")
			}
			if err := def.InitLazyNetwork(); err != nil {
				return errors.Wrap(err, "init network")
			}

			return app.FundEOAAccounts(cmd.Context(), *def, dryRun)
		},
	}

	cmd.Flags().BoolVar(&dryRun, "dry-run", dryRun, "Enables dry-run for testing")

	return cmd
}

// verifyKeyNodeType checks if the node exists in the manifest and if the key type is allowed for the node.
func verifyKeyNodeType(def app.Definition, cfg key.UploadConfig) error {
	if err := cfg.Type.Verify(); err != nil {
		return err
	}

	if cfg.Type == key.EOA {
		eoaRole := eoa.Role(cfg.Name)
		if err := eoaRole.Verify(); err != nil {
			return errors.Wrap(err, "verifying name as eoa type")
		}

		account, ok := eoa.AccountForRole(def.Testnet.Network, eoaRole)
		if !ok {
			return errors.New("eoa account not found", "role", eoaRole)
		}

		if account.Type != eoa.TypeSecret {
			return errors.New("cannot create eoa key for non secret account")
		}

		if account.Address != (common.Address{}) {
			return errors.New("cannot create eoa key already defined", "addr", account.Address.Hex())
		}

		return nil
	}

	for _, node := range def.Testnet.Nodes {
		if node.Name == cfg.Name {
			if cfg.Type == key.P2PExecution {
				return errors.New("cannot create execution key for halo node")
			}

			return nil
		}
	}

	for _, evm := range def.Testnet.OmniEVMs {
		if evm.InstanceName == cfg.Name {
			if cfg.Type != key.P2PExecution {
				return errors.New("only execution keys allowed for evm nodes")
			}

			return nil
		}
	}

	return errors.New("node not found", "name", cfg.Name)
}
