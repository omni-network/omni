//nolint:unused // Partially integrated
package app

import (
	"context"

	"github.com/omni-network/omni/anchor/anchorinbox"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/svmutil"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// svmGetOrder retrieves the order from the Solana chain.
// It converts the order state to the required Order struct.
func svmGetOrder(ctx context.Context, cl *rpc.Client, outboxAddr common.Address, orderID OrderID) (Order, bool, error) {
	state, ok, err := anchorinbox.GetOrderState(ctx, cl, solana.PublicKey(orderID))
	if err != nil || !ok {
		return Order{}, ok, err
	}

	chainID, err := svmutil.ChainID(ctx, cl)
	if err != nil {
		return Order{}, false, err
	}

	status, err := toEVMStatus(state.Status)
	if err != nil {
		return Order{}, false, err
	}

	minReceived := []bindings.IERC7683Output{
		{
			Token:     state.DepositMint,
			Amount:    bi.N(state.DepositAmount),
			Recipient: [32]byte{}, // N/A
			ChainId:   bi.N(chainID),
		},
	}

	maxSpent := []bindings.IERC7683Output{
		{
			Token:     toBz32(state.DestExpense.Token),
			Amount:    state.DestExpense.Amount.BigInt(),
			Recipient: [32]byte{}, // N/A
			ChainId:   bi.N(state.DestChainId),
		},
	}

	fillData, err := anchorinbox.FillData(chainID, state)
	if err != nil {
		return Order{}, false, err
	}

	encoded, err := solvernet.EncodeFillData(fillData)
	if err != nil {
		return Order{}, false, err
	}

	return Order{
		ID:            orderID,
		SourceChainID: chainID,
		Status:        status,
		Offset:        0,                // N/A
		UpdatedBy:     common.Address{}, // N/A
		pendingData: PendingData{
			MinReceived:        minReceived,
			DestinationSettler: outboxAddr,
			DestinationChainID: state.DestChainId,
			FillOriginData:     encoded,
			MaxSpent:           maxSpent,
		},
		filledData: FilledData{
			MinReceived: minReceived,
		},
	}, true, nil
}

func toEVMStatus(status anchorinbox.Status) (solvernet.OrderStatus, error) {
	switch status {
	case anchorinbox.StatusPending:
		return solvernet.StatusPending, nil
	case anchorinbox.StatusFilled:
		return solvernet.StatusFilled, nil
	case anchorinbox.StatusClaimed:
		return solvernet.StatusClaimed, nil
	case anchorinbox.StatusClosed:
		return solvernet.StatusClosed, nil
	case anchorinbox.StatusRejected:
		return solvernet.StatusRejected, nil
	default:
		return 0, errors.New("invalid order status", "status", status)
	}
}

func claimSVMOrder(
	ctx context.Context,
	cl *rpc.Client,
	claimer solana.PrivateKey,
	orderID OrderID,
) error {
	claim, err := anchorinbox.NewClaimOrder(ctx, cl, claimer.PublicKey(), solana.PublicKey(orderID))
	if err != nil {
		return err
	}

	sig, err := svmutil.SendSimple(ctx, cl, claimer, claim.Build())
	if err != nil {
		return err
	}

	_, err = svmutil.AwaitConfirmedTransaction(ctx, cl, sig)
	if err != nil {
		return err
	}

	return nil
}

func rejectSVMOrder(
	ctx context.Context,
	cl *rpc.Client,
	admin solana.PrivateKey,
	orderID OrderID,
	reason types.RejectReason,
) error {
	reject, err := anchorinbox.NewRejectOrder(ctx, cl, admin.PublicKey(), solana.PublicKey(orderID), uint8(reason))
	if err != nil {
		return err
	}

	sig, err := svmutil.SendSimple(ctx, cl, admin, reject.Build())
	if err != nil {
		return err
	}

	_, err = svmutil.AwaitConfirmedTransaction(ctx, cl, sig)
	if err != nil {
		return err
	}

	return nil
}

func MarkFilledSVMOrder(
	ctx context.Context,
	cl *rpc.Client,
	admin solana.PrivateKey,
	claimableBy solana.PublicKey,
	orderID OrderID,
) error {
	mark, err := anchorinbox.NewMarkFilledOrder(ctx, cl, claimableBy, admin.PublicKey(), solana.PublicKey(orderID))
	if err != nil {
		return err
	}

	sig, err := svmutil.SendSimple(ctx, cl, admin, mark.Build())
	if err != nil {
		return err
	}

	_, err = svmutil.AwaitConfirmedTransaction(ctx, cl, sig)
	if err != nil {
		return err
	}

	return nil
}
