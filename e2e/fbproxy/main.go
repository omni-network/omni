// The main entry point for the fireblocks proxy command.
package main

import (
	proxycmd "github.com/omni-network/omni/e2e/fbproxy/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
)

func main() {
	libcmd.Main(proxycmd.New())
}
