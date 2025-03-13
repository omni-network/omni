package util

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/umath"
)

var (
	// ether1 is 1 token in wei (18 decimals).
	Ether1 = dec(1, 18)

	// million Gwei.
	MilliEther = dec(1, 15)
)

func dec(amt float64, decimals int) *big.Int {
	unit := math.Pow10(decimals)

	p := amt * unit

	_, dec := math.Modf(p)
	if dec != 0 {
		panic(fmt.Sprintf("amt float64 must be an int multiple of 1e%d", decimals))
	}

	return new(big.Int).SetUint64(uint64(p))
}

// ApproveToken gives `user` max allowance for `token`.
func ApproveToken(ctx context.Context, backend *ethbackend.Backend, token, user, contract common.Address) error {
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
