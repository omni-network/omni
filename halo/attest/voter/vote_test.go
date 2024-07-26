package voter_test

import (
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

	var attHeader xchain.AttestHeader
	fuzzer.Fuzz(&attHeader)
	attHeader.ChainVersion.ID = block.ChainID            // Align headers
	attHeader.ChainVersion.ConfLevel = xchain.ConfLatest // Pick valid conf level

	att, err := voter.CreateVote(privKey, attHeader, block)
	require.NoError(t, err)
	require.Equal(t, block.BlockHeader, att.BlockHeader.ToXChain())
	require.Equal(t, addr, common.Address(att.Signature.ValidatorAddress))

	// Verify the attestation
	err = att.Verify()
	require.NoError(t, err)
}
