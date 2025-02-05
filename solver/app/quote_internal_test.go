package app

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/ethclient/mock"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// TODO: merge TestQuote & TestShouldReject test cases, as reject cases should match
//
//nolint:tparallel // subtests use same mock controller
func TestQuote(t *testing.T) {
	t.Parallel()

	// static setup
	ctx := context.Background()
	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)

	// mock backends, to manipulate balances
	backends, clients := makeMockBackends(t,
		// mock omega chains for tests
		evmchain.IDOmniOmega,
		evmchain.IDHolesky,
		evmchain.IDBaseSepolia,

		// add one mainnet chain, to make sure testnet ETH cannot be used for mainnet ETH
		evmchain.IDOptimism,
	)

	client := func(chainID uint64) *mock.MockClient {
		c, ok := clients[chainID]
		require.True(t, ok, "client for chainID %d not found", chainID)

		return c
	}

	mockNativeBalance := func(chainID uint64, balance *big.Int) func() {
		return func() {
			client(chainID).EXPECT().BalanceAt(ctx, solver, nil).Return(balance, nil)
		}
	}

	mockERC20Balance := func(chainID uint64, balance *big.Int) func() {
		return func() {
			// TODO: match eth msg param to IERC20(addr).balanceOf call
			ctx := gomock.Any()
			msg := gomock.Any()
			client(chainID).EXPECT().CallContract(ctx, msg, nil).Return(abiEncodeBig(t, balance), nil)
		}
	}

	tests := []struct {
		name string
		mock func()
		req  QuoteRequest
		res  QuoteResponse
	}{
		{
			name: "insufficient native balance",
			mock: mockNativeBalance(evmchain.IDOmniOmega, big.NewInt(0)),
			req: QuoteRequest{
				// quote 1 native OMNI for erc20 OMNI on omega
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDOmniOmega,
				Expenses:           []Expense{{Amount: big.NewInt(1), Token: common.Address{}}}, // native
				DepositToken:       omniERC20(netconf.Omega).Address,
			},
			res: QuoteResponse{
				// solver does not have enough native balance
				Rejected:     true,
				RejectReason: rejectInsufficientInventory.String(),
			},
		},
		{
			name: "sufficient native balance",
			mock: mockNativeBalance(evmchain.IDOmniOmega, big.NewInt(1)),
			req: QuoteRequest{
				// quote 1 native OMNI for erc20 OMNI on omega
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDOmniOmega,
				Expenses:           []Expense{{Amount: big.NewInt(1), Token: common.Address{}}}, // native
				DepositToken:       omniERC20(netconf.Omega).Address,
			},
			res: QuoteResponse{
				// 1 erc20 OMNI required
				Deposit: &Deposit{
					Amount: big.NewInt(1),
					Token:  omniERC20(netconf.Omega).Address,
				},
			},
		},
		{
			name: "insufficient ERC20 balance",
			mock: mockERC20Balance(evmchain.IDHolesky, big.NewInt(0)),
			req: QuoteRequest{
				// request 1 erc20 OMNI for 1 native OMNI on omega
				SourceChainID:      evmchain.IDOmniOmega,
				DestinationChainID: evmchain.IDHolesky,
				Expenses:           []Expense{{Amount: big.NewInt(1), Token: omniERC20(netconf.Omega).Address}},
				DepositToken:       common.Address{}, // native
			},
			res: QuoteResponse{
				// solver does not have enough erc20 balance
				Rejected:     true,
				RejectReason: rejectInsufficientInventory.String(),
			},
		},
		{
			name: "sufficient ERC20 balance",
			mock: mockERC20Balance(evmchain.IDHolesky, big.NewInt(1)),
			req: QuoteRequest{
				// request 1 erc20 OMNI for 1 native OMNI on omega
				SourceChainID:      evmchain.IDOmniOmega,
				DestinationChainID: evmchain.IDHolesky,
				Expenses:           []Expense{{Amount: big.NewInt(1), Token: omniERC20(netconf.Omega).Address}},
				DepositToken:       common.Address{}, // native
			},
			res: QuoteResponse{
				// 1 native OMNI required
				Deposit: &Deposit{
					Amount: big.NewInt(1),
					Token:  common.Address{},
				},
			},
		},
		{
			name: "unsupported expense token",
			req: QuoteRequest{
				// request unsupported erc20 for  native OMNI on omega
				SourceChainID:      evmchain.IDOmniOmega,
				DestinationChainID: evmchain.IDHolesky,
				Expenses:           []Expense{{Amount: big.NewInt(1), Token: common.HexToAddress("0x01")}}, // unsupported token
				DepositToken:       common.Address{},                                                       // native
			},
			res: QuoteResponse{
				// expense token is not supported
				Rejected:     true,
				RejectReason: rejectUnsupportedExpense.String(),
			},
		},
		{
			name: "unsupported dest chain",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDOmniOmega,
				DestinationChainID: 1234567, // unsupported chain
			},
			res: QuoteResponse{
				// destination chain is not supported
				Rejected:     true,
				RejectReason: rejectUnsupportedDestChain.String(),
			},
		},
		{
			name: "invalid deposit (token mismatch)",
			req: QuoteRequest{
				// deposit native ETH for native OMNi
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDOmniOmega,
				Expenses:           []Expense{{Amount: big.NewInt(1), Token: common.Address{}}}, // native
				DepositToken:       common.Address{},                                            // native
			},

			res: QuoteResponse{
				// deposit token does not match expense token
				Rejected:     true,
				RejectReason: rejectInvalidDeposit.String(),
			},
		},
		{
			name: "invalid deposit (mismatch chain class)",
			req: QuoteRequest{
				// deposit native testnet ETH for mainnet ETH
				SourceChainID:      evmchain.IDHolesky,                                          // testnet chain
				DestinationChainID: evmchain.IDOptimism,                                         // mainnet chain
				Expenses:           []Expense{{Amount: big.NewInt(1), Token: common.Address{}}}, // native
				DepositToken:       common.Address{},                                            // native
			},
			res: QuoteResponse{
				// deposit token does not match expense token
				Rejected:     true,
				RejectReason: rejectInvalidDeposit.String(),
			},
		},
		{
			name: "invalid expenses (multiple expenses)",
			req: QuoteRequest{
				SourceChainID:      evmchain.IDBaseSepolia,
				DestinationChainID: evmchain.IDHolesky,
				Expenses: []Expense{
					{Amount: big.NewInt(1), Token: omniERC20(netconf.Omega).Address},
					{Amount: big.NewInt(1), Token: common.Address{}}, // native[
				},
				DepositToken: common.Address{}, // native
			},
			res: QuoteResponse{
				// multiple expenses are not supported
				Rejected:     true,
				RejectReason: rejectInvalidExpense.String(),
			},
		},
	}

	handler := newQuoteHandler(newQuoter(backends, solver))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			body, err := json.Marshal(tt.req)
			require.NoError(t, err)

			ctx := context.Background()
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "api/v1/quote", bytes.NewBuffer(body))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			require.Equal(t, http.StatusOK, rr.Code)

			var res QuoteResponse
			err = json.NewDecoder(rr.Body).Decode(&res)
			require.NoError(t, err)

			require.Equal(t, tt.res.Rejected, res.Rejected)
			require.Equal(t, tt.res.RejectReason, res.RejectReason)
			require.Equal(t, tt.res.Deposit, res.Deposit)
		})
	}
}
