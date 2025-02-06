package types

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type WithdrawalKeeper interface {
	// InsertWithdrawal creates a new withdrawal request into the local DB.
	InsertWithdrawal(ctx context.Context, withdrawalAddr sdk.AccAddress, amountGwei uint64) error
}

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

type WithdrawalsProvider interface {
	SumPendingWithdrawalsByAddress(ctx context.Context, in *SumPendingWithdrawalsByAddressRequest) (*SumPendingWithdrawalsByAddressResponse, error)
}
