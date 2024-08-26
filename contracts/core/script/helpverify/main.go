package main

import (
	"context"
	"os"

	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
)

type Params struct {
	EtherscanAPIKey string
	ArbscanAPIKey   string
}

func main() {
	cmd := newCmd()

	libcmd.SilenceErrUsage(cmd)

	ctx := log.WithCLILogger(context.Background())

	err := cmd.ExecuteContext(ctx)
	if err != nil {
		log.Error(ctx, "❌", err)
		os.Exit(1)
	}
}

func newCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "help-verify",
		Short: "Contract verification helpers",
	}

	cmd.AddCommand(
		newGetCreationTxHashCmd(),
		newParseProxyCreate3TxCmd(),
	)

	return cmd
}
