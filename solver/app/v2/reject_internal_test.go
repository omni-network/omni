package appv2

import (
	"context"
	"crypto/rand"
	"math/big"
	"testing"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/cast"
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
	addr, err := cast.EthAddress(bz[:20])
	require.NoError(t, err)
	require.True(t, cmpAddrs(addr, bz))

	// within 20 bytes
	_, err = rand.Read(bz[:20])
	require.NoError(t, err)
	addr, err = cast.EthAddress(bz[:20])
	require.NoError(t, err)
	require.True(t, cmpAddrs(addr, bz))

	// not within 20 bytes
	_, err = rand.Read(bz[:32])
	bz[31] = 0x01 // just make sure it's not all zeros
	require.NoError(t, err)
	addr, err = cast.EthAddress(bz[:20])
	require.NoError(t, err)
	require.False(t, cmpAddrs(addr, bz))
}

//nolint:tparallel // subtests use same mock controller
func TestShouldReject(t *testing.T) {
	t.Parallel()

	// static setup
	ctx := context.Background()
	srcChainID := evmchain.IDBaseSepolia
	destChainID := evmchain.IDHolesky
	omniERC20Addr := omniERC20(netconf.Omega).Address
	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)
	targetName := func(Order) string { return "target" }
	chainName := func(uint64) string { return "chain" }

	// mock backend, to manipulate balances
	ctrl := gomock.NewController(t)
	mockEthCl := mock.NewMockClient(ctrl)
	mockBackend, err := ethbackend.NewDevBackend("mock", destChainID, time.Minute, mockEthCl)
	require.NoError(t, err)
	mockBackends := ethbackend.BackendsFrom(map[uint64]*ethbackend.Backend{destChainID: mockBackend})

	shouldReject := newShouldRejector(mockBackends, solver, targetName, chainName)

	makeOrder := func(maxSpent ...bindings.IERC7683Output) Order {
		return Order{
			ID:                 [32]byte{0x01},
			SourceChainID:      srcChainID,
			DestinationChainID: destChainID,
			MaxSpent:           maxSpent,
		}
	}

	mockNativeBalance := func(b *big.Int) func() {
		return func() {
			mockEthCl.EXPECT().BalanceAt(ctx, solver, nil).Return(b, nil)
		}
	}

	mockERC20Balance := func(b *big.Int) func() {
		return func() {
			// we don't go through the trouble of matching eth msg param to IERC20(addr).balanceOf call
			ctx := gomock.Any()
			msg := gomock.Any()
			mockEthCl.EXPECT().CallContract(ctx, msg, nil).Return(abiEncodeBig(t, b), nil)
		}
	}

	nativeOutput := func(amount *big.Int) bindings.IERC7683Output {
		return bindings.IERC7683Output{
			Amount:  amount,
			Token:   [32]byte{}, // native
			ChainId: new(big.Int).SetUint64(destChainID),
		}
	}

	erc20Output := func(amount *big.Int) bindings.IERC7683Output {
		// we need to use omniERC20Addr, because that's the only ERC20 supported in tokens.go
		return bindings.IERC7683Output{
			Amount:  amount,
			Token:   toBz32(omniERC20Addr),
			ChainId: new(big.Int).SetUint64(destChainID),
		}
	}

	tests := []struct {
		name   string
		reason rejectReason
		reject bool
		mock   func()
		order  func() Order
	}{
		{
			name:   "insufficient native balance",
			reason: rejectInsufficientInventory,
			reject: true,
			mock:   mockNativeBalance(big.NewInt(0)),
			order:  func() Order { return makeOrder(nativeOutput(big.NewInt(1))) },
		},
		{
			name:   "sufficient native balance",
			reason: rejectNone,
			reject: false,
			mock:   mockNativeBalance(big.NewInt(1)),
			order:  func() Order { return makeOrder(nativeOutput(big.NewInt(1))) },
		},
		{
			name:   "insufficient ERC20 balance",
			reason: rejectInsufficientInventory,
			reject: true,
			mock:   mockERC20Balance(big.NewInt(0)),
			order:  func() Order { return makeOrder(erc20Output(big.NewInt(1))) },
		},
		{
			name:   "sufficient ERC20 balance",
			reason: rejectNone,
			reject: false,
			mock:   mockERC20Balance(big.NewInt(1)),
			order:  func() Order { return makeOrder(erc20Output(big.NewInt(1))) },
		},
		{
			name:   "unsupported token",
			reason: rejectUnsupportedToken,
			reject: true,
			mock:   func() {},
			order: func() Order {
				unknownTkn := [32]byte{0x01}
				return makeOrder(bindings.IERC7683Output{Token: unknownTkn, ChainId: new(big.Int).SetUint64(destChainID)})
			},
		},
		{
			name:   "unsupported dest chain",
			reason: rejectUnsupportedDestChain,
			reject: true,
			mock:   func() {},
			order: func() Order {
				badDestChainID := destChainID + 1
				order := makeOrder(bindings.IERC7683Output{Token: [32]byte{}, ChainId: new(big.Int).SetUint64(0x01)})
				order.DestinationChainID = badDestChainID

				return order
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			reason, reject, err := shouldReject(ctx, srcChainID, tt.order())
			require.NoError(t, err)
			require.Equal(t, tt.reject, reject)
			require.Equal(t, tt.reason, reason)
		})
	}
}

func abiEncodeBig(t *testing.T, n *big.Int) []byte {
	t.Helper()

	abiT, err := abi.NewType("uint256", "", nil)
	require.NoError(t, err)

	data, err := abi.Arguments{{Type: abiT}}.Pack(n)
	require.NoError(t, err)

	return data
}

func toBz32(addr common.Address) [32]byte {
	var bz [32]byte
	copy(bz[:], addr.Bytes())

	return bz
}
