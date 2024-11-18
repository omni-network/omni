package app

import (
	"context"
	"fmt"
	"math/big"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

const saneMaxETH = 113    // Maximum amount of ETH to fund (in ether).
const saneMaxOmni = 56630 // Maximum amount of OMNI to fund (in ether OMNI).

// noAnvilDev returns a list of accounts that are not dev anvil accounts.
func noAnvilDev(accounts []common.Address) []common.Address {
	var nonDevAccounts []common.Address
	for _, account := range accounts {
		if !anvil.IsDevAccount(account) {
			nonDevAccounts = append(nonDevAccounts, account)
		}
	}

	return nonDevAccounts
}

// accountsToFund returns a list of accounts to fund on anvil chains, based on the network.
func accountsToFund(network netconf.ID) []common.Address {
	switch network {
	case netconf.Staging:
		return eoa.MustAddresses(netconf.Staging, eoa.AllRoles()...)
	case netconf.Devnet:
		return eoa.MustAddresses(netconf.Devnet, eoa.AllRoles()...)
	default:
		return []common.Address{}
	}
}

// fundAnvilAccounts funds the EOAs on anvil that need funding.
func fundAnvilAccounts(ctx context.Context, def Definition) error {
	accounts := accountsToFund(def.Testnet.Network)
	for _, chain := range def.Testnet.AnvilChains {
		if err := anvil.FundAccounts(ctx, chain.ExternalRPC, saneMax(tokens.ETH), noAnvilDev(accounts)...); err != nil {
			return errors.Wrap(err, "fund anvil account")
		}
	}

	return nil
}

// FundAccounts funds the EOAs and contracts that need funding to their target balance.
func FundAccounts(ctx context.Context, def Definition, hotOnly bool, dryRun bool) error {
	network := def.Testnet.Network

	var funderRole eoa.Role
	var accounts []eoa.Account
	if hotOnly {
		funder, ok := eoa.AccountForRole(network, eoa.RoleHot)
		if !ok {
			return errors.New("hot wallet not found")
		}
		accounts = []eoa.Account{funder}
		funderRole = eoa.RoleCold
	} else {
		for _, account := range eoa.AllAccounts(network) {
			if account.Role == eoa.RoleCold || account.Role == eoa.RoleHot {
				continue // Don't fund cold or hot
			}
			accounts = append(accounts, account)
		}
		funderRole = eoa.RoleHot
	}
	log.Info(ctx, "Checking accounts to fund", "network", network, "count", len(accounts))

	for _, chain := range def.Testnet.EVMChains() {
		backend, err := def.Backends().Backend(chain.ChainID)
		if err != nil {
			return errors.Wrap(err, "backend")
		}

		funderAddr := eoa.MustAddress(network, funderRole)
		funderBal, err := backend.BalanceAt(ctx, funderAddr, nil)
		if err != nil {
			return err
		}

		log.Info(ctx, "Funder account balance",
			"chain", chain.Name,
			"role", funderRole,
			"balance", etherStr(funderBal),
			"address", funderAddr,
		)

		for _, account := range accounts {
			accCtx := log.WithCtx(ctx,
				"chain", chain.Name,
				"role", account.Role,
				"address", account.Address,
				"type", account.Type,
			)

			if account.Type == eoa.TypeWellKnown {
				log.Info(accCtx, "Skipping well-known anvil account")
				continue
			}

			thresholds, ok := eoa.GetFundThresholds(chain.NativeToken, network, account.Role)
			if !ok {
				log.Warn(accCtx, "Skipping account without fund thresholds", nil)
				continue
			}

			if err := fund(accCtx, fundParams{
				backend:       backend,
				account:       account.Address,
				minBalance:    thresholds.MinBalance(),
				targetBalance: thresholds.TargetBalance(),
				saneMax:       saneMax(chain.NativeToken),
				dryRun:        dryRun,
				funder:        funderAddr,
			}); err != nil {
				return errors.Wrap(err, "fund account")
			}
		}

		if hotOnly {
			continue // Skip contract funding if hotOnly.
		}

		toFund, err := contracts.ToFund(ctx, network)
		if err != nil {
			return errors.Wrap(err, "get contracts to fund")
		}

		log.Info(ctx, "Checking contracts to fund", "network", network, "count", len(toFund))

		for _, contract := range toFund {
			ctrCtx := log.WithCtx(ctx,
				"chain", chain.Name,
				"contract", contract.Name,
				"address", contract.Address,
			)

			isOmniEVM := chain.ChainID == def.Testnet.Network.Static().OmniExecutionChainID

			if contract.OnlyOmniEVM && !isOmniEVM {
				log.Info(ctrCtx, "Skipping non-OmniEVM chain", "chain", chain.ChainID)
				continue
			}

			if err := fund(ctrCtx, fundParams{
				backend:       backend,
				account:       contract.Address,
				minBalance:    contract.Thresholds.MinBalance(),
				targetBalance: contract.Thresholds.TargetBalance(),
				saneMax:       saneMax(chain.NativeToken),
				dryRun:        dryRun,
				funder:        funderAddr,
			}); err != nil {
				return errors.Wrap(err, "fund contract")
			}
		}
	}

	return nil
}

type fundParams struct {
	backend       *ethbackend.Backend
	funder        common.Address
	account       common.Address
	minBalance    *big.Int
	targetBalance *big.Int
	saneMax       *big.Int
	dryRun        bool
}

func fund(ctx context.Context, params fundParams) error {
	backend := params.backend
	account := params.account
	minBalance := params.minBalance
	targetBalance := params.targetBalance
	saneMax := params.saneMax
	dryRun := params.dryRun
	funder := params.funder

	funderBal, err := backend.BalanceAt(ctx, funder, nil)
	if err != nil {
		log.Warn(ctx, "Failed fetching balance, skipping", err)
		return nil
	}

	balance, err := backend.BalanceAt(ctx, account, nil)
	if err != nil {
		log.Warn(ctx, "Failed fetching balance, skipping", err)
		return nil
	} else if minBalance.Cmp(balance) < 0 {
		log.Info(ctx,
			"Not funding account, balance sufficient",
			"balance", etherStr(balance),
			"min_balance", etherStr(minBalance),
		)

		return nil
	}

	amount := new(big.Int).Sub(targetBalance, balance)
	if amount.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("unexpected negative amount [BUG]") // Target balance below minimum balance
	} else if saneMax != nil && amount.Cmp(saneMax) > 0 {
		log.Warn(ctx, "Funding amount exceeds sane max, skipping", nil,
			"amount", etherStr(amount),
			"max", etherStr(saneMax),
		)
	} else if amount.Cmp(funderBal) >= 0 {
		return errors.New("funder balance too low",
			"amount", etherStr(amount),
			"funder", etherStr(funderBal),
		)
	}

	log.Info(ctx, "Funding account",
		"amount", etherStr(amount),
		"balance", etherStr(balance),
		"target_balance", etherStr(targetBalance),
	)

	if dryRun {
		log.Warn(ctx, "Skipping actual funding tx due to dry-run", nil)
		return nil
	}

	tx, rec, err := backend.Send(ctx, funder, txmgr.TxCandidate{
		To:       &account,
		GasLimit: 0,
		Value:    amount,
	})
	if err != nil {
		return errors.Wrap(err, "send tx")
	} else if rec.Status != types.ReceiptStatusSuccessful {
		return errors.New("funding tx failed", "tx", tx.Hash())
	}

	b, err := backend.BalanceAt(ctx, account, nil)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	log.Info(ctx, "Account funded 🎉",
		"amount_funded", etherStr(amount),
		"resulting_balance", etherStr(b),
		"tx", tx.Hash().Hex(),
	)

	return nil
}

func etherStr(amount *big.Int) string {
	b, _ := amount.Float64()
	b /= params.Ether

	return fmt.Sprintf("%.4f", b)
}

func saneMax(token tokens.Token) *big.Int {
	saneETH := new(big.Int).Mul(big.NewInt(saneMaxETH), big.NewInt(params.Ether))
	saneOmni := new(big.Int).Mul(big.NewInt(saneMaxOmni), big.NewInt(params.Ether))

	if token == tokens.OMNI {
		return saneOmni
	}

	return saneETH
}
