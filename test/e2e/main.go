package main

import (
	libcmd "github.com/omni-network/omni/lib/cmd"
	e2ecmd "github.com/omni-network/omni/test/e2e/cmd"
)

func main() {
	libcmd.Main(e2ecmd.New())
}
