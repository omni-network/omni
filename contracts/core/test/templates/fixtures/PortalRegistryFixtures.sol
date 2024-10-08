// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { PortalRegistry } from "src/xchain/PortalRegistry.sol";
import { Test } from "forge-std/Test.sol";

/// @dev PortalRegistry test fixtures
contract PortalRegistryFixtures is Test {
    PortalRegistry reg;
    address owner;

    function setUp() public {
        owner = makeAddr("owner");
        reg = new PortalRegistryHarness(owner);
    }
}

/// @dev Wrapper around PortalRegistry that adds a constructor.
contract PortalRegistryHarness is PortalRegistry {
    constructor(address _owner) {
        _transferOwnership(_owner);
    }
}
