package app

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	libtokens "github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/common"
)

var claimants = map[libtokens.Token]map[netconf.ID]common.Address{
	// for wstETH, we claim orders to a separate rebalancing address
	libtokens.WSTETH: {
		netconf.Mainnet: common.HexToAddress("0x79Ef4d1224a055Ad4Ee5e2226d0cb3720d929AE7"),
		netconf.Omega:   common.HexToAddress("0x521786BE8A0f455700c25FB25F94A4B626E460Ec"),
		netconf.Staging: common.HexToAddress("0x521786BE8A0f455700c25FB25F94A4B626E460Ec"), // same as omega
	},
}

func getClaimant(network netconf.ID, order Order) (common.Address, bool, error) {
	minReceived, err := parseMinReceived(order)
	if err != nil {
		return common.Address{}, false, errors.Wrap(err, "parse min received")
	}

	var cs []common.Address
	for _, output := range minReceived {
		cs = append(cs, claimants[output.Token.Token][network])
	}

	if allEmpty(cs) { // all empty -> solver claims
		return common.Address{}, false, nil
	}

	if !allSame(cs) { // not all the same -> default to solver claims
		return common.Address{}, false, nil
	}

	return cs[0], true, nil
}

func allEmpty(cs []common.Address) bool {
	for _, c := range cs {
		if c != (common.Address{}) {
			return false
		}
	}

	return true
}

func allSame(cs []common.Address) bool {
	for i, c := range cs {
		if i == 0 {
			continue
		}

		if c != cs[i-1] {
			return false
		}
	}

	return true
}
