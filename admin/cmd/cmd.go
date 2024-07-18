package cmd

import (
	"log/slog"

	"github.com/omni-network/omni/admin/app"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	logCfg := log.DefaultConfig()
	logCfg.Level = slog.LevelDebug.String()
	logCfg.Color = log.ColorForce

	cfg := app.DefaultConfig()

	cmd := libcmd.NewRootCmd("admin", "Network admin commands")
	cmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		ctx := cmd.Context()
		if _, err := log.Init(ctx, logCfg); err != nil {
			return err
		}

		return libcmd.LogFlags(ctx, cmd.Flags())
	}

	bindCommmonFlags(cmd.PersistentFlags(), &cfg)
	log.BindFlags(cmd.PersistentFlags(), &logCfg)

	cmd.AddCommand(
		newPausePortalCmd(&cfg),
		newUnpausePortalCmd(&cfg),
		newUpgradePortalCmd(&cfg),
	)

	return cmd
}

func newPausePortalCmd(cfg *app.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause-portal",
		Short: "Pause a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.PausePortal(cmd.Context(), *cfg)
		},
	}

	return cmd
}

func newUnpausePortalCmd(cfg *app.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unpause-portal",
		Short: "Unpause a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.UnpausePortal(cmd.Context(), *cfg)
		},
	}

	return cmd
}

func newUpgradePortalCmd(cfg *app.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-portal",
		Short: "Upgrade a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.UpgradePortal(cmd.Context(), *cfg)
		},
	}

	return cmd
}
