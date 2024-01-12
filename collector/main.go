// Command explorer-api is the main entry point for the explorer-api.
package main

import (
	appcmd "github.com/omni-network/omni/collector/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
)

func main() {
	libcmd.Main(appcmd.New())
}
