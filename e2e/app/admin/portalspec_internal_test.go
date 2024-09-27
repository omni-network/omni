package admin

import (
	"testing"

	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

func TestNetworkPortalSpec(t *testing.T) {
	t.Parallel()
	golden := make(map[netconf.ID]NetworkPortalSpec)

	for _, network := range netconf.All() {
		if network == netconf.Simnet {
			continue
		}

		golden[network] = portalSpec[network]
	}

	tutil.RequireGoldenJSON(t, golden)
}

func TestMakePortalDirectives(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		local      PortalSpec
		live       PortalSpec
		directives PortalDirectives
		isErr      bool
	}{
		{
			"empties - no change",
			PortalSpec{},
			PortalSpec{},
			PortalDirectives{},
			false,
		},
		{
			"should pause all",
			PortalSpec{PauseAll: true},
			PortalSpec{PauseAll: false},
			PortalDirectives{PauseAll: true},
			false,
		},
		{
			"should unpause all",
			PortalSpec{PauseAll: false},
			PortalSpec{PauseAll: true},
			PortalDirectives{UnpauseAll: true},
			false,
		},
		{
			"should pause some xcall to",
			PortalSpec{PauseXCallTo: []uint64{1, 2}},
			PortalSpec{PauseXCallTo: []uint64{1}},
			PortalDirectives{PauseXCallTo: []uint64{2}},
			false,
		},
		{
			"should unpause some xcall to",
			PortalSpec{PauseXCallTo: []uint64{1}},
			PortalSpec{PauseXCallTo: []uint64{1, 2}},
			PortalDirectives{UnpauseXCallTo: []uint64{2}},
			false,
		},
		{
			"unpause all, pause specific",
			PortalSpec{PauseAll: false, PauseXCallTo: []uint64{1}},
			PortalSpec{PauseAll: true},
			PortalDirectives{UnpauseAll: true, PauseXCallTo: []uint64{1}},
			false,
		},
		{
			"unpause specific, pause all",
			PortalSpec{PauseAll: true},
			PortalSpec{PauseAll: false, PauseXCallTo: []uint64{1}, PauseXSubmit: true},
			PortalDirectives{PauseAll: true, UnpauseXCallTo: []uint64{1}, UnpauseXSubmit: true},
			false,
		},
		{
			"invalid spec - pause all and pause xcall to",
			PortalSpec{PauseAll: true, PauseXCallTo: []uint64{1}},
			PortalSpec{},
			PortalDirectives{},
			true,
		},
		{
			"invalid spec - pause xsubmit and pause xsubmit from",
			PortalSpec{PauseXSubmit: true, PauseXSubmitFrom: []uint64{1}},
			PortalSpec{},
			PortalDirectives{},
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			dir, err := makePortalDirectives(test.local, test.live)

			if test.isErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, test.directives, dir)
		})
	}
}
