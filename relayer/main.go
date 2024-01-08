// Command relayer is the main entry point for the relayer.
package main

import (
	libcmd "github.com/omni-network/omni/lib/cmd"
	appcmd "github.com/omni-network/omni/relayer/cmd"
)

func main() {
	libcmd.Main(appcmd.New())
}
