package account

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

// startMonitoring starts the monitoring goroutines.
func startMonitoring(ctx context.Context, network netconf.Network,
	addresses []common.Address, rpcClients map[uint64]ethclient.Client) {
	for _, srcChain := range network.Chains {
		if srcChain.IsOmniConsensus {
			continue // Below monitors only apply to EVM chains.
		}

		for _, addr := range addresses {
			go monitorAccountForever(ctx, addr, srcChain.Name, rpcClients[srcChain.ID])
		}
	}
}

// monitorAccountsForever blocks and periodically monitors the relayer accounts for the given chain.
func monitorAccountForever(ctx context.Context, addr common.Address, chainName string, client ethclient.Client) {
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

// monitorAccountOnce monitors the relayer account for the given chain.
func monitorAccountOnce(ctx context.Context, addr common.Address, chainName string, client ethclient.Client) error {
	balance, err := client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "balance at")
	}

	nonce, err := client.NonceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "nonce at")
	}

	bf, _ := balance.Float64()
	bf /= params.Ether

	accountBalance.WithLabelValues(chainName).Set(bf)
	accountNonce.WithLabelValues(chainName).Set(float64(nonce))

	return nil
}
