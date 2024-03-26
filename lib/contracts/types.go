package contracts

import (
	"github.com/ethereum/go-ethereum/common"
)

type Deployment struct {
	// Address is the address the contract was deployed to.
	Address common.Address

	// BlockHeight is the block height at which the contract was deployed.
	BlockHeight uint64
}

func (d Deployment) IsEmpty() bool {
	return d.Address == common.Address{}
}
