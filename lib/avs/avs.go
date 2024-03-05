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

	if _, err = backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
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

	if _, err = backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}

func verifyRegisterOperator(ctx context.Context, contracts Contracts, operator common.Address) error {
	callOpts := &bind.CallOpts{Context: ctx}

	ok, err := contracts.DelegationManager.IsOperator(callOpts, operator)
	if err != nil {
		return errors.Wrap(err, "is operator")
	} else if !ok {
		return errors.New("operator not registered with eigen layer delegation manager")
	}

	ok, err = contracts.OmniAVS.IsInAllowlist(callOpts, operator)
	if err != nil {
		return errors.Wrap(err, "is in allowlist")
	} else if !ok {
		return errors.New("operator not in omni avs allow list")
	}

	operators, err := contracts.OmniAVS.Operators(callOpts)
	if err != nil {
		return errors.Wrap(err, "operators")
	}
	for _, op := range operators {
		if op.Addr == operator {
			return errors.New("operator already registered with omni avs")
		}
	}

	return nil
}
