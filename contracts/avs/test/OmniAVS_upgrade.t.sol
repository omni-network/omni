// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { Base } from "./common/Base.sol";

import { OmniAVS } from "src/OmniAVS.sol";

import { IDelegationManager } from "src/ext/IDelegationManager.sol";
import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import {
    ITransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract OmniAVS_upgrade_Test is Base {
    function _checkAndHashState() internal view returns (bytes32) {
        address _avsDirectory = omniAVS.avsDirectory();
        uint32 _maxOperatorCount = omniAVS.maxOperatorCount();
        uint64 _omniChainId = omniAVS.omniChainId();
        uint64 _xcallGasLimitPerOperator = omniAVS.xcallGasLimitPerOperator();
        uint64 _xcallBaseGasLimit = omniAVS.xcallBaseGasLimit();
        uint96 _minOperatorStake = omniAVS.minOperatorStake();
        bool _allowlistEnabled = omniAVS.allowlistEnabled();
        address _ethStakeInbox = omniAVS.ethStakeInbox();
        address _omni = address(omniAVS.omni());

        return keccak256(
            abi.encode(
                _avsDirectory,
                _maxOperatorCount,
                _omniChainId,
                _xcallGasLimitPerOperator,
                _xcallBaseGasLimit,
                _minOperatorStake,
                _allowlistEnabled,
                _ethStakeInbox,
                _omni
            )
        );
    }

    function test_upgrade_succeeds() public {
        address impl =
            address(new OmniAVS(IDelegationManager(address(delegation)), IAVSDirectory(address(avsDirectory))));
        ITransparentUpgradeableProxy proxy = ITransparentUpgradeableProxy(address(omniAVS));

        bytes32 beforeUpgradeHash = _checkAndHashState();

        vm.prank(proxyAdminOwner);
        proxyAdmin.upgrade(proxy, impl);

        bytes32 afterUpgradeHash = _checkAndHashState();

        assertEq(proxyAdmin.getProxyImplementation(proxy), impl);
        assertEq(beforeUpgradeHash, afterUpgradeHash);
    }

    function test_upgrade_invalidProxyOwner_reverts() public {
        address impl =
            address(new OmniAVS(IDelegationManager(address(delegation)), IAVSDirectory(address(avsDirectory))));
        ITransparentUpgradeableProxy proxy = ITransparentUpgradeableProxy(address(omniAVS));

        vm.expectRevert("Ownable: caller is not the owner");
        proxyAdmin.upgrade(proxy, impl);
    }
}
