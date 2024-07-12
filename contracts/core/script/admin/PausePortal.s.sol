// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { IOmniPortalAdmin } from "src/interfaces/IOmniPortalAdmin.sol";
import { Admins } from "./Admins.sol";

/**
 * @title PausePortal
 * @notice Pause a portal contract.
 */
contract PausePortal is Script {
    function run(string calldata network, address portal) public {
        vm.startBroadcast(Admins.forNetwork(network));
        IOmniPortalAdmin(portal).pause();
        vm.stopBroadcast();
    }
}
