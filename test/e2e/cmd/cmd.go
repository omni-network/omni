package cmd

import (
	"log/slog"

	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/app"

	cmtdocker "github.com/cometbft/cometbft/test/e2e/pkg/infra/docker"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	logCfg := log.DefaultConfig()
	logCfg.Level = slog.LevelDebug.String()

	defCfg := app.DefaultDefinitionConfig()

	var def app.Definition

	cmd := libcmd.NewRootCmd("e2e", "e2e network generator and test runner")
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if _, err := log.Init(cmd.Context(), logCfg); err != nil {
			return err
		}

		var err error
		def, err = app.MakeDefinition(defCfg)

		return err
	}

	bindDefFlags(cmd.PersistentFlags(), &defCfg)
	log.BindFlags(cmd.PersistentFlags(), &logCfg)

	// Root command runs the full E2E test.
	e2eTestCfg := app.DefaultE2ETestConfig()
	bindE2EFlags(cmd.Flags(), &e2eTestCfg)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		if err := libcmd.LogFlags(ctx, cmd.Flags()); err != nil {
			return err
		}

		return app.E2ETest(ctx, def, e2eTestCfg)
	}

	// Add subcommands
	cmd.AddCommand(
		newDeployCmd(&def),
		newLogsCmd(&def),
		newCleanCmd(&def),
	)

	return cmd
}

func newDeployCmd(def *app.Definition) *cobra.Command {
	var promSecrets app.PromSecrets
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploys the e2e network",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := libcmd.LogFlags(ctx, cmd.Flags()); err != nil {
				return err
			}

			return app.Deploy(ctx, *def, promSecrets)
		},
	}

	bindPromFlags(cmd.Flags(), &promSecrets)

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
			return app.Cleanup(cmd.Context(), def.Testnet.Testnet)
		},
	}
}
