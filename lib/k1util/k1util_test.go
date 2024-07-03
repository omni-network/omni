// Copyright © 2022-2023 Obol Labs Inc. Licensed under the terms of a Business Source License 1.1

package k1util_test

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/tutil"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/cometbft/cometbft/privval"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttypes "github.com/cometbft/cometbft/types"
	cmttime "github.com/cometbft/cometbft/types/time"

	"github.com/ethereum/go-ethereum/crypto"

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

func TestCosmosPbukey(t *testing.T) {
	t.Parallel()

	priv, err := crypto.GenerateKey()
	require.NoError(t, err)

	cosmosPub, err := k1util.StdPubKeyToCosmos(&priv.PublicKey)
	require.NoError(t, err)

	require.NotPanics(t, func() {
		cosmosPub.Address()
	})
}

// TestCometBFT tests that CometBFT and k1util can produce the same signatures.
// This ensures that we can use web3signer to sign votes and proposals.
func TestCometBFT(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	const chainID = "test"

	// Create a key
	key := k1.GenPrivKey()
	pv := privval.NewFilePV(key, filepath.Join(dir, "key"), filepath.Join(dir, "state"))

	// Create a vote
	block := cmttypes.BlockID{Hash: tutil.RandomBytes(32)}
	vote := newVote(pv.Key.Address, 1, 2, 3, cmtproto.PrecommitType, block, []byte("vote extension"))
	votePB1 := vote.ToProto()
	votePB2 := vote.ToProto()

	// Sign it with CometBFT privval.
	err := pv.SignVote(chainID, votePB1)
	require.NoError(t, err)

	require.Nil(t, votePB2.Signature) // Ensure no pointer shenanigans.

	// Sign it with k1util.
	signVote(t, key, chainID, votePB2)

	require.Equal(t, votePB1.Signature, votePB2.Signature)
	require.Equal(t, votePB1.ExtensionSignature, votePB2.ExtensionSignature)
}

func TestPubkey64(t *testing.T) {
	t.Parallel()

	priv, err := crypto.GenerateKey()
	require.NoError(t, err)

	bz64, err := k1util.PubKeyToBytes64(&priv.PublicKey)
	require.NoError(t, err)
	require.Len(t, bz64, 64)

	pub, err := k1util.PubKeyFromBytes64(bz64)
	require.NoError(t, err)

	require.True(t, pub.Equal(&priv.PublicKey))
}

func signVote(t *testing.T, key k1.PrivKey, chainID string, vote *cmtproto.Vote) {
	t.Helper()

	sigBytes := cmttypes.VoteSignBytes(chainID, vote)
	hashed := sha256.Sum256(sigBytes)
	sig, err := k1util.Sign(key, hashed)
	require.NoError(t, err)

	extSignBytes := cmttypes.VoteExtensionSignBytes(chainID, vote)
	extHashed := sha256.Sum256(extSignBytes)
	extSig, err := k1util.Sign(key, extHashed)
	require.NoError(t, err)

	vote.Signature = sig[:64]             // CometBFT drops recovery id.
	vote.ExtensionSignature = extSig[:64] // CometBFT drops recovery id.
}

func fromHex(t *testing.T, hexStr string) []byte {
	t.Helper()
	b, err := hex.DecodeString(hexStr)
	require.NoError(t, err)

	return b
}

func newVote(addr cmttypes.Address, idx int32, height int64, round int32,
	typ cmtproto.SignedMsgType, blockID cmttypes.BlockID, extension []byte) *cmttypes.Vote {
	return &cmttypes.Vote{
		ValidatorAddress: addr,
		ValidatorIndex:   idx,
		Height:           height,
		Round:            round,
		Type:             typ,
		Timestamp:        cmttime.Now(),
		BlockID:          blockID,
		Extension:        extension,
	}
}
