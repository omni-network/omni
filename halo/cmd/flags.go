//nolint:lll,revive // Flags are long but don't wrap since it makes the code harder to read.
package cmd

import (
	"github.com/omni-network/omni/halo/app"
	libcmd "github.com/omni-network/omni/lib/cmd"

	"github.com/spf13/pflag"
)

func bindHaloFlags(flags *pflag.FlagSet, cfg *app.HaloConfig) {
	libcmd.BindHomeFlag(flags, &cfg.HomeDir)
	flags.StringVar(&cfg.EngineJWTFile, "engine-jwt-file", "", "The path to the Engine API JWT file")
	flags.Uint64Var(&cfg.AppStatePersistInterval, "state-persist-interval", 10, "The interval (in blocks) at which to persist the app state")
	flags.Uint64Var(&cfg.SnapshotInterval, "snapshot-interval", 256, "The interval (in blocks) at which to create snapshots")
}
