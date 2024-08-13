package evmupgrade

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"cosmossdk.io/depinject"
	updatekeeper "cosmossdk.io/x/upgrade/keeper"
)

type DIInputs struct {
	depinject.In
	EthCl        ethclient.Client
	UpdateKeeper *updatekeeper.Keeper
}

type DIOutputs struct {
	depinject.Out
	EventProc         EventProcessor
	InjectedEventProc evmenginetypes.InjectedEventProc
}

func DIProvide(input DIInputs) (DIOutputs, error) {
	proc, err := New(
		input.EthCl,
		input.UpdateKeeper,
	)
	if err != nil {
		return DIOutputs{}, errors.Wrap(err, "new")
	}

	return DIOutputs{
		EventProc:         proc,
		InjectedEventProc: evmenginetypes.InjectEventProc(proc),
	}, nil
}
