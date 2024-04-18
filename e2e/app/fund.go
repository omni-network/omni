package app

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

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
		return []common.Address{
			eoa.MustAddress(netconf.Staging, eoa.RoleFbDev),
			eoa.MustAddress(netconf.Staging, eoa.RoleCreate3Deployer),
			eoa.MustAddress(netconf.Staging, eoa.RoleDeployer),
			eoa.MustAddress(netconf.Staging, eoa.RoleProxyAdminOwner),
			eoa.MustAddress(netconf.Staging, eoa.RolePortalAdmin),
			eoa.MustAddress(netconf.Staging, eoa.RoleAVSAdmin),
			eoa.MustAddress(netconf.Staging, eoa.RoleRelayer),
			eoa.MustAddress(netconf.Staging, eoa.RoleMonitor),
		}
	case netconf.Devnet:
		return []common.Address{
			eoa.MustAddress(netconf.Devnet, eoa.RoleFbDev),
			eoa.MustAddress(netconf.Devnet, eoa.RoleCreate3Deployer),
			eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer),
			eoa.MustAddress(netconf.Devnet, eoa.RoleProxyAdminOwner),
			eoa.MustAddress(netconf.Devnet, eoa.RolePortalAdmin),
			eoa.MustAddress(netconf.Devnet, eoa.RoleAVSAdmin),
			eoa.MustAddress(netconf.Devnet, eoa.RoleRelayer),
			eoa.MustAddress(netconf.Devnet, eoa.RoleMonitor),
		}
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
