package symbiotic

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/monitor/flowgen/types"
	"github.com/omni-network/omni/monitor/flowgen/util"
	"github.com/omni-network/omni/solver/app"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	_ "embed"
)

//go:embed symbiotic_abi.json
var abiJSON []byte

func NewJob(
	ctx context.Context,
	backends ethbackend.Backends,
	network netconf.ID,
	inboxAddr common.Address,
	srcChain,
	dstChain uint64,
	srcChainToken common.Address,
	dstChainToken common.Address,
	symbioticContractAddress common.Address,
	owner common.Address,
	amount *big.Int,
) (types.Job, error) {
	backend, err := backends.Backend(srcChain)
	if err != nil {
		return types.Job{}, errors.Wrap(err, "get backend")
	}

	if err := util.ApproveToken(ctx, backend, srcChainToken, owner, inboxAddr); err != nil {
		return types.Job{}, errors.Wrap(err, "token approval")
	}

	data, err := orderData(owner, srcChain, dstChain, srcChainToken, dstChainToken, symbioticContractAddress, amount)
	if err != nil {
		return types.Job{}, errors.Wrap(err, "new job")
	}

	namer := netconf.ChainNamer(network)

	return types.Job{
		Name:    fmt.Sprintf("Symbiotic deposit (%v->%v)", namer(srcChain), namer(dstChain)),
		Cadence: 30 * time.Minute,
		Network: network,

		SrcChain: srcChain,
		DstChain: dstChain,

		Owner: owner,

		OrderData: data,
	}, nil
}

// orderData returns the order data required to do the job.
func orderData(
	owner common.Address,
	srcChain, dstChain uint64,
	srcChainToken common.Address,
	dstChainToken common.Address,
	symbioticContractAddress common.Address,
	amount *big.Int,
) (bindings.SolverNetOrderData, error) {
	srcToken, ok := app.AllTokens().Find(srcChain, srcChainToken)
	if !ok {
		return bindings.SolverNetOrderData{}, errors.New("src token not found", "token", srcToken.Address)
	}
	dstToken, ok := app.AllTokens().Find(dstChain, dstChainToken)
	if !ok {
		return bindings.SolverNetOrderData{}, errors.New("dst token not found", "token", dstToken.Address)
	}

	// Tokens that will be deposited to the user on the destination chain.
	expense := app.TokenAmt{Token: dstToken, Amount: amount}

	depositWithFee, err := app.QuoteDeposit(srcToken, app.TokenAmt{Token: srcToken, Amount: amount})
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "quote deposit")
	}

	abi, err := metaData.GetAbi()
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "get abi")
	}

	data, err := abi.Pack("deposit", owner, expense.Amount)
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "packing")
	}

	orderData := bindings.SolverNetOrderData{
		Owner:       owner,
		DestChainId: dstChain,
		Deposit: solvernet.Deposit{
			Token:  depositWithFee.Token.Address,
			Amount: depositWithFee.Amount,
		},
		Expenses: []solvernet.Expense{{
			Token:   expense.Token.Address,
			Amount:  expense.Amount,
			Spender: symbioticContractAddress,
		}},
		Calls: []bindings.SolverNetCall{
			solvernet.Call{
				Target: symbioticContractAddress,
				Data:   data,
				Value:  new(big.Int),
			}.ToBinding(),
		},
	}

	return orderData, nil
}

var metaData = &bind.MetaData{
	ABI: string(abiJSON),
}
