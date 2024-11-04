// Command solver is the main entry point for the solver service.
package main

import (
	libcmd "github.com/omni-network/omni/lib/cmd"
	solvercmd "github.com/omni-network/omni/solver/cmd"
)

func main() {
	libcmd.Main(solvercmd.New())
}
