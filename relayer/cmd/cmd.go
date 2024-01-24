// Package cmd provides the cli for running the relayer.
package cmd

import (
	libcmd "github.com/omni-network/omni/lib/cmd"
	relayer "github.com/omni-network/omni/relayer/app"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	cmd := libcmd.NewRootCmd(
		"relayer",
		"Relayer is a service that relays txs between the omni network and rollups",
	)

	cfg := relayer.DefaultConfig()
	bindRunFlags(cmd.Flags(), &cfg)

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		if err := libcmd.LogFlags(ctx, cmd.Flags()); err != nil {
			return err
		}

		return relayer.Run(ctx, cfg)
	}

	return cmd
}
