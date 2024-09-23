package keeper

import (
	"github.com/omni-network/omni/lib/errors"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
)

// Validate returns an error if the validator is invalid.
func (v *Validator) Validate() error {
	if v == nil {
		return errors.New("nil validator")
	} else if v.GetPower() < 0 {
		return errors.New("negative power")
	} else if len(v.GetPubKey()) != k1.PubKeySize {
		return errors.New("invalid pubkey size")
	} else if v.GetValsetId() != 0 {
		return errors.New("valset id already set")
	}

	return nil
}
