// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { IDelegationManager } from "src/interfaces/IDelegationManager.sol";

import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { OmniAVS } from "src/protocol/OmniAVS.sol";
import { EigenLayerTestHelper } from "./eigen/EigenLayerTestHelper.t.sol";
import { EigenM2GoerliDeployments } from "./eigen/EigenM2GoerliDeployments.sol";
import { MockPortal } from "./MockPortal.sol";

import { DeployGoerliAVS } from "script/avs/DeployGoerliAVS.s.sol";
import { StrategyParams } from "script/avs/StrategyParams.sol";
import { AVSDeploy } from "script/avs/AVSDeploy.sol";

contract AVSBase is EigenLayerTestHelper {
    address omniAVSOwner = makeAddr("omniAVSOwner");
    address proxyAdminOwner = makeAddr("proxyAdminOwner");

    uint32 maxOperatorCount = 10;
    uint96 minimumOperatorStake = 1 ether;
    uint64 omniChainId = 111;

    ProxyAdmin proxyAdmin;
    MockPortal portal;
    OmniAVS omniAVSImplementation;
    OmniAVS omniAVS;

    function setUp() public override {
        super.setUp();

        portal = new MockPortal();

        _deployProxyAdmin();

        if (isGoerli()) {
            _deployGoerliAVS();
        } else {
            _deployLocalAVS();
        }
    }

    function _deployProxyAdmin() private {
        vm.prank(proxyAdminOwner);
        proxyAdmin = new ProxyAdmin();
    }

    function _deployGoerliAVS() internal {
        DeployGoerliAVS deployer = new DeployGoerliAVS();
        omniAVS = OmniAVS(deployer.deploy(omniAVSOwner, address(proxyAdmin), address(portal), omniChainId));
    }

    function _deployLocalAVS() internal {
        vm.startPrank(proxyAdminOwner);
        omniAVS = OmniAVS(AVSDeploy.proxy(address(proxyAdmin)));
        omniAVSImplementation = OmniAVS(AVSDeploy.impl(address(delegation), address(avsDirectory)));

        IOmniAVS.StrategyParams[] memory strategyParams = _localStrategyParams();
        address[] memory allowlist = new address[](0);

        AVSDeploy.upgradeAndInit(
            address(proxyAdmin),
            address(omniAVS),
            address(omniAVSImplementation),
            // init params
            omniAVSOwner,
            address(portal),
            omniChainId,
            minimumOperatorStake,
            maxOperatorCount,
            allowlist,
            strategyParams
        );

        vm.stopPrank();
    }

    function _localStrategyParams() internal view returns (IOmniAVS.StrategyParams[] memory params) {
        params = new IOmniAVS.StrategyParams[](1);
        params[0] = IOmniAVS.StrategyParams({
            strategy: wethStrat,
            multiplier: uint96(omniAVSImplementation.WEIGHTING_DIVISOR())
        });
        return params;
    }
}
