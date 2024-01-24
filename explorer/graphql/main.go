// Command api is the main entry point for the api.
package main

import (
	appcmd "github.com/omni-network/omni/explorer/graphql/cmd"
	libcmd "github.com/omni-network/omni/lib/cmd"
)

func main() {
	libcmd.Main(appcmd.New())
}
