// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Staking } from "src/octane/Staking.sol";
import { Test, Vm } from "forge-std/Test.sol";
import { Secp256k1 } from "src/libraries/Secp256k1.sol";

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
        bytes32 x = 0x991aec2da95415fd61dff26eb1439618d659e5a5f07f114554217b4265031012;
        bytes32 y = 0x418b0e8b020e76744e231152dada6ab80b4cdf39e55d273510302e8d135bdcbc;
        bytes memory pubkey = Secp256k1.compressPublicKey(x, y);
        bytes memory signature =
            hex"8ac5fe19a8bf5a3c8acd9b9ba0a7f06cb3b162aab01772988e2d179b454badb32f1db3170a117c7e681ef592ea90950d7a85e7278afb8f46dce02cdf26710a2700";
        vm.deal(validator, staking.MinDeposit());

        // allowlist is disabled
        assertFalse(staking.isAllowlistEnabled());

        // enable allowlist
        vm.prank(owner);
        staking.enableAllowlist();
        assertTrue(staking.isAllowlistEnabled());

        // must be in allowlist
        vm.expectRevert("Staking: not allowed");
        staking.createValidator(x, y, signature);

        // add to allowlist
        vm.prank(owner);
        staking.allowValidators(validators);
        assertTrue(staking.isAllowedValidator(validator));

        // requires minimum deposit
        uint256 insufficientDeposit = staking.MinDeposit() - 1;

        vm.expectRevert("Staking: insufficient deposit");
        vm.prank(validator);
        staking.createValidator{ value: insufficientDeposit }(x, y, signature);

        uint256 deposit = staking.MinDeposit();

        // requires a valid pubkey on the secp256k1 curve
        bytes32 badY = bytes32(uint256(y) + 1);
        vm.expectRevert("Staking: invalid pubkey");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(x, badY, signature);

        // requires a valid signature
        bytes memory badSignature = abi.encodePacked(signature);
        badSignature[0] = bytes1(uint8(badSignature[0]) + 1);
        vm.expectRevert("Staking: invalid signature");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(x, y, badSignature);

        // succeeds with valid deposit and pubkey
        vm.expectEmit();
        emit CreateValidator(validator, pubkey, deposit);

        vm.prank(validator);
        staking.createValidator{ value: deposit }(x, y, signature);

        // remove from allowlist
        vm.prank(owner);
        staking.disallowValidators(validators);
        assertFalse(staking.isAllowedValidator(validator));

        // must be in allowlist
        vm.expectRevert("Staking: not allowed");
        vm.deal(validator, deposit);
        vm.prank(validator);
        staking.createValidator{ value: deposit }(x, y, signature);

        // disable allowlist
        vm.prank(owner);
        staking.disableAllowlist();
        assertFalse(staking.isAllowlistEnabled());

        // can create validator with allowlist disabled
        vm.expectEmit();
        emit CreateValidator(validator, pubkey, deposit);

        vm.prank(validator);
        staking.createValidator{ value: deposit }(x, y, signature);
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
