package mybank

import (
	"fmt"

	"cosmossdk.io/depinject"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type DIInputs struct {
	depinject.In

	bankKeeper *bankkeeper.Keeper
}

type DIOutputs struct {
	depinject.Out

	WrappedBankKeeper Keeper
}

func DIProvide(input DIInputs) (DIOutputs, error) {
	fmt.Println("DEBUQ", input)
	k := Keeper{
		Keeper: *input.bankKeeper,
	}

	return DIOutputs{
		WrappedBankKeeper: k,
	}, nil
}
