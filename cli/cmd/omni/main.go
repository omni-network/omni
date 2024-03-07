// Command omni provides the omni command line interface.
package main

import (
	"context"
	"os"

	clicmd "github.com/omni-network/omni/cli/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"
)

func main() {
	cmd := clicmd.New()
	libcmd.SilenceErrUsage(cmd)

	ctx := log.WithCLILogger(context.Background())

	err := cmd.ExecuteContext(ctx)
	if err != nil {
		log.Error(ctx, "❌ "+err.Error(), nil)
		os.Exit(1)
	}
}
