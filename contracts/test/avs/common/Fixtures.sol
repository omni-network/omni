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

import { DeployGoerliAVS } from "script/avs/DeployGoerliAVS.s.sol";
import { StrategyParams } from "script/avs/StrategyParams.sol";
import { EigenLayerFixtures } from "./eigen/EigenLayerFixtures.sol";
import { Empty } from "./Empty.sol";
import { MockPortal } from "./MockPortal.sol";

/**
 * @title Fixtures
 * @dev Common fixtures contract for all AVS tests.
 */
contract Fixtures is EigenLayerFixtures {
    address multisig = makeAddr("multisig");
    address proxyAdminOwner = multisig;
    address omniAVSOwner = multisig;

    uint32 maxOperatorCount = 10;
    uint96 minOperatorStake = 1 ether;
    uint64 omniChainId = 111;

    ProxyAdmin proxyAdmin;
    MockPortal portal;
    OmniAVS omniAVS;

    /// Canonical, virtual beacon chain ETH strategy
    address constant beaconChainETHStrategy = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0;

    function setUp() public override {
        super.setUp();

        portal = new MockPortal();

        vm.prank(proxyAdminOwner);
        proxyAdmin = new ProxyAdmin();

        omniAVS = isGoerli() ? _deployGoerliAVS() : _deployLocalAVS();
    }

    function _deployGoerliAVS() internal returns (OmniAVS) {
        DeployGoerliAVS deployer = new DeployGoerliAVS();
        return OmniAVS(
            deployer.prankDeploy(proxyAdminOwner, omniAVSOwner, address(proxyAdmin), address(portal), omniChainId)
        );
    }

    function _deployLocalAVS() internal returns (OmniAVS) {
        vm.startPrank(proxyAdminOwner);

        address proxy = address(new TransparentUpgradeableProxy(address(new Empty()), address(proxyAdmin), ""));
        address impl =
            address(new OmniAVS(IDelegationManager(address(delegation)), IAVSDirectory(address(avsDirectory))));

        address[] memory allowlist = new address[](0);

        ProxyAdmin(proxyAdmin).upgradeAndCall(
            ITransparentUpgradeableProxy(proxy),
            impl,
            abi.encodeWithSelector(
                OmniAVS.initialize.selector,
                omniAVSOwner,
                portal,
                omniChainId,
                minOperatorStake,
                maxOperatorCount,
                allowlist,
                _localStrategyParams()
            )
        );
        vm.stopPrank();

        return OmniAVS(proxy);
    }

    function _localStrategyParams() internal view returns (IOmniAVS.StrategyParams[] memory params) {
        // add all EigenLayerDeployer.strategies
        params = new IOmniAVS.StrategyParams[](strategies.length);

        for (uint256 i = 0; i < strategies.length; i++) {
            params[i] = IOmniAVS.StrategyParams({
                strategy: IStrategy(strategies[i]),
                multiplier: uint96(1e18) // OmniAVS.STRATEGY_WEIGHTING_DIVISOR
             });
        }

        return params;
    }
}
