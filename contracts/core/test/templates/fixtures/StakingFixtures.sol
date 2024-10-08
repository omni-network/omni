// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Staking } from "src/octane/Staking.sol";
import { Test } from "forge-std/Test.sol";

/// @dev Staking test fixtures
contract StakingFixtures is Test {
    address owner;
    Staking staking;

    function setUp() public {
        owner = makeAddr("owner");
        staking = new StakingHarness(owner);
    }
}

/// @dev Wrapper around Staking.sol that allows setting owner in constructor
contract StakingHarness is Staking {
    constructor(address _owner) {
        _transferOwnership(_owner);
    }
}
