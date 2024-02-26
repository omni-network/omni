// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

import { IOmniEthRestaking } from "src/interfaces/IOmniEthRestaking.sol";
import { OmniPredeploys } from "src/libraries/OmniPredeploys.sol";
import { OmniAVS } from "src/protocol/OmniAVS.sol";
import { AVSBase } from "./AVSBase.sol";
import { AVSUtils } from "./AVSUtils.sol";

contract OmniAVS_Test is AVSBase, AVSUtils {
    uint32 numOperators;
    uint32 numDelegatorsPerOp;
    address[] operators;
    address[] delegators;

    uint96 initialOperatorStake;
    uint96 initialDelegatorStake;
    uint96 delegatorStakeAddition = 1 ether;
    uint96 operatorStakeAddition = 1 ether;

    // feeForSync is msg.value passsed in tests OmniAVS.syncWithOmni()
    // normally, these fee will be calculated by the caller off-chain
    uint256 syncFee = 1 gwei;

    uint96 GWEI_TO_WEI = 1e9;

    /**
     * Test OmniAVS.syncWithOmni() at a number of different points in a "delegation lifecycle".
     *  - no registered operators
     *  - registered operators with some initial stake
     *  - delegators have delegated to operators
     *  - delegators have increased their stake to some operators
     *  - operators have increased their stake
     *  - delegators have undelegated
     *
     * For each point in delegation lifecycle, test that OmniAVS.getValidators() returns the
     * expected validators list, and that OmniAVS.syncWithOmni() makes a call to
     * OmniPortal.xcall with that list of validators.
     */
    function testFuzz_syncWithOmni(
        uint8 numOperators_,
        uint8 numDelegatorsPerOp_,
        uint96 initialOperatorStake_,
        uint96 initialDelegatorStake_
    ) public {
        numOperators = uint32(bound(numOperators_, 2, maxOperatorCount));
        numDelegatorsPerOp = uint32(bound(numDelegatorsPerOp_, 1, 30));

        initialOperatorStake = uint96(bound(initialOperatorStake_, minimumOperatorStake, 100 ether));
        initialDelegatorStake = uint96(bound(initialDelegatorStake_, 500 gwei, 5 ether));

        // round both initialOperatorStake and initialDelegatorStake to the nearest GWEI
        // beaconChainETHStrategy requires round numbers in GWEI
        initialOperatorStake = (initialOperatorStake / GWEI_TO_WEI) * GWEI_TO_WEI;
        initialDelegatorStake = (initialDelegatorStake / GWEI_TO_WEI) * GWEI_TO_WEI;

        for (uint256 i = 0; i < numOperators; i++) {
            operators.push(_operator(i));
        }

        for (uint256 i = 0; i < numOperators * numDelegatorsPerOp; i++) {
            delegators.push(_delegator(i));
        }

        // these functions much be called in this order
        _testRegisterOperators();
        _testRegisterOperatorsWithAVS();
        _testDelegateToOperators();
        _testIncreaseDelegationsToFirstHalfOfOperators();
        _testIncreaseStakeOfSecondHalfOfOperators();
        _testUndelegateAllDelegators();
        _testDeregisterOperators();
    }

    /// @dev Register operators with eigen layer core, assert OmniAVS quorum is still empty
    function _testRegisterOperators() internal {
        // NOTE: it is not necessary for operator to have deposited minimumStakeForQuorum
        // other staker(s) could have deposited and delegated to the operator
        for (uint32 i = 0; i < numOperators; i++) {
            operators[i] = _operator(i);
            _registerAsOperator(operators[i]);
            _depositIntoSupportedStrategy(operators[i], initialOperatorStake);
        }

        OmniAVS.Validator[] memory validators;
        validators = omniAVS.getValidators();

        // assert no operator for omni avs quorum, because no operator has been registered
        assertEq(validators.length, 0, "_testRegisterOperators: no operators should be registered");

        // TODO: should we revert if no operators are registered?
        // current thinking is no, to allow all operators to be deregistered
        _expectXCall(validators);
        omniAVS.syncWithOmni{ value: syncFee }();
    }

    /// @dev Register operators with OmniAVS, assert OmniAVS quorum is populated with initial stake
    function _testRegisterOperatorsWithAVS() internal {
        // register operators with AVS
        for (uint32 i = 0; i < numOperators; i++) {
            _addToAllowlist(operators[i]);
            _registerOperatorWithAVS(operators[i]);
        }

        OmniAVS.Validator[] memory validators;
        validators = omniAVS.getValidators();

        // assert all operators have been registered
        assertEq(validators.length, numOperators, "_testRegisterOperatorsWithAVS: all operators should be registered");

        // assert operator has initial stake
        for (uint32 i = 0; i < numOperators; i++) {
            assertEq(
                validators[i].staked,
                initialOperatorStake,
                "_testRegisterOperatorsWithAVS: validator.staked should be initialOperatorStake"
            );
            assertEq(validators[i].delegated, 0, "_testRegisterOperatorsWithAVS: validator.delegated should be 0");
        }

        _expectXCall(validators);
        omniAVS.syncWithOmni{ value: syncFee }();
    }

    /// @dev Delegate to operators, assert OmniAVS quorum is populated with initial stake + delegations
    function _testDelegateToOperators() internal {
        // initialize delegators
        for (uint32 i = 0; i < numOperators; i++) {
            address operator = operators[i];

            // for each operator, initizialzie numDelegatorsPerOp delegators
            for (uint32 j = 0; j < numDelegatorsPerOp; j++) {
                uint96 idx = i * numDelegatorsPerOp + j;
                address delegator = delegators[idx];

                // should contribute to quorom stake
                _depositIntoSupportedStrategy(delegator, initialDelegatorStake);

                // should NOT contribute to quorom stake
                _depositIntoUnsupportedStrategy(delegator, initialDelegatorStake);

                // all stake is delegated
                _testDelegateToOperator(delegator, operator);
            }
        }

        OmniAVS.Validator[] memory validators;
        validators = omniAVS.getValidators();

        // assert all operators still registered
        assertEq(validators.length, numOperators, "_testDelegateToOperators: all operators should still be registered");

        // assert all operator stake has been updated by initialDelegatorStake
        for (uint32 i = 0; i < numOperators; i++) {
            uint96 totalDelegated = numDelegatorsPerOp * initialDelegatorStake;
            uint96 totalStaked = initialOperatorStake;

            // validator state tracks these separately
            assertEq(validators[i].staked, totalStaked, "_testDelegateToOperators: validator.staked unexpected");
            assertEq(
                validators[i].delegated, totalDelegated, "_testDelegateToOperators: validator.delegated unexpected"
            );
        }

        _expectXCall(validators);
        omniAVS.syncWithOmni{ value: syncFee }();
    }

    /// @dev Increase delegations for first half of operators, assert OmniAVS quorum is updated
    function _testIncreaseDelegationsToFirstHalfOfOperators() internal {
        // increase delegations for first half of operators
        for (uint32 i = 0; i < numOperators / 2; i++) {
            for (uint32 j = 0; j < numDelegatorsPerOp; j++) {
                _depositIntoSupportedStrategy(delegators[i * numDelegatorsPerOp + j], delegatorStakeAddition);
            }
        }

        OmniAVS.Validator[] memory validators;
        validators = omniAVS.getValidators();

        // assert all operators still registered
        assertEq(
            validators.length,
            numOperators,
            "_testIncreaseDelegationsToFirstHalfOfOperators: all operators should still be registered"
        );

        // assert first half of operators have increased delegations
        for (uint32 i = 0; i < numOperators; i++) {
            // initial stake
            uint96 totalStaked = initialOperatorStake;

            // initial delegations
            uint96 totalDelegated = numDelegatorsPerOp * initialDelegatorStake;

            if (i < numOperators / 2) {
                // increase totalDelegated for first half of delegators
                totalDelegated += numDelegatorsPerOp * delegatorStakeAddition;
            }

            // validator state tracks these separately
            assertEq(
                validators[i].staked,
                totalStaked,
                "_testIncreaseDelegationsToFirstHalfOfOperators: validator.staked unexpected"
            );
            assertEq(
                validators[i].delegated,
                totalDelegated,
                "_testIncreaseDelegationsToFirstHalfOfOperators: validator.delegated unexpected"
            );
        }

        _expectXCall(validators);
        omniAVS.syncWithOmni{ value: syncFee }();
    }

    /// @dev Increase stake for second half of operators, assert OmniAVS quorum is updated
    function _testIncreaseStakeOfSecondHalfOfOperators() internal {
        // increase stake of second half of delegators
        for (uint32 i = numOperators / 2; i < numOperators; i++) {
            _depositIntoSupportedStrategy(operators[i], operatorStakeAddition);
        }

        OmniAVS.Validator[] memory validators;
        validators = omniAVS.getValidators();

        // assert all operators still registered
        assertEq(
            validators.length,
            numOperators,
            "_testIncreaseStakeOfSecondHalfOfOperators: all operators should still be registered"
        );

        // assert first half of operators have increased delegations by delegatorStakeAddition
        // assert second half of operators have increased stake by operatorStakeAddition
        for (uint32 i = 0; i < numOperators; i++) {
            // initial stake
            uint96 totalStaked = initialOperatorStake;

            // initial delegations
            uint96 totalDelegated = numDelegatorsPerOp * initialDelegatorStake;

            if (i < numOperators / 2) {
                // increase totalDelegated for first half of operators
                totalDelegated += numDelegatorsPerOp * delegatorStakeAddition;
            } else {
                // increase totalStaked for second half of operators
                totalStaked += operatorStakeAddition;
            }

            // validator state tracks these separately
            assertEq(
                validators[i].staked,
                totalStaked,
                "_testIncreaseStakeOfSecondHalfOfOperators: validator.staked unexpected "
            );
            assertEq(
                validators[i].delegated,
                totalDelegated,
                "_testIncreaseStakeOfSecondHalfOfOperators: validator.delegated unexpected"
            );
        }

        _expectXCall(validators);
        omniAVS.syncWithOmni{ value: syncFee }();
    }

    /// @dev Undelegate all delegators, assert OmniAVS quorum is updated
    function _testUndelegateAllDelegators() internal {
        // undelegate all delegators
        for (uint32 i = 0; i < numOperators; i++) {
            for (uint32 j = 0; j < numDelegatorsPerOp; j++) {
                address delegator = delegators[i * numDelegatorsPerOp + j];
                vm.prank(delegator);
                // removes delegation to operator (stakers can only delegate to one operator at a time)
                delegation.undelegate(delegator);
            }
        }

        OmniAVS.Validator[] memory validators;
        validators = omniAVS.getValidators();

        // assert all operators still registered
        assertEq(
            validators.length, numOperators, "_testUndelegateAllDelegators: all operators should still be registered"
        );

        // assert all operators have no delegations
        // assert first half is back to initial stake
        // assert second half is back to initial stake + operatorStakeAddition
        for (uint32 i = 0; i < numOperators; i++) {
            // initial stake
            uint96 totalStaked = initialOperatorStake;

            if (i >= numOperators / 2) {
                // increase totalStaked for second half of operators
                totalStaked += operatorStakeAddition;
            }

            // validator state tracks these separately
            assertEq(validators[i].staked, totalStaked, "_testUndelegateAllDelegators: validator.staked unexpected ");
            assertEq(validators[i].delegated, 0, "_testtUndelegateAllDelegators: validator.delegated should be 0");
        }

        _expectXCall(validators);
        omniAVS.syncWithOmni{ value: syncFee }();
    }

    /// @dev Deregister operators, assert OmniAVS quorum is updated after each deregistration
    function _testDeregisterOperators() internal {
        OmniAVS.Validator[] memory validators;

        for (uint32 i = 0; i < numOperators; i++) {
            address operator = operators[i];

            _deregisterOperatorFromAVS(operator);

            validators = omniAVS.getValidators();

            uint96 numOperatorsLeft = numOperators - i - 1;

            // assert there are only numOperatorsLeft
            assertEq(validators.length, numOperatorsLeft);

            // assert that none of the operators left is the operator that just deregistered
            for (uint32 j = 0; j < numOperatorsLeft; j++) {
                assertNotEq(
                    validators[j].addr, operator, "_testDeregisterOperators: operator should not be in validators list"
                );
            }
        }
    }

    /// @dev Expect an OmniPortal.xcall to IOmniEthRestaking.sync(validators), with correct fee and gasLimit
    function _expectXCall(OmniAVS.Validator[] memory validators) internal {
        bytes memory data = abi.encodeWithSelector(IOmniEthRestaking.sync.selector, validators);
        uint64 gasLimit = omniAVS.xcallBaseGasLimit() + omniAVS.xcallGasLimitPerValidator() * uint64(validators.length);

        vm.expectCall(
            address(portal),
            syncFee,
            abi.encodeWithSignature(
                "xcall(uint64,address,bytes,uint64)", omniChainId, OmniPredeploys.OMNI_ETH_RESTAKING, data, gasLimit
            )
        );
    }

    /**
     * Unit tests.
     */
    function test_depositBeaconEth_succeeds() public {
        address operator = _operator(0);
        uint96 amount = minimumOperatorStake;

        _registerAsOperator(operator);
        _addToAllowlist(operator);
        _depositBeaconEth(operator, amount);
        _registerOperatorWithAVS(operator);

        OmniAVS.Validator[] memory validators = omniAVS.getValidators();

        assertEq(validators.length, 1);
        assertEq(validators[0].addr, operator);
        assertEq(validators[0].staked, amount);
        assertEq(validators[0].delegated, 0);
    }

    /// @dev Test that an operator cannot register if not in allow list
    function test_registerOperator_notAllowed_reverts() public {
        address operator = _operator(0);

        ISignatureUtils.SignatureWithSaltAndExpiry memory emptySig;

        vm.expectRevert("OmniAVS: not allowed");
        vm.prank(operator);
        omniAVS.registerOperatorToAVS(operator, emptySig);
    }

    /// @dev Test that the owner can deregister an operator
    function test_deregisterOperator_byOwner_succeeds() public {
        address operator = _operator(0);

        // register operator
        _registerAsOperator(operator);
        _addToAllowlist(operator);
        _depositBeaconEth(operator, minimumOperatorStake);
        _registerOperatorWithAVS(operator);

        // assert operator is registered
        OmniAVS.Validator[] memory validators = omniAVS.getValidators();
        assertEq(validators.length, 1);
        assertEq(validators[0].addr, operator);

        // deregister operator
        vm.prank(omniAVSOwner);
        omniAVS.deregisterOperatorFromAVS(operator);

        // assert operator is deregistered
        validators = omniAVS.getValidators();
        assertEq(validators.length, 0);
    }

    /// @dev Test that an operator can be added to the allowlist
    function test_addToAllowlist_succeeds() public {
        address operator = makeAddr("operator");
        _addToAllowlist(operator);
        assertTrue(omniAVS.isInAllowlist(operator));
    }

    /// @dev Test that an operator can be removed from the allowlist
    function test_removeFromAllowlist_succeeds() public {
        address operator1 = makeAddr("operator");
        address operator2 = makeAddr("operator2");

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
        address operator = makeAddr("operator");

        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.addToAllowlist(operator);
    }

    /// @dev Test that only the owner can remove from the allowlist
    function test_removeFromAllowlist_notOwner_reverts() public {
        address operator = makeAddr("operator");

        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.removeFromAllowlist(operator);
    }
}
