package txmgr

import (
	"context"
	"log/slog"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	// PriceBump geth requires a minimum fee bump of 10% for regular tx resubmission.
	PriceBump int64 = 10
)

var (
	ErrBlobFeeLimit = errors.New("blob fee limit reached")
	ErrClosed       = errors.New("transaction manager is closed")
)

// TxManager is an interface that allows callers to reliably publish txs,
// bumping the gas price if needed, and obtain the receipt of the resulting tx.
//
//go:generate mockery --name TxManager --output ./mocks
type TxManager interface {
	// Send is used to create & doSend a transaction. It will handle increasing
	// the gas price & ensuring that the transaction remains in the transaction pool.
	// It can be stopped by canceling the provided context; however, the transaction
	// may be included on L1 even if the context is canceled.
	//
	// NOTE: Send can be called concurrently, the nonce will be managed internally.
	Send(ctx context.Context, candidate TxCandidate) (*types.Receipt, error)

	// From returns the sending address associated with the instance of the transaction manager.
	// It is static for a single instance of a TxManager.
	From() common.Address

	// BlockNumber returns the most recent block number from the underlying network.
	BlockNumber(ctx context.Context) (uint64, error)

	// Close the underlying connection
	Close()
}

// ETHBackend is the set of methods that the transaction manager uses to resubmit gas & determine
// when transactions are included on L1.
type ETHBackend interface {
	// BlockNumber returns the most recent block number.
	BlockNumber(ctx context.Context) (uint64, error)

	// CallContract executes an eth_call against the provided contract.
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)

	// TransactionReceipt queries the backend for a receipt associated with
	// txHash. If lookup does not fail, but the transaction is not found,
	// nil should be returned for both values.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)

	// SendTransaction submits a signed transaction to L1.
	SendTransaction(ctx context.Context, tx *types.Transaction) error

	// These functions are used to estimate what the base fee & priority fee should be set to.
	// TODO(CLI-3318): Maybe need a generic interface to support different RPC providers
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	// NonceAt returns the account nonce of the given account.
	// The block number can be nil, in which case the nonce is taken from the latest known block.
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	// PendingNonceAt returns the pending nonce.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	// EstimateGas returns an estimate of the amount of gas needed to execute the given
	// transaction against the current pending block.
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	// Close the underlying eth connection
	Close()
}

// SimpleTxManager is a implementation of TxManager that performs linear fee
// bumping of a tx until it confirms.
type SimpleTxManager struct {
	Cfg     Config // embed the config directly
	Name    string
	ChainID *big.Int

	Backend ETHBackend

	nonce     *uint64
	nonceLock sync.RWMutex

	pending atomic.Int64

	Closed atomic.Bool
}

// NewSimpleTxManagerFromConfig initializes a new SimpleTxManager with the passed Config.
func NewSimpleTxManagerFromConfig(name string, conf Config) (*SimpleTxManager, error) {
	if err := conf.Check(); err != nil {
		return nil, errors.Wrap(err, "invalid config")
	}

	return &SimpleTxManager{
		ChainID: conf.ChainID,
		Name:    name,
		Cfg:     conf,
		Backend: conf.Backend,
	}, nil
}

func (m *SimpleTxManager) From() common.Address {
	return m.Cfg.From
}

func (m *SimpleTxManager) BlockNumber(ctx context.Context) (uint64, error) {
	return m.Backend.BlockNumber(ctx)
}

// Close closes the underlying connection, and sets the closed flag.
// once closed, the tx manager will refuse to doSend any new transactions, and may abandon pending ones.
func (m *SimpleTxManager) Close() {
	m.Backend.Close()
	m.Closed.Store(true)
}

// txLogger returns a logger with the transaction hash and nonce fields set.
//
//nolint:revive // Might fix
func (*SimpleTxManager) txLogger(ctx context.Context, tx *types.Transaction,
	logGas bool) *slog.Logger {
	fields := []any{"tx", tx.Hash(), "nonce", tx.Nonce()}
	if logGas {
		fields = append(fields, "gasTipCap", tx.GasTipCap(), "gasFeeCap", tx.GasFeeCap(), "gasLimit", tx.Gas())
	}
	if len(tx.BlobHashes()) != 0 {
		// log the number of blobs a tx has only if it's a blob tx
		fields = append(fields, "blobs", len(tx.BlobHashes()))
	}

	return log.GetLogger(ctx).With(fields...)
}

// TxCandidate is a transaction candidate that can be submitted to ask the
// [TxManager] to construct a transaction with gas price bounds.
type TxCandidate struct {
	// TxData is the transaction calldata to be used in the constructed tx.
	TxData []byte
	// To is the recipient of the constructed tx. Nil means contract creation.
	To *common.Address
	// GasLimit is the gas limit to be used in the constructed tx.
	GasLimit uint64
	// Value is the value to be used in the constructed tx.
	Value *big.Int
}

// Send is used to publish a transaction with incrementally higher gas prices
// until the transaction eventually confirms. This method blocks until an
// invocation of sendTx returns (called with differing gas prices). The method
// may be canceled using the passed context.
//
// The transaction manager handles all signing. If and only if the gas limit is 0, the
// transaction manager will do a gas estimation.
//
// NOTE: Send can be called concurrently, the nonce will be managed internally.
func (m *SimpleTxManager) Send(ctx context.Context, candidate TxCandidate) (*types.Receipt, error) {
	// refuse new requests if the tx manager is closed
	if m.Closed.Load() {
		return nil, ErrClosed
	}
	m.pending.Add(1)
	defer func() {
		m.pending.Add(-1)
	}()
	receipt, err := m.doSend(ctx, candidate)
	if err != nil {
		m.resetNonce()
	}

	return receipt, err
}

// doSend performs the actual transaction creation and sending.
func (m *SimpleTxManager) doSend(ctx context.Context, candidate TxCandidate) (*types.Receipt, error) {
	if m.Cfg.TxSendTimeout != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, m.Cfg.TxSendTimeout)
		defer cancel()
	}
	tx, err := Do(ctx, 30, Fixed(2*time.Second), func() (*types.Transaction, error) {
		if m.Closed.Load() {
			return nil, ErrClosed
		}
		tx, err := m.CraftTx(ctx, candidate)
		if err != nil {
			log.Debug(ctx, "Failed to create a transaction, will retry", "err", err)
		}

		return tx, err
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create the tx")
	}

	return m.SendTx(ctx, tx)
}

// CraftTx creates the signed transaction
// It queries L1 for the current fee market conditions as well as for the nonce.
// NOTE: This method SHOULD NOT publish the resulting transaction.
// NOTE: If the [TxCandidate.GasLimit] is non-zero, it will be used as the transaction's gas.
// NOTE: Otherwise, the [SimpleTxManager] will query the specified Backend for an estimate.
func (m *SimpleTxManager) CraftTx(ctx context.Context, candidate TxCandidate) (*types.Transaction, error) {
	gasTipCap, baseFee, _, err := m.SuggestGasPriceCaps(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get gas price info")
	}
	gasFeeCap := CalcGasFeeCap(baseFee, gasTipCap)

	gasLimit := candidate.GasLimit

	// If the gas limit is set, we can use that as the gas
	if gasLimit == 0 {
		// Calculate the intrinsic gas for the transaction
		gas, err := m.Backend.EstimateGas(ctx, ethereum.CallMsg{
			From:      m.Cfg.From,
			To:        candidate.To,
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
			Data:      candidate.TxData,
			Value:     candidate.Value,
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to estimate gas")
		}
		gasLimit = gas
	}

	txMessage := &types.DynamicFeeTx{
		ChainID:   m.ChainID,
		To:        candidate.To,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     candidate.Value,
		Data:      candidate.TxData,
		Gas:       gasLimit,
	}

	return m.signWithNextNonce(ctx, txMessage) // signer sets the nonce field of the tx
}

// signWithNextNonce returns a signed transaction with the next available nonce.
// The nonce is fetched once using eth_getTransactionCount with "latest", and
// then subsequent calls simply increment this number. If the transaction manager
// is reset, it will query the eth_getTransactionCount nonce again. If signing
// fails, the nonce is not incremented.
func (m *SimpleTxManager) signWithNextNonce(ctx context.Context, txMessage types.TxData) (*types.Transaction, error) {
	m.nonceLock.Lock()
	defer m.nonceLock.Unlock()

	if m.nonce == nil {
		// Fetch the sender's nonce from the latest known block (nil `blockNumber`)
		childCtx, cancel := context.WithTimeout(ctx, m.Cfg.NetworkTimeout)
		defer cancel()
		nonce, err := m.Backend.NonceAt(childCtx, m.Cfg.From, nil)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get nonce")
		}
		m.nonce = &nonce
	} else {
		*m.nonce++
	}

	switch x := txMessage.(type) {
	case *types.DynamicFeeTx:
		x.Nonce = *m.nonce
	default:
		return nil, errors.New("unrecognized tx type", x)
	}
	ctx, cancel := context.WithTimeout(ctx, m.Cfg.NetworkTimeout)
	defer cancel()
	tx, err := m.Cfg.Signer(ctx, m.Cfg.From, types.NewTx(txMessage))
	if err != nil {
		// decrement the nonce, so we can retry signing with the same nonce next time
		// signWithNextNonce is called
		*m.nonce--
	}

	return tx, err
}

// resetNonce resets the internal nonce tracking. This is called if any pending doSend
// returns an error.
func (m *SimpleTxManager) resetNonce() {
	m.nonceLock.Lock()
	defer m.nonceLock.Unlock()
	m.nonce = nil
}

// doSend submits the same transaction several times with increasing gas prices as necessary.
// It waits for the transaction to be confirmed on chain.
func (m *SimpleTxManager) SendTx(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	var wg sync.WaitGroup
	defer wg.Wait()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sendState := NewSendState(m.Cfg.SafeAbortNonceTooLowCount, m.Cfg.TxNotInMempoolTimeout)
	receiptChan := make(chan *types.Receipt, 1)
	publishAndWait := func(tx *types.Transaction, bumpFees bool) *types.Transaction {
		wg.Add(1)
		tx, published := m.publishTx(ctx, tx, sendState, bumpFees)
		if published {
			go func() {
				defer wg.Done()
				m.waitForTx(ctx, tx, sendState, receiptChan)
			}()
		} else {
			wg.Done()
		}

		return tx
	}

	// Immediately publish a transaction before starting the resumbission loop
	tx = publishAndWait(tx, false)

	ticker := time.NewTicker(m.Cfg.ResubmissionTimeout)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Don't resubmit a transaction if it has been mined, but we are waiting for the conf depth.
			if sendState.IsWaitingForConfirmation() {
				continue
			}
			// If we see lots of unrecoverable errors (and no pending transactions) abort sending the transaction.
			if sendState.ShouldAbortImmediately() {
				m.txLogger(ctx, tx, false).Warn("Aborting transaction submission")
				return nil, errors.New("aborted transaction sending")
			}
			// if the tx manager closed while we were waiting for the tx, give up
			if m.Closed.Load() {
				m.txLogger(ctx, tx, false).Warn("TxManager closed, aborting transaction submission")
				return nil, ErrClosed
			}
			tx = publishAndWait(tx, true)

		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "context canceled")

		case receipt := <-receiptChan:
			return receipt, nil
		}
	}
}

// publishTx publishes the transaction to the transaction pool. If it receives any underpriced errors
// it will bump the fees and retry.
// Returns the latest fee bumped tx, and a boolean indicating whether the tx was sent or not.
//
//nolint:revive // Might fix
func (m *SimpleTxManager) publishTx(ctx context.Context, tx *types.Transaction, sendState *SendState,
	bumpFeesImmediately bool) (*types.Transaction, bool) {
	for {
		// if the tx manager closed, give up without bumping fees or retrying
		if m.Closed.Load() {
			return tx, false
		}
		if bumpFeesImmediately {
			newTx, err := m.IncreaseGasPrice(ctx, tx)
			if err != nil {
				log.Info(ctx, "unable to increase gas", err)
				return tx, false
			}
			tx = newTx
			sendState.bumpCount++
		}
		bumpFeesImmediately = true // bump fees next loop

		if sendState.IsWaitingForConfirmation() {
			// there is a chance the previous tx goes into "waiting for confirmation" state
			// during the increaseGasPrice call; continue waiting rather than resubmit the tx
			return tx, false
		}

		cCtx, cancel := context.WithTimeout(ctx, m.Cfg.NetworkTimeout)
		err := m.Backend.SendTransaction(cCtx, tx)
		cancel()
		sendState.ProcessSendError(err)

		if err == nil {
			return tx, true
		}

		switch {
		case ErrStringMatch(err, core.ErrNonceTooLow):
			log.Warn(ctx, "nonce too low", err)
		case ErrStringMatch(err, context.Canceled):
			log.Warn(ctx, "transaction doSend canceled", err)
		case ErrStringMatch(err, txpool.ErrAlreadyKnown):
			log.Warn(ctx, "resubmitted already known transaction", err)
		case ErrStringMatch(err, txpool.ErrReplaceUnderpriced):
			log.Warn(ctx, "transaction replacement is underpriced", err)
			continue // retry with fee bump
		case ErrStringMatch(err, txpool.ErrUnderpriced):
			log.Warn(ctx, "transaction is underpriced", err)
			continue // retry with fee bump
		default:
			log.Warn(ctx, "unknown error publishing transaction", err)
		}

		// on non-underpriced error return immediately; will retry on next resubmission timeout
		return tx, false
	}
}

// waitForTx calls WaitMined, and then sends the receipt to receiptChan in a non-blocking way if a receipt is found
// for the transaction. It should be called in a separate goroutine.
func (m *SimpleTxManager) waitForTx(ctx context.Context, tx *types.Transaction, sendState *SendState,
	receiptChan chan *types.Receipt) {
	// Poll for the transaction to be ready & then doSend the result to receiptChan
	receipt, err := m.WaitMined(ctx, tx, sendState)
	if err != nil {
		// this will happen if the tx was successfully replaced by a tx with bumped fees
		m.txLogger(ctx, tx, true).Debug("Transaction receipt not found", "err", err)
		return
	}
	select {
	case receiptChan <- receipt:
	default:
	}
}

// WaitMined waits for the transaction to be mined or for the context to be canceled.
func (m *SimpleTxManager) WaitMined(ctx context.Context, tx *types.Transaction,
	sendState *SendState) (*types.Receipt, error) {
	txHash := tx.Hash()
	queryTicker := time.NewTicker(m.Cfg.ReceiptQueryInterval)
	defer queryTicker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "context canceled")
		case <-queryTicker.C:
			if receipt := m.queryReceipt(ctx, txHash, sendState); receipt != nil {
				return receipt, nil
			}
		}
	}
}

// queryReceipt queries for the receipt and returns the receipt if it has passed the confirmation depth.
func (m *SimpleTxManager) queryReceipt(ctx context.Context, txHash common.Hash, sendState *SendState) *types.Receipt {
	ctx, cancel := context.WithTimeout(ctx, m.Cfg.NetworkTimeout)
	defer cancel()
	receipt, err := m.Backend.TransactionReceipt(ctx, txHash)
	if errors.Is(err, ethereum.NotFound) {
		sendState.TxNotMined(txHash)
		log.Info(ctx, "Transaction not yet mined", "tx", txHash)

		return nil
	} else if err != nil {
		log.Warn(ctx, "Receipt retrieval failed", err, txHash)

		return nil
	} else if receipt == nil {
		log.Info(ctx, "Receipt is nil", txHash)

		return nil
	}

	// Receipt is confirmed to be valid from this point on
	sendState.TxMined(txHash)

	txHeight := receipt.BlockNumber.Uint64()
	tip, err := m.Backend.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Warn(ctx, "Unable to fetch tip", err)

		return nil
	}

	// The transaction is considered confirmed when
	// txHeight+numConfirmations-1 <= tipHeight. Note that the -1 is
	// needed to account for the fact that confirmations have an
	// inherent off-by-one, i.e. when using 1 confirmation the
	// transaction should be confirmed when txHeight is equal to
	// tipHeight. The equation is rewritten in this form to avoid
	// underflows.
	tipHeight := tip.Number.Uint64()
	if txHeight+m.Cfg.NumConfirmations <= tipHeight+1 {
		return receipt
	}

	return nil
}

// IncreaseGasPrice returns a new transaction that is equivalent to the input transaction but with
// higher fees that should satisfy geth's tx replacement rules. It also computes an updated gas
// limit estimate. To avoid runaway price increases, fees are capped at a `feeLimitMultiplier`
// multiple of the suggested values.
func (m *SimpleTxManager) IncreaseGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	m.txLogger(ctx, tx, true).Info("bumping gas price for transaction")
	tip, baseFee, _, err := m.SuggestGasPriceCaps(ctx)
	if err != nil {
		m.txLogger(ctx, tx, false).Warn("failed to get suggested gas tip and base fee", "err", err)
		return nil, err
	}
	bumpedTip, bumpedFee := UpdateFees(ctx, tx.GasTipCap(), tx.GasFeeCap(), tip, baseFee)

	if err := m.checkLimits(tip, baseFee, bumpedTip, bumpedFee); err != nil {
		return nil, err
	}

	// Re-estimate gaslimit in case things have changed or a previous gaslimit estimate was wrong
	gas, err := m.Backend.EstimateGas(ctx, ethereum.CallMsg{
		From:      m.Cfg.From,
		To:        tx.To(),
		GasTipCap: bumpedTip,
		GasFeeCap: bumpedFee,
		Data:      tx.Data(),
		Value:     tx.Value(),
	})
	if err != nil {
		// If this is a transaction resubmission, we sometimes see this outcome because the
		// original tx can get included in a block just before the above call. In this case the
		// error is due to the tx reverting with message "block number must be equal to next
		// expected block number"
		log.Warn(ctx, "failed to re-estimate gas", err, "tx", tx.Hash(), "gaslimit", tx.Gas(),
			"gasFeeCap", bumpedFee, "gasTipCap", bumpedTip)

		return nil, err
	}
	if tx.Gas() != gas {
		log.Debug(ctx, "re-estimated gas differs", "tx", tx.Hash(), "old_gas", tx.Gas(), "new_gas", gas,
			"gas_fee_Cap", bumpedFee, "gas_tip_cap", bumpedTip)
	}

	newTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   tx.ChainId(),
		Nonce:     tx.Nonce(),
		To:        tx.To(),
		GasTipCap: bumpedTip,
		GasFeeCap: bumpedFee,
		Value:     tx.Value(),
		Data:      tx.Data(),
		Gas:       gas,
	})

	ctx, cancel := context.WithTimeout(ctx, m.Cfg.NetworkTimeout)
	defer cancel()
	signedTx, err := m.Cfg.Signer(ctx, m.Cfg.From, newTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// SuggestGasPriceCaps suggests what the new tip, base fee, and blob base fee should be based on
// the current L1 conditions. blobfee will be nil if 4844 is not yet active.
func (m *SimpleTxManager) SuggestGasPriceCaps(ctx context.Context) (*big.Int, *big.Int, *big.Int, error) {
	cCtx, cancel := context.WithTimeout(ctx, m.Cfg.NetworkTimeout)
	defer cancel()
	tip, err := m.Backend.SuggestGasTipCap(cCtx)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to fetch the suggested gas tip cap")
	} else if tip == nil {
		return nil, nil, nil, errors.New("the suggested tip was nil")
	}
	cCtx, cancel = context.WithTimeout(ctx, m.Cfg.NetworkTimeout)
	defer cancel()
	head, err := m.Backend.HeaderByNumber(cCtx, nil)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to fetch the suggested base fee")
	} else if head.BaseFee == nil {
		return nil, nil, nil, errors.New("txmgr does not support pre-london blocks that do not have a base fee")
	}

	baseFee := head.BaseFee

	// Enforce minimum base fee and tip cap
	if minTipCap := m.Cfg.MinTipCap; minTipCap != nil && tip.Cmp(minTipCap) == -1 {
		log.Debug(ctx, "Enforcing min tip cap", "minTipCap", m.Cfg.MinTipCap, "origTipCap", tip)
		tip = new(big.Int).Set(m.Cfg.MinTipCap)
	}
	if minBaseFee := m.Cfg.MinBaseFee; minBaseFee != nil && baseFee.Cmp(minBaseFee) == -1 {
		log.Debug(ctx, "Enforcing min base fee", "minBaseFee", m.Cfg.MinBaseFee, "origBaseFee", baseFee)
		baseFee = new(big.Int).Set(m.Cfg.MinBaseFee)
	}

	var blobFee *big.Int
	if head.ExcessBlobGas != nil {
		blobFee = eip4844.CalcBlobFee(*head.ExcessBlobGas)
	}

	return tip, baseFee, blobFee, nil
}

// checkLimits checks that the tip and baseFee have not increased by more than the configured multipliers
// if FeeLimitThreshold is specified in config, any increase which stays under the threshold are allowed.
func (m *SimpleTxManager) checkLimits(tip, baseFee, bumpedTip, bumpedFee *big.Int) error {
	threshold := m.Cfg.FeeLimitThreshold
	limit := big.NewInt(int64(m.Cfg.FeeLimitMultiplier))
	maxTip := new(big.Int).Mul(tip, limit)
	maxFee := CalcGasFeeCap(new(big.Int).Mul(baseFee, limit), maxTip)
	var errs error
	// generic check function to check tip and fee, and build up an error
	check := func(v, max *big.Int, name string) {
		// if threshold is specified and the value is under the threshold, no need to check the max
		if threshold != nil && threshold.Cmp(v) > 0 {
			return
		}
		// if the value is over the max, add an error message
		if v.Cmp(max) > 0 {
			errs = errors.New("bumped cap is over multiple of the suggested value", name, v, limit)
		}
	}
	check(bumpedTip, maxTip, "tip")
	check(bumpedFee, maxFee, "fee")

	return errs
}

// calcThresholdValue returns ceil(x * priceBumpPercent / 100) for non-blob txs, or
// It guarantees that x is increased by at least 1.
func calcThresholdValue(x *big.Int) *big.Int {
	threshold := new(big.Int)
	ninetyNine := big.NewInt(99)
	oneHundred := big.NewInt(100)
	priceBumpPercent := big.NewInt(100 + PriceBump)
	threshold.Set(priceBumpPercent)

	return threshold.Mul(threshold, x).Add(threshold, ninetyNine).Div(threshold, oneHundred)
}

// UpdateFees takes an old transaction's tip & fee cap plus a new tip & base fee, and returns
// a suggested tip and fee cap such that:
//
//	(a) each satisfies geth's required tx-replacement fee bumps, and
//	(b) gasTipCap is no less than new tip, and
//	(c) gasFeeCap is no less than calcGasFee(newBaseFee, newTip)
func UpdateFees(ctx context.Context, oldTip, oldFeeCap, newTip, newBaseFee *big.Int) (*big.Int, *big.Int) {
	newFeeCap := CalcGasFeeCap(newBaseFee, newTip)
	log.Debug(ctx, "updating fees", "old_gas_tip_cap", oldTip, "old_gas_fee_cap", oldFeeCap,
		"new_gas_tip_cap", newTip, "new_gas_fee_cap", newFeeCap, "new_base_fee", newBaseFee)
	thresholdTip := calcThresholdValue(oldTip)
	thresholdFeeCap := calcThresholdValue(oldFeeCap)
	if newTip.Cmp(thresholdTip) >= 0 && newFeeCap.Cmp(thresholdFeeCap) >= 0 {
		log.Debug(ctx, "Using new tip and feecap")
		return newTip, newFeeCap
	} else if newTip.Cmp(thresholdTip) >= 0 && newFeeCap.Cmp(thresholdFeeCap) < 0 {
		// Tip has gone up, but base fee is flat or down.
		// TODO(CLI-3714): Do we need to recalculate the FC here?
		log.Debug(ctx, "Using new tip and threshold feecap")
		return newTip, thresholdFeeCap
	} else if newTip.Cmp(thresholdTip) < 0 && newFeeCap.Cmp(thresholdFeeCap) >= 0 {
		// Base fee has gone up, but the tip hasn't. Recalculate the feecap because if the tip went up a lot
		// not enough of the feecap may be dedicated to paying the base fee.
		log.Debug(ctx, "Using threshold tip and recalculated feecap")
		return thresholdTip, CalcGasFeeCap(newBaseFee, thresholdTip)
	}

	log.Debug(ctx, "Using threshold tip and threshold feecap")

	return thresholdTip, thresholdFeeCap
}

// CalcGasFeeCap deterministically computes the recommended gas fee cap given
// the base fee and gasTipCap. The resulting gasFeeCap is equal to:
//
//	gasTipCap + 2*baseFee.
func CalcGasFeeCap(baseFee, gasTipCap *big.Int) *big.Int {
	return new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(baseFee, big.NewInt(2)),
	)
}

// ErrStringMatch returns true if err.Error() is a substring in target.Error() or if both are nil.
// It can accept nil errors without issue.
func ErrStringMatch(err, target error) bool {
	if err == nil && target == nil {
		return true
	} else if err == nil || target == nil {
		return false
	}

	return strings.Contains(err.Error(), target.Error())
}
