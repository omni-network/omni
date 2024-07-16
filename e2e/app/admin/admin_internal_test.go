package admin

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

type runCall struct {
	calldata []byte
	senders  []common.Address
}

type mockRunner struct {
	calls []runCall
}

var _ runner = &mockRunner{}

func (r *mockRunner) run(ctx context.Context, calldata []byte, senders ...common.Address) (string, error) {
	r.calls = append(r.calls, runCall{calldata, senders})

	return "", nil
}

func TestPausePortalAction(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	r := &mockRunner{}
	s := mockShared()
	c := mockChain()

	_, err := pausePortal(ctx, s, c, r)
	require.NoError(t, err)
	require.Len(t, r.calls, 1)
	require.Equal(t, mustPack(t, adminABI, "pausePortal", s.admin, c.PortalAddress), r.calls[0].calldata)
	require.Equal(t, []common.Address{s.admin}, r.calls[0].senders)
}

func TestUnpausePortalAction(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	r := &mockRunner{}
	s := mockShared()
	c := mockChain()

	_, err := unpausePortal(ctx, s, c, r)
	require.NoError(t, err)
	require.Len(t, r.calls, 1)
	require.Equal(t, mustPack(t, adminABI, "unpausePortal", s.admin, c.PortalAddress), r.calls[0].calldata)
	require.Equal(t, []common.Address{s.admin}, r.calls[0].senders)
}

func TestUpgradePortalAction(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	r := &mockRunner{}
	s := mockShared()
	c := mockChain()

	_, err := upgradePortal(ctx, s, c, r)
	require.NoError(t, err)
	require.Len(t, r.calls, 1)
	require.Equal(t, mustPack(t, adminABI, "upgradePortal", s.admin, s.deployer, c.PortalAddress, []byte{}), r.calls[0].calldata)
	require.Equal(t, []common.Address{s.admin, s.deployer}, r.calls[0].senders)
}

func mockShared() shared {
	return shared{
		admin:    common.HexToAddress("0x1"),
		deployer: common.HexToAddress("0x2"),
	}
}

func mockChain() chain {
	return chain{
		Chain: netconf.Chain{
			PortalAddress: common.HexToAddress("0x3"),
		},
		rpc: "http://localhost:8545",
	}
}

func mustPack(t *testing.T, abi *abi.ABI, method string, args ...interface{}) []byte {
	t.Helper()
	data, err := abi.Pack(method, args...)
	require.NoError(t, err)

	return data
}
