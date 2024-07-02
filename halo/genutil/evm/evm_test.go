package evm_test

import (
	"testing"

	"github.com/omni-network/omni/halo/genutil/evm"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"

	_ "github.com/omni-network/omni/halo/app" // To init SDK config.
)

//go:generate go test . -golden -clean

func TestMakeGenesis(t *testing.T) {
	t.Parallel()

	genesis, err := evm.MakeGenesis(netconf.Staging)
	require.NoError(t, err)
	tutil.RequireGoldenJSON(t, genesis)
}
