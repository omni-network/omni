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
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/uniswap"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// minSend is the minimum amount of surplus USDC to send to other chains.
	minSend = bi.Dec6(1000) // 1000 USDC

	// maxSend is the maximum amount of USDC allowed in a single send.
	maxSend = bi.Dec6(5000) // 5000 USDC
)

// rebalanceForever starts rebalancing loops for each chain in the network.
func rebalanceForever(
	ctx context.Context,
	interval time.Duration,
	db *cctpdb.DB,
	network netconf.Network,
	pricer tokenpricer.Pricer,
	backends ethbackend.Backends,
	solver common.Address,
) {
	for _, chain := range network.EVMChains() {
		ctx := log.WithCtx(ctx, "chain", evmchain.Name(chain.ID))

		go swapSurplusForever(ctx, interval, backends, solver, chain.ID)
		go sendSurplusForever(ctx, interval, db, network, backends, solver, chain.ID)
		go fillDeficitForever(ctx, interval, db, network, pricer, backends, solver, chain.ID)
	}
}

// swapSurplusForever swaps surplus tokens to USDC on `chainID` forever.
func swapSurplusForever(
	ctx context.Context,
	interval time.Duration,
	backends ethbackend.Backends,
	solver common.Address,
	chainID uint64,
) {
	ctx = log.WithCtx(ctx, "loop", "swapSurplus")

	ticker := time.NewTimer(0)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ticker.Reset(interval)

			do := func() error {
				return swapSurplusOnce(ctx, backends, chainID, solver)
			}

			if err := expbackoff.Retry(ctx, do); err != nil {
				log.Error(ctx, "Swap surplus failed", err)
			}
		}
	}
}

// sendSurplusForever sends surplus USDC on `chainID` to chains in deficit forever.
func sendSurplusForever(
	ctx context.Context,
	interval time.Duration,
	db *cctpdb.DB,
	network netconf.Network,
	backends ethbackend.Backends,
	solver common.Address,
	chainID uint64,
) {
	ctx = log.WithCtx(ctx, "loop", "sendSurplus")

	ticker := time.NewTimer(0)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ticker.Reset(interval)

			do := func() error {
				return sendSurplusOnce(ctx, db, network.ID, backends, chainID, solver)
			}

			if err := expbackoff.Retry(ctx, do); err != nil {
				log.Error(ctx, "Send surplus failed", err)
			}
		}
	}
}

// fillDeficitForever fills token deficits from surplus USDC on `chainID` forever.
func fillDeficitForever(
	ctx context.Context,
	interval time.Duration,
	db *cctpdb.DB,
	network netconf.Network,
	pricer tokenpricer.Pricer,
	backends ethbackend.Backends,
	solver common.Address,
	chainID uint64,
) {
	ctx = log.WithCtx(ctx, "loop", "fillDeficit")

	ticker := time.NewTimer(0)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ticker.Reset(interval)

			do := func() error {
				return fillDeficitOnce(ctx, db, network.ID, pricer, backends, chainID, solver)
			}

			if err := expbackoff.Retry(ctx, do); err != nil {
				log.Error(ctx, "Fill deficit failed", err)
			}
		}
	}
}

// sendSurplusOnce sends surplus USDC on `chainID` to chains in deficit.
func sendSurplusOnce(
	ctx context.Context,
	db *cctpdb.DB,
	networkID netconf.ID,
	backends ethbackend.Backends,
	chainID uint64,
	solver common.Address,
) error {
	if chainID != evmchain.IDBase {
		// Only sending surplus wstETH on Base.
		return nil
	}

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

	toSend := surplus
	if bi.GT(toSend, maxSend) { // Cap send to maxSend.
		log.Debug(ctx, "Surplus > maxSend, capping send", "amount", usdc.FormatAmt(toSend), "max", usdc.FormatAmt(maxSend))
		toSend = maxSend
	}

	log.Debug(ctx, "Sending surplus", "amount", usdc.FormatAmt(toSend))

	// Just send surplus to Ethereum for now.
	// TODO: calculate chain deficits, send to most in-need chain.

	if _, err = cctp.SendUSDC(ctx, db, networkID, backend, cctp.SendUSDCArgs{
		Sender:      solver,
		Recipient:   solver,
		SrcChainID:  chainID,
		DestChainID: evmchain.IDEthereum,
		Amount:      toSend,
	}); err != nil {
		return errors.Wrap(err, "send usdc")
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
	backend, err := backends.Backend(chainID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	for _, token := range TokensByChain(chainID) {
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

	maxSwap := GetFundThreshold(token).MaxSwap()
	minSwap := GetFundThreshold(token).MinSwap()

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
	db *cctpdb.DB,
	networkID netconf.ID,
	pricer tokenpricer.Pricer,
	backends ethbackend.Backends,
	chainID uint64,
	solver common.Address,
) error {
	// will be used when calculating deficits
	_ = db
	_ = networkID

	backend, err := backends.Backend(chainID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	for _, token := range TokensByChain(chainID) {
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

	usdc, ok := tokens.ByAsset(chainID, tokens.USDC)
	if !ok {
		return errors.New("token not found")
	}

	surplusUSDC, err := GetSurplus(ctx, backend, usdc, solver)
	if err != nil {
		return errors.Wrap(err, "get surplus")
	}

	price, err := pricer.USDPrice(ctx, token.Asset)
	if err != nil {
		return errors.Wrap(err, "get price")
	}

	// Use USD deficit to inform swap input.
	deficitUSD := bi.MulF64(deficit, price)

	toSwap := deficitUSD
	if bi.GT(toSwap, surplusUSDC) { // Deficit > surplus, cap swap to surplus.
		log.Debug(ctx, "Deficit > surplus, capping swap", "deficit", usdc.FormatAmt(deficitUSD), "surplus", usdc.FormatAmt(surplusUSDC))
		toSwap = surplusUSDC
	}

	minSwap := GetFundThreshold(usdc).MinSwap()
	maxSwap := GetFundThreshold(usdc).MaxSwap()

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
