package evm_test

import (
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"

	_ "github.com/omni-network/omni/halo/app" // To init SDK config.
)

//go:generate go test . -golden -clean

func TestMakeGenesis(t *testing.T) {
	t.Parallel()

	network := netconf.Staging

	admin, err := eoa.Admin(network)
	require.NoError(t, err)

	genesis, err := evm.MakeGenesis(network, admin)
	require.NoError(t, err)
	tutil.RequireGoldenJSON(t, genesis)
}
