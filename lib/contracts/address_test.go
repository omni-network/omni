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

		if network == netconf.Staging {
			continue // Skip staging because salt version is dynamic.
		}

		addrs, err := contracts.GetAddresses(t.Context(), network)
		require.NoError(t, err)

		addrsJSON := map[string]common.Address{
			"create3":           addrs.Create3Factory,
			"portal":            addrs.Portal,
			"avs":               addrs.AVS,
			"l1bridge":          addrs.L1Bridge,
			"token":             addrs.Token,
			"nom":               addrs.NomToken,
			"gaspump":           addrs.GasPump,
			"gasstation":        addrs.GasStation,
			"solvernetinbox":    addrs.SolverNetInbox,
			"solvernetoutbox":   addrs.SolverNetOutbox,
			"solvernetexecutor": addrs.SolverNetExecutor,
			"feeoraclev2":       addrs.FeeOracleV2,
		}

		for name, addr := range addrsJSON {
			require.NotEmpty(t, addr, "empty address %s in networks: network=%v", name, network)
			if duplicate[addr] {
				t.Errorf("duplicate address %s in networks: network=%v", name, network)
			}
			duplicate[addr] = true
		}

		golden[network] = addrsJSON
	}

	tutil.RequireGoldenJSON(t, golden)
}
