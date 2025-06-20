// Command trade is the main entry point for the trade service.
package main

import (
	libcmd "github.com/omni-network/omni/lib/cmd"
	appcmd "github.com/omni-network/omni/scripts/trade/cmd"
)

func main() {
	libcmd.Main(appcmd.New())
}
