package keeper

import (
	"context"

	"github.com/omni-network/omni/lib/log"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/slashing/types"
)

var _ types.StakingHooks = Hooks{}

// Hooks implements the staking hooks. It just logs at this point.
type Hooks struct{}

// AfterValidatorBonded updates the signing info start height or create a new signing info.
func (Hooks) AfterValidatorBonded(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	log.Debug(ctx, "📚 Validator bonded", "cons_addr", consAddr, "val_addr", valAddr)
	return nil
}

// AfterValidatorRemoved deletes the address-pubkey relation when a validator is removed,.
func (Hooks) AfterValidatorRemoved(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	log.Debug(ctx, "📚 Validator removed", "cons_addr", consAddr, "val_addr", valAddr)
	return nil
}

// AfterValidatorCreated adds the address-pubkey relation when a validator is created.
func (Hooks) AfterValidatorCreated(ctx context.Context, valAddr sdk.ValAddress) error {
	log.Debug(ctx, "📚 Validator created", "val_addr", valAddr)
	return nil
}

func (Hooks) AfterValidatorBeginUnbonding(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	log.Debug(ctx, "📚 Validator begin unbonding", "cons_addr", consAddr, "val_addr", valAddr)
	return nil
}

func (Hooks) BeforeValidatorModified(ctx context.Context, valAddr sdk.ValAddress) error {
	log.Debug(ctx, "📚 Validator modified", "val_addr", valAddr)
	return nil
}

func (Hooks) BeforeDelegationCreated(ctx context.Context, accAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	log.Debug(ctx, "📚 Delegation created", "acc_addr", accAddr, "val_addr", valAddr)
	return nil
}

func (Hooks) BeforeDelegationSharesModified(ctx context.Context, accAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	log.Debug(ctx, "📚 Delegation shares modified", "acc_addr", accAddr, "val_addr", valAddr)
	return nil
}

func (Hooks) BeforeDelegationRemoved(ctx context.Context, accAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	log.Debug(ctx, "📚 Delegation removed", "acc_addr", accAddr, "val_addr", valAddr)
	return nil
}

func (Hooks) AfterDelegationModified(ctx context.Context, accAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	log.Debug(ctx, "📚 Delegation modified", "acc_addr", accAddr, "val_addr", valAddr)
	return nil
}

func (Hooks) BeforeValidatorSlashed(ctx context.Context, valAddr sdk.ValAddress, amount sdkmath.LegacyDec) error {
	log.Debug(ctx, "📚 Validator slashed", "val_addr", valAddr, "amount", amount)
	return nil
}

func (Hooks) AfterUnbondingInitiated(ctx context.Context, id uint64) error {
	log.Debug(ctx, "📚 Unbonding initiated", "id", id)
	return nil
}
