package solvernet

import (
	"context"
	"crypto/ecdsa"
	"math"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// SolverNetInbox.ORDER_DATA_TYPEHASH
	// keccak256("OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)");.
	onchainOrderDataTypeHash = cast.Must32(hexutil.MustDecode("0x2e7de755ca70cb933dc80103af16cc3303580e5712f1a8927d6461441e99a1e6"))

	// keccak256("OmniOrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Call(address target,bytes4 selector,uint256 value,bytes params)Deposit(address token,uint96 amount)TokenExpense(address spender,address token,uint96 amount)");.
	OmniOrderDataTypeHash = cast.Must32(hexutil.MustDecode("0xfb3edd456f65c82530757ae6d57d2d138659a0a24917d9e9e0e1f5070a065069"))
)

type OpenOpts struct {
	FillDeadline time.Time
}

func WithFillDeadline(t time.Time) func(*OpenOpts) {
	return func(o *OpenOpts) {
		o.FillDeadline = t
	}
}

func DefaultOpenOpts() *OpenOpts {
	return &OpenOpts{
		FillDeadline: time.Now().Add(24 * time.Hour),
	}
}

// OpenOrder opens an order on chainID for user.
// user pays for the order, and must be in the backend for chainID.
// It returns the order id.
func OpenOrder(
	ctx context.Context,
	network netconf.ID,
	chainID uint64,
	backends ethbackend.Backends,
	user common.Address,
	orderData bindings.SolverNetOrderData,
	opts ...func(*OpenOpts),
) (OrderID, error) {
	backend, err := backends.Backend(chainID)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "get backend")
	}

	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "get addrs")
	}

	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "bind opts")
	}

	// if native deposit, add tx value
	deposit := orderData.Deposit
	if deposit.Token == (common.Address{}) {
		txOpts.Value = deposit.Amount
	}

	contract, err := bindings.NewSolverNetInbox(addrs.SolverNetInbox, backend)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "bind contract")
	}

	o := DefaultOpenOpts()
	for _, opt := range opts {
		opt(o)
	}

	packed, err := PackOrderData(orderData)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "pack order data")
	}

	// fill deadline is currently not enforced by the contract
	fillDeadline := o.FillDeadline.Unix()
	if fillDeadline < time.Now().Unix() {
		return OrderID{}, errors.New("fill deadline must be in the future")
	} else if fillDeadline > math.MaxUint32 {
		return OrderID{}, errors.New("fill deadline too far in the future")
	}

	order := bindings.IERC7683OnchainCrossChainOrder{
		//nolint:gosec // overflow is checked above
		FillDeadline:  uint32(fillDeadline),
		OrderData:     packed,
		OrderDataType: onchainOrderDataTypeHash,
	}

	tx, err := contract.Open(txOpts, order)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "open tx", "custom", DetectCustomError(err))
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "wait mined")
	}

	for _, log := range receipt.Logs {
		if log == nil {
			return OrderID{}, errors.New("invalid log")
		}

		orderID, status, err := ParseEvent(*log)
		if err == nil && status == StatusPending {
			return orderID, nil
		}
	}

	return OrderID{}, errors.New("no pending event in logs")
}

// OpenGaslessOpts contains options for opening gasless orders.
type OpenGaslessOpts struct {
	FillDeadline time.Time
	OpenDeadline time.Time
	Nonce        *big.Int
	SolverAddr   common.Address // Solver address that will relay the transaction
}

func WithGaslessFillDeadline(t time.Time) func(*OpenGaslessOpts) {
	return func(o *OpenGaslessOpts) {
		o.FillDeadline = t
	}
}

func WithGaslessOpenDeadline(t time.Time) func(*OpenGaslessOpts) {
	return func(o *OpenGaslessOpts) {
		o.OpenDeadline = t
	}
}

func WithGaslessNonce(nonce *big.Int) func(*OpenGaslessOpts) {
	return func(o *OpenGaslessOpts) {
		o.Nonce = nonce
	}
}

func WithSolverAddr(addr common.Address) func(*OpenGaslessOpts) {
	return func(o *OpenGaslessOpts) {
		o.SolverAddr = addr
	}
}

func DefaultOpenGaslessOpts() *OpenGaslessOpts {
	return &OpenGaslessOpts{
		FillDeadline: time.Now().Add(24 * time.Hour),
		OpenDeadline: time.Now().Add(1 * time.Hour),
		Nonce:        big.NewInt(1),    // Default nonce, should be managed properly
		SolverAddr:   common.Address{}, // Must be provided
	}
}

// OpenGaslessOrder prepares a gasless order and returns the relay request that should be submitted.
// Unlike OpenOrder, this doesn't submit directly but returns the components needed for relay submission.
func OpenGaslessOrder(
	ctx context.Context,
	network netconf.ID,
	chainID uint64,
	backends ethbackend.Backends,
	userKey *ecdsa.PrivateKey, // User's private key for signing
	orderData bindings.SolverNetOrderData,
	opts ...func(*OpenGaslessOpts),
) (bindings.IERC7683GaslessCrossChainOrder, []byte, error) {
	backend, err := backends.Backend(chainID)
	if err != nil {
		return bindings.IERC7683GaslessCrossChainOrder{}, nil, errors.Wrap(err, "get backend")
	}

	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return bindings.IERC7683GaslessCrossChainOrder{}, nil, errors.Wrap(err, "get addrs")
	}

	// Gasless orders don't support native deposits, only ERC20
	if orderData.Deposit.Token == (common.Address{}) {
		return bindings.IERC7683GaslessCrossChainOrder{}, nil, errors.New("gasless orders only support ERC20 deposits")
	}

	o := DefaultOpenGaslessOpts()
	for _, opt := range opts {
		opt(o)
	}

	if o.SolverAddr == (common.Address{}) {
		return bindings.IERC7683GaslessCrossChainOrder{}, nil, errors.New("solver address must be provided")
	}

	packed, err := PackOrderData(orderData)
	if err != nil {
		return bindings.IERC7683GaslessCrossChainOrder{}, nil, errors.Wrap(err, "pack order data")
	}

	// Validate deadlines
	fillDeadline := o.FillDeadline.Unix()
	openDeadline := o.OpenDeadline.Unix()
	if openDeadline < time.Now().Unix() {
		return bindings.IERC7683GaslessCrossChainOrder{}, nil, errors.New("open deadline must be in the future")
	}
	if fillDeadline <= openDeadline {
		return bindings.IERC7683GaslessCrossChainOrder{}, nil, errors.New("fill deadline must be after open deadline")
	}
	if fillDeadline > math.MaxUint32 || openDeadline > math.MaxUint32 {
		return bindings.IERC7683GaslessCrossChainOrder{}, nil, errors.New("deadlines too far in the future")
	}

	// Get user address from private key
	userAddr := crypto.PubkeyToAddress(userKey.PublicKey)

	// Create the gasless cross-chain order
	gaslessOrder := bindings.IERC7683GaslessCrossChainOrder{
		OriginSettler: addrs.SolverNetInbox,
		User:          userAddr,
		Nonce:         o.Nonce,
		OriginChainId: new(big.Int).SetUint64(chainID),
		//nolint:gosec // overflow is checked above
		OpenDeadline: uint32(openDeadline),
		//nolint:gosec // overflow is checked above
		FillDeadline:  uint32(fillDeadline),
		OrderDataType: OmniOrderDataTypeHash, // Use the newer typehash
		OrderData:     packed,
	}

	// Generate the EIP-712 signature for Permit2
	signature, err := GeneratePermit2Signature(ctx, backend, userKey, gaslessOrder, orderData, addrs.SolverNetInbox)
	if err != nil {
		return bindings.IERC7683GaslessCrossChainOrder{}, nil, errors.Wrap(err, "generate permit2 signature")
	}

	return gaslessOrder, signature, nil
}
