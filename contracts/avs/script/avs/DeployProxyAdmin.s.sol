// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

// solhint-disable no-console
// solhint-disable var-name-mixedcase
// solhint-disable max-states-count
// solhint-disable state-visibility

import { ProxyAdmin } from "src/ProxyAdmin.sol";

import { Script } from "forge-std/Script.sol";
import { console } from "forge-std/console.sol";

contract DeployProxyAdmin is Script {
    address _deployer = 0x00000072e2740F8a9A4D20Ed05C1832d12498642;
    address _owner = 0xFf89C654846B2E4BC572cEABE77056daf7b299a3;

    function run() external {
        uint256 deployerKey = vm.envUint("PROXY_ADMIN_DEPLOYER_KEY");

        require(block.chainid == 1, "Only mainnet deployment.");
        require(vm.addr(deployerKey) == _deployer, "Wrong deployer key");

        vm.startBroadcast(deployerKey);
        (ProxyAdmin proxyAdmin) = deploy(_owner);
        vm.stopBroadcast();

        console.log("ProxyAdmin deployed at: ", address(proxyAdmin));
        console.log("ProxyAdmin owner: ", _owner);
    }

    function deploy(address owner) public returns (ProxyAdmin proxyAdmin) {
        require(owner != address(0), "Owner not set");
        proxyAdmin = new ProxyAdmin(owner);
        require(address(proxyAdmin) != address(0), "ProxyAdmin not deployed");
        require(proxyAdmin.owner() == owner, "ProxyAdmin owner not set");
    }
}
