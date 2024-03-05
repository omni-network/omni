package avs

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

//nolint:gochecknoglobals // Static mapping of zero contsants
var (
	zeroDelegationApprover       = common.Address{}
	zeroStakerOptOutWindowBlocks = uint32(0) // Currently unused by Eigen
)

// RegisterOperatorWithEigen registers the operator with the eigen layer delegation manager.
func RegisterOperatorWithEigen(ctx context.Context, avs Contracts, backend *ethbackend.Backend, operator common.Address, metadataURI string) (*ethtypes.Receipt, error) {
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

// Undelegate undelegates WETH from the Eigen Layer strategy manager's WETH strategy.
func Undelegate(ctx context.Context, contracts Contracts, backend *ethbackend.Backend, delegator common.Address) error {
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
