package app

import (
	"context"
	"fmt"
	"math/big"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

const saneMaxEther = 5 // Maximum amount to fund in ether. // TODO(corver): Increase this.

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

// FundEOAAccounts funds the EOAs that need funding to their target balance.
func FundEOAAccounts(ctx context.Context, def Definition) error {
	network := networkFromDef(def)
	accounts, ok := eoa.AllAccounts(network.ID)
	if !ok {
		return errors.New("no accounts found", "network", network.ID)
	}

	for _, account := range accounts {
		if account.Address == common.HexToAddress(eoa.ZeroXDead) {
			log.Info(ctx, "Skipping 0xdead account", "role", account.Role)
			continue
		}

		for _, chain := range network.EVMChains() {
			thresholds, ok := eoa.GetFundThresholds(network.ID, account.Role)
			if !ok {
				log.Warn(ctx, "Skipping account without fund thresholds", nil, "role", account.Role)
				continue
			}

			backend, err := def.Backends().Backend(chain.ID)
			if err != nil {
				return errors.Wrap(err, "backend")
			}

			balance, err := backend.BalanceAt(ctx, account.Address, nil)
			if err != nil {
				// skip if we have rpc errors
				continue
			}

			if thresholds.MinBalance().Cmp(balance) < 0 {
				log.Info(ctx,
					"Not funding account, balance sufficient",
					"chain", chain.Name,
					"role", account.Role,
					"address", account.Address,
					"type", account.Type,
					"balance", etherStr(balance),
					"min_threshold", etherStr(thresholds.MinBalance()),
				)

				continue
			}

			saneMax := new(big.Int).Mul(big.NewInt(saneMaxEther), big.NewInt(params.Ether))

			amount := new(big.Int).Sub(thresholds.TargetBalance(), balance)
			if amount.Cmp(big.NewInt(0)) <= 0 {
				return errors.New("unexpected negative amount")
			} else if amount.Cmp(saneMax) > 0 {
				log.Warn(ctx, "Funding amount exceeds sane max, skipping", nil,
					"chain", chain.Name,
					"role", account.Role,
					"amount", etherStr(amount),
					"max", etherStr(saneMax),
				)

				continue
			}

			tx, _, err := backend.Send(ctx, eoa.Funder(), txmgr.TxCandidate{
				To:       &account.Address,
				GasLimit: 100_000,
				Value:    amount,
			})

			if err != nil {
				return errors.Wrap(err, "send tx")
			}

			b, err := backend.BalanceAt(ctx, account.Address, nil)
			if err != nil {
				return errors.Wrap(err, "get balance")
			}

			log.Info(ctx, "Account funded",
				"chain", chain.Name,
				"role", account.Role,
				"address", account.Address,
				"type", account.Type,
				"amount_funded", etherStr(amount),
				"resulting_balance", etherStr(b),
				"tx", tx.Hash().Hex(),
			)
		}
	}

	return nil
}

func etherStr(amount *big.Int) string {
	b, _ := amount.Float64()
	b /= params.Ether

	return fmt.Sprintf("%.4f", b)
}
