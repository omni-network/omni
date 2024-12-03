// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { MockToken } from "test/utils/MockToken.sol";
import { Script } from "forge-std/Script.sol";

contract MintMockToken is Script {
    function run(address token, address to) public {
        vm.startBroadcast();
        MockToken(token).mint(to, 1000 ether);
        vm.stopBroadcast();
    }
}
