// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.13;

import { Create3 } from "src/deploy/Create3.sol";
import { Script } from "forge-std/Script.sol";

contract DeployCreate3Factory is Script {
    function run() external {
        uint256 deployerKey = vm.envUint("DEPLOYER_KEY");
        vm.startBroadcast(deployerKey);
        new Create3();
    }
}
