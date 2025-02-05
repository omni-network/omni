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

//nolint:paralleltest // Global docker dir container
func TestIntegration(t *testing.T) {
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()

	cl, privKey0, stop, err := solcompose.Start(ctx, dir)
	require.NoError(t, err)
	defer stop()
	t.Logf("Pubkey: %s", privKey0.PublicKey())

	bal0, err := cl.GetBalance(ctx, privKey0.PublicKey(), rpc.CommitmentConfirmed)
	require.NoError(t, err)
	t.Logf("Balance: %d lamports, %d sol", bal0.Value, bal0.Value/solana.LAMPORTS_PER_SOL)

	height, err := cl.GetBlockHeight(ctx, rpc.CommitmentConfirmed)
	require.NoError(t, err)
	t.Logf("Finalized Block: %d", height)

	slot, err := cl.GetSlot(ctx, rpc.CommitmentConfirmed)
	require.NoError(t, err)
	t.Logf("Processed Slot: %d", slot)

	privKey1, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	bal1, err := cl.GetBalance(ctx, privKey1.PublicKey(), rpc.CommitmentConfirmed)
	require.NoError(t, err)
	t.Logf("Balance: %d lamports, %d sol", bal1.Value, bal1.Value/solana.LAMPORTS_PER_SOL)

	txSig, err := cl.RequestAirdrop(ctx, privKey1.PublicKey(), solana.LAMPORTS_PER_SOL, rpc.CommitmentConfirmed)
	require.NoError(t, err)

	var v0 uint64
	tx, err := cl.GetTransaction(ctx, txSig, &rpc.GetTransactionOpts{
		Encoding:                       solana.EncodingBase64,
		Commitment:                     rpc.CommitmentConfirmed,
		MaxSupportedTransactionVersion: &v0,
	})
	require.NoError(t, err)
	t.Logf("Airdrop Tx: slot=%d, time=%v, sig=%v", tx.Slot, tx.BlockTime, txSig)

	slot, err = cl.GetSlot(ctx, rpc.CommitmentConfirmed)
	require.NoError(t, err)
	t.Logf("Finalized Slot: %d", slot)

	bal1, err = cl.GetBalance(ctx, privKey1.PublicKey(), rpc.CommitmentConfirmed)
	require.NoError(t, err)
	t.Logf("Balance: %d lamports, %d sol", bal1.Value, bal1.Value/solana.LAMPORTS_PER_SOL)
	time.Sleep(time.Minute * 10)
}
