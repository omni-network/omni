// Package evmdistribution monitors the Distribution pre-deploy contract and converts
// the Withdraw log events to cosmosSDK x/distribution logic.
package evmdistribution

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	dkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	dtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

const ModuleName = "evmdistribution"

var _ evmenginetypes.EvmEventProcessor = EventProcessor{}

var (
	distributionABI = mustGetABI(bindings.DistributionMetaData)
	withdrawEvent   = mustGetEvent(distributionABI, "Withdraw")
)

// EventProcessor implements the evmenginetypes.EvmEventProcessor interface.
type EventProcessor struct {
	contract *bindings.Distribution
	ethCl    ethclient.Client
	address  common.Address
	dKeeper  dkeeper.Keeper
}

// New returns a new EventProcessor.
func New(ethCl ethclient.Client, dKeeper dkeeper.Keeper) (EventProcessor, error) {
	address := common.HexToAddress(predeploys.Distribution)
	contract, err := bindings.NewDistribution(address, ethCl)
	if err != nil {
		return EventProcessor{}, errors.Wrap(err, "new upgrade")
	}

	return EventProcessor{
		contract: contract,
		ethCl:    ethCl,
		address:  address,
		dKeeper:  dKeeper,
	}, nil
}

func (EventProcessor) Name() string {
	return ModuleName
}

// FilterParams defines the matching EVM log events, see github.com/ethereum/go-ethereum#FilterQuery.
func (p EventProcessor) FilterParams() ([]common.Address, [][]common.Hash) {
	return []common.Address{p.address}, [][]common.Hash{{withdrawEvent.ID}}
}

// Deliver processes a rewards withdrawal log event.
func (p EventProcessor) Deliver(ctx context.Context, _ common.Hash, elog evmenginetypes.EVMEvent) error {
	ethlog, err := elog.ToEthLog()
	if err != nil {
		return err
	}

	switch ethlog.Topics[0] {
	case withdrawEvent.ID:
		event, err := p.contract.ParseWithdraw(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse withdraw")
		}

		return p.deliverWithdraw(ctx, event)
	default:
		return errors.New("unknown event")
	}
}

// deliverWithdraw processes a Withdraw event.
func (p EventProcessor) deliverWithdraw(ctx context.Context, event *bindings.DistributionWithdraw) error {
	log.Info(ctx, "EVM rewards withdrawal detected", "delegator", event.Delegator, "validator", event.Validator)

	msg := dtypes.MsgWithdrawDelegatorReward{
		DelegatorAddress: event.Delegator.String(),
		ValidatorAddress: event.Validator.String(),
	}

	_, err := dkeeper.NewMsgServerImpl(p.dKeeper).WithdrawDelegatorReward(ctx, &msg)
	if err != nil {
		return errors.Wrap(err, "withdraw rewards")
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
