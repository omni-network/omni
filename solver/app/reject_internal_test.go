package app

import (
	"context"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/ethclient/mock"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCmdAddrs(t *testing.T) {
	t.Parallel()

	// zero addr
	var bz [32]byte
	addr := toEthAddr(bz)
	require.True(t, cmpAddrs(addr, bz))

	// within 20 bytes
	_, err := rand.Read(bz[12:])
	require.NoError(t, err)
	addr = toEthAddr(bz)
	require.True(t, cmpAddrs(addr, bz))

	// not within 20 bytes
	_, err = rand.Read(bz[:32])
	bz[31] = 0x01 // just make sure it's not all zeros
	require.NoError(t, err)
	addr = toEthAddr(bz)
	require.False(t, cmpAddrs(addr, bz))
}

//nolint:tparallel // subtests use same mock controller
func TestShouldReject(t *testing.T) {
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

	shouldReject := newShouldRejector(backends, solver, targetName, chainName)

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

	makeOutputs := func(payments ...Payment) []bindings.IERC7683Output {
		var outputs []bindings.IERC7683Output

		for _, p := range payments {
			outputs = append(outputs, bindings.IERC7683Output{
				Amount:  p.Amount,
				Token:   toBz32(p.Token.Address),
				ChainId: new(big.Int).SetUint64(p.Token.ChainID),
			})
		}

		return outputs
	}

	tests := []struct {
		name   string
		reason rejectReason
		reject bool
		mock   func()
		order  Order
	}{
		{
			name:   "insufficient native balance",
			reason: rejectInsufficientInventory,
			reject: true,
			mock:   mockNativeBalance(evmchain.IDOmniOmega, big.NewInt(0)),
			order: Order{
				// request 1 native OMNI for 1 erc20 OMNI on omega
				ID:                 [32]byte{0x01},
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDOmniOmega,
				MinReceived:        makeOutputs(Payment{Amount: big.NewInt(1), Token: omniERC20(netconf.Omega)}),
				MaxSpent:           makeOutputs(Payment{Amount: big.NewInt(1), Token: nativeOMNI(evmchain.IDOmniOmega)}),
			},
		},
		{
			name:   "sufficient native balance",
			reason: rejectNone,
			reject: false,
			mock:   mockNativeBalance(evmchain.IDOmniOmega, big.NewInt(1)),
			order: Order{
				// request 1 native OMNI for 1 erc20 OMNI on omega
				ID:                 [32]byte{0x01},
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDOmniOmega,
				MinReceived:        makeOutputs(Payment{Amount: big.NewInt(1), Token: omniERC20(netconf.Omega)}),
				MaxSpent:           makeOutputs(Payment{Amount: big.NewInt(1), Token: nativeOMNI(evmchain.IDOmniOmega)}),
			},
		},
		{
			name:   "insufficient ERC20 balance",
			reason: rejectInsufficientInventory,
			reject: true,
			mock:   mockERC20Balance(evmchain.IDHolesky, big.NewInt(0)),
			order: Order{
				// request 1 erc20 OMNI for 1 native OMNI on omega
				ID:                 [32]byte{0x01},
				SourceChainID:      evmchain.IDOmniOmega,
				DestinationChainID: evmchain.IDHolesky,
				MinReceived:        makeOutputs(Payment{Amount: big.NewInt(1), Token: nativeOMNI(evmchain.IDOmniOmega)}),
				MaxSpent:           makeOutputs(Payment{Amount: big.NewInt(1), Token: omniERC20(netconf.Omega)}),
			},
		},
		{
			name:   "sufficient ERC20 balance",
			reason: rejectNone,
			reject: false,
			mock:   mockERC20Balance(evmchain.IDHolesky, big.NewInt(1)),
			order: Order{
				// request 1 erc20 OMNI for 1 native OMNI on omega
				ID:                 [32]byte{0x01},
				SourceChainID:      evmchain.IDOmniOmega,
				DestinationChainID: evmchain.IDHolesky,
				MinReceived:        makeOutputs(Payment{Amount: big.NewInt(1), Token: nativeOMNI(evmchain.IDOmniOmega)}),
				MaxSpent:           makeOutputs(Payment{Amount: big.NewInt(1), Token: omniERC20(netconf.Omega)}),
			},
		},
		{
			name:   "unsupported expense token",
			reason: rejectUnsupportedExpense,
			reject: true,
			order: Order{
				// request unsupported erc20 for 1 native OMNI on omega
				ID:                 [32]byte{0x01},
				SourceChainID:      evmchain.IDOmniOmega,
				DestinationChainID: evmchain.IDHolesky,
				MinReceived:        makeOutputs(Payment{Amount: big.NewInt(1), Token: nativeOMNI(evmchain.IDOmniOmega)}),
				MaxSpent: makeOutputs(Payment{
					Amount: big.NewInt(1),
					Token: Token{
						ChainID: evmchain.IDHolesky,
						Address: common.HexToAddress("0x01"), // unsupported token
					},
				}),
			},
		},
		{
			name:   "unsupported dest chain",
			reason: rejectUnsupportedDestChain,
			reject: true,
			order: Order{
				ID:                 [32]byte{0x01},
				SourceChainID:      evmchain.IDOmniOmega,
				DestinationChainID: 1234567, // unsupported chain
			},
		},
		{
			name:   "invalid deposit (token mismatch)",
			reason: rejectInvalidDeposit,
			reject: true,
			order: Order{
				// deposit native ETH for native OMNi
				ID:                 [32]byte{0x01},
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDOmniOmega,
				MinReceived:        makeOutputs(Payment{Amount: big.NewInt(1), Token: nativeETH(evmchain.IDHolesky)}),
				MaxSpent:           makeOutputs(Payment{Amount: big.NewInt(1), Token: nativeOMNI(evmchain.IDOmniOmega)}),
			},
		},
		{
			name:   "invalid deposit (multiple tokens)",
			reason: rejectInvalidDeposit,
			reject: true,
			order: Order{
				// multiple deposits are not supported
				ID:                 [32]byte{0x01},
				SourceChainID:      evmchain.IDHolesky,
				DestinationChainID: evmchain.IDOmniOmega,
				MinReceived: makeOutputs(
					Payment{Amount: big.NewInt(1), Token: omniERC20(netconf.Omega)},
					Payment{Amount: big.NewInt(1), Token: nativeETH(evmchain.IDHolesky)},
				),
				MaxSpent: makeOutputs(Payment{Amount: big.NewInt(1), Token: nativeOMNI(evmchain.IDOmniOmega)}),
			},
		},
		{
			name:   "invalid deposit (mismatch chain class)",
			reason: rejectInvalidDeposit,
			reject: true,
			order: Order{
				// deposit native testnet ETH for mainnet ETH
				ID:                 [32]byte{0x01},
				SourceChainID:      evmchain.IDHolesky,  // testnet chain
				DestinationChainID: evmchain.IDOptimism, // mainnet chain
				MinReceived:        makeOutputs(Payment{Amount: big.NewInt(1), Token: nativeETH(evmchain.IDHolesky)}),
				MaxSpent:           makeOutputs(Payment{Amount: big.NewInt(1), Token: nativeETH(evmchain.IDOptimism)}),
			},
		},
		{
			name:   "invalid expense",
			reason: rejectInvalidExpense,
			reject: true,
			order: Order{
				// multiple expenses are not supported
				ID:                 [32]byte{0x01},
				SourceChainID:      evmchain.IDBaseSepolia,
				DestinationChainID: evmchain.IDHolesky,
				MinReceived:        makeOutputs(Payment{Amount: big.NewInt(1), Token: nativeETH(evmchain.IDBaseSepolia)}),
				MaxSpent: makeOutputs(
					Payment{Amount: big.NewInt(1), Token: nativeETH(evmchain.IDHolesky)},
					Payment{Amount: big.NewInt(1), Token: omniERC20(netconf.Omega)},
				),
			},
		},
		{
			name:   "insufficient deposit",
			reason: rejectInsufficientDeposit,
			reject: true,
			order: Order{
				// request 2 holesky ETH for 1 base sepolia ETH
				ID:                 [32]byte{0x01},
				SourceChainID:      evmchain.IDBaseSepolia,
				DestinationChainID: evmchain.IDHolesky,
				MinReceived:        makeOutputs(Payment{Amount: big.NewInt(1), Token: nativeETH(evmchain.IDBaseSepolia)}),
				MaxSpent:           makeOutputs(Payment{Amount: big.NewInt(2), Token: nativeETH(evmchain.IDHolesky)}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
			reason, reject, err := shouldReject(ctx, tt.order.SourceChainID, tt.order)
			require.NoError(t, err)
			require.Equal(t, tt.reject, reject, "expected reject reason %s, got %s", tt.reason, reason)
			require.Equal(t, tt.reason, reason, "expected reject reason %s, got %s", tt.reason, reason)
		})
	}
}

func makeMockBackends(t *testing.T, chainIDs ...uint64) (ethbackend.Backends, map[uint64]*mock.MockClient) {
	t.Helper()

	clients := make(map[uint64]*mock.MockClient)
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

	return ethbackend.BackendsFrom(backends), clients
}

func targetName(o Order) string {
	fill, err := o.ParsedFillOriginData()
	if err != nil {
		return unknown
	}

	// use last call target
	return fill.Calls[len(fill.Calls)-1].Target.Hex()
}

func chainName(chainID uint64) string {
	metadata, ok := evmchain.MetadataByID(chainID)
	if !ok {
		return unknown
	}

	return metadata.Name
}

func abiEncodeBig(t *testing.T, n *big.Int) []byte {
	t.Helper()

	abiT, err := abi.NewType("uint256", "", nil)
	require.NoError(t, err)

	data, err := abi.Arguments{{Type: abiT}}.Pack(n)
	require.NoError(t, err)

	return data
}
