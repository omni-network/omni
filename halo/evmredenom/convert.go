package evmredenom

import (
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Factor is the conversion factor between $NATIVE_EVM tokens and bonded $STAKE tokens.
const Factor = 75 // 75 $NATIVE_EVM == 1 bonded $STAKE

// ToStakeCoin converts the $NATIVE_EVM amount into a $STAKE coin.
func ToStakeCoin(amount *big.Int) sdk.Coin {
	n := bi.DivRaw(amount, Factor)
	return sdk.NewCoin(sdk.DefaultBondDenom, math.NewIntFromBigInt(n))
}

// ToEVMAmount converts a $STAKE coin into a $NATIVE_EVM amount.
func ToEVMAmount(coin sdk.Coin) (*big.Int, error) {
	if coin.Denom != sdk.DefaultBondDenom {
		return nil, errors.New("not bond denom [BUG]", "denom", coin.Denom)
	}

	return bi.MulRaw(coin.Amount.BigInt(), Factor), nil
}
