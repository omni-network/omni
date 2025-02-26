package admin

import (
	"testing"

	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

func TestNetworkSolverNetSpec(t *testing.T) {
	t.Parallel()
	golden := make(map[netconf.ID]NetworkSolverNetSpec)

	for _, network := range netconf.All() {
		if network == netconf.Simnet {
			continue
		}

		golden[network] = solverNetSpec[network]
	}

	tutil.RequireGoldenJSON(t, golden)
}

func TestMakeSolverNetDirectives(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		local      SolverNetSpec
		live       SolverNetSpec
		directives SolverNetDirectives
		isErr      bool
	}{
		{
			"empties - no change",
			SolverNetSpec{},
			SolverNetSpec{},
			SolverNetDirectives{},
			false,
		},
		{
			"should pause all",
			SolverNetSpec{PauseAll: true},
			SolverNetSpec{PauseAll: false},
			SolverNetDirectives{PauseAll: true},
			false,
		},
		{
			"should unpause all",
			SolverNetSpec{PauseAll: false},
			SolverNetSpec{PauseAll: true},
			SolverNetDirectives{UnpauseAll: true},
			false,
		},
		{
			"should pause open",
			SolverNetSpec{PauseOpen: true},
			SolverNetSpec{PauseOpen: false},
			SolverNetDirectives{PauseOpen: true},
			false,
		},
		{
			"should unpause open",
			SolverNetSpec{PauseOpen: false},
			SolverNetSpec{PauseOpen: true},
			SolverNetDirectives{UnpauseOpen: true},
			false,
		},
		{
			"should pause close",
			SolverNetSpec{PauseClose: true},
			SolverNetSpec{PauseClose: false},
			SolverNetDirectives{PauseClose: true},
			false,
		},
		{
			"should unpause close",
			SolverNetSpec{PauseClose: false},
			SolverNetSpec{PauseClose: true},
			SolverNetDirectives{UnpauseClose: true},
			false,
		},
		{
			"unpause all, pause open",
			SolverNetSpec{PauseAll: false, PauseOpen: true},
			SolverNetSpec{PauseAll: true},
			SolverNetDirectives{UnpauseAll: true},
			false,
		},
		{
			"unpause specific, pause all",
			SolverNetSpec{PauseAll: true},
			SolverNetSpec{PauseAll: false, PauseOpen: true, PauseClose: true},
			SolverNetDirectives{PauseAll: true},
			false,
		},
		{
			"multiple specific changes",
			SolverNetSpec{PauseOpen: true, PauseClose: false},
			SolverNetSpec{PauseOpen: false, PauseClose: true},
			SolverNetDirectives{PauseOpen: true, UnpauseClose: true},
			false,
		},
		{
			"invalid spec - pause all and pause open",
			SolverNetSpec{PauseAll: true, PauseOpen: true},
			SolverNetSpec{},
			SolverNetDirectives{},
			true,
		},
		{
			"invalid spec - pause all and pause close",
			SolverNetSpec{PauseAll: true, PauseClose: true},
			SolverNetSpec{},
			SolverNetDirectives{},
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			dir, err := makeSolverNetDirectives(test.local, test.live)

			if test.isErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, test.directives, dir)
		})
	}
}
