package symbiotic

import (
	"context"
	"fmt"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/uni"
	"github.com/omni-network/omni/monitor/flowgen/types"
	"github.com/omni-network/omni/monitor/flowgen/util"
	sclient "github.com/omni-network/omni/solver/client"
	stypes "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	_ "embed"
)

//go:embed symbiotic_abi.json
var abiJSON []byte

// newJob returns a symbiotic deposit job if a config for the given network exists, or it returns false.
func newJob(
	ctx context.Context,
	backends ethbackend.Backends,
	networkID netconf.ID,
	scl sclient.Client,
) (types.Job, bool, error) {
	addrs, err := contracts.GetAddresses(ctx, networkID)
	if err != nil {
		return types.Job{}, false, errors.Wrap(err, "contract addresses")
	}

	conf, ok := config[networkID]
	if !ok {
		return types.Job{}, false, nil
	}

	backend, err := backends.Backend(conf.srcChain)
	if err != nil {
		return types.Job{}, false, errors.Wrap(err, "src chain backend")
	}

	token := tokens.WSTETH

	srcChainTkn, ok := tokens.BySymbol(conf.srcChain, token.Symbol)
	if !ok {
		return types.Job{}, false, errors.Wrap(err, "src token not found")
	}

	dstChainTkn, ok := tokens.BySymbol(conf.dstChain, token.Symbol)
	if !ok {
		return types.Job{}, false, errors.Wrap(err, "dst token not found")
	}

	flowgenAddr := eoa.MustAddress(networkID, eoa.RoleFlowgen)

	if err := util.ApproveToken(ctx, backend, srcChainTkn.Address, flowgenAddr, addrs.SolverNetInbox); err != nil {
		return types.Job{}, false, errors.Wrap(err, "token approval")
	}

	namer := netconf.ChainNamer(networkID)
	cadence := 30 * time.Minute
	if networkID == netconf.Devnet {
		cadence = time.Second * 10
	}

	return types.Job{
		Name:       fmt.Sprintf("Symbiotic deposit (%v->%v)", namer(conf.srcChain), namer(conf.dstChain)),
		Cadence:    cadence,
		SrcChainID: conf.srcChain,
		OpenOrdersFunc: func(ctx context.Context) ([]types.Result, error) {
			result, ok, err := openOrder(ctx, scl, backends, networkID, flowgenAddr, srcChainTkn, dstChainTkn, conf)
			if err != nil {
				return nil, errors.Wrap(err, "open order")
			} else if !ok {
				return nil, nil
			}

			return []types.Result{result}, nil
		},
	}, true, nil
}

// openOrder returns the order id if an order was opened successfully,
// it returns false if no order was opened or an error in case of an error.
func openOrder(
	ctx context.Context,
	scl sclient.Client,
	backends ethbackend.Backends,
	networkID netconf.ID,
	owner common.Address,
	srcToken, dstToken tokens.Token,
	conf flowConfig,
) (types.Result, bool, error) {
	quoteReq := stypes.QuoteRequest{
		SourceChainID:      srcToken.ChainID,
		DestinationChainID: dstToken.ChainID,
		Deposit: stypes.AddrAmt{
			Token: uni.EVMAddress(srcToken.Address),
			// Amount left empty, quote will return the required amount.
		},
		Expense: stypes.AddrAmt{
			Token:  uni.EVMAddress(dstToken.Address),
			Amount: conf.orderAmount,
		},
	}

	quoteResp, err := scl.Quote(ctx, quoteReq)
	if err != nil {
		return types.Result{}, false, errors.Wrap(err, "quote deposit")
	} else if quoteResp.Rejected {
		return types.Result{}, false, errors.New("quote rejected", "description", quoteResp.RejectDescription, "reason", quoteResp.RejectCode)
	}

	abi, err := metaData.GetAbi()
	if err != nil {
		return types.Result{}, false, errors.Wrap(err, "get abi")
	}

	data, err := abi.Pack("deposit", owner, quoteReq.Expense.Amount)
	if err != nil {
		return types.Result{}, false, errors.Wrap(err, "packing")
	}

	orderData := bindings.SolverNetOrderData{
		Owner:       owner,
		DestChainId: conf.dstChain,
		Deposit: solvernet.Deposit{
			Token:  quoteResp.Deposit.Token.EVM(),
			Amount: quoteResp.Deposit.Amount,
		},
		Expenses: []solvernet.Expense{{
			Token:   quoteReq.Expense.Token.EVM(),
			Amount:  quoteReq.Expense.Amount,
			Spender: conf.vaultAddr,
		}},
		Calls: []bindings.SolverNetCall{
			solvernet.Call{
				Target: conf.vaultAddr,
				Data:   data,
				Value:  bi.Zero(),
			}.ToBinding(),
		},
	}

	orderID, err := solvernet.OpenOrder(ctx, networkID, conf.srcChain, backends, owner, orderData)
	if err != nil {
		return types.Result{}, false, errors.Wrap(err, "open order")
	}

	return types.Result{OrderID: orderID, Data: orderData}, true, nil
}

var metaData = &bind.MetaData{
	ABI: string(abiJSON),
}

// Jobs creates the following jobs:
// - deposit wstETH from the source to the destination chain.
func Jobs(
	ctx context.Context,
	backends ethbackend.Backends,
	networkID netconf.ID,
	scl sclient.Client,
) ([]types.Job, error) {
	var jobs []types.Job
	job, ok, err := newJob(ctx, backends, networkID, scl)
	if err != nil {
		return jobs, errors.Wrap(err, "symbiotic job")
	}
	if ok {
		jobs = append(jobs, job)
	}

	return jobs, nil
}
