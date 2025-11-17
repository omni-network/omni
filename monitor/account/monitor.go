package account

import (
	"context"
	"time"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	solver "github.com/omni-network/omni/solver/app"

	"github.com/ethereum/go-ethereum/common"
)

// StartMonitoring starts the monitoring goroutines.
func StartMonitoring(ctx context.Context, network netconf.Network, rpcClients map[uint64]ethclient.Client) error {
	accounts := eoa.AllAccounts(network.ID)
	chains := network.EVMChains()
	log.Info(ctx, "Monitoring accounts", "accounts", len(accounts), "chains", len(chains))

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
			if solvernet.SkipRole(chain.ID, account.Role) {
				// Do not monitor non-HL roles on HL-only chains
				continue
			}

			go monitorAccountForever(ctx, network.ID, account, chain.Name, rpcClients[chain.ID])

			if isSolverNetRole(account.Role) {
				solverCtx := log.WithCtx(ctx, "chain", chain.Name, "role", account.Role)
				go monitorSolverNetRoleForever(solverCtx, network.ID, backend, account.Role, account.Address)
			}
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
	if bi.LT(balance, thresholds.MinBalance()) {
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
			for _, token := range tokens.UniqueAssets() {
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
	meta tokens.Asset,
	role eoa.Role,
	address common.Address,
) error {
	chainName, chainID := backend.Chain()
	token, ok := tokens.BySymbol(chainID, meta.Symbol)
	if !ok {
		// Not all tokens are present on all chains.
		return nil
	}

	thresh, ok := eoa.GetSolverNetThreshold(role, network, chainID, meta)
	if !ok {
		// Thresholds only exist for solver and flowgen role on dest fill chains
		return nil
	}

	if !bi.GT(thresh.MinBalance(), bi.Zero()) && !solver.IsSupportedToken(token) {
		// No need to monitor unsupported tokens.
		// Still monitor "unsupported" tokens with min balance > 0 (needed for gas).
		return nil
	}

	balance, err := tokenutil.BalanceOf(ctx, backend, token, address)
	if err != nil {
		return err
	}

	// Convert to float64 ether
	balF64 := token.AmtToF64(balance)
	tokenBalance.WithLabelValues(chainName, string(role), meta.Symbol).Set(balF64)

	var isLow float64
	if bi.LT(balance, thresh.MinBalance()) {
		isLow = 1
	}

	tokenBalanceLow.WithLabelValues(chainName, string(role), meta.Symbol).Set(isLow)

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
