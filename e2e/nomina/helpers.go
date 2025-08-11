package nomina

import "github.com/ethereum/go-ethereum/common"

func isEmpty(addr common.Address) bool {
	return addr == common.Address{}
}
