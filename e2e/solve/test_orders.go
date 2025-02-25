package solve

import (
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/evmchain"
	solver "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

func makeOrders() []TestOrder {
	users := anvil.DevAccounts()

	orders := insufficientDepositTestOrders(users)
	orders = append(orders, expenseUnderMinOrOverMaxTestOrders(users)...)
	// orders = append(orders, unsupportedSrcChainTestOrder(users[0]))
	orders = append(orders, unsupportedDestChainTestOrder(users[0]))
	// orders = append(orders, sameChainTestOrder(users[0]))

	return orders
}

func expenseUnderMinOrOverMaxTestOrders(users []common.Address) []TestOrder {
	var orders []TestOrder

	// native ETH transfers
	for i, user := range users {
		amt := math.NewInt(1).MulRaw(params.Ether).BigInt()

		// make some under min or over max expense
		overMax := i < 3
		underMin := i > 6

		if overMax {
			// max is 1 ETH
			amt = math.NewInt(2).MulRaw(params.Ether).BigInt()
		}

		if underMin {
			// min is 0.001 ETH
			amt = big.NewInt(1)
		}

		var rejectReason string
		if underMin {
			rejectReason = solver.RejectExpenseUnderMin.String()
		}
		if overMax {
			rejectReason = solver.RejectExpenseOverMax.String()
		}

		order := TestOrder{
			Owner:         user,
			FillDeadline:  time.Now().Add(1 * time.Hour),
			SourceChainID: evmchain.IDMockL1,
			DestChainID:   evmchain.IDMockL2,
			Expenses:      nativeExpense(amt),
			Calls:         nativeTransferCall(amt, user),
			Deposit:       nativeDeposit(new(big.Int).Add(amt, big.NewInt(1e17))), // add enough to cover fee
			RejectReason:  rejectReason,
		}

		orders = append(orders, order)
	}

	return orders
}

func insufficientDepositTestOrders(users []common.Address) []TestOrder {
	var orders []TestOrder

	// erc20 OMNI -> native OMNI orders
	for i, user := range users {
		requestAmt := math.NewInt(10).MulRaw(params.Ether).BigInt()

		// make some insufficient (should reject)
		insufficientDeposit := i%2 == 0
		depositAmt := new(big.Int).Set(requestAmt)
		if insufficientDeposit {
			depositAmt = depositAmt.Div(depositAmt, big.NewInt(2))
		}

		var rejectReason string
		if insufficientDeposit {
			rejectReason = solver.RejectInsufficientDeposit.String()
		}

		order := TestOrder{
			Owner:         user,
			FillDeadline:  time.Now().Add(1 * time.Hour),
			SourceChainID: evmchain.IDMockL1,
			DestChainID:   evmchain.IDOmniDevnet,
			Expenses:      nativeExpense(requestAmt),
			Calls:         nativeTransferCall(requestAmt, user),
			Deposit:       erc20Deposit(depositAmt, addrs.Token),
			RejectReason:  rejectReason,
		}

		orders = append(orders, order)
	}

	return orders
}

func unsupportedDestChainTestOrder(user common.Address) TestOrder {
	return TestOrder{
		Owner:         user,
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   1234, // invalid
		Expenses:      nativeExpense(big.NewInt(1)),
		Calls:         nativeTransferCall(big.NewInt(1), user),
		Deposit:       erc20Deposit(big.NewInt(1), addrs.Token),
		RejectReason:  solver.RejectUnsupportedDestChain.String(),
	}
}

// func unsupportedSrcChainTestOrder(user common.Address) TestOrder {
//	return TestOrder{
//		Owner:         user,
//		FillDeadline:  time.Now().Add(1 * time.Hour),
//		SourceChainID: 1234, // invalid
//		DestChainID:   evmchain.IDMockL1,
//		Expenses:      nativeExpense(big.NewInt(1)),
//		Calls:         nativeTransferCall(big.NewInt(1), user),
//		Deposit:       erc20Deposit(big.NewInt(1), addrs.Token),
//		RejectReason:  solver.RejectUnsupportedSrcChain.String(),
//	}
// }
//
// func sameChainTestOrder(user common.Address) TestOrder {
//	return TestOrder{
//		Owner:         user,
//		FillDeadline:  time.Now().Add(1 * time.Hour),
//		SourceChainID: evmchain.IDMockL1,
//		DestChainID:   evmchain.IDMockL1,
//		Expenses:      nativeExpense(big.NewInt(1)),
//		Calls:         nativeTransferCall(big.NewInt(1), user),
//		Deposit:       erc20Deposit(big.NewInt(1), addrs.Token),
//		RejectReason:  solver.RejectSameChain.String(),
//	}
// }
