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
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

const saneEOAMaxEther = 50       // Maximum amount of ETH to fund an eoe (in ether).
const saneContractMaxEther = 50  // Maximum amount of ETH to fund a contract (in ether).
const saneContractMaxOMNI = 5000 // Maximum amount of OMNI to fund a contract (in OMNI).

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

// fundAccounts funds the EOAs that need funding (just on anvil chains, for now).
func fundAccounts(ctx context.Context, def Definition) error {
	accounts := accountsToFund(def.Testnet.Network)
	eth100 := new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(100))
	for _, chain := range def.Testnet.AnvilChains {
		if err := anvil.FundAccounts(ctx, chain.ExternalRPC, eth100, noAnvilDev(accounts)...); err != nil {
			return errors.Wrap(err, "fund anvil account")
		}
	}

	return nil
}

// FundAccounts funds the EOAs and contracts that need funding to their target balance.
func FundAccounts(ctx context.Context, def Definition, dryRun bool) error {
	network := NetworkFromDef(def)
	accounts := eoa.AllAccounts(network.ID)

	log.Info(ctx, "Checking accounts to fund", "network", network.ID, "count", len(accounts))

	for _, chain := range network.EVMChains() {
		backend, err := def.Backends().Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "backend")
		}

		funder := eoa.MustAddress(network.ID, eoa.RoleFunder)
		funderBal, err := backend.BalanceAt(ctx, funder, nil)
		if err != nil {
			return err
		}

		log.Info(ctx, "Funder balance",
			"chain", chain.Name,
			"funder", funder,
			"balance", etherStr(funderBal),
		)

		for _, account := range accounts {
			accCtx := log.WithCtx(ctx,
				"chain", chain.Name,
				"role", account.Role,
				"address", account.Address,
				"type", account.Type,
			)

			if account.Address == common.HexToAddress(eoa.ZeroXDead) {
				log.Info(accCtx, "Skipping 0xdead account")
				continue
			} else if account.Type == eoa.TypeWellKnown {
				log.Info(accCtx, "Skipping well-known anvil account")
				continue
			}

			thresholds, ok := eoa.GetFundThresholds(network.ID, account.Role)
			if !ok {
				log.Warn(accCtx, "Skipping account without fund thresholds", nil)
				continue
			}

			saneMax := new(big.Int).Mul(big.NewInt(saneEOAMaxEther), big.NewInt(params.Ether))

			if err := fund(accCtx, fundParams{
				backend:       backend,
				account:       account.Address,
				minBalance:    thresholds.MinBalance(),
				targetBalance: thresholds.TargetBalance(),
				saneMax:       saneMax,
				dryRun:        dryRun,
				funder:        funder,
			}); err != nil {
				return errors.Wrap(err, "fund account")
			}
		}

		toFund, err := contracts.ToFund(ctx, network.ID)
		if err != nil {
			return errors.Wrap(err, "get contracts to fund")
		}

		for _, contract := range toFund {
			ctrCtx := log.WithCtx(ctx,
				"chain", chain.Name,
				"contract", contract.Name,
				"address", contract.Address,
			)

			isOmniEVM := chain.ID == network.ID.Static().OmniExecutionChainID

			if contract.OnlyOmniEVM && !isOmniEVM {
				log.Info(ctrCtx, "Skipping non-OmniEVM chain", "chain", chain.ID)
				continue
			}

			saneMax := new(big.Int).Mul(big.NewInt(saneContractMaxEther), big.NewInt(params.Ether))
			if isOmniEVM {
				saneMax = new(big.Int).Mul(big.NewInt(saneContractMaxOMNI), big.NewInt(params.Ether))
			}

			if err := fund(ctrCtx, fundParams{
				backend:       backend,
				account:       contract.Address,
				minBalance:    contract.Thresholds.MinBalance(),
				targetBalance: contract.Thresholds.TargetBalance(),
				saneMax:       saneMax,
				dryRun:        dryRun,
				funder:        funder,
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
