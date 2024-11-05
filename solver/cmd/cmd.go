// Package cmd provides the cli for running the solver service
package cmd

import (
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"
	solver "github.com/omni-network/omni/solver/app"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	cmd := libcmd.NewRootCmd(
		"solver",
		"Service that solves Omni network intents",
	)

	cfg := solver.DefaultConfig()
	bindRunFlags(cmd.Flags(), &cfg)

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

		return solver.Run(ctx, cfg)
	}

	return cmd
}
