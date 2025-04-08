package account

import (
	"context"
	"time"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokenmeta"
	"github.com/omni-network/omni/lib/tokens"
	solver "github.com/omni-network/omni/solver/app"

	"github.com/ethereum/go-ethereum/common"
)

// StartMonitoring starts the monitoring goroutines.
func StartMonitoring(ctx context.Context, network netconf.Network, rpcClients map[uint64]ethclient.Client) error {
	accounts := eoa.AllAccounts(network.ID)
	sponsors := eoa.AllSponsors(network.ID)
	chains := network.EVMChains()
	log.Info(ctx, "Monitoring accounts", "accounts", len(accounts), "sponsors", len(sponsors), "chains", len(chains))

	for _, chain := range chains {
		meta, ok := evmchain.MetadataByID(chain.ID)
		if !ok {
			return errors.New("chain metadata not found", "chain", chain.ID)
		}
		backend, err := ethbackend.NewBackend(chain.Name, chain.ID, meta.BlockPeriod, rpcClients[chain.ID])
		if err != nil {
			return errors.Wrap(err, "new backend")
		}

		for _, account := range accounts {
			go monitorAccountForever(ctx, network.ID, account, chain.Name, rpcClients[chain.ID])

			if isSolverNetRole(account.Role) {
				solverCtx := log.WithCtx(ctx, "chain", chain.Name, "role", account.Role)
				go monitorSolverNetRoleForever(solverCtx, network.ID, backend, account.Role, account.Address)
			}
		}

		for _, sponsor := range sponsors {
			if chain.ID != sponsor.ChainID {
				continue
			}

			go monitorSponsorForever(ctx, sponsor, chain.Name, rpcClients[chain.ID])
		}

		// Also monitor solvernet claimant balances
		// These claimants should maybe be added as proper roles and added to isSolverNetRole
		for _, token := range tokens.UniqueMetas() {
			claimantAddress, ok := solver.Claimant(network.ID, token)
			if !ok {
				continue
			}

			claimantRole := eoa.Role("claimant")
			solverCtx := log.WithCtx(ctx, "chain", chain.Name, "role", claimantRole)
			go monitorSolverNetRoleForever(solverCtx, network.ID, backend, claimantRole, claimantAddress)
		}
	}

	return nil
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

	accountBalance.WithLabelValues(chainName, string(account.Role)).Set(bi.ToEtherF64(balance))
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
	if bi.LTE(balance, thresholds.MinBalance()) {
		isLow = 1
	}

	accountBalanceLow.WithLabelValues(chainName, string(account.Role)).Set(isLow)

	return nil
}

// monitorSolverNetRoleTokenOnce monitors solvernet role for provided token and chain.
func monitorSolverNetRoleForever(
	ctx context.Context,
	network netconf.ID,
	backend *ethbackend.Backend,
	role eoa.Role,
	address common.Address,
) {
	log.Info(ctx, "Monitoring solvernet role")

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			for _, token := range tokens.UniqueMetas() {
				loopCtx := log.WithCtx(ctx, "token", token)
				err := monitorSolverNetRoleTokenOnce(loopCtx, network, backend, token, role, address)
				if ctx.Err() != nil {
					return
				} else if err != nil {
					log.Warn(loopCtx, "Monitoring solvernet role token failed (will retry)", err)
				}
			}
		}
	}
}

// monitorSolverNetRoleTokenOnce monitors solvernet role for provided token and chain.
func monitorSolverNetRoleTokenOnce(
	ctx context.Context,
	network netconf.ID,
	backend *ethbackend.Backend,
	meta tokenmeta.Meta,
	role eoa.Role,
	address common.Address,
) error {
	chainName, chainID := backend.Chain()
	token, ok := tokens.BySymbol(chainID, meta.Symbol)
	if !ok {
		// Not all tokens are present on all chains.
		return nil
	}

	if !solver.IsSupportedToken(token) {
		// No need to monitor unsupported tokens.
		return nil
	}

	balance, err := tokens.BalanceOf(ctx, backend, token, address)
	if err != nil {
		return err
	}

	// Convert to float64 ether
	balF64 := token.AmtToF64(balance)
	tokenBalance.WithLabelValues(chainName, string(role), meta.Symbol).Set(balF64)

	thresh, ok := eoa.GetSolverNetThreshold(role, network, chainID, meta)
	if !ok {
		// Thresholds only exist for solver and flowgen role on dest fill chains
		return nil
	}

	var isLow float64
	if bi.LTE(balance, thresh.MinBalance()) {
		isLow = 1
	}

	tokenBalanceLow.WithLabelValues(chainName, string(role), meta.Symbol).Set(isLow)

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
	balanceEth := bi.ToEtherF64(balance)

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
	if bi.LTE(balance, thresholds.MinBalance()) {
		isLow = 1
	}

	accountBalanceLow.WithLabelValues(chainName, sponsor.Name).Set(isLow)

	return nil
}

func isSolverNetRole(role eoa.Role) bool {
	for _, r := range eoa.SolverNetRoles() {
		if role == r {
			return true
		}
	}

	return false
}
