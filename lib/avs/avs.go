package avs

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/k1util"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func RegisterOperatorWithAVS(ctx context.Context, contracts Contracts, backend *ethbackend.Backend, operator common.Address) error {
	if err := verifyRegisterOperator(ctx, contracts, operator); err != nil {
		return err
	}

	salt := crypto.Keccak256Hash(operator.Bytes())         // Salt can be anything, it should just be unique.
	expiry := big.NewInt(time.Now().Add(time.Hour).Unix()) // Sig is 1 Hour valid

	digestHash, err := contracts.AVSDirectory.CalculateOperatorAVSRegistrationDigestHash(&bind.CallOpts{},
		operator,
		contracts.OmniAVSAddr,
		salt,
		expiry)
	if err != nil {
		return errors.Wrap(err, "calculate digest hash")
	}

	sig, err := backend.Sign(operator, digestHash)
	if err != nil {
		return errors.Wrap(err, "sign")
	}

	txOpts, err := backend.BindOpts(ctx, operator)
	if err != nil {
		return err
	}

	pubkey, err := backend.PublicKey(operator)
	if err != nil {
		return errors.Wrap(err, "public key")
	}

	pubkey64 := k1util.PubKeyToBytes64(pubkey)

	tx, err := contracts.OmniAVS.RegisterOperator(txOpts, pubkey64, bindings.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: sig[:],
		Salt:      salt,
		Expiry:    expiry,
	})
	if err != nil {
		return errors.Wrap(err, "register operator to avs")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}

func DeregisterOperatorFromAVS(ctx context.Context, contracts Contracts, backend *ethbackend.Backend, operator common.Address) error {
	txOpts, err := backend.BindOpts(ctx, operator)
	if err != nil {
		return err
	}

	tx, err := contracts.OmniAVS.DeregisterOperator(txOpts)
	if err != nil {
		return errors.Wrap(err, "deregister operator from avs")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}

func verifyRegisterOperator(ctx context.Context, contracts Contracts, operator common.Address) error {
	canRegister, reason, err := contracts.OmniAVS.CanRegister(&bind.CallOpts{Context: ctx}, operator)
	if err != nil {
		return errors.Wrap(err, "can register")
	}

	if !canRegister {
		return errors.New(reason)
	}

	return nil
}
