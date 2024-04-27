// Package cmd provides the cli for running the indexer.
package cmd

import (
	"github.com/omni-network/omni/explorer/indexer/app"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	cmd := libcmd.NewRootCmd(
		"indexer",
		"Indexer is a service that will initialize our streams to listen to our portals and index "+
			"data and put it in our Omni Blocks DB",
	)

	cfg := app.DefaultConfig()
	bindIndexerFlags(cmd.Flags(), &cfg)

	logCfg := log.DefaultConfig()
	log.BindFlags(cmd.Flags(), &logCfg)

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		ctx, err := log.Init(cmd.Context(), logCfg)
		if err != nil {
			return err
		}

		if err := libcmd.LogFlags(ctx, cmd.Flags()); err != nil {
			return err
		}

		return app.Run(ctx, cfg)
	}

	return cmd
}

func bindIndexerFlags(flags *pflag.FlagSet, cfg *app.Config) {
	netconf.BindFlag(flags, &cfg.Network)
	xchain.BindFlags(flags, &cfg.RPCEndpoints)
	flags.StringVar(&cfg.ExplorerDBConn, "explorer-db-conn", cfg.ExplorerDBConn, "URL to the database")
	flags.StringVar(&cfg.MonitoringAddr, "monitoring-addr", cfg.MonitoringAddr, "The address to bind the monitoring server")
}
