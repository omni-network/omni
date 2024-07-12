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

	cmd.AddCommand(newPausePortalCmd(def))

	return cmd
}

func newPausePortalCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultPausePortalConfig()

	cmd := &cobra.Command{
		Use:   "pause-portal",
		Short: "Pause a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.PausePortal(cmd.Context(), *def, cfg)
		},
	}

	bindPausePortalFlags(cmd.Flags(), &cfg)

	return cmd
}
