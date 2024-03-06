// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniAVS_serviceManager_Test
 * @dev Test suite for the AVS "service manager" interface - functions expected by eigenlayer's
 *      frontend, matching eigenlayer-middleware's IServiceManager. These include
 *          - getRestakeableStrategies()
 *          - getOperatorRestakedStrategies(address operator)
 *          - avsDirectory()
 */
contract OmniAVS_admin_Test is Base {
    IStrategy mockStrat1;
    IStrategy mockStrat2;
    IStrategy mockStrat3;

    uint96 mockStrat1Mul = 1e18;
    uint96 mockStrat2Mul = 2e18;
    uint96 mockStrat3Mul = 3e18;

    IOmniAVS.StrategyParam[] strategyParams;

    function setUp() public override {
        super.setUp();

        mockStrat1 = IStrategy(makeAddr("mockStrat1"));
        mockStrat2 = IStrategy(makeAddr("mockStrat2"));
        mockStrat3 = IStrategy(makeAddr("mockStrat3"));

        strategyParams.push(IOmniAVS.StrategyParam(mockStrat1, mockStrat1Mul));
        strategyParams.push(IOmniAVS.StrategyParam(mockStrat2, mockStrat2Mul));
        strategyParams.push(IOmniAVS.StrategyParam(mockStrat3, mockStrat3Mul));
    }

    /// @dev Test that getRestakeableStrategies() matches OmniAVS.strategyParams
    function test_getRestakeableStrategies_succeeds() public {
        _setStrategyParams(strategyParams);

        address[] memory restakeableStrategies = omniAVS.getRestakeableStrategies();

        assertEq(restakeableStrategies.length, 3);
        assertEq(restakeableStrategies[0], address(mockStrat1));
        assertEq(restakeableStrategies[1], address(mockStrat2));
        assertEq(restakeableStrategies[2], address(mockStrat3));
    }

    /// @dev Test that getRestakeableStrategies() returns an empty list when AVS.strategyParams is empty
    function test_getRestakeableStrategies_noStrategies_succeeds() public {
        _setStrategyParams(new IOmniAVS.StrategyParam[](0));

        address[] memory restakeableStrategies = omniAVS.getRestakeableStrategies();

        assertEq(restakeableStrategies.length, 0);
    }

    /// @dev Test that getOperatorRestakedStrategies() matches OmniAVS.strategyParams
    ///
    ///      Note that no work is done to determine if the operator has actually restaked in the strategy
    ///      This behavior matches that of the middleware's ServiceManagerBase.getOperatorRestakedStrategies()
    ///      In ServiceManagerBase, they return the aggregate list of strategies for the quorums the operator
    ///      is a member of. We only have one "quorum", so we return all strategies.
    function test_getOperatorRestakedStrategies_succeeds() public {
        address operator = _operator(0);
        _registerAsOperator(operator);
        _addToAllowlist(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);
        _registerOperatorWithAVS(operator);

        _setStrategyParams(strategyParams);
        address[] memory operatorRestakedStrategies = omniAVS.getOperatorRestakedStrategies(operator);

        assertEq(operatorRestakedStrategies.length, 3);
        assertEq(operatorRestakedStrategies[0], address(mockStrat1));
        assertEq(operatorRestakedStrategies[1], address(mockStrat2));
        assertEq(operatorRestakedStrategies[2], address(mockStrat3));
    }

    /// @dev Test that getOperatorRestakedStrategies() returns an empty list when AVS.strategyParams is empty
    function test_getOperatorRestakedStrategies_noStrategies_succeeds() public {
        address operator = _operator(0);
        _registerAsOperator(operator);
        _addToAllowlist(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);
        _registerOperatorWithAVS(operator);

        _setStrategyParams(new IOmniAVS.StrategyParam[](0));
        address[] memory operatorRestakedStrategies = omniAVS.getOperatorRestakedStrategies(operator);

        assertEq(operatorRestakedStrategies.length, 0);
    }

    /// @dev Test that getOperatorRestakedStrategies() returns an empty list when the operator is not registered
    function test_getOPeratorRestakedStrategies_notOperator_succeeds() public {
        address operator = _operator(0);
        address[] memory operatorRestakedStrategies = omniAVS.getOperatorRestakedStrategies(operator);

        assertEq(operatorRestakedStrategies.length, 0);
    }

    /// @dev Test that avsDirectory() returns the correct address
    function test_avsDirectory_succeeds() public {
        assertEq(omniAVS.avsDirectory(), address(avsDirectory));
    }

    /// @dev Set omniAVS.strategyParams
    function _setStrategyParams(IOmniAVS.StrategyParam[] memory _strategyParams) internal {
        vm.prank(omniAVSOwner);
        omniAVS.setStrategyParams(_strategyParams);
    }
}
