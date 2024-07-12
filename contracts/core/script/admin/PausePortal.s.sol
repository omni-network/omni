// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { IOmniPortalAdmin } from "src/interfaces/IOmniPortalAdmin.sol";

/**
 * @title PausePortal
 * @notice Pause a portal contract.
 */
contract PausePortal is Script {
    function run(address portal) public {
        vm.startBroadcast();
        IOmniPortalAdmin(portal).pause();
        vm.stopBroadcast();
    }
}
