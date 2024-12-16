package devapp

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
	OnBehalfOf common.Address
	Amount     *big.Int
}

func (App) Name() string {
	return "devapp"
}

func (a App) ChainID() uint64 {
	return a.L1.ChainID
}

func (a App) Address() common.Address {
	return a.L1Vault
}

func (a App) LogMetadata(ctx context.Context) {
	log.Info(ctx, "Target app",
		"name", a.Name(),
		"l1_chain", a.L1.Name,
		"l1_token", a.L1Token,
		"l1_vault", a.L1Vault,
		"l2_chain", a.L2.Name,
		"l2_token", a.L2Token,
	)
}

func (a App) TokenPrereqs(call bindings.SolveCall) ([]bindings.SolveTokenPrereq, error) {
	args, err := unpackDeposit(call.Data)
	if err != nil {
		return nil, errors.Wrap(err, "unpack deposit")
	}

	return []bindings.SolveTokenPrereq{
		{
			Token:   a.L1Token,
			Spender: a.L1Vault,
			Amount:  args.Amount,
		},
	}, nil
}

func (a App) Verify(srcChainID uint64, call bindings.SolveCall, deposits []bindings.SolveDeposit) error {
	// we only accept deposits from mock L2
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

	var l2token *bindings.SolveDeposit
	for _, deposit := range deposits {
		if deposit.Token == a.L2Token {
			l2token = &deposit
		}
	}

	// if no l2 deposit, we can't accept
	if l2token == nil {
		return errors.New("no L2 token deposit")
	}

	// if l2 deposit amount does not match call amount, we can't accept
	if l2token.Amount.Cmp(args.Amount) != 0 {
		return errors.New("insufficient L2 token deposit")
	}

	// TODO: require native deposit that covers gas / risk / overhead

	return nil
}

func (a App) LogCall(ctx context.Context, call bindings.SolveCall) error {
	args, err := unpackDeposit(call.Data)
	if err != nil {
		return errors.Wrap(err, "unpack deposit")
	}

	if call.Target != a.L1Vault {
		return errors.New("unexpected target", "expected", a.L1Vault, "actual", call.Target)
	}

	log.Debug(ctx, "Devapp mock vault deposit", "on_behalf_of", args.OnBehalfOf, "amount", args.Amount, "target", call.Target)

	return nil
}

func unpackDeposit(data []byte) (DepositArgs, error) {
	trimmed := bytes.TrimPrefix(data, vaultDeposit.ID)
	if bytes.Equal(trimmed, data) {
		return DepositArgs{}, errors.New("data not prefixed with deposit method id")
	}

	unpacked, err := vaultDeposit.Inputs.Unpack(trimmed)
	if err != nil {
		return DepositArgs{}, errors.Wrap(err, "unpack data")
	}

	var args DepositArgs
	if err := vaultDeposit.Inputs.Copy(&args, unpacked); err != nil {
		return DepositArgs{}, errors.Wrap(err, "copy args")
	}

	return args, nil
}
