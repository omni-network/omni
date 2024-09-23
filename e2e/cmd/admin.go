package cmd

import (
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/admin"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

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
		newPauseXCallCmd(def),
		newUnpauseXCallCmd(def),
		newPauseXSubmitCmd(def),
		newUnpauseXSubmitCmd(def),
		newAllowValidatorsCmd(def),
		newAdminTestCmd(def),
	)

	return cmd
}

func newAllowValidatorsCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allow-operators",
		Short: "Ensure all operators are allowed as validators",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.AllowOperators(cmd.Context(), *def)
		},
	}

	return cmd
}

func newPausePortalCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultConfig()

	cmd := &cobra.Command{
		Use:   "pause-portal",
		Short: "Pause a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.PausePortal(cmd.Context(), *def, cfg)
		},
	}

	bindAdminFlags(cmd.Flags(), &cfg)

	return cmd
}

func newUnpausePortalCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultConfig()

	cmd := &cobra.Command{
		Use:   "unpause-portal",
		Short: "Unpause a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UnpausePortal(cmd.Context(), *def, cfg)
		},
	}

	bindAdminFlags(cmd.Flags(), &cfg)

	return cmd
}

func newUpgradePortalCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultConfig()

	cmd := &cobra.Command{
		Use:   "upgrade-portal",
		Short: "Upgrade a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradePortal(cmd.Context(), *def, cfg)
		},
	}

	bindAdminFlags(cmd.Flags(), &cfg)

	return cmd
}

func newPauseXCallCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultConfig()
	xcallCfg := admin.XCallConfig{}

	cmd := &cobra.Command{
		Use:   "pause-xcall",
		Short: "Pause cross-chain calls",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.PauseXCall(cmd.Context(), *def, cfg, xcallCfg)
		},
	}

	bindAdminFlags(cmd.Flags(), &cfg)
	bindAdminXCallFlags(cmd.Flags(), &xcallCfg)

	return cmd
}

func newUnpauseXCallCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultConfig()
	xcallCfg := admin.XCallConfig{}

	cmd := &cobra.Command{
		Use:   "unpause-xcall",
		Short: "Unpause cross-chain calls",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UnpauseXCall(cmd.Context(), *def, cfg, xcallCfg)
		},
	}

	bindAdminFlags(cmd.Flags(), &cfg)
	bindAdminXCallFlags(cmd.Flags(), &xcallCfg)

	return cmd
}

func newPauseXSubmitCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultConfig()
	xsubCfg := admin.XSubmitConfig{}

	cmd := &cobra.Command{
		Use:   "pause-xsubmit",
		Short: "Pause cross-chain submissions",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.PauseXSubmit(cmd.Context(), *def, cfg, xsubCfg)
		},
	}

	bindAdminFlags(cmd.Flags(), &cfg)
	bindAdminXSubmitFlags(cmd.Flags(), &xsubCfg)

	return cmd
}

func newUnpauseXSubmitCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultConfig()
	xsubCfg := admin.XSubmitConfig{}

	cmd := &cobra.Command{
		Use:   "unpause-xsubmit",
		Short: "Unpause cross-chain submissions",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UnpauseXSubmit(cmd.Context(), *def, cfg, xsubCfg)
		},
	}

	bindAdminFlags(cmd.Flags(), &cfg)
	bindAdminXSubmitFlags(cmd.Flags(), &xsubCfg)

	return cmd
}

func newAdminTestCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Test contract admin commands",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			if def.Testnet.Network != netconf.Devnet {
				return errors.New("only devnet")
			}

			if _, err := app.Deploy(ctx, *def, app.DeployConfig{
				PingPongN: 0,
				PingPongP: 0,
				PingPongL: 0}); err != nil {
				return errors.Wrap(err, "deploy")
			}

			if err := admin.Test(ctx, *def); err != nil {
				return err
			}

			return app.CleanInfra(ctx, *def)
		},
	}

	return cmd
}
