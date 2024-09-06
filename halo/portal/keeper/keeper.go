package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	blockTable  BlockTable
	msgTable    MsgTable
	offsetTable OffsetTable
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
		blockTable:  portalStore.BlockTable(),
		msgTable:    portalStore.MsgTable(),
		offsetTable: portalStore.OffsetTable(),
	}, nil
}

func (k Keeper) EmitMsg(ctx sdk.Context, typ types.MsgType, msgTypeID uint64, destChainID uint64, shardID xchain.ShardID) (uint64, error) {
	if (destChainID == xchain.BroadcastChainID) != shardID.Broadcast() {
		return 0, errors.New("dest chain and shard broadcast flag mismatch [BUG]")
	}

	height := uint64(ctx.BlockHeight())

	// Get or create a block to add the message to
	var blockID uint64
	if block, err := k.blockTable.GetByCreatedHeight(ctx, height); ormerrors.IsNotFound(err) {
		blockID, err = k.blockTable.InsertReturningId(ctx, &Block{CreatedHeight: height})
		if err != nil {
			return 0, errors.Wrap(err, "insert block")
		}
	} else if err != nil {
		return 0, errors.Wrap(err, "get block")
	} else {
		blockID = block.GetId()
	}

	offset, err := k.incAndGetOffset(ctx, destChainID, shardID)
	if err != nil {
		return 0, errors.Wrap(err, "increment offset")
	}

	err = k.msgTable.Insert(ctx, &Msg{
		BlockId:      blockID,
		MsgType:      uint32(typ),
		MsgTypeId:    msgTypeID,
		DestChainId:  destChainID,
		ShardId:      uint64(shardID),
		StreamOffset: offset,
	})
	if err != nil {
		return 0, errors.Wrap(err, "insert message")
	}

	return blockID, nil
}

func (k Keeper) incAndGetOffset(ctx context.Context, destChainID uint64, shardID xchain.ShardID) (uint64, error) {
	offset, err := k.offsetTable.GetByDestChainIdShardId(ctx, destChainID, uint64(shardID))
	if ormerrors.IsNotFound(err) {
		offset = &Offset{
			DestChainId: destChainID,
			ShardId:     uint64(shardID),
			Offset:      0,
		}
	} else if err != nil {
		return 0, errors.Wrap(err, "get next offset")
	}

	// Increment the offset
	offset.Offset++

	// Save the new offset
	if err := k.offsetTable.Save(ctx, offset); err != nil {
		return 0, errors.Wrap(err, "save next offset")
	}

	return offset.GetOffset(), nil
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

	return block, msgs, nil
}
