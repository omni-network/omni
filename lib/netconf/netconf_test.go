package netconf_test

import (
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/test/tutil"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestSaveLoad(t *testing.T) {
	t.Parallel()

	var net netconf.Network
	fuzz.NewWithSeed(0).NilChance(0).NumElements(1, 5).Fuzz(&net)

	path := filepath.Join(t.TempDir(), "network.json")
	err := netconf.Save(net, path)
	require.NoError(t, err)

	net2, err := netconf.Load(path)
	require.NoError(t, err)

	require.Equal(t, net, net2)
	tutil.RequireGoldenJSON(t, net)
}
