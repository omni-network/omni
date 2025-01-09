// Package evmslashing monitors the Slashing pre-deploy contract and converts
// its log events to cosmosSDK x/slashing logic.
package evmslashing

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

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
	address  common.Address
	sKeeper  skeeper.Keeper
}

// New returns a new EventProcessor.
func New(sKeeper skeeper.Keeper) (EventProcessor, error) {
	address := common.HexToAddress(predeploys.Slashing)
	contract, err := bindings.NewSlashing(address, nil) // Passing nil backend if safe since only Parse functions are used.
	if err != nil {
		return EventProcessor{}, errors.Wrap(err, "new staking")
	}

	return EventProcessor{
		contract: contract,
		address:  address,
		sKeeper:  sKeeper,
	}, nil
}

func (EventProcessor) Name() string {
	return ModuleName
}

// FilterParams defines the matching EVM log events, see github.com/ethereum/go-ethereum#FilterQuery.
func (p EventProcessor) FilterParams() ([]common.Address, [][]common.Hash) {
	return []common.Address{p.address}, [][]common.Hash{{unjailEvent.ID}}
}

// Deliver processes a omni deposit log event, which must be one of:
// - Unjail.
func (p EventProcessor) Deliver(ctx context.Context, _ common.Hash, elog evmenginetypes.EVMEvent) error {
	ethlog, err := elog.ToEthLog()
	if err != nil {
		return err
	}

	switch ethlog.Topics[0] {
	case unjailEvent.ID:
		unjail, err := p.contract.ParseUnjail(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse create validator")
		}

		return p.deliverUnjail(ctx, unjail)
	default:
		return errors.New("unknown event")
	}
}

// deliverUnjail processes a Unjail event, and unjails an existing validator.
func (p EventProcessor) deliverUnjail(ctx context.Context, unjail *bindings.SlashingUnjail) error {
	valAddr := sdk.ValAddress(unjail.Validator.Bytes())

	log.Info(ctx, "EVM unjail detected, unjailing validator", "validator", unjail.Validator.Hex())

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
