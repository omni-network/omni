package keeper

import (
	"github.com/omni-network/omni/lib/k1util"

	abci "github.com/cometbft/cometbft/abci/types"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/cometbft/cometbft/proto/tendermint/crypto"

	"github.com/ethereum/go-ethereum/common"
)

// ValidatorUpdate returns the validator as a cometBFT validator update.
func (v *Validator) ValidatorUpdate() abci.ValidatorUpdate {
	return abci.UpdateValidator(
		v.GetPubKey(),
		v.GetPower(),
		k1.KeyType,
	)
}

// CometPubkey returns the cometBFT public key of the validator.
func (v *Validator) CometPubkey() crypto.PublicKey {
	return v.ValidatorUpdate().PubKey
}

// Address returns the validator ethereum 20 byte address.
func (v *Validator) Address() (common.Address, error) {
	return k1util.PubKeyPBToAddress(v.CometPubkey())
}
