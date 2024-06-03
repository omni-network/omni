// Command anvilproxy is the main entry point for the anvilproxy.
package main

import (
	anvilproxycmd "github.com/omni-network/omni/e2e/anvilproxy/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
)

func main() {
	libcmd.Main(anvilproxycmd.New())
}
