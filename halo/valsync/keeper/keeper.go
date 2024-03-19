package keeper

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
	sKeeper      *skeeper.Keeper
	txConfig     client.TxConfig
	valsetTable  ValidatorSetTable
	valTable     ValidatorTable
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	txConfig client.TxConfig,
	sKeeper *skeeper.Keeper,
) (Keeper, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_halo_valsync_keeper_valset_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		return Keeper{}, errors.Wrap(err, "create module db")
	}

	valsetStore, err := NewValsetStore(modDB)
	if err != nil {
		return Keeper{}, errors.Wrap(err, "create valset store")
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		txConfig:     txConfig,
		valsetTable:  valsetStore.ValidatorSetTable(),
		valTable:     valsetStore.ValidatorTable(),
		sKeeper:      sKeeper,
	}, nil
}

// InsertValsetIfUpdated inserts a new valset if the validator set has been updated in the current block.
func (k Keeper) InsertValsetIfUpdated(ctx context.Context) error {
	valUpdates, err := k.sKeeper.GetValidatorUpdates(ctx)
	if err != nil {
		return errors.Wrap(err, "get validator updates")
	} else if len(valUpdates) == 0 {
		// No validator updates, nothing to do
		return nil
	}

	return k.InsertValset(ctx)
}

// InsertValset inserts the current validator set into the database.
func (k Keeper) InsertValset(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	valset, err := k.sKeeper.GetLastValidators(ctx)
	if err != nil {
		return errors.Wrap(err, "get last validators")
	} else if len(valset) == 0 {
		return errors.New("empty validator set")
	}

	// TODO(corver): Ensure we are not inserting the same validator set twice.

	valsetID, err := k.valsetTable.InsertReturningId(ctx, &ValidatorSet{CreatedHeight: uint64(sdkCtx.BlockHeight())})
	if err != nil {
		return errors.Wrap(err, "insert valset")
	}

	var totalPower int64
	for _, val := range valset {
		pubkey, err := val.CmtConsPublicKey()
		if err != nil {
			return errors.Wrap(err, "get consensus public key")
		}

		addr, err := k1util.PubKeyPBToAddress(pubkey)
		if err != nil {
			return errors.Wrap(err, "pubkey to address")
		}

		power := val.ConsensusPower(sdk.DefaultPowerReduction)
		err = k.valTable.Insert(ctx, &Validator{
			ValsetId: valsetID,
			Address:  addr.Bytes(),
			Power:    uint64(power),
		})
		if err != nil {
			return errors.Wrap(err, "insert validator")
		}

		totalPower += power
	}

	log.Info(ctx, "Created new validator set",
		"valset_id", valsetID,
		"len", len(valset),
		"total_power", totalPower,
		"height", sdkCtx.BlockHeight(),
	)

	return nil
}
