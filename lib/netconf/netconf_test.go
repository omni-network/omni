package netconf_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestSaveLoad(t *testing.T) {
	t.Parallel()

	var net netconf.Network
	fuzz.NewWithSeed(0).NilChance(0).NumElements(1, 5).Fuzz(&net)

	path := filepath.Join(t.TempDir(), "network.json")
	err := netconf.Save(context.Background(), net, path)
	require.NoError(t, err)

	net2, err := netconf.Load(path)
	require.NoError(t, err)

	require.Equal(t, net, net2)
	tutil.RequireGoldenJSON(t, net)
}

// TestStrats ensures the netconf.StratX matches ethclient.HeadX.
// Netconf shouldn't import ethclient, so using this test to keep in-sync.
func TestStrats(t *testing.T) {
	t.Parallel()

	require.EqualValues(t, ethclient.HeadLatest, netconf.StratLatest)
	require.EqualValues(t, ethclient.HeadFinalized, netconf.StratFinalized)
}
