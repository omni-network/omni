package types

import (
	"context"

	atypes "github.com/omni-network/omni/halo/attest/types"

	abci "github.com/cometbft/cometbft/abci/types"

	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type AttestKeeper interface {
	ListAttestationsFrom(ctx context.Context, chainID uint64, height uint64, max uint64) ([]*atypes.Attestation, error)
}

type StakingKeeper interface {
	EndBlocker(ctx context.Context) ([]abci.ValidatorUpdate, error)
	GetLastValidators(ctx context.Context) ([]stypes.Validator, error)
}

type ValSetSubscriber interface {
	UpdateValidators(valset []abci.ValidatorUpdate)
}
