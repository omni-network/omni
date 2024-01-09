// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import {Test} from "forge-std/Test.sol";
import {OmniPortal} from "src/OmniPortal.sol";
import {Events} from "test/common/Events.sol";

/**
 * @title CommonTest
 * @dev An extension of forge Test that includes Omni specifc setup, utils, and events.
 */
contract CommonTest is Test, Events {
    address deployer;
    address xcaller;

    OmniPortal portal;

    function setUp() public {
        deployer = makeAddr("deployer");
        xcaller = makeAddr("xcaller");

        vm.prank(deployer);
        portal = new OmniPortal();
    }
}
