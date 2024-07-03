package eoa

import (
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/netconf"
)

const (
	// fbFunder is the address of the fireblocks "funder" account.
	fbFunder  = "0xf63316AA39fEc9D2109AB0D9c7B1eE3a6F60AEA4"
	ZeroXDead = "0x000000000000000000000000000000000000dead"
)

//nolint:gochecknoglobals // Static mappings.
var statics = map[netconf.ID][]Account{
	netconf.Devnet: flatten(
		wellKnown(anvil.DevPrivateKey0(), RoleCreate3Deployer, RoleDeployer, RoleAdmin),
		wellKnown(anvil.DevPrivateKey5(), RoleRelayer),
		wellKnown(anvil.DevPrivateKey6(), RoleMonitor),
		wellKnown(anvil.DevPrivateKey7(), RoleTester),
	),
	netconf.Staging: flatten(
		remote("0x4891925c4f13A34FC26453FD168Db80aF3273014", RoleAdmin),
		remote("0xC8103859Ac7CB547d70307EdeF1A2319FC305fdC", RoleCreate3Deployer),
		remote("0x274c4B3e5d27A65196d63964532366872F81D261", RoleDeployer),
		remote("0x7a6cF389082dc698285474976d7C75CAdE08ab7e", RoleTester), // Legacy fbdev account
		secret("0xfE921e06Ed0a22c035b4aCFF0A5D3a434A330c96", RoleRelayer),
		secret("0x0De553555Fa19d787Af4273B18bDB77282D618c4", RoleMonitor),
	),
	netconf.Omega: flatten(
		remote("0xEAD625eB2011394cdD739E91Bf9D51A7169C22F5", RoleAdmin),
		remote("0xeC5134556da0797A5C5cD51DD622b689Cac97Fe9", RoleCreate3Deployer),
		remote("0x0CdCc644158b7D03f40197f55454dc7a11Bd92c1", RoleDeployer),
		remote("0x7a6cF389082dc698285474976d7C75CAdE08ab7e", RoleTester), // Identical address to staging. Concurrent usage will result in nonce clashes.
		secret("0x37AD6f7267454cac494C177127aC017750c8A7DB", RoleRelayer),
		secret("0xcef2a2c477Ec8473E4DeB9a8c2DF1D0585ea1040", RoleMonitor),
	),
	netconf.Mainnet: flatten(
		dummy(RoleAdmin, RoleCreate3Deployer, RoleDeployer, RoleTester),
		secret("0x07804D7B8be635c0C68Cdf3E946114221B12f4F7", RoleRelayer),
		secret("0x07082fcbFA5F5AC9FBc03A48B7f6391441DB8332", RoleMonitor),
	),
}
