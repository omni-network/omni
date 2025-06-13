//nolint:paralleltest // Global docker dir container
package svmutil_test

import (
	"context"
	"encoding/json"
	"flag"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/anchor/anchorinbox"
	"github.com/omni-network/omni/anchor/localnet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/svmutil"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/umath"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/memo"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", true, "Include integration tests")

// dir is subdirectory to store the docker compose file and solana generated artifacts (ignored from repo).
const dir = "compose"

func TestIntegration(t *testing.T) {
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()

	cl, _, privKey0, stop, err := svmutil.Start(ctx, dir)
	if err != nil {
		t.Skip("Skip if docker unhealthy")
	}
	defer stop()
	t.Logf("Pubkey: %s", privKey0.PublicKey())

	id, err := svmutil.ChainID(ctx, cl)
	require.NoError(t, err)
	require.Equal(t, evmchain.IDSolanaLocal, id)

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

		tx, err := svmutil.AwaitConfirmedTransaction(ctx, cl, txSig)
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

		txResp, err := svmutil.AwaitConfirmedTransaction(t.Context(), cl, txSig)
		require.NoError(t, err)
		t.Logf("Memo Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp.Slot, txResp.BlockTime, txSig, txResp.Meta.LogMessages)

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

func TestUpgrade(t *testing.T) {
	t.Skip("Skipping upgrade test as this is very slow") // Uncomment to test manually.

	// Start svm
	ctx := t.Context()
	cl, rpcAddr, privkey, stop, err := svmutil.Start(ctx, dir)
	require.NoError(t, err)
	defer stop()

	// Create two new accounts to deploy and upgrade programs.
	deployer, err := fundNewAccount(ctx, cl)
	require.NoError(t, err)
	upgrader, err := fundNewAccount(ctx, cl)
	require.NoError(t, err)

	// dummyProgram is a valid "empty" program; it is result of `anchor init`.
	dummyProgram := svmutil.Program{
		Name:         "dummy",
		SharedObject: localnet.DummySO,
		KeyPairJSON:  localnet.DummyKeyPairJSON,
	}

	// Deploy the dummy program
	_, err = svmutil.Deploy(ctx, rpcAddr, dummyProgram, deployer, upgrader)
	require.NoError(t, err)

	anchorDir, err := filepath.Abs("../../anchor")
	require.NoError(t, err)

	// Rebuild anchor inbox with dummy key pair
	prog, err := svmutil.Rebuild(ctx, anchorinbox.Program(), dummyProgram.MustPrivateKey(), anchorDir)
	require.NoError(t, err)
	anchorinbox.SetProgramID(prog.MustPublicKey())

	// Redeploy/upgrade dummy program to anchorinbox
	_, err = svmutil.Redeploy(ctx, rpcAddr, prog, upgrader)
	require.NoError(t, err)

	// Init the anchor program
	init, err := anchorinbox.NewInit(0, 0, privkey.PublicKey())
	require.NoError(t, err)

	txSig, err := svmutil.SendSimple(ctx, cl, privkey, init.Build())
	require.NoError(t, err)

	_, err = svmutil.AwaitConfirmedTransaction(ctx, cl, txSig)
	require.NoError(t, err)
}

func fundNewAccount(ctx context.Context, cl *rpc.Client) (solana.PrivateKey, error) {
	key, err := solana.NewRandomPrivateKey()
	if err != nil {
		return nil, err
	}

	txSig, err := cl.RequestAirdrop(ctx, key.PublicKey(), 1e6*solana.LAMPORTS_PER_SOL, rpc.CommitmentConfirmed)
	if err != nil {
		return nil, err
	}

	_, err = svmutil.AwaitConfirmedTransaction(ctx, cl, txSig)
	if err != nil {
		return nil, err
	}

	log.Debug(ctx, "Funded new account", "address", key.PublicKey())

	return key, nil
}

func TestInbox(t *testing.T) {
	if !*integration {
		t.Skip("skipping integration test")
	}

	prog := anchorinbox.Program()

	ctx := t.Context()
	cl, rpcAddr, privKey0, stop, err := svmutil.Start(ctx, dir)
	if err != nil {
		t.Skip("Skip if docker unhealthy")
	}
	defer stop()

	random, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	// Deploy events program
	t.Run("deploy", func(t *testing.T) {
		tx0, err := svmutil.Deploy(ctx, rpcAddr, prog, privKey0, random)
		tutil.RequireNoError(t, err)
		t.Logf("Deployed inbox program: slot=%d, account=%s", tx0.Slot, prog.MustPublicKey())
	})

	inboxStateAddr, bump, err := anchorinbox.FindInboxStateAddress()
	require.NoError(t, err)
	chainID, err := svmutil.ChainID(ctx, cl)
	require.NoError(t, err)

	const closeBuffer = 0

	var initSig solana.Signature
	t.Run("init", func(t *testing.T) {
		// Initialize inbox state
		init, err := anchorinbox.NewInit(chainID, closeBuffer, privKey0.PublicKey())
		require.NoError(t, err)
		txSig0, err := svmutil.SendSimple(ctx, cl, privKey0, init.Build())
		require.NoError(t, err)
		txResp0, err := svmutil.AwaitConfirmedTransaction(ctx, cl, txSig0)
		require.NoError(t, err)
		t.Logf("Init Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp0.Slot, txResp0.BlockTime, txSig0, txResp0.Meta.LogMessages)

		// Get InboxState account
		inboxState, ok, err := anchorinbox.GetInboxState(ctx, cl)
		require.NoError(t, err)
		require.True(t, ok)
		// Ensure inbox state is as expected
		expInboxState := anchorinbox.InboxStateAccount{
			Admin:           privKey0.PublicKey(),
			DeployedAt:      txResp0.Slot,
			Bump:            bump,
			ChainId:         chainID,
			CloseBufferSecs: closeBuffer,
		}
		require.Equal(t, expInboxState, inboxState)

		initSig, err = anchorinbox.GetInitSig(ctx, cl)
		require.NoError(t, err)
		require.Equal(t, txSig0, initSig)
	})

	// Start streaming program tx sigs async
	async := make(chan solana.Signature, 1000)
	go func() {
		streamReq := svmutil.StreamReq{
			AfterSig:   initSig,
			Account:    prog.MustPublicKey(),
			Commitment: rpc.CommitmentConfirmed,
		}
		err := svmutil.Stream(ctx, cl, streamReq, func(ctx context.Context, sig *rpc.TransactionSignature) error {
			t.Logf("Streamed Tx: slot=%d, sig=%v", sig.Slot, sig.Signature)
			async <- sig.Signature

			return nil
		})
		if err != nil {
			t.Errorf("stream error: %v", errors.Format(err))
		}
		close(async) // Signal end of stream
	}()

	owner := privKey0.PublicKey()
	depositAmount := uint64(1e3) // 1K tokens

	claimer, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	var mintResp svmutil.CreateMintResp
	var claimerATA solana.PublicKey
	t.Run("mint", func(t *testing.T) {
		// Create mint and depositor token account
		mintResp, err = svmutil.CreateMint(ctx, cl, privKey0, svmutil.DevnetUSDCMint, 0, claimer.PublicKey(), privKey0.PublicKey())
		require.NoError(t, err)
		claimerATA = mintResp.ATAs[claimer.PublicKey()]

		// Airdrop 1 SOL to claimer (to pay for claim)
		txSig, err := cl.RequestAirdrop(ctx, claimer.PublicKey(), solana.LAMPORTS_PER_SOL, rpc.CommitmentConfirmed)
		require.NoError(t, err)
		_, err = svmutil.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.NoError(t, err)
	})

	// Prep Open instruction
	var openOrder anchorinbox.OpenOrder
	var orderFillHash solana.PublicKey
	t.Run("open", func(t *testing.T) {
		var p anchorinbox.OpenParams
		fuzz.New().Funcs(func(u *bin.Uint128, c fuzz.Continue) {
			*u = svmutil.RandomU96() // Limit u128s to u96 (since that is max expense value)
		}).Fuzz(&p)
		p.DepositAmount = depositAmount
		p.Call.Params = tutil.RandomBytes(4)

		openOrder, err = anchorinbox.NewOpenOrder(p, owner, mintResp.MintAccount)
		require.NoError(t, err)

		// Send Open instruction
		txSig1, err := svmutil.SendSimple(ctx, cl, privKey0, openOrder.Build())
		require.NoError(t, err)

		txResp1, err := svmutil.AwaitConfirmedTransaction(ctx, cl, txSig1)
		t.Logf("Open Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp1.Slot, txResp1.BlockTime, txSig1, txResp1.Meta.LogMessages)
		require.NoError(t, err)
		assertStreamedTx(t, async, txResp1)

		// Ensure updated event
		event := anchorinbox.EventUpdated{
			OrderId: openOrder.ID,
			Status:  anchorinbox.StatusPending,
		}
		ensureInboxEvent(t, prog, txResp1, anchorinbox.EventNameUpdated, event)

		// Get OrderState account
		var orderState2 anchorinbox.OrderStateAccount
		_, err = svmutil.GetAccountDataInto(ctx, cl, openOrder.StateAddress, &orderState2)
		require.NoError(t, err)

		orderState, ok, err := anchorinbox.GetOrderState(ctx, cl, openOrder.ID)
		require.NoError(t, err)
		require.True(t, ok)

		// Ensure OrderState account is correct
		expOrderState := anchorinbox.OrderStateAccount{
			OrderId:       openOrder.ID,
			Status:        anchorinbox.StatusPending,
			Owner:         privKey0.PublicKey(),
			Bump:          openOrder.StateBump,
			CreatedAt:     txResp1.BlockTime.Time().Unix(),
			ClosableAt:    txResp1.BlockTime.Time().Unix() + closeBuffer,
			DepositAmount: depositAmount,
			DepositMint:   mintResp.MintAccount,
			DestChainId:   openOrder.Params.DestChainId,
			DestCall:      openOrder.Params.Call,
			DestExpense:   openOrder.Params.Expense,
			FillHash:      fillHash(t, chainID, openOrder.Params, txResp1.BlockTime.Time().Unix()+closeBuffer),
		}
		require.Equal(t, expOrderState, orderState)
		orderFillHash = orderState.FillHash

		// Ensure deposit amount transferred to order token account
		bal, err := cl.GetTokenAccountBalance(ctx, openOrder.TokenAddress, rpc.CommitmentConfirmed)
		require.NoError(t, err)
		require.Equal(t, depositAmount, uint64(*bal.Value.UiAmount))

		bal, err = cl.GetTokenAccountBalance(ctx, mintResp.AuthATA(), rpc.CommitmentConfirmed)
		require.NoError(t, err)
		require.Equal(t, 1e9-int64(depositAmount), int64(*bal.Value.UiAmount))
	})

	// Send MarkFilled instruction
	t.Run("mark filled", func(t *testing.T) {
		markFilled, err := anchorinbox.NewMarkFilledOrder(ctx, cl, claimer.PublicKey(), privKey0.PublicKey(), openOrder.ID)
		require.NoError(t, err)
		txSig2, err := svmutil.SendSimple(ctx, cl, privKey0, markFilled.Build())
		require.NoError(t, err)
		txResp2, err := svmutil.AwaitConfirmedTransaction(ctx, cl, txSig2)
		t.Logf("MarkFilled Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp2.Slot, txResp2.BlockTime, txSig2, txResp2.Meta.LogMessages)
		customErr := anchorinbox.DecodeMetaError(txResp2)
		require.NoError(t, customErr)
		require.NoError(t, err)
		assertStreamedTx(t, async, txResp2)

		// Ensure updated event
		event := anchorinbox.EventUpdated{
			OrderId: openOrder.ID,
			Status:  anchorinbox.StatusFilled,
		}
		ensureInboxEvent(t, prog, txResp2, anchorinbox.EventNameUpdated, event)

		orderState, ok, err := anchorinbox.GetOrderState(ctx, cl, openOrder.ID)
		require.NoError(t, err)
		require.True(t, ok)

		// Ensure OrderState account is correct
		expOrderState := anchorinbox.OrderStateAccount{
			OrderId:       openOrder.ID,
			Status:        anchorinbox.StatusFilled,
			Owner:         privKey0.PublicKey(),
			Bump:          openOrder.StateBump,
			CreatedAt:     orderState.CreatedAt,
			ClosableAt:    orderState.ClosableAt,
			DepositAmount: depositAmount,
			DepositMint:   mintResp.MintAccount,
			ClaimableBy:   claimer.PublicKey(),
			DestChainId:   openOrder.Params.DestChainId,
			DestCall:      openOrder.Params.Call,
			DestExpense:   openOrder.Params.Expense,
			FillHash:      orderFillHash,
		}
		require.Equal(t, expOrderState, orderState)
	})

	t.Run("claim", func(t *testing.T) {
		claim, err := anchorinbox.NewClaimOrder(ctx, cl, claimer.PublicKey(), openOrder.ID)
		require.NoError(t, err)
		txSig3, err := svmutil.SendSimple(ctx, cl, claimer, claim.Build())
		require.NoError(t, err)

		txResp3, err := svmutil.AwaitConfirmedTransaction(ctx, cl, txSig3)
		require.NoError(t, err)
		t.Logf("Claim Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp3.Slot, txResp3.BlockTime, txSig3, txResp3.Meta.LogMessages)
		assertStreamedTx(t, async, txResp3)

		bal, err := cl.GetTokenAccountBalance(ctx, claimerATA, rpc.CommitmentConfirmed)
		require.NoError(t, err)
		const mintAmount = 1e9
		require.Equal(t, depositAmount+mintAmount, uint64(*bal.Value.UiAmount))

		_, err = cl.GetTokenAccountBalance(ctx, openOrder.TokenAddress, rpc.CommitmentConfirmed)
		require.Contains(t, errors.Format(svmutil.WrapRPCError(err, "getTokenAccountBalance")), "could not find account")
	})

	// Prep Open instruction
	t.Run("open and close", func(t *testing.T) {
		openOrder, err := anchorinbox.NewOpenOrder(anchorinbox.OpenParams{DepositAmount: depositAmount}, owner, mintResp.MintAccount)
		require.NoError(t, err)
		closeOrder := anchorinbox.NewCloseInstruction(
			openOrder.ID,
			openOrder.StateAddress,
			openOrder.TokenAddress,
			mintResp.AuthATA(),
			owner,
			token.ProgramID,
		)
		// Send Open instruction
		txSig, err := svmutil.SendSimple(ctx, cl, privKey0, openOrder.Build(), closeOrder.Build())
		require.NoError(t, err)

		txResp, err := svmutil.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.NoError(t, err)
		t.Logf("Open and Close Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp.Slot, txResp.BlockTime, txSig, txResp.Meta.LogMessages)
		assertStreamedTx(t, async, txResp)

		orderState, ok, err := anchorinbox.GetOrderState(ctx, cl, openOrder.ID)
		require.NoError(t, err)
		require.True(t, ok)

		// Ensure OrderState account is correct
		expOrderState := anchorinbox.OrderStateAccount{
			OrderId:       openOrder.ID,
			Status:        anchorinbox.StatusClosed,
			Owner:         privKey0.PublicKey(),
			Bump:          openOrder.StateBump,
			CreatedAt:     orderState.CreatedAt,
			ClosableAt:    orderState.ClosableAt,
			DepositAmount: depositAmount,
			DepositMint:   mintResp.MintAccount,
			FillHash:      fillHash(t, chainID, openOrder.Params, orderState.ClosableAt),
		}
		require.Equal(t, expOrderState, orderState)

		// Ensure deposit amount transferred to back to owner
		_, err = cl.GetTokenAccountBalance(ctx, openOrder.TokenAddress, rpc.CommitmentConfirmed)
		require.Contains(t, errors.Format(svmutil.WrapRPCError(err, "getTokenAccountBalance")), "could not find account")

		bal, err := cl.GetTokenAccountBalance(ctx, mintResp.AuthATA(), rpc.CommitmentConfirmed)
		require.NoError(t, err)
		require.Equal(t, 1e9-int64(depositAmount), int64(*bal.Value.UiAmount))
	})

	t.Run("open and reject", func(t *testing.T) {
		const reason uint8 = 99
		openOrder, err := anchorinbox.NewOpenOrder(anchorinbox.OpenParams{DepositAmount: depositAmount}, owner, mintResp.MintAccount)
		require.NoError(t, err)

		// Send open instructions
		txSig, err := svmutil.SendSimple(ctx, cl, privKey0, openOrder.Build())
		require.NoError(t, err)

		txResp, err := svmutil.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.NoError(t, err)
		t.Logf("Open Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp.Slot, txResp.BlockTime, txSig, txResp.Meta.LogMessages)
		assertStreamedTx(t, async, txResp)

		rejectOrder, err := anchorinbox.NewRejectOrder(ctx, cl, privKey0.PublicKey(), openOrder.ID, reason)
		require.NoError(t, err)

		// Send Reject instruction
		txSig, err = svmutil.SendSimple(ctx, cl, privKey0, rejectOrder.Build())
		require.NoError(t, err)
		txResp, err = svmutil.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.NoError(t, err)
		t.Logf("Reject Tx: slot=%d, time=%v, sig=%v, logs=%#v", txResp.Slot, txResp.BlockTime, txSig, txResp.Meta.LogMessages)
		assertStreamedTx(t, async, txResp)

		// Ensure token account closed
		_, err = cl.GetTokenAccountBalance(ctx, openOrder.TokenAddress, rpc.CommitmentConfirmed)
		require.Contains(t, errors.Format(svmutil.WrapRPCError(err, "getTokenAccountBalance")), "could not find account")

		// Ensure deposit amount transferred back to owner
		bal, err := cl.GetTokenAccountBalance(ctx, mintResp.AuthATA(), rpc.CommitmentConfirmed)
		require.NoError(t, err)
		require.Equal(t, 1e9-int64(depositAmount), int64(*bal.Value.UiAmount))

		// Get OrderState account
		orderState, ok, err := anchorinbox.GetOrderState(ctx, cl, openOrder.ID)
		require.NoError(t, err)
		require.True(t, ok)

		// Ensure OrderState account is correct
		expOrderState := anchorinbox.OrderStateAccount{
			OrderId:       openOrder.ID,
			Status:        anchorinbox.StatusRejected,
			Owner:         privKey0.PublicKey(),
			Bump:          openOrder.StateBump,
			CreatedAt:     orderState.CreatedAt,
			ClosableAt:    orderState.ClosableAt,
			DepositAmount: depositAmount,
			DepositMint:   mintResp.MintAccount,
			FillHash:      fillHash(t, chainID, openOrder.Params, orderState.ClosableAt),
			RejectReason:  reason,
		}
		require.Equal(t, expOrderState, orderState)
	})

	t.Run("init fail", func(t *testing.T) {
		init, err := anchorinbox.NewInit(chainID, closeBuffer, privKey0.PublicKey())
		require.NoError(t, err)
		txSig, err := svmutil.SendSimple(ctx, cl, privKey0, init.Build())
		require.NoError(t, err)
		txResp, err := svmutil.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.ErrorContains(t, err, "transaction failed")
		assertStreamedTx(t, async, txResp)
	})

	t.Run("open fail", func(t *testing.T) {
		params := anchorinbox.OpenParams{OrderId: openOrder.ID, Nonce: openOrder.Params.Nonce}
		open := anchorinbox.NewOpenInstruction(params, openOrder.StateAddress, privKey0.PublicKey(), mintResp.MintAccount, mintResp.AuthATA(), openOrder.TokenAddress, token.ProgramID, inboxStateAddr, solana.SystemProgramID)
		txSig, err := svmutil.SendSimple(ctx, cl, privKey0, open.Build())
		require.NoError(t, err)
		txResp, err := svmutil.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.ErrorContains(t, err, "transaction failed")
		assertStreamedTx(t, async, txResp)
	})

	t.Run("mark filled fail", func(t *testing.T) {
		markFilled := anchorinbox.NewMarkFilledInstruction(openOrder.ID, claimer.PublicKey(), orderFillHash, openOrder.StateAddress, inboxStateAddr, privKey0.PublicKey())
		txSig, err := svmutil.SendSimple(ctx, cl, privKey0, markFilled.Build())
		require.NoError(t, err)
		txResp, err := svmutil.AwaitConfirmedTransaction(ctx, cl, txSig)
		require.ErrorContains(t, err, "transaction failed")
		assertStreamedTx(t, async, txResp)
	})
}

func assertStreamedTx(t *testing.T, async <-chan solana.Signature, txResp *rpc.GetTransactionResult) {
	t.Helper()
	require.Equal(t, mustFirstTxSig(txResp).String(), (<-async).String())
}

func fillHash(t *testing.T, chainID uint64, params *anchorinbox.OpenParams, closableAt int64) solana.PublicKey {
	t.Helper()

	closableAtU32, err := umath.ToUint32(closableAt)
	require.NoError(t, err)

	hash, err := svmutil.FillHash(
		params.OrderId,
		chainID,
		params.DestChainId,
		closableAtU32,
		params.Call.Target,
		params.Call.Selector,
		params.Call.Value.BigInt(),
		params.Call.Params,
		params.Expense.Spender,
		params.Expense.Token,
		params.Expense.Amount.BigInt(),
	)
	require.NoError(t, err)

	return solana.PublicKey(hash)
}

func TestChainIDs(t *testing.T) {
	ctx := t.Context()

	cl := rpc.New("https://api.mainnet-beta.solana.com")
	id, err := svmutil.ChainID(ctx, cl)
	if err != nil {
		log.Error(ctx, "get chain ID", err)
		t.Skip("Skip 3rd party network issues")
	}
	require.Equal(t, evmchain.IDSolana, id)

	cl = rpc.New("https://api.devnet.solana.com")
	id, err = svmutil.ChainID(ctx, cl)
	if err != nil {
		log.Error(ctx, "get chain ID", err)
		t.Skip("Skip 3rd party network issues")
	}
	require.Equal(t, evmchain.IDSolanaDevnet, id)
}

func ensureInboxEvent(t *testing.T, prog svmutil.Program, txRes *rpc.GetTransactionResult, expectName string, expectData any) {
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
