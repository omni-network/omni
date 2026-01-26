package rebalance

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cctp"
	cctpdb "github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/uniswap"
	"github.com/omni-network/omni/solver/fundthresh"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// minSend is the minimum amount of surplus USDC to send to other chains.
	minSend = bi.Dec6(1000) // 1k USDC

	// maxSend is the maximum amount of USDC allowed in a single send.
	maxSend = bi.Dec6(10000) // 10k USDC
)

// rebalanceCCTPForever starts rebalancing loops for each chain in the cctp network.
// Full CCTP rebalancing also requires Uniswap V3 liquidity.
func rebalanceCCTPForever(
	ctx context.Context,
	interval time.Duration,
	db *cctpdb.DB,
	network netconf.Network,
	pricer tokenpricer.Pricer,
	backends ethbackend.Backends,
	solver common.Address,
) {
	for {
		start := time.Now()
		rebalanceCCTPOnce(ctx, db, network, pricer, backends, solver)
		elapsed := time.Since(start)

		// Sleep for the remaining time in the interval, if any.
		if elapsed < interval {
			time.Sleep(interval - elapsed)
		}
	}
}

// rebalanceCCTPOnce rebalances all chains once.
func rebalanceCCTPOnce(
	ctx context.Context,
	db *cctpdb.DB,
	network netconf.Network,
	pricer tokenpricer.Pricer,
	backends ethbackend.Backends,
	solver common.Address,
) {
	for _, chain := range network.EVMChains() {
		func() {
			ctx := log.WithCtx(ctx, "chain", evmchain.Name(chain.ID))
			log.Debug(ctx, "Rebalancing chain; trying lock")
			defer lock(chain.ID)() // Lock the chain to prevent concurrent rebalancing.
			log.Info(ctx, "Rebalancing chain; locked")

			// First, swap surplus tokens USDC.
			if err := swapSurplusOnce(ctx, backends, chain.ID, solver); err != nil {
				log.Warn(ctx, "Failed to swap surplus", err)
			}

			// Then, fill deficits from surplus USDC.
			if err := fillDeficitOnce(ctx, pricer, backends, chain.ID, solver); err != nil {
				log.Warn(ctx, "Failed to fill deficit", err)
			}

			// Finally, send remaining surplus USDC to other chains.
			if err := sendSurplusOnce(ctx, db, network, pricer, backends, chain.ID, solver); err != nil {
				log.Warn(ctx, "Failed to send surplus", err)
			}
		}()
	}
}

// sendSurplusOnce sends surplus USDC on `chainID` to chains in deficit.
func sendSurplusOnce(
	ctx context.Context,
	db *cctpdb.DB,
	network netconf.Network,
	pricer tokenpricer.Pricer,
	backends ethbackend.Backends,
	chainID uint64,
	solver common.Address,
) error {
	ctx = log.WithCtx(ctx, "step", "sendSurplus")

	backend, err := backends.Backend(chainID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	usdc, ok := tokens.ByAsset(chainID, tokens.USDC)
	if !ok {
		return errors.New("token not found")
	}

	surplus, err := GetSurplus(ctx, backend, usdc, solver)
	if err != nil {
		return errors.Wrap(err, "get surplus")
	}

	if bi.LT(surplus, minSend) { // Only send if surplus is above minSend.
		log.Debug(ctx, "No surplus to send", "amount", usdc.FormatAmt(surplus))
		return nil
	}

	deficits, err := GetUSDChainDeficits(ctx, db, network, backends.Clients(), pricer, solver)
	if err != nil {
		return errors.Wrap(err, "get deficits")
	}

	// For clarity
	thisChainID := chainID

	deficitHere, ok := find(deficits, func(d ChainAmount) bool { return d.ChainID == thisChainID })
	if !ok {
		return errors.New("missing deficit")
	}

	if bi.LT(surplus, deficitHere.Amount) { // Surplus < deficit here, don't send it elsewhere.
		log.Debug(ctx, "Surplus < deficit here, skipping send",
			"deficit", formatUSD(deficitHere.Amount),
			"amount", usdc.FormatAmt(surplus))

		return nil
	}

	usdcBalance, err := tokenutil.BalanceOf(ctx, backend, usdc, solver)
	if err != nil {
		return errors.Wrap(err, "get usdc balance")
	}

	log.Debug(ctx, "Preparing to send surplus",
		"deficit_here", formatUSD(deficitHere.Amount),
		"usdc_balance", usdc.FormatAmt(usdcBalance),
		"surplus", usdc.FormatAmt(surplus),
		"min_send", usdc.FormatAmt(minSend),
		"max_send", usdc.FormatAmt(maxSend))

	// Subtract deficit here from available surplus (if positive).
	// Decrement surplus each time we send USDC.
	if bi.GT(deficitHere.Amount, bi.Zero()) {
		surplus = bi.Sub(surplus, deficitHere.Amount)
	}

	for _, d := range deficits {
		if d.ChainID == thisChainID { // Skip self.
			continue
		}

		ctx := log.WithCtx(ctx,
			"dest", evmchain.Name(d.ChainID),
			"deficit", formatUSD(d.Amount),
			"surplus", usdc.FormatAmt(surplus),
			"min", usdc.FormatAmt(minSend),
			"max", usdc.FormatAmt(maxSend))

		if bi.LTE(d.Amount, bi.Zero()) { // No deficit, no need to send.
			log.Debug(ctx, "No deficit, skipping send")
			continue
		}

		toSend := d.Amount

		if bi.GT(toSend, surplus) { // Cap send to available surplus.
			log.Debug(ctx, "Deficit > surplus, capping send")
			toSend = surplus
		}

		if bi.LT(toSend, minSend) { // Not enough worth sending.
			log.Debug(ctx, "Send < minSend, skipping send")
			break // Can break, because deficits is sorted by descending amount.
		}

		if bi.GT(toSend, maxSend) { // Cap send to maxSend.
			log.Debug(ctx, "Send > maxSend, capping send")
			toSend = maxSend
		}

		if _, err = cctp.SendUSDC(ctx, db, network.ID, backend, cctp.SendUSDCArgs{
			Sender:      solver,
			Recipient:   solver,
			SrcChainID:  chainID,
			DestChainID: d.ChainID,
			Amount:      toSend,
		}); err != nil {
			return errors.Wrap(err, "send usdc", "dest", evmchain.Name(d.ChainID))
		}

		// Decrement surplus by the amount sent.
		surplus = bi.Sub(surplus, toSend)
	}

	return nil
}

// swapSurplusOnce swaps surplus tokens to USDC on `chainID`.
func swapSurplusOnce(
	ctx context.Context,
	backends ethbackend.Backends,
	chainID uint64,
	solver common.Address,
) error {
	ctx = log.WithCtx(ctx, "step", "swapSurplus")

	backend, err := backends.Backend(chainID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	for _, token := range SwappableTokensByChain(chainID) {
		if token.Is(tokens.USDC) { // Already USDC, skip.
			continue
		}

		if err := swapTokenSurplusOnce(ctx, backend, token, solver); err != nil {
			return errors.Wrap(err, "swap surplus", "token", token)
		}
	}

	return nil
}

// swapTokenSurplusOnce swaps any surplus of `token` to USDC.
func swapTokenSurplusOnce(
	ctx context.Context,
	backend *ethbackend.Backend,
	token tokens.Token,
	solver common.Address,
) error {
	surplus, err := GetSurplus(ctx, backend, token, solver)
	if err != nil {
		return errors.Wrap(err, "get surplus")
	}

	maxSwap := fundthresh.Get(token).MaxSwap()
	minSwap := fundthresh.Get(token).MinSwap()

	if bi.IsZero(maxSwap) { // Require max swap.
		log.Warn(ctx, "No max swap set, skipping", errors.New("missing max swap"), "token", token)
		return nil
	}

	if bi.LTE(surplus, minSwap) { // Surplus <= minSwap, do nothing.
		log.Debug(ctx, "Surplus < minSwap, skipping", "amount", token.FormatAmt(surplus), "min", token.FormatAmt(minSwap))
		return nil
	}

	toSwap := surplus
	if bi.GT(toSwap, maxSwap) { // Cap swap to maxSwap.
		log.Debug(ctx, "Surplus > maxSwap, capping swap", "amount", token.FormatAmt(toSwap), "max", token.FormatAmt(maxSwap))
		toSwap = maxSwap
	}

	// If native, leave some buffer for gas.
	nativeBuffer := bi.Ether(0.05) // 0.05 ETH
	if token.IsNative() {
		if bi.LTE(toSwap, nativeBuffer) { // Not enough to leave buffer, skip.
			log.Debug(ctx, "Surplus <= native buffer, skipping", "amount", token.FormatAmt(toSwap), "buffer", token.FormatAmt(nativeBuffer))
			return nil
		}

		toSwap = bi.Sub(toSwap, nativeBuffer)
		log.Debug(ctx, "Token is native, leaving gas buffer", "buffer", token.FormatAmt(nativeBuffer), "to_swap", token.FormatAmt(toSwap))
	}

	log.Debug(ctx, "Swapping surplus", "amount", token.FormatAmt(toSwap))

	// Swap surplus to USDC.
	_, err = uniswap.SwapToUSDC(ctx, backend, solver, token, toSwap)
	if err != nil {
		return errors.Wrap(err, "swap to usdc")
	}

	return nil
}

// fillDeficitOnce fills token deficits from surplus USDC on `chainID`.
func fillDeficitOnce(
	ctx context.Context,
	pricer tokenpricer.Pricer,
	backends ethbackend.Backends,
	chainID uint64,
	solver common.Address,
) error {
	ctx = log.WithCtx(ctx, "step", "fillDeficit")

	backend, err := backends.Backend(chainID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	for _, token := range SwappableTokensByChain(chainID) {
		if token.Is(tokens.USDC) { // Already USDC, skip.
			continue
		}

		if err := fillTokenDeficitOnce(ctx, pricer, backend, token, solver); err != nil {
			return errors.Wrap(err, "fill deficit", "token", token)
		}
	}

	return nil
}

// fillTokenDeficitOnce fills any deficit of `token` from surplus USDC.
func fillTokenDeficitOnce(
	ctx context.Context,
	pricer tokenpricer.Pricer,
	backend *ethbackend.Backend,
	token tokens.Token,
	solver common.Address,
) error {
	chainID := token.ChainID

	deficit, err := GetDeficit(ctx, backend, token, solver)
	if err != nil {
		return errors.Wrap(err, "get deficit")
	}

	if bi.IsZero(deficit) { // No deficit, nothing to fill.
		log.Debug(ctx, "No deficit", "amount", token.FormatAmt(deficit))
		return nil
	}

	deficitUSD, err := AmtToUSD(ctx, pricer, token, deficit)
	if err != nil {
		return errors.Wrap(err, "get deficit in usd")
	}

	usdc, ok := tokens.ByAsset(chainID, tokens.USDC)
	if !ok {
		return errors.New("token not found")
	}

	surplusUSDC, err := GetSurplus(ctx, backend, usdc, solver)
	if err != nil {
		return errors.Wrap(err, "get surplus")
	}

	toSwap := deficitUSD
	if bi.GT(toSwap, surplusUSDC) { // Deficit > surplus, cap swap to surplus.
		log.Debug(ctx, "Deficit > surplus, capping swap", "deficit", formatUSD(deficitUSD), "surplus", usdc.FormatAmt(surplusUSDC))
		toSwap = surplusUSDC
	}

	minSwap := fundthresh.Get(usdc).MinSwap()
	maxSwap := fundthresh.Get(usdc).MaxSwap()

	if bi.IsZero(maxSwap) { // Require max swap.
		log.Warn(ctx, "No max swap set, skipping", errors.New("missing max swap"), "token", usdc)
		return nil
	}

	if bi.LT(toSwap, minSwap) { // Surplus < minSwap, do nothing.
		log.Debug(ctx, "Surplus < minSwap, skipping", "amount", usdc.FormatAmt(toSwap), "min", usdc.FormatAmt(minSwap))
		return nil
	}

	if ok && bi.GT(toSwap, maxSwap) { // Cap swap to maxSwap.
		log.Debug(ctx, "Surplus > maxSwap, capping", "amount", usdc.FormatAmt(toSwap), "max", usdc.FormatAmt(maxSwap))
		toSwap = maxSwap
	}

	log.Debug(ctx, "Filling deficit", "deficit", token.FormatAmt(deficit), "usdc", usdc.FormatAmt(toSwap))

	if _, err = uniswap.SwapFromUSDC(ctx, backend, solver, token, toSwap); err != nil {
		return errors.Wrap(err, "swap from usdc")
	}

	return nil
}

func find[T any](xs []T, f func(T) bool) (T, bool) {
	for _, x := range xs {
		if f(x) {
			return x, true
		}
	}

	return *new(T), false
}
