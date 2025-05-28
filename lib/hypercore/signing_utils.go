// This file implements python Hyperliquid's signing utils in golang.
// ref: https://github.com/hyperliquid-dex/hyperliquid-python-sdk/tree/a8edca1
package hypercore

import (
	"context"
	"encoding/binary"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"

	"github.com/ugorji/go/codec"
)

var emptyAddress = common.Address{}

// l1Payload returns an EIP-712 typed data payload for the PhantomAgent.
// ref: https://github.com/hyperliquid-dex/hyperliquid-python-sdk/blob/a8edca1/hyperliquid/utils/signing.py#L173
func l1Payload(agent PhantomAgent) apitypes.TypedData {
	return apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			ChainId:           math.NewHexOrDecimal256(1337),
			Name:              "Exchange",
			VerifyingContract: "0x0000000000000000000000000000000000000000",
			Version:           "1",
		},
		Types: apitypes.Types{
			"Agent": []apitypes.Type{
				{Name: "source", Type: "string"},
				{Name: "connectionId", Type: "bytes32"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "Agent",
		Message: map[string]any{
			"source":       agent.Source,
			"connectionId": agent.ConnectionID[:],
		},
	}
}

// phantomAgent creates a PhantomAgent with the given hash and network type.
// ref: https://github.com/hyperliquid-dex/hyperliquid-python-sdk/blob/a8edca1/hyperliquid/utils/signing.py#L169C5-L169C28
func phantomAgent(hash [32]byte, isMainnet bool) PhantomAgent {
	source := "a"
	if !isMainnet {
		source = "b"
	}

	return PhantomAgent{Source: source, ConnectionID: hash}
}

// signL1Action signs a Hyperliquid L1 (core) action using the provided signer.
// ref: https://github.com/hyperliquid-dex/hyperliquid-python-sdk/blob/a8edca1/hyperliquid/utils/signing.py#L221
func signL1Action(
	ctx context.Context,
	signer Signer,
	action any,
	activePool common.Address,
	nonce uint64,
	expiresAfter uint64,
	isMainnet bool,
) (SigRSV, error) {
	hash, err := actionHash(action, activePool, nonce, expiresAfter)
	if err != nil {
		return SigRSV{}, errors.Wrap(err, "action hash")
	}

	agent := phantomAgent(hash, isMainnet)
	data := l1Payload(agent)

	return signInner(ctx, signer, data)
}

// actionHash returns the keccak256 hash of encoded action data
// ref: https://github.com/hyperliquid-dex/hyperliquid-python-sk/blob/a8edca1/hyperliquid/utils/signing.py#L155
func actionHash(
	action any,
	vault common.Address,
	nonce uint64,
	expiresAfter uint64,
) ([32]byte, error) {
	data, err := msgpack(action)
	if err != nil {
		return [32]byte{}, err
	}

	data = binary.BigEndian.AppendUint64(data, nonce)

	if vault == emptyAddress {
		data = append(data, 0)
	} else {
		data = append(data, 1)
		data = append(data, vault.Bytes()...)
	}

	if expiresAfter != 0 {
		data = append(data, 0)
		data = binary.BigEndian.AppendUint64(data, expiresAfter)
	}

	return crypto.Keccak256Hash(data), nil
}

// signInner signs typed data using the provided signer.
// ref: https://github.com/hyperliquid-dex/hyperliquid-python-sdk/blob/a8edca1/hyperliquid/utils/signing.py#L413
func signInner(
	ctx context.Context,
	signer Signer,
	data apitypes.TypedData,
) (SigRSV, error) {
	hash, _, err := apitypes.TypedDataAndHash(data)
	if err != nil {
		return SigRSV{}, errors.Wrap(err, "typed data and hash")
	}

	sig, err := signer.Sign(ctx, hash)
	if err != nil {
		return SigRSV{}, errors.Wrap(err, "sign")
	}

	return bzToRSV(sig), nil
}

// timestampMS returns the current timestamp in milliseconds (used for nonce generation).
// ref: https://github.com/hyperliquid-dex/hyperliquid-python-sdk/blob/a8edca1/hyperliquid/utils/signing.py#L462
func timestampMS() uint64 {
	return uint64(time.Now().UnixMilli()) //nolint:gosec // time not negative
}

func bzToRSV(sig [65]byte) SigRSV {
	// Convert to big.Int, to trim leading zeros
	// This matches the behavior of Hyperliquid's Python SDK
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:64])

	return SigRSV{
		R: "0x" + r.Text(16),
		S: "0x" + s.Text(16),
		V: sig[64] + 27,
	}
}

func msgpack(v any) ([]byte, error) {
	var data []byte
	handle := &codec.MsgpackHandle{}
	enc := codec.NewEncoderBytes(&data, handle)
	if err := enc.Encode(v); err != nil {
		return nil, errors.Wrap(err, "msgpack encode")
	}

	return data, nil
}
