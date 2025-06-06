//nolint:contextcheck // Not critical for testing
package solve

import (
	"context"
	"time"

	"github.com/omni-network/omni/anchor/anchorinbox"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/svmutil"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func TestSVM(ctx context.Context, testnet types.Testnet) error {
	if len(testnet.SVMChains) == 0 {
		return nil
	} else if len(testnet.SVMChains) > 1 {
		return errors.New("multiple SVM chains")
	} else if testnet.Network != netconf.Devnet {
		return errors.New("not devnet")
	}

	svmChain := testnet.SVMChains[0]
	cl := rpc.New(svmChain.ExternalRPC)

	for i, svmOrder := range makeSVMOrders(ctx, testnet.Network) {
		txResp, err := svmutil.SendSimple(ctx, cl, svmOrder.PrivateKey, svmOrder.Build())
		if err != nil {
			return err
		}

		_, err = svmutil.AwaitConfirmedTransaction(ctx, cl, txResp)
		if err != nil {
			return err
		}

		for {
			state, _, err := anchorinbox.GetOrderState(ctx, cl, svmOrder.OpenOrder.ID)
			if err != nil {
				return err
			}

			if state.Status == anchorinbox.StatusRejected {
				return errors.New("svm order rejected", "index", i, "order_id", svmOrder.ID(), "status", state.Status, "reason", state.RejectReason)
			} else if state.Status != anchorinbox.StatusClaimed {
				log.Debug(ctx, "SVM order status", "index", i, "order_id", svmOrder.ID(), "status", state.Status)
				time.Sleep(time.Second) // Wait before checking again

				continue
			}

			break
		}
	}

	return nil
}

type svmOrder struct {
	anchorinbox.OpenOrder
	PrivateKey solana.PrivateKey
}

func (o svmOrder) ID() solvernet.OrderID {
	return solvernet.OrderID(o.OpenOrder.ID)
}

func makeSVMOrders(ctx context.Context, network netconf.ID) []svmOrder {
	var keys []solana.PrivateKey
	for _, role := range []eoa.Role{eoa.RoleFlowgen, eoa.RoleHot} {
		pk, err := eoa.PrivateKey(ctx, network, role)
		if err != nil {
			panic(errors.Wrap(err, "get private key for role", "role", role, "network", network))
		}
		keys = append(keys, svmutil.MapEVMKey(pk))
	}

	return []svmOrder{
		// Create order 1: svm USDC -> mockL1 USDC
		mustSVMOrder(
			keys[0],
			mustToken(evmchain.IDSolanaLocal, tokens.USDC),
			mustToken(evmchain.IDMockL1, tokens.USDC),
			10, // 10 USDC
		),
		// Create order 2: svm USDC -> mockL2 ETH
		mustSVMOrder(
			keys[1],
			mustToken(evmchain.IDSolanaLocal, tokens.USDC),
			mustToken(evmchain.IDMockL2, tokens.WSTETH),
			0.01, // 0.01 WSTETH
		),
	}
}

func mustSVMOrder(owner solana.PrivateKey, src, dst tokens.Token, expenseF64 float64) svmOrder {
	params, err := svmTokenParams(src, dst, expenseF64)
	if err != nil {
		panic(errors.Wrap(err, "create SVM order params"))
	}

	openOrder, err := anchorinbox.NewOpenOrder(params, owner.PublicKey(), src.SVMAddress)
	if err != nil {
		panic(errors.Wrap(err, "create SVM open order"))
	}

	return svmOrder{
		OpenOrder:  openOrder,
		PrivateKey: owner,
	}
}

func mustToken(chainID uint64, asset tokens.Asset) tokens.Token {
	token, ok := tokens.ByAsset(chainID, asset)
	if !ok {
		panic(errors.New("get token by asset", "chain_id", chainID, "asset", asset))
	}

	return token
}

func svmTokenParams(src, dst tokens.Token, expenseF64 float64) (anchorinbox.OpenParams, error) {
	expenseAmount := dst.F64ToAmt(expenseF64)
	depositAmount := mustDepositAmount(expenseAmount, src, dst)

	expenses, calls := expenseAndCall(expenseAmount, dst, common.Address{}) // Send to zero address for now

	svmCall, err := toSVMCall(calls)
	if err != nil {
		return anchorinbox.OpenParams{}, err
	}

	svmExpense, err := toSVMExpense(expenses)
	if err != nil {
		return anchorinbox.OpenParams{}, err
	}

	return anchorinbox.OpenParams{
		DepositAmount: depositAmount.Uint64(),
		DestChainId:   dst.ChainID,
		Call:          svmCall,
		Expense:       svmExpense,
	}, nil
}

func toSVMExpense(expenses []solvernet.Expense) (anchorinbox.EVMTokenExpense, error) {
	if len(expenses) != 1 {
		return anchorinbox.EVMTokenExpense{}, errors.New("expected exactly one expense")
	}
	expense := expenses[0]

	if expense.Amount == nil {
		expense.Amount = bi.Zero()
	}
	amount, err := svmutil.U128(expense.Amount)
	if err != nil {
		return anchorinbox.EVMTokenExpense{}, err
	}

	return anchorinbox.EVMTokenExpense{
		Spender: expense.Spender,
		Token:   expense.Token,
		Amount:  amount,
	}, nil
}

func toSVMCall(calls []solvernet.Call) (anchorinbox.EVMCall, error) {
	if len(calls) != 1 {
		return anchorinbox.EVMCall{}, errors.New("expected exactly one call")
	}
	call := calls[0].ToBinding()

	if call.Value == nil {
		call.Value = bi.Zero()
	}
	value, err := svmutil.U128(call.Value)
	if err != nil {
		return anchorinbox.EVMCall{}, err
	}

	return anchorinbox.EVMCall{
		Target:   call.Target,
		Selector: call.Selector,
		Value:    value,
		Params:   call.Params,
	}, nil
}
