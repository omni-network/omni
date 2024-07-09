package types

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"

	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type StakingKeeper interface {
	EndBlocker(ctx context.Context) ([]abci.ValidatorUpdate, error)
	GetLastValidators(ctx context.Context) ([]stypes.Validator, error)
}

type ValSetSubscriber interface {
	UpdateValidatorSet(valset *ValidatorSetResponse) error
}
