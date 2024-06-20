// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Slashing } from "src/octane/Slashing.sol";
import { Test, Vm } from "forge-std/Test.sol";

/**
 * @title Slashing_Test
 * @notice Test suite for Slashing.sol
 */
contract Slashing_Test is Test {
    /// @dev Matches Slashing.Unjail event
    event Unjail(address indexed validator);

    Slashing slashing;

    function setUp() public {
        slashing = new Slashing();
    }

    function test_unjail() public {
        address validator = makeAddr("validator");
        uint256 fee = slashing.Fee();
        vm.deal(address(validator), fee);

        vm.expectEmit();
        emit Unjail(validator);

        vm.prank(validator);
        slashing.unjail{ value: fee }();
    }
}
