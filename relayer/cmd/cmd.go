// Package cmd provides the cli for running the relayer.
package cmd

import (
	"github.com/omni-network/omni/lib/buildinfo"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"
	relayer "github.com/omni-network/omni/relayer/app"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	cmd := libcmd.NewRootCmd(
		"relayer",
		"Relayer is a service that relays txs between the omni network and rollups",
		buildinfo.NewVersionCmd(),
	)

	cfg := relayer.DefaultConfig()
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

		return relayer.Run(ctx, cfg)
	}

	return cmd
}
