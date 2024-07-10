package account

import (
	"context"
	"time"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/params"
)

// StartMonitoring starts the monitoring goroutines.
func StartMonitoring(ctx context.Context, network netconf.Network, rpcClients map[uint64]ethclient.Client) {
	accounts := eoa.AllAccounts(network.ID)
	chains := network.EVMChains()
	log.Info(ctx, "Monitoring accounts", "accounts", len(accounts), "chains", len(chains))

	for _, chain := range chains {
		for _, account := range accounts {
			go monitorAccountForever(ctx, network.ID, account, chain.Name, rpcClients[chain.ID])
		}
	}
}

// monitorAccountsForever blocks and periodically monitors the account for the given chain.
func monitorAccountForever(
	ctx context.Context,
	network netconf.ID,
	account eoa.Account,
	chainName string,
	client ethclient.Client,
) {
	ctx = log.WithCtx(ctx,
		"chain", chainName,
		"role", account.Role,
		"address", account.Address,
	)

	log.Info(ctx, "Monitoring account")

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorAccountOnce(ctx, network, account, chainName, client)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Monitoring account failed (will retry)", err)

				continue
			}
		}
	}
}

// monitorAccountOnce monitors account for the given chain.
func monitorAccountOnce(
	ctx context.Context,
	network netconf.ID,
	account eoa.Account,
	chainName string,
	client ethclient.Client,
) error {
	balance, err := client.BalanceAt(ctx, account.Address, nil)
	if err != nil {
		return err
	}
	// Convert to ether units
	bf, _ := balance.Float64()
	balanceEth := bf / params.Ether

	nonce, err := client.NonceAt(ctx, account.Address, nil)
	if err != nil {
		return err
	}

	accountBalance.WithLabelValues(chainName, string(account.Role)).Set(balanceEth)
	accountNonce.WithLabelValues(chainName, string(account.Role)).Set(float64(nonce))

	thresholds, ok := eoa.GetFundThresholds(network, account.Role)
	if !ok {
		// Skip accounts without thresholds
		return nil
	}

	var isLow float64
	if balance.Cmp(thresholds.MinBalance()) <= 0 {
		isLow = 1
	}

	accountBalanceLow.WithLabelValues(chainName, string(account.Role)).Set(isLow)

	return nil
}
