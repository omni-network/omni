package app_test

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/omni-network/omni/e2e/fbproxy/app"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

type mockTxSigner struct {
	pk *ecdsa.PrivateKey
}

const (
	testSigner   = "0x71562b71999873db5b286df957af199ec94617f7"
	testSignerPK = "b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291"
)

var (
	chainID = big.NewInt(1337)
)

func (b *mockTxSigner) Sign(ctx context.Context, digest common.Hash, signer common.Address) ([65]byte, error) {
	if signer != common.HexToAddress(testSigner) {
		return [65]byte{}, errors.New("invalid signer")
	}

	sig, err := crypto.Sign(digest.Bytes(), b.pk)
	if err != nil {
		return [65]byte{}, errors.Wrap(err, "sign")
	}

	return [65]byte(sig), nil
}

func newMockTxSigner() *mockTxSigner {
	pk, err := crypto.HexToECDSA(testSignerPK)
	if err != nil {
		panic(err)
	}

	return &mockTxSigner{pk: pk}
}

// testSendTxMiddlewareOnce tests that the middleware correctly signs a transaction
// and replaces the eth_sendTransaction request with an eth_sendRawTransaction request.
func testSendTxMiddlewareOnce(t *testing.T, tt testTx) {
	t.Helper()

	txsigner := newMockTxSigner()

	var txargs app.TransactionArgs
	err := json.Unmarshal([]byte(tt.txArgsJSON), &txargs)
	require.NoError(t, err)

	mw := app.NewSendTxMiddleware(txsigner, chainID.Uint64())
	req := app.JSONRPCMessage{
		Method:  "eth_sendTransaction",
		Params:  mustMarshal(t, []app.TransactionArgs{txargs}),
		ID:      mustMarshal(t, 1),
		Version: "2.0",
	}

	// Test that the middleware signs the transaction
	resp, err := mw(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, "eth_sendRawTransaction", resp.Method)

	var params []string
	err = json.Unmarshal(resp.Params, &params)
	require.NoError(t, err)

	// tx returned from middleware
	var tx1 types.Transaction
	err = tx1.UnmarshalBinary(hexutil.MustDecode(params[0]))
	require.NoError(t, err)

	// tx signed now
	signer := types.LatestSignerForChainID(chainID)
	tx2, err := types.SignNewTx(txsigner.pk, signer, tt.tx)
	require.NoError(t, err)

	// tx with hardcoded sig values
	var tx3 types.Transaction
	err = tx3.UnmarshalJSON([]byte(tt.txJSON))
	require.NoError(t, err)

	// check hashes are equal
	require.Equal(t, tx1.Hash(), tx2.Hash())
	require.Equal(t, tx1.Hash(), tx3.Hash())

	// check that the signature values are equal
	v1, r1, s1 := tx1.RawSignatureValues()
	v2, r2, s2 := tx2.RawSignatureValues()
	v3, r3, s3 := tx3.RawSignatureValues()

	require.Equal(t, v1, v2)
	require.Equal(t, r1, r2)
	require.Equal(t, s1, s2)
	require.Equal(t, v1, v3)
	require.Equal(t, r1, r3)
	require.Equal(t, s1, s3)
}

func TestSendTxMiddleware(t *testing.T) {
	t.Parallel()

	for _, tt := range testTxns() {
		testSendTxMiddlewareOnce(t, tt)
	}
}

type testTx struct {
	tx         types.TxData
	txJSON     string
	txArgsJSON string
}

// testTxns returns a list of test transactions.
// Repurposed from test cases go-ethereum/internal/ethapi/api_test.go
// TxArgsJson is trimmed TxJson, removing fields that are not part of the TransactionArgs.
func testTxns() []testTx {
	dead := common.HexToAddress("0xdead000000000000000000000000000000000000")

	return []testTx{
		{
			tx: &types.LegacyTx{
				Nonce:    5,
				GasPrice: big.NewInt(6),
				Gas:      7,
				To:       &dead,
				Value:    big.NewInt(8),
				Data:     []byte{0, 1, 2, 3, 4},
				V:        big.NewInt(9),
				R:        big.NewInt(10),
				S:        big.NewInt(11),
			},
			txArgsJSON: `{
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x6",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": "0xdead000000000000000000000000000000000000",
				"value": "0x8",
				"chainId": "0x539"
			}`,
			txJSON: `{
				"blockHash": null,
				"blockNumber": null,
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x6",
				"hash": "0x5f3240454cd09a5d8b1c5d651eefae7a339262875bcd2d0e6676f3d989967008",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": "0xdead000000000000000000000000000000000000",
				"transactionIndex": null,
				"value": "0x8",
				"type": "0x0",
				"chainId": "0x539",
				"v": "0xa96",
				"r": "0xbc85e96592b95f7160825d837abb407f009df9ebe8f1b9158a4b8dd093377f75",
				"s": "0x1b55ea3af5574c536967b039ba6999ef6c89cf22fc04bcb296e0e8b0b9b576f5"
			}`,
		},
		{
			tx: &types.LegacyTx{
				Nonce:    5,
				GasPrice: big.NewInt(6),
				Gas:      7,
				To:       nil,
				Value:    big.NewInt(8),
				Data:     []byte{0, 1, 2, 3, 4},
				V:        big.NewInt(32),
				R:        big.NewInt(10),
				S:        big.NewInt(11),
			},
			txArgsJSON: `{
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x6",
				"hash": "0x806e97f9d712b6cb7e781122001380a2837531b0fc1e5f5d78174ad4cb699873",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": null,
				"value": "0x8",
				"chainId": "0x539"
			}`,
			txJSON: `{
				"blockHash": null,
				"blockNumber": null,
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x6",
				"hash": "0x806e97f9d712b6cb7e781122001380a2837531b0fc1e5f5d78174ad4cb699873",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": null,
				"transactionIndex": null,
				"value": "0x8",
				"type": "0x0",
				"chainId": "0x539",
				"v": "0xa96",
				"r": "0x9dc28b267b6ad4e4af6fe9289668f9305c2eb7a3241567860699e478af06835a",
				"s": "0xa0b51a071aa9bed2cd70aedea859779dff039e3630ea38497d95202e9b1fec7"
			}`,
		},
		{
			tx: &types.AccessListTx{
				ChainID:  chainID,
				Nonce:    5,
				GasPrice: big.NewInt(6),
				Gas:      7,
				To:       &dead,
				Value:    big.NewInt(8),
				Data:     []byte{0, 1, 2, 3, 4},
				AccessList: types.AccessList{
					types.AccessTuple{
						Address:     common.Address{0x2},
						StorageKeys: []common.Hash{types.EmptyRootHash},
					},
				},
				V: big.NewInt(32),
				R: big.NewInt(10),
				S: big.NewInt(11),
			},
			txArgsJSON: `{
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x6",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": "0xdead000000000000000000000000000000000000",
				"value": "0x8",
				"accessList": [
					{
						"address": "0x0200000000000000000000000000000000000000",
						"storageKeys": [
							"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"
						]
					}
				],
				"chainId": "0x539"
			}`,
			txJSON: `{
				"blockHash": null,
				"blockNumber": null,
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x6",
				"hash": "0x121347468ee5fe0a29f02b49b4ffd1c8342bc4255146bb686cd07117f79e7129",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": "0xdead000000000000000000000000000000000000",
				"transactionIndex": null,
				"value": "0x8",
				"type": "0x1",
				"accessList": [
					{
						"address": "0x0200000000000000000000000000000000000000",
						"storageKeys": [
							"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"
						]
					}
				],
				"chainId": "0x539",
				"v": "0x0",
				"r": "0xf372ad499239ae11d91d34c559ffc5dab4daffc0069e03afcabdcdf231a0c16b",
				"s": "0x28573161d1f9472fa0fd4752533609e72f06414f7ab5588699a7141f65d2abf",
				"yParity": "0x0"
			}`,
		},
		{
			tx: &types.AccessListTx{
				ChainID:  chainID,
				Nonce:    5,
				GasPrice: big.NewInt(6),
				Gas:      7,
				To:       nil,
				Value:    big.NewInt(8),
				Data:     []byte{0, 1, 2, 3, 4},
				AccessList: types.AccessList{
					types.AccessTuple{
						Address:     common.Address{0x2},
						StorageKeys: []common.Hash{types.EmptyRootHash},
					},
				},
				V: big.NewInt(32),
				R: big.NewInt(10),
				S: big.NewInt(11),
			},
			txArgsJSON: `{
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x6",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": null,
				"value": "0x8",
				"accessList": [
					{
						"address": "0x0200000000000000000000000000000000000000",
						"storageKeys": [
							"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"
						]
					}
				],
				"chainId": "0x539"
			}`,
			txJSON: `{
				"blockHash": null,
				"blockNumber": null,
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x6",
				"hash": "0x067c3baebede8027b0f828a9d933be545f7caaec623b00684ac0659726e2055b",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": null,
				"transactionIndex": null,
				"value": "0x8",
				"type": "0x1",
				"accessList": [
					{
						"address": "0x0200000000000000000000000000000000000000",
						"storageKeys": [
							"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"
						]
					}
				],
				"chainId": "0x539",
				"v": "0x1",
				"r": "0x542981b5130d4613897fbab144796cb36d3cb3d7807d47d9c7f89ca7745b085c",
				"s": "0x7425b9dd6c5deaa42e4ede35d0c4570c4624f68c28d812c10d806ffdf86ce63",
				"yParity": "0x1"
			}`,
		},
		{
			tx: &types.DynamicFeeTx{
				ChainID:   chainID,
				Nonce:     5,
				GasTipCap: big.NewInt(6),
				GasFeeCap: big.NewInt(9),
				Gas:       7,
				To:        &dead,
				Value:     big.NewInt(8),
				Data:      []byte{0, 1, 2, 3, 4},
				AccessList: types.AccessList{
					types.AccessTuple{
						Address:     common.Address{0x2},
						StorageKeys: []common.Hash{types.EmptyRootHash},
					},
				},
				V: big.NewInt(32),
				R: big.NewInt(10),
				S: big.NewInt(11),
			},
			txArgsJSON: `{
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x9",
				"maxFeePerGas": "0x9",
				"maxPriorityFeePerGas": "0x6",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": "0xdead000000000000000000000000000000000000",
				"value": "0x8",
				"accessList": [
					{
						"address": "0x0200000000000000000000000000000000000000",
						"storageKeys": [
							"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"
						]
					}
				],
				"chainId": "0x539"
			}`,
			txJSON: `{
				"blockHash": null,
				"blockNumber": null,
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x9",
				"maxFeePerGas": "0x9",
				"maxPriorityFeePerGas": "0x6",
				"hash": "0xb63e0b146b34c3e9cb7fbabb5b3c081254a7ded6f1b65324b5898cc0545d79ff",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": "0xdead000000000000000000000000000000000000",
				"transactionIndex": null,
				"value": "0x8",
				"type": "0x2",
				"accessList": [
					{
						"address": "0x0200000000000000000000000000000000000000",
						"storageKeys": [
							"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"
						]
					}
				],
				"chainId": "0x539",
				"v": "0x1",
				"r": "0x3b167e05418a8932cd53d7578711fe1a76b9b96c48642402bb94978b7a107e80",
				"s": "0x22f98a332d15ea2cc80386c1ebaa31b0afebfa79ebc7d039a1e0074418301fef",
				"yParity": "0x1"
			}`,
		},
		{
			tx: &types.DynamicFeeTx{
				ChainID:    chainID,
				Nonce:      5,
				GasTipCap:  big.NewInt(6),
				GasFeeCap:  big.NewInt(9),
				Gas:        7,
				To:         nil,
				Value:      big.NewInt(8),
				Data:       []byte{0, 1, 2, 3, 4},
				AccessList: types.AccessList{},
				V:          big.NewInt(32),
				R:          big.NewInt(10),
				S:          big.NewInt(11),
			},
			txArgsJSON: `{
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x9",
				"maxFeePerGas": "0x9",
				"maxPriorityFeePerGas": "0x6",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": null,
				"value": "0x8",
				"accessList": [],
				"chainId": "0x539"
			}`,
			txJSON: `{
				"blockHash": null,
				"blockNumber": null,
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x9",
				"maxFeePerGas": "0x9",
				"maxPriorityFeePerGas": "0x6",
				"hash": "0xcbab17ee031a9d5b5a09dff909f0a28aedb9b295ac0635d8710d11c7b806ec68",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": null,
				"transactionIndex": null,
				"value": "0x8",
				"type": "0x2",
				"accessList": [],
				"chainId": "0x539",
				"v": "0x0",
				"r": "0x6446b8a682db7e619fc6b4f6d1f708f6a17351a41c7fbd63665f469bc78b41b9",
				"s": "0x7626abc15834f391a117c63450047309dbf84c5ce3e8e609b607062641e2de43",
				"yParity": "0x0"
			}`,
		},
	}
}

func mustMarshal(t *testing.T, v interface{}) []byte {
	t.Helper()
	bz, err := json.Marshal(v)
	require.NoError(t, err)

	return bz
}
