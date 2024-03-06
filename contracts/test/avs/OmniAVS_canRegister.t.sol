// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniAVS_canRegister_Test
 * @dev Test suite for the AVS canRegister function
 */
contract OmniAVS_canRegister_Test is Base {
    function test_canRegister_notOperator() public {
        address operator = _operator(0);

        (bool canRegister, string memory reason) = omniAVS.canRegister(operator);
        assertFalse(canRegister);
        assertEq(reason, "not an operator");
    }

    function test_canRegister_notAllowed() public {
        address operator = _operator(0);

        if (!omniAVS.allowlistEnabled()) _enableAllowlist();

        _registerAsOperator(operator);

        (bool canRegister, string memory reason) = omniAVS.canRegister(operator);
        assertFalse(canRegister);
        assertEq(reason, "not in allowlist");
    }

    function test_canRegister_maxOperatorsReached() public {
        address operator = _operator(0);

        _registerAsOperator(operator);
        _addToAllowlist(operator);

        vm.prank(omniAVSOwner);
        omniAVS.setMaxOperatorCount(0);

        (bool canRegister, string memory reason) = omniAVS.canRegister(operator);
        assertFalse(canRegister);
        assertEq(reason, "max operators reached");
    }

    function test_canRegister_minStakeNotMet() public {
        address operator = _operator(0);

        if (omniAVS.allowlistEnabled()) _addToAllowlist(operator);

        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake - 1);

        (bool canRegister, string memory reason) = omniAVS.canRegister(operator);
        assertFalse(canRegister);
        assertEq(reason, "min stake not met");
    }

    function test_canRegister_allowed() public {
        address operator = _operator(0);

        if (omniAVS.allowlistEnabled()) _addToAllowlist(operator);

        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);

        (bool canRegister, string memory reason) = omniAVS.canRegister(operator);
        assertTrue(canRegister);
        assertEq(reason, "");
    }

    function test_canRegister_allowlistDisabled() public {
        address operator = _operator(0);

        if (omniAVS.allowlistEnabled()) _disableAllowlist();

        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);

        (bool canRegister, string memory reason) = omniAVS.canRegister(operator);
        assertTrue(canRegister);
        assertEq(reason, "");
    }

    function test_canRegister_alreadyRegistered() public {
        address operator = _operator(0);

        if (omniAVS.allowlistEnabled()) _addToAllowlist(operator);

        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);
        _registerOperatorWithAVS(operator);

        (bool canRegister, string memory reason) = omniAVS.canRegister(operator);
        assertFalse(canRegister);
        assertEq(reason, "already registered");
    }
}
