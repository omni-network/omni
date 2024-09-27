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
		newEnsurePortalSpecCmd(def),
		newUpgradePortalCmd(def),
		newUpgradeFeeOracleV1Cmd(def),
		newUpgradeGasStationCmd(def),
		newUpgradeGasPumpCmd(def),
		newUpgradeStakingCmd(def),
		newUpgradeSlashingCmd(def),
		newUpgradeBridgeNativeCmd(def),
		newUpgradeBridgeL1(def),
		newUpgradePortalRegistryCmd(def),
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

func newEnsurePortalSpecCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultConfig()

	cmd := &cobra.Command{
		Use:   "ensure-portal-spec",
		Short: "Ensure live portals match local the specs, defined in e2e/app/admin/portalspec.go",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.EnsurePortalSpec(cmd.Context(), *def, cfg, nil)
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

func newUpgradeFeeOracleV1Cmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultConfig()

	cmd := &cobra.Command{
		Use:   "upgrade-fee-oracle-v1",
		Short: "Upgrade FeeOracleV1 contracts.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeFeeOracleV1(cmd.Context(), *def, cfg)
		},
	}

	bindAdminFlags(cmd.Flags(), &cfg)

	return cmd
}

func newUpgradeGasStationCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-gas-station",
		Short: "Upgrade the OmniGasStation contract.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeGasStation(cmd.Context(), *def)
		},
	}

	return cmd
}

func newUpgradeGasPumpCmd(def *app.Definition) *cobra.Command {
	cfg := admin.DefaultConfig()

	cmd := &cobra.Command{
		Use:   "upgrade-gas-pump",
		Short: "Upgrade OmniGasPump contracts.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeGasPump(cmd.Context(), *def, cfg)
		},
	}

	bindAdminFlags(cmd.Flags(), &cfg)

	return cmd
}

func newUpgradeStakingCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-staking",
		Short: "Upgrade the Staking predeploy.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeStaking(cmd.Context(), *def)
		},
	}

	return cmd
}

func newUpgradeSlashingCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-slashing",
		Short: "Upgrade the Slashing predeploy.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeSlashing(cmd.Context(), *def)
		},
	}

	return cmd
}

func newUpgradeBridgeNativeCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-bridge-native",
		Short: "Upgrade the OmniBridgeNative predeploy.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeBridgeNative(cmd.Context(), *def)
		},
	}

	return cmd
}

func newUpgradeBridgeL1(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-bridge-l1",
		Short: "Upgrade the OmniBridgeL1 contract.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeBridgeL1(cmd.Context(), *def)
		},
	}

	return cmd
}

func newUpgradePortalRegistryCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-portal-registry",
		Short: "Upgrade the PortalRegistry predeploy.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradePortalRegistry(cmd.Context(), *def)
		},
	}

	return cmd
}

func newAdminTestCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Test contract admin commands",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			if !def.Testnet.Network.IsEphemeral() {
				return errors.New("only ephemeral networks")
			}

			// deploy devnet, but not staging
			if def.Testnet.Network == netconf.Devnet {
				if _, err := app.Deploy(ctx, *def, app.DeployConfig{
					PingPongN: 0,
					PingPongP: 0,
					PingPongL: 0}); err != nil {
					return errors.Wrap(err, "deploy")
				}
			}

			if err := admin.Test(ctx, *def); err != nil {
				return err
			}

			return app.CleanInfra(ctx, *def)
		},
	}

	return cmd
}
