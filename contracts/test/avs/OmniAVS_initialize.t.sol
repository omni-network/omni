// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";

import { IDelegationManager } from "src/interfaces/IDelegationManager.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { OmniAVS } from "src/protocol/OmniAVS.sol";

import { Empty } from "./common/Empty.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniAVS_initialize_Test
 * @dev Test suite for the AVS initialization
 */
contract OmniAVS_initialize_Test is Base {
    struct InitializeParams {
        address owner;
        IOmniPortal omni;
        uint64 omniChainId;
        uint96 minOperatorStake;
        uint32 maxOperatorCount;
        address[] allowlist;
        IOmniAVS.StrategyParams[] strategyParams;
    }

    function _defaultInitializeParams() internal view returns (InitializeParams memory) {
        return InitializeParams({
            owner: omniAVSOwner,
            omni: IOmniPortal(address(portal)),
            omniChainId: omniChainId,
            minOperatorStake: minOperatorStake,
            maxOperatorCount: maxOperatorCount,
            allowlist: new address[](0),
            strategyParams: _localStrategyParams()
        });
    }

    /// @dev Deploy a new OmniAVS proxy and initialize it with the given parameters
    function _deployAndInitialize(InitializeParams memory params) internal returns (OmniAVS) {
        vm.startPrank(proxyAdminOwner);

        address proxy = address(new TransparentUpgradeableProxy(address(new Empty()), address(proxyAdmin), ""));
        address impl =
            address(new OmniAVS(IDelegationManager(address(delegation)), IAVSDirectory(address(avsDirectory))));

        ProxyAdmin(proxyAdmin).upgradeAndCall(
            ITransparentUpgradeableProxy(proxy),
            impl,
            abi.encodeWithSelector(
                OmniAVS.initialize.selector,
                params.owner,
                params.omni,
                params.omniChainId,
                params.minOperatorStake,
                params.maxOperatorCount,
                params.allowlist,
                params.strategyParams
            )
        );
        vm.stopPrank();

        return OmniAVS(proxy);
    }

    /// @dev Test that the default initialization parameters are set correctly
    function test_initialize_defaultParams_succeeds() public {
        InitializeParams memory params = _defaultInitializeParams();
        OmniAVS omniAVS = _deployAndInitialize(params);

        assertEq(omniAVS.owner(), params.owner);
        assertEq(address(omniAVS.omni()), address(params.omni));
        assertEq(omniAVS.omniChainId(), params.omniChainId);
        assertEq(omniAVS.minOperatorStake(), params.minOperatorStake);
        assertEq(omniAVS.maxOperatorCount(), params.maxOperatorCount);

        IOmniAVS.StrategyParams[] memory strategyParams = omniAVS.strategyParams();
        assertEq(strategyParams.length, params.strategyParams.length);
        for (uint256 i = 0; i < strategyParams.length; i++) {
            assertEq(address(strategyParams[i].strategy), address(params.strategyParams[i].strategy));
            assertEq(strategyParams[i].multiplier, params.strategyParams[i].multiplier);
        }

        assertEq(omniAVS.paused(), false); // assert not paused
    }

    // TODO: add more initialization tests
}
