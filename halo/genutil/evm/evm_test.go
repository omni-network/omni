package evm_test

import (
	"testing"

	"github.com/omni-network/omni/halo/genutil/evm"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestMakeEVMGenesis(t *testing.T) {
	t.Parallel()

	genesis, err := evm.MakeGenesis(netconf.Staging)
	require.NoError(t, err)
	tutil.RequireGoldenJSON(t, genesis)
}
