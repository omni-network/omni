// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniAVS_registration_Test
 * @dev Test suite for the AVS registration functionality
 */
contract OmniAVS_allowlist_Test is Base {
    /// @dev Test that an operator can register
    function test_registerOperator_succeeds() public {
        address operator = _operator(0);

        _addToAllowlist(operator);
        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);
        _registerOperatorWithAVS(operator);

        IOmniAVS.Operator[] memory operators = omniAVS.operators();

        assertEq(operators.length, 1);
        assertEq(operators[0].addr, operator);
        assertEq(operators[0].pubkey, _pubkey(operator));
    }

    /// @dev That that an operator cannot register with the wrong pubkey
    function test_registerOperator_wrongPubkey_reverts() public {
        address operator1 = _operator(0);
        address operator2 = _operator(1);

        _addToAllowlist(operator1);
        _registerAsOperator(operator1);
        _depositIntoSupportedStrategy(operator1, minOperatorStake);

        ISignatureUtils.SignatureWithSaltAndExpiry memory emptySig;

        vm.expectRevert("OmniAVS: pubkey != sender");
        vm.prank(operator1);
        omniAVS.registerOperator(_pubkey(operator2), emptySig);
    }

    /// @dev Test that an operator cannot register with an invalid pubkey
    function test_registerOperator_invalidPubkey_reverts() public {
        address operator = _operator(0);

        _addToAllowlist(operator);
        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);

        ISignatureUtils.SignatureWithSaltAndExpiry memory emptySig;
        bytes memory pubkey = bytes.concat(hex"04", _pubkey(operator));

        vm.expectRevert("Secp256k1: invalid pubkey length");
        vm.prank(operator);
        omniAVS.registerOperator(pubkey, emptySig);
    }
}
