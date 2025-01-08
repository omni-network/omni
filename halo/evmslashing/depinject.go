package evmslashing

import (
	"github.com/omni-network/omni/lib/errors"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"cosmossdk.io/depinject"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
)

type DIInputs struct {
	depinject.In
	SlashingKeeper slashingkeeper.Keeper
}

type DIOutputs struct {
	depinject.Out
	EventProc         EventProcessor
	InjectedEventProc evmenginetypes.InjectedEventProc
}

func DIProvide(input DIInputs) (DIOutputs, error) {
	proc, err := New(
		input.SlashingKeeper,
	)
	if err != nil {
		return DIOutputs{}, errors.Wrap(err, "new")
	}

	return DIOutputs{
		EventProc:         proc,
		InjectedEventProc: evmenginetypes.InjectEventProc(proc),
	}, nil
}
