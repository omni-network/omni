package xbridge

import (
	"github.com/omni-network/omni/lib/ethclient"

	"github.com/ethereum/go-ethereum/common"
)

func isEmpty(addr common.Address) bool {
	return addr == common.Address{}
}

func maybeTxHash(receipt *ethclient.Receipt) string {
	if receipt != nil {
		return receipt.TxHash.Hex()
	}

	return "nil"
}
