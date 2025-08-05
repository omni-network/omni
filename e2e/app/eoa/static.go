package eoa

import (
	"github.com/omni-network/omni/lib/netconf"
)

//nolint:gochecknoglobals // Static mappings.
var statics = map[netconf.ID][]Account{
	netconf.Devnet: flatten(
		wellKnown(DevPrivateKey0(), RoleCreate3Deployer, RoleDeployer, RoleManager, RoleUpgrader),
		wellKnown(DevPrivateKey1(), RoleRedenomizer),
		wellKnown(DevPrivateKey2(), RoleFlowgen),
		wellKnown(DevPrivateKey3(), RoleXCaller),
		wellKnown(DevPrivateKey4(), RoleSolver),
		wellKnown(DevPrivateKey5(), RoleRelayer),
		wellKnown(DevPrivateKey6(), RoleMonitor),
		wellKnown(DevPrivateKey7(), RoleTester),
		wellKnown(DevPrivateKey8(), RoleHot),
		wellKnown(DevPrivateKey9(), RoleCold),
	),
	netconf.Staging: flatten(
		remote("0x64Bf40F5E6C4DE0dfe8fE6837F6339455657A2F5", RoleCold), // we use shared-cold
		remote("0xf63316AA39fEc9D2109AB0D9c7B1eE3a6F60AEA4", RoleHot),  // we use shared-hot
		remote("0xCC43713c9C9c565Fd4830cC85F7f254979F64518", RoleManager),
		remote("0x4891925c4f13A34FC26453FD168Db80aF3273014", RoleUpgrader),
		remote("0xC8103859Ac7CB547d70307EdeF1A2319FC305fdC", RoleCreate3Deployer),
		remote("0x274c4B3e5d27A65196d63964532366872F81D261", RoleDeployer),
		remote("0x7a6cF389082dc698285474976d7C75CAdE08ab7e", RoleTester), // reused shared-test with omega. Concurrent usage will result in nonce clashes.
		secret("0xfE921e06Ed0a22c035b4aCFF0A5D3a434A330c96", RoleRelayer),
		secret("0x0De553555Fa19d787Af4273B18bDB77282D618c4", RoleMonitor),
		secret("0x9a8BF80057c8E5B5526a87389cF21A631a420998", RoleFlowgen),
		secret("0xbC0F36A57B666922CF7C01003a01a613D44e993C", RoleXCaller),
		secret("0xa2C64d844c177C814b6F0423b41D8644539f5F58", RoleSolver),
		secret("0x2ced744c894eBdB09D15F4d426bCE07726602416", RoleRedenomizer),
	),
	netconf.Omega: flatten(
		remote("0x64Bf40F5E6C4DE0dfe8fE6837F6339455657A2F5", RoleCold),     // we use shared-cold
		remote("0xf63316AA39fEc9D2109AB0D9c7B1eE3a6F60AEA4", RoleHot),      // we use shared-hot
		remote("0xEAD625eB2011394cdD739E91Bf9D51A7169C22F5", RoleManager),  // we reuse omega-owner-upgrader
		remote("0xEAD625eB2011394cdD739E91Bf9D51A7169C22F5", RoleUpgrader), // we reuse omega-owner-upgrader
		remote("0xeC5134556da0797A5C5cD51DD622b689Cac97Fe9", RoleCreate3Deployer),
		remote("0x0CdCc644158b7D03f40197f55454dc7a11Bd92c1", RoleDeployer),
		remote("0x7a6cF389082dc698285474976d7C75CAdE08ab7e", RoleTester), // reused shared-test with staging. Concurrent usage will result in nonce clashes.
		secret("0x37AD6f7267454cac494C177127aC017750c8A7DB", RoleRelayer),
		secret("0xcef2a2c477Ec8473E4DeB9a8c2DF1D0585ea1040", RoleMonitor),
		secret("0x16DA241141D70290F043eED299140b6EC5942CAb", RoleFlowgen),
		secret("0x01A1A1C3Fe5369bc4DF3B5a5bbC10639a14113ab", RoleXCaller),
		secret("0x2677B9165c426eF0d3EC74E0e853e2F3A3d525cd", RoleSolver),
	),
	netconf.Mainnet: flatten(
		remote("0x8b6b217572582C57616262F9cE02A951A1D1b951", RoleCold),
		remote("0x8F609f4d58355539c48C98464E1e54ab2709aCfe", RoleHot),
		remote("0xd09DD1126385877352d24B669Fd68f462200756E", RoleManager),
		remote("0xF8740c09f25E2cbF5C9b34Ef142ED7B343B42360", RoleUpgrader),
		remote("0x992b9de7D42981B90A75C523842C01e27875b65B", RoleCreate3Deployer),
		remote("0x9496Bf1Bd2Fa5BCba72062cC781cC97eA6930A13", RoleDeployer),
		secret("0xfD62020Cee216Dc543E29752058Ee9f60f7D9Ff9", RoleMonitor),
		secret("0x7f409BaC75F5260340EbEC91066a845631Dc4859", RoleFlowgen),
		secret("0x6191442101086253A636aecBCC870e4778490AaB", RoleRelayer),
		secret("0x835c36774B28563b9a1d1ae83dD6F671F51DCb5c", RoleXCaller),
		secret("0x8cC81c5C09394CEaCa7a53be5f547AE719D75dFC", RoleSolver),
	),
}
