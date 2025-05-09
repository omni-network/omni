package voter_test

import (
	"math/rand"
	"testing"

	"github.com/omni-network/omni/halo/attest/voter"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestCreateVerifyVotes(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 64)

	privKey := k1.GenPrivKey()
	addr, err := k1util.PubKeyToAddress(privKey.PubKey())
	require.NoError(t, err)

	var block xchain.Block
	fuzzer.Fuzz(&block)
	// Ensure msg.LogIndex is increasing
	for i := 1; i < len(block.Msgs); i++ {
		block.Msgs[i].LogIndex = block.Msgs[i-1].LogIndex + 1 + uint64(rand.Intn(1000))
	}

	var attHeader xchain.AttestHeader
	fuzzer.Fuzz(&attHeader)
	attHeader.ChainVersion.ID = block.ChainID            // Align headers
	attHeader.ChainVersion.ConfLevel = xchain.ConfLatest // Pick valid conf level

	att, err := voter.CreateVote(privKey, attHeader, block)
	require.NoError(t, err)
	header, err := att.BlockHeader.ToXChain()
	require.NoError(t, err)
	require.Equal(t, block.BlockHeader, header)
	require.Equal(t, addr, common.Address(att.Signature.ValidatorAddress))

	// Verify the attestation
	err = att.Verify()
	require.NoError(t, err)
}
