package app

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/ethclient/mock"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	deposits   []Deposit
	expenses   []Expense
}

// orderTestCase is a test case for both shouldReject and quote handlers.
type orderTestCase struct {
	name        string
	reason      rejectReason
	reject      bool
	fillReverts bool
	mock        func(clients MockClients)
	order       testOrder
}

// rejectTestCase is a test case for shouldReject(...)
type rejectTestCase struct {
	name        string
	reason      rejectReason
	reject      bool
	fillReverts bool
	mock        func(clients MockClients)
	order       Order
}

// quoteTestCase is test case for quote handler.
type quoteTestCase struct {
	name string
	mock func(clients MockClients)
	req  QuoteRequest
	res  QuoteResponse
}

func toQuoteTestCase(t *testing.T, tt orderTestCase) quoteTestCase {
	t.Helper()

	require.Len(t, tt.order.deposits, 1)

	var deposit *Deposit
	if !tt.reject {
		deposit = &tt.order.deposits[0]
	}

	var rejectReason string
	if tt.reject {
		rejectReason = tt.reason.String()
	}

	return quoteTestCase{
		name: tt.name,
		mock: tt.mock,
		req: QuoteRequest{
			SourceChainID:      tt.order.srcChainID,
			DestinationChainID: tt.order.dstChainID,
			Expenses:           tt.order.expenses,
			DepositToken:       tt.order.deposits[0].Token,
		},
		res: QuoteResponse{
			Deposit:      deposit,
			Rejected:     tt.reject,
			RejectReason: rejectReason,
		},
	}
}

func toRejectTestCase(t *testing.T, tt orderTestCase, outbox common.Address) rejectTestCase {
	t.Helper()

	maxSpent := make([]bindings.IERC7683Output, len(tt.order.expenses))
	for i, e := range tt.order.expenses {
		maxSpent[i] = bindings.IERC7683Output{
			Amount:    e.Amount,
			Token:     toBz32(e.Token),
			Recipient: toBz32(outbox),
			ChainId:   new(big.Int).SetUint64(tt.order.dstChainID),
		}
	}

	minReceived := make([]bindings.IERC7683Output, len(tt.order.deposits))
	for i, d := range tt.order.deposits {
		minReceived[i] = bindings.IERC7683Output{
			Amount:  d.Amount,
			Token:   toBz32(d.Token),
			ChainId: new(big.Int).SetUint64(tt.order.srcChainID),
		}
	}

	return rejectTestCase{
		name:        tt.name,
		reason:      tt.reason,
		reject:      tt.reject,
		fillReverts: tt.fillReverts,
		mock:        tt.mock,
		order: Order{
			ID:                 [32]byte{0x01},
			SourceChainID:      tt.order.srcChainID,
			DestinationChainID: tt.order.dstChainID,
			DestinationSettler: outbox,
			MaxSpent:           maxSpent,
			MinReceived:        minReceived,
		},
	}
}

func quoteTestCases(t *testing.T, solver common.Address) []quoteTestCase {
	t.Helper()

	var tests []quoteTestCase

	for _, tt := range orderTestCases(t, solver) {
		if len(tt.order.deposits) != 1 {
			// quote requires single deposit token
			continue
		}

		tests = append(tests, toQuoteTestCase(t, tt))
	}

	additional := []quoteTestCase{
		{
			name: "unsupported source chain",
			req: QuoteRequest{
				SourceChainID:      1234567,
				DestinationChainID: evmchain.IDHolesky,
			},
			res: QuoteResponse{
				Rejected:     true,
				RejectReason: rejectUnsupportedSrcChain.String(),
			},
		},
		{
			name: "same chain",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDHolesky,
			},
			res: QuoteResponse{
				Rejected:     true,
				RejectReason: rejectSameChain.String(),
			},
		},
	}

	return append(tests, additional...)
}

func rejectTestCases(t *testing.T, solver, outbox common.Address) []rejectTestCase {
	t.Helper()

	tests := []rejectTestCase{}

	for _, tt := range orderTestCases(t, solver) {
		tests = append(tests, toRejectTestCase(t, tt, outbox))
	}

	// not shared with quote
	additional := []rejectTestCase{
		toRejectTestCase(t, orderTestCase{
			name:   "insufficient deposit",
			reason: rejectInsufficientDeposit,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDBaseSepolia,
				dstChainID: evmchain.IDHolesky,
				deposits:   []Deposit{{Amount: big.NewInt(1)}},
				expenses:   []Expense{{Amount: big.NewInt(2)}},
			},
		}, outbox),
	}

	return append(tests, additional...)
}

func orderTestCases(t *testing.T, solver common.Address) []orderTestCase {
	t.Helper()

	omegaOMNIAddr := omniERC20(netconf.Omega).Address

	return []orderTestCase{
		{
			name:   "insufficient native balance",
			reason: rejectInsufficientInventory,
			reject: true,
			order: testOrder{
				// request 1 native OMNI for 1 erc20 OMNI on omega
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDOmniOmega,
				deposits:   []Deposit{{Amount: big.NewInt(1), Token: omegaOMNIAddr}},
				expenses:   []Expense{{Amount: big.NewInt(1)}},
			},
			mock: func(clients MockClients) {
				mockNativeBalance(t, clients.Client(t, evmchain.IDOmniOmega), solver, big.NewInt(0))
			},
		},
		{
			name:   "sufficient native balance",
			reason: rejectNone,
			reject: false,
			order: testOrder{
				// request 1 native OMNI for 1 erc20 OMNI on omega
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDOmniOmega,
				deposits:   []Deposit{{Amount: big.NewInt(1), Token: omegaOMNIAddr}},
				expenses:   []Expense{{Amount: big.NewInt(1)}},
			},
			mock: func(clients MockClients) {
				mockNativeBalance(t, clients.Client(t, evmchain.IDOmniOmega), solver, big.NewInt(1))
			},
		},
		{
			name:   "insufficient ERC20 balance",
			reason: rejectInsufficientInventory,
			reject: true,
			order: testOrder{
				// request 1 erc20 OMNI for 1 native OMNI on omega
				srcChainID: evmchain.IDOmniOmega,
				dstChainID: evmchain.IDHolesky,
				deposits:   []Deposit{{Amount: big.NewInt(1)}},
				expenses:   []Expense{{Amount: big.NewInt(1), Token: omegaOMNIAddr}},
			},
			mock: func(clients MockClients) {
				mockERC20Balance(t, clients.Client(t, evmchain.IDHolesky), omegaOMNIAddr, big.NewInt(0))
			},
		},
		{
			name:   "sufficient ERC20 balance",
			reason: rejectNone,
			reject: false,
			order: testOrder{
				// request 1 erc20 OMNI for 1 native OMNI on omega
				srcChainID: evmchain.IDOmniOmega,
				dstChainID: evmchain.IDHolesky,
				deposits:   []Deposit{{Amount: big.NewInt(1)}},
				expenses:   []Expense{{Amount: big.NewInt(1), Token: omegaOMNIAddr}},
			},
			mock: func(clients MockClients) {
				mockERC20Balance(t, clients.Client(t, evmchain.IDHolesky), omegaOMNIAddr, big.NewInt(1))
				mockERC20Allowance(t, clients.Client(t, evmchain.IDHolesky), omegaOMNIAddr)
			},
		},
		{
			name:        "fill reverts",
			fillReverts: true,
			reason:      rejectDestCallReverts,
			reject:      true,

			// rest same as above
			order: testOrder{
				// request 1 erc20 OMNI for 1 native OMNI on omega
				srcChainID: evmchain.IDOmniOmega,
				dstChainID: evmchain.IDHolesky,
				deposits:   []Deposit{{Amount: big.NewInt(1)}},
				expenses:   []Expense{{Amount: big.NewInt(1), Token: omegaOMNIAddr}},
			},
			mock: func(clients MockClients) {
				mockERC20Balance(t, clients.Client(t, evmchain.IDHolesky), omegaOMNIAddr, big.NewInt(1))
				mockERC20Allowance(t, clients.Client(t, evmchain.IDHolesky), omegaOMNIAddr)
			},
		},
		{
			name:   "unsupported expense token",
			reason: rejectUnsupportedExpense,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDOmniOmega,
				dstChainID: evmchain.IDHolesky,
				deposits:   []Deposit{{Amount: big.NewInt(1)}},
				expenses:   []Expense{{Amount: big.NewInt(1), Token: common.HexToAddress("0x01")}},
			},
		},
		{
			name:   "unsupported dest chain",
			reason: rejectUnsupportedDestChain,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDOmniOmega,
				dstChainID: 1234567,
			},
		},
		{
			name:   "invalid deposit (token mismatch)",
			reason: rejectInvalidDeposit,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDOmniOmega,
				deposits:   []Deposit{{Amount: big.NewInt(1)}},
				expenses:   []Expense{{Amount: big.NewInt(1)}},
			},
		},
		{
			name:   "invalid deposit (multiple tokens)",
			reason: rejectInvalidDeposit,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDOmniOmega,
				deposits:   []Deposit{{Amount: big.NewInt(1), Token: omegaOMNIAddr}, {Amount: big.NewInt(1)}},
				expenses:   []Expense{{Amount: big.NewInt(1)}},
			},
		},
		{
			name:   "invalid deposit (mismatch chain class)",
			reason: rejectInvalidDeposit,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDHolesky,
				dstChainID: evmchain.IDOptimism,
				deposits:   []Deposit{{Amount: big.NewInt(1)}},
				expenses:   []Expense{{Amount: big.NewInt(1)}},
			},
		},
		{
			name:   "invalid expense (multiple tokens)",
			reason: rejectInvalidExpense,
			reject: true,
			order: testOrder{
				srcChainID: evmchain.IDBaseSepolia,
				dstChainID: evmchain.IDHolesky,
				deposits:   []Deposit{{Amount: big.NewInt(1)}},
				expenses:   []Expense{{Amount: big.NewInt(1)}, {Amount: big.NewInt(1), Token: omegaOMNIAddr}},
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
	fee := big.NewInt(1e9)

	ctx := gomock.Any()
	msg := newCallMatcher("Outbox.fillFee", outbox, outboxABI.Methods["fillFee"].ID)

	client.EXPECT().CallContract(ctx, msg, nil).Return(abiEncodeBig(t, fee), nil).AnyTimes()
}

// mockGetNextID mocks an Inbox.getNextId(...) call.
func mockGetNextID(t *testing.T, client *mock.MockClient, inbox common.Address) {
	t.Helper()

	// next id does not matter, just return 1
	id := big.NewInt(1)

	ctx := gomock.Any()
	msg := newCallMatcher("Inbox.getNextId", inbox, inboxABI.Methods["getNextId"].ID)

	client.EXPECT().CallContract(ctx, msg, nil).Return(abiEncodeBig(t, id), nil).AnyTimes()
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

func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}
