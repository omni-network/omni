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
	minDec6  float64 // Only applicable to USDC (6 decimals)
}

func (t SolverNetThreshold) MinBalance() *big.Int {
	if t.minDec6 > 0 {
		return bi.Dec6(t.minDec6)
	}

	return bi.Ether(t.minEther)
}

var (
	// solverThresholds defines the solvernet  thresholds RoleSolver: network -> chain -> token -> threshold.
	solverThresholds = map[netconf.ID]map[uint64]map[tokens.Asset]SolverNetThreshold{
		netconf.Mainnet: {
			evmchain.IDEthereum: {
				tokens.WSTETH: {minEther: 10},     // 10 wstETH
				tokens.ETH:    {minEther: 70},     // 70 ETH
				tokens.USDC:   {minDec6: 110_000}, // 110k USDC
			},
			evmchain.IDBase: {
				tokens.USDC: {minDec6: 10_000}, // 10k USDC
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
)

// SolverNetRoles returns the roles that have solvernet thresholds.
func SolverNetRoles() []Role {
	return []Role{RoleSolver}
}

// GetSolverNetThreshold returns the solvernet threshold for the given role, network, chain, and token.
func GetSolverNetThreshold(role Role, network netconf.ID, chainID uint64, tkn tokens.Asset) (SolverNetThreshold, bool) {
	m := map[Role]map[netconf.ID]map[uint64]map[tokens.Asset]SolverNetThreshold{
		RoleSolver: solverThresholds,
	}

	resp, ok := m[role][network][chainID][tkn]

	return resp, ok
}
