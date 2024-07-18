package main

import (
	admcmd "github.com/omni-network/omni/admin/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
)

func main() {
	libcmd.Main(admcmd.New())
}
