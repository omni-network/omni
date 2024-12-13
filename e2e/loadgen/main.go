package main

import (
	loadgencmd "github.com/omni-network/omni/e2e/loadgen/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
)

func main() {
	libcmd.Main(loadgencmd.New())
}
