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

// RegisterOperatorWithAVS registers an operator with the AVS contract.
// Note RegisterOperatorWithAVS is tested indirectly via cli tests, see cli/cmd/register_test.go.
func RegisterOperatorWithAVS(ctx context.Context, addr common.Address, backend *ethbackend.Backend, operator common.Address) error {
	avs, err := bindings.NewOmniAVS(addr, backend)
	if err != nil {
		return errors.Wrap(err, "new avs")
	}

	if err := verifyRegisterOperator(ctx, avs, operator); err != nil {
		return err
	}

	salt := crypto.Keccak256Hash(operator.Bytes())         // Salt can be anything, it should just be unique.
	expiry := big.NewInt(time.Now().Add(time.Hour).Unix()) // Sig is 1 Hour valid

	avsDirAddr, err := avs.AvsDirectory(&bind.CallOpts{})
	if err != nil {
		return errors.Wrap(err, "avs directory")
	}

	avsDir, err := bindings.NewAVSDirectory(avsDirAddr, backend)
	if err != nil {
		return errors.Wrap(err, "new avs directory")
	}

	digestHash, err := avsDir.CalculateOperatorAVSRegistrationDigestHash(&bind.CallOpts{},
		operator,
		addr,
		salt,
		expiry)
	if err != nil {
		return errors.Wrap(err, "calculate digest hash")
	}

	sig, err := backend.Sign(ctx, operator, digestHash)
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

	pubkey64, err := k1util.PubKeyToBytes64(pubkey)
	if err != nil {
		return err
	}

	tx, err := avs.RegisterOperator(txOpts, pubkey64, bindings.ISignatureUtilsSignatureWithSaltAndExpiry{
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

func verifyRegisterOperator(ctx context.Context, avs *bindings.OmniAVS, operator common.Address) error {
	canRegister, reason, err := avs.CanRegister(&bind.CallOpts{Context: ctx}, operator)
	if err != nil {
		return errors.Wrap(err, "can register")
	}

	if !canRegister {
		return errors.New(reason)
	}

	return nil
}
