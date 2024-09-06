// Package evmslashing monitors the Slashing pre-deploy contract and converts
// its log events to cosmosSDK x/slashing logic.
package evmslashing

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	skeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

const ModuleName = "evmslashing"

var _ evmenginetypes.EvmEventProcessor = EventProcessor{}

var (
	slashingABI = mustGetABI(bindings.SlashingMetaData)
	unjailEvent = mustGetEvent(slashingABI, "Unjail")
)

// EventProcessor implements the evmenginetypes.EvmEventProcessor interface.
type EventProcessor struct {
	contract *bindings.Slashing
	ethCl    ethclient.Client
	address  common.Address
	sKeeper  skeeper.Keeper
}

// New returns a new EventProcessor.
func New(ethCl ethclient.Client, sKeeper skeeper.Keeper) (EventProcessor, error) {
	address := common.HexToAddress(predeploys.Slashing)
	contract, err := bindings.NewSlashing(address, ethCl)
	if err != nil {
		return EventProcessor{}, errors.Wrap(err, "new staking")
	}

	return EventProcessor{
		contract: contract,
		ethCl:    ethCl,
		address:  address,
		sKeeper:  sKeeper,
	}, nil
}

// Prepare returns all omni stake contract EVM event logs from the provided block hash.
func (p EventProcessor) Prepare(ctx context.Context, blockHash common.Hash) ([]*evmenginetypes.EVMEvent, error) {
	logs, err := p.ethCl.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Addresses: p.Addresses(),
		Topics:    [][]common.Hash{{unjailEvent.ID}},
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
// - Unjail.
func (p EventProcessor) Deliver(ctx context.Context, _ common.Hash, elog *evmenginetypes.EVMEvent) error {
	ethlog := elog.ToEthLog()

	switch ethlog.Topics[0] {
	case unjailEvent.ID:
		ev, err := p.contract.ParseUnjail(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse create validator")
		}

		return p.deliverUnjail(ctx, ev)
	default:
		return errors.New("unknown event")
	}
}

// deliverUnjail processes a Unjail event, and unjails an existing validator.
func (p EventProcessor) deliverUnjail(ctx context.Context, ev *bindings.SlashingUnjail) error {
	valAddr := sdk.ValAddress(ev.Validator.Bytes())

	log.Info(ctx, "EVM unjail detected, unjailing validator", "validator", ev.Validator.Hex())

	msg := stypes.NewMsgUnjail(valAddr.String())

	_, err := skeeper.NewMsgServerImpl(p.sKeeper).Unjail(ctx, msg)
	if err != nil {
		return errors.Wrap(err, "unjail validator")
	}

	return nil
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
