package app

import (
	"context"
	"fmt"
	"math/big"
	"time"

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

// Relay error codes and messages.
const (
	// Error codes.
	RelayErrorInvalidOrder         = "INVALID_ORDER"
	RelayErrorMissingSignature     = "MISSING_SIGNATURE"
	RelayErrorUnsupportedChain     = "UNSUPPORTED_CHAIN"
	RelayErrorBackendError         = "BACKEND_ERROR"
	RelayErrorUnsupportedBackend   = "UNSUPPORTED_BACKEND"
	RelayErrorInvalidOriginSettler = "INVALID_ORIGIN_SETTLER"
	RelayErrorInvalidOrderData     = "INVALID_ORDER_DATA"
	RelayErrorSubmissionFailed     = "SUBMISSION_FAILED"

	// Error messages.
	RelayMsgOrderValidationFailed  = "Order validation failed"
	RelayMsgSignatureRequired      = "Signature is required for gasless orders"
	RelayMsgChainNotSupported      = "Chain not supported for relay"
	RelayMsgBackendFailed          = "Failed to get backend for chain"
	RelayMsgEVMOnly                = "Only EVM chains are supported for relay"
	RelayMsgOriginSettlerMismatch  = "Origin settler mismatch"
	RelayMsgDecodeOrderDataFailed  = "Failed to decode order data"
	RelayMsgConvertOrderDataFailed = "Failed to convert order data"
	RelayMsgOrderRejected          = "Order rejected by solver"
	RelayMsgSubmissionFailed       = "Failed to submit gasless order"

	// Error descriptions.
	RelayDescSignatureEmpty    = "The signature field cannot be empty"
	RelayDescChainNotSupported = "The specified origin chain is not supported"
	RelayDescEVMRequired       = "The specified chain does not have EVM support"
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
// It integrates with the existing check logic to pre-validate orders before submission.
func newRelayer(
	inboxContracts map[uint64]*bindings.SolverNetInbox,
	backends unibackend.Backends,
	solverAddr common.Address,
	inboxAddr common.Address,
	checkFunc checkFunc,
) relayFunc {
	return func(ctx context.Context, req types.RelayRequest) (types.RelayResponse, error) {
		ctx, span := tracer.Start(ctx, "relay/submit_gasless_order")
		defer span.End()

		// Extract chain ID from the order
		chainID := req.Order.OriginChainId.Uint64()

		// Basic validation of the gasless order structure
		if err := validateGaslessOrder(req.Order); err != nil {
			return types.RelayResponse{
				Success: false,
				Error: &types.RelayError{
					Code:        RelayErrorInvalidOrder,
					Message:     RelayMsgOrderValidationFailed,
					Description: err.Error(),
				},
			}, nil
		}

		// Validate that the signature is present
		if len(req.Signature) == 0 {
			return types.RelayResponse{
				Success: false,
				Error: &types.RelayError{
					Code:        RelayErrorMissingSignature,
					Message:     RelayMsgSignatureRequired,
					Description: RelayDescSignatureEmpty,
				},
			}, nil
		}

		// Get the appropriate inbox contract
		inbox, ok := inboxContracts[chainID]
		if !ok {
			return types.RelayResponse{
				Success: false,
				Error: &types.RelayError{
					Code:        RelayErrorUnsupportedChain,
					Message:     RelayMsgChainNotSupported,
					Description: RelayDescChainNotSupported,
				},
			}, nil
		}

		// Get backend for the chain
		uniBackend, err := backends.Backend(chainID)
		if err != nil {
			return types.RelayResponse{
				Success: false,
				Error: &types.RelayError{
					Code:        RelayErrorBackendError,
					Message:     RelayMsgBackendFailed,
					Description: err.Error(),
				},
			}, nil
		} else if !uniBackend.IsEVM() {
			return types.RelayResponse{
				Success: false,
				Error: &types.RelayError{
					Code:        RelayErrorUnsupportedBackend,
					Message:     RelayMsgEVMOnly,
					Description: RelayDescEVMRequired,
				},
			}, nil
		}

		// Validate that the inbox address matches the order's origin settler
		if req.Order.OriginSettler != inboxAddr {
			return types.RelayResponse{
				Success: false,
				Error: &types.RelayError{
					Code:        RelayErrorInvalidOriginSettler,
					Message:     RelayMsgOriginSettlerMismatch,
					Description: fmt.Sprintf("Expected %s, got %s", inboxAddr.Hex(), req.Order.OriginSettler.Hex()),
				},
			}, nil
		}

		// Decode the order data
		orderData, err := solvernet.ParseOrderData(req.Order.OrderData)
		if err != nil {
			return types.RelayResponse{
				Success: false,
				Error: &types.RelayError{
					Code:        RelayErrorInvalidOrderData,
					Message:     RelayMsgDecodeOrderDataFailed,
					Description: err.Error(),
				},
			}, nil
		}

		// Convert to CheckRequest for order validation
		checkReq, err := types.CheckRequestFromOrderData(chainID, orderData)
		if err != nil {
			return types.RelayResponse{
				Success: false,
				Error: &types.RelayError{
					Code:        RelayErrorInvalidOrderData,
					Message:     RelayMsgConvertOrderDataFailed,
					Description: err.Error(),
				},
			}, nil
		}

		// types.CheckRequestFromOrderData overrides fill deadline for 1 hour from now
		// so we need to override it with the fill deadline from the gasless order
		checkReq.FillDeadline = req.Order.FillDeadline

		// Validate the order using existing check logic
		if err := checkFunc(ctx, checkReq); err != nil {
			// Check if it's an order rejection vs an actual error
			if r := new(RejectionError); errors.As(err, &r) {
				return types.RelayResponse{
					Success: false,
					Error: &types.RelayError{
						Code:        r.Reason.String(),
						Message:     RelayMsgOrderRejected,
						Description: errors.Format(r.Err),
					},
				}, nil
			}
			// It's an actual error, not a rejection
			return types.RelayResponse{}, err
		}

		// All validations passed, submit the gasless order
		txOpts, err := uniBackend.EVMBackend().BindOpts(ctx, solverAddr)
		if err != nil {
			return types.RelayResponse{}, errors.Wrap(err, "failed to get bind opts")
		}

		// Submit the gasless order via openFor
		tx, err := inbox.OpenFor(txOpts, req.Order, req.Signature, req.OriginFillerData)
		if err != nil {
			// Any openFor failure is treated as a submission error
			return types.RelayResponse{
				Success: false,
				Error: &types.RelayError{
					Code:        RelayErrorSubmissionFailed,
					Message:     RelayMsgSubmissionFailed,
					Description: errors.Format(err),
				},
			}, nil
		}

		// Wait for transaction confirmation
		rec, err := uniBackend.EVMBackend().WaitMined(ctx, tx)
		if err != nil {
			return types.RelayResponse{}, errors.Wrap(err, "transaction submitted but confirmation failed")
		}

		// Extract order ID from transaction receipt
		orderID, err := extractOrderIDFromReceipt(rec)
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
func validateGaslessOrder(order bindings.IERC7683GaslessCrossChainOrder) error {
	if order.User == (common.Address{}) {
		return errors.New("user address cannot be zero")
	}
	if order.OriginSettler == (common.Address{}) {
		return errors.New("origin settler cannot be zero")
	}
	if order.Nonce == nil || order.Nonce.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("nonce must be positive")
	}
	if order.OriginChainId == nil || order.OriginChainId.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("origin chain ID must be positive")
	}
	if order.OpenDeadline <= uint32(0) {
		return errors.New("open deadline must be positive")
	}
	// Check if the open deadline has already expired
	if order.OpenDeadline < uint32(time.Now().Unix()) {
		return errors.New("open deadline has already expired")
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
func extractOrderIDFromReceipt(rec *ethclient.Receipt) (common.Hash, error) {
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
