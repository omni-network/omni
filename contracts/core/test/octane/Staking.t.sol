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

    address owner;
    StakingHarness staking;

    function setUp() public {
        owner = makeAddr("owner");
        staking = new StakingHarness(owner);
    }

    function test_createValidator() public {
        address validator = makeAddr("validator");
        address[] memory validators = new address[](1);
        validators[0] = validator;
        bytes memory pubkey = abi.encodePacked(hex"03", keccak256("pubkey"));
        vm.deal(validator, staking.MinDeposit());

        // allowlist is disabled
        assertFalse(staking.isAllowlistEnabled());

        // enable allowlist
        vm.prank(owner);
        staking.enableAllowlist();
        assertTrue(staking.isAllowlistEnabled());

        // must be in allowlist
        vm.expectRevert("Staking: not allowed");
        staking.createValidator(pubkey);

        // add to allowlist
        vm.prank(owner);
        staking.allowValidators(validators);
        assertTrue(staking.isAllowedValidator(validator));

        // requires minimum deposit
        uint256 insufficientDeposit = staking.MinDeposit() - 1;

        vm.expectRevert("Staking: insufficient deposit");
        vm.prank(validator);
        staking.createValidator{ value: insufficientDeposit }(pubkey);

        // requires 33 byte
        uint256 deposit = staking.MinDeposit();
        bytes memory pubkey32 = abi.encodePacked(keccak256("pubkey"));

        vm.expectRevert("Staking: invalid pubkey length");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey32);

        // succeeds with valid deposit and pubkey
        vm.expectEmit();
        emit CreateValidator(validator, pubkey, deposit);

        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey);

        // remove from allowlist
        vm.prank(owner);
        staking.disallowValidators(validators);
        assertFalse(staking.isAllowedValidator(validator));

        // must be in allowlist
        vm.expectRevert("Staking: not allowed");
        vm.deal(validator, deposit);
        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey);

        // disable allowlist
        vm.prank(owner);
        staking.disableAllowlist();
        assertFalse(staking.isAllowlistEnabled());

        // can create validator with allowlist disabled
        vm.expectEmit();
        emit CreateValidator(validator, pubkey, deposit);

        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey);
    }
}

/**
 * @title StakingHarness
 * @notice Wrapper around Staking.sol that allows setting owner in constructor
 */
contract StakingHarness is Staking {
    constructor(address _owner) {
        _transferOwnership(_owner);
    }
}
