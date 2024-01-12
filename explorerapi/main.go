// Command explorerapi is the main entry point for the explorerapi.
package main

import (
	appcmd "github.com/omni-network/omni/explorerapi/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
)

func main() {
	libcmd.Main(appcmd.New())
}
