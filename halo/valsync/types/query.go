package types

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"

	"github.com/cometbft/cometbft/crypto"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"
)

func (v *Validator) CometPubkey() (k1.PubKey, error) {
	if len(v.ConsensusPubkey) != k1.PubKeySize {
		return nil, errors.New("invalid cometBFT public key size")
	}

	return v.ConsensusPubkey, nil
}

func (v *Validator) EthereumAddress() (common.Address, error) {
	pk, err := v.CometPubkey()
	if err != nil {
		return common.Address{}, err
	}

	return k1util.PubKeyToAddress(pk)
}

func (v *Validator) CometAddress() (crypto.Address, error) {
	pk, err := v.CometPubkey()
	if err != nil {
		return nil, err
	}

	return pk.Address(), nil
}

// IsValidator returns true if the provided address is a validator in the set.
func (v *ValidatorSetResponse) IsValidator(addr common.Address) (bool, error) {
	if len(v.Validators) == 0 {
		return false, errors.New("empty validators")
	}

	for _, val := range v.Validators {
		if val.Power == 0 {
			return false, errors.New("invalid active validator [BUG]")
		}

		ethAddr, err := val.EthereumAddress()
		if err != nil {
			return false, err
		}

		if ethAddr == addr {
			return true, nil
		}
	}

	return false, nil
}
