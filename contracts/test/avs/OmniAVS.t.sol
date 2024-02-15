// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { OperatorStateRetriever } from "eigenlayer-middleware/src/OperatorStateRetriever.sol";

import { OmniAVS } from "src/protocol/OmniAVS.sol";
import { AVSBase } from "./AVSBase.sol";
import { AVSUtils } from "./AVSUtils.sol";

import { console } from "forge-std/console.sol";

contract OmniAVS_Test is AVSBase, AVSUtils {
    uint32 numOperators;
    uint32 numDelegatorsPerOp;
    address[] operators;
    address[] delegators;

    uint96 initialOperatorStake;
    uint96 initialDelegatorStake;
    uint96 delegatorStakeAddition = 1 ether;
    uint96 operatorStakeAddition = 1 ether;

    /**
     * Test OmniAVS.getValidators() at a number of different points in a "delegation lifecycle"
     *  - no registered operators
     *  - registered operators with some initial stake
     *  - delegators have delegated to operators
     *  - delegators have increased their stake to some operators
     *  - operators have increased their stake
     *  - delegators have undelegated
     *
     * NOTES:
     *
     *  - Test Omni AVS uses the WETH strategy, configured in OmniDelegationAVSBase WETH counts 1:1 for stake.
     *  - We test both OmniAVS.getValidators() and OmniAVS._getOperatorState() (exposed in harness) to ensure
     *    that getValidators (the function we wrote) is consistent with getOperatorState, which proxies
     *    OperatorStateRetriever.getOperatorState (a function Eigen wrote).
     */
    function testFuzz_getValidators(
        uint8 numOperators_,
        uint8 numDelegatorsPerOp_,
        uint96 initialOperatorStake_,
        uint96 initialDelegatorStake_
    ) public {
        numOperators = uint32(bound(numOperators_, 2, defaultMaxOperatorCount));
        numDelegatorsPerOp = uint32(bound(numDelegatorsPerOp_, 1, 30));
        initialOperatorStake = uint96(bound(initialOperatorStake_, minimumStakeForQuorum, 100 ether));
        initialDelegatorStake = uint96(bound(initialDelegatorStake_, 500 gwei, 5 ether));

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
            _depositWeth(operators[i], initialOperatorStake);
        }

        OperatorStateRetriever.Operator[][] memory operatorState;
        OmniAVS.Validator[] memory validators;

        omniAVS.syncWithEigenLayer();
        operatorState = omniAVS.getOperatorState();
        validators = omniAVS.getValidators();

        // assert no operator for omni avs quorum, because no operator has been registered
        assertEq(operatorState.length, 1); // only one quorum
        assertEq(operatorState[0].length, 0); // no operators
        assertEq(validators.length, 0); // no validators
    }

    /// @dev Register operators with OmniAVS, assert OmniAVS quorum is populated with initial stake
    function _testRegisterOperatorsWithAVS() internal {
        // register operators with AVS
        for (uint32 i = 0; i < numOperators; i++) {
            _registerOperatorWithAVS(operators[i]);
        }

        OperatorStateRetriever.Operator[][] memory operatorState;
        OmniAVS.Validator[] memory validators;

        omniAVS.syncWithEigenLayer();
        operatorState = omniAVS.getOperatorState();
        validators = omniAVS.getValidators();

        // assert all operators have been registered
        assertEq(operatorState.length, 1);
        assertEq(operatorState[0].length, numOperators);
        assertEq(validators.length, numOperators);

        // assert operator has initial stake
        for (uint32 i = 0; i < numOperators; i++) {
            assertEq(operatorState[0][i].stake, initialOperatorStake);
            assertEq(validators[i].staked, initialOperatorStake);
            assertEq(validators[i].delegated, 0);
        }
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
                _depositWeth(delegator, initialDelegatorStake);

                // should NOT contribute to quorom stake
                _depositEigen(delegator, initialDelegatorStake);

                // all stake is delegated
                _testDelegateToOperator(delegator, operator);
            }
        }

        OperatorStateRetriever.Operator[][] memory operatorState;
        OmniAVS.Validator[] memory validators;

        omniAVS.syncWithEigenLayer();
        operatorState = omniAVS.getOperatorState();
        validators = omniAVS.getValidators();

        // assert all operators still registered
        assertEq(operatorState.length, 1); // only one quorum
        assertEq(operatorState[0].length, numOperators);
        assertEq(validators.length, numOperators);

        // assert all operator stake has been updated by initialDelegatorStake
        for (uint32 i = 0; i < numOperators; i++) {
            uint96 totalDelegated = numDelegatorsPerOp * initialDelegatorStake;
            uint96 totalStaked = initialOperatorStake;

            // operator state is delegations + stake
            assertEq(operatorState[0][i].stake, totalStaked + totalDelegated);

            // validator state tracks these separately
            assertEq(validators[i].staked, totalStaked);
            assertEq(validators[i].delegated, totalDelegated);
        }
    }

    /// @dev Increase delegations for first half of operators, assert OmniAVS quorum is updated
    function _testIncreaseDelegationsToFirstHalfOfOperators() internal {
        // increase delegations for first half of operators
        for (uint32 i = 0; i < numOperators / 2; i++) {
            for (uint32 j = 0; j < numDelegatorsPerOp; j++) {
                _depositWeth(delegators[i * numDelegatorsPerOp + j], delegatorStakeAddition);
            }
        }

        OperatorStateRetriever.Operator[][] memory operatorState;
        OmniAVS.Validator[] memory validators;

        omniAVS.syncWithEigenLayer();
        operatorState = omniAVS.getOperatorState();
        validators = omniAVS.getValidators();

        // assert all operators still registered
        assertEq(operatorState.length, 1); // only one quorum
        assertEq(operatorState[0].length, numOperators);
        assertEq(validators.length, numOperators);

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

            // operator state is delegations + stake
            assertEq(operatorState[0][i].stake, totalStaked + totalDelegated);

            // validator state tracks these separately
            assertEq(validators[i].staked, totalStaked);
            assertEq(validators[i].delegated, totalDelegated);
        }
    }

    /// @dev Increase stake for second half of operators, assert OmniAVS quorum is updated
    function _testIncreaseStakeOfSecondHalfOfOperators() internal {
        // increase stake of second half of delegators
        for (uint32 i = numOperators / 2; i < numOperators; i++) {
            _depositWeth(operators[i], operatorStakeAddition);
        }

        OperatorStateRetriever.Operator[][] memory operatorState;
        OmniAVS.Validator[] memory validators;

        omniAVS.syncWithEigenLayer();
        operatorState = omniAVS.getOperatorState();
        validators = omniAVS.getValidators();

        // assert all operators still registered
        assertEq(operatorState.length, 1); // only one quorum
        assertEq(operatorState[0].length, numOperators);
        assertEq(validators.length, numOperators);

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

            // operator state is delegations + stake
            assertEq(operatorState[0][i].stake, totalStaked + totalDelegated);

            // validator state tracks these separately
            assertEq(validators[i].staked, totalStaked);
            assertEq(validators[i].delegated, totalDelegated);
        }
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

        OperatorStateRetriever.Operator[][] memory operatorState;
        OmniAVS.Validator[] memory validators;

        omniAVS.syncWithEigenLayer();
        operatorState = omniAVS.getOperatorState();
        validators = omniAVS.getValidators();

        // assert all operators still registered
        assertEq(operatorState.length, 1); // only one quorum
        assertEq(operatorState[0].length, numOperators);
        assertEq(validators.length, numOperators);

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

            // operator state is delegations + stake
            assertEq(operatorState[0][i].stake, totalStaked);

            // validator state tracks these separately
            assertEq(validators[i].staked, totalStaked);
            assertEq(validators[i].delegated, 0);
        }
    }

    /// @dev Deregister operators, assert OmniAVS quorum is updated after each deregistration
    function _testDeregisterOperators() internal {
        OperatorStateRetriever.Operator[][] memory operatorState;
        OmniAVS.Validator[] memory validators;

        for (uint32 i = 0; i < numOperators; i++) {
            address operator = operators[i];

            _deregisterOperatorFromAVS(operators[i]);

            omniAVS.syncWithEigenLayer();
            operatorState = omniAVS.getOperatorState();
            validators = omniAVS.getValidators();

            uint96 numOperatorsLeft = numOperators - i - 1;

            // assert there are only numOperatorsLeft
            assertEq(operatorState.length, 1); // only one quorum
            assertEq(operatorState[0].length, numOperatorsLeft);
            assertEq(validators.length, numOperatorsLeft);

            // assert that none of the operators left is the operator that just deregistered
            for (uint32 j = 0; j < numOperatorsLeft; j++) {
                assertNotEq(operatorState[0][j].operator, operator);
                assertNotEq(validators[j].addr, operator);
            }
        }
    }
}
