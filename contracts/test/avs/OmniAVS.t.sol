// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

import { IOmniEthRestaking } from "src/interfaces/IOmniEthRestaking.sol";
import { OmniPredeploys } from "src/libraries/OmniPredeploys.sol";
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

    // feeForSync is msg.value passsed in tests OmniAVS.syncWithOmni()
    // normally, these fee will be calculated by the caller off-chain
    uint256 syncFee = 1 gwei;

    /**
     * Test OmniAVS.syncWithOmnie() at a number of different points in a "delegation lifecycle".
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

        OmniAVS.Validator[] memory validators;
        validators = omniAVS.getValidators();

        // assert no operator for omni avs quorum, because no operator has been registered
        assertEq(validators.length, 0); // no validators

        // TODO: should we revert if no operators are registered?
        // current thinking is no, to allow all operators to be deregistered
        _expectXCall(validators);
        omniAVS.syncWithOmni{ value: syncFee }();
    }

    /// @dev Register operators with OmniAVS, assert OmniAVS quorum is populated with initial stake
    function _testRegisterOperatorsWithAVS() internal {
        // register operators with AVS
        for (uint32 i = 0; i < numOperators; i++) {
            _registerOperatorWithAVS(operators[i]);
        }

        OmniAVS.Validator[] memory validators;
        validators = omniAVS.getValidators();

        // assert all operators have been registered
        assertEq(validators.length, numOperators);

        // assert operator has initial stake
        for (uint32 i = 0; i < numOperators; i++) {
            assertEq(validators[i].staked, initialOperatorStake);
            assertEq(validators[i].delegated, 0);
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
                _depositWeth(delegator, initialDelegatorStake);

                // should NOT contribute to quorom stake
                _depositEigen(delegator, initialDelegatorStake);

                // all stake is delegated
                _testDelegateToOperator(delegator, operator);
            }
        }

        OmniAVS.Validator[] memory validators;
        validators = omniAVS.getValidators();

        // assert all operators still registered
        assertEq(validators.length, numOperators);

        // assert all operator stake has been updated by initialDelegatorStake
        for (uint32 i = 0; i < numOperators; i++) {
            uint96 totalDelegated = numDelegatorsPerOp * initialDelegatorStake;
            uint96 totalStaked = initialOperatorStake;

            // validator state tracks these separately
            assertEq(validators[i].staked, totalStaked);
            assertEq(validators[i].delegated, totalDelegated);
        }

        _expectXCall(validators);
        omniAVS.syncWithOmni{ value: syncFee }();
    }

    /// @dev Increase delegations for first half of operators, assert OmniAVS quorum is updated
    function _testIncreaseDelegationsToFirstHalfOfOperators() internal {
        // increase delegations for first half of operators
        for (uint32 i = 0; i < numOperators / 2; i++) {
            for (uint32 j = 0; j < numDelegatorsPerOp; j++) {
                _depositWeth(delegators[i * numDelegatorsPerOp + j], delegatorStakeAddition);
            }
        }

        OmniAVS.Validator[] memory validators;
        validators = omniAVS.getValidators();

        // assert all operators still registered
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

            // validator state tracks these separately
            assertEq(validators[i].staked, totalStaked);
            assertEq(validators[i].delegated, totalDelegated);
        }

        _expectXCall(validators);
        omniAVS.syncWithOmni{ value: syncFee }();
    }

    /// @dev Increase stake for second half of operators, assert OmniAVS quorum is updated
    function _testIncreaseStakeOfSecondHalfOfOperators() internal {
        // increase stake of second half of delegators
        for (uint32 i = numOperators / 2; i < numOperators; i++) {
            _depositWeth(operators[i], operatorStakeAddition);
        }

        OmniAVS.Validator[] memory validators;
        validators = omniAVS.getValidators();

        // assert all operators still registered
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

            // validator state tracks these separately
            assertEq(validators[i].staked, totalStaked);
            assertEq(validators[i].delegated, totalDelegated);
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

            // validator state tracks these separately
            assertEq(validators[i].staked, totalStaked);
            assertEq(validators[i].delegated, 0);
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
                assertNotEq(validators[j].addr, operator);
            }
        }
    }

    /// @dev Expect an OmniPortal.xcall to IOmniEthRestaking.sync(validators), with correct fee and gasLimit
    function _expectXCall(OmniAVS.Validator[] memory validators) internal {
        bytes memory data = abi.encodeWithSelector(IOmniEthRestaking.sync.selector, validators);
        uint64 gasLimit = omniAVS.xcallGasLimitFor(validators.length);

        vm.expectCall(
            address(portal),
            syncFee,
            abi.encodeWithSignature(
                "xcall(uint64,address,bytes,uint64)", omniChainId, OmniPredeploys.OMNI_ETH_RESTAKING, data, gasLimit
            )
        );
    }
}
