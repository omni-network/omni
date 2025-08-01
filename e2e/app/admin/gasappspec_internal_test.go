package admin

import (
	"testing"

	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

func TestNetworkGasAppSpec(t *testing.T) {
	t.Parallel()
	golden := make(map[netconf.ID]NetworkGasAppSpec)

	for _, network := range netconf.All() {
		if network == netconf.Simnet {
			continue
		}

		golden[network] = gasAppSpec[network]
	}

	tutil.RequireGoldenJSON(t, golden)
}

func TestMakeGasAppDirectives(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		local      GasAppSpec
		live       GasAppSpec
		directives GasAppDirectives
	}{
		{
			"empties - no change",
			GasAppSpec{},
			GasAppSpec{},
			GasAppDirectives{},
		},
		{
			"should pause gas pump",
			GasAppSpec{PauseGasPump: true},
			GasAppSpec{PauseGasPump: false},
			GasAppDirectives{PauseGasPump: true},
		},
		{
			"should unpause gas pump",
			GasAppSpec{PauseGasPump: false},
			GasAppSpec{PauseGasPump: true},
			GasAppDirectives{UnpauseGasPump: true},
		},
		{
			"should pause gas station",
			GasAppSpec{PauseGasStation: true},
			GasAppSpec{PauseGasStation: false},
			GasAppDirectives{PauseGasStation: true},
		},
		{
			"should unpause gas station",
			GasAppSpec{PauseGasStation: false},
			GasAppSpec{PauseGasStation: true},
			GasAppDirectives{UnpauseGasStation: true},
		},
		{
			"pause both - different chains",
			GasAppSpec{PauseGasPump: true, PauseGasStation: true},
			GasAppSpec{PauseGasPump: false, PauseGasStation: false},
			GasAppDirectives{PauseGasPump: true, PauseGasStation: true},
		},
		{
			"unpause both - different chains",
			GasAppSpec{PauseGasPump: false, PauseGasStation: false},
			GasAppSpec{PauseGasPump: true, PauseGasStation: true},
			GasAppDirectives{UnpauseGasPump: true, UnpauseGasStation: true},
		},
		{
			"mixed states - pause pump, unpause station",
			GasAppSpec{PauseGasPump: true, PauseGasStation: false},
			GasAppSpec{PauseGasPump: false, PauseGasStation: true},
			GasAppDirectives{PauseGasPump: true, UnpauseGasStation: true},
		},
		{
			"mixed states - unpause pump, pause station",
			GasAppSpec{PauseGasPump: false, PauseGasStation: true},
			GasAppSpec{PauseGasPump: true, PauseGasStation: false},
			GasAppDirectives{UnpauseGasPump: true, PauseGasStation: true},
		},
		{
			"no change - both already paused",
			GasAppSpec{PauseGasPump: true, PauseGasStation: true},
			GasAppSpec{PauseGasPump: true, PauseGasStation: true},
			GasAppDirectives{},
		},
		{
			"no change - both already unpaused",
			GasAppSpec{PauseGasPump: false, PauseGasStation: false},
			GasAppSpec{PauseGasPump: false, PauseGasStation: false},
			GasAppDirectives{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			dir, err := makeGasAppDirectives(test.local, test.live)

			require.NoError(t, err)
			require.Equal(t, test.directives, dir)
		})
	}
}
