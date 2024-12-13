package symbiotic

import (
	"bytes"
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	solver "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
)

var _ solver.Target = (*App)(nil)

type DepositArgs struct {
	Recipient common.Address
	Amount    *big.Int
}

func (App) Name() string {
	return "symbiotic"
}

func (a App) ChainID() uint64 {
	return a.L1.ChainID
}

func (a App) Address() common.Address {
	// Live ymbiotic deposits are for erc20 "collateral" wrappers, not vaults.
	return a.L1wstETHCollateral
}

func (a App) LogMetadata(ctx context.Context) {
	log.Info(ctx, "Target app",
		"name", a.Name(),
		"l1_chain", a.L1.Name,
		"l1_collateral", a.L1wstETHCollateral,
		"l1_token", a.L1wstETH,
		"l2_chain", a.L2.Name,
		"l2_token", a.L2wstETH,
	)
}

func (a App) TokenPrereqs(call bindings.SolveCall) ([]bindings.SolveTokenPrereq, error) {
	args, err := unpackDeposit(call.Data)
	if err != nil {
		return nil, errors.Wrap(err, "unpack deposit")
	}

	return []bindings.SolveTokenPrereq{
		{
			Token:   a.L1wstETH,
			Spender: a.L1wstETHCollateral,
			Amount:  args.Amount,
		},
	}, nil
}

func (a App) Verify(srcChainID uint64, call bindings.SolveCall, deposits []bindings.SolveDeposit) error {
	// for now, we only accept deposits from a single, explicit l2
	if srcChainID != a.L2.ChainID {
		return errors.New("source chain not supported", "src", srcChainID)
	}

	args, err := unpackDeposit(call.Data)
	if err != nil {
		return errors.Wrap(err, "invalid deposit")
	}

	if _, err := a.TokenPrereqs(call); err != nil {
		return errors.Wrap(err, "token prereqs")
	}

	var l2Deposit *bindings.SolveDeposit
	for _, deposit := range deposits {
		if deposit.Token == a.L2wstETH {
			l2Deposit = &deposit
		}
	}

	// if no l2 deposit, we can'a accept
	if l2Deposit == nil {
		return errors.New("no L2 token deposit")
	}

	// if l2 deposit amount does not match call amount, we can'a accept
	if l2Deposit.Amount.Cmp(args.Amount) < 0 {
		return errors.New("insufficient L2 token deposit",
			"expected", args.Amount,
			"actual", l2Deposit.Amount,
		)
	}

	// TODO: require native deposit that covers gas / risk / overhead

	return nil
}

func (App) LogCall(ctx context.Context, call bindings.SolveCall) error {
	dep, err := unpackDeposit(call.Data)
	if err != nil {
		return errors.Wrap(err, "unpack deposit")
	}

	log.Debug(ctx, "Symbiotic deposit", "method", "wstETH_collateral.deposit",
		"recipient", dep.Recipient, "amount", dep.Amount)

	return nil
}

func unpackDeposit(data []byte) (DepositArgs, error) {
	trimmed := bytes.TrimPrefix(data, depositABI.ID)
	if bytes.Equal(trimmed, data) {
		return DepositArgs{}, errors.New("data not prefixed with deposit method id")
	}

	unpacked, err := depositABI.Inputs.Unpack(trimmed)
	if err != nil {
		return DepositArgs{}, errors.Wrap(err, "unpack data")
	}

	var args DepositArgs
	if err := depositABI.Inputs.Copy(&args, unpacked); err != nil {
		return DepositArgs{}, errors.Wrap(err, "copy args")
	}

	return args, nil
}
