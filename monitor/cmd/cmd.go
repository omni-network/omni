// Package cmd provides the cli for running the monitor service
package cmd

import (
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"
	monitor "github.com/omni-network/omni/monitor/app"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	cmd := libcmd.NewRootCmd(
		"monitor",
		"Service that monitors Omni network metrics not measured in individual apps (e.g. smart contracts).",
	)

	cfg := monitor.DefaultConfig()
	bindRunFlags(cmd.Flags(), &cfg)
	bindLoadGenFlags(cmd.Flags(), &cfg.LoadGen)

	logCfg := log.DefaultConfig()
	log.BindFlags(cmd.Flags(), &logCfg)

	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		ctx, err := log.Init(cmd.Context(), logCfg)
		if err != nil {
			return err
		}

		if err := libcmd.LogFlags(ctx, cmd.Flags()); err != nil {
			return err
		}

		return monitor.Run(ctx, cfg)
	}

	return cmd
}
