package eoa

import (
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/solver/fundthresh"
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
	// testnetThresholds defines the testnt solvernet thresholds RoleSolver: network -> chain -> token -> threshold.
	testnetThresholds = map[netconf.ID]map[uint64]map[tokens.Asset]SolverNetThreshold{
		// NOTE: mainnet thresholds are defined in solver/fundthresh
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
func GetSolverNetThreshold(role Role, network netconf.ID, chainID uint64, asset tokens.Asset) (SolverNetThreshold, bool) {
	if network == netconf.Mainnet { // If mainnet, used thresholds defined in solver/fundthresh
		if role != RoleSolver { // Only RoleSolver has solvernet thresholds
			return SolverNetThreshold{}, false
		}

		tkn, ok := tokens.ByAsset(chainID, asset)
		if !ok {
			return SolverNetThreshold{}, false
		}

		minB := fundthresh.Get(tkn).Min()

		if tkn.Decimals == 6 {
			return SolverNetThreshold{minDec6: bi.ToF64(minB, 6)}, true
		}

		return SolverNetThreshold{minEther: bi.ToF64(minB, 18)}, true
	}

	m := map[Role]map[netconf.ID]map[uint64]map[tokens.Asset]SolverNetThreshold{
		RoleSolver: testnetThresholds,
	}

	resp, ok := m[role][network][chainID][asset]

	return resp, ok
}
