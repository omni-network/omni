package bankwrap

import (
	"github.com/omni-network/omni/octane/evmengine/types"

	"cosmossdk.io/depinject"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type DIInputs struct {
	depinject.In

	BankKeeper bankkeeper.BaseKeeper
}

type DIOutputs struct {
	depinject.Out

	Wrapper Wrapper
}

func DependencyInjector(wrapper *Wrapper, wk *types.WithdrawalKeeper) {
	if wrapper == nil || wk == nil {
		return
	}

	wrapper.WithdrawalKeeper = *wk
}

func DIProvide(input DIInputs) (DIOutputs, error) {
	return DIOutputs{
		Wrapper: Wrapper{
			Keeper: input.BankKeeper,
		},
	}, nil
}

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

const WrapperImpl = "github.com/omni-network/omni/octane/bankwrap/bankwrap.Wrapper"

// SDKBindInterfaces returns a list of depinject.Configs that bind the bankwrap.Wrapper
// to the x/bank Keeper interface.
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
