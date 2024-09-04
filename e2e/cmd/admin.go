package cmd

import (
	"github.com/omni-network/omni/e2e/app"

	"github.com/spf13/cobra"
)

func newAdminCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "admin",
		Short: "Network admin commands",
	}

	cmd.AddCommand(
		newPortalCmd(def),
		newStakingCmd(def),
		newAdminTestCmd(def),
	)

	return cmd
}

func newPortalCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "portal",
		Short: "OmniPortal.sol admin commands",
	}

	cfg := app.DefaultPortalAdminConfig()
	bindPortalAdminFlags(cmd.PersistentFlags(), &cfg)

	cmd.AddCommand(
		newPortalPauseCmd(def, &cfg),
		newPortalUnpauseCmd(def, &cfg),
		newPortalUpgradeCmd(def, &cfg),
	)

	return cmd
}

func newPortalPauseCmd(def *app.Definition, cfg *app.PortalAdminConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause",
		Short: "Pause a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.PausePortal(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newPortalUnpauseCmd(def *app.Definition, cfg *app.PortalAdminConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unpause",
		Short: "Unpause a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.UnpausePortal(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newPortalUpgradeCmd(def *app.Definition, cfg *app.PortalAdminConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.UpgradePortal(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newStakingCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "staking",
		Short: "Staking.sol admin commands",
	}

	cmd.AddCommand(
		newStakingUpgradeCmd(def),
		newStakingConfigureCmd(def),
	)

	return cmd
}

func newStakingUpgradeCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade a staking contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.UpgradeStaking(cmd.Context(), *def)
		},
	}

	return cmd
}

func newStakingConfigureCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Update on-chain Staking config, if it differs from config defined in lib/contracts/staking",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.ConfigureStaking(cmd.Context(), *def)
		},
	}

	return cmd
}

func newAdminTestCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Test all admin commands",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.TestAdminCommands(cmd.Context(), *def)
		},
	}

	return cmd
}
