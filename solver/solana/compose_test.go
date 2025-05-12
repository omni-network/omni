//nolint:paralleltest // Global docker dir container
package solana_test

import (
	"context"
	"encoding/json"
	"flag"
	"math/rand/v2"
	"sync"
	"testing"
	"time"

	"github.com/omni-network/omni/anchor/anchorinbox"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tutil"
	solcompose "github.com/omni-network/omni/solver/solana"
	"github.com/omni-network/omni/solver/solana/events"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	associatedtokenaccount "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/memo"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
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

	// Start streaming program tx sigs async
	async := make(chan solana.Signature, 1000)
	go func() {
		streamReq := solcompose.StreamReq{
			FromSlot:   ptr(uint64(0)),
			Account:    prog.MustPublicKey(),
			Commitment: rpc.CommitmentConfirmed,
		}
		err := solcompose.Stream(ctx, cl, streamReq, func(ctx context.Context, sig *rpc.TransactionSignature) error {
			t.Logf("Streamed Tx: slot=%d, sig=%v", sig.Slot, sig.Signature)
			async <- sig.Signature

			return nil
		})
		if err != nil {
			t.Errorf("stream error: %v", err)
		}
	}()

	// Deploy events program
	t.Run("deploy", func(t *testing.T) {
		tx0, err := solcompose.Deploy(ctx, cl, dir, prog)
		tutil.RequireNoError(t, err)
		t.Logf("Deployed events program: slot=%d, account=%s", tx0.Slot, prog.MustPublicKey())
		require.Equal(t, mustFirstTxSig(tx0), <-async)

		// Sent Init instruction
		txSig1, err := solcompose.SendSimple(ctx, cl, privKey0, events.NewInitializeInstruction().Build())
		require.NoError(t, err)

		txResp1, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig1)
		require.NoError(t, err)
		t.Logf("Init Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp1.Slot, txResp1.BlockTime, txSig1, txResp1.Meta.LogMessages)

		ensureEvent(t, prog, txResp1, events.EventMyEvent, events.MyEventEventData{Data: 5, Label: "hello"})
		require.Equal(t, mustFirstTxSig(txResp1), <-async)
	})

	t.Run("send event", func(t *testing.T) {
		// Send TestEvent instruction
		txSig2, err := solcompose.SendSimple(ctx, cl, privKey0, events.NewTestEventInstruction().Build())
		require.NoError(t, err)

		txResp2, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig2)
		require.NoError(t, err)
		t.Logf("TestEvent Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp2.Slot, txResp2.BlockTime, txSig2, txResp2.Meta.LogMessages)

		ensureEvent(t, prog, txResp2, events.EventMyOtherEvent, events.MyOtherEventEventData{Data: 6, Label: "bye"})
		require.Equal(t, mustFirstTxSig(txResp2), <-async)
	})

	t.Run("send multi-concurrent", func(t *testing.T) {
		// Send N async txs
		const N = 16
		var sigs sync.Map
		for i := range N {
			go func() {
				time.Sleep(time.Millisecond * time.Duration(rand.IntN(2000))) // Delay up to 2s
				txSig, err := solcompose.SendSimple(ctx, cl, privKey0,
					events.NewTestEventInstruction().Build(),
					memo.NewMemoInstruction([]byte{byte(i)}, privKey0.PublicKey()).Build(), // Add uniqueness to tx
				)
				if err != nil {
					t.Error("error sending tx:", err)
				}
				t.Logf("Async sent tx: %s", txSig)
				sigs.Store(txSig, true)
			}()
		}

		for range N {
			txSig := <-async
			require.NotNil(t, txSig)
			require.Eventuallyf(t, func() bool {
				_, ok := sigs.LoadAndDelete(txSig)
				return ok
			}, time.Second*10, time.Second, "tx sig not found in map: %s", txSig)
		}

		sigs.Range(func(k, _ any) bool {
			require.Fail(t, "tx sig map not empty")
			return true
		})
	})
}

func TestInbox(t *testing.T) {
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()
	cl, privKey0, stop, err := solcompose.Start(ctx, dir)
	require.NoError(t, err)
	defer stop()

	prog := solcompose.ProgramInbox

	// Start streaming program tx sigs async
	async := make(chan solana.Signature, 1000)
	go func() {
		streamReq := solcompose.StreamReq{
			FromSlot:   ptr(uint64(0)),
			Account:    prog.MustPublicKey(),
			Commitment: rpc.CommitmentConfirmed,
		}
		err := solcompose.Stream(ctx, cl, streamReq, func(ctx context.Context, sig *rpc.TransactionSignature) error {
			t.Logf("Streamed Tx: slot=%d, sig=%v", sig.Slot, sig.Signature)
			async <- sig.Signature

			return nil
		})
		if err != nil {
			t.Errorf("stream error: %v", err)
		}
	}()

	// Deploy events program
	t.Run("deploy", func(t *testing.T) {
		tx0, err := solcompose.Deploy(ctx, cl, dir, prog)
		tutil.RequireNoError(t, err)
		t.Logf("Deployed inbox program: slot=%d, account=%s", tx0.Slot, prog.MustPublicKey())
		require.Equal(t, mustFirstTxSig(tx0), <-async)
	})

	inboxStateAddr, bump, err := anchorinbox.FindInboxStateAddress()
	require.NoError(t, err)

	const closeBuffer = 0

	t.Run("init", func(t *testing.T) {
		// Initialize inbox state
		init := anchorinbox.NewInitInstruction(closeBuffer, inboxStateAddr, privKey0.PublicKey(), solana.SystemProgramID)
		txSig0, err := solcompose.SendSimple(ctx, cl, privKey0, init.Build())
		require.NoError(t, err)
		txResp0, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig0)
		require.NoError(t, err)
		t.Logf("Init Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp0.Slot, txResp0.BlockTime, txSig0, txResp0.Meta.LogMessages)
		require.Equal(t, mustFirstTxSig(txResp0), <-async)

		// Get InboxState account
		info, err := cl.GetAccountInfoWithOpts(ctx, inboxStateAddr, &rpc.GetAccountInfoOpts{Commitment: rpc.CommitmentConfirmed})
		require.NoError(t, err)
		inboxState := anchorinbox.InboxStateAccount{}
		err = bin.NewBinDecoder(info.Value.Data.GetBinary()).Decode(&inboxState)
		require.NoError(t, err)
		// Ensure inbox state is as expected
		expInboxState := anchorinbox.InboxStateAccount{
			Admin:           privKey0.PublicKey(),
			DeployedAt:      txResp0.Slot,
			Bump:            bump,
			CloseBufferSecs: closeBuffer,
		}
		require.Equal(t, expInboxState, inboxState)
	})

	owner := privKey0.PublicKey()
	depositAmount := uint64(1e3) // 1K tokens

	claimer, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	var mintResp createMintResp
	var claimerATA solana.PublicKey
	t.Run("mint", func(t *testing.T) {
		// Create mint and depositor token account
		mintResp = createMint(ctx, t, cl, privKey0)

		// Airdrop 1 SOL to claimer (to pay for claim)
		txSig, err := cl.RequestAirdrop(ctx, claimer.PublicKey(), solana.LAMPORTS_PER_SOL, rpc.CommitmentConfirmed)
		require.NoError(t, err)
		_, err = solcompose.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.NoError(t, err)

		claimerATA = ensureATA(ctx, t, cl, claimer, mintResp.MintAccount)
	})

	// Prep Open instruction
	var openOrder anchorinbox.OpenOrder
	t.Run("open", func(t *testing.T) {
		openOrder, err = anchorinbox.NewOpenOrder(anchorinbox.OpenParams{
			DepositAmount: depositAmount,
		}, owner, mintResp.MintAccount, mintResp.TokenAccount)
		require.NoError(t, err)

		// Send Open instruction
		txSig1, err := solcompose.SendSimple(ctx, cl, privKey0, openOrder.Build())
		require.NoError(t, err)

		txResp1, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig1)
		require.NoError(t, err)
		t.Logf("Open Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp1.Slot, txResp1.BlockTime, txSig1, txResp1.Meta.LogMessages)
		require.Equal(t, mustFirstTxSig(txResp1), <-async)

		// Ensure Opened event
		openEvent := anchorinbox.EventOpened{
			OrderId:    openOrder.ID,
			OrderState: openOrder.StateAddress,
			Status:     anchorinbox.StatusPending,
		}
		ensureInboxEvent(t, prog, txResp1, anchorinbox.EventNameOpened, openEvent)

		// Get OrderState account
		var orderState anchorinbox.OrderStateAccount
		_, err = solcompose.GetAccountDataInto(ctx, cl, openOrder.StateAddress, &orderState)
		require.NoError(t, err)

		// Ensure OrderState account is correct
		expOrderState := anchorinbox.OrderStateAccount{
			OrderId:    openOrder.ID,
			Status:     anchorinbox.StatusPending,
			Owner:      privKey0.PublicKey(),
			Bump:       openOrder.StateBump,
			CreatedAt:  txResp1.BlockTime.Time().Unix(),
			ClosableAt: txResp1.BlockTime.Time().Unix() + closeBuffer,
			Deposit: anchorinbox.TokenAmount{
				Mint:   mintResp.MintAccount,
				Amount: depositAmount,
			},
		}
		require.Equal(t, expOrderState, orderState)

		// Ensure deposit amount transferred to order token account
		bal, err := cl.GetTokenAccountBalance(ctx, openOrder.TokenAddress, rpc.CommitmentConfirmed)
		require.NoError(t, err)
		require.Equal(t, depositAmount, uint64(*bal.Value.UiAmount))

		bal, err = cl.GetTokenAccountBalance(ctx, mintResp.TokenAccount, rpc.CommitmentConfirmed)
		require.NoError(t, err)
		require.Equal(t, 1e9-int64(depositAmount), int64(*bal.Value.UiAmount))
	})

	// Send MarkFilled instruction
	t.Run("mark filled", func(t *testing.T) {
		markFilled := anchorinbox.NewMarkFilledInstruction(openOrder.ID, claimer.PublicKey(), openOrder.StateAddress, inboxStateAddr, privKey0.PublicKey())
		txSig2, err := solcompose.SendSimple(ctx, cl, privKey0, markFilled.Build())
		require.NoError(t, err)
		txResp2, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig2)
		require.NoError(t, err)
		t.Logf("MarkFilled Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp2.Slot, txResp2.BlockTime, txSig2, txResp2.Meta.LogMessages)
		require.Equal(t, mustFirstTxSig(txResp2), <-async)

		// Ensure MarkFilled event
		markFilledEvent := anchorinbox.EventMarkFilled{
			OrderId:    openOrder.ID,
			OrderState: openOrder.StateAddress,
			Status:     anchorinbox.StatusFilled,
		}
		ensureInboxEvent(t, prog, txResp2, anchorinbox.EventNameMarkFilled, markFilledEvent)

		var orderState anchorinbox.OrderStateAccount
		_, err = solcompose.GetAccountDataInto(ctx, cl, openOrder.StateAddress, &orderState)
		require.NoError(t, err)

		// Ensure OrderState account is correct
		expOrderState := anchorinbox.OrderStateAccount{
			OrderId:    openOrder.ID,
			Status:     anchorinbox.StatusFilled,
			Owner:      privKey0.PublicKey(),
			Bump:       openOrder.StateBump,
			CreatedAt:  orderState.CreatedAt,
			ClosableAt: orderState.ClosableAt,
			Deposit: anchorinbox.TokenAmount{
				Mint:   mintResp.MintAccount,
				Amount: depositAmount,
			},
			ClaimableBy: claimer.PublicKey(),
		}
		require.Equal(t, expOrderState, orderState)
	})

	t.Run("claim", func(t *testing.T) {
		claim := anchorinbox.NewClaimInstruction(
			openOrder.ID,
			openOrder.StateAddress,
			openOrder.TokenAddress,
			claimer.PublicKey(),
			claimerATA,
			token.ProgramID,
		)
		txSig3, err := solcompose.SendSimple(ctx, cl, claimer, claim.Build())
		require.NoError(t, err)

		txResp3, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig3)
		require.NoError(t, err)
		t.Logf("Claim Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp3.Slot, txResp3.BlockTime, txSig3, txResp3.Meta.LogMessages)
		require.Equal(t, mustFirstTxSig(txResp3), <-async)

		bal, err := cl.GetTokenAccountBalance(ctx, claimerATA, rpc.CommitmentConfirmed)
		require.NoError(t, err)
		require.Equal(t, depositAmount, uint64(*bal.Value.UiAmount))

		bal, err = cl.GetTokenAccountBalance(ctx, openOrder.TokenAddress, rpc.CommitmentConfirmed)
		require.NoError(t, err)
		require.Equal(t, uint64(0), uint64(*bal.Value.UiAmount))
	})

	// Prep Open instruction
	t.Run("open and close", func(t *testing.T) {
		openOrder, err := anchorinbox.NewOpenOrder(anchorinbox.OpenParams{DepositAmount: depositAmount}, owner, mintResp.MintAccount, mintResp.TokenAccount)
		require.NoError(t, err)
		closeOrder := anchorinbox.NewCloseInstruction(
			openOrder.ID,
			openOrder.StateAddress,
			openOrder.TokenAddress,
			mintResp.TokenAccount,
			owner,
			token.ProgramID,
		)
		// Send Open instruction
		txSig, err := solcompose.SendSimple(ctx, cl, privKey0, openOrder.Build(), closeOrder.Build())
		require.NoError(t, err)

		txResp, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.NoError(t, err)
		t.Logf("Open and Close Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp.Slot, txResp.BlockTime, txSig, txResp.Meta.LogMessages)
		require.Equal(t, mustFirstTxSig(txResp), <-async)

		// Get OrderState account
		var orderState anchorinbox.OrderStateAccount
		_, err = solcompose.GetAccountDataInto(ctx, cl, openOrder.StateAddress, &orderState)
		require.NoError(t, err)

		// Ensure OrderState account is correct
		expOrderState := anchorinbox.OrderStateAccount{
			OrderId:    openOrder.ID,
			Status:     anchorinbox.StatusClosed,
			Owner:      privKey0.PublicKey(),
			Bump:       openOrder.StateBump,
			CreatedAt:  orderState.CreatedAt,
			ClosableAt: orderState.ClosableAt,
			Deposit: anchorinbox.TokenAmount{
				Mint:   mintResp.MintAccount,
				Amount: depositAmount,
			},
		}
		require.Equal(t, expOrderState, orderState)

		// Ensure deposit amount transferred to back to owner
		_, err = cl.GetTokenAccountBalance(ctx, openOrder.TokenAddress, rpc.CommitmentConfirmed)
		require.Contains(t, errors.Format(solcompose.WrapRPCError(err, "getTokenAccountBalance")), "could not find account")

		bal, err := cl.GetTokenAccountBalance(ctx, mintResp.TokenAccount, rpc.CommitmentConfirmed)
		require.NoError(t, err)
		require.Equal(t, 1e9-int64(depositAmount), int64(*bal.Value.UiAmount))
	})

	t.Run("init fail", func(t *testing.T) {
		init := anchorinbox.NewInitInstruction(0, inboxStateAddr, privKey0.PublicKey(), solana.SystemProgramID)
		txSig, err := solcompose.SendSimple(ctx, cl, privKey0, init.Build())
		require.NoError(t, err)
		txResp, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.ErrorContains(t, err, "transaction failed")
		require.Equal(t, mustFirstTxSig(txResp), <-async)
	})

	t.Run("open fail", func(t *testing.T) {
		params := anchorinbox.OpenParams{OrderId: openOrder.ID, Nonce: openOrder.Params.Nonce}
		open := anchorinbox.NewOpenInstruction(params, openOrder.StateAddress, privKey0.PublicKey(), mintResp.MintAccount, mintResp.TokenAccount, openOrder.TokenAddress, token.ProgramID, inboxStateAddr, solana.SystemProgramID)
		txSig, err := solcompose.SendSimple(ctx, cl, privKey0, open.Build())
		require.NoError(t, err)
		txResp, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.ErrorContains(t, err, "transaction failed")
		require.Equal(t, mustFirstTxSig(txResp), <-async)
	})

	t.Run("mark filled fail", func(t *testing.T) {
		markFilled := anchorinbox.NewMarkFilledInstruction(openOrder.ID, claimer.PublicKey(), openOrder.StateAddress, inboxStateAddr, privKey0.PublicKey())
		txSig, err := solcompose.SendSimple(ctx, cl, privKey0, markFilled.Build())
		require.NoError(t, err)
		txResp, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.ErrorContains(t, err, "transaction failed")
		require.Equal(t, mustFirstTxSig(txResp), <-async)
	})
}

func TestCreateMint(t *testing.T) {
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()
	cl, privKey0, stop, err := solcompose.Start(ctx, dir)
	require.NoError(t, err)
	defer stop()

	_ = createMint(ctx, t, cl, privKey0)
}

type createMintResp struct {
	MintAccount  solana.PublicKey
	TokenAccount solana.PublicKey
	Authority    solana.PrivateKey
}

func createMint(ctx context.Context, t *testing.T, cl *rpc.Client, privkey solana.PrivateKey) createMintResp {
	t.Helper()

	// Create new random mint and associated token account
	mintPrivKey, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	tokenAccount, _, err := solana.FindAssociatedTokenAddress(privkey.PublicKey(), mintPrivKey.PublicKey())
	require.NoError(t, err)
	const mintAmount uint64 = 1e9 // 1G tokens

	// Calculate rent
	rent, err := cl.GetMinimumBalanceForRentExemption(ctx, token.MINT_SIZE, rpc.CommitmentConfirmed)
	require.NoError(t, err)

	// Create mint account, and initialize it, create associated token account, and mint 1G tokens
	createAccount := system.NewCreateAccountInstruction(rent, token.MINT_SIZE, solana.TokenProgramID, privkey.PublicKey(), mintPrivKey.PublicKey())
	initMint := token.NewInitializeMint2Instruction(0, privkey.PublicKey(), privkey.PublicKey(), mintPrivKey.PublicKey())
	createATA := associatedtokenaccount.NewCreateInstruction(privkey.PublicKey(), privkey.PublicKey(), mintPrivKey.PublicKey())
	mintTo := token.NewMintToInstruction(mintAmount, mintPrivKey.PublicKey(), tokenAccount, privkey.PublicKey(), nil)

	txSig, err := solcompose.Send(ctx, cl,
		solcompose.WithInstructions(
			createAccount.Build(),
			initMint.Build(),
			createATA.Build(),
			mintTo.Build(),
		),
		solcompose.WithPrivateKeys(privkey, mintPrivKey),
	)
	require.NoError(t, err)

	txResp, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig)
	require.NoError(t, err)
	t.Logf("Create Mint Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp.Slot, txResp.BlockTime, txSig, txResp.Meta.LogMessages)

	bal, err := cl.GetTokenAccountBalance(ctx, tokenAccount, rpc.CommitmentConfirmed)
	require.NoError(t, err)
	require.Equal(t, mintAmount, uint64(*bal.Value.UiAmount))

	return createMintResp{
		MintAccount:  mintPrivKey.PublicKey(),
		TokenAccount: tokenAccount,
		Authority:    privkey,
	}
}

// ensureATA finds or creates an associated token account.
func ensureATA(ctx context.Context, t *testing.T, cl *rpc.Client,
	wallet solana.PrivateKey, mintAccount solana.PublicKey,
) solana.PublicKey {
	t.Helper()
	ata, _, err := solana.FindAssociatedTokenAddress(wallet.PublicKey(), mintAccount)
	require.NoError(t, err)

	_, err = cl.GetAccountInfoWithOpts(ctx, ata, &rpc.GetAccountInfoOpts{Commitment: rpc.CommitmentConfirmed})
	if errors.Is(err, rpc.ErrNotFound) {
		// Create it below
	} else if err != nil {
		require.NoError(t, err)
	} else {
		// Already exists
		return ata
	}

	create := associatedtokenaccount.NewCreateInstruction(wallet.PublicKey(), wallet.PublicKey(), mintAccount)
	txSig, err := solcompose.SendSimple(ctx, cl, wallet, create.Build())
	require.NoError(t, err)

	txResp, err := solcompose.AwaitConfirmedTransaction(ctx, cl, txSig)
	require.NoError(t, err)
	t.Logf("ATA Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp.Slot, txResp.BlockTime, txSig, txResp.Meta.LogMessages)

	return ata
}

func ensureInboxEvent(t *testing.T, prog solcompose.Program, txRes *rpc.GetTransactionResult, expectName string, expectData any) {
	t.Helper()

	evnts, err := anchorinbox.DecodeEvents(txRes, prog.MustPublicKey(), nil)
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

func mustFirstTxSig(txResp *rpc.GetTransactionResult) solana.Signature {
	tx, err := txResp.Transaction.GetTransaction()
	if err != nil {
		panic(err)
	}

	return tx.Signatures[0]
}
