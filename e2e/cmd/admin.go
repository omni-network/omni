package cmd

import (
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/admin"

	"github.com/spf13/cobra"
)

func newAdminCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "admin",
		Short: "Network admin commands",
	}

	cmd.AddCommand(
		newPausePortalCmd(def),
		newUnpausePortalCmd(def),
		newUpgradePortalCmd(def),
	)

	return cmd
}

func newPausePortalCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultPortalAdminConfig()

	cmd := &cobra.Command{
		Use:   "pause-portal",
		Short: "Pause a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.PausePortal(cmd.Context(), *def, cfg)
		},
	}

	bindPortalAdminFlags(cmd.Flags(), &cfg)

	return cmd
}

func newUnpausePortalCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultPortalAdminConfig()

	cmd := &cobra.Command{
		Use:   "unpause-portal",
		Short: "Unpause a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UnpausePortal(cmd.Context(), *def, cfg)
		},
	}

	bindPortalAdminFlags(cmd.Flags(), &cfg)

	return cmd
}

func newUpgradePortalCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultPortalAdminConfig()

	cmd := &cobra.Command{
		Use:   "upgrade-portal",
		Short: "Upgrade a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradePortal(cmd.Context(), *def, cfg)
		},
	}

	bindPortalAdminFlags(cmd.Flags(), &cfg)

	return cmd
}
