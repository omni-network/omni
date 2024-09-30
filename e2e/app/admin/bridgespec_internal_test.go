package admin

import (
	"testing"

	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
)

func TestNetworkBridgeSpec(t *testing.T) {
	t.Parallel()
	golden := make(map[netconf.ID]NetworkBridgeSpec)

	for _, network := range netconf.All() {
		if network == netconf.Simnet {
			continue
		}

		golden[network] = bridgeSpec[network]
	}

	tutil.RequireGoldenJSON(t, golden)
}
