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

    /// @dev Matches Staking.Delegate event
    event Delegate(address indexed delegator, address indexed validator, uint256 amount);

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
        bytes memory pubkey = abi.encodePacked(hex"03440d290e4394cd9832cc7025769be18ab7975e34e4514c31c07da3d370fe0b05");
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

        // requires valid pubkey prefix
        bytes memory invalidPubkey = abi.encodePacked(hex"01", pubkey32);
        vm.expectRevert("Staking: invalid pubkey prefix");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(invalidPubkey);

        // requires a valid pubkey on the secp256k1 curve
        invalidPubkey = abi.encodePacked(hex"03", pubkey32);
        vm.expectRevert("Staking: invalid pubkey");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(invalidPubkey);

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

    function test_delegate() public {
        // requires min delegation
        address validator = makeAddr("validator");
        uint256 minDelegation = staking.MinDelegation();

        vm.deal(validator, minDelegation);

        vm.expectRevert("Staking: insufficient deposit");
        staking.delegate{ value: minDelegation - 1 }(validator);

        // requires self-delegation
        vm.expectRevert("Staking: only self delegation");
        vm.prank(validator);
        staking.delegate{ value: minDelegation }(makeAddr("someone else"));

        // if allowlist enabled, must be in allowlist
        vm.prank(owner);
        staking.enableAllowlist();

        vm.expectRevert("Staking: not allowed val");
        vm.prank(validator);
        staking.delegate{ value: minDelegation }(validator);

        // succeeds
        address[] memory validators = new address[](1);
        validators[0] = validator;
        vm.prank(owner);
        staking.allowValidators(validators);

        vm.expectEmit();
        emit Delegate(validator, validator, minDelegation);

        vm.prank(validator);
        staking.delegate{ value: minDelegation }(validator);
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
