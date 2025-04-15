package cctp

import (
	crand "crypto/rand"
	"math/big"
	mrand "math/rand"
	"testing"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/types"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/mock"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	cosmosdb "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestStreamEventProc(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	ctrl := gomock.NewController(t)
	client := mock.NewMockClient(ctrl)

	chainID := evmchain.IDEthereum
	chainVer := xchain.ChainVersion{ID: chainID, ConfLevel: xchain.ConfFinalized}

	msgTransmitter, msgTransmitterAddr := mustMessageTransmitter(chainID, client)
	tknMessenger, tknMessengerAddr := mustTokenMessenger(chainID, client)

	type testCase struct {
		name          string
		chainID       uint64
		recipient     common.Address
		inititalState []types.MsgSendUSDC
		logs          []ethtypes.Log
		expectedState []types.MsgSendUSDC
	}
	tests := []func() testCase{
		// inserts single new
		func() testCase {
			msg := randMsg(chainID, randAddr())

			ev1 := depositForBurnLog(tknMessengerAddr, msg)
			ev2 := messageSentLog(msgTransmitterAddr, msg)

			return testCase{
				name:          "inserts single new",
				chainID:       chainID,
				recipient:     msg.Recipient,
				inititalState: []types.MsgSendUSDC{},
				logs:          []ethtypes.Log{ev1, ev2},
				expectedState: []types.MsgSendUSDC{msg},
			}
		},

		// inserts multiple new
		func() testCase {
			recipient := randAddr()

			msg1 := randMsg(chainID, recipient)
			msg2 := randMsg(chainID, recipient)

			ev1 := depositForBurnLog(tknMessengerAddr, msg1)
			ev2 := messageSentLog(msgTransmitterAddr, msg1)
			ev3 := depositForBurnLog(tknMessengerAddr, msg2)
			ev4 := messageSentLog(msgTransmitterAddr, msg2)

			return testCase{
				name:          "inserts multiple new",
				chainID:       chainID,
				recipient:     recipient,
				inititalState: []types.MsgSendUSDC{},
				logs:          []ethtypes.Log{ev1, ev2, ev3, ev4},
				expectedState: []types.MsgSendUSDC{msg1, msg2},
			}
		},

		// ignore different tx hash
		func() testCase {
			msg := randMsg(chainID, randAddr())

			ev1 := depositForBurnLog(tknMessengerAddr, msg)
			ev2 := messageSentLog(msgTransmitterAddr, msg)
			ev2.TxHash = common.BytesToHash(mustRandBytes(32)) // different tx hash

			return testCase{
				name:          "ignore different tx hash",
				chainID:       chainID,
				recipient:     msg.Recipient,
				inititalState: []types.MsgSendUSDC{},
				logs:          []ethtypes.Log{ev1, ev2},
				expectedState: []types.MsgSendUSDC{},
			}
		},

		// ignore unknown recipient
		func() testCase {
			msg := randMsg(chainID, randAddr())

			ev1 := depositForBurnLog(tknMessengerAddr, msg)
			ev2 := messageSentLog(msgTransmitterAddr, msg)

			return testCase{
				name:          "ignore unknown recipient",
				chainID:       chainID,
				recipient:     randAddr(), // different recipient
				inititalState: []types.MsgSendUSDC{},
				logs:          []ethtypes.Log{ev1, ev2},
				expectedState: []types.MsgSendUSDC{},
			}
		},

		// ignore unknown topic
		func() testCase {
			msg := randMsg(chainID, randAddr())

			ev1 := depositForBurnLog(tknMessengerAddr, msg)
			ev2 := messageSentLog(msgTransmitterAddr, msg)
			ev2.Topics[0] = common.HexToHash("0x123") // different topic

			return testCase{
				name:          "ignore unknown topic",
				chainID:       chainID,
				recipient:     msg.Recipient,
				inititalState: []types.MsgSendUSDC{},
				logs:          []ethtypes.Log{ev1, ev2},
				expectedState: []types.MsgSendUSDC{},
			}
		},

		// updates existing message hash (sims reorg)
		func() testCase {
			// initial state
			msg1 := randMsg(chainID, randAddr())

			// expected state, used for lgos
			msg2 := msg1
			msg2.MessageBytes = mustRandBytes(32 * 10)
			msg2.MessageHash = crypto.Keccak256Hash(msg2.MessageBytes)

			ev1 := depositForBurnLog(tknMessengerAddr, msg2)
			ev2 := messageSentLog(msgTransmitterAddr, msg2)

			return testCase{
				name:          "updates existing message hash",
				chainID:       chainID,
				recipient:     msg1.Recipient,
				inititalState: []types.MsgSendUSDC{msg1},
				logs:          []ethtypes.Log{ev1, ev2},
				expectedState: []types.MsgSendUSDC{msg2},
			}
		},

		// keeps existing db messages
		func() testCase {
			initialMsgs := []types.MsgSendUSDC{randMsg(chainID, randAddr()), randMsg(chainID, randAddr())}
			newMsg := randMsg(chainID, randAddr())

			ev1 := depositForBurnLog(tknMessengerAddr, newMsg)
			ev2 := messageSentLog(msgTransmitterAddr, newMsg)

			return testCase{
				name:          "keeps existing db messages",
				chainID:       chainID,
				recipient:     newMsg.Recipient,
				inititalState: initialMsgs,
				logs:          []ethtypes.Log{ev1, ev2},
				expectedState: append(initialMsgs, newMsg),
			}
		},
	}
	for _, tf := range tests {
		tt := tf()

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			memDB := cosmosdb.NewMemDB()
			db, err := db.New(memDB)
			require.NoError(t, err)

			for _, msg := range tt.inititalState {
				err := db.InsertMsg(ctx, msg)
				require.NoError(t, err)
			}

			proc := newEventProc(db, chainVer,
				newDepositForBurnGetter(tknMessenger, tknMessengerAddr, tt.recipient),
				newMessageSentGetter(msgTransmitter, msgTransmitterAddr),
			)

			// TODO(kevin): move to test case when source block number tracked in db
			header := &ethtypes.Header{Number: randBigInt()}

			err = proc(ctx, header, tt.logs)
			require.NoError(t, err)

			list, err := db.ListMsgs(ctx)
			require.NoError(t, err)

			assertSameMsgs(t, tt.expectedState, list)
		})
	}
}

// assertSameMsgs asserts that the expected messages are the same as the actual messages.
func assertSameMsgs(t *testing.T, expected, got []types.MsgSendUSDC) {
	t.Helper()

	require.Len(t, got, len(expected), "expected %d messages, got %d", len(expected), len(got))

	for _, expectedMsg := range expected {
		var match types.MsgSendUSDC
		found := false

		for _, gotMsg := range got {
			if expectedMsg.TxHash == gotMsg.TxHash {
				found = true
				match = gotMsg

				break
			}
		}

		require.Truef(t, found, "expected message not found in result: %s", expectedMsg.TxHash.Hex())
		require.Equalf(t, expectedMsg, match, "expected message does not match: %s", expectedMsg.TxHash.Hex())
	}
}

// depositForBurnLog creates a mock DepositForBurn event log.
func depositForBurnLog(addr common.Address, msg types.MsgSendUSDC) ethtypes.Log {
	mustPackData := func(args ...any) []byte {
		packed, err := depositForBurnEvent.Inputs.NonIndexed().Pack(args...)
		if err != nil {
			panic(err)
		}

		return packed
	}

	return ethtypes.Log{
		TxHash:  msg.TxHash,
		Address: addr,
		Topics: []common.Hash{
			depositForBurnEvent.ID,
			common.HexToHash("0x123"),             // nonce (doesn't matter)
			common.HexToHash("0x456"),             // burnToken (doesn't matter)
			common.HexToHash(msg.Recipient.Hex()), // depositor (use recipient, doesn't matter)
		},
		Data: mustPackData(
			msg.Amount,                       // amount
			cast.EthAddress32(msg.Recipient), // mintRecipient
			uint32(msg.DestChainID),          // destinationDomain
			common.HexToHash("0x1"),          // destinationTokenMessenger (doesn't matter)
			common.HexToHash("0x1"),          // destinationCaller (doesn't matter)
		),
	}
}

// messageSentLog creates a mock MessageSent event log.
func messageSentLog(addr common.Address, msg types.MsgSendUSDC) ethtypes.Log {
	return ethtypes.Log{
		TxHash:  msg.TxHash,
		Address: addr,
		Topics:  []common.Hash{messageSentEvent.ID},
		Data:    mustABIEncodeBytes(msg.MessageBytes),
	}
}

func randAddr() common.Address { return cast.MustEthAddress(mustRandBytes(20)) }
func randBigInt() *big.Int     { return big.NewInt(mrand.Int63()) }

func randMsg(srcChainID uint64, recipient common.Address) types.MsgSendUSDC {
	msgBz := mustRandBytes(32 * 10)
	msgHash := crypto.Keccak256Hash(msgBz)

	return types.MsgSendUSDC{
		TxHash:       common.BytesToHash(mustRandBytes(32)),
		SrcChainID:   srcChainID,
		DestChainID:  uint64(mrand.Uint32()),
		Amount:       randBigInt(),
		MessageBytes: msgBz,
		MessageHash:  msgHash,
		Recipient:    recipient,
	}
}

func mustMessageTransmitter(chainID uint64, client ethclient.Client) (*MessageTransmitter, common.Address) {
	contract, addr, err := newMessageTransmitter(chainID, client)
	if err != nil {
		panic("missing message transmitter")
	}

	return contract, addr
}

func mustTokenMessenger(chainID uint64, client ethclient.Client) (*TokenMessenger, common.Address) {
	contract, addr, err := newTokenMessenger(chainID, client)
	if err != nil {
		panic("missing token messenger")
	}

	return contract, addr
}

// mustRandBytes generates n random bytes.
func mustRandBytes(n int) []byte {
	bz := make([]byte, n)

	_, err := crand.Read(bz)
	if err != nil {
		panic(err)
	}

	return bz
}

// mustABIEncodeBytes encodes a byte slice into an ABI-encoded byte array.
func mustABIEncodeBytes(bz []byte) []byte {
	typ, err := abi.NewType("bytes", "", nil)
	if err != nil {
		panic(err)
	}

	packed, err := abi.Arguments{{Type: typ}}.Pack(bz)
	if err != nil {
		panic(err)
	}

	return packed
}
