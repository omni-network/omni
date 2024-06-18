package app

import (
	"context"
	"testing"

	atypes "github.com/omni-network/omni/halo/attest/types"
	etypes "github.com/omni-network/omni/octane/evmengine/types"

	abci "github.com/cometbft/cometbft/abci/types"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

func TestProcessProposalRouter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		payloadMsgs int
		voteMsgs    int
		stakingMsgs int
	}{
		{
			name: "no messages",
		},
		{
			name:        "one payload message",
			payloadMsgs: 1,
		},
		{
			name:     "one vote message",
			voteMsgs: 1,
		},
		{
			name:        "one of each message",
			payloadMsgs: 1,
			voteMsgs:    1,
		},
		{
			name:        "two payload messages",
			payloadMsgs: 2,
		},
		{
			name:     "two vote messages",
			voteMsgs: 2,
		},
		{
			name:     "staking messages",
			voteMsgs: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			key := storetypes.NewKVStoreKey("test")
			ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))

			srv := &mockServer{}
			encCfg := moduletestutil.MakeTestEncodingConfig()
			txConfig := encCfg.TxConfig

			reg := encCfg.InterfaceRegistry
			etypes.RegisterInterfaces(reg)
			atypes.RegisterInterfaces(reg)
			stypes.RegisterInterfaces(reg)

			router := baseapp.NewMsgServiceRouter()
			router.SetInterfaceRegistry(reg)
			etypes.RegisterMsgServiceServer(router, srv)
			atypes.RegisterMsgServiceServer(router, srv)
			stypes.RegisterMsgServer(router, srv)

			handler := makeProcessProposalHandler(router, txConfig)

			var msgs []types.Msg
			for i := 0; i < tt.payloadMsgs; i++ {
				msgs = append(msgs, &etypes.MsgExecutionPayload{})
			}
			for i := 0; i < tt.voteMsgs; i++ {
				msgs = append(msgs, &atypes.MsgAddVotes{})
			}
			for i := 0; i < tt.stakingMsgs; i++ {
				msgs = append(msgs, &stypes.MsgDelegate{})
			}

			newReq := func(msgs ...types.Msg) *abci.RequestProcessProposal {
				b := txConfig.NewTxBuilder()
				err := b.SetMsgs(msgs...)
				require.NoError(t, err)

				tx, err := txConfig.TxEncoder()(b.GetTx())
				require.NoError(t, err)

				return &abci.RequestProcessProposal{
					Height: 0,
					Txs:    [][]byte{tx},
				}
			}

			accept := abci.ResponseProcessProposal_ACCEPT
			reject := abci.ResponseProcessProposal_REJECT

			res, err := handler(ctx, newReq(msgs...))
			require.NoError(t, err)
			if tt.stakingMsgs > 0 {
				require.Equal(t, reject, res.Status)
				require.Empty(t, srv.addVotes)
				require.Empty(t, srv.payload)
			} else if tt.payloadMsgs > 1 {
				require.Equal(t, reject, res.Status)
				require.Equal(t, 1, srv.payload)
				require.Empty(t, srv.addVotes)
			} else if tt.voteMsgs > 1 {
				require.Equal(t, reject, res.Status)
				require.Equal(t, 1, srv.addVotes)
				require.Empty(t, srv.payload)
			} else {
				require.Equal(t, accept, res.Status)
				require.Equal(t, tt.payloadMsgs, srv.payload)
				require.Equal(t, tt.voteMsgs, srv.addVotes)
			}
		})
	}
}

var _ atypes.MsgServiceServer = &mockServer{}
var _ etypes.MsgServiceServer = &mockServer{}
var _ stypes.MsgServer = &mockServer{}

type mockServer struct {
	stypes.MsgServer
	addVotes int
	payload  int
}

func (s *mockServer) ExecutionPayload(context.Context, *etypes.MsgExecutionPayload) (*etypes.ExecutionPayloadResponse, error) {
	s.payload++
	return &etypes.ExecutionPayloadResponse{}, nil
}

func (s *mockServer) AddVotes(context.Context, *atypes.MsgAddVotes) (*atypes.AddVotesResponse, error) {
	s.addVotes++
	return &atypes.AddVotesResponse{}, nil
}
