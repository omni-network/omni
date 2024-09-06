// Package sdk wraps the cosmos-sdk/types package (via dot-import) and initializes the Cosmos SDK configuration.
// Always use this package instead of cosmos-sdk/types.
package sdk

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Bech32HRP is the human-readable-part of the cosmos bech32 address format.
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
}
