// Package sdk implements the Cosmos SDK configuration.
package sdk

import (
	"github.com/ethereum/go-ethereum/params"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Bech32HRP is the human-readable-part of the Bech32 address format.
const Bech32HRP = "omni"

// init initializes the Cosmos SDK configuration.
//
//nolint:gochecknoinits // Cosmos-style
func init() {
	// Set prefixes
	accountPubKeyPrefix := Bech32HRP + "pub"
	validatorAddressPrefix := Bech32HRP + "valoper"
	validatorPubKeyPrefix := Bech32HRP + "valoperpub"
	consNodeAddressPrefix := Bech32HRP + "valcons"
	consNodePubKeyPrefix := Bech32HRP + "valconspub"

	// Set and seal config
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(Bech32HRP, accountPubKeyPrefix)
	cfg.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	cfg.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	cfg.Seal()

	// Override default power reduction: 1 ether (1e18) $STAKE == 1 power.
	sdk.DefaultPowerReduction = sdkmath.NewInt(params.Ether)
}
