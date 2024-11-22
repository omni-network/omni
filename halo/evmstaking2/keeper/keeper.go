package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/evmstaking2/types"
	"github.com/omni-network/omni/lib/errors"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/common"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
)

// Keeper also implements the evmenginetypes.EvmEventProcessor interface.
type Keeper struct {
	eventsTable EVMEventTable
}

func NewKeeper(storeService store.KVStoreService) (*Keeper, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_halo_evmstaking2_keeper_evmstaking_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		return nil, errors.Wrap(err, "create module db")
	}

	evmstakingStore, err := NewEvmstakingStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create valsync store")
	}

	return &Keeper{
		eventsTable: evmstakingStore.EVMEventTable(),
	}, nil
}

func (*Keeper) EndBlock(context.Context) ([]abci.ValidatorUpdate, error) {
	return nil, nil
}

// Prepare returns all omni stake contract EVM event logs from the provided block hash.
func (Keeper) Prepare(context.Context, common.Hash) ([]evmenginetypes.EVMEvent, error) {
	return nil, nil
}

func (Keeper) Name() string {
	return types.ModuleName
}

func (Keeper) Addresses() []common.Address {
	return nil
}

// Deliver processes a omni deposit log event, which must be one of:
// - CreateValidator
// - Delegate.
// Note that the event delivery is not immediate. Instead, every event is
// first stored in keeper's state. Then all stored events are periodically delivered
// from `EndBlock` at once.
func (k Keeper) Deliver(ctx context.Context, _ common.Hash, elog evmenginetypes.EVMEvent) error {
	err := k.eventsTable.Insert(ctx, &EVMEvent{
		Event: &elog,
	})
	if err != nil {
		return errors.Wrap(err, "insert evm event")
	}

	return nil
}
