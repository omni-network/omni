package keeper

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/core/store"
)

type Keeper struct{}

func NewKeeper(store.KVStoreService) (*Keeper, error) {
	return &Keeper{}, nil
}

func (*Keeper) EndBlock(context.Context) ([]abci.ValidatorUpdate, error) {
	return nil, nil
}
