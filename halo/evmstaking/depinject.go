package evmstaking

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"cosmossdk.io/depinject"
	accountkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

type DIInputs struct {
	depinject.In
	EthCl         ethclient.Client
	StakingKeeper *stakingkeeper.Keeper
	BankKeeper    bankkeeper.Keeper
	AccountKeeper accountkeeper.AccountKeeper
}

type DIOutputs struct {
	depinject.Out
	EventProc         EventProcessor
	InjectedEventProc evmenginetypes.InjectedEventProc
}

func DIProvide(input DIInputs) (DIOutputs, error) {
	proc, err := New(
		input.EthCl,
		input.StakingKeeper,
		input.BankKeeper,
		input.AccountKeeper,
	)
	if err != nil {
		return DIOutputs{}, errors.Wrap(err, "new")
	}

	return DIOutputs{
		EventProc:         proc,
		InjectedEventProc: evmenginetypes.InjectEventProc(proc),
	}, nil
}
