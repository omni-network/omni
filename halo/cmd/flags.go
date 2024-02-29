package cmd

import (
	halocfg "github.com/omni-network/omni/halo/config"
	libcmd "github.com/omni-network/omni/lib/cmd"

	"github.com/spf13/pflag"
)

func bindRunFlags(flags *pflag.FlagSet, cfg *halocfg.Config) {
	libcmd.BindHomeFlag(flags, &cfg.HomeDir)
	flags.StringVar(&cfg.EngineJWTFile, "engine-jwt-file", cfg.EngineJWTFile, "The path to the Engine API JWT file")
	flags.Uint64Var(&cfg.SnapshotInterval, "snapshot-interval", cfg.SnapshotInterval, "The interval (in blocks) at which to create snapshots")
	flags.Uint64Var(&cfg.MinRetainBlocks, "min-retain-blocks", cfg.MinRetainBlocks, "Minimum block height offset during ABCI commit to prune CometBFT blocks")
	flags.StringVar(&cfg.BackendType, "app-db-backend", cfg.BackendType, "The type of database for application and snapshots databases")
	flags.StringVar(&cfg.PruningOption, "pruning", cfg.PruningOption, "Pruning strategy (default|nothing|everything)")
	flags.DurationVar(&cfg.EVMBuildDelay, "evm-build-delay", cfg.EVMBuildDelay, "Minimum delay between triggering and fetching a EVM payload build")
	flags.BoolVar(&cfg.EVMBuildOptimistic, "evm-build-optimistic", cfg.EVMBuildOptimistic, "Enables optimistic building of EVM payloads on previous block finalize")
}

func bindInitFlags(flags *pflag.FlagSet, cfg *InitConfig) {
	libcmd.BindHomeFlag(flags, &cfg.HomeDir)
	flags.StringVar(&cfg.Network, "network", cfg.Network, "The network to initialize")
	flags.BoolVar(&cfg.Force, "force", cfg.Force, "Force initialization (overwrite existing files)")
	flags.BoolVar(&cfg.Clean, "clean", cfg.Clean, "Delete home directory before initialization")
}

func bindOperatorFlags(flags *pflag.FlagSet, cfg *OperatorConfig) {
	libcmd.BindHomeFlag(flags, &cfg.HomeDir)
	flags.StringVar(&cfg.OmniAVSAddr, "omni-avs", cfg.OmniAVSAddr, "The contract address of OmniAVS")
}
