package account

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/params"
)

// startMonitoring starts the monitoring goroutines.
func startMonitoring(ctx context.Context, network netconf.Network,
	addresses []accountType, rpcClients map[uint64]ethclient.Client) {
	for _, chain := range network.Chains {
		if chain.IsOmniConsensus {
			continue // skip non-EVM chains.
		}

		for _, at := range addresses {
			go monitorAccountForever(ctx, at, chain.Name, rpcClients[chain.ID])
		}
	}
}

// monitorAccountsForever blocks and periodically monitors accounts for the given chain.
func monitorAccountForever(ctx context.Context, addr accountType, chainName string, client ethclient.Client) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorAccountOnce(ctx, addr, chainName, client)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Monitoring account failed (will retry)", err,
					"chain", chainName)

				continue
			}
		}
	}
}

// monitorAccountOnce monitors account for the given chain.
func monitorAccountOnce(ctx context.Context, at accountType, chainName string, client ethclient.Client) error {
	balance, err := client.BalanceAt(ctx, at.addr, nil)
	if err != nil {
		return errors.Wrap(err, "balance at")
	}

	nonce, err := client.NonceAt(ctx, at.addr, nil)
	if err != nil {
		return errors.Wrap(err, "nonce at")
	}

	bf, _ := balance.Float64()
	bf /= params.Ether

	accountBalance.WithLabelValues(chainName, string(at.addressType)).Set(bf)
	accountNonce.WithLabelValues(chainName, string(at.addressType)).Set(float64(nonce))

	return nil
}
