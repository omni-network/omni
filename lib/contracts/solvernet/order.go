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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	// SolverNetInbox.ORDER_DATA_TYPEHASH
	// keccak256("OrderData(Call call,Deposit[] deposits)Call(uint64 chainId,bytes32 target,uint256 value,bytes data,TokenExpense[] expenses)TokenExpense(bytes32 token,bytes32 spender,uint256 amount)Deposit(bytes32 token,uint256 amount)").
	orderDataTypeHash = cast.Must32(hexutil.MustDecode("0xe5be6bd381a38cd250f9aa1a05935cbcd261fe0e77e9ef6f6d07bf3b7e5d22e2"))
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
func OpenOrder(
	ctx context.Context,
	network netconf.ID,
	chainID uint64,
	backends ethbackend.Backends,
	user common.Address,
	orderData bindings.ISolverNetOrderData,
	opts ...func(*OpenOpts),
) error {
	backend, err := backends.Backend(chainID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	contract, err := bindings.NewSolverNetInbox(addrs.SolverNetInbox, backend)
	if err != nil {
		return errors.Wrap(err, "bind contract")
	}

	o := DefaultOpenOpts()
	for _, opt := range opts {
		opt(o)
	}

	packed, err := PackOrderData(orderData)
	if err != nil {
		return errors.Wrap(err, "pack order data")
	}

	// fill deadline is currently not enforced by the contract
	fillDeadline := o.FillDeadline.Unix()
	if fillDeadline < time.Now().Unix() {
		return errors.New("fill deadline must be in the future")
	} else if fillDeadline > math.MaxUint32 {
		return errors.New("fill deadline too far in the future")
	}

	order := bindings.IERC7683OnchainCrossChainOrder{
		//nolint:gosec // overflow is checked above
		FillDeadline:  uint32(fillDeadline),
		OrderData:     packed,
		OrderDataType: orderDataTypeHash,
	}

	tx, err := contract.Open(txOpts, order)
	if err != nil {
		return errors.Wrap(err, "open tx", "custom", DetectCustomError(err))
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}
