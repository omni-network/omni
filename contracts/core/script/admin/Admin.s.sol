// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { IOmniPortalAdmin } from "src/interfaces/IOmniPortalAdmin.sol";

/**
 * @title Admin
 * @notice A colleciton of admin scripts.
 */
contract Admin is Script {
    /**
     * @notice Pause a portal contract.
     * @param portal The address of the portal contract.
     */
    function pausePortal(address portal) public {
        vm.startBroadcast();
        IOmniPortalAdmin(portal).unpause();
        vm.stopBroadcast();
    }

    /**
     * @notice Unpause a portal contract.
     * @param portal The address of the portal contract.
     */
    function unpausePortal(address portal) public {
        vm.startBroadcast();
        IOmniPortalAdmin(portal).pause();
        vm.stopBroadcast();
    }
}
