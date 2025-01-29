package keeper

import (
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

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

// catch executes the function, returning an error if it panics.
func catch(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("recovered", "panic", r)
		}
	}()

	return fn()
}

// eventName returns the name of the EVM event log or "unknown".
func eventName(elog *evmenginetypes.EVMEvent) string {
	const unknown = "unknown"

	ethlog, err := elog.ToEthLog()
	if err != nil {
		return unknown
	}

	event, ok := eventsByID[ethlog.Topics[0]]
	if !ok {
		return unknown
	}

	return event.Name
}
