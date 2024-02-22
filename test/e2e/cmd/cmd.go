package cmd

import (
	"log/slog"

	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/app"
	"github.com/omni-network/omni/test/e2e/types"

	cmtdocker "github.com/cometbft/cometbft/test/e2e/pkg/infra/docker"

	"github.com/spf13/cobra"
)

const (
	defaultPingPongDeploy uint64 = 1000
)

func New() *cobra.Command {
	// E2E app is aimed at devs and CI, so debug level and force colors by default.
	logCfg := log.DefaultConfig()
	logCfg.Level = slog.LevelDebug.String()
	logCfg.Color = log.ColorForce

	defCfg := app.DefaultDefinitionConfig()

	var def app.Definition
	var depCfg app.DeployConfig

	cmd := libcmd.NewRootCmd("e2e", "e2e network generator and test runner")
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if _, err := log.Init(cmd.Context(), logCfg); err != nil {
			return err
		}

		if err := libcmd.LogFlags(cmd.Context(), cmd.Flags()); err != nil {
			return err
		}

		var err error
		def, err = app.MakeDefinition(defCfg)

		return err
	}

	bindDefFlags(cmd.PersistentFlags(), &defCfg)
	bindDeployFlags(cmd.PersistentFlags(), &depCfg)
	log.BindFlags(cmd.PersistentFlags(), &logCfg)

	// Root command runs the full E2E test.
	e2eTestCfg := app.DefaultE2ETestConfig()
	bindE2EFlags(cmd.Flags(), &e2eTestCfg)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return app.E2ETest(cmd.Context(), def, e2eTestCfg, depCfg)
	}

	// Add subcommands
	cmd.AddCommand(
		newDeployCmd(&def),
		newLogsCmd(&def),
		newCleanCmd(&def),
		newTestCmd(&def),
		newUpgradeCmd(&def),
	)

	return cmd
}

func newDeployCmd(def *app.Definition) *cobra.Command {
	var cfg app.DeployConfig

	pingPongN := defaultPingPongDeploy // Default to 1000 ping pongs.
	cfg.PingPongN = &pingPongN

	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploys the e2e network",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := app.Deploy(cmd.Context(), *def, cfg)
			return err
		},
	}

	bindDeployFlags(cmd.Flags(), &cfg)
	cmd.Flags().Uint64Var(&pingPongN, "ping-pong", pingPongN, "Number of ping pongs messages to send. 0 disables it")

	return cmd
}

func newLogsCmd(def *app.Definition) *cobra.Command {
	return &cobra.Command{
		Use:   "logs",
		Short: "Prints the infrastructure logs (of a previously preserved network)",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmtdocker.ExecComposeVerbose(cmd.Context(), def.Testnet.Dir, "logs")
		},
	}
}

func newCleanCmd(def *app.Definition) *cobra.Command {
	return &cobra.Command{
		Use:   "clean",
		Short: "Cleans (deletes) previously preserved network infrastructure",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.Cleanup(cmd.Context(), *def)
		},
	}
}

func newTestCmd(def *app.Definition) *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Runs go tests against the a previously preserved network",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.Test(cmd.Context(), *def, types.DeployInfos{}, true)
		},
	}
}

func newUpgradeCmd(def *app.Definition) *cobra.Command {
	return &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrades docker containers of a previously preserved network",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return def.Infra.Upgrade(cmd.Context())
		},
	}
}
