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

    /// @dev Test that an operator can register if in allow list
    function test_registerOperator_succeeds() public {
        address operator = _operator(0);

        _addToAllowlist(operator);
        _depositIntoSupportedStrategy(operator, 1 ether);
        _registerAsOperator(operator);
        _registerOperatorWithAVS(operator);

        IOmniAVS.Operator[] memory operators = omniAVS.operators();

        assertEq(operators.length, 1);
        assertEq(operators[0].addr, operator);
    }
}
