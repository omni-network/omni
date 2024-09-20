package cmd

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	cmtconfig "github.com/cometbft/cometbft/config"
	cmtjson "github.com/cometbft/cometbft/libs/json"
	cmthttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdkflags "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

type statusConfig struct {
	Node   string
	Output string
}

func defaultStatusConfig() statusConfig {
	return statusConfig{
		Output: sdkflags.OutputFormatJSON,
		Node:   cmtconfig.DefaultRPCConfig().ListenAddress,
	}
}

func newStatusCmd() *cobra.Command {
	cfg := defaultStatusConfig()

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Query remote node for status",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			err := printStatus(cmd.Context(), cfg)
			if err != nil {
				return errors.Wrap(err, "status failed")
			}

			return nil
		},
	}

	bindStatusFlags(cmd, &cfg)

	return cmd
}

func printStatus(ctx context.Context, cfg statusConfig) error {
	rpcCl, err := cmthttp.New(cfg.Node, "/websocket")
	if err != nil {
		return errors.Wrap(err, "create rpc client", "address", cfg.Node)
	}

	status, err := rpcCl.Status(ctx)
	if err != nil {
		return errors.Wrap(err, "query status", "address", cfg.Node)
	}

	output, err := cmtjson.Marshal(status)
	if err != nil {
		return errors.Wrap(err, "marshal status")
	}

	err = new(client.Context).WithOutputFormat(cfg.Output).PrintRaw(output)
	if err != nil {
		return errors.Wrap(err, "print status")
	}

	return nil
}
