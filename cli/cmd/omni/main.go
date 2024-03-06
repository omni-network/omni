// Command omni provides the omni command line interface.
package main

import (
	"context"

	clicmd "github.com/omni-network/omni/cli/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/log"
)

func main() {
	libcmd.MainWithCtx(log.WithNoopLogger(context.Background()), clicmd.New())
}
