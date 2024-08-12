package evmslashing

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"cosmossdk.io/depinject"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
)

type DIInputs struct {
	depinject.In
	EthCl          ethclient.Client
	SlashingKeeper slashingkeeper.Keeper
}

type DIOutputs struct {
	depinject.Out
	EventProc         EventProcessor
	InjectedEventProc evmenginetypes.InjectedEventProc
}

func DIProvide(input DIInputs) (DIOutputs, error) {
	proc, err := New(
		input.EthCl,
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
