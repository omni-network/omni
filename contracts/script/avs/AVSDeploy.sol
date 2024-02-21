// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IDelegationManager } from "src/interfaces/IDelegationManager.sol";

import { Empty } from "test/common/Empty.sol";
import { OmniAVS } from "src/protocol/OmniAVS.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";

library AVSDeploy {
    /// @dev Deploys an OmniAVS implementation contract
    function impl(address delegationManager, address avsDirectory) internal returns (address) {
        return address(new OmniAVS(IDelegationManager(delegationManager), IAVSDirectory(avsDirectory)));
    }

    /// Deploy a TransparentUpgradeableProxy with an Empty implementation
    function proxy(address admin) internal returns (address) {
        return address(new TransparentUpgradeableProxy(address(new Empty()), admin, ""));
    }

    /// @dev Upgrades a proxy to the OmniAVS implementation, initializing it with the given parameters
    function upgradeAndInit(
        address proxyAdmin,
        address proxy_,
        address impl_,
        // initialize params
        address owner,
        address portal,
        uint64 omniChainId,
        uint96 minimumOperatorStake,
        uint32 maxOperatorCount,
        address[] memory allowlist,
        IOmniAVS.StrategyParams[] memory strategyParams
    ) internal {
        ProxyAdmin(proxyAdmin).upgradeAndCall(
            ITransparentUpgradeableProxy(proxy_),
            impl_,
            abi.encodeWithSelector(
                OmniAVS.initialize.selector,
                owner,
                portal,
                omniChainId,
                minimumOperatorStake,
                maxOperatorCount,
                allowlist,
                strategyParams
            )
        );
    }
}
