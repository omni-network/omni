// Copyright © 2022-2023 Obol Labs Inc. Licensed under the terms of a Business Source License 1.1

package k1util_test

import (
	"crypto/rand"
	"encoding/hex"
	"testing"

	"github.com/omni-network/omni/lib/k1util"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/stretchr/testify/require"
)

//nolint:lll // No wrap
const (
	privKey1 = "41d3ff12045b73c870529fe44f70dca2745bafbe1698ffc3c8759eef3cfbaee1"
	pubKey1  = "02bc8e7cdb50e0ffd52a54faf984d6ac8fe5ee6856d38a5f8acd9bd33fc9c7d50d"
	digest1  = "52fdfc072182654f163f5f0f9a621d729566c74d10037c4d7bbb0407d1e2c649" // 32 byte digest.
	sig1     = "e08097bed6dc40d70aa0076f9d8250057566cdf40c652b3785ad9c06b1e38d584f8f331bf46f68e3737823a3bda905e90ca96735d510a6934b215753c09acec21c"
	addr1    = "0xF88D5892faF084DCF4143566d9C9b3F047c153Ca"
)

func TestK1Util(t *testing.T) {
	t.Parallel()
	key := k1.PrivKey(fromHex(t, privKey1))

	require.Equal(t, fromHex(t, privKey1), key.Bytes())
	require.Equal(t, fromHex(t, pubKey1), key.PubKey().Bytes())

	digest := fromHex(t, digest1)

	sig, err := k1util.Sign(key, [32]byte(digest))
	require.NoError(t, err)
	require.EqualValues(t, fromHex(t, sig1), sig[:])

	addr, err := k1util.PubKeyToAddress(key.PubKey())
	require.NoError(t, err)
	require.Equal(t, addr1, addr.Hex())

	ok, err := k1util.Verify(addr, [32]byte(digest), sig)
	require.NoError(t, err)
	require.True(t, ok)
}

func TestRandom(t *testing.T) {
	t.Parallel()
	key := k1.GenPrivKey()

	var digest [32]byte
	_, _ = rand.Read(digest[:])

	sig, err := k1util.Sign(key, digest)
	require.NoError(t, err)

	addr, err := k1util.PubKeyToAddress(key.PubKey())
	require.NoError(t, err)

	ok, err := k1util.Verify(addr, digest, sig)
	require.NoError(t, err)
	require.True(t, ok)
}

func fromHex(t *testing.T, hexStr string) []byte {
	t.Helper()
	b, err := hex.DecodeString(hexStr)
	require.NoError(t, err)

	return b
}
