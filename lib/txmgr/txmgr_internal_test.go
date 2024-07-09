package txmgr

import (
	"context"
	"errors"
	"io"
	"math/big"
	"math/rand/v2"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/ethclient"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/require"
)

const (
	startingNonce uint64 = 1 // we pick something other than 0 so we can confirm nonces are getting set properly
)

type sendTransactionFunc func(ctx context.Context, tx *types.Transaction) error

func testSendState() *SendState {
	return NewSendState(100, time.Hour)
}

// testHarness houses the necessary resources to test the simple.
type testHarness struct {
	cfg       Config
	mgr       *simple
	backend   *mockBackend
	gasPricer *gasPricer
}

// newTestHarnessWithConfig initializes a testHarness with a specific
// configuration.
func newTestHarnessWithConfig(_ *testing.T, cfg Config) *testHarness {
	g := newGasPricer(3)
	backend := newMockBackend(g)
	cfg.Backend = backend
	mgr := &simple{
		chainID:   cfg.ChainID,
		chainName: "TEST",
		cfg:       cfg,
		backend:   cfg.Backend,
	}

	return &testHarness{
		cfg:       cfg,
		mgr:       mgr,
		backend:   backend,
		gasPricer: g,
	}
}

// newTestHarness initializes a testHarness with a default configuration that is
// suitable for most tests.
func newTestHarness(t *testing.T) *testHarness {
	t.Helper()
	return newTestHarnessWithConfig(t, configWithNumConfs(1))
}

// createTxCandidate creates a mock [TxCandidate].
func (h testHarness) createTxCandidate() TxCandidate {
	inbox := common.HexToAddress("0x42000000000000000000000000000000000000ff")
	return TxCandidate{
		To:       &inbox,
		TxData:   []byte{0x00, 0x01, 0x02},
		GasLimit: uint64(1337),
	}
}

func configWithNumConfs(numConfirmations uint64) Config {
	return Config{
		ResubmissionTimeout:       time.Second,
		ReceiptQueryInterval:      50 * time.Millisecond,
		NumConfirmations:          numConfirmations,
		SafeAbortNonceTooLowCount: 3,
		FeeLimitMultiplier:        5,
		TxNotInMempoolTimeout:     1 * time.Hour,
		Signer: func(ctx context.Context, from common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
		From: common.Address{},
	}
}

type gasPricer struct {
	epoch         int64
	mineAtEpoch   int64
	baseGasTipFee *big.Int
	baseBaseFee   *big.Int
	excessBlobGas uint64
	err           error
	mu            sync.RWMutex
}

func newGasPricer(mineAtEpoch int64) *gasPricer {
	return &gasPricer{
		mineAtEpoch:   mineAtEpoch,
		baseGasTipFee: big.NewInt(5),
		baseBaseFee:   big.NewInt(7),
		// Simulate 100 excess blobs, which results in a blobBaseFee of 50 wei.  This default means
		// blob txs will be subject to the geth minimum blobgas fee of 1 gwei.
		excessBlobGas: 100 * (params.BlobTxBlobGasPerBlob),
	}
}

func (g *gasPricer) getEpoch() int64 {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.epoch
}

func (g *gasPricer) expGasFeeCap() *big.Int {
	_, gasFeeCap := g.feesForEpoch(g.mineAtEpoch)
	return gasFeeCap
}

func (g *gasPricer) shouldMine(gasFeeCap *big.Int) bool {
	return g.expGasFeeCap().Cmp(gasFeeCap) <= 0
}

func (g *gasPricer) feesForEpoch(epoch int64) (*big.Int, *big.Int) {
	e := big.NewInt(epoch)
	epochBaseFee := new(big.Int).Mul(g.baseBaseFee, e)
	epochGasTipCap := new(big.Int).Mul(g.baseGasTipFee, e)
	epochGasFeeCap := calcGasFeeCap(epochBaseFee, epochGasTipCap)

	return epochGasTipCap, epochGasFeeCap
}

func (g *gasPricer) baseFee() *big.Int {
	return new(big.Int).Mul(g.baseBaseFee, big.NewInt(g.getEpoch()))
}

func (g *gasPricer) sample() (*big.Int, *big.Int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.epoch++
	epochGasTipCap, epochGasFeeCap := g.feesForEpoch(g.epoch)

	return epochGasTipCap, epochGasFeeCap
}

type minedTxInfo struct {
	gasFeeCap   *big.Int
	blobFeeCap  *big.Int
	blockNumber uint64
}

// mockBackend implements ReceiptSource that tracks mined transactions
// along with the gas price used.
type mockBackend struct {
	ethclient.Client
	mu sync.RWMutex

	g    *gasPricer
	send sendTransactionFunc

	// blockHeight tracks the current height of the chain.
	blockHeight uint64

	// minedTxs maps the hash of a mined transaction to its details.
	minedTxs map[common.Hash]minedTxInfo
}

// newMockBackend initializes a new mockBackend.
func newMockBackend(g *gasPricer) *mockBackend {
	return &mockBackend{
		g:        g,
		minedTxs: make(map[common.Hash]minedTxInfo),
	}
}

// setTxSender sets the implementation for the sendTransactionFunction.
func (b *mockBackend) setTxSender(s sendTransactionFunc) {
	b.send = s
}

// mine records a (txHash, gasFeeCap) as confirmed. Subsequent calls to
// TransactionReceipt with a matching txHash will result in a non-nil receipt.
// If a nil txHash is supplied this has the effect of mining an empty block.
func (b *mockBackend) mine(txHash *common.Hash, gasFeeCap *big.Int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.blockHeight++
	if txHash != nil {
		b.minedTxs[*txHash] = minedTxInfo{
			gasFeeCap:   gasFeeCap,
			blockNumber: b.blockHeight,
		}
	}
}

// BlockNumber returns the most recent block number.
func (b *mockBackend) BlockNumber(ctx context.Context) (uint64, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.blockHeight, nil
}

// SyncProgress returns the nil syncing progress; i.e. not syncing.
func (b *mockBackend) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return nil, nil //nolint:nilnil // This is how geth does it.
}

func (b *mockBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	num := big.NewInt(int64(b.blockHeight))
	if number != nil {
		num.Set(number)
	}

	bg := b.g.excessBlobGas + uint64(b.g.getEpoch())

	return &types.Header{
		Number:        num,
		BaseFee:       b.g.baseFee(),
		ExcessBlobGas: &bg,
	}, nil
}

func (b *mockBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	if b.g.err != nil {
		return 0, b.g.err
	}
	if msg.GasFeeCap.Cmp(msg.GasTipCap) < 0 {
		return 0, core.ErrTipAboveFeeCap
	}

	return b.g.baseFee().Uint64(), nil
}

func (b *mockBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	tip, _ := b.g.sample()
	return tip, nil
}

func (b *mockBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if b.send == nil {
		panic("set sender function was not set")
	}

	return b.send(ctx, tx)
}

func (b *mockBackend) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return startingNonce, nil
}

func (b *mockBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return startingNonce, nil
}

func (*mockBackend) ChainID(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1), nil
}

// TransactionReceipt queries the mockBackend for a mined txHash. If none is found, nil is returned
// for both return values. Otherwise, it returns a receipt containing the txHash, the gasFeeCap
// used in GasUsed, and the blobFeeCap in CumuluativeGasUsed to make the values accessible from our
// test framework.
func (b *mockBackend) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	txInfo, ok := b.minedTxs[txHash]
	if !ok {
		return nil, errors.New("not found")
	}

	// Return the gas fee cap for the transaction in the GasUsed field so that
	// we can assert the proper tx confirmed in our tests.
	var blobFeeCap uint64
	if txInfo.blobFeeCap != nil {
		blobFeeCap = txInfo.blobFeeCap.Uint64()
	}

	return &types.Receipt{
		TxHash:            txHash,
		GasUsed:           txInfo.gasFeeCap.Uint64(),
		CumulativeGasUsed: blobFeeCap,
		BlockNumber:       big.NewInt(int64(txInfo.blockNumber)),
	}, nil
}

func (b *mockBackend) Close() {
}

// TestTxMgrConfirmAtMinGasPrice asserts that Send returns the min gas price tx
// if the tx is mined instantly.
func TestTxMgrConfirmAtMinGasPrice(t *testing.T) {
	t.Parallel()

	h := newTestHarness(t)

	gasPricer := newGasPricer(1)

	gasTipCap, gasFeeCap := gasPricer.sample()
	tx := types.NewTx(&types.DynamicFeeTx{
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
	})

	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		if gasPricer.shouldMine(tx.GasFeeCap()) {
			txHash := tx.Hash()
			h.backend.mine(&txHash, tx.GasFeeCap())
		}

		return nil
	}
	h.backend.setTxSender(sendTx)

	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()
	resTx, receipt, err := h.mgr.sendTx(ctx, tx)
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, resTx.Hash(), tx.Hash())
	require.Equal(t, gasPricer.expGasFeeCap().Uint64(), receipt.GasUsed)
}

// TestTxMgrNeverConfirmCancel asserts that a Send can be canceled even if no
// transaction is mined. This is done to ensure the tx mgr can properly
// abort on shutdown, even if a txn is in the process of being published.
func TestTxMgrNeverConfirmCancel(t *testing.T) {
	// t.SkipNow()
	t.Parallel()

	h := newTestHarness(t)

	gasTipCap, gasFeeCap := h.gasPricer.sample()
	tx := types.NewTx(&types.DynamicFeeTx{
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
	})
	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		// Don't publish tx to backend, simulating never being mined.
		return nil
	}
	h.backend.setTxSender(sendTx)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	_, receipt, err := h.mgr.sendTx(ctx, tx)
	require.Error(t, err)
	require.Nil(t, receipt)
}

// TestTxMgrConfirmsAtMaxGasPrice asserts that Send properly returns the max gas
// price receipt if none of the lower gas price txs were mined.
func TestTxMgrConfirmsAtHigherGasPrice(t *testing.T) {
	t.Parallel()

	h := newTestHarness(t)

	gasTipCap, gasFeeCap := h.gasPricer.sample()
	tx := types.NewTx(&types.DynamicFeeTx{
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
	})
	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		if h.gasPricer.shouldMine(tx.GasFeeCap()) {
			txHash := tx.Hash()
			h.backend.mine(&txHash, tx.GasFeeCap())
		}

		return nil
	}
	h.backend.setTxSender(sendTx)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resTx, receipt, err := h.mgr.sendTx(ctx, tx)
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.NotEqual(t, resTx.Hash(), tx.Hash()) // Different resulting tx
	require.Equal(t, h.gasPricer.expGasFeeCap().Uint64(), receipt.GasUsed)
}

// errRPCFailure is a sentinel error used in testing to fail publications.
var errRPCFailure = errors.New("rpc failure")

// TestTxMgrBlocksOnFailingRpcCalls asserts that if all the publication
// attempts fail due to rpc failures, that the tx manager will return
// ErrPublishTimeout.
func TestTxMgrBlocksOnFailingRpcCalls(t *testing.T) {
	t.Parallel()
	h := newTestHarness(t)
	gasTipCap, gasFeeCap := h.gasPricer.sample()
	orig := types.NewTx(&types.DynamicFeeTx{
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
	})
	errFirst := true
	var sent *types.Transaction
	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		if errFirst {
			errFirst = false
			return errRPCFailure
		}
		sent = tx
		hash := tx.Hash()
		h.backend.mine(&hash, tx.GasFeeCap())

		return nil
	}
	h.backend.setTxSender(sendTx)
	resTx, receipt, err := h.mgr.sendTx(context.Background(), orig)
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, receipt.TxHash, sent.Hash())
	require.NotEqual(t, resTx.Hash(), orig.Hash()) // Different resulting tx
	require.GreaterOrEqual(t, sent.GasTipCap().Uint64(), orig.GasTipCap().Uint64())
}

// TestTxMgr_CraftTx ensures that the tx manager will create transactions as expected.
func TestTxMgr_CraftTx(t *testing.T) {
	t.Parallel()
	h := newTestHarness(t)
	candidate := h.createTxCandidate()
	candidate.Nonce = uint64Ptr(startingNonce)

	// Craft the transaction.
	gasTipCap, gasFeeCap := h.gasPricer.feesForEpoch(h.gasPricer.getEpoch() + 1)
	tx, err := h.mgr.craftTx(context.Background(), candidate)
	require.NoError(t, err)
	require.NotNil(t, tx)
	require.Equal(t, byte(types.DynamicFeeTxType), tx.Type())

	// Validate the gas tip cap and fee cap.
	require.Equal(t, gasTipCap, tx.GasTipCap())
	require.Equal(t, gasFeeCap, tx.GasFeeCap())

	// Craft transaction uses candidate nonce.
	require.EqualValues(t, startingNonce, tx.Nonce())

	// Check that the gas was set using the gas limit.
	require.Equal(t, candidate.GasLimit, tx.Gas())
}

// TestTxMgr_EstimateGas ensures that the tx manager will estimate
// the gas when candidate gas limit is zero in [craftTx].
func TestTxMgr_EstimateGas(t *testing.T) {
	t.Parallel()
	h := newTestHarness(t)
	candidate := h.createTxCandidate()
	candidate.Nonce = uint64Ptr(startingNonce)

	// Set the gas limit to zero to trigger gas estimation.
	candidate.GasLimit = 0

	// Gas estimate
	gasEstimate := h.gasPricer.baseBaseFee.Uint64()

	// Craft the transaction.
	tx, err := h.mgr.craftTx(context.Background(), candidate)
	require.NoError(t, err)
	require.NotNil(t, tx)

	// Check that the gas was estimated correctly.
	require.Equal(t, gasEstimate, tx.Gas())
}

func TestTxMgr_EstimateGasFails(t *testing.T) {
	t.Parallel()
	h := newTestHarness(t)
	candidate := h.createTxCandidate()

	// Set the gas limit to zero to trigger gas estimation.
	candidate.GasLimit = 0
	candidate.Nonce = uint64Ptr(startingNonce)

	// Craft a successful transaction.
	tx, err := h.mgr.craftTx(context.Background(), candidate)
	require.NoError(t, err)
	require.EqualValues(t, startingNonce, tx.Nonce())

	// Mock gas estimation failure.
	h.gasPricer.err = errors.New("execution error")
	_, err = h.mgr.craftTx(context.Background(), candidate)
	require.ErrorContains(t, err, "failed to estimate gas")

	// Ensure successful craft uses the correct nonce
	h.gasPricer.err = nil
	tx, err = h.mgr.craftTx(context.Background(), candidate)
	require.NoError(t, err)
	require.EqualValues(t, startingNonce, tx.Nonce())
}

func TestTxMgr_SigningFails(t *testing.T) {
	t.Parallel()
	errorSigning := false
	cfg := configWithNumConfs(1)
	cfg.Signer = func(ctx context.Context, from common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if errorSigning {
			return nil, errors.New("signer error")
		} else {
			return tx, nil
		}
	}
	h := newTestHarnessWithConfig(t, cfg)
	candidate := h.createTxCandidate()

	const nonce uint64 = 99
	candidate.Nonce = uint64Ptr(nonce)
	// Set the gas limit to zero to trigger gas estimation.
	candidate.GasLimit = 0

	// Craft a successful transaction.
	tx, err := h.mgr.craftTx(context.Background(), candidate)
	require.NoError(t, err)
	require.Equal(t, nonce, tx.Nonce())

	// Mock signer failure.
	errorSigning = true
	_, err = h.mgr.craftTx(context.Background(), candidate)
	require.ErrorContains(t, err, "signer error")

	// Ensure successful craft uses the correct nonce
	errorSigning = false
	tx, err = h.mgr.craftTx(context.Background(), candidate)
	require.NoError(t, err)
	require.Equal(t, nonce, tx.Nonce())
}

// TestTxMgrOnlyOnePublicationSucceeds asserts that the tx manager will return a
// receipt so long as at least one of the publications is able to succeed with a
// simulated rpc failure.
func TestTxMgrOnlyOnePublicationSucceeds(t *testing.T) {
	t.Parallel()

	h := newTestHarness(t)

	gasTipCap, gasFeeCap := h.gasPricer.sample()
	tx := types.NewTx(&types.DynamicFeeTx{
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
	})

	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		// Fail all but the final attempt.
		if !h.gasPricer.shouldMine(tx.GasFeeCap()) {
			return errRPCFailure
		}

		txHash := tx.Hash()
		h.backend.mine(&txHash, tx.GasFeeCap())

		return nil
	}
	h.backend.setTxSender(sendTx)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resTx, receipt, err := h.mgr.sendTx(ctx, tx)
	require.NoError(t, err)
	require.NotEqual(t, resTx.Hash(), tx.Hash()) // Different resulting tx
	require.NotNil(t, receipt)
	require.Equal(t, h.gasPricer.expGasFeeCap().Uint64(), receipt.GasUsed)
}

// TestTxMgrConfirmsMinGasPriceAfterBumping delays the mining of the initial tx
// with the minimum gas price, and asserts that its receipt is returned even
// though if the gas price has been bumped in other goroutines.
func TestTxMgrConfirmsMinGasPriceAfterBumping(t *testing.T) {
	t.Parallel()

	h := newTestHarness(t)

	gasTipCap, gasFeeCap := h.gasPricer.sample()
	tx := types.NewTx(&types.DynamicFeeTx{
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
	})

	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		// Delay mining the tx with the min gas price.
		if h.gasPricer.shouldMine(tx.GasFeeCap()) {
			time.AfterFunc(1*time.Second, func() {
				txHash := tx.Hash()
				h.backend.mine(&txHash, tx.GasFeeCap())
			})
		}

		return nil
	}
	h.backend.setTxSender(sendTx)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resTx, receipt, err := h.mgr.sendTx(ctx, tx)
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.NotEqual(t, resTx.Hash(), tx.Hash()) // Different resulting tx
	require.Equal(t, h.gasPricer.expGasFeeCap().Uint64(), receipt.GasUsed)
}

// TestTxMgrDoesntAbortNonceTooLowAfterMiningTx.
func TestTxMgrDoesntAbortNonceTooLowAfterMiningTx(t *testing.T) {
	t.Parallel()

	h := newTestHarnessWithConfig(t, configWithNumConfs(2))

	gasTipCap, gasFeeCap := h.gasPricer.sample()
	tx := types.NewTx(&types.DynamicFeeTx{
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
	})

	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		switch {
		// If the txn's gas fee cap is less than the one we expect to mine,
		// accept the txn to the mempool.
		case tx.GasFeeCap().Cmp(h.gasPricer.expGasFeeCap()) < 0:
			return nil

		// Accept and mine the actual txn we expect to confirm.
		case h.gasPricer.shouldMine(tx.GasFeeCap()):
			txHash := tx.Hash()
			h.backend.mine(&txHash, tx.GasFeeCap())
			time.AfterFunc(1*time.Second, func() {
				h.backend.mine(nil, nil)
			})

			return nil

		// For gas prices greater than our expected, return ErrNonceTooLow since
		// the prior txn confirmed and will invalidate subsequent publications.
		default:
			return core.ErrNonceTooLow
		}
	}
	h.backend.setTxSender(sendTx)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resTx, receipt, err := h.mgr.sendTx(ctx, tx)
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.NotEqual(t, resTx.Hash(), tx.Hash()) // Different resulting tx
	require.Equal(t, h.gasPricer.expGasFeeCap().Uint64(), receipt.GasUsed)
}

// TestWaitMinedReturnsReceiptOnFirstSuccess insta-mines a transaction and
// asserts that waitMined returns the appropriate receipt.
func TestWaitMinedReturnsReceiptOnFirstSuccess(t *testing.T) {
	t.Parallel()

	h := newTestHarnessWithConfig(t, configWithNumConfs(1))

	// Create a tx and mine it immediately using the default backend.
	tx := types.NewTx(&types.LegacyTx{})
	txHash := tx.Hash()
	h.backend.mine(&txHash, new(big.Int))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	receipt, err := h.mgr.waitMined(ctx, tx, testSendState())
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, receipt.TxHash, txHash)
}

// TestWaitMinedCanBeCanceled ensures that waitMined exits if the passed context
// is canceled before a receipt is found.
func TestWaitMinedCanBeCanceled(t *testing.T) {
	t.Parallel()

	h := newTestHarness(t)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	// Create an unimined tx.
	tx := types.NewTx(&types.LegacyTx{})

	receipt, err := h.mgr.waitMined(ctx, tx, NewSendState(10, time.Hour))
	require.Error(t, err)
	require.Nil(t, receipt)
}

// TestWaitMinedMultipleConfs asserts that waitMined will properly wait for more
// than one confirmation.
func TestWaitMinedMultipleConfs(t *testing.T) {
	t.Parallel()

	const numConfs = 2

	h := newTestHarnessWithConfig(t, configWithNumConfs(numConfs))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create an unimined tx.
	tx := types.NewTx(&types.LegacyTx{})
	txHash := tx.Hash()
	h.backend.mine(&txHash, new(big.Int))

	receipt, err := h.mgr.waitMined(ctx, tx, NewSendState(10, time.Hour))
	require.Error(t, err)
	require.Nil(t, receipt)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Mine an empty block, tx should now be confirmed.
	h.backend.mine(nil, nil)
	receipt, err = h.mgr.waitMined(ctx, tx, NewSendState(10, time.Hour))
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, txHash, receipt.TxHash)
}

// failingBackend implements ReceiptSource, returning a failure on the
// first call but a success on the second call. This allows us to test that the
// inner loop of waitMined properly handles this case.
type failingBackend struct {
	ethclient.Client
	txReceiptErr             error // Defaults to ethereum.NotFound if nil
	returnSuccessBlockNumber bool
	returnSuccessHeader      bool
	returnSuccessReceipt     bool
	baseFee, gasTip          *big.Int
	excessBlobGas            *uint64
}

// BlockNumber for the failingBackend returns errRPCFailure on the first
// invocation and a fixed block height on subsequent calls.
func (b *failingBackend) BlockNumber(_ context.Context) (uint64, error) {
	if !b.returnSuccessBlockNumber {
		b.returnSuccessBlockNumber = true
		return 0, errRPCFailure
	}

	return 1, nil
}

// TransactionReceipt for the failingBackend returns errRPCFailure on the first
// invocation, and a receipt containing the passed TxHash on the second.
func (b *failingBackend) TransactionReceipt(
	_ context.Context, txHash common.Hash,
) (*types.Receipt, error) {
	if !b.returnSuccessReceipt {
		b.returnSuccessReceipt = true

		err := ethereum.NotFound
		if b.txReceiptErr != nil {
			err = b.txReceiptErr
		}

		return nil, err
	}

	return &types.Receipt{
		TxHash:      txHash,
		BlockNumber: big.NewInt(1),
	}, nil
}

func (b *failingBackend) HeaderByNumber(_ context.Context, _ *big.Int) (*types.Header, error) {
	return &types.Header{
		Number:        big.NewInt(1),
		BaseFee:       b.baseFee,
		ExcessBlobGas: b.excessBlobGas,
	}, nil
}

func (b *failingBackend) SuggestGasTipCap(_ context.Context) (*big.Int, error) {
	return b.gasTip, nil
}

func (b *failingBackend) EstimateGas(_ context.Context, msg ethereum.CallMsg) (uint64, error) {
	return b.baseFee.Uint64(), nil
}

// TestWaitMinedReturnsReceiptAfterNotFound asserts that waitMined is able to
// recover from failed calls to the backend. It uses the failedBackend to
// simulate an rpc call failure, followed by the successful return of a receipt.
func TestWaitMinedReturnsReceiptAfterNotFound(t *testing.T) {
	t.Parallel()

	var borkedBackend failingBackend

	mgr := &simple{
		cfg: Config{
			ResubmissionTimeout:       time.Second,
			ReceiptQueryInterval:      50 * time.Millisecond,
			NumConfirmations:          1,
			SafeAbortNonceTooLowCount: 3,
		},
		chainName: "TEST",
		backend:   &borkedBackend,
	}

	// Don't mine the tx with the default backend. The failingBackend will
	// return the txHash on the second call.
	tx := types.NewTx(&types.LegacyTx{})
	txHash := tx.Hash()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	receipt, err := mgr.waitMined(ctx, tx, testSendState())
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, receipt.TxHash, txHash)
}

func TestWaitMinedRetriesNetError(t *testing.T) {
	t.Parallel()

	// Temporary errors or any net.OpError results in retries.
	netErr := context.DeadlineExceeded
	if rand.Float64() > 0.5 {
		netErr = &net.OpError{Op: "read", Err: io.EOF}
	}
	netErrBackend := failingBackend{txReceiptErr: netErr}

	mgr := &simple{
		cfg: Config{
			ResubmissionTimeout:       time.Second,
			ReceiptQueryInterval:      50 * time.Millisecond,
			NumConfirmations:          1,
			SafeAbortNonceTooLowCount: 3,
		},
		chainName: "TEST",
		backend:   &netErrBackend,
	}

	// Don't mine the tx with the default backend. The failingBackend will
	// return the txHash on the second call.
	tx := types.NewTx(&types.LegacyTx{})
	txHash := tx.Hash()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	receipt, err := mgr.waitMined(ctx, tx, testSendState())
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, receipt.TxHash, txHash)
}

func doGasPriceIncrease(_ *testing.T, txTipCap, txFeeCap, newTip,
	newBaseFee int64) (*types.Transaction, *types.Transaction, error) {
	borkedBackend := failingBackend{
		gasTip:              big.NewInt(newTip),
		baseFee:             big.NewInt(newBaseFee),
		returnSuccessHeader: true,
	}

	mgr := &simple{
		cfg: Config{
			ResubmissionTimeout:       time.Second,
			ReceiptQueryInterval:      50 * time.Millisecond,
			NumConfirmations:          1,
			SafeAbortNonceTooLowCount: 3,
			FeeLimitMultiplier:        5,
			Signer: func(ctx context.Context, from common.Address, tx *types.Transaction) (*types.Transaction, error) {
				return tx, nil
			},
			From: common.Address{},
		},
		chainName: "TEST",
		backend:   &borkedBackend,
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		GasTipCap: big.NewInt(txTipCap),
		GasFeeCap: big.NewInt(txFeeCap),
	})
	newTx, err := mgr.increaseGasPrice(context.Background(), tx)

	return tx, newTx, err
}

func TestIncreaseGasPrice(t *testing.T) {
	t.Parallel()
	require.Equal(t, int64(10), PriceBump, "test must be updated if priceBump is adjusted")
	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "bump at least 1",
			run: func(t *testing.T) {
				t.Helper()
				tx, newTx, err := doGasPriceIncrease(t, 1, 3, 1, 1)
				require.Greater(t, newTx.GasFeeCap().Uint64(), tx.GasFeeCap().Uint64(), "new tx fee cap must be larger")
				require.Greater(t, newTx.GasTipCap().Uint64(), tx.GasTipCap().Uint64(), "new tx tip must be larger")
				require.NoError(t, err)
			},
		},
		{
			name: "enforces min bump",
			run: func(t *testing.T) {
				t.Helper()
				tx, newTx, err := doGasPriceIncrease(t, 100, 1000, 101, 460)
				require.Greater(t, newTx.GasFeeCap().Uint64(), tx.GasFeeCap().Uint64(), "new tx fee cap must be larger")
				require.Greater(t, newTx.GasTipCap().Uint64(), tx.GasTipCap().Uint64(), "new tx tip must be larger")
				require.NoError(t, err)
			},
		},
		{
			name: "enforces min bump on only tip increase",
			run: func(t *testing.T) {
				t.Helper()
				tx, newTx, err := doGasPriceIncrease(t, 100, 1000, 101, 440)
				require.Greater(t, newTx.GasFeeCap().Uint64(), tx.GasFeeCap().Uint64(), "new tx fee cap must be larger")
				require.Greater(t, newTx.GasTipCap().Uint64(), tx.GasTipCap().Uint64(), "new tx tip must be larger")
				require.NoError(t, err)
			},
		},
		{
			name: "enforces min bump on only base fee increase",
			run: func(t *testing.T) {
				t.Helper()
				tx, newTx, err := doGasPriceIncrease(t, 100, 1000, 99, 460)
				require.Greater(t, newTx.GasFeeCap().Uint64(), tx.GasFeeCap().Uint64(), "new tx fee cap must be larger")
				require.Greater(t, newTx.GasTipCap().Uint64(), tx.GasTipCap().Uint64(), "new tx tip must be larger")
				require.NoError(t, err)
			},
		},
		{
			name: "uses L1 values when larger",
			run: func(t *testing.T) {
				t.Helper()
				_, newTx, err := doGasPriceIncrease(t, 10, 100, 50, 200)
				require.Equal(t, newTx.GasFeeCap().Uint64(), big.NewInt(450).Uint64(), "new tx fee cap must be equal L1")
				require.Equal(t, newTx.GasTipCap().Uint64(), big.NewInt(50).Uint64(), "new tx tip must be equal L1")
				require.NoError(t, err)
			},
		},
		{
			name: "uses L1 tip when larger and threshold FC",
			run: func(t *testing.T) {
				t.Helper()
				_, newTx, err := doGasPriceIncrease(t, 100, 2200, 120, 1050)
				require.Equal(t, newTx.GasTipCap().Uint64(), big.NewInt(120).Uint64(), "new tx tip must be equal L1")
				require.Equal(t, newTx.GasFeeCap().Uint64(), big.NewInt(2420).Uint64(),
					"new tx fee cap must be equal to the threshold value")
				require.NoError(t, err)
			},
		},
		{
			name: "bumped fee above multiplier limit",
			run: func(t *testing.T) {
				t.Helper()
				_, _, err := doGasPriceIncrease(t, 1, 9999, 1, 1)
				require.ErrorContains(t, err, "bumped cap")
				require.NotContains(t, err.Error(), "tip cap")
			},
		},
		{
			name: "bumped tip above multiplier limit",
			run: func(t *testing.T) {
				t.Helper()
				_, _, err := doGasPriceIncrease(t, 9999, 0, 0, 9999)
				require.ErrorContains(t, err, "bumped cap")
				require.NotContains(t, err.Error(), "fee cap")
			},
		},
		{
			name: "bumped fee and tip above multiplier limit",
			run: func(t *testing.T) {
				t.Helper()
				_, _, err := doGasPriceIncrease(t, 9999, 9999, 1, 1)
				require.ErrorContains(t, err, "bumped cap")
			},
		},
		{
			name: "uses L1 FC when larger and threshold tip",
			run: func(t *testing.T) {
				t.Helper()
				_, newTx, err := doGasPriceIncrease(t, 100, 2200, 100, 2000)
				require.Equal(t, big.NewInt(110).Uint64(),
					newTx.GasTipCap().Uint64(), "new tx tip must be equal the threshold value")
				require.Equal(t, big.NewInt(4110).Uint64(),
					newTx.GasFeeCap().Uint64(), "new tx fee cap must be equal L1")

				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.run(t)
		})
	}
}

// TestIncreaseGasPriceLimits asserts that if the L1 base fee & tip remain the
// same, repeated calls to increaseGasPrice eventually hit a limit.
func TestIncreaseGasPriceLimits(t *testing.T) {
	t.Parallel()
	t.Run("no-threshold", func(t *testing.T) {
		t.Parallel()
		testIncreaseGasPriceLimit(t, gasPriceLimitTest{
			expTipCap:     46,
			expFeeCap:     354, // just below 5*100
			expBlobFeeCap: 4 * params.GWei,
		})
	})
	t.Run("with-threshold", func(t *testing.T) {
		t.Parallel()
		testIncreaseGasPriceLimit(t, gasPriceLimitTest{
			thr:           big.NewInt(params.GWei * 10),
			expTipCap:     1_293_535_754,
			expFeeCap:     9_192_620_686, // just below 10 gwei
			expBlobFeeCap: 8 * params.GWei,
		})
	})
}

type gasPriceLimitTest struct {
	thr                  *big.Int
	expTipCap, expFeeCap int64
	expBlobFeeCap        int64
}

// testIncreaseGasPriceLimit runs a gas bumping test that increases the gas price until it hits an error.
// It starts with a tx that has a tip cap of 10 wei and fee cap of 100 wei.
func testIncreaseGasPriceLimit(t *testing.T, lt gasPriceLimitTest) {
	t.Helper()

	borkedTip := int64(10)
	borkedFee := int64(45)
	// simulate 100 excess blobs which yields a 50 wei blob base fee
	borkedExcessBlobGas := uint64(100 * params.BlobTxBlobGasPerBlob)
	borkedBackend := failingBackend{
		gasTip:              big.NewInt(borkedTip),
		baseFee:             big.NewInt(borkedFee),
		excessBlobGas:       &borkedExcessBlobGas,
		returnSuccessHeader: true,
	}

	mgr := &simple{
		cfg: Config{
			ResubmissionTimeout:       time.Second,
			ReceiptQueryInterval:      50 * time.Millisecond,
			NumConfirmations:          1,
			SafeAbortNonceTooLowCount: 3,
			FeeLimitMultiplier:        5,
			FeeLimitThreshold:         lt.thr,
			Signer: func(ctx context.Context, from common.Address, tx *types.Transaction) (*types.Transaction, error) {
				return tx, nil
			},
			From: common.Address{},
		},
		chainName: "TEST",
		backend:   &borkedBackend,
	}
	lastGoodTx := types.NewTx(&types.DynamicFeeTx{
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(100),
	})

	// Run increaseGasPrice a bunch of times in a row to simulate a very fast resubmit loop to make
	// sure it errors out without a runaway fee increase.
	ctx := context.Background()
	var err error
	for {
		var tmpTx *types.Transaction
		tmpTx, err = mgr.increaseGasPrice(ctx, lastGoodTx)
		if err != nil {
			break
		}
		lastGoodTx = tmpTx
	}
	require.Error(t, err)

	// Confirm that fees only rose until expected threshold
	require.Equal(t, lt.expTipCap, lastGoodTx.GasTipCap().Int64())
	require.Equal(t, lt.expFeeCap, lastGoodTx.GasFeeCap().Int64())
}

func TestErrStringMatch(t *testing.T) {
	t.Parallel()
	tests := []struct {
		err    error
		target error
		match  bool
	}{
		{err: nil, target: nil, match: true},
		{err: errors.New("exists"), target: nil, match: false},
		{err: nil, target: errors.New("exists"), match: false},
		{err: errors.New("exact match"), target: errors.New("exact match"), match: true},
		{err: errors.New("partial: match"), target: errors.New("match"), match: true},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			require.Equal(t, test.match, errStringMatch(test.err, test.target))
		})
	}
}

func TestNonceReset(t *testing.T) {
	t.Parallel()
	conf := configWithNumConfs(1)
	conf.SafeAbortNonceTooLowCount = 1
	h := newTestHarnessWithConfig(t, conf)

	index := -1
	var nonces []uint64
	sendTx := func(ctx context.Context, tx *types.Transaction) error {
		index++
		nonces = append(nonces, tx.Nonce())
		// fail every 3rd tx
		if index%3 == 0 {
			return core.ErrNonceTooLow
		}
		txHash := tx.Hash()
		h.backend.mine(&txHash, tx.GasFeeCap())

		return nil
	}
	h.backend.setTxSender(sendTx)

	ctx := context.Background()
	for i := 0; i < 8; i++ {
		_, _, err := h.mgr.Send(ctx, TxCandidate{
			To: &common.Address{},
		})
		// expect every 3rd tx to fail
		if i%3 == 0 {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}

	// internal nonce tracking should be reset to startingNonce value every 3rd tx
	require.Equal(t, []uint64{1, 1, 2, 3, 1, 2, 3, 1}, nonces)
}

func TestMinFees(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		desc             string
		minBaseFee       *big.Int
		minTipCap        *big.Int
		expectMinBaseFee bool
		expectMinTipCap  bool
	}{
		{
			desc: "no-mins",
		},
		{
			desc:             "high-min-basefee",
			minBaseFee:       big.NewInt(10_000_000),
			expectMinBaseFee: true,
		},
		{
			desc:            "high-min-tipcap",
			minTipCap:       big.NewInt(1_000_000),
			expectMinTipCap: true,
		},
		{
			desc:             "high-mins",
			minBaseFee:       big.NewInt(10_000_000),
			minTipCap:        big.NewInt(1_000_000),
			expectMinBaseFee: true,
			expectMinTipCap:  true,
		},
		{
			desc:       "low-min-basefee",
			minBaseFee: big.NewInt(1),
		},
		{
			desc:      "low-min-tipcap",
			minTipCap: big.NewInt(1),
		},
		{
			desc:       "low-mins",
			minBaseFee: big.NewInt(1),
			minTipCap:  big.NewInt(1),
		},
	} {
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			conf := configWithNumConfs(1)
			conf.MinBaseFee = tt.minBaseFee
			conf.MinTipCap = tt.minTipCap
			h := newTestHarnessWithConfig(t, conf)

			tip, baseFee, err := h.mgr.suggestGasPriceCaps(context.Background())
			require.NoError(t, err)

			if tt.expectMinBaseFee {
				require.Equal(t, tt.minBaseFee, baseFee, "expect suggested base fee to equal MinBaseFee")
			} else {
				require.Equal(t, h.gasPricer.baseBaseFee, baseFee, "expect suggested base fee to equal mock base fee")
			}

			if tt.expectMinTipCap {
				require.Equal(t, tt.minTipCap, tip, "expect suggested tip to equal MinTipCap")
			} else {
				require.Equal(t, h.gasPricer.baseGasTipFee, tip, "expect suggested tip to equal mock tip")
			}
		})
	}
}

func uint64Ptr(i uint64) *uint64 {
	return &i
}
