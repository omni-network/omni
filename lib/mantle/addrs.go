package mantle

import (
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/common"
)

var l1Bridges = map[uint64]common.Address{
	evmchain.IDEthereum: addr("0x95fC37A27a2f68e3A647CDc081F0A89bb47c3012"),
}

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}
