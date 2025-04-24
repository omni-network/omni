// This file mirrors e2e/app/eoa/solver.go and extends
// The two should be merged in the future, or reconciled in tests.
//
//nolint:unused // WIP
package rebalance

import (
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
)

const noMax = -1

type FundThreshold struct {
	token  tokens.Token
	min    float64 // alert if below
	target float64 // fund target, below which, consider deficit
	max    float64 // above which, consider surplus (-1 for no max)
}

func (t FundThreshold) MinBalance() *big.Int {
	return t.balance(t.min)
}

func (t FundThreshold) TargetBalance() *big.Int {
	return t.balance(t.target)
}

func (t FundThreshold) MaxBalance() (*big.Int, bool) {
	if t.max == noMax {
		return nil, false
	}

	return t.balance(t.max), true
}

func (t FundThreshold) balance(f float64) *big.Int {
	if t.token.Decimals == 6 {
		return bi.Dec6(f)
	}

	return bi.Ether(f)
}

var (
	// starting with just wstETH on L1 and Base, for simplicity.
	thresholds = map[netconf.ID]map[tokens.Token]FundThreshold{
		netconf.Mainnet: {
			mustToken(evmchain.IDEthereum, tokens.WSTETH): {
				min:    10,
				target: 50,
				max:    noMax,
			},
			mustToken(evmchain.IDBase, tokens.WSTETH): {
				max: 1,
			},
		},
	}
)

func mustToken(chainID uint64, asset tokens.Asset) tokens.Token {
	tkn, ok := tokens.ByAsset(chainID, asset)
	if !ok {
		panic("token not found")
	}

	return tkn
}
