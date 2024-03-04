// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniAVS_admin_Test
 * @dev Test suite for the AVS admin functionality
 */
contract OmniAVS_admin_Test is Base {
    /// @dev Test that the owner can deregister an operator
    function test_deregisterOperator_byOwner_succeeds() public {
        address operator = _operator(0);

        // register operator
        _registerAsOperator(operator);
        _addToAllowlist(operator);
        _depositBeaconEth(operator, 1 ether);
        _registerOperatorWithAVS(operator);

        // assert operator is registered
        IOmniAVS.Operator[] memory operators = omniAVS.operators();
        assertEq(operators.length, 1);
        assertEq(operators[0].addr, operator);

        // deregister operator
        vm.prank(omniAVSOwner);
        omniAVS.deregisterOperatorFromAVS(operator);

        // assert operator is deregistered
        operators = omniAVS.operators();
        assertEq(operators.length, 0);
    }

    /// @dev Test that the owner can set the strategy parameters
    function test_setStrategyParams_succeds() public {
        IOmniAVS.StrategyParam[] memory params = new IOmniAVS.StrategyParam[](3);
        params[0] = IOmniAVS.StrategyParam({ strategy: IStrategy(address(1)), multiplier: 100 });
        params[1] = IOmniAVS.StrategyParam({ strategy: IStrategy(address(2)), multiplier: 200 });
        params[2] = IOmniAVS.StrategyParam({ strategy: IStrategy(address(3)), multiplier: 300 });

        vm.prank(omniAVSOwner);
        omniAVS.setStrategyParams(params);

        IOmniAVS.StrategyParam[] memory actualParams = omniAVS.strategyParams();
        assertEq(actualParams.length, 3);
        assertEq(address(actualParams[0].strategy), address(1));
        assertEq(actualParams[0].multiplier, 100);
        assertEq(address(actualParams[1].strategy), address(2));
        assertEq(actualParams[1].multiplier, 200);
        assertEq(address(actualParams[2].strategy), address(3));
        assertEq(actualParams[2].multiplier, 300);
    }

    /// @dev Test that only the owner can set the strategy parameters
    function test_setStrategyParams_notOwner_reverts() public {
        IOmniAVS.StrategyParam[] memory params;

        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.setStrategyParams(params);
    }

    /// @dev Test that a stratey cannot be the 0 address
    function test_setStrategyParams_zeroAddress_reverts() public {
        IOmniAVS.StrategyParam[] memory params = new IOmniAVS.StrategyParam[](1);
        params[0] = IOmniAVS.StrategyParam({ strategy: IStrategy(address(0)), multiplier: 100 });

        vm.prank(omniAVSOwner);
        vm.expectRevert("OmniAVS: no zero strategy");
        omniAVS.setStrategyParams(params);
    }

    /// @dev Test that there cannot be duplicate strategies
    function test_setStrategyParams_duplicateStrategy_reverts() public {
        IOmniAVS.StrategyParam[] memory params = new IOmniAVS.StrategyParam[](2);
        params[0] = IOmniAVS.StrategyParam({ strategy: IStrategy(address(1)), multiplier: 100 });
        params[1] = IOmniAVS.StrategyParam({ strategy: IStrategy(address(1)), multiplier: 200 });

        vm.prank(omniAVSOwner);
        vm.expectRevert("OmniAVS: no duplicate strategy");
        omniAVS.setStrategyParams(params);
    }

    /// @dev Test that the owner can set the strategy parameters twice
    function test_setStrategyParams_twice_succeeds() public {
        IOmniAVS.StrategyParam[] memory params = new IOmniAVS.StrategyParam[](2);
        params[0] = IOmniAVS.StrategyParam({ strategy: IStrategy(address(1)), multiplier: 100 });
        params[1] = IOmniAVS.StrategyParam({ strategy: IStrategy(address(2)), multiplier: 200 });

        vm.prank(omniAVSOwner);
        omniAVS.setStrategyParams(params);

        params = new IOmniAVS.StrategyParam[](2);
        params[0] = IOmniAVS.StrategyParam({ strategy: IStrategy(address(3)), multiplier: 300 });
        params[1] = IOmniAVS.StrategyParam({ strategy: IStrategy(address(4)), multiplier: 400 });

        vm.prank(omniAVSOwner);
        omniAVS.setStrategyParams(params);

        IOmniAVS.StrategyParam[] memory actualParams = omniAVS.strategyParams();
        assertEq(actualParams.length, 2);
        assertEq(address(actualParams[0].strategy), address(3));
        assertEq(actualParams[0].multiplier, 300);
        assertEq(address(actualParams[1].strategy), address(4));
        assertEq(actualParams[1].multiplier, 400);
    }
}
