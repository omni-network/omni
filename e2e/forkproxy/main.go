// Command forkproxy is the main entry point for the forkproxy.
package main

import (
	forkproxycmd "github.com/omni-network/omni/e2e/forkproxy/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
)

func main() {
	libcmd.Main(forkproxycmd.New())
}
