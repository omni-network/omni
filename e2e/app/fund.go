package app

import (
	"context"
	"fmt"
	"math/big"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const saneMaxETH = 121    // Maximum amount of ETH to fund (in ether).
const saneMaxOmni = 60420 // Maximum amount of OMNI to fund (in ether OMNI).

// FundAccounts funds the EOAs and contracts that need funding to their target balance.
func FundAccounts(ctx context.Context, def Definition, hotOnly bool, dryRun bool) error {
	network, err := SolverNetworkFromDef(ctx, def)
	if err != nil {
		return errors.Wrap(err, "get hl network")
	}

	endpoints := ExternalEndpoints(def)

	network, backends, err := AddSolverNetworkAndBackends(ctx, network, endpoints, def.Cfg, "fund")
	if err != nil {
		return errors.Wrap(err, "get solver network and backends")
	}

	var funderRole eoa.Role
	var accounts []eoa.Account
	if hotOnly {
		funder, ok := eoa.AccountForRole(network.ID, eoa.RoleHot)
		if !ok {
			return errors.New("hot wallet not found")
		}
		accounts = []eoa.Account{funder}
		funderRole = eoa.RoleCold
	} else {
		for _, account := range eoa.AllAccounts(network.ID) {
			if account.Role == eoa.RoleCold || account.Role == eoa.RoleHot {
				continue // Don't fund cold or hot
			}
			accounts = append(accounts, account)
		}
		funderRole = eoa.RoleHot
	}
	log.Info(ctx, "Checking accounts to fund", "network", network, "count", len(accounts))

	for _, chain := range network.EVMChains() {
		if evmchain.IsDisabled(chain.ID) {
			log.Warn(ctx, "Ignoring disabled chain", nil, "chain", chain.Name)
			continue
		}

		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "backend")
		}

		metadata, ok := evmchain.MetadataByID(chain.ID)
		if !ok {
			return errors.New("unknown chain", "chain", chain.Name)
		}

		funderAddr := eoa.MustAddress(network.ID, funderRole)
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
			if solvernet.SkipRole(chain.ID, account.Role) {
				log.Info(ctx, "Skipping non-solvernet role on HL chain", "chain", chain.Name, "role", account.Role, "address", account.Address)
				continue
			}

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

			thresholds, ok := eoa.GetFundThresholds(metadata.NativeToken, network.ID, account.Role)
			if !ok {
				log.Warn(accCtx, "Skipping account without fund thresholds", nil)
				continue
			}

			if err := fund(accCtx, fundParams{
				backend:       backend,
				account:       account.Address,
				minBalance:    thresholds.MinBalance(),
				targetBalance: thresholds.TargetBalance(),
				saneMax:       saneMax(metadata.NativeToken),
				dryRun:        dryRun,
				funder:        funderAddr,
			}); err != nil {
				return errors.Wrap(err, "fund account")
			}
		}

		if hotOnly {
			continue // Skip contract / sponsor funding if hotOnly.
		}

		for _, sponsor := range eoa.AllSponsors(network.ID) {
			if sponsor.ChainID != chain.ID {
				continue
			}

			sponsorCtx := log.WithCtx(ctx,
				"chain", chain.Name,
				"role", sponsor.Name,
				"address", sponsor.Address,
			)

			if err := fund(sponsorCtx, fundParams{
				backend:       backend,
				account:       sponsor.Address,
				minBalance:    sponsor.FundThresholds.MinBalance(),
				targetBalance: sponsor.FundThresholds.TargetBalance(),
				saneMax:       saneMax(metadata.NativeToken),
				dryRun:        dryRun,
				funder:        funderAddr,
			}); err != nil {
				return errors.Wrap(err, "fund sponsor")
			}
		}

		toFund, err := contracts.ToFund(ctx, network.ID)
		if err != nil {
			return errors.Wrap(err, "get contracts to fund")
		}

		log.Debug(ctx, "Checking contracts to fund", "chain", chain.Name, "count", len(toFund))

		for _, contract := range toFund {
			ctrCtx := log.WithCtx(ctx,
				"chain", chain.Name,
				"contract", contract.Name,
				"address", contract.Address,
			)

			if !contract.IsDeployedOn(chain.ID, network.ID) {
				log.Info(ctrCtx, "Skipping chain without deplyment", "chain", chain.ID, "contract", contract.Name)
				continue
			}

			if err := fund(ctrCtx, fundParams{
				backend:       backend,
				account:       contract.Address,
				minBalance:    contract.FundThresholds.MinBalance(),
				targetBalance: contract.FundThresholds.TargetBalance(),
				saneMax:       saneMax(metadata.NativeToken),
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
	} else if bi.LT(minBalance, balance) {
		log.Info(ctx,
			"Not funding account, balance sufficient",
			"balance", etherStr(balance),
			"min_balance", etherStr(minBalance),
		)

		return nil
	}

	amount := bi.Sub(targetBalance, balance)
	if amount.Sign() <= 0 {
		return errors.New("unexpected negative amount [BUG]") // Target balance below minimum balance
	} else if saneMax != nil && bi.GT(amount, saneMax) {
		log.Warn(ctx, "Funding amount exceeds sane max, skipping", nil,
			"amount", etherStr(amount),
			"max", etherStr(saneMax),
		)
	} else if bi.GTE(amount, funderBal) {
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
	return fmt.Sprintf("%.4f", bi.ToEtherF64(amount))
}

func saneMax(token tokens.Asset) *big.Int {
	saneETH := bi.Ether(saneMaxETH)
	saneOmni := bi.Ether(saneMaxOmni)

	if token == tokens.OMNI {
		return saneOmni
	}

	return saneETH
}
