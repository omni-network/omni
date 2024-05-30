package cmd

import (
	"github.com/omni-network/omni/halo/app"
	halocfg "github.com/omni-network/omni/halo/config"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func bindRunFlags(cmd *cobra.Command, cfg *halocfg.Config) {
	flags := cmd.Flags()

	libcmd.BindHomeFlag(flags, &cfg.HomeDir)
	tracer.BindFlags(flags, &cfg.Tracer)
	xchain.BindFlags(flags, &cfg.RPCEndpoints)
	netconf.BindFlag(flags, &cfg.Network)
	flags.StringVar(&cfg.EngineEndpoint, "engine-endpoint", cfg.EngineEndpoint, "An EVM execution client Engine API http endpoint")
	flags.StringVar(&cfg.EngineJWTFile, "engine-jwt-file", cfg.EngineJWTFile, "The path to the Engine API JWT file")
	flags.Uint64Var(&cfg.SnapshotInterval, "snapshot-interval", cfg.SnapshotInterval, "State sync snapshot interval")
	flags.Uint64Var(&cfg.SnapshotKeepRecent, "snapshot-keep-recent", cfg.SnapshotKeepRecent, "State sync snapshot to keep")
	flags.Uint64Var(&cfg.MinRetainBlocks, "min-retain-blocks", cfg.MinRetainBlocks, "Minimum block height offset during ABCI commit to prune CometBFT blocks")
	flags.StringVar(&cfg.BackendType, "app-db-backend", cfg.BackendType, "The type of database for application and snapshots databases")
	flags.StringVar(&cfg.PruningOption, "pruning", cfg.PruningOption, "Pruning strategy (default|nothing|everything)")
	flags.DurationVar(&cfg.EVMBuildDelay, "evm-build-delay", cfg.EVMBuildDelay, "Minimum delay between triggering and fetching a EVM payload build")
	flags.BoolVar(&cfg.EVMBuildOptimistic, "evm-build-optimistic", cfg.EVMBuildOptimistic, "Enables optimistic building of EVM payloads on previous block finalize")
}

func bindRollbackFlags(flags *pflag.FlagSet, cfg *app.RollbackConfig) {
	flags.BoolVar(&cfg.RemoveCometBlock, "hard", cfg.RemoveCometBlock, "Remove last block as well as state")
}

func bindInitFlags(flags *pflag.FlagSet, cfg *InitConfig) {
	libcmd.BindHomeFlag(flags, &cfg.HomeDir)
	netconf.BindFlag(flags, &cfg.Network)
	flags.BoolVar(&cfg.TrustedSync, "trusted-sync", cfg.TrustedSync, "Initialize trusted state-sync height and hash by querying the Omni RPC")
	flags.BoolVar(&cfg.Force, "force", cfg.Force, "Force initialization (overwrite existing files)")
	flags.BoolVar(&cfg.Clean, "clean", cfg.Clean, "Delete home directory before initialization")
}
