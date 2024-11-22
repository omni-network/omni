package evm_test

import (
	"testing"

	"github.com/omni-network/omni/halo/genutil/evm"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"

	_ "github.com/omni-network/omni/halo/sdk" // To init SDK config.
)

//go:generate go test . -golden -clean

func TestMakeEVMGenesis(t *testing.T) {
	t.Parallel()

	genesis, err := evm.MakeGenesis(netconf.Staging)
	require.NoError(t, err)
	tutil.RequireGoldenJSON(t, genesis)

	t.Run("backwards", func(t *testing.T) {
		t.Parallel()
		backwards, err := evm.MarshallBackwardsCompatible(genesis)
		require.NoError(t, err)
		tutil.RequireGoldenBytes(t, backwards)
	})
}
