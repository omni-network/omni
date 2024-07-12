// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { IOmniPortalAdmin } from "src/interfaces/IOmniPortalAdmin.sol";

/**
 * @title UnpausePortal
 * @notice Unpause a portal contract.
 */
contract UnpausePortal is Script {
    function run(address portal) public {
        vm.startBroadcast();
        IOmniPortalAdmin(portal).unpause();
        vm.stopBroadcast();
    }
}
