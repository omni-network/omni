package sdk

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/ethereum/go-ethereum/params"

	sdklog "cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO(corver): Add all other consts/vars/funcs

const (
	ExecModeFinalize = sdk.ExecModeFinalize
)

var (
	// DefaultPowerReduction is the default amount of staking tokens required for 1 unit of consensus-engine power
	// Override default power reduction: 1 ether (1e18) $STAKE == 1 power.
	DefaultPowerReduction = sdkmath.NewInt(params.Ether)

	// DefaultBondDenom is the default bondable coin denomination (defaults to stake)
	// Overwriting this value has the side effect of changing the default denomination in genesis.
	DefaultBondDenom = sdk.DefaultBondDenom

	// MsgTypeURL returns the TypeURL of a `sdk.Msg`.
	MsgTypeURL = sdk.MsgTypeURL
)

type (
	Msg                        = sdk.Msg
	Context                    = sdk.Context
	ConsAddress                = sdk.ConsAddress
	ValAddress                 = sdk.ValAddress
	AccAddress                 = sdk.AccAddress
	ExtendVoteHandler          = sdk.ExtendVoteHandler
	VerifyVoteExtensionHandler = sdk.VerifyVoteExtensionHandler
	ProcessProposalHandler     = sdk.ProcessProposalHandler
	Tx                         = sdk.Tx
	Coin                       = sdk.Coin
	Coins                      = sdk.Coins
)

// ValAddressFromBech32 creates a ValAddress from a Bech32 string.
func ValAddressFromBech32(address string) (sdk.ValAddress, error) {
	addr, err := sdk.ValAddressFromBech32(address)
	if err != nil {
		return sdk.ValAddress{}, errors.Wrap(err, "failed to convert address")
	}

	return addr, nil
}

// UnwrapSDKContext retrieves a Context from a context.Context instance
// attached with WrapSDKContext. It panics if a Context was not properly
// attached.
func UnwrapSDKContext(ctx context.Context) sdk.Context {
	return sdk.UnwrapSDKContext(ctx)
}

// NewInt64Coin returns a new coin with a denomination and amount. It will panic
// if the amount is negative.
func NewInt64Coin(denom string, amount int64) sdk.Coin {
	return sdk.NewInt64Coin(denom, amount)
}

// NewCoin returns a new coin with a denomination and amount. It will panic if
// the amount is negative or if the denomination is invalid.
func NewCoin(denom string, amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(denom, amount)
}

// NewCoins constructs a new coin set. The provided coins will be sanitized by removing
// zero coins and sorting the coin set. A panic will occur if the coin set is not valid.
func NewCoins(coins ...sdk.Coin) sdk.Coins {
	return sdk.NewCoins(coins...)
}

// GetConfig returns the config instance for the SDK.
func GetConfig() *sdk.Config {
	return sdk.GetConfig()
}

// NewContext creates a new sdk.Context instance.
func NewContext(ms storetypes.MultiStore, header cmtproto.Header, isCheckTx bool, logger sdklog.Logger) sdk.Context {
	return sdk.NewContext(ms, header, isCheckTx, logger)
}
