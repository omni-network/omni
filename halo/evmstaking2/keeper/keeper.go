package keeper

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/evmstaking2/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/feature"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum/common"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/math"
	"cosmossdk.io/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	stakingABI           = mustGetABI(bindings.StakingMetaData)
	createValidatorEvent = mustGetEvent(stakingABI, "CreateValidator")
	delegateEvent        = mustGetEvent(stakingABI, "Delegate")
)

// Keeper also implements the evmenginetypes.EvmEventProcessor interface.
type Keeper struct {
	eventsTable     EVMEventTable
	address         common.Address
	contract        *bindings.Staking
	aKeeper         types.AuthKeeper
	bKeeper         types.BankKeeper
	sKeeper         types.StakingKeeper
	sServer         types.StakingMsgServer
	deliverInterval int64
}

func NewKeeper(
	storeService store.KVStoreService,
	aKeeper types.AuthKeeper,
	bKeeper types.BankKeeper,
	sKeeper types.StakingKeeper,
	sServer types.StakingMsgServer,
	deliverInterval int64,
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
	contract, err := bindings.NewStaking(address, nil) // Passing nil backend if safe since only Parse functions are used.
	if err != nil {
		return &Keeper{}, errors.Wrap(err, "new staking")
	}

	return &Keeper{
		eventsTable:     evmstakingStore.EVMEventTable(),
		aKeeper:         aKeeper,
		bKeeper:         bKeeper,
		sKeeper:         sKeeper,
		sServer:         sServer,
		address:         address,
		contract:        contract,
		deliverInterval: deliverInterval,
	}, nil
}

// EndBlock delivers all pending EVM events on every `k.deliverInterval`'th block.
func (k Keeper) EndBlock(ctx context.Context) error {
	if !feature.FlagEVMStakingModule.Enabled(ctx) {
		return errors.New("unexpected code path [BUG]")
	}

	blockHeight := sdk.UnwrapSDKContext(ctx).BlockHeight()

	if blockHeight%k.deliverInterval != 0 {
		return nil
	}

	iter, err := k.eventsTable.List(ctx, EVMEventIdIndexKey{})
	if err != nil {
		return errors.Wrap(err, "fetch evm events")
	}
	defer iter.Close()

	delivered := false
	for iter.Next() {
		val, err := iter.Value()
		if err != nil {
			return errors.Wrap(err, "get event")
		}
		k.processBufferedEvent(ctx, val.GetEvent())
		err = k.eventsTable.Delete(ctx, val)
		if err != nil {
			return errors.Wrap(err, "delete evm event")
		}
		delivered = true
	}
	if delivered {
		eventDeliveries.Inc()
	}

	bufferedEvents.Set(0)

	return nil
}

func (Keeper) Name() string {
	return types.ModuleName
}

// FilterParams defines the matching EVM log events, see github.com/ethereum/go-ethereum#FilterQuery.
func (k Keeper) FilterParams() ([]common.Address, [][]common.Hash) {
	return []common.Address{k.address}, [][]common.Hash{{createValidatorEvent.ID, delegateEvent.ID}}
}

// Deliver processes a omni deposit log event, which must be one of:
// - CreateValidator
// - Delegate.
// Note that the event delivery is not immediate. Instead, every event is
// first stored in keeper's state. Then all stored events are periodically delivered
// from `EndBlock` at once.
func (k Keeper) Deliver(ctx context.Context, _ common.Hash, elog evmenginetypes.EVMEvent) error {
	if !feature.FlagEVMStakingModule.Enabled(ctx) {
		return errors.New("unexpected code path [BUG]")
	}
	err := k.eventsTable.Insert(ctx, &EVMEvent{
		Event: &elog,
	})
	if err != nil {
		return errors.Wrap(err, "insert evm event")
	}

	bufferedEvents.Inc()

	return nil
}

// processBufferedEvent branches the multi-store, parses the EVM event and tries to deliver it.
// If the delivery succeeds, the multi store branch is committed; if it fails, the corresponding error is logged.
// Panics are intercepted and logged.
func (k Keeper) processBufferedEvent(ctx context.Context, elog *evmenginetypes.EVMEvent) {
	// Branch the store in case processing fails.
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	branchMS := sdkCtx.MultiStore().CacheMultiStore()
	branchCtx := sdkCtx.WithMultiStore(branchMS)

	if err := catch(func() error { //nolint:contextcheck // False positive wrt ctx
		return k.parseAndDeliver(branchCtx, elog)
	}); err != nil {
		log.InfoErr(ctx, "Delivering EVM log event failed", err,
			"name", k.Name(),
			"height", branchCtx.BlockHeight(),
		)
		failedEvents.Inc()

		return
	}

	branchMS.Write()
}

// parseAndDeliver parses the provided event and tries to deliver it on a state branch.
// If the delivery fails, the error will be logged and the state branch will be discarded.
func (k Keeper) parseAndDeliver(ctx context.Context, elog *evmenginetypes.EVMEvent) error {
	ethlog, err := elog.ToEthLog()
	if err != nil {
		return err
	}

	switch ethlog.Topics[0] {
	case createValidatorEvent.ID:
		delegate, err := k.contract.ParseCreateValidator(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse create validator")
		}

		if err := k.deliverCreateValidator(ctx, delegate); err != nil {
			return errors.Wrap(err, "create validator")
		}
	case delegateEvent.ID:
		delegate, err := k.contract.ParseDelegate(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse delegate")
		}

		if err := k.deliverDelegate(ctx, delegate); err != nil {
			return errors.Wrap(err, "delegate")
		}
	default:
		return errors.New("unknown event")
	}

	return nil
}

// deliverDelegate processes a Delegate event, and delegates to an existing validator.
// - Mint the corresponding amount of $STAKE coins.
// - Send the minted coins to the delegator's account.
// - Delegate the minted coins to the validator.
//
// NOTE: if we error, the deposit is lost (on EVM). consider recovery methods.
func (k Keeper) deliverDelegate(ctx context.Context, ev *bindings.StakingDelegate) error {
	if err := verifyStakingDelegate(ctx, ev); err != nil {
		return err
	}

	delAddr := sdk.AccAddress(ev.Delegator.Bytes())
	valAddr := sdk.ValAddress(ev.Validator.Bytes())

	if _, err := k.sKeeper.GetValidator(ctx, valAddr); err != nil {
		return errors.New("validator does not exist", "validator", valAddr.String())
	}

	amountCoin, amountCoins := omniToBondCoin(ev.Amount)

	k.createAccIfNone(ctx, delAddr)

	if err := k.bKeeper.MintCoins(ctx, k.Name(), amountCoins); err != nil {
		return errors.Wrap(err, "mint coins")
	}

	if err := k.bKeeper.SendCoinsFromModuleToAccount(ctx, k.Name(), delAddr, amountCoins); err != nil {
		return errors.Wrap(err, "send coins")
	}

	log.Info(ctx, "EVM staking delegation detected, delegating",
		"delegator", ev.Delegator.Hex(),
		"validator", ev.Validator.Hex(),
		"amount", ev.Amount.String())

	msg := stypes.NewMsgDelegate(delAddr.String(), valAddr.String(), amountCoin)
	_, err := k.sServer.Delegate(ctx, msg)
	if err != nil {
		return errors.Wrap(err, "delegate")
	}

	return nil
}

func (k Keeper) createAccIfNone(ctx context.Context, addr sdk.AccAddress) {
	if !k.aKeeper.HasAccount(ctx, addr) {
		acc := k.aKeeper.NewAccountWithAddress(ctx, addr)
		k.aKeeper.SetAccount(ctx, acc)
	}
}

// deliverCreateValidator processes a CreateValidator event, and creates a new validator.
// - Mint the corresponding amount of $STAKE coins.
// - Send the minted coins to the depositor's account.
// - Create a new validator with the depositor's account.
//
// NOTE: if we error, the deposit is lost (on EVM). consider recovery methods.
func (k Keeper) deliverCreateValidator(ctx context.Context, createValidator *bindings.StakingCreateValidator) error {
	pubkey, err := k1util.PubKeyBytesToCosmos(createValidator.Pubkey)
	if err != nil {
		return errors.Wrap(err, "pubkey to cosmos")
	}

	accAddr := sdk.AccAddress(createValidator.Validator.Bytes())
	valAddr := sdk.ValAddress(createValidator.Validator.Bytes())

	amountCoin, amountCoins := omniToBondCoin(createValidator.Deposit)

	if _, err := k.sKeeper.GetValidator(ctx, valAddr); err == nil {
		return errors.New("validator already exists")
	}

	k.createAccIfNone(ctx, accAddr)

	if err := k.bKeeper.MintCoins(ctx, k.Name(), amountCoins); err != nil {
		return errors.Wrap(err, "mint coins")
	}

	if err := k.bKeeper.SendCoinsFromModuleToAccount(ctx, k.Name(), accAddr, amountCoins); err != nil {
		return errors.Wrap(err, "send coins")
	}

	log.Info(ctx, "EVM staking deposit detected, adding new validator",
		"depositor", createValidator.Validator.Hex(),
		"amount", createValidator.Deposit.String())

	msg, err := stypes.NewMsgCreateValidator(
		valAddr.String(),
		pubkey,
		amountCoin,
		stypes.Description{Moniker: createValidator.Validator.Hex()},
		stypes.NewCommissionRates(math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec()),
		math.NewInt(1)) // Stub out minimum self delegation for now, just use 1.
	if err != nil {
		return errors.Wrap(err, "create validator message")
	}

	_, err = k.sServer.CreateValidator(ctx, msg)
	if err != nil {
		return errors.Wrap(err, "create validator")
	}

	return nil
}

func verifyStakingDelegate(ctx context.Context, delegate *bindings.StakingDelegate) error {
	if !feature.FlagDelegations.Enabled(ctx) && delegate.Delegator != delegate.Validator {
		return errors.New("only self delegation")
	}

	if delegate.Amount == nil {
		return errors.New("stake amount missing")
	}

	return nil
}
