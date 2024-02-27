// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniAVS_admin_Test
 * @dev Test suite for the AVS amdin functionality
 */
contract OmniAVS_admin_Test is Base {
    /// @dev Test that the owner can deregister an operator
    function test_deregisterOperator_byOwner_succeeds() public {
        address operator = _operator(0);

        // register operator
        _registerAsOperator(operator);
        _addToAllowlist(operator);
        _depositBeaconEth(operator, minimumOperatorStake);
        _registerOperatorWithAVS(operator);

        // assert operator is registered
        IOmniAVS.Validator[] memory validators = omniAVS.getValidators();
        assertEq(validators.length, 1);
        assertEq(validators[0].addr, operator);

        // deregister operator
        vm.prank(omniAVSOwner);
        omniAVS.deregisterOperatorFromAVS(operator);

        // assert operator is deregistered
        validators = omniAVS.getValidators();
        assertEq(validators.length, 0);
    }
}
