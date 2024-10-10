package cchain

import (
	"crypto/ecdsa"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/umath"

	cmtcrypto "github.com/cometbft/cometbft/crypto"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	cosmosk1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sltypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/gogoproto/proto"
)

// PortalValidator is a consensus chain validator in a validator set emitted/submitted by/tp portals .
type PortalValidator struct {
	Address common.Address
	Power   int64
}

// Verify returns an error if the validator is invalid.
func (v PortalValidator) Verify() error {
	if v.Address == (common.Address{}) {
		return errors.New("empty validator address")
	}
	if v.Power <= 0 {
		return errors.New("invalid validator power")
	}

	return nil
}

// SDKSigningInfo wraps the cosmos slashing signing info type and extends it with convenience functions.
type SDKSigningInfo struct {
	sltypes.ValidatorSigningInfo
	// Uptime is the percentage [0,1] of blocks signed in the previous <SignedBlockWindow> (1000).
	// Note this is 100% if the validator isn't bonded, since it can't technically miss blocks.
	Uptime float64
}

func (i SDKSigningInfo) Jailed() bool {
	return i.JailedUntil.Unix() != 0
}

func (i SDKSigningInfo) ConsensusCmtAddr() (cmtcrypto.Address, error) {
	valAddr, err := sdk.ConsAddressFromBech32(i.Address)
	if err != nil {
		return cmtcrypto.Address{}, errors.Wrap(err, "parse validator address")
	} else if len(valAddr) != cmtcrypto.AddressSize {
		return nil, errors.New("invalid consensus address length")
	}

	return cmtcrypto.Address(valAddr), nil
}

// SDKValidator wraps the cosmos staking validator type and extends it with
// convenience functions.
type SDKValidator struct {
	stypes.Validator
}

// Power returns the validators cometBFT power.
func (v SDKValidator) Power() (uint64, error) {
	return umath.ToUint64(v.ConsensusPower(sdk.DefaultPowerReduction))
}

// OperatorEthAddr returns the validator operator ethereum address.
func (v SDKValidator) OperatorEthAddr() (common.Address, error) {
	opAddr, err := sdk.ValAddressFromBech32(v.OperatorAddress)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "parse operator address")
	} else if len(opAddr) != common.AddressLength {
		return common.Address{}, errors.New("invalid operator address length")
	}

	return cast.EthAddress(opAddr)
}

// ConsensusEthAddr returns the validator consensus eth address.
func (v SDKValidator) ConsensusEthAddr() (common.Address, error) {
	pk, err := v.ConsensusPublicKey()
	if err != nil {
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(*pk), nil
}

// ConsensusCmtAddr returns the validator consensus cometBFT address.
func (v SDKValidator) ConsensusCmtAddr() (cmtcrypto.Address, error) {
	pk := new(cosmosk1.PubKey)
	err := proto.Unmarshal(v.ConsensusPubkey.Value, pk)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal consensus pubkey")
	}

	if len(pk.Bytes()) != k1.PubKeySize {
		return nil, errors.Wrap(err, "invalid public key size after unmarshal consensus pubkey")
	}

	return pk.Address(), nil
}

// ConsensusPublicKey returns the validator consensus public key (eth ecdsa style).
func (v SDKValidator) ConsensusPublicKey() (*ecdsa.PublicKey, error) {
	pk := new(cosmosk1.PubKey)
	err := proto.Unmarshal(v.ConsensusPubkey.Value, pk)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal consensus pubkey")
	}

	pubkey, err := crypto.DecompressPubkey(pk.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "decompress pubkey")
	}

	return pubkey, nil
}
