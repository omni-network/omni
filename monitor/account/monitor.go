package account

import (
	"context"
	"time"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
)

// StartMonitoring starts the monitoring goroutines.
func StartMonitoring(ctx context.Context, network netconf.Network, rpcClients map[uint64]ethclient.Client) {
	accounts := eoa.AllAccounts(network.ID)
	sponsors := eoa.AllSponsors(network.ID)
	chains := network.EVMChains()
	log.Info(ctx, "Monitoring accounts", "accounts", len(accounts), "sponsors", len(sponsors), "chains", len(chains))

	for _, chain := range chains {
		for _, account := range accounts {
			go monitorAccountForever(ctx, network.ID, account, chain.Name, rpcClients[chain.ID])
		}

		for _, sponsor := range sponsors {
			if chain.ID != sponsor.ChainID {
				continue
			}

			go monitorSponsorForever(ctx, sponsor, chain.Name, rpcClients[chain.ID])
		}
	}
}

// monitorAccountForever blocks and periodically monitors the account for the given chain.
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

	nonce, err := client.NonceAt(ctx, account.Address, nil)
	if err != nil {
		return err
	}

	accountBalance.WithLabelValues(chainName, string(account.Role)).Set(umath.ToEtherF64(balance))
	accountNonce.WithLabelValues(chainName, string(account.Role)).Set(float64(nonce))

	meta, ok := evmchain.MetadataByName(chainName)
	if !ok {
		return errors.New("invalid chain name [BUG]", "name", chainName)
	}

	thresholds, ok := eoa.GetFundThresholds(meta.NativeToken, network, account.Role)
	if !ok {
		// Skip accounts without thresholds
		return nil
	}

	var isLow float64
	if umath.LTE(balance, thresholds.MinBalance()) {
		isLow = 1
	}

	accountBalanceLow.WithLabelValues(chainName, string(account.Role)).Set(isLow)

	return nil
}

// monitorSponsorForever blocks and periodically monitors the sponsor account for the given chain.
func monitorSponsorForever(
	ctx context.Context,
	sponsor eoa.Sponsor,
	chainName string,
	client ethclient.Client,
) {
	ctx = log.WithCtx(ctx,
		"chain", chainName,
		"role", sponsor.Name,
		"address", sponsor.Address,
	)

	log.Info(ctx, "Monitoring account")

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorSponsorOnce(ctx, sponsor, chainName, client)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Monitoring account failed (will retry)", err)

				continue
			}
		}
	}
}

// monitorSponsorOnce monitors sponsor account for the given chain.
func monitorSponsorOnce(
	ctx context.Context,
	sponsor eoa.Sponsor,
	chainName string,
	client ethclient.Client,
) error {
	balance, err := client.BalanceAt(ctx, sponsor.Address, nil)
	if err != nil {
		return err
	}
	// Convert to ether units
	balanceEth := umath.ToEtherF64(balance)

	nonce, err := client.NonceAt(ctx, sponsor.Address, nil)
	if err != nil {
		return err
	}

	accountBalance.WithLabelValues(chainName, sponsor.Name).Set(balanceEth)
	accountNonce.WithLabelValues(chainName, sponsor.Name).Set(float64(nonce))

	meta, ok := evmchain.MetadataByName(chainName)
	if !ok {
		return errors.New("invalid chain name [BUG]", "name", chainName)
	}

	if sponsor.ChainID != meta.ChainID {
		return errors.New("sponsor chain mismatch [BUG]", "expected", meta.ChainID, "got", sponsor.ChainID)
	}

	thresholds := sponsor.FundThresholds

	var isLow float64
	if umath.LTE(balance, thresholds.MinBalance()) {
		isLow = 1
	}

	accountBalanceLow.WithLabelValues(chainName, sponsor.Name).Set(isLow)

	return nil
}
