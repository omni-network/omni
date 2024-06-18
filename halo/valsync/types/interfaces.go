package types

import (
	"context"

	atypes "github.com/omni-network/omni/halo/attest/types"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type AttestKeeper interface {
	ListAttestationsFrom(ctx context.Context, chainID uint64, confLevel uint32, offset uint64, max uint64) ([]*atypes.Attestation, error)
}

type StakingKeeper interface {
	EndBlocker(ctx context.Context) ([]abci.ValidatorUpdate, error)
	GetLastValidators(ctx context.Context) ([]stypes.Validator, error)
	GetValidatorByConsAddr(ctx context.Context, consAddr sdk.ConsAddress) (validator stypes.Validator, err error)
}

type ValSetSubscriber interface {
	UpdateValidators(valset []abci.ValidatorUpdate)
}
