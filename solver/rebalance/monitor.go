package rebalance

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/solver/fundthresh"

	"github.com/ethereum/go-ethereum/common"
)

func monitorForever(
	ctx context.Context,
	network netconf.Network,
	clients map[uint64]ethclient.Client,
	solver common.Address,
) error {
	for _, chain := range network.EVMChains() {
		client, ok := clients[chain.ID]
		if !ok {
			return errors.New("missing client", "chain_id", chain.ID)
		}

		go monitorChainForever(ctx, chain.ID, client, solver)
	}

	return nil
}

func monitorChainForever(
	ctx context.Context,
	chainID uint64,
	client ethclient.Client,
	solver common.Address,
) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := monitorChainOnce(ctx, chainID, client, solver); err != nil {
				log.Warn(ctx, "Failed monitoring chain, will retry", err, "chain")
			}
		}
	}
}

func monitorChainOnce(
	ctx context.Context,
	chainID uint64,
	client ethclient.Client,
	solver common.Address,
) error {
	chainName := evmchain.Name(chainID)

	for _, token := range tokens.ByChain(chainID) {
		thresh := fundthresh.Get(token)

		balance, err := tokenutil.BalanceOf(ctx, client, token, solver)
		if err != nil {
			return errors.Wrap(err, "get balance")
		}

		surplus, err := GetSurplus(ctx, client, token, solver)
		if err != nil {
			return errors.Wrap(err, "get surplus")
		}

		deficit, err := GetDeficit(ctx, client, token, solver)
		if err != nil {
			return errors.Wrap(err, "get deficit")
		}

		balanceDecifit.WithLabelValues(chainName, token.Asset.String()).Set(bi.ToF64(deficit, token.Decimals))
		balanceSurplus.WithLabelValues(chainName, token.Asset.String()).Set(bi.ToF64(surplus, token.Decimals))
		balanceCurrent.WithLabelValues(chainName, token.Asset.String()).Set(bi.ToF64(balance, token.Decimals))

		thresholdTarget.WithLabelValues(chainName, token.Asset.String()).Set(bi.ToF64(thresh.Target(), token.Decimals))
		thresholdMin.WithLabelValues(chainName, token.Asset.String()).Set(bi.ToF64(thresh.Min(), token.Decimals))

		// only gauge surplus if it is not infinite
		if !thresh.NeverSurplus() {
			thresholdSurplus.WithLabelValues(chainName, token.Asset.String()).Set(bi.ToF64(thresh.Surplus(), token.Decimals))
		}
	}

	return nil
}
