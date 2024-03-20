package keeper

import (
	"context"
	"strconv"

	akeeper "github.com/omni-network/omni/halo/attest/keeper"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/proto/tendermint/crypto"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
	sKeeper      *skeeper.Keeper
	aKeeper      *akeeper.Keeper
	txConfig     client.TxConfig
	valsetTable  ValidatorSetTable
	valTable     ValidatorTable
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	txConfig client.TxConfig,
	sKeeper *skeeper.Keeper,
	aKeeper *akeeper.Keeper,
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
		aKeeper:      aKeeper,
	}, nil
}

// EndBlock has two responsibilities:
//
// 1. It wraps the staking module EndBlocker, intercepting the resulting validator updates and storing it as
// the next unattested validator set (to be attested to by current validator set before it can be sent to cometBFT).
//
// 2. It checks if any previously unattested validator set has been attested to, marks it as so, and returns its updates
// to pass along to cometBFT to activate that new set.
func (k Keeper) EndBlock(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	updates, err := k.sKeeper.EndBlocker(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "staking keeper end block")
	}

	if len(updates) > 0 {
		isUpdate := func(pubkey crypto.PublicKey) bool {
			for _, update := range updates {
				if update.PubKey.Equal(pubkey) {
					return true
				}
			}

			return false
		}

		// Insert the new validator set.
		if err := k.insertValidatorSet(ctx, isUpdate); err != nil {
			return nil, errors.Wrap(err, "insert updates")
		}
	}

	// Check if any unattested set has been attested to (and return its updates).
	return k.processAttested(ctx)
}

// InsertGenesisSet inserts the current genesis validator set into the database.
// This should only be called during InitGenesis AFTER the staking module's InitGenesis.
func (k Keeper) InsertGenesisSet(ctx context.Context) error {
	// All validators are "updated" in the genesis set.
	allUpdated := func(crypto.PublicKey) bool { return true }

	return k.insertValidatorSet(ctx, allUpdated)
}

// insertValidatorSet inserts the current validator set into the database.
func (k Keeper) insertValidatorSet(ctx context.Context, isUpdate func(crypto.PublicKey) bool) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	valset, err := k.sKeeper.GetLastValidators(ctx)
	if err != nil {
		return errors.Wrap(err, "get last validators")
	} else if len(valset) == 0 {
		return errors.New("empty validator set")
	}

	// TODO(corver): Ensure we are not inserting the same validator set twice.

	valsetID, err := k.valsetTable.InsertReturningId(ctx, &ValidatorSet{
		CreatedHeight: uint64(sdkCtx.BlockHeight()),
		Attested:      isGenesis(ctx), // Only genesis set is automatically attested.
	})
	if err != nil {
		return errors.Wrap(err, "insert valset")
	}

	var totalPower, totalUpdates int64
	for _, val := range valset {
		pubkey, err := val.CmtConsPublicKey()
		if err != nil {
			return errors.Wrap(err, "get consensus public key")
		}

		power := val.ConsensusPower(sdk.DefaultPowerReduction)
		err = k.valTable.Insert(ctx, &Validator{
			ValsetId: valsetID,
			PubKey:   pubkey.GetSecp256K1(),
			Power:    power,
			Updated:  isUpdate(pubkey),
		})
		if err != nil {
			return errors.Wrap(err, "insert validator")
		}

		totalPower += power
		if isUpdate(pubkey) {
			totalUpdates++
		}
	}

	log.Info(ctx, "💫 Storing new unattested validator set",
		"valset_id", valsetID,
		"len", len(valset),
		"total_updates", totalUpdates,
		"total_power", totalPower,
		"height", sdkCtx.BlockHeight(),
	)

	return nil
}

// processAttested possibly marks the next unattested set as attested by querying approved attestations.
// If found, it returns the validator updates for that set.
//
// Note the order doesn't match that of the staking keeper's original updates.
func (k Keeper) processAttested(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	valset, ok, err := k.nextUnattestedSet(ctx)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, nil // No unattested set, so no updates.
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	chainID, err := strconv.ParseUint(sdkCtx.ChainID(), 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "parse chain id")
	}

	// Check if this unattested set was attested to (valSet.Id == attestation.BlockHeight)
	if atts, err := k.aKeeper.ListAttestationsFrom(ctx, chainID, valset.GetId(), 1); err != nil {
		return nil, errors.Wrap(err, "list attestations")
	} else if len(atts) == 0 {
		return nil, nil // No attested set, so no updates.
	}

	// Mark the valset as attested.
	valset.Attested = true
	if err := k.valsetTable.Update(ctx, valset); err != nil {
		return nil, errors.Wrap(err, "update valset")
	}

	// Get its validator updates.
	valIter, err := k.valTable.List(ctx, ValidatorValsetIdIndexKey{}.WithValsetId(valset.GetId()))
	if err != nil {
		return nil, errors.Wrap(err, "list validators")
	}
	defer valIter.Close()

	var updates []abci.ValidatorUpdate
	for valIter.Next() {
		val, err := valIter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "get validator")
		} else if !val.GetUpdated() {
			continue // Only return updated validators.
		}

		updates = append(updates, val.ValidatorUpdate())
	}

	log.Info(ctx, "💫 Activating attested validator set",
		"valset_id", valset.GetId(),
		"created_height", valset.GetCreatedHeight(),
		"height", sdkCtx.BlockHeight())

	return updates, nil
}

// nextUnattestedSet returns the next unattested validator set (lowest id), or false if none are found.
func (k Keeper) nextUnattestedSet(ctx context.Context) (*ValidatorSet, bool, error) {
	iter, err := k.valsetTable.List(ctx, ValidatorSetAttestedCreatedHeightIndexKey{}.WithAttested(false), ormlist.DefaultLimit(1))
	if err != nil {
		return nil, false, errors.Wrap(err, "list unattested")
	}
	defer iter.Close()

	if !iter.Next() {
		return nil, false, nil
	}

	valset, err := iter.Value()
	if err != nil {
		return nil, false, errors.Wrap(err, "get unattested")
	}

	return valset, true, nil
}

// isGenesis returns true if the current block is the genesis block (0).
func isGenesis(ctx context.Context) bool {
	return sdk.UnwrapSDKContext(ctx).BlockHeight() == 0
}
