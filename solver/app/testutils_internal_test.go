package app

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/ethclient/mock"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var (
	erc20ABI  = mustGetABI(bindings.IERC20MetaData)
	inboxABI  = mustGetABI(bindings.SolverNetInboxMetaData)
	outboxABI = mustGetABI(bindings.SolverNetOutboxMetaData)
)

type OrderData = bindings.SolverNetOrderData

type testOrder struct {
	srcChainID uint64
	dstChainID uint64
	calls      []types.Call // calls not tested explicitly, but required to test native expense
	deposits   []types.AddrAmt
	expenses   []types.Expense
}

// orderTestCase is a test case for both shouldReject and quote handlers.
type orderTestCase struct {
	name         string
	reason       types.RejectReason
	reject       bool
	fillReverts  bool
	disallowCall bool
	shouldErr    bool
	mock         func(clients MockClients)
	order        testOrder
	testdata     bool
}

// rejectTestCase is a test case for shouldReject(...)
type rejectTestCase struct {
	name         string
	reason       types.RejectReason
	reject       bool
	fillReverts  bool
	disallowCall bool
	shouldErr    bool
	mock         func(clients MockClients)
	order        Order
}

// checkTestCase is test case for quote handler.
type checkTestCase struct {
	name         string
	disallowCall bool
	mock         func(clients MockClients)
	req          types.CheckRequest
	res          types.CheckResponse
	testdata     bool
}

func toCheckTestCase(t *testing.T, tt orderTestCase) checkTestCase {
	t.Helper()

	require.Len(t, tt.order.deposits, 1)

	deposit := tt.order.deposits[0]

	var rejectReason string
	if tt.reject {
		rejectReason = tt.reason.String()
	}

	return checkTestCase{
		name:         tt.name,
		mock:         tt.mock,
		disallowCall: tt.disallowCall,
		testdata:     tt.testdata,
		req: types.CheckRequest{
			SourceChainID:      tt.order.srcChainID,
			DestinationChainID: tt.order.dstChainID,
			Calls:              tt.order.calls,
			Expenses:           tt.order.expenses,
			Deposit:            deposit,
		},
		res: types.CheckResponse{
			Accepted:     !tt.reject,
			Rejected:     tt.reject,
			RejectReason: rejectReason,
		},
	}
}

func toRejectTestCase(t *testing.T, tt orderTestCase, outbox common.Address) rejectTestCase {
	t.Helper()

	// NOTE: in tests we map expenses directly to Order.MaxSpent.
	// if there is a native expense, it should be included in expenses.
	// this differs from inbox order opening / resolution, where native
	// expenses are derived from sum of native call value

	var maxSpent []bindings.IERC7683Output
	for _, e := range tt.order.expenses {
		maxSpent = append(maxSpent, bindings.IERC7683Output{
			Amount:    e.Amount,
			Token:     toBz32(e.Token),
			Recipient: toBz32(outbox),
			ChainId:   bi.N(tt.order.dstChainID),
		})
	}

	var minReceived []bindings.IERC7683Output
	for _, d := range tt.order.deposits {
		minReceived = append(minReceived, bindings.IERC7683Output{
			Amount:  d.Amount,
			Token:   toBz32(d.Token),
			ChainId: bi.N(tt.order.srcChainID),
		})
	}

	fillOriginData, err := solvernet.PackFillOriginData(
		bindings.SolverNetFillOriginData{
			SrcChainId:   tt.order.srcChainID,
			DestChainId:  tt.order.dstChainID,
			FillDeadline: uint32(4294967295), // max, does not matter
			Calls:        types.CallsToBindings(tt.order.calls),
			Expenses:     types.ExpensesToBindings(tt.order.expenses),
		},
	)
	require.NoError(t, err)

	return rejectTestCase{
		name:         tt.name,
		reason:       tt.reason,
		reject:       tt.reject,
		fillReverts:  tt.fillReverts,
		disallowCall: tt.disallowCall,
		shouldErr:    tt.shouldErr,
		mock:         tt.mock,
		order: Order{
			ID:            [32]byte{0x01},
			SourceChainID: tt.order.srcChainID,
			Status:        solvernet.StatusPending,
			pendingData: PendingData{
				DestinationChainID: tt.order.dstChainID,
				DestinationSettler: outbox,
				MaxSpent:           maxSpent,
				MinReceived:        minReceived,
				FillOriginData:     fillOriginData,
			},
			filledData: FilledData{
				MinReceived: minReceived,
			},
		},
	}
}

func checkTestCases(t *testing.T, solver common.Address) []checkTestCase {
	t.Helper()

	var tests []checkTestCase

	for _, tt := range orderTestCases(t, solver) {
		if len(tt.order.deposits) != 1 {
			// quote requires single deposit token
			continue
		}

		tests = append(tests, toCheckTestCase(t, tt))
	}

	additional := []checkTestCase{
		{
			name: "unsupported source chain",
			req: types.CheckRequest{
				SourceChainID:      1234567,
				DestinationChainID: evmchain.IDHolesky,
			},
			res: types.CheckResponse{
				Rejected:     true,
				RejectReason: types.RejectUnsupportedSrcChain.String(),
			},
		},
		{
			name: "same chain",
			req: types.CheckRequest{
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDHolesky,
			},
			res: types.CheckResponse{
				Rejected:     true,
				RejectReason: types.RejectSameChain.String(),
			},
		},
	}

	return append(tests, additional...)
}

func rejectTestCases(t *testing.T, solver, outbox common.Address) []rejectTestCase {
	t.Helper()

	var tests []rejectTestCase
	for _, tt := range orderTestCases(t, solver) {
		tests = append(tests, toRejectTestCase(t, tt, outbox))
	}

	additional := []rejectTestCase{
		// special case: insufficient native OMNI should error, not reject
		toRejectTestCase(t, orderTestCase{
			name:      "insufficient native OMNI",
			shouldErr: true,
			order: testOrder{
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDOmniOmega,
				deposits:   []types.AddrAmt{{Amount: ether(1), Token: omniERC20(netconf.Omega).Address}},
				calls:      []types.Call{{Value: ether(1)}},
				expenses:   []types.Expense{{Amount: ether(1)}},
			},
			mock: func(clients MockClients) {
				mockNativeBalance(t, clients.Client(t, evmchain.IDOmniOmega), solver, ether(0))
			},
		}, outbox),
	}

	return append(tests, additional...)
}

func erc20(chainID uint64, asset tokens.Asset) tokens.Token {
	token, ok := tokens.ByAsset(chainID, asset)
	if !ok {
		panic("OMNI token not found")
	}

	return token
}

func omniERC20(network netconf.ID) tokens.Token {
	token, ok := tokens.BySymbol(netconf.EthereumChainID(network), "OMNI")
	if !ok {
		panic("OMNI token not found")
	}

	return token
}

func orderTestCases(t *testing.T, solver common.Address) []orderTestCase {
	t.Helper()

	omegaOMNIAddr := omniERC20(netconf.Omega).Address
	holeskySTETH := common.HexToAddress("0x3f1c547b21f65e10480de3ad8e19faac46c95034")
	arbSepoliaUSDC := common.HexToAddress("0x75faf114eafb1BDbe2F0316DF893fd58CE46AA4d")
	baseSepoliaUSDC := common.HexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e")

	// dummy calldata / target. unused but for /check calls, and to build valid FillOriginData
	dummyCallData := hexutil.MustDecode("0x70a08231000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc7")

	return []orderTestCase{
		{
			name:   "insufficient native balance",
			reason: types.RejectInsufficientInventory,
			reject: true,
			order: testOrder{
				// erqeust 1 ETH for 2 ETH (large deposit to cover fee)
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDBaseSepolia,
				deposits:   []types.AddrAmt{{Amount: ether(2)}},
				calls:      []types.Call{{Value: ether(1)}},
				expenses:   []types.Expense{{Amount: ether(1)}},
			},
			mock: func(clients MockClients) {
				mockNativeBalance(t, clients.Client(t, evmchain.IDBaseSepolia), solver, ether(0))
			},
			testdata: true,
		},
		{
			name:   "sufficient native balance",
			reason: types.RejectNone,
			reject: false,
			order: testOrder{
				// request 1 native OMNI for 1 erc20 OMNI on omega
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDOmniOmega,
				// OMNI does not require fee
				deposits: []types.AddrAmt{{Amount: ether(1), Token: omegaOMNIAddr}},
				calls:    []types.Call{{Value: ether(1)}},
				expenses: []types.Expense{{Amount: ether(1)}},
			},
			mock: func(clients MockClients) {
				mockNativeBalance(t, clients.Client(t, evmchain.IDOmniOmega), solver, ether(1))
			},
			testdata: true,
		},
		{
			name:   "insufficient ERC20 balance",
			reason: types.RejectInsufficientInventory,
			reject: true,
			order: testOrder{
				// request 1 erc20 OMNI for 1 native OMNI on omega
				srcChainID: evmchain.IDOmniOmega,
				dstChainID: evmchain.IDHolesky,
				deposits:   []types.AddrAmt{{Amount: ether(1)}},
				calls:      []types.Call{{Data: dummyCallData}},
				expenses:   []types.Expense{{Amount: ether(1), Token: omegaOMNIAddr}},
			},
			mock: func(clients MockClients) {
				mockERC20Balance(t, clients.Client(t, evmchain.IDHolesky), omegaOMNIAddr, ether(0))
			},
		},
		{
			name:   "sufficient ERC20 balance",
			reason: types.RejectNone,
			reject: false,
			order: testOrder{
				// request 1 erc20 OMNI for 1 native OMNI on omega
				srcChainID: evmchain.IDOmniOmega,
				dstChainID: evmchain.IDHolesky,
				// OMNI does not require fee
				deposits: []types.AddrAmt{{Amount: ether(1)}},
				calls:    []types.Call{{Data: dummyCallData}},
				expenses: []types.Expense{{Amount: ether(1), Token: omegaOMNIAddr}},
			},
			mock: func(clients MockClients) {
				mockERC20Balance(t, clients.Client(t, evmchain.IDHolesky), omegaOMNIAddr, ether(1))
				mockERC20Allowance(t, clients.Client(t, evmchain.IDHolesky), omegaOMNIAddr)
			},
			testdata: true,
		},
		{
			name:        "fill reverts",
			fillReverts: true,
			reason:      types.RejectDestCallReverts,
			reject:      true,

			// rest same as above
			order: testOrder{
				// request 1 erc20 OMNI for 1 native OMNI on omega
				srcChainID: evmchain.IDOmniOmega,
				dstChainID: evmchain.IDHolesky,
				deposits:   []types.AddrAmt{{Amount: ether(1)}},
				calls:      []types.Call{{Data: dummyCallData}},
				expenses:   []types.Expense{{Amount: ether(1), Token: omegaOMNIAddr}},
			},
			mock: func(clients MockClients) {
				mockERC20Balance(t, clients.Client(t, evmchain.IDHolesky), omegaOMNIAddr, ether(1))
				mockERC20Allowance(t, clients.Client(t, evmchain.IDHolesky), omegaOMNIAddr)
			},
		},
		{
			name:   "unsupported expense token",
			reason: types.RejectUnsupportedExpense,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDOmniOmega,
				dstChainID: evmchain.IDHolesky,
				deposits:   []types.AddrAmt{{Amount: ether(1)}},
				calls:      []types.Call{{Data: dummyCallData}},
				expenses:   []types.Expense{{Amount: ether(1), Token: common.HexToAddress("0x01")}},
			},
		},
		{
			name:   "unsupported dest chain",
			reason: types.RejectUnsupportedDestChain,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDOmniOmega,
				dstChainID: 1234567,
			},
		},
		{
			name:   "invalid deposit (native token mismatch)",
			reason: types.RejectInvalidDeposit,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDOmniOmega,
				deposits:   []types.AddrAmt{{Amount: ether(1)}},
				calls:      []types.Call{{Value: ether(1)}},
				expenses:   []types.Expense{{Amount: ether(1)}},
			},
		},
		{
			name:   "invalid deposit (token mismatch)",
			reason: types.RejectInvalidDeposit,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDBaseSepolia,
				// wstETH on holesky
				deposits: []types.AddrAmt{{Amount: ether(1), Token: common.HexToAddress("0x8d09a4502cc8cf1547ad300e066060d043f6982d")}},
				// native eth on base
				calls:    []types.Call{{Value: ether(1)}},
				expenses: []types.Expense{{Amount: ether(1)}},
			},
		},
		{
			name:   "invalid deposit (multiple tokens)",
			reason: types.RejectInvalidDeposit,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDOmniOmega,
				deposits:   []types.AddrAmt{{Amount: ether(1), Token: omegaOMNIAddr}, {Amount: ether(1)}},
				calls:      []types.Call{{Value: ether(1)}},
				expenses:   []types.Expense{{Amount: ether(1)}},
			},
		},
		{
			name:   "invalid deposit (mismatch chain class)",
			reason: types.RejectInvalidDeposit,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDOptimism,
				deposits:   []types.AddrAmt{{Amount: ether(1)}},
				calls:      []types.Call{{Value: ether(1)}},
				expenses:   []types.Expense{{Amount: ether(1)}},
			},
		},
		{
			name:   "invalid expense (token and native)",
			reason: types.RejectInvalidExpense,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDOmniOmega,
				dstChainID: evmchain.IDHolesky,
				deposits:   []types.AddrAmt{{Amount: bi.Ether(0.3)}},
				calls:      []types.Call{{Value: bi.Ether(0.1)}},
				expenses: []types.Expense{
					{Amount: bi.Ether(0.1)},
					{Amount: bi.Ether(0.01), Token: common.HexToAddress("0x8d09a4502cc8cf1547ad300e066060d043f6982d")},
				},
			},
		},
		{
			name:   "insufficient deposit",
			reason: types.RejectInsufficientDeposit,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDBaseSepolia,
				dstChainID: evmchain.IDHolesky,
				// does not include fee
				deposits: []types.AddrAmt{{Amount: ether(1)}},
				calls:    []types.Call{{Value: ether(1)}},
				expenses: []types.Expense{{Amount: ether(1)}},
			},
			testdata: true,
		},
		{
			name: "sufficient deposit",
			order: testOrder{
				srcChainID: evmchain.IDBaseSepolia,
				dstChainID: evmchain.IDHolesky,
				// includes fee
				deposits: []types.AddrAmt{{Amount: depositFor(t, ether(1))}},
				calls:    []types.Call{{Value: ether(1)}},
				expenses: []types.Expense{{Amount: ether(1)}},
			},
			mock: func(clients MockClients) {
				mockNativeBalance(t, clients.Client(t, evmchain.IDHolesky), solver, bi.Add(bi.Ether(1), minSafeETH)) // add min safe eth
			},
		},
		{
			name:   "spend below min safe eth",
			reason: types.RejectInsufficientInventory,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDBaseSepolia,
				dstChainID: evmchain.IDHolesky,
				// includes fee
				deposits: []types.AddrAmt{{Amount: depositFor(t, ether(1))}},
				calls:    []types.Call{{Value: ether(1)}},
				expenses: []types.Expense{{Amount: ether(1)}},
			},
			mock: func(clients MockClients) {
				mockNativeBalance(t, clients.Client(t, evmchain.IDHolesky), solver, bi.Ether(1)) // do not add min safe eth
			},
		},
		{
			name: "more than sufficient deposit",
			order: testOrder{
				srcChainID: evmchain.IDBaseSepolia,
				dstChainID: evmchain.IDHolesky,
				deposits: []types.AddrAmt{{
					Amount: bi.Add(
						depositFor(t, ether(1)), // required deposit
						gwei(1),                 // a little more
					),
				}},
				calls:    []types.Call{{Value: ether(1)}},
				expenses: []types.Expense{{Amount: ether(1)}},
			},
			mock: func(clients MockClients) {
				mockNativeBalance(t, clients.Client(t, evmchain.IDHolesky), solver, ether(2))
			},
		},
		{
			name:   "expense over max",
			reason: types.RejectExpenseOverMax,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDBaseSepolia,
				dstChainID: evmchain.IDHolesky,
				deposits:   []types.AddrAmt{{Amount: ether(3000)}},
				calls:      []types.Call{{Value: ether(2000)}},
				expenses:   []types.Expense{{Amount: ether(2000)}},
			},
		},
		{
			name:   "expense under min",
			reason: types.RejectExpenseUnderMin,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDBaseSepolia,
				dstChainID: evmchain.IDHolesky,
				deposits:   []types.AddrAmt{{Amount: bi.N(2)}},
				calls:      []types.Call{{Value: bi.N(1)}},
				expenses:   []types.Expense{{Amount: bi.N(1)}},
			},
		},
		{
			name:         "disallowed call",
			reason:       types.RejectCallNotAllowed,
			reject:       true,
			disallowCall: true,
			// rest does not matter
			order: testOrder{
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDBaseSepolia,
				deposits:   []types.AddrAmt{{Amount: ether(2)}},
				calls:      []types.Call{{Value: ether(1)}},
				expenses:   []types.Expense{{Amount: ether(1)}},
			},
			mock: func(clients MockClients) {
				mockNativeBalance(t, clients.Client(t, evmchain.IDBaseSepolia), solver, ether(2))
			},
		},
		{
			name:   "ETH covers STETH",
			reason: types.RejectNone,
			reject: false,
			order: testOrder{
				srcChainID: evmchain.IDBaseSepolia,
				dstChainID: evmchain.IDHolesky,
				deposits:   []types.AddrAmt{{Amount: depositFor(t, bi.Ether(0.1))}},
				calls:      []types.Call{{Target: common.HexToAddress("0x01"), Data: dummyCallData}}, // does not matter
				expenses:   []types.Expense{{Amount: bi.Ether(0.1), Token: holeskySTETH}},
			},
			mock: func(clients MockClients) {
				mockERC20Balance(t, clients.Client(t, evmchain.IDHolesky), holeskySTETH, ether(1))
				mockERC20Allowance(t, clients.Client(t, evmchain.IDHolesky), holeskySTETH)
			},
		},
		{
			// note USDC has 6 decimals
			name:   "USDC sufficient deposit",
			reason: types.RejectNone,
			reject: false,
			order: testOrder{
				srcChainID: evmchain.IDBaseSepolia,
				dstChainID: evmchain.IDArbSepolia,
				deposits:   []types.AddrAmt{{Amount: depositFor(t, bi.Dec6(1)), Token: baseSepoliaUSDC}},
				calls:      []types.Call{{Target: common.HexToAddress("0x01"), Data: dummyCallData}}, // does not matter
				expenses:   []types.Expense{{Amount: bi.Dec6(1), Token: arbSepoliaUSDC}},
			},
			mock: func(clients MockClients) {
				mockERC20Balance(t, clients.Client(t, evmchain.IDArbSepolia), arbSepoliaUSDC, bi.Dec6(1))
				mockERC20Allowance(t, clients.Client(t, evmchain.IDArbSepolia), arbSepoliaUSDC)
			},
		},
		{
			// note USDC has 6 decimals
			name:   "USDC insufficient deposit",
			reason: types.RejectInsufficientDeposit,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDBaseSepolia,
				dstChainID: evmchain.IDArbSepolia,
				deposits:   []types.AddrAmt{{Amount: bi.Dec6(1), Token: baseSepoliaUSDC}},
				calls:      []types.Call{{Target: common.HexToAddress("0x01"), Data: dummyCallData}}, // does not matter
				expenses:   []types.Expense{{Amount: bi.Dec6(1), Token: arbSepoliaUSDC}},
			},
			mock: func(clients MockClients) {
				mockERC20Balance(t, clients.Client(t, evmchain.IDArbSepolia), arbSepoliaUSDC, bi.Dec6(1))
				mockERC20Allowance(t, clients.Client(t, evmchain.IDArbSepolia), arbSepoliaUSDC)
			},
		},
		{
			name:   "USDC expense over max", // note us of ether(1), not bi.Dec6(1)
			reason: types.RejectExpenseOverMax,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDBaseSepolia,
				dstChainID: evmchain.IDArbSepolia,
				deposits:   []types.AddrAmt{{Amount: depositFor(t, ether(1)), Token: baseSepoliaUSDC}},
				calls:      []types.Call{{Target: common.HexToAddress("0x01"), Data: dummyCallData}}, // does not matter
				expenses:   []types.Expense{{Amount: ether(1), Token: arbSepoliaUSDC}},
			},
			mock: func(clients MockClients) {
				mockERC20Balance(t, clients.Client(t, evmchain.IDArbSepolia), arbSepoliaUSDC, ether(1))
				mockERC20Allowance(t, clients.Client(t, evmchain.IDArbSepolia), arbSepoliaUSDC)
			},
		},
	}
}

// testBackends returns test backends / clients required for test cases above.
func testBackends(t *testing.T) (ethbackend.Backends, MockClients) {
	t.Helper()

	return makeMockBackends(t,
		// mock omega chains for tests
		evmchain.IDOmniOmega,
		evmchain.IDHolesky,
		evmchain.IDBaseSepolia,
		evmchain.IDArbSepolia,

		// add one mainnet chain, to make sure testnet ETH cannot be used for mainnet ETH
		evmchain.IDOptimism,
	)
}

// mockDidFill mocks an Outbox.didFill(...) call.
func mockDidFill(t *testing.T, client *mock.MockClient, outbox common.Address, didFill bool) {
	t.Helper()

	ctx := gomock.Any()
	msg := newCallMatcher("Outbox.didFill", outbox, outboxABI.Methods["didFill"].ID)

	client.EXPECT().CallContract(ctx, msg, nil).Return(abiEncodeBool(t, didFill), nil).AnyTimes()
}

// mockFill mocks an Outbox.fill(...) call.
func mockFill(t *testing.T, client *mock.MockClient, outbox common.Address, shouldErr bool) {
	t.Helper()

	var err error
	if shouldErr {
		err = errors.New("mock execution reverted")
	}

	// no return data
	returnData := []byte{}

	ctx := gomock.Any()
	msg := newCallMatcher("Outbox.fill", outbox, outboxABI.Methods["fill"].ID)

	client.EXPECT().CallContract(ctx, msg, nil).Return(returnData, err).AnyTimes()
}

// mockFillFee mocks an Outbox.fillFee(...) call.
func mockFillFee(t *testing.T, client *mock.MockClient, outbox common.Address) {
	t.Helper()

	// always return a fee of 1 gwei
	fee := bi.Gwei(1)

	ctx := gomock.Any()
	msg := newCallMatcher("Outbox.fillFee", outbox, outboxABI.Methods["fillFee"].ID)

	client.EXPECT().CallContract(ctx, msg, nil).Return(abiEncodeBig(t, fee), nil).AnyTimes()
}

// mockNativeBalance mocks a Client.BalanceAt(...) call.
func mockNativeBalance(t *testing.T, client *mock.MockClient, address common.Address, balance *big.Int) {
	t.Helper()

	ctx := gomock.Any()

	client.EXPECT().BalanceAt(ctx, address, nil).Return(balance, nil).AnyTimes()
}

// mockERC20Balance mocks an ERC20.balanceOf(...)
func mockERC20Balance(t *testing.T, client *mock.MockClient, token common.Address, balance *big.Int) {
	t.Helper()

	ctx := gomock.Any()
	msg := newCallMatcher("ERC20.balanceOf", token, erc20ABI.Methods["balanceOf"].ID)

	client.EXPECT().CallContract(ctx, msg, nil).Return(abiEncodeBig(t, balance), nil).AnyTimes()
}

// mockERC20Allowance mocks an ERC20.allowance(...)
func mockERC20Allowance(t *testing.T, client *mock.MockClient, token common.Address) {
	t.Helper()

	ctx := gomock.Any()

	// match allowance, return max uint256, so no approval tx required
	msg := CallMatcher{
		To:       token,
		Name:     "ERC20.allowance",
		Selector: erc20ABI.Methods["allowance"].ID,
	}

	client.EXPECT().CallContract(ctx, msg, nil).Return(abiEncodeBig(t, umath.MaxUint256), nil).AnyTimes()
}

func makeMockBackends(t *testing.T, chainIDs ...uint64) (ethbackend.Backends, MockClients) {
	t.Helper()

	clients := make(map[uint64]*mock.MockClient)
	controllers := make(map[uint64]*gomock.Controller)
	backends := make(map[uint64]*ethbackend.Backend)

	for _, chainID := range chainIDs {
		chain, ok := evmchain.MetadataByID(chainID)
		require.True(t, ok)

		ctrl := gomock.NewController(t)
		ethcl := mock.NewMockClient(ctrl)
		clients[chainID] = ethcl

		backend, err := ethbackend.NewDevBackend(chain.Name, chain.ChainID, chain.BlockPeriod, ethcl)
		require.NoError(t, err)

		backends[chainID] = backend
	}

	return ethbackend.BackendsFrom(backends), MockClients{clients, controllers}
}

type MockClients struct {
	clients     map[uint64]*mock.MockClient
	controllers map[uint64]*gomock.Controller
}

// Client returns a mock client for the given chainID.
func (m MockClients) Client(t *testing.T, chainID uint64) *mock.MockClient {
	t.Helper()

	c, ok := m.clients[chainID]
	require.True(t, ok, "client for chainID %d not found", chainID)

	return c
}

// Finish finishes all mock controllers.
func (m MockClients) Finish(t *testing.T) {
	t.Helper()

	for _, ctrl := range m.controllers {
		ctrl.Finish()
	}
}

// CallMatcher is a gomock.Matcher for ethereum.CallMsg.
type CallMatcher struct {
	Name string

	To common.Address

	// if omitted, matches any selector
	Selector []byte
}

var _ gomock.Matcher = CallMatcher{}

func newCallMatcher(name string, to common.Address, selector []byte) gomock.Matcher {
	matcher := CallMatcher{
		Name:     name,
		To:       to,
		Selector: selector,
	}

	return gomock.GotFormatterAdapter(CallGotFormatter{}, matcher)
}

func (m CallMatcher) Matches(x any) bool {
	msg, ok := x.(ethereum.CallMsg)
	if !ok {
		return false
	}

	if msg.To == nil || *msg.To != m.To {
		return false
	}

	if len(m.Selector) > 0 && !bytes.Equal(m.Selector, msg.Data[:4]) {
		return false
	}

	return true
}

func (m CallMatcher) String() string {
	return "matches call to " + m.Name + " at " + m.To.Hex()
}

// CallGotFormatter formats the got value of a CallMatcher.
type CallGotFormatter struct{}

var _ gomock.GotFormatter = CallGotFormatter{}

func (CallGotFormatter) Got(got any) string {
	msg, ok := got.(ethereum.CallMsg)
	if !ok {
		return "unknown"
	}

	selector := msg.Data[:4]
	abis := map[string]*abi.ABI{"Inbox": inboxABI, "Oubox": outboxABI, "ERC20": erc20ABI}

	friendly := hexutil.Encode(selector)
	for name, abi := range abis {
		for _, method := range abi.Methods {
			if bytes.Equal(selector, method.ID) {
				friendly = name + "." + method.Name
				break
			}
		}
	}

	return "call " + friendly + " at " + msg.To.Hex()
}

func abiEncodeBig(t *testing.T, n *big.Int) []byte {
	t.Helper()

	abiT, err := abi.NewType("uint256", "", nil)
	require.NoError(t, err)

	data, err := abi.Arguments{{Type: abiT}}.Pack(n)
	require.NoError(t, err)

	return data
}

func abiEncodeBool(t *testing.T, b bool) []byte {
	t.Helper()

	abiT, err := abi.NewType("bool", "", nil)
	require.NoError(t, err)

	data, err := abi.Arguments{{Type: abiT}}.Pack(b)
	require.NoError(t, err)

	return data
}

func ether(x int64) *big.Int {
	return bi.Ether(x)
}

func gwei(x int64) *big.Int {
	return bi.Gwei(x)
}

// depositFor is equivalent to legacy DepositFor(...) function.
func depositFor(t *testing.T, expense *big.Int) *big.Int {
	t.Helper()

	// feePrice is unary price function that is 1-to-1 WITH fees.
	feePrice, err := wrapPriceHandlerFunc(unaryPrice)(t.Context(),
		types.PriceRequest{
			SourceChainID:      evmchain.IDMockL2,
			DestinationChainID: evmchain.IDMockL1,
		})
	require.NoError(t, err)

	return feePrice.ToDeposit(expense)
}

// expenseFor is equivalent to legacy ExpenseFor(...) function.
func expenseFor(t *testing.T, deposit *big.Int) *big.Int {
	t.Helper()

	// feePrice is unary price function that is 1-to-1 with fees.
	feePrice, err := wrapPriceHandlerFunc(unaryPrice)(t.Context(),
		types.PriceRequest{
			SourceChainID:      evmchain.IDMockL2,
			DestinationChainID: evmchain.IDMockL1,
		})
	require.NoError(t, err)

	return feePrice.ToExpense(deposit)
}
