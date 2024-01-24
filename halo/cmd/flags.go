//nolint:lll,revive // Flags are long but don't wrap since it makes the code harder to read.
package cmd

import (
	"github.com/omni-network/omni/halo/app"
	libcmd "github.com/omni-network/omni/lib/cmd"

	"github.com/spf13/pflag"
)

func bindRunFlags(flags *pflag.FlagSet, cfg *app.HaloConfig) {
	libcmd.BindHomeFlag(flags, &cfg.HomeDir)
	flags.StringVar(&cfg.EngineJWTFile, "engine-jwt-file", cfg.EngineJWTFile, "The path to the Engine API JWT file")
	flags.Uint64Var(&cfg.AppStatePersistInterval, "state-persist-interval", cfg.AppStatePersistInterval, "The interval (in blocks) at which to persist the app state")
	flags.Uint64Var(&cfg.SnapshotInterval, "snapshot-interval", cfg.SnapshotInterval, "The interval (in blocks) at which to create snapshots")
}

func bindInitFlags(flags *pflag.FlagSet, cfg *InitConfig) {
	libcmd.BindHomeFlag(flags, &cfg.HomeDir)
	flags.StringVar(&cfg.Network, "network", cfg.Network, "The network to initialize")
	flags.BoolVar(&cfg.Force, "force", cfg.Force, "Force initialization (overwrite existing files)")
	flags.BoolVar(&cfg.Clean, "clean", cfg.Clean, "Delete home directory before initialization")
}
