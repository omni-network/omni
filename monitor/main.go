// Command monitor is the main entry point for the monitor service.
package main

import (
	libcmd "github.com/omni-network/omni/lib/cmd"
	appcmd "github.com/omni-network/omni/monitor/cmd"
)

func main() {
	libcmd.Main(appcmd.New())
}
