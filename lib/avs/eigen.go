package avs

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/test/e2e/backend"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

//nolint:gochecknoglobals // Static mapping of zero contsants
var (
	zeroDelegationApprover       = common.Address{}
	zeroStakerOptOutWindowBlocks = uint32(0) // Currently unused by Eigen
)

// RegisterOperatorWithEigen registers the operator with the eigen layer delegation manager.
func RegisterOperatorWithEigen(ctx context.Context, avs Contracts, backend backend.Backend, operator common.Address, metadataURI string) (*ethtypes.Receipt, error) {
	operatorDetails := bindings.IDelegationManagerOperatorDetails{
		EarningsReceiver:         operator,
		DelegationApprover:       zeroDelegationApprover,
		StakerOptOutWindowBlocks: zeroStakerOptOutWindowBlocks,
	}

	txOpts, err := backend.BindOpts(ctx, operator)
	if err != nil {
		return nil, err
	}

	tx, err := avs.DelegationManager.RegisterAsOperator(txOpts, operatorDetails, metadataURI)
	if err != nil {
		return nil, errors.Wrap(err, "register as operator")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return nil, errors.Wrap(err, "wait mined")
	}

	return receipt, nil
}

// DelegateWETH delegates WETH to the Eigen Layer strategy manager's WETH strategy.
func DelegateWETH(ctx context.Context, contracts Contracts, backend backend.Backend, delegator common.Address, amount int64) error {
	txOpts, err := backend.BindOpts(ctx, delegator)
	if err != nil {
		return err
	}

	// First approve the strategy manager to "assign" the WETH to itself.
	tx, err := contracts.WETHToken.Approve(txOpts, contracts.StrategyManagerAddr, big.NewInt(amount))
	if err != nil {
		return errors.Wrap(err, "approve")
	}

	if _, err = backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	// Then deposit the WETH into the strategy (it will assign it to itself).
	tx, err = contracts.StrategyManager.DepositIntoStrategy(txOpts, contracts.WETHStrategyAddr, contracts.WETHTokenAddr, big.NewInt(amount))
	if err != nil {
		return errors.Wrap(err, "deposit into strategy")
	}

	if _, err = backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}

// Undelegate undelegates WETH from the Eigen Layer strategy manager's WETH strategy.
func Undelegate(ctx context.Context, contracts Contracts, backend backend.Backend, delegator common.Address) error {
	txOpts, err := backend.BindOpts(ctx, delegator)
	if err != nil {
		return err
	}

	tx, err := contracts.DelegationManager.Undelegate(txOpts, delegator)
	if err != nil {
		return errors.Wrap(err, "deposit into strategy")
	}

	if _, err = backend.WaitMined(ctx, tx); err != nil {
		return err
	}

	return nil
}
