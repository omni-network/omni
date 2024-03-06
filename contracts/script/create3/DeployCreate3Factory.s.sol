// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { Create3 } from "src/deploy/Create3.sol";
import { Script } from "forge-std/Script.sol";

contract DeployCreate3Factory is Script {
    address internal _goerliDeployer = 0x0c6f5DD88D87F703d18d3D52397d20a885324A91;
    address internal _anvilDeployer = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266; // account 0

    function run() external {
        uint256 deployerKey = vm.envUint("CREATE3_DEPLOYER_KEY");

        require(block.chainid == 5 || block.chainid == 31_337, "Don't deploy on this chain yes!");
        if (block.chainid == 5) require(vm.addr(deployerKey) == _goerliDeployer, "wrong goerli deployer key");
        if (block.chainid == 31_337) require(vm.addr(deployerKey) == _anvilDeployer, "wrong anvil deployer key");

        vm.startBroadcast(deployerKey);
        new Create3();
        vm.stopBroadcast();
    }
}
