package avs

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func RegisterOperatorWithAVS(ctx context.Context, contracts Contracts, backend *ethbackend.Backend, operator common.Address) error {
	if err := verifyRegisterOperator(ctx, contracts, operator); err != nil {
		return errors.Wrap(err, "verify register operator")
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

	tx, err := contracts.OmniAVS.RegisterOperatorToAVS(txOpts, operator, bindings.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: sig[:],
		Salt:      salt,
		Expiry:    expiry,
	})
	if err != nil {
		return errors.Wrap(err, "register operator to avs")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return errors.New("receipt status failed")
	}

	return nil
}

func DeregisterOperatorFromAVS(ctx context.Context, contracts Contracts, backend *ethbackend.Backend, operator common.Address) error {
	txOpts, err := backend.BindOpts(ctx, operator)
	if err != nil {
		return err
	}

	tx, err := contracts.OmniAVS.DeregisterOperatorFromAVS(txOpts, operator)
	if err != nil {
		return errors.Wrap(err, "deregister operator from avs")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return errors.New("receipt status failed")
	}

	return nil
}

func verifyRegisterOperator(ctx context.Context, contracts Contracts, operator common.Address) error {
	canRegister, reason, err := contracts.OmniAVS.CanRegister(&bind.CallOpts{Context: ctx}, operator)
	if err != nil {
		return errors.Wrap(err, "can register")
	}

	if !canRegister {
		return errors.New("cannot register", "reason", reason)
	}

	return nil
}
