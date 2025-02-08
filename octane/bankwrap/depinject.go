package bankwrap

import (
	"cosmossdk.io/depinject"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type DIInputs struct {
	depinject.In

	BankKeeper bankkeeper.BaseKeeper
}

type DIOutputs struct {
	depinject.Out

	Wrapper *Wrapper
}

// DIInvoke injects the EVMEngineKeeper into the Wrapper.
// This mitigates cyclic dependency: EVMEngineKeeper -> EVMEventProc(EVMStakingKeeper) -> BankWrapperKeeper -> EVMEngineKeeper.
func DIInvoke(wrapper *Wrapper, evmEngKeeper EVMEngineKeeper) {
	if evmEngKeeper == nil {
		panic("nil EVMEngineKeeper")
	}

	wrapper.SetEVMEngineKeeper(evmEngKeeper)
}

func DIProvide(input DIInputs) (DIOutputs, error) {
	return DIOutputs{
		Wrapper: NewWrapper(input.BankKeeper),
	}, nil
}

// sdkOverrides is a list of standard cosmosSDK module x/bank Keeper interfaces that are overridden by the Wrapper.
// Note that all these modules require Burner permissions in order to create withdrawals.
var sdkOverrides = []string{
	"github.com/cosmos/cosmos-sdk/x/distribution/types/types.BankKeeper",
	"github.com/cosmos/cosmos-sdk/x/slashing/types/types.BankKeeper",
	"github.com/cosmos/cosmos-sdk/x/mint/types/types.BankKeeper",
	"github.com/cosmos/cosmos-sdk/x/staking/types/types.BankKeeper",
	"github.com/cosmos/cosmos-sdk/x/auth/types/types.BankKeeper",
	"github.com/cosmos/cosmos-sdk/x/auth/types/types.BankKeeper",
	"github.com/cosmos/cosmos-sdk/x/auth/tx/config/tx.BankKeeper",
	"github.com/cosmos/cosmos-sdk/x/mint/types/types.BankKeeper",
	"github.com/cosmos/cosmos-sdk/x/bank/keeper/keeper.Keeper",
}

// WrapperImpl is the full path to the bankwrap.Wrapper implementation.
const WrapperImpl = "github.com/omni-network/omni/octane/bankwrap/*bankwrap.Wrapper"

// SDKBindInterfaces returns a list of depinject.Configs that bind the bankwrap.Wrapper
// to the x/bank Keeper interfaces.
func SDKBindInterfaces() []depinject.Config {
	var resp []depinject.Config
	for _, override := range sdkOverrides {
		resp = append(resp, depinject.BindInterface(
			override,
			WrapperImpl,
		))
	}

	return resp
}
