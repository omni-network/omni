package keeper

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/evmredenom"
	"github.com/omni-network/omni/halo/evmstaking/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/umath"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
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
	undelegateEvent      = mustGetEvent(stakingABI, "Undelegate")
	editValidatorEvent   = mustGetEvent(stakingABI, "EditValidator")

	eventsByID = map[common.Hash]abi.Event{
		createValidatorEvent.ID: createValidatorEvent,
		delegateEvent.ID:        delegateEvent,
		undelegateEvent.ID:      undelegateEvent,
		editValidatorEvent.ID:   editValidatorEvent,
	}
)

// Keeper also implements the evmenginetypes.EvmEventProcessor interface.
type Keeper struct {
	eventsTable     EVMEventTable
	address         common.Address
	contract        *bindings.Staking
	aKeeper         types.AuthKeeper
	bKeeper         types.WrappedBankKeeper
	sKeeper         types.StakingKeeper
	sServer         types.StakingMsgServer
	deliverInterval int64
}

func NewKeeper(
	storeService store.KVStoreService,
	aKeeper types.AuthKeeper,
	bKeeper types.WrappedBankKeeper,
	sKeeper types.StakingKeeper,
	sServer types.StakingMsgServer,
	deliverInterval int64,
) (*Keeper, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_halo_evmstaking_keeper_evmstaking_proto.Path()},
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

// nextDeliverHeight returns the next deliver height for the EVM events.
// It returns the current block height if the current block height is divisible by `k.deliverInterval`.
// Else it returns the next block height that is divisible by `k.deliverInterval`.
func (k Keeper) nextDeliverHeight(ctx context.Context) int64 {
	blockHeight := sdk.UnwrapSDKContext(ctx).BlockHeight()
	offset := blockHeight % k.deliverInterval
	if offset == 0 {
		return blockHeight
	}

	// Else return next deliver height.
	return blockHeight - offset + k.deliverInterval
}

// shouldDeliver returns true if the EVM events should be delivered on the current block.
func (k Keeper) shouldDeliver(ctx context.Context) bool {
	return k.nextDeliverHeight(ctx) == sdk.UnwrapSDKContext(ctx).BlockHeight()
}

// EndBlock delivers all pending EVM events on every `k.deliverInterval`'th block.
func (k Keeper) EndBlock(ctx context.Context) error {
	if !k.shouldDeliver(ctx) {
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

	return nil
}

func (Keeper) Name() string {
	return types.ModuleName
}

// FilterParams defines the matching EVM log events, see github.com/ethereum/go-ethereum#FilterQuery.
func (k Keeper) FilterParams() ([]common.Address, [][]common.Hash) {
	return []common.Address{k.address}, [][]common.Hash{{
		createValidatorEvent.ID, delegateEvent.ID, undelegateEvent.ID, editValidatorEvent.ID,
	}}
}

// Deliver processes a omni deposit log event, which must be one of:
// - CreateValidator,
// - EditValidator,
// - Delegate,
// - Undelegate.
// Note that the event delivery is not immediate. Instead, every event is
// first stored in keeper's state. Then all stored events are periodically delivered
// from `EndBlock` at once.
func (k Keeper) Deliver(ctx context.Context, _ common.Hash, elog evmenginetypes.EVMEvent) error {
	log.Debug(ctx, "Buffering EVM staking event",
		"name", eventName(&elog),
		"deliver_height", k.nextDeliverHeight(ctx),
	)

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
		log.InfoErr(ctx, "Delivering EVM staking event failed", err,
			"name", eventName(elog),
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
		createVal, err := k.contract.ParseCreateValidator(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse create validator")
		}

		if err := k.deliverCreateValidator(ctx, createVal); err != nil {
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
	case undelegateEvent.ID:
		undelegate, err := k.contract.ParseUndelegate(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse undelegate")
		}

		if err := k.deliverUndelegate(ctx, undelegate); err != nil {
			return errors.Wrap(err, "undelegate")
		}
	case editValidatorEvent.ID:
		editVal, err := k.contract.ParseEditValidator(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse edit validator")
		}

		if err := k.deliverEditValidator(ctx, editVal); err != nil {
			return errors.Wrap(err, "edit validator")
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
	if ev.Amount == nil {
		return errors.New("stake amount missing")
	}

	delAddr := sdk.AccAddress(ev.Delegator.Bytes())
	valAddr := sdk.ValAddress(ev.Validator.Bytes())

	if _, err := k.sKeeper.GetValidator(ctx, valAddr); err != nil {
		return errors.New("validator does not exist", "validator", valAddr.String())
	}

	stake := evmredenom.ToBondCoin(ev.Amount)

	k.createAccIfNone(ctx, delAddr)

	if err := k.bKeeper.MintCoins(ctx, k.Name(), sdk.NewCoins(stake)); err != nil {
		return errors.Wrap(err, "mint coins")
	}

	if err := k.bKeeper.SendCoinsFromModuleToAccountNoWithdrawal(ctx, k.Name(), delAddr, sdk.NewCoins(stake)); err != nil {
		return errors.Wrap(err, "send coins")
	}

	log.Info(ctx, "EVM staking delegation detected, delegating",
		"delegator", ev.Delegator.Hex(),
		"validator", ev.Validator.Hex(),
		"amount", ev.Amount.String(),
		"stake", stake.String())

	msg := stypes.NewMsgDelegate(delAddr.String(), valAddr.String(), stake)
	_, err := k.sServer.Delegate(ctx, msg)
	if err != nil {
		return errors.Wrap(err, "delegate")
	}

	return nil
}

// deliverUndelegate processes an Unelegate event.
func (k Keeper) deliverUndelegate(ctx context.Context, ev *bindings.StakingUndelegate) error {
	if ev.Amount == nil {
		return errors.New("unstake amount missing")
	}

	delAddr := sdk.AccAddress(ev.Delegator.Bytes())
	valAddr := sdk.ValAddress(ev.Validator.Bytes())

	stake := evmredenom.ToBondCoin(ev.Amount)

	log.Info(ctx, "EVM staking undelegation detected, undelegating",
		"delegator", ev.Delegator.Hex(),
		"validator", ev.Validator.Hex(),
		"amount", ev.Amount.String(),
		"stake", stake.String(),
	)

	msg := stypes.NewMsgUndelegate(delAddr.String(), valAddr.String(), stake)
	_, err := k.sServer.Undelegate(ctx, msg)
	if err != nil {
		return errors.Wrap(err, "undelegate")
	}

	return nil
}

func (k Keeper) deliverEditValidator(ctx context.Context, ev *bindings.StakingEditValidator) error {
	valAddr := sdk.ValAddress(ev.Validator.Bytes())

	p := ev.Params
	description := stypes.Description{
		Moniker:         p.Moniker,
		Identity:        p.Identity,
		Website:         p.Website,
		SecurityContact: p.SecurityContact,
		Details:         p.Details,
	}

	var rateOptional *math.LegacyDec
	if p.CommissionRatePercentage != -1 {
		rateI64, err := umath.ToInt64(p.CommissionRatePercentage)
		if err != nil {
			return errors.Wrap(err, "convert commission rate")
		}
		rateDec := math.LegacyNewDec(rateI64)
		rateOptional = &rateDec
	}

	var minSelfOptional *math.Int
	if p.MinSelfDelegation == nil {
		return errors.New("min self delegation missing")
	} else if p.MinSelfDelegation.Int64() != -1 {
		minSelfInt := math.NewIntFromBigInt(p.MinSelfDelegation)
		minSelfOptional = &minSelfInt
	}

	log.Info(ctx, "EVM staking editing validator",
		"validator", ev.Validator.Hex(),
		"moniker", description.Moniker,
		"rate", p.CommissionRatePercentage,
		"min_self", p.MinSelfDelegation,
	)

	msg := stypes.NewMsgEditValidator(valAddr.String(), description, rateOptional, minSelfOptional)
	if _, err := k.sServer.EditValidator(ctx, msg); err != nil {
		return errors.Wrap(err, "edit validator")
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

	stake := evmredenom.ToBondCoin(createValidator.Deposit)

	if _, err := k.sKeeper.GetValidator(ctx, valAddr); err == nil {
		return errors.New("validator already exists")
	}

	k.createAccIfNone(ctx, accAddr)

	if err := k.bKeeper.MintCoins(ctx, k.Name(), sdk.NewCoins(stake)); err != nil {
		return errors.Wrap(err, "mint coins")
	}

	if err := k.bKeeper.SendCoinsFromModuleToAccountNoWithdrawal(ctx, k.Name(), accAddr, sdk.NewCoins(stake)); err != nil {
		return errors.Wrap(err, "send coins")
	}

	log.Info(ctx, "EVM staking deposit detected, adding new validator",
		"depositor", createValidator.Validator.Hex(),
		"amount", createValidator.Deposit.String(),
		"stake", stake.String())

	msg, err := stypes.NewMsgCreateValidator(
		valAddr.String(),
		pubkey,
		stake,
		stypes.Description{Moniker: createValidator.Validator.Hex()},
		stypes.NewCommissionRates(math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec()),
		math.NewInt(1)) // Omni has trusted validator set, so use minimum valid minSelfDelegation.
	if err != nil {
		return errors.Wrap(err, "create validator message")
	}

	_, err = k.sServer.CreateValidator(ctx, msg)
	if err != nil {
		return errors.Wrap(err, "create validator")
	}

	return nil
}
