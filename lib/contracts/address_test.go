package contracts_test

import (
	"testing"

	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestContractAddressReference(t *testing.T) {
	t.Parallel()
	golden := make(map[netconf.ID]map[string]common.Address)
	duplicate := make(map[common.Address]bool)
	for _, network := range netconf.All() {
		if network == netconf.Simnet {
			continue // Skip simnet since it doesn't have eoas.
		}

		addrs := map[string]common.Address{
			"create3":  contracts.Create3Factory(network),
			"portal":   contracts.Portal(network),
			"avs":      contracts.AVS(network),
			"l1bridge": contracts.L1Bridge(network),
			"token":    contracts.Token(network),
		}

		for name, addr := range addrs {
			require.NotEmpty(t, addr, "empty address %s in networks: network=%v", name, network)
			if duplicate[addr] {
				t.Errorf("duplicate address %s in networks: network=%v", name, network)
			}
			duplicate[addr] = true
		}

		if !network.IsEphemeral() {
			golden[network] = addrs // Don't add ephemeral networks to golden since it's not deterministic.
		}
	}

	tutil.RequireGoldenJSON(t, golden)
}
