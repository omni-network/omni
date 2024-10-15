// Command omni provides the omni command line interface.
package main

import (
	"context"
	"os"

	clicmd "github.com/omni-network/omni/cli/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	cmd := clicmd.New()
	libcmd.WrapRunE(cmd, func(ctx context.Context, err error) {
		if cliErr := new(clicmd.CliError); errors.As(err, &cliErr) {
			// Easy on the eyes CLI error message with suggestions.
			log.Error(ctx, "❌ "+cliErr.Error(), nil)
		} else {
			log.Error(ctx, "❌ Error", err)
		}
	})

	fig := figure.NewFigure("omni", "", true)
	cmd.SetHelpTemplate(fig.String() + "\n" + cmd.HelpTemplate())

	ctx := log.WithCLILogger(context.Background())

	err := cmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}
