// Package evmstaking monitors the omni stake pre-deploy contract and converts
// its log events to cosmosSDK logic.
package evmstaking

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	evmenginetypes "github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	akeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const AccountName = "evmstaking"

var _ evmenginetypes.EvmEventProcessor = EventProcessor{}

// EventProcessor implements the evmenginetypes.EvmEventProcessor interface.
type EventProcessor struct {
	contract *bindings.OmniStake
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
	address := common.HexToAddress(predeploys.OmniStake)
	contract, err := bindings.NewOmniStake(address, ethCl)
	if err != nil {
		return EventProcessor{}, errors.Wrap(err, "new omni stake")
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
	// TODO(corver): Maybe we should filter by expected topic only.
	logs, err := p.ethCl.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Addresses: p.Addresses(),
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

func (p EventProcessor) Addresses() []common.Address {
	return []common.Address{p.address}
}

// Deliver processes a omni deposit log event:
// - Mint the corresponding amount of $STAKE coins.
// - Send the minted coins to the depositor's account.
// - Create a new validator with the depositor's account.ccbelunhfhcrcegtfvbgfibjjllhjurctgfrfjhhh.
func (p EventProcessor) Deliver(ctx context.Context, _ common.Hash, elog *evmenginetypes.EVMEvent) error {
	deposit, err := p.contract.ParseDeposit(elog.ToEthLog())
	if err != nil {
		return errors.Wrap(err, "parse deposit")
	}

	stdPubkey, err := k1util.PubKeyFromBytes64(deposit.Pubkey)
	if err != nil {
		return errors.Wrap(err, "deposit pubkey")
	}

	pubkey, err := k1util.StdPubKeyToCosmos(stdPubkey)
	if err != nil {
		return errors.Wrap(err, "pubkey to cosmos")
	}

	ethAddr := crypto.PubkeyToAddress(*stdPubkey)
	accAddr := sdk.AccAddress(ethAddr.Bytes())
	valAddr := sdk.ValAddress(ethAddr.Bytes())

	amountCoin, amountCoins := omniToBondCoin(deposit.Amount)

	_ = p.getOrCreateAccount(ctx, accAddr)

	if err := p.bKeeper.MintCoins(ctx, AccountName, amountCoins); err != nil {
		return errors.Wrap(err, "mint coins")
	}

	if err := p.bKeeper.SendCoinsFromModuleToAccount(ctx, AccountName, accAddr, amountCoins); err != nil {
		return errors.Wrap(err, "send coins")
	}

	if _, err := p.sKeeper.GetValidator(ctx, valAddr); err == nil {
		log.Info(ctx, "EVM staking deposit detected, adding self-delegation",
			"depositor", ethAddr.Hex(),
			"amount", deposit.Amount.String())

		// Validator already exists, add deposit to self delegation
		msg := stypes.NewMsgDelegate(accAddr.String(), valAddr.String(), amountCoin)
		_, err = skeeper.NewMsgServerImpl(p.sKeeper).Delegate(ctx, msg)
		if err != nil {
			return errors.Wrap(err, "delegate")
		}

		return nil
	}

	log.Info(ctx, "EVM staking deposit detected, adding new validator",
		"depositor", ethAddr.Hex(),
		"amount", deposit.Amount.String())

	msg, err := stypes.NewMsgCreateValidator(
		valAddr.String(),
		pubkey,
		amountCoin,
		stypes.Description{Moniker: ethAddr.Hex()},
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

func (p EventProcessor) getOrCreateAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI {
	if p.aKeeper.HasAccount(ctx, addr) {
		return p.aKeeper.GetAccount(ctx, addr)
	}

	acc := p.aKeeper.NewAccountWithAddress(ctx, addr)

	p.aKeeper.SetAccount(ctx, acc)

	return acc
}

// omniToBondCoin converts the $OMNI amount into a $STAKE coin.
// TODO(corver): At this point, it is 1-to1, but this might change in the future.
func omniToBondCoin(amount *big.Int) (sdk.Coin, sdk.Coins) {
	coin := sdk.NewCoin(sdk.DefaultBondDenom, math.NewIntFromBigInt(amount))
	return coin, sdk.NewCoins(coin)
}
