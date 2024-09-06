// Package evmstaking monitors the Staking pre-deploy contract and converts
// its log events to cosmosSDK x/staking logic.
package evmstaking

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	akeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const ModuleName = "evmstaking"

var _ evmenginetypes.EvmEventProcessor = EventProcessor{}

var (
	stakingABI           = mustGetABI(bindings.StakingMetaData)
	createValidatorEvent = mustGetEvent(stakingABI, "CreateValidator")
	delegateEvent        = mustGetEvent(stakingABI, "Delegate")
)

// EventProcessor implements the evmenginetypes.EvmEventProcessor interface.
type EventProcessor struct {
	contract *bindings.Staking
	ethCl    ethclient.Client
	address  common.Address
	sKeeper  *skeeper.Keeper
	bKeeper  bkeeper.Keeper
	aKeeper  akeeper.AccountKeeper
}

// New returns a new EventProcessor.
func New(
	ethCl ethclient.Client,
	sKeeper *skeeper.Keeper,
	bKeeper bkeeper.Keeper,
	aKeeper akeeper.AccountKeeper,
) (EventProcessor, error) {
	address := common.HexToAddress(predeploys.Staking)
	contract, err := bindings.NewStaking(address, ethCl)
	if err != nil {
		return EventProcessor{}, errors.Wrap(err, "new staking")
	}

	return EventProcessor{
		contract: contract,
		ethCl:    ethCl,
		address:  address,
		sKeeper:  sKeeper,
		bKeeper:  bKeeper,
		aKeeper:  aKeeper,
	}, nil
}

// Prepare returns all omni stake contract EVM event logs from the provided block hash.
func (p EventProcessor) Prepare(ctx context.Context, blockHash common.Hash) ([]*evmenginetypes.EVMEvent, error) {
	logs, err := p.ethCl.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Addresses: p.Addresses(),
		Topics:    [][]common.Hash{{createValidatorEvent.ID, delegateEvent.ID}},
	})
	if err != nil {
		return nil, errors.Wrap(err, "filter logs")
	}

	resp := make([]*evmenginetypes.EVMEvent, 0, len(logs))
	for _, l := range logs {
		topics := make([][]byte, 0, len(l.Topics))
		for _, t := range l.Topics {
			topics = append(topics, t.Bytes())
		}
		resp = append(resp, &evmenginetypes.EVMEvent{
			Address: l.Address.Bytes(),
			Topics:  topics,
			Data:    l.Data,
		})
	}

	return resp, nil
}

func (EventProcessor) Name() string {
	return ModuleName
}

func (p EventProcessor) Addresses() []common.Address {
	return []common.Address{p.address}
}

// Deliver processes a omni deposit log event, which must be one of:
// - CreateValidator
// - Delegate.
func (p EventProcessor) Deliver(ctx context.Context, _ common.Hash, elog *evmenginetypes.EVMEvent) error {
	ethlog := elog.ToEthLog()

	switch ethlog.Topics[0] {
	case createValidatorEvent.ID:
		ev, err := p.contract.ParseCreateValidator(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse create validator")
		}

		if err := p.deliverCreateValidator(ctx, ev); err != nil {
			return errors.Wrap(err, "create validator")
		}
	case delegateEvent.ID:
		ev, err := p.contract.ParseDelegate(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse delegate")
		}

		if err := p.deliverDelegate(ctx, ev); err != nil {
			return errors.Wrap(err, "delegate")
		}
	default:
		return errors.New("unknown event")
	}

	return nil
}

// deliverCreateValidator processes a CreateValidator event, and creates a new validator.
// - Mint the corresponding amount of $STAKE coins.
// - Send the minted coins to the depositor's account.
// - Create a new validator with the depositor's account.
//
// NOTE: if we error, the deposit is lost (on EVM). consider recovery methods.
func (p EventProcessor) deliverCreateValidator(ctx context.Context, ev *bindings.StakingCreateValidator) error {
	pubkey, err := k1util.PubKeyBytesToCosmos(ev.Pubkey)
	if err != nil {
		return errors.Wrap(err, "pubkey to cosmos")
	}

	accAddr := sdk.AccAddress(ev.Validator.Bytes())
	valAddr := sdk.ValAddress(ev.Validator.Bytes())

	amountCoin, amountCoins := omniToBondCoin(ev.Deposit)

	if _, err := p.sKeeper.GetValidator(ctx, valAddr); err == nil {
		return errors.New("validator already exists")
	}

	p.createAccIfNone(ctx, accAddr)

	if err := p.bKeeper.MintCoins(ctx, ModuleName, amountCoins); err != nil {
		return errors.Wrap(err, "mint coins")
	}

	if err := p.bKeeper.SendCoinsFromModuleToAccount(ctx, ModuleName, accAddr, amountCoins); err != nil {
		return errors.Wrap(err, "send coins")
	}

	log.Info(ctx, "EVM staking deposit detected, adding new validator",
		"depositor", ev.Validator.Hex(),
		"amount", ev.Deposit.String())

	msg, err := stypes.NewMsgCreateValidator(
		valAddr.String(),
		pubkey,
		amountCoin,
		stypes.Description{Moniker: ev.Validator.Hex()},
		stypes.NewCommissionRates(math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec()),
		math.NewInt(1)) // Stub out minimum self delegation for now, just use 1.
	if err != nil {
		return errors.Wrap(err, "create validator message")
	}

	_, err = skeeper.NewMsgServerImpl(p.sKeeper).CreateValidator(ctx, msg)
	if err != nil {
		return errors.Wrap(err, "create validator")
	}

	return nil
}

// deliverDelegate processes a Delegate event, and delegates to an existing validator.
// - Mint the corresponding amount of $STAKE coins.
// - Send the minted coins to the delegator's account.
// - Delegate the minted coins to the validator.
//
// NOTE: if we error, the deposit is lost (on EVM). consider recovery methods.
func (p EventProcessor) deliverDelegate(ctx context.Context, ev *bindings.StakingDelegate) error {
	if ev.Delegator != ev.Validator {
		return errors.New("only self delegation")
	}

	delAddr := sdk.AccAddress(ev.Delegator.Bytes())
	valAddr := sdk.ValAddress(ev.Validator.Bytes())

	if _, err := p.sKeeper.GetValidator(ctx, valAddr); err != nil {
		return errors.New("validator does not exist", "validator", valAddr.String())
	}

	amountCoin, amountCoins := omniToBondCoin(ev.Amount)

	p.createAccIfNone(ctx, delAddr)

	if err := p.bKeeper.MintCoins(ctx, ModuleName, amountCoins); err != nil {
		return errors.Wrap(err, "mint coins")
	}

	if err := p.bKeeper.SendCoinsFromModuleToAccount(ctx, ModuleName, delAddr, amountCoins); err != nil {
		return errors.Wrap(err, "send coins")
	}

	log.Info(ctx, "EVM staking delegation detected, delegating",
		"delegator", ev.Delegator.Hex(),
		"validator", ev.Validator.Hex(),
		"amount", ev.Amount.String())

	// Validator already exists, add deposit to self delegation
	msg := stypes.NewMsgDelegate(delAddr.String(), valAddr.String(), amountCoin)
	_, err := skeeper.NewMsgServerImpl(p.sKeeper).Delegate(ctx, msg)
	if err != nil {
		return errors.Wrap(err, "delegate")
	}

	return nil
}

func (p EventProcessor) createAccIfNone(ctx context.Context, addr sdk.AccAddress) {
	if !p.aKeeper.HasAccount(ctx, addr) {
		acc := p.aKeeper.NewAccountWithAddress(ctx, addr)
		p.aKeeper.SetAccount(ctx, acc)
	}
}

// omniToBondCoin converts the $OMNI amount into a $STAKE coin.
// TODO(corver): At this point, it is 1-to1, but this might change in the future.
func omniToBondCoin(amount *big.Int) (sdk.Coin, sdk.Coins) {
	coin := sdk.NewCoin(sdk.DefaultBondDenom, math.NewIntFromBigInt(amount))
	return coin, sdk.NewCoins(coin)
}

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
