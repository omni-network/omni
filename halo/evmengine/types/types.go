package types

import "github.com/ethereum/go-ethereum/common"

type AddressProvider interface {
	// LocalAddress returns the local validator's ethereum address.
	LocalAddress() common.Address
}
