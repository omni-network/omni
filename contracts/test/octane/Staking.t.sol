// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Staking } from "src/octane/Staking.sol";
import { Test, Vm } from "forge-std/Test.sol";

/**
 * @title Staking_Test
 * @notice Test suite for Staking.sol
 */
contract Staking_Test is Test {
    /// @dev Matches Staking.CreateValidator event
    event CreateValidator(address indexed validator, bytes pubkey, uint256 deposit);

    Staking staking;

    function setUp() public {
        staking = new Staking();
    }

    function test_createValidator() public {
        address validator = makeAddr("validator");
        bytes memory pubkey = abi.encodePacked(hex"03", keccak256("pubkey"));
        vm.deal(validator, staking.MIN_DEPOSIT());

        // requires minimum deposit
        uint256 insufficientDeposit = staking.MIN_DEPOSIT() - 1;

        vm.expectRevert("Staking: insufficient deposit");
        vm.prank(validator);
        staking.createValidator{ value: insufficientDeposit }(pubkey);

        // requires 33 byte
        uint256 deposit = staking.MIN_DEPOSIT();
        bytes memory pubkey32 = abi.encodePacked(keccak256("pubkey"));

        vm.expectRevert("Staking: invalid pubkey length");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey32);

        // succeeds with valid deposit and pubkey
        vm.expectEmit();
        emit CreateValidator(validator, pubkey, deposit);

        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey);
    }
}
