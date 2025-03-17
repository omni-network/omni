package eoa

import (
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
)

type SolverNetThreshold struct {
	minEther float64
}

func (t SolverNetThreshold) MinBalance() *big.Int {
	return bi.Ether(t.minEther)
}

var (
	// solverThresholds defines the solvernet  thresholds RoleSolver: network -> chain -> token -> threshold.
	solverThresholds = map[netconf.ID]map[uint64]map[tokens.Token]SolverNetThreshold{
		netconf.Mainnet: {
			evmchain.IDEthereum: {
				tokens.WSTETH: {minEther: 10}, // 10 wstETH
			},
		},
		netconf.Omega: {
			evmchain.IDHolesky: {
				tokens.WSTETH: {minEther: 1}, // 1 wstETH
			},
		},
		netconf.Staging: {
			evmchain.IDHolesky: {
				tokens.WSTETH: {minEther: 1}, // 1 wstETH
			},
		},
	}

	// flowgenThresholds defines the solvernet thresholds RoleFlowgen: network -> chain -> token -> threshold.
	flowgenThresholds = map[netconf.ID]map[uint64]map[tokens.Token]SolverNetThreshold{
		netconf.Mainnet: {
			evmchain.IDEthereum: {
				tokens.WSTETH: {minEther: 0.001}, // 0.001 wstETH
			},
		},
		netconf.Omega: {
			evmchain.IDHolesky: {
				tokens.WSTETH: {minEther: 0.001}, // 0.001 wstETH
			},
		},
		netconf.Staging: {
			evmchain.IDHolesky: {
				tokens.WSTETH: {minEther: 0.001}, // 0.001 wstETH
			},
		},
	}
)

// SolverNetTokens returns the tokens that have solvernet thresholds.
func SolverNetTokens() []tokens.Token {
	return []tokens.Token{tokens.WSTETH}
}

// SolverNetRoles returns the roles that have solvernet thresholds.
func SolverNetRoles() []Role {
	return []Role{RoleSolver, RoleFlowgen}
}

// GetSolverNetThreshold returns the solvernet threshold for the given role, network, chain, and token.
func GetSolverNetThreshold(role Role, network netconf.ID, chainID uint64, tkn tokens.Token) (SolverNetThreshold, bool) {
	m := map[Role]map[netconf.ID]map[uint64]map[tokens.Token]SolverNetThreshold{
		RoleSolver:  solverThresholds,
		RoleFlowgen: flowgenThresholds,
	}

	resp, ok := m[role][network][chainID][tkn]

	return resp, ok
}
