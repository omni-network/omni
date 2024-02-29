package cmd

import (
	halocfg "github.com/omni-network/omni/halo/config"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/cometbft/cometbft/config"

	"github.com/spf13/cobra"
)

// OperatorConfig is the config required for operator registration and deregistration.
type OperatorConfig struct {
	HomeDir     string
	OmniAVSAddr string
	HaloConfig  halocfg.Config
	CometConfig config.Config
	LogConfig   log.Config
}

// newOperatorCmd returns a new cobra command that is used to perform eigen layer relates tasks.
func newOperatorCmd() *cobra.Command {
	cfg := OperatorConfig{
		HomeDir:    halocfg.DefaultHomeDir,
		HaloConfig: halocfg.DefaultConfig(),
		LogConfig:  log.DefaultConfig(),
	}

	operatorCmd := &cobra.Command{
		Use:   "operator",
		Short: "operator commands to interact with OmniAVS",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := log.Init(cmd.Context(), cfg.LogConfig)
			if err != nil {
				return err
			}

			return libcmd.LogFlags(ctx, cmd.Flags())
		},
	}

	operatorCmd.AddCommand(
		RegisterOperatorToOmniAVS(&cfg),
		DeRegisterOperatorFromOmniAVS(),
	)

	return operatorCmd
}
