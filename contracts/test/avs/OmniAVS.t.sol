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
    function testFuzz_getOperatorState_succeeds(uint8 numOperators, uint8 numDelegatorsPerOp) public {
        vm.assume(numOperators > 0 && numOperators < 10);
        vm.assume(numDelegatorsPerOp > 0 && numDelegatorsPerOp < 10);

        // Test Omni AVS uses the WETH strategy, configured in OmniDelegationAVSBase
        // WETH counts 1:1 for `stake`

        OperatorStateRetriever.Operator[][] memory avsOperatorState;
        OmniAVS.Validator[] memory validators;

        // register operators
        //
        // NOTE: it is not necessary for operator to have deposited minimumStakeForQuorum
        // other staker(s) could have deposited and delegated to the operator
        address[] memory operators = new address[](numOperators);
        for (uint8 i = 0; i < numOperators; i++) {
            operators[i] = _operator(i);
            _registerAsOperator(operators[i]);
            _depositWeth(operators[i], minimumStakeForQuorum);
        }

        // assert no operator for omni avs quorum, because no operator has been registered
        avsOperatorState = omniAVS.getOperatorState();
        validators = omniAVS.getValidators();

        assertEq(avsOperatorState.length, 1); // only one quorum
        assertEq(avsOperatorState[0].length, 0); // no operators
        assertEq(validators.length, 0); // no validators

        // register operators with AVS
        for (uint8 i = 0; i < numOperators; i++) {
            _registerOperatorWithAVS(operators[i]);
        }

        // assert all operators have been registered, and their stake is viewable by our avs
        // every operator has the same stake - minimumStakeForQuorum
        avsOperatorState = omniAVS.getOperatorState();
        validators = omniAVS.getValidators();

        assertEq(avsOperatorState.length, 1);
        assertEq(avsOperatorState[0].length, numOperators);
        assertEq(validators.length, numOperators);

        for (uint8 i = 0; i < numOperators; i++) {
            assertEq(avsOperatorState[0][i].stake, minimumStakeForQuorum);
            assertEq(validators[i].staked, minimumStakeForQuorum);
            assertEq(validators[i].delegated, 0);
        }

        // initialize delegators
        address[] memory delegators = new address[](numDelegatorsPerOp * numOperators);
        for (uint8 i = 0; i < numOperators; i++) {
            address operator = operators[i];

            for (uint8 j = 0; j < numDelegatorsPerOp; j++) {
                uint8 idx = i * numDelegatorsPerOp + j;
                address delegator = _delegator(idx);
                delegators[idx] = _delegator(idx);

                // should contribute to quorom stake
                _depositWeth(delegator, 1 ether);

                // should NOT contribute to quorom stake
                _depositEigen(delegator, 1 ether);

                // all stake is delegated
                _testDelegateToOperator(delegator, operator);
            }
        }

        // we need to update stake registry view of operator stake, before
        // querying getOperatorState / getValidators again.
        // operator list can be tracked in the AVS, and udpated before syncing with omni
        registryCoordinator.updateOperators(operators);

        // assert all operator stake has been updated by 1 ether per delegator
        avsOperatorState = omniAVS.getOperatorState();
        validators = omniAVS.getValidators();

        assertEq(avsOperatorState.length, 1); // only one quorum
        assertEq(avsOperatorState[0].length, numOperators);
        assertEq(validators.length, numOperators);
        for (uint8 i = 0; i < numOperators; i++) {
            uint96 totalDelegated = numDelegatorsPerOp * 1 ether;
            uint96 totalStaked = minimumStakeForQuorum;

            // operator state is delegations + stake
            assertEq(avsOperatorState[0][i].stake, minimumStakeForQuorum + totalDelegated);

            // validator state tracks these separately
            assertEq(validators[i].staked, totalStaked);
            assertEq(validators[i].delegated, totalDelegated);
        }
    }
}
