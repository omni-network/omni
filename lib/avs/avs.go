package avs

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/test/e2e/backend"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func RegisterOperatorWithAVS(ctx context.Context, contracts Contracts, backend backend.Backend, operator common.Address) error {
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

func DeregisterOperatorFromAVS(ctx context.Context, contracts Contracts, backend backend.Backend, operator common.Address) error {
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
