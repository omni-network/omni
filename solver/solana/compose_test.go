//nolint:paralleltest // Global docker dir container
package solana_test

import (
	"encoding/json"
	"flag"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/tutil"
	solcompose "github.com/omni-network/omni/solver/solana"
	"github.com/omni-network/omni/solver/solana/events"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/memo"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "Include integration tests")

// dir is subdirectory to store the docker compose file and solana generated artifacts (ignored from repo).
const dir = "compose"

var v0 uint64

func TestIntegration(t *testing.T) {
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()

	cl, privKey0, stop, err := solcompose.Start(ctx, dir)
	require.NoError(t, err)
	defer stop()
	t.Logf("Pubkey: %s", privKey0.PublicKey())

	t.Run("prefunded 500M SOL", func(t *testing.T) {
		bal0, err := cl.GetBalance(ctx, privKey0.PublicKey(), rpc.CommitmentConfirmed)
		require.NoError(t, err)
		t.Logf("Balance: %d lamports, %d sol", bal0.Value, bal0.Value/solana.LAMPORTS_PER_SOL)

		require.Equal(t, solana.LAMPORTS_PER_SOL*500_000_000, bal0.Value)
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

		tx, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.NoError(t, err)

		t.Logf("Airdrop Tx: slot=%d, time=%v, sig=%v", tx.Slot, tx.BlockTime, txSig)

		bal1, err = cl.GetBalance(ctx, privKey1.PublicKey(), rpc.CommitmentConfirmed)
		require.NoError(t, err)
		t.Logf("Balance: %d lamports, %d sol", bal1.Value, bal1.Value/solana.LAMPORTS_PER_SOL)

		require.Equal(t, airdropVal, bal1.Value)
	})

	t.Run("send memo", func(t *testing.T) {
		// Ensure memo program is deployed.
		memoAccount, err := cl.GetAccountInfo(ctx, memo.ProgramID)
		require.NoError(t, err)
		require.NotNil(t, memoAccount.Value)
		require.True(t, memoAccount.Value.Executable)

		recent, err := cl.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
		require.NoError(t, err)

		msg1 := []byte("Hello, Solana!")
		msg2 := []byte("Hello, Omni!")
		tx, err := solana.NewTransaction(
			[]solana.Instruction{
				memo.NewMemoInstruction(msg1, privKey0.PublicKey()).Build(),
				memo.NewMemoInstruction(msg2, privKey0.PublicKey()).Build(),
			},
			recent.Value.Blockhash,
			solana.TransactionPayer(privKey0.PublicKey()),
		)
		require.NoError(t, err)

		sigs, err := tx.Sign(func(pub solana.PublicKey) *solana.PrivateKey {
			require.Equal(t, privKey0.PublicKey(), pub)
			return &privKey0
		})
		require.NoError(t, err)
		require.Len(t, sigs, 1)

		txSig, err := cl.SendTransactionWithOpts(ctx, tx, rpc.TransactionOpts{
			SkipPreflight: true, // This fails with Memo Program not deployed if false :/
		})
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
			t.Logf("Memo Tx: slot=%d, time=%v, sig=%v, logs=%#v", tx.Slot, tx.BlockTime, txSig, tx.Meta.LogMessages)

			return true
		}, time.Second*10, time.Second)

		// Get above Memo Tx (using sender address), and assert memo contains messages.
		txSigs, err := cl.GetSignaturesForAddressWithOpts(ctx, privKey0.PublicKey(), &rpc.GetSignaturesForAddressOpts{
			Limit:      ptr(1),
			Commitment: rpc.CommitmentConfirmed,
		})
		require.NoError(t, err)
		require.Len(t, txSigs, 1)
		require.Equal(t, txSig, txSigs[0].Signature)
		require.Contains(t, *txSigs[0].Memo, string(msg1))
		require.Contains(t, *txSigs[0].Memo, string(msg2))

		// Getting sigs by memo program address results in same txSig
		memoSigs, err := cl.GetSignaturesForAddressWithOpts(ctx, memo.ProgramID, &rpc.GetSignaturesForAddressOpts{
			Limit:      ptr(10),
			Commitment: rpc.CommitmentConfirmed,
		})
		require.NoError(t, err)
		require.Len(t, memoSigs, 1)
		require.Equal(t, txSig, memoSigs[0].Signature)
		require.Contains(t, *memoSigs[0].Memo, string(msg1))
		require.Contains(t, *memoSigs[0].Memo, string(msg2))
	})
}

func TestEventsProgram(t *testing.T) {
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()
	cl, privKey0, stop, err := solcompose.Start(ctx, dir)
	require.NoError(t, err)
	defer stop()

	prog := solcompose.ProgramEvents

	// Deploy events program
	tx0, err := solcompose.Deploy(ctx, cl, dir, prog)
	tutil.RequireNoError(t, err)
	t.Logf("Deployed events program: slot=%d, account=%s", tx0.Slot, prog.MustPublicKey())

	// Sent Initialize instruction
	txSig1, err := solcompose.SendSimple(ctx, cl, privKey0, events.NewInitializeInstruction().Build())
	require.NoError(t, err)

	txResp1, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig1)
	require.NoError(t, err)
	t.Logf("Initialize Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp1.Slot, txResp1.BlockTime, txSig1, txResp1.Meta.LogMessages)

	ensureEvent(t, prog, txResp1, events.EventMyEvent, events.MyEventEventData{Data: 5, Label: "hello"})

	// Send TestEvent instruction
	txSig2, err := solcompose.SendSimple(ctx, cl, privKey0, events.NewTestEventInstruction().Build())
	require.NoError(t, err)

	txResp2, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig2)
	require.NoError(t, err)
	t.Logf("TestEvent Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp2.Slot, txResp2.BlockTime, txSig2, txResp2.Meta.LogMessages)

	ensureEvent(t, prog, txResp2, events.EventMyOtherEvent, events.MyOtherEventEventData{Data: 6, Label: "bye"})
}

func ensureEvent(t *testing.T, prog solcompose.Program, txRes *rpc.GetTransactionResult, expectName string, expectData any) {
	t.Helper()

	evnts, err := events.DecodeEvents(txRes, prog.MustPublicKey(), nil)
	require.NoError(t, err)
	require.Len(t, evnts, 1)

	for _, evnt := range evnts {
		require.Equal(t, expectName, evnt.Name)

		expectJSON, err := json.MarshalIndent(expectData, "", "  ")
		require.NoError(t, err)
		actualJSON, err := json.MarshalIndent(evnt.Data, "", "  ")
		require.NoError(t, err)

		require.JSONEq(t, string(expectJSON), string(actualJSON))
	}
}

func ptr[A any](a A) *A {
	return &a
}
