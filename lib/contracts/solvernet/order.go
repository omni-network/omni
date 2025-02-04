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
	// keccak256( "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,Expense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)Expense(address spender,address token,uint96 amount)");.
	orderDataTypeHash = cast.Must32(hexutil.MustDecode("0xfb2fd076b12750c4bfa77af573b3791b17a5e94805b6300eff13aae1c9ebaeb0"))
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
	orderData bindings.SolverNetOrderData,
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
