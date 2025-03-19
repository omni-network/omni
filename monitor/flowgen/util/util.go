package util

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// TODO (christian): consolidate with `ApproveToken` in the solver package.
// ApproveToken gives the 'contract' max allowance to spend the 'user's 'tokens'.
func ApproveToken(ctx context.Context, backend *ethbackend.Backend, token, user, contract common.Address) error {
	log.Debug(ctx, "Approving token", "token", token, "user", user, "contract", contract)

	erc20, err := bindings.NewIERC20(token, backend)
	if err != nil {
		return errors.Wrap(err, "new token")
	}

	isAppproved := func() (bool, error) {
		tkn, err := bindings.NewIERC20(token, backend)
		if err != nil {
			return false, errors.Wrap(err, "new token")
		}

		allowance, err := tkn.Allowance(&bind.CallOpts{Context: ctx}, user, contract)
		if err != nil {
			return false, errors.Wrap(err, "get allowance")
		}

		return umath.MaxUint256.Cmp(allowance) <= 0, nil
	}

	if approved, err := isAppproved(); err != nil {
		return err
	} else if approved {
		return nil
	}

	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return err
	}

	tx, err := erc20.Approve(txOpts, contract, umath.MaxUint256)
	if err != nil {
		return errors.Wrap(err, "approve token")
	} else if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	if approved, err := isAppproved(); err != nil {
		return err
	} else if !approved {
		return errors.New("approve failed")
	}

	return nil
}
