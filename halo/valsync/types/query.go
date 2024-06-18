package types

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"

	"github.com/cometbft/cometbft/crypto"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (v *Validator) CometPubkey() (k1.PubKey, error) {
	if len(v.ConsensusPubkey) != k1.PubKeySize {
		return nil, errors.New("invalid cometBFT public key size")
	}

	return v.ConsensusPubkey, nil
}

// ConsEthereumAddress returns the ethereum address of the validator.
func (v *Validator) EthConsensusAddress() (common.Address, error) {
	pk, err := v.CometPubkey()
	if err != nil {
		return common.Address{}, err
	}

	return k1util.PubKeyToAddress(pk)
}

func (v *Validator) EthOperatorAddress() (common.Address, error) {
	addr, err := sdk.ValAddressFromBech32(v.OperatorAddr)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "operator address")
	}

	return common.BytesToAddress(addr.Bytes()), nil
}

func (v *Validator) CometConsensusAddress() (crypto.Address, error) {
	pk, err := v.CometPubkey()
	if err != nil {
		return nil, err
	}

	return pk.Address(), nil
}
