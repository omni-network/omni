package main

import (
	"fmt"
	"os"

	"github.com/omni-network/omni/cmd/omni/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
