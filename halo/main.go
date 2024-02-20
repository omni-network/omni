// Command halo is the main entry point for the halo consensus client.
package main

import (
	halocmd "github.com/omni-network/omni/halo/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
)

func main() {
	libcmd.Main(halocmd.New())
}
