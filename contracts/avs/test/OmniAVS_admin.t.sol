// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

import { IOmniPortal } from "core/interfaces/IOmniPortal.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniAVS_admin_Test
 * @dev Test suite for the AVS admin functionality
 */
contract OmniAVS_admin_Test is Base {
    /// @dev Test that the owner can eject an operator
    function test_ejectOwner_byOwner_succeeds() public {
        address operator = _operator(0);

        // register operator
        _registerAsOperator(operator);
        _addToAllowlist(operator);
        _depositBeaconEth(operator, minOperatorStake);
        _registerOperatorWithAVS(operator);

        // assert operator is registered
        IOmniAVS.Operator[] memory operators = omniAVS.operators();
        assertEq(operators.length, 1);
        assertEq(operators[0].addr, operator);

        // deregister operator
        vm.prank(omniAVSOwner);
        omniAVS.ejectOperator(operator);

        // assert operator is deregistered
        operators = omniAVS.operators();
        assertEq(operators.length, 0);
    }

    /// @dev Test that only the owner can eject an operator
    function test_ejectOwner_notOwner_reverts() public {
        address operator = _operator(0);
        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.ejectOperator(operator);
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

    /// @dev Test that a stratey cannot have a multiplier of 0
    function test_setStrategyParams_zeroMul_reverts() public {
        IOmniAVS.StrategyParam[] memory params = new IOmniAVS.StrategyParam[](1);
        params[0] = IOmniAVS.StrategyParam({ strategy: IStrategy(address(1)), multiplier: 0 });

        vm.prank(omniAVSOwner);
        vm.expectRevert("OmniAVS: no zero multiplier");
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

    /// @dev Emitted by the avs directory when an AVS's metadata URI is updated
    event AVSMetadataURIUpdated(address indexed avs, string metadataURI);

    /// @dev Test setMetadataURI sets the metadata URI with the avs directory
    function test_setMetadataURI_succeeds() public {
        string memory uri = "https://example.com/avs/";

        vm.expectEmit(address(avsDirectory));
        emit AVSMetadataURIUpdated(address(omniAVS), uri);

        vm.prank(omniAVSOwner);
        omniAVS.setMetadataURI(uri);
    }

    /// @dev Test that only the owner can set the metadata URI
    function test_setMetadataURI_notOwner_reverts() public {
        string memory uri = "https://example.com/avs/";
        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.setMetadataURI(uri);
    }

    /// @dev Test that the owner can set omniChainId
    function test_setOmniChainId_succeeds() public {
        uint64 chainId = 1;
        vm.prank(omniAVSOwner);
        omniAVS.setOmniChainId(chainId);
        assertEq(omniAVS.omniChainId(), chainId);
    }

    /// @dev Test that only the owner can set omniChainId
    function test_setOmniChainId_notOwner_reverts() public {
        uint64 chainId = 1;
        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.setOmniChainId(chainId);
    }

    /// @dev Test that the owner can set ethStakeInbox
    function test_setEthStakeInbox_succeeds() public {
        address ethStakeInbox = address(1);
        vm.prank(omniAVSOwner);
        omniAVS.setEthStakeInbox(ethStakeInbox);
        assertEq(omniAVS.ethStakeInbox(), ethStakeInbox);
    }

    /// @dev Test that only the owner can set ethStakeInbox
    function test_setEthStakeInbox_notOwner_reverts() public {
        address ethStakeInbox = address(1);
        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.setEthStakeInbox(ethStakeInbox);
    }

    /// @dev Test that the ethStakeInbox cannot be the 0 address
    function test_setEthStakeInbox_zeroAddress_reverts() public {
        vm.expectRevert("OmniAVS: no zero inbox");
        vm.prank(omniAVSOwner);
        omniAVS.setEthStakeInbox(address(0));
    }

    /// @dev Test thath the owner can set the portal address
    function test_setPortal_succeeds() public {
        address portal = address(1);
        vm.prank(omniAVSOwner);
        omniAVS.setOmniPortal(IOmniPortal(portal));
        assertEq(address(omniAVS.omni()), portal);
    }

    /// @dev Test that only the owner can set the portal address
    function test_setPortal_notOwner_reverts() public {
        address portal = address(1);
        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.setOmniPortal(IOmniPortal(portal));
    }

    /// @dev Test that the portal address cannot be the 0 address
    function test_setPortal_zeroAddress_reverts() public {
        vm.expectRevert("OmniAVS: no zero portal");
        vm.prank(omniAVSOwner);
        omniAVS.setOmniPortal(IOmniPortal(address(0)));
    }

    /// @dev Test that the owner can set the xcall gas limit params
    function test_setXCallGasLimits_succeeds() public {
        uint64 base = omniAVS.xcallBaseGasLimit() + 10_000;
        uint64 perOperator = omniAVS.xcallGasLimitPerOperator() + 20_000;

        vm.prank(omniAVSOwner);
        omniAVS.setXCallGasLimits(base, perOperator);

        assertEq(omniAVS.xcallBaseGasLimit(), base);
        assertEq(omniAVS.xcallGasLimitPerOperator(), perOperator);
    }

    /// @dev Test that only the owner can set the xcall gas limit params
    function test_setXCallGasLimits_notOwner_reverts() public {
        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.setXCallGasLimits(0, 0);
    }

    /// @dev Test that the owner can set the min operator stake
    function test_setMinOperatorStake_succeeds() public {
        uint96 stake = minOperatorStake + 10_000;

        vm.prank(omniAVSOwner);
        omniAVS.setMinOperatorStake(stake);
        assertEq(omniAVS.minOperatorStake(), stake);
    }

    /// @dev Test that only the owner can set the min operator stake
    function test_setMinOperatorStake_notOwner_reverts() public {
        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.setMinOperatorStake(0);
    }

    /// @dev Test that the owner can set the max operator count
    function test_setMaxOperatorCount_succeeds() public {
        uint32 count = omniAVS.maxOperatorCount() + 10;

        vm.prank(omniAVSOwner);
        omniAVS.setMaxOperatorCount(count);
        assertEq(omniAVS.maxOperatorCount(), count);
    }

    /// @dev Test that only the owner can set the max operator count
    function test_setMaxOperatorCount_notOwner_reverts() public {
        vm.expectRevert("Ownable: caller is not the owner");
        omniAVS.setMaxOperatorCount(0);
    }
}
