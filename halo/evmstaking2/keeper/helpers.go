package keeper

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// mustGetABI returns the metadata's ABI as an abi.ABI type.
// It panics on error.
func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

// mustGetEvent returns the event with the given name from the ABI.
// It panics if the event is not found.
func mustGetEvent(abi *abi.ABI, name string) abi.Event {
	event, ok := abi.Events[name]
	if !ok {
		panic("event not found")
	}

	return event
}

// omniToBondCoin converts the $OMNI amount into a $STAKE coin.
// TODO(corver): At this point, it is 1-to1, but this might change in the future.
func omniToBondCoin(amount *big.Int) (sdk.Coin, sdk.Coins) {
	coin := sdk.NewCoin(sdk.DefaultBondDenom, math.NewIntFromBigInt(amount))
	return coin, sdk.NewCoins(coin)
}
