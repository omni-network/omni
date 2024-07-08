// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";

import { IDelegationManager } from "src/ext/IDelegationManager.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { OmniAVS } from "src/OmniAVS.sol";
import { IOmniPortal } from "core/interfaces/IOmniPortal.sol";

import { Empty } from "./common/Empty.sol";
import { Base } from "./common/Base.sol";

import { Vm } from "forge-std/Vm.sol";

/**
 * @title OmniAVS_initialize_Test
 * @dev Test suite for the AVS initialization
 */
contract OmniAVS_initialize_Test is Base {
    struct InitializeParams {
        address owner;
        IOmniPortal omni;
        uint64 omniChainId;
        address ethStakeInbox;
        uint96 minOperatorStake;
        uint32 maxOperatorCount;
        IOmniAVS.StrategyParam[] strategyParams;
        string metadataURI;
        bool allowlistEnabled;
    }

    function _defaultInitializeParams() internal view returns (InitializeParams memory) {
        return InitializeParams({
            owner: omniAVSOwner,
            omni: IOmniPortal(address(portal)),
            omniChainId: omniChainId,
            ethStakeInbox: address(1234),
            minOperatorStake: minOperatorStake,
            maxOperatorCount: maxOperatorCount,
            strategyParams: _localStrategyParams(),
            allowlistEnabled: true,
            metadataURI: metadataURI
        });
    }

    /// @dev Deploy a new OmniAVS proxy and initialize it with the given parameters
    function _deployAndInitialize(InitializeParams memory params) internal returns (OmniAVS) {
        vm.startPrank(proxyAdminOwner);

        address proxy = address(new TransparentUpgradeableProxy(address(new Empty()), address(proxyAdmin), ""));
        address impl =
            address(new OmniAVS(IDelegationManager(address(delegation)), IAVSDirectory(address(avsDirectory))));

        ProxyAdmin(proxyAdmin).upgradeAndCall(ITransparentUpgradeableProxy(proxy), impl, _initializer(params));
        vm.stopPrank();

        return OmniAVS(proxy);
    }

    function _initializer(InitializeParams memory params) internal pure returns (bytes memory) {
        return abi.encodeWithSelector(
            OmniAVS.initialize.selector,
            params.owner,
            params.omni,
            params.omniChainId,
            params.ethStakeInbox,
            params.minOperatorStake,
            params.maxOperatorCount,
            params.strategyParams,
            params.metadataURI,
            params.allowlistEnabled
        );
    }

    /// @dev Test that the default initialization parameters are set correctly
    function test_initialize_defaultParams_succeeds() public {
        InitializeParams memory params = _defaultInitializeParams();

        vm.recordLogs();
        OmniAVS avs = _deployAndInitialize(params);
        _assertMetadataURIUpdated(vm.getRecordedLogs(), params.metadataURI, address(avs));

        assertEq(avs.owner(), params.owner);
        assertEq(address(avs.omni()), address(params.omni));
        assertEq(avs.omniChainId(), params.omniChainId);
        assertEq(avs.ethStakeInbox(), params.ethStakeInbox);
        assertEq(avs.minOperatorStake(), params.minOperatorStake);
        assertEq(avs.maxOperatorCount(), params.maxOperatorCount);
        assertEq(avs.allowlistEnabled(), params.allowlistEnabled);

        IOmniAVS.StrategyParam[] memory strategyParams = avs.strategyParams();
        assertEq(strategyParams.length, params.strategyParams.length);
        for (uint256 i = 0; i < strategyParams.length; i++) {
            assertEq(address(strategyParams[i].strategy), address(params.strategyParams[i].strategy));
            assertEq(strategyParams[i].multiplier, params.strategyParams[i].multiplier);
        }

        assertFalse(avs.paused());
    }

    function _assertMetadataURIUpdated(Vm.Log[] memory logs, string memory metadataURI, address avs) internal {
        bool sawMetadataURIUpdated = false;
        for (uint256 i = 0; i < logs.length; i++) {
            Vm.Log memory log = logs[i];
            if (
                log.emitter == address(avsDirectory)
                    && log.topics[0] == keccak256("AVSMetadataURIUpdated(address,string)")
            ) {
                assertEq(avs, address(uint160(uint256(log.topics[1]))));
                assertEq(metadataURI, abi.decode(log.data, (string)));
                sawMetadataURIUpdated = true;
            }
        }
        assertTrue(sawMetadataURIUpdated);
    }
}
