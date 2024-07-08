// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { IEthStakeInbox } from "core/interfaces/IEthStakeInbox.sol";
import { ConfLevel } from "core/libraries/ConfLevel.sol";

import { Base } from "./common/Base.sol";

/**
 * @title OmniAVS_syncWithOmni_Test
 * @dev Test suite for OmniAVS.syncWithOmni(), and by extension, OmniAVS.operators()
 */
contract OmniAVS_syncWithOmni_Test is Base {
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
     * For each point in delegation lifecycle, test that OmniAVS.operators() returns the
     * expected ops list, and that OmniAVS.syncWithOmni() makes a call to
     * OmniPortal.xcall with that list of ops.
     *
     * forge-config: default.fuzz.runs = 10
     */
    function testFuzz_syncWithOmni(
        uint8 numOperators_,
        uint8 numDelegatorsPerOp_,
        uint96 initialOperatorStake_,
        uint96 initialDelegatorStake_
    ) public {
        numOperators = uint32(bound(numOperators_, 2, maxOperatorCount));
        numDelegatorsPerOp = uint32(bound(numDelegatorsPerOp_, 1, 30));

        initialOperatorStake = uint96(bound(initialOperatorStake_, minOperatorStake, 100 ether));
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
        _testEjectOperators();
    }

    /// @dev Register operators with eigen layer core, assert OmniAVS quorum is still empty
    function _testRegisterOperators() internal {
        // NOTE: it is not necessary for operator to have deposited minimumStakeForQuorum
        // other staker(s) could have deposited and delegated to the operator
        for (uint32 i = 0; i < numOperators; i++) {
            _registerAsOperator(operators[i]);
            _depositIntoSupportedStrategy(operators[i], initialOperatorStake);
        }

        IOmniAVS.Operator[] memory ops;
        ops = omniAVS.operators();

        // assert no operator for omni avs quorum, because no operator has been registered
        assertEq(ops.length, 0, "_testRegisterOperators: no operators should be registered");

        _assertSyncWithOmni(ops);
    }

    /// @dev Register operators with OmniAVS, assert OmniAVS quorum is populated with initial stake
    function _testRegisterOperatorsWithAVS() internal {
        // register operators with AVS
        for (uint32 i = 0; i < numOperators; i++) {
            _addToAllowlist(operators[i]);
            _registerOperatorWithAVS(operators[i]);
        }

        IOmniAVS.Operator[] memory ops;
        ops = omniAVS.operators();

        // assert all operators have been registered
        assertEq(ops.length, numOperators, "_testRegisterOperatorsWithAVS: all operators should be registered");

        // assert operator has initial stake
        for (uint32 i = 0; i < numOperators; i++) {
            assertEq(
                ops[i].staked,
                initialOperatorStake,
                "_testRegisterOperatorsWithAVS: validator.staked should be initialOperatorStake"
            );
            assertEq(ops[i].delegated, 0, "_testRegisterOperatorsWithAVS: validator.delegated should be 0");
        }

        _assertSyncWithOmni(ops);
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

        IOmniAVS.Operator[] memory ops;
        ops = omniAVS.operators();

        // assert all operators still registered
        assertEq(ops.length, numOperators, "_testDelegateToOperators: all operators should still be registered");

        // assert all operator stake has been updated by initialDelegatorStake
        for (uint32 i = 0; i < numOperators; i++) {
            uint96 totalDelegated = numDelegatorsPerOp * initialDelegatorStake;
            uint96 totalStaked = initialOperatorStake;

            // validator state tracks these separately
            assertEq(ops[i].staked, totalStaked, "_testDelegateToOperators: validator.staked unexpected");
            assertEq(ops[i].delegated, totalDelegated, "_testDelegateToOperators: validator.delegated unexpected");
        }

        _assertSyncWithOmni(ops);
    }

    /// @dev Increase delegations for first half of operators, assert OmniAVS quorum is updated
    function _testIncreaseDelegationsToFirstHalfOfOperators() internal {
        // increase delegations for first half of operators
        for (uint32 i = 0; i < numOperators / 2; i++) {
            for (uint32 j = 0; j < numDelegatorsPerOp; j++) {
                _depositIntoSupportedStrategy(delegators[i * numDelegatorsPerOp + j], delegatorStakeAddition);
            }
        }

        IOmniAVS.Operator[] memory ops;
        ops = omniAVS.operators();

        // assert all operators still registered
        assertEq(
            ops.length,
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
                ops[i].staked,
                totalStaked,
                "_testIncreaseDelegationsToFirstHalfOfOperators: validator.staked unexpected"
            );
            assertEq(
                ops[i].delegated,
                totalDelegated,
                "_testIncreaseDelegationsToFirstHalfOfOperators: validator.delegated unexpected"
            );
        }

        _assertSyncWithOmni(ops);
    }

    /// @dev Increase stake for second half of operators, assert OmniAVS quorum is updated
    function _testIncreaseStakeOfSecondHalfOfOperators() internal {
        // increase stake of second half of delegators
        for (uint32 i = numOperators / 2; i < numOperators; i++) {
            _depositIntoSupportedStrategy(operators[i], operatorStakeAddition);
        }

        IOmniAVS.Operator[] memory ops;
        ops = omniAVS.operators();

        // assert all operators still registered
        assertEq(
            ops.length,
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
                ops[i].staked, totalStaked, "_testIncreaseStakeOfSecondHalfOfOperators: validator.staked unexpected "
            );
            assertEq(
                ops[i].delegated,
                totalDelegated,
                "_testIncreaseStakeOfSecondHalfOfOperators: validator.delegated unexpected"
            );
        }

        _assertSyncWithOmni(ops);
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

        IOmniAVS.Operator[] memory ops;
        ops = omniAVS.operators();

        // assert all operators still registered
        assertEq(ops.length, numOperators, "_testUndelegateAllDelegators: all operators should still be registered");

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
            assertEq(ops[i].staked, totalStaked, "_testUndelegateAllDelegators: validator.staked unexpected ");
            assertEq(ops[i].delegated, 0, "_testtUndelegateAllDelegators: validator.delegated should be 0");
        }

        _assertSyncWithOmni(ops);
    }

    /// @dev Eject operators, assert OmniAVS quorum is updated after each deregistration
    function _testEjectOperators() internal {
        IOmniAVS.Operator[] memory ops;

        for (uint32 i = 0; i < numOperators; i++) {
            address operator = operators[i];

            _ejectOperatorFromAVS(operator);

            ops = omniAVS.operators();

            uint96 numOperatorsLeft = numOperators - i - 1;

            // assert there are only numOperatorsLeft
            assertEq(ops.length, numOperatorsLeft);

            // assert that none of the operators left is the operator that was just deregistered
            for (uint32 j = 0; j < numOperatorsLeft; j++) {
                assertNotEq(ops[j].addr, operator, "_testEjectOperators: operator should not be in validators list");
            }
        }
    }

    /// @dev Assert syncWithOmni() makes an xcall to OmniPortal with correct ops
    function _assertSyncWithOmni(IOmniAVS.Operator[] memory ops) internal {
        // skip fork tests in which portal is not set
        if (address(omniAVS.omni()) == address(0)) return;

        // portal not deployed on mainnet (but address is not 0)
        if (isMainnet()) return;

        _expectXCall(ops);
        omniAVS.syncWithOmni{ value: syncFee }();
    }

    /// @dev Expect an OmniPortal.xcall to IOmniEthRestaking.sync(ops), with correct fee and gasLimit
    function _expectXCall(IOmniAVS.Operator[] memory ops) internal {
        bytes memory data = abi.encodeWithSelector(IEthStakeInbox.sync.selector, ops);
        uint64 gasLimit = omniAVS.xcallBaseGasLimit() + omniAVS.xcallGasLimitPerOperator() * uint64(ops.length);

        vm.expectCall(
            address(portal),
            syncFee,
            abi.encodeWithSignature(
                "xcall(uint64,uint8,address,bytes,uint64)",
                omniChainId,
                ConfLevel.Finalized,
                ethStakeInbox,
                data,
                gasLimit
            )
        );
    }

    /// @dev Unit test for beacon eth deposit
    function test_depositBeaconEth_succeeds() public {
        address operator = _operator(0);
        uint96 amount = 1 ether;

        _registerAsOperator(operator);
        _addToAllowlist(operator);
        _depositBeaconEth(operator, amount);
        _registerOperatorWithAVS(operator);

        IOmniAVS.Operator[] memory ops = omniAVS.operators();

        assertEq(ops.length, 1);
        assertEq(ops[0].addr, operator);
        assertEq(ops[0].staked, amount);
        assertEq(ops[0].delegated, 0);

        _assertSyncWithOmni(ops);
    }

    /// @dev Test that delegations (self & other) to unsupported strategies are not counted in AVS stake
    function test_unsupportedStrategyDeposit_succeeds() public {
        address operator = _operator(0);
        address delegator = _delegator(0);

        uint96 stakeAmt = 10 ether;
        uint96 delegateAmt = 1 ether;

        _registerAsOperator(operator);

        // should be counted by avs
        _depositIntoSupportedStrategy(operator, stakeAmt);
        _depositIntoSupportedStrategy(delegator, delegateAmt);

        // should NOT be counted by avs
        _depositIntoUnsupportedStrategy(operator, stakeAmt);
        _depositIntoUnsupportedStrategy(delegator, delegateAmt);

        _testDelegateToOperator(delegator, operator);
        _addToAllowlist(operator);
        _registerOperatorWithAVS(operator);

        IOmniAVS.Operator[] memory ops = omniAVS.operators();

        // assert unsupported strategy deposit does not affect operator stake
        assertEq(ops.length, 1);
        assertEq(ops[0].addr, operator);
        assertEq(ops[0].staked, stakeAmt);
        assertEq(ops[0].delegated, delegateAmt);
    }
}
