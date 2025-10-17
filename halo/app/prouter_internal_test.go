package app

import (
	"context"
	"testing"

	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/netconf"
	etypes "github.com/omni-network/omni/octane/evmengine/types"

	abci "github.com/cometbft/cometbft/abci/types"
	cmttypes "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

func TestProcessProposalRouter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		first       bool
		payloadMsgs int
		voteMsgs    int
		stakingMsgs int
		accept      bool
		multiTx     bool
	}{
		{
			name:   "first empty",
			first:  true,
			accept: true,
		},
		{
			name:        "first not empty",
			first:       true,
			accept:      false,
			payloadMsgs: 1,
		},
		{
			name:        "too many txs",
			payloadMsgs: 1,
			accept:      false,
			multiTx:     true,
		},
		{
			name:        "one payload message",
			payloadMsgs: 1,
			accept:      true,
		},
		{
			name:     "one vote message",
			voteMsgs: 1,
			accept:   true,
		},
		{
			name:        "one of each message",
			payloadMsgs: 1,
			voteMsgs:    1,
			accept:      true,
		},
		{
			name:        "two payload messages",
			payloadMsgs: 2,
			accept:      false,
		},
		{
			name:     "two vote messages",
			voteMsgs: 2,
			accept:   false,
		},
		{
			name:        "staking messages",
			stakingMsgs: 1,
			accept:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authority := authtypes.NewModuleAddress("test").String()
			key := storetypes.NewKVStoreKey("test")
			ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))

			srv := &mockServer{}
			encCfg, err := ClientEncodingConfig(ctx, netconf.Devnet)
			require.NoError(t, err)
			txConfig := encCfg.TxConfig

			router := baseapp.NewMsgServiceRouter()
			router.SetInterfaceRegistry(encCfg.InterfaceRegistry)
			etypes.RegisterMsgServiceServer(router, srv)
			atypes.RegisterMsgServiceServer(router, srv)
			stypes.RegisterMsgServer(router, srv)

			handler := makeProcessProposalHandler(router, txConfig)

			var msgs []types.Msg
			for i := 0; i < tt.payloadMsgs; i++ {
				msgs = append(msgs, &etypes.MsgExecutionPayload{Authority: authority})
			}
			for i := 0; i < tt.voteMsgs; i++ {
				msgs = append(msgs, &atypes.MsgAddVotes{Authority: authority})
			}
			for i := 0; i < tt.stakingMsgs; i++ {
				msgs = append(msgs, &stypes.MsgDelegate{DelegatorAddress: authority})
			}

			newReq := func(msgs ...types.Msg) *abci.RequestProcessProposal {
				b := txConfig.NewTxBuilder()
				err := b.SetMsgs(msgs...)
				require.NoError(t, err)

				tx, err := txConfig.TxEncoder()(b.GetTx())
				require.NoError(t, err)

				txs := [][]byte{tx}
				if len(msgs) == 0 {
					txs = nil
				} else if tt.multiTx {
					txs = append(txs, tx)
				}

				height := int64(99)
				if tt.first {
					height = 1
				}

				return &abci.RequestProcessProposal{
					Height: height,
					Txs:    txs,
					ProposedLastCommit: abci.CommitInfo{
						Votes: []abci.VoteInfo{
							{BlockIdFlag: cmttypes.BlockIDFlagCommit, Validator: abci.Validator{Power: 1}},
						},
					},
				}
			}

			accept := abci.ResponseProcessProposal_ACCEPT
			reject := abci.ResponseProcessProposal_REJECT

			res, err := handler(ctx, newReq(msgs...))
			require.NoError(t, err)
			require.Equal(t, tt.accept, res.Status == accept)
			require.Equal(t, !tt.accept, res.Status == reject)
			if tt.stakingMsgs > 0 {
				require.Empty(t, srv.addVotes)
				require.Empty(t, srv.payload)
			} else if tt.payloadMsgs > 1 {
				require.Equal(t, 1, srv.payload)
				require.Empty(t, srv.addVotes)
			} else if tt.voteMsgs > 1 {
				require.Equal(t, 1, srv.addVotes)
				require.Empty(t, srv.payload)
			} else if tt.accept {
				require.Equal(t, tt.payloadMsgs, srv.payload)
				require.Equal(t, tt.voteMsgs, srv.addVotes)
			}
		})
	}
}

func TestVerifyTx(t *testing.T) {
	t.Parallel()

	authority := authtypes.NewModuleAddress("test").String()

	tests := []struct {
		Name     string
		Msgs     []types.Msg
		Callback func(client.TxBuilder)
		Error    string
	}{
		{
			Name:  "empty",
			Msgs:  nil,
			Error: "",
		},
		{
			Name: "one payload message",
			Msgs: []types.Msg{&etypes.MsgExecutionPayload{Authority: authority}},
		},
		{
			Name: "two payload messages",
			Msgs: []types.Msg{
				&etypes.MsgExecutionPayload{Authority: authority},
				&etypes.MsgExecutionPayload{Authority: authority},
			},
			Error: "",
		},
		{
			Name:  "no authority",
			Msgs:  []types.Msg{&etypes.MsgExecutionPayload{}},
			Error: "get signers: empty address string is not allowed",
		},
		{
			Name: "fee not empty",
			Msgs: []types.Msg{&atypes.MsgAddVotes{Authority: authority}},
			Callback: func(b client.TxBuilder) {
				b.SetFeeAmount(types.Coins{types.NewCoin(types.DefaultBondDenom, math.NewInt(1))})
			},
			Error: "fee not empty",
		},
		{
			Name: "gas not empty",
			Msgs: []types.Msg{&etypes.MsgExecutionPayload{Authority: authority}},
			Callback: func(b client.TxBuilder) {
				b.SetGasLimit(1)
			},
			Error: "gas not empty",
		},
		{
			Name: "timeout height",
			Msgs: []types.Msg{&etypes.MsgExecutionPayload{Authority: authority}},
			Callback: func(b client.TxBuilder) {
				b.SetTimeoutHeight(1)
			},
			Error: "timeout height not empty",
		},
		{
			Name: "fee granter",
			Msgs: []types.Msg{&etypes.MsgExecutionPayload{Authority: authority}},
			Callback: func(b client.TxBuilder) {
				b.SetFeeGranter(authtypes.NewModuleAddress("granter"))
			},
			Error: "fee granter not empty",
		},
		{
			Name: "memo",
			Msgs: []types.Msg{&etypes.MsgExecutionPayload{Authority: authority}},
			Callback: func(b client.TxBuilder) {
				b.SetMemo("memo")
			},
			Error: "memo not empty",
		},
		{
			Name: "extension options",
			Msgs: []types.Msg{&etypes.MsgExecutionPayload{Authority: authority}},
			Callback: func(b client.TxBuilder) {
				extBuilder := b.(authtx.ExtensionOptionsTxBuilder)
				extBuilder.SetExtensionOptions(&codectypes.Any{})
			},
			Error: "extension options not empty",
		},
		{
			Name: "proto tx auth info nil",
			Msgs: []types.Msg{&etypes.MsgExecutionPayload{Authority: authority}},
			Callback: func(b client.TxBuilder) {
				ptx := b.(protoTxProvider).GetProtoTx()
				ptx.AuthInfo = nil
			},
			Error: "proto tx auth info is nil",
		},
		{
			Name: "proto tx auth info fee nil",
			Msgs: []types.Msg{&etypes.MsgExecutionPayload{Authority: authority}},
			Callback: func(b client.TxBuilder) {
				ptx := b.(protoTxProvider).GetProtoTx()
				ptx.AuthInfo.Fee = nil
			},
			Error: "proto tx auth info fee is nil",
		},
		{
			Name: "auth info tip",
			Msgs: []types.Msg{&etypes.MsgExecutionPayload{Authority: authority}},
			Callback: func(b client.TxBuilder) {
				ptx := b.(protoTxProvider).GetProtoTx()
				ptx.AuthInfo.Tip = &txtypes.Tip{} //nolint:staticcheck // For testing purposes
			},
			Error: "proto tx tip not nil",
		},
		{
			Name: "raw signatures",
			Msgs: []types.Msg{&etypes.MsgExecutionPayload{Authority: authority}},
			Callback: func(b client.TxBuilder) {
				ptx := b.(protoTxProvider).GetProtoTx()
				ptx.Signatures = [][]byte{{}, {}}
			},
			Error: "proto tx signatures not empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			key := storetypes.NewKVStoreKey("test")
			ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))

			encCfg, err := ClientEncodingConfig(ctx, netconf.Devnet)
			require.NoError(t, err)
			txConfig := encCfg.TxConfig

			b := txConfig.NewTxBuilder()
			if tt.Callback != nil {
				tt.Callback(b)
			}
			err = b.SetMsgs(tt.Msgs...)
			require.NoError(t, err)

			err = verifyTX(b.GetTx())

			if tt.Error != "" {
				require.ErrorContains(t, err, tt.Error)
			} else {
				require.NoError(t, err)
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
