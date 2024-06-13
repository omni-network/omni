package types

import "github.com/ethereum/go-ethereum/common"

type AddressProvider interface {
	// LocalAddress returns the local validator's ethereum address.
	LocalAddress() common.Address
}

type FeeRecipientProvider interface {
	// LocalFeeRecipient returns the local validator's fee recipient address.
	LocalFeeRecipient() common.Address
	// VerifyFeeRecipient returns true if the given address is a valid fee recipient
	VerifyFeeRecipient(proposedFeeRecipient common.Address) error
}
