package app

import (
	"crypto/rand"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestCmdAddrs(t *testing.T) {
	t.Parallel()

	// zero addr
	var bz [32]byte
	addr, err := toEthAddr(bz)
	require.NoError(t, err)
	require.Equal(t, common.Address{}, addr)

	// within 20 bytes
	_, err = rand.Read(bz[12:])
	require.NoError(t, err)
	_, err = toEthAddr(bz)
	require.NoError(t, err)

	// not within 20 bytes
	_, err = rand.Read(bz[:32])
	bz[31] = 0x01 // just make sure it's not all zeros
	require.NoError(t, err)
	_, err = toEthAddr(bz)
	require.Error(t, err)
}

//nolint:tparallel,paralleltest // subtests use same mock controller
func TestShouldReject(t *testing.T) {
	t.Parallel()

	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)

	// outbox addr only matters for mocks, using devnet
	addrs, err := contracts.GetAddresses(t.Context(), netconf.Devnet)
	require.NoError(t, err)
	outbox := addrs.SolverNetOutbox

	priceFunc := unaryPrice

	for _, tt := range rejectTestCases(t, solver, outbox) {
		// TODO(zodomo): Remove this once network upgrade is complete
		if tt.order.SourceChainID == evmchain.IDOmniMainnet {
			tt.reject = true
			tt.reason = types.RejectUnsupportedSrcChain
		}

		// TODO(zodomo): Remove this once network upgrade is complete
		if tt.order.pendingData.DestinationChainID == evmchain.IDOmniMainnet {
			tt.reject = true
			tt.reason = types.RejectUnsupportedDestChain
		}

		t.Run(tt.name, func(t *testing.T) {
			backends, clients := testBackends(t)

			callAllower := func(_ uint64, _ common.Address, _ []byte) bool { return !tt.disallowCall }
			shouldReject := newShouldRejector(backends, callAllower, priceFunc, solver, outbox)

			if tt.mock != nil {
				tt.mock(clients)

				destClient := clients.Client(t, tt.order.pendingData.DestinationChainID)
				mockFill(t, destClient, outbox, tt.fillReverts)
				mockFillFee(t, destClient, outbox)

				// didFill check should be made before shouldReject
				mockDidFill(t, destClient, outbox, false)
			}

			reason, reject, err := shouldReject(t.Context(), tt.order)
			if tt.reject {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.reason, reason, "expected reject reason %s, got %s", tt.reason, reason)
			require.Equal(t, tt.reject, reject, "expected reject %s, got %s", tt.reject, reject)

			clients.Finish(t)
		})
	}
}
