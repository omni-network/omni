package module

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	attestkeeper "github.com/omni-network/omni/halo/attest/keeper"
	v3 "github.com/omni-network/omni/halo/attest/migrations/v3"
)

type Migrator struct {
	keeper *attestkeeper.Keeper
}

func NewMigrator(keeper *attestkeeper.Keeper) Migrator {
	return Migrator{keeper: keeper}
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.MigrateStore(ctx, m.keeper)
}
