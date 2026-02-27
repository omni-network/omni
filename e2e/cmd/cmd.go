package cmd

import (
	"context"
	"log/slog"
	"regexp"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/bridge"
	"github.com/omni-network/omni/e2e/docker"
	"github.com/omni-network/omni/e2e/gasstation"
	"github.com/omni-network/omni/e2e/nomina"
	"github.com/omni-network/omni/e2e/solve"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/e2e/xbridge"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/xchain"

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
	cachedPreRun := cmd.PersistentPreRunE
	cmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		if err := cachedPreRun(cmd, nil); err != nil {
			return err
		}

		ctx := cmd.Context()
		if _, err := log.Init(ctx, logCfg); err != nil {
			return err
		}

		if err := libcmd.LogFlags(ctx, cmd.Flags()); err != nil {
			return err
		}

		// Some commands do not require a full definition.
		if matchAny(cmd.Use, "hyperliquid-use-big-blocks", "drain-relayer-monitor") {
			return nil
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
		newKeyCreateAll(&def),
		newAdminCmd(&def),
		newERC20FaucetCmd(&def),
		newDeployGasAppCmd(&def),
		newDeployBridgeCmd(&def),
		newDeployFeeOracleV2Cmd(&def),
		newDeployXBridgeCmd(&def),
		newDeploySolverNetCmd(&def),
		newDeployNominaCmd(&def),
		newSetSolverNetRoutesCmd(&def),
		newHyperliquidUseBigBlocksCmd(&defCfg),
		fundAccounts(&def),
		newFundOpsFromSolverCmd(&def),
		newConvertOmniCmd(&def),
		newDrainRelayerMonitorCmd(&defCfg),
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
		Use:     "deploy",
		Aliases: []string{"reset"},
		Short:   "Deploys/Resets the e2e network to start from genesis",
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
			err := docker.ExecComposeVerbose(cmd.Context(), def.Testnet.Dir, "logs")
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
	cfg := app.TestConfig{}
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Runs go tests against the a previously preserved network",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.Test(cmd.Context(), *def, cfg)
		},
	}

	bindTestFlags(cmd.Flags(), &cfg)

	return cmd
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
func fundAccounts(def *app.Definition) *cobra.Command {
	var dryRun bool
	var hotOnly bool

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

			return app.FundAccounts(cmd.Context(), *def, hotOnly, dryRun)
		},
	}

	cmd.Flags().BoolVar(&dryRun, "dry-run", dryRun, "Enables dry-run for testing")
	cmd.Flags().BoolVar(&hotOnly, "hot-only", hotOnly, "Only fund the hot wallet (from the cold)")

	return cmd
}

func newERC20FaucetCmd(def *app.Definition) *cobra.Command {
	cfg := app.DefaultRunERC20FaucetConfig()

	cmd := &cobra.Command{
		Use:   "run-erc20-faucet",
		Short: "Runs the ERC20 faucet",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.RunERC20Faucet(cmd.Context(), *def, cfg)
		},
	}

	bindERC20FaucetFlags(cmd.Flags(), &cfg)

	return cmd
}

func newDeployGasAppCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy-gas-app",
		Short: "Deploys gas pump and gas station contracts",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if def.Testnet.Network.IsEphemeral() {
				return errors.New("only permanent networks")
			}

			return gasstation.DeployEphemeralGasApp(cmd.Context(), def.Testnet, def.Backends())
		},
	}

	return cmd
}

func newDeployBridgeCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy-bridge",
		Short: "Deploys l1 bridge, setups native bridge.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return bridge.DeployBridge(cmd.Context(), def.Testnet, def.Backends())
		},
	}

	return cmd
}

func newDeployFeeOracleV2Cmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy-feeoraclev2",
		Short: "Deploys the FeeOracleV2 contract",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.DeployFeeOracleV2(cmd.Context(), *def)
		},
	}

	return cmd
}

func newDeployXBridgeCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy-xbridge",
		Short: "Deploys the XBridge contracts",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			network, err := networkFromDef(ctx, *def)
			if err != nil {
				return errors.Wrap(err, "network")
			}

			return xbridge.Deploy(ctx, network, def.Backends())
		},
	}

	return cmd
}

func newDeploySolverNetCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy-solvernet",
		Short: "Deploys the SolverNet contracts",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			network, err := networkFromDef(ctx, *def)
			if err != nil {
				return errors.Wrap(err, "network from def")
			}

			endpoints := app.ExternalEndpoints(*def)

			network, backends, err := app.AddSolverNetworkAndBackends(ctx, network, endpoints, def.Cfg, cmd.Name())
			if err != nil {
				return errors.Wrap(err, "get solver network and backends")
			}

			err = app.DeployAllCreate3(ctx, network, backends)
			if err != nil {
				return errors.Wrap(err, "deploy create3")
			}

			return solve.Deploy(cmd.Context(), network, backends)
		},
	}

	return cmd
}

func newDeployNominaCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy-nomina",
		Short: "Deploys the Nomina contracts",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			network, err := networkFromDef(ctx, *def)
			if err != nil {
				return errors.Wrap(err, "network from def")
			}

			return nomina.DeployNomina(ctx, network, def.Backends())
		},
	}

	return cmd
}

func newSetSolverNetRoutesCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-solvernet-routes",
		Short: "Set the SolverNet routes for the given chain IDs.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			network, err := networkFromDef(ctx, *def)
			if err != nil {
				return errors.Wrap(err, "network from def")
			}

			endpoints := app.ExternalEndpoints(*def)

			network, backends, err := app.AddSolverNetworkAndBackends(ctx, network, endpoints, def.Cfg, cmd.Name())
			if err != nil {
				return errors.Wrap(err, "get solver network and backends")
			}

			return solve.SetSolverNetRoutes(cmd.Context(), network, backends)
		},
	}

	return cmd
}

func newHyperliquidUseBigBlocksCmd(cfg *app.DefinitionConfig) *cobra.Command {
	var networkID netconf.ID

	cmd := &cobra.Command{
		Use:   "hyperliquid-use-big-blocks",
		Short: "Enables big HyperEVM blocks for configured accounts",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			if err := networkID.Verify(); err != nil {
				return errors.Wrap(err, "invalid network ID")
			}

			fireCl, err := app.NewFireblocksClient(*cfg, networkID, cmd.Name())
			if err != nil {
				return errors.Wrap(err, "new fireblocks client")
			}

			if err := app.HyperliquidUseBigBlocks(ctx, networkID, fireCl); err != nil {
				return errors.Wrap(err, "use big blocks")
			}

			return nil
		},
	}

	cmd.Flags().StringVar((*string)(&networkID), "network", "", "Network ID to use for Hyperliquid big blocks")
	_ = cmd.MarkFlagRequired("network")

	return cmd
}

func newFundOpsFromSolverCmd(def *app.Definition) *cobra.Command {
	var (
		tokenSymbol string
		chainName   string
		amount      float64
	)

	cmd := &cobra.Command{
		Use:   "fund-ops-from-solver",
		Short: "Funds operations wallet from the solver",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			network, err := networkFromDef(ctx, *def)
			if err != nil {
				return errors.Wrap(err, "network from def")
			}

			endpoints := app.ExternalEndpoints(*def)

			endpoints, err = app.AddSolverEndpoints(ctx, network.ID, endpoints, def.Cfg.RPCOverrides)
			if err != nil {
				return errors.Wrap(err, "add solver endpoints")
			}

			network = solvernet.AddNetwork(ctx, network, solvernet.FilterByEndpoints(endpoints))

			chain, ok := network.ChainByName(chainName)
			if !ok {
				return errors.New("unknown chain", "name", chainName)
			}

			token, ok := tokens.BySymbol(chain.ID, tokenSymbol)
			if !ok {
				return errors.New("unknown token", "symbol", tokenSymbol, "chain", chainName)
			}

			return solve.FundOpsFromSolver(ctx, network, endpoints, token, token.F64ToAmt(amount))
		},
	}

	cmd.Flags().StringVar(&tokenSymbol, "token", "", "Token symbol to fund (e.g. USDC, ETH)")
	cmd.Flags().StringVar(&chainName, "chain", "", "Chain name to fund on (e.g. ethereum, optimism)")
	cmd.Flags().Float64Var(&amount, "amount", 0, "Amount to fund in canonical units (e.g 1 USDC, 1 ETH)")

	_ = cmd.MarkFlagRequired("token")
	_ = cmd.MarkFlagRequired("chain")
	_ = cmd.MarkFlagRequired("amount")

	return cmd
}

func newConvertOmniCmd(def *app.Definition) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-omni",
		Short: "Converts OMNI tokens to NOM tokens",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			network, err := networkFromDef(ctx, *def)
			if err != nil {
				return errors.Wrap(err, "network from def")
			}

			backends := def.Backends()

			// solver key is not included in backends from def, so we add it manually
			solverKey, err := eoa.PrivateKey(ctx, network.ID, eoa.RoleSolver)
			if err != nil {
				return errors.Wrap(err, "solver private key")
			}
			_, err = backends.AddAccount(solverKey)
			if err != nil {
				return errors.Wrap(err, "add solver account")
			}

			return nomina.ConvertOmni(ctx, network, backends)
		},
	}

	return cmd
}

func newDrainRelayerMonitorCmd(defCfg *app.DefinitionConfig) *cobra.Command {
	var dryRun bool

	cmd := &cobra.Command{
		Use:   "drain-relayer-monitor",
		Short: "Transfers relayer and monitor ETH balances to ops wallet on all chains",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()

			manifest, err := app.LoadManifest(defCfg.ManifestFile)
			if err != nil {
				return errors.Wrap(err, "load manifest")
			}

			networkID := manifest.Network

			// Build network and endpoints.
			// Partial & inline, because old utils rely on halted infra.
			endpoints := make(xchain.RPCEndpoints)
			var chains []netconf.Chain
			for _, name := range manifest.PublicChains {
				if rpc, ok := defCfg.RPCOverrides[name]; ok {
					endpoints[name] = rpc
				} else {
					endpoints[name] = types.PublicRPCByName(name)
				}

				chain, err := types.PublicChainByName(name)
				if err != nil {
					return errors.Wrap(err, "public chain", "name", name)
				}

				chains = append(chains, netconf.Chain{
					ID:   chain.ChainID,
					Name: chain.Name,
				})
			}

			network := netconf.Network{
				ID:     networkID,
				Chains: chains,
			}

			return app.DrainRelayerMonitor(ctx, network, endpoints, dryRun)
		},
	}

	cmd.Flags().BoolVar(&dryRun, "dry-run", dryRun, "Enables dry-run mode (no transactions sent)")

	return cmd
}
