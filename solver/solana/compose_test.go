package solana_test

import (
	"context"
	"flag"
	"testing"
	"time"

	solcompose "github.com/omni-network/omni/solver/solana"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "Include integration tests")

// dir is subdirectory to store the docker compose file and solana generated artifacts (excluded from repo).
const dir = "compose"

var v0 uint64

//nolint:paralleltest // Global docker dir container
func TestIntegration(t *testing.T) {
	if !*integration {
		// t.Skip("skipping integration test")
	}

	ctx := context.Background()

	cl, privKey0, stop, err := solcompose.Start(ctx, dir)
	require.NoError(t, err)
	defer stop()
	t.Logf("Pubkey: %s", privKey0.PublicKey())

	t.Run("prefunded 500M SOL", func(t *testing.T) {
		bal0, err := cl.GetBalance(ctx, privKey0.PublicKey(), rpc.CommitmentConfirmed)
		require.NoError(t, err)
		t.Logf("Balance: %d lamports, %d sol", bal0.Value, bal0.Value/solana.LAMPORTS_PER_SOL)

		require.Equal(t, bal0.Value, solana.LAMPORTS_PER_SOL*500_000_000)
	})

	t.Run("airdrop new key", func(t *testing.T) {
		privKey1, err := solana.NewRandomPrivateKey()
		require.NoError(t, err)

		bal1, err := cl.GetBalance(ctx, privKey1.PublicKey(), rpc.CommitmentConfirmed)
		require.NoError(t, err)
		t.Logf("Balance: %d lamports, %d sol", bal1.Value, bal1.Value/solana.LAMPORTS_PER_SOL)
		require.Zero(t, bal1.Value)

		const airdropVal = solana.LAMPORTS_PER_SOL // 1 SOL
		txSig, err := cl.RequestAirdrop(ctx, privKey1.PublicKey(), airdropVal, rpc.CommitmentConfirmed)
		require.NoError(t, err)

		require.Eventually(t, func() bool {
			tx, err := cl.GetTransaction(ctx, txSig, &rpc.GetTransactionOpts{
				Encoding:                       solana.EncodingBase64,
				Commitment:                     rpc.CommitmentConfirmed,
				MaxSupportedTransactionVersion: &v0,
			})
			if err != nil {
				return false
			}
			t.Logf("Airdrop Tx: slot=%d, time=%v, sig=%v", tx.Slot, tx.BlockTime, txSig)
			return true
		}, time.Second*10, time.Second)

		bal1, err = cl.GetBalance(ctx, privKey1.PublicKey(), rpc.CommitmentConfirmed)
		require.NoError(t, err)
		t.Logf("Balance: %d lamports, %d sol", bal1.Value, bal1.Value/solana.LAMPORTS_PER_SOL)

		require.EqualValues(t, bal1.Value, airdropVal)
	})

}
