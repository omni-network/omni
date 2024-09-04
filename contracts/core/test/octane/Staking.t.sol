// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { Staking } from "src/octane/Staking.sol";
import { Test, Vm } from "forge-std/Test.sol";

/**
 * @title Staking_Test
 * @notice Test suite for Staking.sol
 */
contract Staking_Test is Test {
    /// @dev Matches Staking.CreateValidator event
    event CreateValidator(address indexed validator, bytes pubkey, uint256 deposit);

    function test_createValidator() public {
        _testCreateValidator(_deployStaking());
    }

    function test_allowlist() public {
        _testAllowlist(_deployStaking());
    }

    function _deployStaking() internal returns (Staking) {
        address impl = address(new Staking());
        address proxy = address(
            new TransparentUpgradeableProxy(
                impl, address(this), abi.encodeCall(Staking.initialize, (makeAddr("owner"), false))
            )
        );

        return Staking(proxy);
    }

    /// @dev Test createValidator. Includes allowlist tests.
    function _testCreateValidator(Staking staking) internal {
        address owner = staking.owner();

        address validator = makeAddr("validator");
        address[] memory validators = new address[](1);
        validators[0] = validator;
        bytes memory pubkey = abi.encodePacked(hex"03", keccak256("pubkey"));
        vm.deal(validator, staking.MinDeposit());

        // enable allowlist
        vm.startPrank(owner);
        if (!staking.isAllowlistEnabled()) staking.setAllowlistEnabled(true);
        assertTrue(staking.isAllowlistEnabled());
        vm.stopPrank();

        // must be in allowlist
        vm.expectRevert("Staking: not allowed");
        staking.createValidator(pubkey);

        // add to allowlist
        vm.prank(owner);
        staking.setAllowlist(validators);
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
        staking.setAllowlist(new address[](0));
        assertFalse(staking.isAllowedValidator(validator));

        // must be in allowlist
        vm.expectRevert("Staking: not allowed");
        vm.deal(validator, deposit);
        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey);

        // disable allowlist
        vm.prank(owner);
        staking.setAllowlistEnabled(false);
        assertFalse(staking.isAllowlistEnabled());

        // can create validator with allowlist disabled
        vm.expectEmit();
        emit CreateValidator(validator, pubkey, deposit);

        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey);
    }

    function _testAllowlist(Staking staking) internal {
        // set allowlist
        address owner = staking.owner();

        // enable allowlist
        vm.prank(owner);
        staking.setAllowlistEnabled(true);
        assertTrue(staking.isAllowlistEnabled());

        address val1 = makeAddr("validator1");
        address val2 = makeAddr("validator2");
        address[] memory validators = new address[](2);
        validators[0] = val1;
        validators[1] = val2;

        // allow validators
        vm.prank(owner);
        staking.setAllowlist(validators);

        // check allowlist
        address[] memory allowlist = staking.allowlist();
        assertEq(allowlist.length, 2);
        assertEq(allowlist[0], validators[0]);
        assertEq(allowlist[1], validators[1]);
        assertTrue(staking.isAllowedValidator(validators[0]));
        assertTrue(staking.isAllowedValidator(validators[1]));

        // remove val1
        validators = new address[](1);
        validators[0] = val2;

        // reset
        vm.prank(owner);
        staking.setAllowlist(validators);

        // check allowlist
        allowlist = staking.allowlist();
        assertEq(allowlist.length, 1);
        assertEq(allowlist[0], validators[0]);
        assertFalse(staking.isAllowedValidator(val1));
        assertTrue(staking.isAllowedValidator(val2));

        // remove all
        validators = new address[](0);

        vm.prank(owner);
        staking.setAllowlist(validators);

        // check allowlist
        allowlist = staking.allowlist();
        assertEq(allowlist.length, 0);
        assertFalse(staking.isAllowedValidator(val1));
        assertFalse(staking.isAllowedValidator(val2));

        // disable allowlist
        vm.prank(owner);
        staking.setAllowlistEnabled(false);
        assertFalse(staking.isAllowlistEnabled());
    }

    /// @dev Run test suite
    function run(address staking) public {
        _testCreateValidator(Staking(staking));
        _testAllowlist(Staking(staking));
    }
}
