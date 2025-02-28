package solvernet

import (
	"context"
	"math"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	// SolverNetInbox.ORDER_DATA_TYPEHASH
	// keccak256("OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)");.
	orderDataTypeHash = cast.Must32(hexutil.MustDecode("0x2e7de755ca70cb933dc80103af16cc3303580e5712f1a8927d6461441e99a1e6"))
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
		OrderDataType: orderDataTypeHash,
	}

	// Simulate the call first.
	callOpts := bind.CallOpts{
		From:    txOpts.From,
		Context: txOpts.Context,
	}
	contractCaller := bindings.SolverNetInboxRaw{
		Contract: contract,
	}
	err = contractCaller.Call(&callOpts, nil, "open", order)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "order simulation")
	}

	tx, err := contract.Open(txOpts, order)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "open tx", "custom", DetectCustomError(err))
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "wait mined")
	}

	logs := receipt.Logs
	lastLog := func() *ethtypes.Log { return logs[len(logs)-1] }
	if len(logs) < 1 || len(lastLog().Topics) < 2 {
		return OrderID{}, errors.New("unexpeced open order logs [BUG]")
	}

	// order id is first topic of Open(...) log, which is the last log
	orderID := lastLog().Topics[1]

	return OrderID(orderID), nil
}
