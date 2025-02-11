package withdraw

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

	BankWrapper *BankWrapper
}

func DIProvide(input DIInputs) (DIOutputs, error) {
	return DIOutputs{
		BankWrapper: NewBankWrapper(input.BankKeeper),
	}, nil
}

// DIInvoke injects the EVMEngineKeeper into the BankWrapper.
// This mitigates cyclic dependency: EVMEngineKeeper -> EVMEventProc(EVMStakingKeeper) -> BankWrapper -> EVMEngineKeeper.
func DIInvoke(bankWrapper *BankWrapper, evmEngKeeper EVMEngineKeeper) {
	if evmEngKeeper == nil {
		panic("nil EVMEngineKeeper")
	} else if bankWrapper == nil {
		panic("nil BankWrapper")
	}

	bankWrapper.SetEVMEngineKeeper(evmEngKeeper)
}

// overrides is a list of x/bank Keeper interfaces to be overridden by the BankWrapper.
// Only modules that call SendCoinsFromModuleToAccount or UndelegateCoinsFromModuleToAccount need to be overridden.
// Note that the `gov` and `protocolpool` modules would also require this if/when used.
// Note that these modules require additional Burner permission to burn coins.
var overrides = []string{
	// SendCoinsFromModuleToAccount
	"github.com/cosmos/cosmos-sdk/x/distribution/types/types.BankKeeper",
	// UndelegateCoinsFromModuleToAccount
	"github.com/cosmos/cosmos-sdk/x/staking/types/types.BankKeeper",

	// Note the x/auth Keepers don't actually need to be wrapped, but
	// it fails with "No type for explicit binding found" if added to noOverrides :(
	"github.com/cosmos/cosmos-sdk/x/auth/types/types.BankKeeper",
	"github.com/cosmos/cosmos-sdk/x/auth/tx/config/tx.BankKeeper",
}

// noOverrides is a list of standard cosmosSDK module x/bank Keeper interfaces
// that do not need to be overridden since they don't call transfer methods.
// They remain bound to the default x/bank Keeper.
var noOverrides = []string{
	"github.com/cosmos/cosmos-sdk/x/bank/keeper/keeper.Keeper",
	"github.com/cosmos/cosmos-sdk/x/mint/types/types.BankKeeper",
	"github.com/cosmos/cosmos-sdk/x/slashing/types/types.BankKeeper",
}

const (
	defaultBankKeeper = "github.com/cosmos/cosmos-sdk/x/bank/keeper/keeper.BaseKeeper"
	bankWrapper       = "github.com/omni-network/omni/halo/withdraw/*withdraw.BankWrapper"
)

// BindInterfaces returns a list of depinject.BindInterface that bind the x/bank Keeper
// interfaces to either this BankWrapper or the default x/bank keeper.
//
// Note that the `gov` and `protocolpool` modules would also require this if/when used.
// Other SDK modules don't call transfer methods so they don't need this.
func BindInterfaces() []depinject.Config {
	var resp []depinject.Config

	// Bind the overrides to the BankWrapper.
	for _, iface := range overrides {
		resp = append(resp, depinject.BindInterface(iface, bankWrapper))
	}

	// Bind other modules to the default x/bank Keeper.
	for _, iface := range noOverrides {
		resp = append(resp, depinject.BindInterface(iface, defaultBankKeeper))
	}

	return resp
}
