// Package cmd provides the cli for running the relayer.
package cmd

import (
	"context"

	libcmd "github.com/omni-network/omni/lib/cmd"
	relayer "github.com/omni-network/omni/relayer/app"

	"github.com/spf13/cobra"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"relayer",
		"Relayer is a service that relays txs between the omni network and rollups",
		newRunCmd(relayer.Run),
	)
}

// newRunCmd returns a new cobra command that runs the relayer.
func newRunCmd(runFunc func(context.Context, relayer.Config) error) *cobra.Command {
	cfg := relayer.DefaultRelayerConfig()

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Runs the relayer",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			return runFunc(ctx, cfg)
		},
	}

	bindRunFlags(cmd.Flags(), &cfg)

	return cmd
}
