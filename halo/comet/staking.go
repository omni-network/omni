package comet

import (
	"context"
	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/halo/cosmos"
	"github.com/omni-network/omni/lib/engine"
	"slices"
)

type stakingEvents struct {
	Stakes        []*stypes.MsgCreateValidator
	Unstakes      []*stypes.MsgUndelegate // Validator exits are triggered by undelegations by validators of 100% of self-delegation.
	Delegations   []*stypes.MsgDelegate
	Undelegations []*stypes.MsgUndelegate
}

// queryStakingLogEvents returns the staking log events from the omni evm at the given height.
func queryStakingLogEvents(ctx context.Context, ethClient engine.API, stakingAddr common.Address, height uint64,
) (stakingEvents, error) {
	// TODO: Query EVM and convert logs into cosmos types.
	// This is probably a two step process:
	//  - Query raw event logs
	//  - Convert them to cosmos types (this include business logic such as $ETH and $OMNI to $POWER conversion.

	return stakingEvents{}, nil
}

func (a *App) processStakingEvents(ctx context.Context, events stakingEvents, req *abci.RequestFinalizeBlock) ([]abci.ValidatorUpdate, error) {
	// Populate more header fields if required
	header := cmtproto.Header{
		Height:             req.Height,
		Time:               req.Time,
		LastBlockId:        cmtproto.BlockID{},
		NextValidatorsHash: req.NextValidatorsHash,
		ProposerAddress:    req.ProposerAddress,
	}

	sdkCtx := cosmos.SDKContext(ctx, a.db, header)

	keeper := a.modules.Staking
	server := skeeper.NewMsgServerImpl(keeper)

	if err := keeper.BeginBlocker(sdkCtx); err != nil {
		return nil, err
	}

	for _, create := range events.Stakes {
		_, err := server.CreateValidator(sdkCtx, create)
		if err != nil {
			return nil, err
		}
	}

	for _, undelegate := range slices.Concat(events.Unstakes, events.Undelegations) {
		_, err := server.Undelegate(sdkCtx, undelegate)
		if err != nil {
			return nil, err
		}
	}

	valUpdates, err := keeper.EndBlocker(sdkCtx)
	if err != nil {
		return nil, err
	}

	return valUpdates, nil
}
