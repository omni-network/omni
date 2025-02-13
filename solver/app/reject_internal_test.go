package app

import (
	"context"
	"crypto/rand"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/stretchr/testify/require"
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

	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)

	// outbox addr only matters for mocks, using devnet
	addrs, err := contracts.GetAddresses(context.Background(), netconf.Devnet)
	require.NoError(t, err)
	outbox := addrs.SolverNetOutbox

	for _, tt := range rejectTestCases(t, solver, outbox) {
		t.Run(tt.name, func(t *testing.T) {
			backends, clients := testBackends(t)

			shouldReject := newShouldRejector(backends, solver, outbox)

			if tt.mock != nil {
				tt.mock(clients)

				destClient := clients.Client(t, tt.order.DestinationChainID)
				mockFill(t, destClient, outbox, tt.fillReverts)
				mockFillFee(t, destClient, outbox)

				// didFill check should be made before shouldReject
				mockDidFill(t, destClient, outbox, false)
			}

			reason, reject, err := shouldReject(context.Background(), tt.order)

			require.NoError(t, err)
			require.Equal(t, tt.reason, reason, "expected reject reason %s, got %s", tt.reason, reason)
			require.Equal(t, tt.reject, reject, "expected reject %s, got %s", tt.reject, reject)

			clients.Finish(t)
		})
	}
}
