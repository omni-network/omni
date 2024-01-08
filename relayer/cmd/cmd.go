// Package cmd provides the cli for running the relayer.
package cmd

import (
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"
	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"relayer",
		"Relayer is a service that relays txs between the omni network and rollups",
		newRunCmd(),
	)
}

// newRunCmd returns a new cobra command that runs the relayer.
// TODO(@lazar955): Implement this.
func newRunCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Runs the relayer",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			log.Info(ctx, "Relayer started")
			log.Info(ctx, "Press Ctrl+C to stop")
			<-ctx.Done()
			log.Info(ctx, "Relayer stopped")

			return nil
		},
	}
}
