// Command omni provides the omni command line interface.
package main

import (
	"context"
	"os"

	clicmd "github.com/omni-network/omni/cli/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	cmd := clicmd.New()

	fig := figure.NewFigure("omni", "", true)
	cmd.SetHelpTemplate(fig.String() + "\n" + cmd.HelpTemplate())

	libcmd.SilenceErrUsage(cmd)

	ctx := log.WithCLILogger(context.Background())

	err := cmd.ExecuteContext(ctx)
	if err != nil {
		log.Error(ctx, "❌ "+err.Error(), nil)
		os.Exit(1)
	}
}
