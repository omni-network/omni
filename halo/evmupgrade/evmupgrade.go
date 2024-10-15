// Package evmupgrade monitors the Upgrade pre-deploy contract and converts
// its log events to cosmosSDK x/upgrade logic.
package evmupgrade

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/umath"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	ukeeper "cosmossdk.io/x/upgrade/keeper"
	utypes "cosmossdk.io/x/upgrade/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const ModuleName = "evmupgrade"

var _ evmenginetypes.EvmEventProcessor = EventProcessor{}

var (
	upgradeABI         = mustGetABI(bindings.UpgradeMetaData)
	planUpgradeEvent   = mustGetEvent(upgradeABI, "PlanUpgrade")
	cancelUpgradeEvent = mustGetEvent(upgradeABI, "CancelUpgrade")
)

// EventProcessor implements the evmenginetypes.EvmEventProcessor interface.
type EventProcessor struct {
	contract *bindings.Upgrade
	ethCl    ethclient.Client
	address  common.Address
	uKeeper  *ukeeper.Keeper
}

// New returns a new EventProcessor.
func New(ethCl ethclient.Client, uKeeper *ukeeper.Keeper) (EventProcessor, error) {
	address := common.HexToAddress(predeploys.Upgrade)
	contract, err := bindings.NewUpgrade(address, ethCl)
	if err != nil {
		return EventProcessor{}, errors.Wrap(err, "new staking")
	}

	return EventProcessor{
		contract: contract,
		ethCl:    ethCl,
		address:  address,
		uKeeper:  uKeeper,
	}, nil
}

// Prepare returns all omni stake contract EVM event logs from the provided block hash.
func (p EventProcessor) Prepare(ctx context.Context, blockHash common.Hash) ([]evmenginetypes.EVMEvent, error) {
	logs, err := p.ethCl.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Addresses: p.Addresses(),
		Topics:    [][]common.Hash{{planUpgradeEvent.ID, cancelUpgradeEvent.ID}},
	})
	if err != nil {
		return nil, errors.Wrap(err, "filter logs")
	}

	resp := make([]evmenginetypes.EVMEvent, 0, len(logs))
	for _, l := range logs {
		topics := make([][]byte, 0, len(l.Topics))
		for _, t := range l.Topics {
			topics = append(topics, t.Bytes())
		}
		resp = append(resp, evmenginetypes.EVMEvent{
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

// Deliver processes a upgrade log event, which must be one of:
// - PlanUpgrade.
// - CancelUpgrade.
func (p EventProcessor) Deliver(ctx context.Context, _ common.Hash, elog evmenginetypes.EVMEvent) error {
	ethlog, err := elog.ToEthLog()
	if err != nil {
		return err
	}

	switch ethlog.Topics[0] {
	case planUpgradeEvent.ID:
		plan, err := p.contract.ParsePlanUpgrade(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse plan upgrade")
		}

		return p.deliverPlanUpgrade(ctx, plan)
	case cancelUpgradeEvent.ID:
		cancel, err := p.contract.ParseCancelUpgrade(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse cancel upgrade")
		}

		return p.deliverCancelUpgrade(ctx, cancel)
	default:
		return errors.New("unknown event")
	}
}

// deliverCancelUpgrade processes a CancelUpgrade event.
func (p EventProcessor) deliverCancelUpgrade(ctx context.Context, _ *bindings.UpgradeCancelUpgrade) error {
	log.Info(ctx, "EVM cancel upgrade detected")

	msg := utypes.MsgCancelUpgrade{
		Authority: authtypes.NewModuleAddress(ModuleName).String(),
	}

	_, err := ukeeper.NewMsgServerImpl(p.uKeeper).CancelUpgrade(ctx, &msg)
	if err != nil {
		return errors.Wrap(err, "cancel software upgrade")
	}

	return nil
}

// deliverPlanUpgrade processes a PlanUpgrade event.
func (p EventProcessor) deliverPlanUpgrade(ctx context.Context, plan *bindings.UpgradePlanUpgrade) error {
	log.Info(ctx, "EVM plan upgrade detected", "name", plan.Name, "height", plan.Height)

	heightInt64, err := umath.ToInt64(plan.Height)
	if err != nil {
		return err
	}

	msg := utypes.MsgSoftwareUpgrade{
		Authority: authtypes.NewModuleAddress(ModuleName).String(),
		Plan: utypes.Plan{
			Name:   plan.Name,
			Height: heightInt64,
			Info:   plan.Info,
		},
	}

	_, err = ukeeper.NewMsgServerImpl(p.uKeeper).SoftwareUpgrade(ctx, &msg)
	if err != nil {
		return errors.Wrap(err, "plan software upgrade")
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
