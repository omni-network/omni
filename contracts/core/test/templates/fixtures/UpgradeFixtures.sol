// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Upgrade } from "src/octane/Upgrade.sol";
import { Test } from "forge-std/Test.sol";

/// @dev Upgrade test fixtures
contract UpgradeFixtures is Test {
    address owner;
    Upgrade upgrade;

    function setUp() public {
        owner = makeAddr("owner");
        upgrade = new UpgradeHarness(owner);
    }
}

/// @dev Wrapper around Upgrade.sol that allows setting owner in constructor
contract UpgradeHarness is Upgrade {
    constructor(address _owner) {
        _transferOwnership(_owner);
    }
}
