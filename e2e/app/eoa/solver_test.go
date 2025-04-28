package eoa_test

import (
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/manifests"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -run=TestSolverThresholds -golden

func TestSolverThresholds(t *testing.T) {
	t.Parallel()

	solverGolden := make(map[netconf.ID]map[string]map[string]map[string]string)
	for _, network := range []netconf.ID{netconf.Devnet, netconf.Staging, netconf.Omega, netconf.Mainnet} {
		chains, err := manifests.EVMChains(network)
		require.NoError(t, err)

		for _, chain := range chains {
			for _, role := range eoa.SolverNetRoles() {
				for _, token := range tokens.UniqueAssets() {
					thresholds, ok := eoa.GetSolverNetThreshold(role, network, chain.ChainID, token)
					if !ok {
						continue
					}
					mini := thresholds.MinBalance()
					t.Logf("Thresholds: network=%s, role=%s, chain=%s, token=%s, min=%s",
						network, role, chain.Name, token.Name, token.FormatAmt(mini))

					if role == eoa.RoleSolver {
						addSolverKV(solverGolden, network, chain.Name, token.Symbol, token.FormatAmt(mini))
					}
				}
			}
		}
	}

	tutil.RequireGoldenJSON(t, solverGolden, tutil.WithFilename("solver_reference.json"))
}

func addSolverKV(m map[netconf.ID]map[string]map[string]map[string]string, network netconf.ID, chainID string, token string, min string) {
	if _, ok := m[network]; !ok {
		m[network] = make(map[string]map[string]map[string]string)
	}
	if _, ok := m[network][chainID]; !ok {
		m[network][chainID] = make(map[string]map[string]string)
	}
	m[network][chainID][token] = map[string]string{
		"min": min,
	}
}
