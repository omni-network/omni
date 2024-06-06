package keeper

import (
	"context"
	"sort"

	"github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/lib/errors"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	blockTable BlockTable
	msgTable   MsgTable
}

func NewKeeper(storeService store.KVStoreService) (Keeper, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_halo_portal_keeper_portal_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		return Keeper{}, errors.Wrap(err, "create module db")
	}

	portalStore, err := NewPortalStore(modDB)
	if err != nil {
		return Keeper{}, errors.Wrap(err, "create portal store")
	}

	return Keeper{
		blockTable: portalStore.BlockTable(),
		msgTable:   portalStore.MsgTable(),
	}, nil
}

func (k Keeper) CreateMsg(ctx sdk.Context, typ types.MsgType, msgTypeID uint64) error {
	if err := typ.Validate(); err != nil {
		return err
	}

	height := uint64(ctx.BlockHeight())

	// Get or create a block to add the message to
	var blockID uint64
	if block, err := k.blockTable.GetByCreatedHeight(ctx, height); ormerrors.IsNotFound(err) {
		blockID, err = k.blockTable.InsertReturningId(ctx, &Block{CreatedHeight: height})
		if err != nil {
			return errors.Wrap(err, "insert block")
		}
	} else if err != nil {
		return errors.Wrap(err, "get block")
	} else {
		blockID = block.GetId()
	}

	err := k.msgTable.Insert(ctx, &Msg{
		BlockId:   blockID,
		MsgType:   uint32(typ),
		MsgTypeId: msgTypeID,
	})
	if err != nil {
		return errors.Wrap(err, "insert message")
	}

	return nil
}

func (k Keeper) getBlockAndMsgs(ctx context.Context, blockID uint64) (*Block, []*Msg, error) {
	block, err := k.blockTable.Get(ctx, blockID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "get block")
	}

	iter, err := k.msgTable.List(ctx, MsgBlockIdIndexKey{}.WithBlockId(blockID))
	if err != nil {
		return nil, nil, errors.Wrap(err, "list messages")
	}

	var msgs []*Msg
	for iter.Next() {
		msg, err := iter.Value()
		if err != nil {
			return nil, nil, errors.Wrap(err, "get msg value")
		}

		msgs = append(msgs, msg)
	}

	sort.Slice(msgs, func(i, j int) bool {
		return msgs[i].GetId() < msgs[j].GetId()
	})

	return block, msgs, nil
}
