// Command omni provides the omni command line interface.
package main

import (
	clicmd "github.com/omni-network/omni/cli/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
)

func main() {
	libcmd.Main(clicmd.New())
}
