package keeper

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	ptypes "github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// cometValidatorActiveDelay is the number of blocks after a validator update is provided to
// cometBFT before that set becomes the active set.
// If a validator update is provided to cometBFT in block X, then the new set will become active in block X+2.
const cometValidatorActiveDelay = 2

type Keeper struct {
	sKeeper           types.StakingKeeper
	aKeeper           atypes.AttestKeeper
	valsetTable       ValidatorSetTable
	valTable          ValidatorTable
	subscriber        types.ValSetSubscriber
	emilPortal        ptypes.EmitPortal
	subscriberInitted bool

	ethCl           ethclient.Client
	portalRegAdress common.Address
	portalRegistry  *bindings.PortalRegistryFilterer
}

func NewKeeper(
	storeService store.KVStoreService,
	sKeeper types.StakingKeeper,
	aKeeper atypes.AttestKeeper,
	subscriber types.ValSetSubscriber,
	portal ptypes.EmitPortal,
	ethCl ethclient.Client,
) (*Keeper, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_halo_valsync_keeper_valsync_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		return nil, errors.Wrap(err, "create module db")
	}

	valSyncStore, err := NewValsyncStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create valsync store")
	}

	address := common.HexToAddress(predeploys.PortalRegistry)
	protalReg, err := bindings.NewPortalRegistryFilterer(address, ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "new portal registry")
	}

	return &Keeper{
		valsetTable:     valSyncStore.ValidatorSetTable(),
		valTable:        valSyncStore.ValidatorTable(),
		sKeeper:         sKeeper,
		aKeeper:         aKeeper,
		subscriber:      subscriber,
		emilPortal:      portal,
		ethCl:           ethCl,
		portalRegAdress: address,
		portalRegistry:  protalReg,
	}, nil
}

// EndBlock has two responsibilities:
//
// 1. It wraps the staking module EndBlocker, intercepting the resulting validator updates and storing it as
// the next unattested validator set (to be attested to by current validator set before it can be sent to cometBFT).
//
// 2. It checks if any previously unattested validator set has been attested to, marks it as so, and returns its updates
// to pass along to cometBFT to activate that new set.
func (k *Keeper) EndBlock(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	updates, err := k.sKeeper.EndBlocker(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "staking keeper end block")
	}

	if err := k.maybeStoreValidatorUpdates(ctx, updates); err != nil {
		return nil, err
	}

	// The subscriber is only added after `InitGenesis`, so ensure we notify it of the latest valset.
	if err := k.maybeInitSubscriber(ctx); err != nil {
		return nil, err
	}

	// Check if any unattested set has been attested to (and return its updates).
	return k.processAttested(ctx)
}

// maybeStoreValidatorUpdates stores the provided validator updates as the next unattested validator set if not empty.
func (k *Keeper) maybeStoreValidatorUpdates(ctx context.Context, updates []abci.ValidatorUpdate) error {
	if len(updates) == 0 {
		return nil
	}

	valset, err := k.sKeeper.GetLastValidators(ctx)
	if err != nil {
		return errors.Wrap(err, "get last validators")
	} else if len(valset) == 0 {
		return errors.New("empty validator set")
	}

	merged, err := mergeValidatorSet(valset, updates)
	if err != nil {
		return errors.Wrap(err, "merge validator set")
	}

	valsetID, err := k.insertValidatorSet(ctx, merged, false)
	if err != nil {
		return errors.Wrap(err, "insert updates")
	}

	stats := setStats(merged)
	log.Info(ctx, "💫 Storing new unattested validator set",
		"valset_id", valsetID,
		"len", stats.TotalLen,
		"updated", stats.TotalUpdated,
		"removed", stats.TotalRemoved,
		"total_power", stats.TotalPower,
		"height", sdk.UnwrapSDKContext(ctx).BlockHeight(),
	)

	return nil
}

// InsertGenesisSet inserts the current genesis validator set and empty network as the initial epoch.
// Note: This MUST only be called during InitGenesis AFTER the staking module's InitGenesis.
func (k *Keeper) InsertGenesisSet(ctx context.Context) error {
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

	valsetID, err := k.insertValidatorSet(ctx, vals, true)
	if err != nil {
		return errors.Wrap(err, "insert valset")
	} else if valsetID != 1 {
		return errors.New("genesis valset id not 1 [BUG]")
	}

	stats := setStats(vals)
	log.Info(ctx, "💫 Storing genesis validator set",
		"valset_id", valsetID,
		"len", stats.TotalLen,
		"updated", stats.TotalUpdated,
		"removed", stats.TotalRemoved,
		"total_power", stats.TotalPower,
		"height", sdk.UnwrapSDKContext(ctx).BlockHeight(),
	)

	return nil
}

// insertValidatorSet inserts the current validator set into the database.
func (k *Keeper) insertValidatorSet(ctx context.Context, vals []*Validator, isGenesis bool) (uint64, error) {
	var err error
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if len(vals) == 0 {
		return 0, errors.New("empty validators")
	}

	valset := &ValidatorSet{
		CreatedHeight: uint64(sdkCtx.BlockHeight()),
		Attested:      isGenesis, // Only genesis set is automatically attested.
	}

	valset.Id, err = k.valsetTable.InsertReturningId(ctx, valset)
	if err != nil {
		return 0, errors.Wrap(err, "insert valset")
	}

	// Emit this validator set message to portals, updating the resulting block offset/height/id.
	valset.BlockOffset, err = k.emilPortal.EmitMsg(
		sdkCtx,
		ptypes.MsgTypeValSet,
		valset.GetId(),
		xchain.BroadcastChainID,
		xchain.ShardBroadcast0,
	)
	if err != nil {
		return 0, errors.Wrap(err, "emit message")
	}
	if err := k.valsetTable.Update(ctx, valset); err != nil {
		return 0, errors.Wrap(err, "update valset")
	}

	var totalPower, totalUpdated, totalLen, totalRemoved int64
	powers := make(map[common.Address]int64)
	for _, val := range vals {
		val.ValsetId = valset.GetId()
		err = k.valTable.Insert(ctx, val)
		if err != nil {
			return 0, errors.Wrap(err, "insert validator")
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
			return 0, errors.New("negative power")
		}

		pubkey, err := crypto.DecompressPubkey(val.GetPubKey())
		if err != nil {
			return 0, errors.Wrap(err, "get pubkey")
		}
		powers[crypto.PubkeyToAddress(*pubkey)] = val.GetPower()
	}

	// Log a warn if any validator has 1/3 or more of the total power.
	// This is a potential attack vector, as a single validator could halt the chain.
	for address, power := range powers {
		if power >= totalPower/3 && len(powers) > 1 {
			log.Warn(ctx, "🚨 Validator has 1/3 or more of total power", nil,
				"address", address.Hex(),
				"power", power,
				"total_power", totalPower,
			)
		}
	}

	return valset.GetId(), nil
}

func (k *Keeper) maybeInitSubscriber(ctx context.Context) error {
	if k.subscriberInitted {
		return nil
	}

	set, err := k.ActiveSetByHeight(ctx, uint64(sdk.UnwrapSDKContext(ctx).BlockHeight()))
	if err != nil {
		return err
	}

	if err := k.subscriber.UpdateValidatorSet(set); err != nil {
		return err
	}

	k.subscriberInitted = true

	return nil
}

// processAttested possibly marks the next unattested set as attested by querying approved attestations.
// If found, it returns the validator updates for that set.
//
// Note the order doesn't match that of the staking keeper's original updates.
func (k *Keeper) processAttested(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	valset, ok, err := k.nextUnattestedSet(ctx)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, nil // No unattested set, so no updates.
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	chainID, err := netconf.ConsensusChainIDStr2Uint64(sdkCtx.ChainID())
	if err != nil {
		return nil, errors.Wrap(err, "parse chain id")
	}
	conf := xchain.ConfFinalized // TODO(corver): Move this to static netconf.

	// Check if this unattested set was attested to
	if atts, err := k.aKeeper.ListAttestationsFrom(ctx, chainID, uint32(conf), valset.GetBlockOffset(), 1); err != nil {
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
	var activeVals []*Validator
	for valIter.Next() {
		val, err := valIter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "get validator")
		}

		if val.GetPower() > 0 {
			// Skip zero power validators (removed from previous set).
			activeVals = append(activeVals, val)
		}

		if val.GetUpdated() {
			// Only add updated validators to updates.
			updates = append(updates, val.ValidatorUpdate())
		}
	}

	if k.subscriberInitted {
		if err := k.subscriber.UpdateValidatorSet(valSetResponse(valset, activeVals)); err != nil {
			return nil, err
		}
	}

	log.Info(ctx, "💫 Activating attested validator set",
		"valset_id", valset.GetId(),
		"created_height", valset.GetCreatedHeight(),
		"height", sdkCtx.BlockHeight())

	return updates, nil
}

// nextUnattestedSet returns the next unattested validator set (lowest id), or false if none are found.
func (k *Keeper) nextUnattestedSet(ctx context.Context) (*ValidatorSet, bool, error) {
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

type stats struct {
	TotalPower, TotalUpdated, TotalLen, TotalRemoved int64
}

func setStats(vals []*Validator) stats {
	var resp stats
	for _, val := range vals {
		resp.TotalPower += val.GetPower()
		if val.GetUpdated() {
			resp.TotalUpdated++
		}
		if val.GetPower() > 0 {
			resp.TotalLen++
		} else if val.GetPower() == 0 {
			resp.TotalRemoved++
		}
	}

	return resp
}
