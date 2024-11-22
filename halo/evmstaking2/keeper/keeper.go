package keeper

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	abci "github.com/cometbft/cometbft/abci/types"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
)

type Keeper struct {
	eventsTable EVMEventsTable
}

func NewKeeper(storeService store.KVStoreService) (*Keeper, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_halo_evmstaking2_keeper_evmstaking2_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		return nil, errors.Wrap(err, "create module db")
	}

	evmstakingStore, err := NewEvmstaking2Store(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create valsync store")
	}

	return &Keeper{
		eventsTable: evmstakingStore.EVMEventsTable(),
	}, nil
}

func (*Keeper) EndBlock(context.Context) ([]abci.ValidatorUpdate, error) {
	return nil, nil
}
