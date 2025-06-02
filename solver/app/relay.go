package app

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/unibackend"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
)

// relayFunc abstracts the relay processing function.
type relayFunc func(ctx context.Context, req types.RelayRequest) (types.RelayResponse, error)

// RelayError represents a relay submission error with structured error information.
type RelayError struct {
	Code        string
	Message     string
	Description string
}

func (e RelayError) Error() string {
	return e.Message
}

// newRelayer creates a new relay function that can submit gasless orders on behalf of users.
// Note: Currently uses SolverNetInbox with a plan to upgrade to SolverNetInboxV2 when available
func newRelayer(
	inboxContracts map[uint64]*bindings.SolverNetInbox,
	backends unibackend.Backends,
	solverAddr common.Address,
) relayFunc {
	return func(ctx context.Context, req types.RelayRequest) (types.RelayResponse, error) {
		ctx, span := tracer.Start(ctx, "relay/submit_gasless_order")
		defer span.End()

		// Extract chain ID from the order
		chainID := req.Order.OriginChainId.ToInt().Uint64()

		// Get the appropriate inbox contract
		inbox, ok := inboxContracts[chainID]
		if !ok {
			return types.RelayResponse{}, RelayError{
				Code:        "UNSUPPORTED_CHAIN",
				Message:     "Chain not supported for relay",
				Description: "The specified origin chain is not supported by this relay service",
			}
		}

		// Get backend for the chain
		uniBackend, err := backends.Backend(chainID)
		if err != nil {
			return types.RelayResponse{}, RelayError{
				Code:        "BACKEND_ERROR",
				Message:     "Failed to get backend for chain",
				Description: err.Error(),
			}
		} else if !uniBackend.IsEVM() {
			return types.RelayResponse{}, RelayError{
				Code:        "UNSUPPORTED_BACKEND",
				Message:     "Only EVM chains are supported for relay",
				Description: "The specified chain does not have EVM support",
			}
		}

		// Validate the order before submission
		if err := validateGaslessOrder(req.Order); err != nil {
			return types.RelayResponse{}, RelayError{
				Code:        "INVALID_ORDER",
				Message:     "Order validation failed",
				Description: err.Error(),
			}
		}

		txOpts, err := uniBackend.EVMBackend().BindOpts(ctx, solverAddr)
		if err != nil {
			return types.RelayResponse{}, RelayError{
				Code:        "BACKEND_ERROR",
				Message:     "Failed to get backend for chain",
				Description: err.Error(),
			}
		}

		// Submit the gasless order via openFor (placeholder for when bindings are available)
		tx, err := inbox.OpenFor(txOpts, req.Order, req.Signature, req.OriginFillerData)
		if err != nil {
			return types.RelayResponse{}, RelayError{
				Code:        "SUBMISSION_FAILED",
				Message:     "Failed to submit gasless order",
				Description: errors.Format(err),
			}
		}

		// Wait for transaction confirmation
		rec, err := uniBackend.EVMBackend().WaitMined(ctx, tx)
		if err != nil {
			return types.RelayResponse{}, RelayError{
				Code:        "CONFIRMATION_FAILED",
				Message:     "Transaction submitted but confirmation failed",
				Description: err.Error(),
			}
		}

		// Extract order ID from transaction receipt
		orderID, err := extractOrderIDFromReceipt(rec, inbox)
		if err != nil {
			log.Warn(ctx, "Failed to extract order ID from receipt", err)
			// Don't fail the request since transaction was successful
		}

		log.Info(ctx, "Successfully relayed gasless order",
			"tx_hash", tx.Hash().Hex(),
			"order_id", orderID.Hex(),
			"user", req.Order.User.Hex(),
			"chain_id", chainID,
		)

		return types.RelayResponse{
			Success: true,
			TxHash:  tx.Hash(),
			OrderID: orderID,
		}, nil
	}
}

// validateGaslessOrder performs basic validation on the gasless order.
func validateGaslessOrder(order types.GaslessCrossChainOrder) error {
	if order.User == (common.Address{}) {
		return errors.New("user address cannot be zero")
	}
	if order.OriginSettler == (common.Address{}) {
		return errors.New("origin settler cannot be zero")
	}
	if order.Nonce == nil || order.Nonce.ToInt().Cmp(big.NewInt(0)) <= 0 {
		return errors.New("nonce must be positive")
	}
	if order.OriginChainId == nil || order.OriginChainId.ToInt().Cmp(big.NewInt(0)) <= 0 {
		return errors.New("origin chain ID must be positive")
	}
	if order.OpenDeadline <= uint32(0) {
		return errors.New("open deadline must be positive")
	}
	if order.FillDeadline <= order.OpenDeadline {
		return errors.New("fill deadline must be after open deadline")
	}
	if len(order.OrderData) == 0 {
		return errors.New("order data cannot be empty")
	}
	return nil
}

// extractOrderIDFromReceipt attempts to extract the order ID from the transaction receipt.
func extractOrderIDFromReceipt(rec *ethclient.Receipt, inbox *bindings.SolverNetInbox) (common.Hash, error) {
	// Look for the Open event in the receipt logs
	for _, log := range rec.Logs {
		if len(log.Topics) > 0 && log.Topics[0] == solvernet.TopicOpened {
			// The order ID is the first indexed parameter (topics[1])
			if len(log.Topics) > 1 {
				return log.Topics[1], nil
			}
		}
	}
	return common.Hash{}, errors.New("order ID not found in transaction receipt")
}
