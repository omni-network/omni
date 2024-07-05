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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

const saneMaxEther = 20 // Maximum amount to fund in ether. // TODO(corver): Increase this.

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
func FundEOAAccounts(ctx context.Context, def Definition, dryRun bool) error {
	if def.Testnet.Network == netconf.Mainnet {
		return errors.New("mainnet funding not supported yet")
	}

	network := networkFromDef(def)
	accounts := eoa.AllAccounts(network.ID)

	log.Info(ctx, "Checking accounts to fund", "network", network.ID, "count", len(accounts))

	for _, chain := range network.EVMChains() {
		backend, err := def.Backends().Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "backend")
		}

		funder := eoa.Funder()
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

			balance, err := backend.BalanceAt(accCtx, account.Address, nil)
			if err != nil {
				log.Warn(accCtx, "Failed fetching balance, skipping", err)
				continue
			} else if thresholds.MinBalance().Cmp(balance) < 0 {
				log.Info(accCtx,
					"Not funding account, balance sufficient",
					"balance", etherStr(balance),
					"min_balance", etherStr(thresholds.MinBalance()),
				)

				continue
			}

			saneMax := new(big.Int).Mul(big.NewInt(saneMaxEther), big.NewInt(params.Ether))

			amount := new(big.Int).Sub(thresholds.TargetBalance(), balance)
			if amount.Cmp(big.NewInt(0)) <= 0 {
				return errors.New("unexpected negative amount [BUG]") // Target balance below minimum balance
			} else if amount.Cmp(saneMax) > 0 {
				log.Warn(accCtx, "Funding amount exceeds sane max, skipping", nil,
					"amount", etherStr(amount),
					"max", etherStr(saneMax),
				)

				continue
			} else if amount.Cmp(funderBal) >= 0 {
				return errors.New("funder balance too low",
					"amount", etherStr(amount),
					"funder", etherStr(funderBal),
				)
			}

			log.Info(accCtx, "Funding account",
				"amount", etherStr(amount),
				"balance", etherStr(balance),
				"target_balance", etherStr(thresholds.TargetBalance()),
			)

			if dryRun {
				log.Warn(accCtx, "Skipping actual funding tx due to dry-run", nil)
				continue
			}

			tx, rec, err := backend.Send(accCtx, eoa.Funder(), txmgr.TxCandidate{
				To:       &account.Address,
				GasLimit: 0,
				Value:    amount,
			})
			if err != nil {
				return errors.Wrap(err, "send tx")
			} else if rec.Status != types.ReceiptStatusSuccessful {
				return errors.New("funding tx failed", "tx", tx.Hash())
			}

			b, err := backend.BalanceAt(accCtx, account.Address, nil)
			if err != nil {
				return errors.Wrap(err, "get balance")
			}

			log.Info(accCtx, "Account funded 🎉",
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
