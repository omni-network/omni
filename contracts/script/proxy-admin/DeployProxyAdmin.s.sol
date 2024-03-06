// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

// solhint-disable no-console

import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import { Script } from "forge-std/Script.sol";
import { console } from "forge-std/console.sol";

contract DeployProxyAdmin is Script {
    address internal _goerliOwner = 0x46D005a3D18740Cd63e8072078796f043d040191;
    address internal _anvilOwner = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266; // account 0
    address internal _goerliDeployer = 0x6706acd456f199E2ce3Ef8e4688D83A24Ef35489;
    address internal _anvilDeployer = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266; // account 0

    function run() external {
        uint256 deployerKey = vm.envUint("PROXY_ADMIN_DEPLOYER_KEY");

        require(block.chainid == 5 || block.chainid == 31_337, "Don't deploy on this chain yes!");

        if (block.chainid == 5) require(vm.addr(deployerKey) == _goerliDeployer, "wrong goerli deployer key");
        if (block.chainid == 31_337) require(vm.addr(deployerKey) == _anvilDeployer, "wrong anvil deployer key");

        address owner;
        if (block.chainid == 5) owner = _goerliOwner;
        if (block.chainid == 31_337) owner = _anvilOwner;
        require(owner != address(0), "Owner not set");

        vm.startBroadcast(deployerKey);
        (ProxyAdmin proxyAdmin) = deploy(owner);
        vm.stopBroadcast();

        console.log("ProxyAdmin deployed at: ", address(proxyAdmin));
        console.log("ProxyAdmin owner: ", owner);
        console.log("Chain ID: ", block.chainid);
    }

    function deploy(address owner) public returns (ProxyAdmin proxyAdmin) {
        require(owner != address(0), "Owner not set");

        proxyAdmin = new ProxyAdmin();
        proxyAdmin.transferOwnership(owner);

        require(proxyAdmin.owner() == owner, "ProxyAdmin owner not set");
    }
}
