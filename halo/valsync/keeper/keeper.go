package keeper

import (
	"context"
	"strconv"

	akeeper "github.com/omni-network/omni/halo/attest/keeper"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// cometValidatorActiveDelay is the number of blocks after a validator update is provided to
// cometBFT before that set becomes the active set.
// If a validator update is provided to cometBFT in block X, then the new set will become active in block X+2.
const cometValidatorActiveDelay = 2

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

	// Insert the new validator set.
	if len(updates) > 0 {
		valset, err := k.sKeeper.GetLastValidators(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get last validators")
		} else if len(valset) == 0 {
			return nil, errors.New("empty validator set")
		}

		merged, err := mergeValidatorSet(valset, updates)
		if err != nil {
			return nil, errors.Wrap(err, "merge validator set")
		}

		if err := k.insertValidatorSet(ctx, merged, false); err != nil {
			return nil, errors.Wrap(err, "insert updates")
		}
	}

	// Check if any unattested set has been attested to (and return its updates).
	return k.processAttested(ctx)
}

// InsertGenesisSet inserts the current genesis validator set into the database.
// This should only be called during InitGenesis AFTER the staking module's InitGenesis.
func (k Keeper) InsertGenesisSet(ctx context.Context) error {
	valset, err := k.sKeeper.GetLastValidators(ctx)
	if err != nil {
		return errors.Wrap(err, "get genesis validators")
	} else if len(valset) == 0 {
		return errors.New("empty validator set")
	}

	// Convert
	vals := make([]*Validator, 0, len(valset))
	for _, val := range valset {
		pubkey, err := val.ConsPubKey()
		if err != nil {
			return errors.Wrap(err, "get consensus public key")
		}
		vals = append(vals, &Validator{
			PubKey:  pubkey.Bytes(),
			Power:   val.ConsensusPower(sdk.DefaultPowerReduction),
			Updated: true, // All validators are "updated" in the genesis set.
		})
	}

	return k.insertValidatorSet(ctx, vals, true)
}

// insertValidatorSet inserts the current validator set into the database.
func (k Keeper) insertValidatorSet(ctx context.Context, vals []*Validator, isGenesis bool) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if len(vals) == 0 {
		return errors.New("empty validators")
	}

	// TODO(corver): Ensure we are not inserting the same validator set twice.

	valsetID, err := k.valsetTable.InsertReturningId(ctx, &ValidatorSet{
		CreatedHeight: uint64(sdkCtx.BlockHeight()),
		Attested:      isGenesis, // Only genesis set is automatically attested.
	})
	if err != nil {
		return errors.Wrap(err, "insert valset")
	}

	var totalPower, totalUpdated, totalLen, totalRemoved int64
	for _, val := range vals {
		val.ValsetId = valsetID
		err = k.valTable.Insert(ctx, val)
		if err != nil {
			return errors.Wrap(err, "insert validator")
		}

		totalPower += val.GetPower()
		if val.GetUpdated() {
			totalUpdated++
		}
		if val.GetPower() > 0 {
			totalLen++
		} else if val.GetPower() == 0 {
			totalRemoved++
		} else {
			return errors.New("negative power")
		}
	}

	msg := "💫 Storing new unattested validator set"
	if isGenesis {
		msg = "💫 Storing genesis validator set"
	}

	log.Info(ctx, msg,
		"valset_id", valsetID,
		"len", totalLen,
		"updated", totalUpdated,
		"removed", totalRemoved,
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
	valset.ActivatedHeight = uint64(sdkCtx.BlockHeight()) + cometValidatorActiveDelay
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

// mergeValidatorSet returns the validator set with any zero power updates merged in.
// The valsetID is not set.
func mergeValidatorSet(valset []stypes.Validator, updates []abci.ValidatorUpdate) ([]*Validator, error) {
	var resp []*Validator //nolint:prealloc // We don't know the length of the result.

	added := make(map[string]bool)
	for _, update := range updates {
		resp = append(resp, &Validator{
			PubKey:  update.PubKey.GetSecp256K1(),
			Power:   update.Power,
			Updated: true,
		})
		added[update.PubKey.String()] = true
	}

	for _, val := range valset {
		pubkey, err := val.CmtConsPublicKey()
		if err != nil {
			return nil, errors.Wrap(err, "get consensus public key")
		}

		if added[pubkey.String()] {
			continue
		}

		resp = append(resp, &Validator{
			PubKey:  pubkey.GetSecp256K1(),
			Power:   val.ConsensusPower(sdk.DefaultPowerReduction),
			Updated: false,
		})
	}

	return resp, nil
}
