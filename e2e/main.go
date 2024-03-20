package main

import (
	e2ecmd "github.com/omni-network/omni/e2e/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
)

func main() {
	libcmd.Main(e2ecmd.New())
}
