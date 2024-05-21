// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniAVS_allowlist_Test
 * @dev Test suite for the AVS allowlist functionality
 */
contract OmniAVS_allowlist_Test is Base {
    /// @dev Test that an operator can be added to the allowlist
    function test_addToAllowlist_succeeds() public {
        address operator = _operator(0);
        _addToAllowlist(operator);
        assertTrue(omniAVS.isInAllowlist(operator));
    }

    /// @dev Test that an operator can be removed from the allowlist
    function test_removeFromAllowlist_succeeds() public {
        address operator1 = _operator(0);
        address operator2 = _operator(1);

        _addToAllowlist(operator1);
        _addToAllowlist(operator2);
        assertTrue(omniAVS.isInAllowlist(operator1));
        assertTrue(omniAVS.isInAllowlist(operator2));

        _removeFromAllowlist(operator1);
        assertFalse(omniAVS.isInAllowlist(operator1));
        assertTrue(omniAVS.isInAllowlist(operator2));
    }

    /// @dev Test that only the owner can add to the allowlist
    function test_addToAllowlist_notOwner_reverts() public {
        address operator = _operator(0);

        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.addToAllowlist(operator);
    }

    /// @dev Test that only the owner can remove from the allowlist
    function test_removeFromAllowlist_notOwner_reverts() public {
        address operator = _operator(0);

        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.removeFromAllowlist(operator);
    }

    /// @dev Test that an operator can register if in allowlist
    function test_registerOperator_succeeds() public {
        address operator = _operator(0);

        _addToAllowlist(operator);
        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);
        _registerOperatorWithAVS(operator);

        IOmniAVS.Operator[] memory operators = omniAVS.operators();

        assertEq(operators.length, 1);
        assertEq(operators[0].addr, operator);
    }

    /// @dev Test that an operator can't register if not in allowlist
    function test_registerOperator_nowAllowed_reverts() public {
        if (!omniAVS.allowlistEnabled()) _enableAllowlist();

        address operator = _operator(0);
        ISignatureUtils.SignatureWithSaltAndExpiry memory emptySig;

        vm.expectRevert("OmniAVS: not allowed");
        vm.prank(operator);
        omniAVS.registerOperator(_pubkey(operator), emptySig);
    }

    /// @dev Test that the owner can disable the allowlist
    function test_disableAllowlist_succeeds() public {
        if (!omniAVS.allowlistEnabled()) _enableAllowlist();

        vm.prank(omniAVSOwner);
        omniAVS.disableAllowlist();
        assertFalse(omniAVS.allowlistEnabled());
    }

    /// @dev Test that only the owner can disable the allowlist
    function test_disableAllowlist_notOwner_reverts() public {
        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.disableAllowlist();
    }

    /// @dev Test that the allowlist can't be disabled if already disabled
    function test_disableAllowlist_alreadyDisabled_reverts() public {
        if (omniAVS.allowlistEnabled()) _disableAllowlist();

        assertFalse(omniAVS.allowlistEnabled());
        vm.expectRevert("OmniAVS: already disabled");
        vm.prank(omniAVSOwner);
        omniAVS.disableAllowlist();
    }

    /// @dev Test that the owner can enable the allowlist
    function test_enableAllowlist_succeeds() public {
        if (omniAVS.allowlistEnabled()) _disableAllowlist();

        vm.prank(omniAVSOwner);
        omniAVS.enableAllowlist();
        assertTrue(omniAVS.allowlistEnabled());
    }

    /// @dev Test that only the owner can enable the allowlist
    function test_enableAllowlist_notOwner_reverts() public {
        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.enableAllowlist();
    }

    /// @dev Test that the allowlist can't be enabled if already enabled
    function test_enableAllowlist_alreadyEnabled_reverts() public {
        if (!omniAVS.allowlistEnabled()) _enableAllowlist();

        assertTrue(omniAVS.allowlistEnabled());
        vm.expectRevert("OmniAVS: already enabled");
        vm.prank(omniAVSOwner);
        omniAVS.enableAllowlist();
    }

    /// @dev Test that an operator can register if not in allowlist, if allowlist is disabled
    function test_registerOperator_allowlistDisabled_succeeds() public {
        if (omniAVS.allowlistEnabled()) _disableAllowlist();

        address operator = _operator(0);

        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);
        _registerOperatorWithAVS(operator);

        IOmniAVS.Operator[] memory operators = omniAVS.operators();

        assertEq(operators.length, 1);
        assertEq(operators[0].addr, operator);
    }
}
