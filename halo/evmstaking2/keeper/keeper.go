package keeper

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/evmstaking2/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	stakingABI           = mustGetABI(bindings.StakingMetaData)
	createValidatorEvent = mustGetEvent(stakingABI, "CreateValidator")
	delegateEvent        = mustGetEvent(stakingABI, "Delegate")
)

// Keeper also implements the evmenginetypes.EvmEventProcessor interface.
type Keeper struct {
	eventsTable     EVMEventTable
	ethCl           ethclient.Client
	address         common.Address
	submissionDelay int64
}

func NewKeeper(
	storeService store.KVStoreService,
	ethCl ethclient.Client,
	submissionDelay int64,
) (*Keeper, error) {
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

	address := common.HexToAddress(predeploys.Staking)

	return &Keeper{
		eventsTable:     evmstakingStore.EVMEventTable(),
		ethCl:           ethCl,
		address:         address,
		submissionDelay: submissionDelay,
	}, nil
}

// EndBlock delivers all pending EVM events on every `k.submissionDelay`'th block.
func (k *Keeper) EndBlock(ctx context.Context) error {
	blockHeight := sdk.UnwrapSDKContext(ctx).BlockHeight()

	if blockHeight%k.submissionDelay != 0 {
		return nil
	}

	iter, err := k.eventsTable.List(ctx, EVMEventIdIndexKey{})
	if err != nil {
		return errors.Wrap(err, "fetch evm events")
	}
	defer iter.Close()

	for iter.Next() {
		val, err := iter.Value()
		if err != nil {
			return errors.Wrap(err, "get event")
		}
		parseAndDeliver(ctx, val.GetEvent())
		err = k.eventsTable.Delete(ctx, val)
		if err != nil {
			return errors.Wrap(err, "delete evm event")
		}
	}

	return nil
}

// Prepare returns all omni stake contract EVM event logs from the provided block hash.
func (k Keeper) Prepare(ctx context.Context, blockHash common.Hash) ([]evmenginetypes.EVMEvent, error) {
	logs, err := k.ethCl.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Addresses: k.Addresses(),
		Topics:    [][]common.Hash{{createValidatorEvent.ID, delegateEvent.ID}},
	})
	if err != nil {
		return nil, errors.Wrap(err, "filter logs")
	}

	resp := make([]evmenginetypes.EVMEvent, 0, len(logs))
	for _, l := range logs {
		topics := make([][]byte, 0, len(l.Topics))
		for _, t := range l.Topics {
			topics = append(topics, t.Bytes())
		}
		resp = append(resp, evmenginetypes.EVMEvent{
			Address: l.Address.Bytes(),
			Topics:  topics,
			Data:    l.Data,
		})
	}

	return resp, nil
}

func (Keeper) Name() string {
	return types.ModuleName
}

func (k Keeper) Addresses() []common.Address {
	return []common.Address{k.address}
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

// parseAndDeliver parses the provided event and tries to deliver it on a state branch.
// If the delivery fails, the error will be logged and the state branch will be discarded.
func parseAndDeliver(context.Context, *evmenginetypes.EVMEvent) {}

// mustGetABI returns the metadata's ABI as an abi.ABI type.
// It panics on error.
func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

// mustGetEvent returns the event with the given name from the ABI.
// It panics if the event is not found.
func mustGetEvent(abi *abi.ABI, name string) abi.Event {
	event, ok := abi.Events[name]
	if !ok {
		panic("event not found")
	}

	return event
}
