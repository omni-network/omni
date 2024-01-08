// Command halo is the main entry point for the halo consensus client.
package main

import (
	appcmd "github.com/omni-network/omni/halo/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
)

func main() {
	libcmd.Main(appcmd.New())
}
