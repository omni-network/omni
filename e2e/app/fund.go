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
	network := externalNetwork(def)
	accounts, ok := eoa.AllAccounts(network.ID)
	if !ok {
		return errors.New("no accounts found", "network", network.ID)
	}

	for _, account := range accounts {
		for _, chain := range account.Chains(network) {
			backend, err := def.Backends().Backend(chain.ID)
			if err != nil {
				return errors.Wrap(err, "backend")
			}

			balance, err := backend.BalanceAt(ctx, account.Address, nil)
			if err != nil {
				// skip if we have rpc errors
				continue
			}

			bf, _ := balance.Float64()
			bf /= params.Ether

			fund := account.MinBalance.Cmp(balance) > 0

			log.Info(ctx,
				"Account",
				"address", account.Address,
				"type", account.Type,
				"balance", fmt.Sprintf("%.2f ETH", bf),
				"funding", fund,
			)

			if fund {
				continue
			}

			target := new(big.Int).Sub(account.TargetBalance, balance)
			if target.Cmp(big.NewInt(0)) <= 0 {
				continue
			}

			tx, _, err := backend.Send(ctx, eoa.Funder(), txmgr.TxCandidate{
				To:       &account.Address,
				GasLimit: 100_000,
				Value:    target,
			})

			if err != nil {
				return errors.Wrap(err, "send tx")
			} else if _, err := backend.WaitMined(ctx, tx); err != nil {
				return errors.Wrap(err, "wait mined")
			}

			b, err := backend.EtherBalanceAt(ctx, account.Address)
			if err != nil {
				return errors.Wrap(err, "get balance")
			}

			log.Info(ctx, "Account funded",
				"address", account.Address,
				"type", account.Type,
				"balance", fmt.Sprintf("%.2f ETH", b),
			)
		}
	}

	return nil
}
