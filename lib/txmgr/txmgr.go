package txmgr

import (
	"context"
	"log/slog"
	"math/big"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/p2p/netutil"
	"github.com/ethereum/go-ethereum/params"
)

const (
	// PriceBump geth requires a minimum fee bump of 10% for regular tx resubmission.
	PriceBump int64 = 10
)

// TxManager is an interface that allows callers to reliably publish txs,
// bumping the gas price if needed, and obtain the receipt of the resulting tx.
type TxManager interface {
	// Send is used to create & doSend a transaction. It will handle increasing
	// the gas price & ensuring that the transaction remains in the transaction pool.
	// It can be stopped by canceling the provided context; however, the transaction
	// may be included on L1 even if the context is canceled.
	//
	// NOTE: Send can be called concurrently, the nonce will be managed internally.
	Send(ctx context.Context, candidate TxCandidate) (*types.Transaction, *types.Receipt, error)

	// From returns the sending address associated with the instance of the transaction manager.
	// It is static for a single instance of a TxManager.
	From() common.Address

	// ReserveNextNonce returns the next available nonce and increments the available nonce.
	ReserveNextNonce(ctx context.Context) (uint64, error)
}

// simple is a implementation of TxManager that performs linear fee
// bumping of a tx until it confirms.
type simple struct {
	cfg       Config // embed the config directly
	chainName string
	chainID   *big.Int

	backend ethclient.Client

	nonce     *uint64 // nil == unset, 0 == unused account
	nonceLock sync.Mutex
}

// NewSimple initializes a new simple with the passed Config.
func NewSimple(chainName string, conf Config) (TxManager, error) {
	if err := conf.Check(); err != nil {
		return nil, errors.Wrap(err, "invalid config")
	}

	return &simple{
		chainID:   conf.ChainID,
		chainName: chainName,
		cfg:       conf,
		backend:   conf.Backend,
	}, nil
}

func (m *simple) ReserveNextNonce(ctx context.Context) (uint64, error) {
	m.nonceLock.Lock()
	defer m.nonceLock.Unlock()

	if m.nonce == nil {
		// Ensure the node is synced before fetching the nonce
		syncing, err := m.backend.SyncProgress(ctx)
		if ethclient.IsErrMethodNotAvailable(err) { //nolint:revive // Empty block skips error handling below.
			// Skip if method not available.
		} else if err != nil {
			return 0, errors.Wrap(err, "sync progress")
		} else if syncing != nil && !syncing.Done() { // Note syncing is nil if node is not syncing.
			return 0, errors.New("backend not synced",
				"lag", umath.SubtractOrZero(syncing.HighestBlock, syncing.CurrentBlock),
				"indexing", syncing.TxIndexRemainingBlocks,
			)
		}

		nonce, err := m.backend.NonceAt(ctx, m.cfg.From, nil)
		if err != nil {
			return 0, errors.Wrap(err, "failed to get nonce")
		}

		log.Debug(ctx, "Txmgr initialized nonce", "nonce", nonce)
		m.nonce = &nonce
	}

	defer func() {
		*m.nonce++
	}()

	return *m.nonce, nil
}

func (m *simple) From() common.Address {
	return m.cfg.From
}

// txFields returns a logger with the transaction hash and nonce fields set.
func txFields(tx *types.Transaction, logGas bool) []any {
	fields := []any{
		slog.Int64("nonce", int64(tx.Nonce())),
		slog.String("tx", tx.Hash().String()),
	}
	if logGas {
		fields = append(fields,
			slog.String("gas_tip_cap", tx.GasTipCap().String()),
			slog.String("gas_fee_cap", tx.GasFeeCap().String()),
			slog.Int64("gas_limit", int64(tx.Gas())),
		)
	}

	return fields
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
	// Nonce to use for the transaction. If nil, the current nonce is used.
	Nonce *uint64
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
func (m *simple) Send(ctx context.Context, candidate TxCandidate) (*types.Transaction, *types.Receipt, error) {
	tx, rec, err := m.doSend(ctx, candidate)
	if err != nil {
		m.resetNonce()
		return nil, nil, err
	}

	return tx, rec, nil
}

// doSend performs the actual transaction creation and sending.
func (m *simple) doSend(ctx context.Context, candidate TxCandidate) (*types.Transaction, *types.Receipt, error) {
	ctx, cancel := maybeSetTimeout(ctx, m.cfg.TxSendTimeout)
	defer cancel()

	backoff := expbackoff.New(ctx, expbackoff.WithFastConfig())
	for {
		if ctx.Err() != nil {
			return nil, nil, errors.Wrap(ctx.Err(), "send timeout")
		}

		// Set a candidate nonce if not already set
		if candidate.Nonce == nil {
			nonce, err := m.ReserveNextNonce(ctx)
			if err != nil {
				log.Warn(ctx, "Failed to reserve nonce (will retry)", err)
				backoff()

				continue
			}
			candidate.Nonce = &nonce
		}

		// Create the initial transaction
		tx, err := m.craftTx(ctx, candidate)
		if err != nil {
			log.Warn(ctx, "Failed to create transaction (will retry)", err, "nonce", *candidate.Nonce)
			backoff()

			continue
		}

		// Send it (note this has internal retries bumping fees, so might return a different tx)
		tx, rec, err := m.sendTx(ctx, tx)
		if err != nil {
			return nil, nil, errors.Wrap(err, "send tx")
		}

		return tx, rec, nil
	}
}

// craftTx creates the signed transaction
// It queries L1 for the current fee market conditions as well as for the nonce.
// NOTE: This method SHOULD NOT publish the resulting transaction.
// NOTE: If the [TxCandidate.GasLimit] is non-zero, it will be used as the transaction's gas.
// NOTE: Otherwise, the [simple] will query the specified backend for an estimate.
func (m *simple) craftTx(ctx context.Context, candidate TxCandidate) (*types.Transaction, error) {
	if candidate.Nonce == nil {
		return nil, errors.New("invalid nil nonce")
	}

	gasTipCap, baseFee, err := m.suggestGasPriceCaps(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get gas price info")
	}
	gasFeeCap := calcGasFeeCap(baseFee, gasTipCap)

	gasLimit := candidate.GasLimit

	// If the gas limit is set, we can use that as the gas
	if gasLimit == 0 {
		// Calculate the intrinsic gas for the transaction
		gas, err := m.backend.EstimateGas(ctx, ethereum.CallMsg{
			From:      m.cfg.From,
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
		ChainID:   m.chainID,
		To:        candidate.To,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     candidate.Value,
		Data:      candidate.TxData,
		Gas:       gasLimit,
		Nonce:     *candidate.Nonce,
	}

	ctx, cancel := context.WithTimeout(ctx, m.cfg.NetworkTimeout)
	defer cancel()

	tx, err := m.cfg.Signer(ctx, m.cfg.From, types.NewTx(txMessage))
	if err != nil {
		return nil, errors.Wrap(err, "sign tx")
	}

	log.Debug(ctx, "Crafted tx", "sender", m.cfg.From, "nonce", tx.Nonce(), "tx", tx.Hash())

	return tx, nil
}

// resetNonce resets the internal nonce tracking. This is called if any pending doSend
// returns an error.
func (m *simple) resetNonce() {
	m.nonceLock.Lock()
	defer m.nonceLock.Unlock()
	m.nonce = nil
}

// sendTx submits the same transaction several times with increasing gas prices as necessary.
// It waits for the transaction to be confirmed on chain.
// It returns the confirmed transaction.
func (m *simple) sendTx(ctx context.Context, tx *types.Transaction) (*types.Transaction, *types.Receipt, error) {
	var wg sync.WaitGroup
	defer wg.Wait()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sendState := NewSendState(m.cfg.SafeAbortNonceTooLowCount, m.cfg.TxNotInMempoolTimeout)
	minedChan := make(chan minedTuple, 1)
	publishAndWait := func(tx *types.Transaction, bumpFees bool) *types.Transaction {
		if bumpFees {
			resendTotal.WithLabelValues(m.chainName).Inc()
		}

		wg.Add(1)
		tx, published := m.publishTx(ctx, tx, sendState, bumpFees)
		if published {
			go func() {
				defer wg.Done()
				m.waitForTx(ctx, tx, sendState, minedChan)
			}()
		} else {
			wg.Done()
		}

		return tx
	}

	// Immediately publish a transaction before starting the resubmission loop
	tx = publishAndWait(tx, false)

	ticker := time.NewTicker(m.cfg.ResubmissionTimeout)
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
				return nil, nil, errors.New("aborted transaction sending", txFields(tx, false)...)
			}
			tx = publishAndWait(tx, true)

		case <-ctx.Done():
			return nil, nil, errors.Wrap(ctx.Err(), "timeout")

		case mined := <-minedChan:
			if mined.Tx.Hash() != mined.Rec.TxHash { // Sanity check
				return nil, nil, errors.New("mined hash mismatch [BUG]", "tx", mined.Tx.Hash(), "receipt", mined.Rec.TxHash)
			}
			if mined.Rec.EffectiveGasPrice != nil {
				txEffectiveGasPrice.
					WithLabelValues(m.chainName).
					Set(float64(mined.Rec.EffectiveGasPrice.Uint64() / params.GWei))
				txGasUsed.WithLabelValues(m.chainName).Observe(float64(mined.Rec.GasUsed))
			}

			return mined.Tx, mined.Rec, nil
		}
	}
}

// publishTx publishes the transaction to the transaction pool. If it receives any underpriced errors
// it will bump the fees and retry.
// Returns the latest fee bumped tx, and a boolean indicating whether the tx was sent or not.
func (m *simple) publishTx(ctx context.Context, tx *types.Transaction, sendState *SendState, bumpFeesImmediately bool) (*types.Transaction, bool) {
	for {
		if ctx.Err() != nil {
			return tx, false
		}

		if bumpFeesImmediately {
			newTx, err := m.increaseGasPrice(ctx, tx)
			if err != nil {
				log.Info(ctx, "Unable to increase gas", err)
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

		cCtx, cancel := context.WithTimeout(ctx, m.cfg.NetworkTimeout)
		err := m.backend.SendTransaction(cCtx, tx)
		cancel()
		sendState.ProcessSendError(err)

		if err == nil {
			return tx, true
		}

		// Handle known errors. The errors are returned by json-rpc and therefore we
		// can't rely on the error type and need to compare the error message instead.
		switch {
		case errStringMatch(err, core.ErrNonceTooLow):
			log.Warn(ctx, "Nonce too low", err)
		case errStringMatch(err, context.Canceled) || errStringMatch(err, context.DeadlineExceeded):
			log.Warn(ctx, "Transaction send canceled", err)
		case errStringMatch(err, txpool.ErrAlreadyKnown):
			log.Warn(ctx, "Resubmitted already known transaction", err)
		case errStringMatch(err, txpool.ErrReplaceUnderpriced):
			log.Warn(ctx, "Transaction replacement is underpriced", err)
			continue // retry with fee bump
		case errStringMatch(err, txpool.ErrUnderpriced):
			log.Warn(ctx, "Transaction is underpriced", err)
			continue // retry with fee bump
		case errStringMatch(err, core.ErrIntrinsicGas):
			log.Warn(ctx, "Intrinsic gas too low", err)
			continue // retry with fee bump
		default:
			log.Error(ctx, "Unknown error publishing transaction", err)
		}

		// on non-underpriced error return immediately; will retry on next resubmission timeout
		return tx, false
	}
}

// waitForTx calls waitMined, and then sends the receipt to receiptChan in a non-blocking way if a receipt is found
// for the transaction. It should be called in a separate goroutine.
func (m *simple) waitForTx(ctx context.Context, tx *types.Transaction, sendState *SendState, minedChan chan minedTuple) {
	t0 := time.Now()

	// Poll for the transaction to be ready & then doSend the result to receiptChan
	receipt, err := m.waitMined(ctx, tx, sendState)
	if ctx.Err() != nil {
		return
	} else if err != nil {
		// this will happen if the tx was successfully replaced by a tx with bumped fees
		log.Warn(ctx, "Transaction receipt not mined, probably replaced", err)
		return
	}
	txConfirmationLatency.WithLabelValues(m.chainName).Set(time.Since(t0).Seconds())

	select {
	case minedChan <- minedTuple{
		Tx:  tx,
		Rec: receipt,
	}:
	default:
		log.Error(ctx, "Multiple txs mined for same nonce [BUG]", nil)
	}
}

// waitMined waits for the transaction to be mined or for the context to be canceled.
func (m *simple) waitMined(ctx context.Context, tx *types.Transaction,
	sendState *SendState) (*types.Receipt, error) {
	txHash := tx.Hash()
	const logFreqFactor = 10 // Log every 10th attempt
	attempt := 1

	netErrBackoff := expbackoff.New(ctx) // Additional backoff on network errors.

	queryTicker := time.NewTicker(m.cfg.ReceiptQueryInterval)
	defer queryTicker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "context canceled")
		case <-queryTicker.C:
			receipt, ok, err := m.queryReceipt(ctx, txHash, sendState)
			if netutil.IsTemporaryError(err) || isNetworkError(err) {
				// Treat all network errors as temporary, since we know we submitted the tx already.
				// Network issues might resolve, and the tx might still be mined, don't give up.
				log.Warn(ctx, "Temporary network error querying receipt (will retry)", err, "attempt", attempt)
				netErrBackoff()
			} else if err != nil {
				return nil, err
			} else if !ok && attempt%logFreqFactor == 0 {
				log.Warn(ctx, "Transaction not yet mined", nil, "attempt", attempt)
			} else if ok {
				return receipt, nil
			}

			attempt++
		}
	}
}

// queryReceipt queries for the receipt and returns the receipt if it has passed the confirmation depth.
func (m *simple) queryReceipt(ctx context.Context, txHash common.Hash,
	sendState *SendState) (*types.Receipt, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, m.cfg.NetworkTimeout)
	defer cancel()
	receipt, err := m.backend.TransactionReceipt(ctx, txHash)
	if err != nil {
		if strings.Contains(err.Error(), "transaction indexing is in progress") {
			return nil, false, nil // Just back off here
		} else if errors.Is(err, ethereum.NotFound) || strings.Contains(err.Error(), ethereum.NotFound.Error()) {
			sendState.TxNotMined(txHash)
			return nil, false, nil
		}

		return nil, false, err
	}

	// Receipt is confirmed to be valid from this point on
	sendState.TxMined(txHash)

	txHeight := receipt.BlockNumber.Uint64()
	tip, err := m.backend.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, false, err
	}

	// The transaction is considered confirmed when
	// txHeight+numConfirmations-1 <= tipHeight. Note that the -1 is
	// needed to account for the fact that confirmations have an
	// inherent off-by-one, i.e. when using 1 confirmation the
	// transaction should be confirmed when txHeight is equal to
	// tipHeight. The equation is rewritten in this form to avoid
	// underflow'm.
	tipHeight := tip.Number.Uint64()
	if txHeight+m.cfg.NumConfirmations <= tipHeight+1 {
		return receipt, true, nil
	}

	return nil, false, nil
}

// increaseGasPrice returns a new transaction that is equivalent to the input transaction but with
// higher fees that should satisfy geth's tx replacement rules. It also computes an updated gas
// limit estimate. To avoid runaway price increases, fees are capped at a `feeLimitMultiplier`
// multiple of the suggested values.
func (m *simple) increaseGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	log.Debug(ctx, "Bumping gas price")
	tip, baseFee, err := m.suggestGasPriceCaps(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get gas price info", txFields(tx, true)...)
	}
	bumpedTip, bumpedFee := updateFees(ctx, tx.GasTipCap(), tx.GasFeeCap(), tip, baseFee)

	if err := m.checkLimits(tip, baseFee, bumpedTip, bumpedFee); err != nil {
		return nil, err
	}

	// Re-estimate gaslimit in case things have changed or a previous gaslimit estimate was wrong
	gas, err := m.backend.EstimateGas(ctx, ethereum.CallMsg{
		From:      m.cfg.From,
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
		log.Warn(ctx, "Failed to re-estimate gas", err,
			"gas_limit", tx.Gas(),
			"gas_fee_cap", bumpedFee,
			"gas_tip_cap", bumpedTip,
		)
		// just log and carry on
		gas = tx.Gas()
	}
	if tx.Gas() != gas {
		log.Debug(ctx, "Re-estimated gas differs",
			"old_gas", tx.Gas(),
			"new_gas", gas,
			"gas_fee_cap", bumpedFee,
			"gas_tip_cap", bumpedTip,
		)
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

	ctx, cancel := context.WithTimeout(ctx, m.cfg.NetworkTimeout)
	defer cancel()
	signedTx, err := m.cfg.Signer(ctx, m.cfg.From, newTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// suggestGasPriceCaps suggests what the new tip, base fee, and blob base fee should be based on
// the current L1 conditions. blobfee will be nil if 4844 is not yet active.
func (m *simple) suggestGasPriceCaps(ctx context.Context) (*big.Int, *big.Int, error) {
	cCtx, cancel := context.WithTimeout(ctx, m.cfg.NetworkTimeout)
	defer cancel()
	tip, err := m.backend.SuggestGasTipCap(cCtx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to fetch the suggested gas tip cap")
	} else if tip == nil {
		return nil, nil, errors.New("the suggested tip was nil")
	}
	cCtx, cancel = context.WithTimeout(ctx, m.cfg.NetworkTimeout)
	defer cancel()
	head, err := m.backend.HeaderByNumber(cCtx, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to fetch the suggested base fee")
	} else if head.BaseFee == nil {
		return nil, nil, errors.New("txmgr does not support pre-london blocks that do not have a base fee")
	}

	baseFee := head.BaseFee

	// Enforce minimum base fee and tip cap
	if minTipCap := m.cfg.MinTipCap; minTipCap != nil && tip.Cmp(minTipCap) == -1 {
		log.Debug(ctx, "Enforcing min tip cap", "min_tip_cap", m.cfg.MinTipCap, "orig_tip_cap", tip)
		tip = new(big.Int).Set(m.cfg.MinTipCap)
	}
	if minBaseFee := m.cfg.MinBaseFee; minBaseFee != nil && baseFee.Cmp(minBaseFee) == -1 {
		log.Debug(ctx, "Enforcing min base fee", "min_base_fee", m.cfg.MinBaseFee, "orig_base_fee", baseFee)
		baseFee = new(big.Int).Set(m.cfg.MinBaseFee)
	}

	return tip, baseFee, nil
}

// checkLimits checks that the tip and baseFee have not increased by more than the configured multipliers
// if FeeLimitThreshold is specified in config, any increase which stays under the threshold are allowed.
func (m *simple) checkLimits(tip, baseFee, bumpedTip, bumpedFee *big.Int) error {
	threshold := m.cfg.FeeLimitThreshold
	limit := big.NewInt(int64(m.cfg.FeeLimitMultiplier))
	maxTip := new(big.Int).Mul(tip, limit)
	maxFee := calcGasFeeCap(new(big.Int).Mul(baseFee, limit), maxTip)
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

// updateFees takes an old transaction's tip & fee cap plus a new tip & base fee, and returns
// a suggested tip and fee cap such that:
//
//	(a) each satisfies geth's required tx-replacement fee bumps, and
//	(b) gasTipCap is no less than new tip, and
//	(c) gasFeeCap is no less than calcGasFee(newBaseFee, newTip)
func updateFees(ctx context.Context, oldTip, oldFeeCap, newTip, newBaseFee *big.Int) (*big.Int, *big.Int) {
	newFeeCap := calcGasFeeCap(newBaseFee, newTip)
	log.Debug(ctx, "Updating fees", "old_gas_tip_cap", oldTip, "old_gas_fee_cap", oldFeeCap,
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
		return thresholdTip, calcGasFeeCap(newBaseFee, thresholdTip)
	}

	log.Debug(ctx, "Using threshold tip and threshold feecap")

	return thresholdTip, thresholdFeeCap
}

// calcGasFeeCap deterministically computes the recommended gas fee cap given
// the base fee and gasTipCap. The resulting gasFeeCap is equal to:
//
//	gasTipCap + 2*baseFee.
func calcGasFeeCap(baseFee, gasTipCap *big.Int) *big.Int {
	return new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(baseFee, big.NewInt(2)),
	)
}

// errStringMatch returns true if err.Error() is a substring in target.Error() or if both are nil.
// It can accept nil errors without issue.
func errStringMatch(err, target error) bool {
	if err == nil && target == nil {
		return true
	} else if err == nil || target == nil {
		return false
	}

	return strings.Contains(err.Error(), target.Error())
}

// maybeSetTimeout returns a copy of the context with the timeout set if timeout is not zero.
// If the timeout is zero, it doesn't set it and just returns the context.
func maybeSetTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	if timeout == 0 {
		return ctx, func() {}
	}

	return context.WithTimeout(ctx, timeout)
}

// isNetworkError returns true if the error is a network error.
func isNetworkError(err error) bool {
	if err == nil {
		return false
	}

	opErr := new(net.OpError)

	return errors.As(err, &opErr)
}

// minedTTuple groups the mined/confirmed tx with its receipt.
type minedTuple struct {
	Tx  *types.Transaction
	Rec *types.Receipt
}
