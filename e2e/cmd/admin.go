package cmd

import (
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/admin"
	"github.com/omni-network/omni/e2e/bridge"
	"github.com/omni-network/omni/e2e/gasstation"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/spf13/cobra"
)

func newAdminCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "admin",
		Short: "Network admin commands",
	}

	cfg := admin.DefaultConfig()
	bindAdminFlags(cmd.PersistentFlags(), &cfg)

	cmd.AddCommand(
		newEnsurePortalSpecCmd(def, &cfg),
		newEnsureBridgeSpecCmd(def, &cfg),
		newEnsureSolverNetSpecCmd(def, &cfg),
		newEnsureGasAppSpecCmd(def, &cfg),
		newUpgradePortalCmd(def, &cfg),
		newUpgradeFeeOracleV1Cmd(def, &cfg),
		newUpgradeGasStationCmd(def, &cfg),
		newUpgradeGasPumpCmd(def, &cfg),
		newUpgradeStakingCmd(def, &cfg),
		newUpgradeSlashingCmd(def, &cfg),
		newUpgradeDistributionCmd(def, &cfg),
		newUpgradeRedenomCmd(def, &cfg),
		newUpgradeBridgeNativeCmd(def, &cfg),
		newUpgradeBridgeL1(def, &cfg),
		newUpgradePortalRegistryCmd(def, &cfg),
		newUpgradeSolverNetInboxCmd(def, &cfg),
		newUpgradeSolverNetOutboxCmd(def, &cfg),
		newUpgradeSolverNetExecutorCmd(def, &cfg),
		newUpgradeSolverNetAllCmd(def, &cfg),
		newSetPortalFeeOracleV2Cmd(def, &cfg),
		newAllowValidatorsCmd(def, &cfg),
		newPlanUpgradeCmd(def, &cfg),
		newCancelUpgradeCmd(def, &cfg),
		newAdminTestCmd(def),
	)

	return cmd
}

func newAllowValidatorsCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allow-operators",
		Short: "Ensure all operators are allowed as validators",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.AllowOperators(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newPlanUpgradeCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan-upgrade",
		Short: "Plan a network upgrade",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.PlanUpgrade(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newCancelUpgradeCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-upgrade",
		Short: "Cancel the current network upgrade",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.CancelPlannedUpgrade(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newEnsurePortalSpecCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ensure-portal-spec",
		Short: "Ensure live portals match local spec, defined in e2e/app/admin/portalspec.go",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.EnsurePortalSpec(cmd.Context(), *def, *cfg, nil)
		},
	}

	return cmd
}

func newEnsureBridgeSpecCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ensure-bridge-spec",
		Short: "Ensure live bridge contracts (l1 and native) match local spec, defined in e2e/app/admin/bridgespec.go",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.EnsureBridgeSpec(cmd.Context(), *def, *cfg, nil)
		},
	}

	return cmd
}

func newEnsureSolverNetSpecCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ensure-solvernet-spec",
		Short: "Ensure live SolverNetInbox contracts match local spec, defined in e2e/app/admin/solvernetspec.go",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.EnsureSolverNetSpec(cmd.Context(), *def, *cfg, nil)
		},
	}

	return cmd
}

func newEnsureGasAppSpecCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ensure-gas-app-spec",
		Short: "Ensure live gas app contracts match local spec, defined in e2e/app/admin/gasappspec.go",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.EnsureGasAppSpec(cmd.Context(), *def, *cfg, nil)
		},
	}

	return cmd
}

func newUpgradePortalCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-portal",
		Short: "Upgrade a portal contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradePortal(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeFeeOracleV1Cmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-fee-oracle-v1",
		Short: "Upgrade FeeOracleV1 contracts.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeFeeOracleV1(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeGasStationCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-gas-station",
		Short: "Upgrade the OmniGasStation contract.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeGasStation(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeGasPumpCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-gas-pump",
		Short: "Upgrade OmniGasPump contracts.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeGasPump(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeStakingCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-staking",
		Short: "Upgrade the Staking predeploy.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeStaking(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeSlashingCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-slashing",
		Short: "Upgrade the Slashing predeploy.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeSlashing(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeDistributionCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-distribution",
		Short: "Upgrade the Distribution predeploy.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeDistribution(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeRedenomCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-redenom",
		Short: "Upgrade the Redenom predeploy.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeDistribution(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeBridgeNativeCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-bridge-native",
		Short: "Upgrade the OmniBridgeNative predeploy.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeBridgeNative(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeBridgeL1(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-bridge-l1",
		Short: "Upgrade the OmniBridgeL1 contract.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeBridgeL1(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradePortalRegistryCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-portal-registry",
		Short: "Upgrade the PortalRegistry predeploy.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradePortalRegistry(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeSolverNetInboxCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-solvernet-inbox",
		Short: "Upgrade the SolverNetInbox contract.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeSolverNetInbox(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeSolverNetOutboxCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-solvernet-outbox",
		Short: "Upgrade the SolverNetOutbox contract.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeSolverNetOutbox(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeSolverNetExecutorCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-solvernet-executor",
		Short: "Upgrade the SolverNetExecutor contract.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeSolverNetExecutor(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newUpgradeSolverNetAllCmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-solvernet-all",
		Short: "Upgrade all of the SolverNet contracts.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.UpgradeSolverNetAll(cmd.Context(), *def, *cfg)
		},
	}

	return cmd
}

func newSetPortalFeeOracleV2Cmd(def *app.Definition, cfg *admin.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-portal-fee-oracle-v2",
		Short: "Sets OmniPortal's FeeOracle to the FeeOracleV2 contract.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return admin.SetPortalFeeOracleV2(cmd.Context(), *def, *cfg)
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

				if err := bridge.DeployBridge(ctx, def.Testnet, def.Backends()); err != nil {
					return errors.Wrap(err, "deploy bridge")
				}

				if err := gasstation.DeployEphemeralGasApp(ctx, def.Testnet, def.Backends()); err != nil {
					return errors.Wrap(err, "deploy gas station")
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
