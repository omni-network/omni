// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Slashing } from "src/octane/Slashing.sol";
import { Test } from "forge-std/Test.sol";

/// @dev Slashing test fixtures
contract SlashingFixtures is Test {
    Slashing slashing;

    function setUp() public {
        slashing = new Slashing();
    }
}
