package evmstaking

import (
	"github.com/omni-network/omni/halo/mybank"
	"github.com/omni-network/omni/lib/errors"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"cosmossdk.io/depinject"
	accountkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

type DIInputs struct {
	depinject.In
	StakingKeeper *stakingkeeper.Keeper
	BankKeeper    mybank.Keeper
	AccountKeeper accountkeeper.AccountKeeper
}

type DIOutputs struct {
	depinject.Out
	EventProc         EventProcessor
	InjectedEventProc evmenginetypes.InjectedEventProc
}

func DIProvide(input DIInputs) (DIOutputs, error) {
	proc, err := New(
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
