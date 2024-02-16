// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {
    ITransparentUpgradeableProxy,
    TransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { OmniAVSHarness } from "./OmniAVSHarness.sol";
import { EigenLayerTestHelper } from "./eigen/EigenLayerTestHelper.t.sol";
import { MockPortal } from "./MockPortal.sol";

contract AVSBase is EigenLayerTestHelper {
    address public omniAVSOwner = makeAddr("omniAVSOwner");
    address public proxyAdminOwner = makeAddr("proxyAdminOwner");

    uint32 maxOperatorCount = 10;
    uint96 minimumOperatorStake = 1 ether;
    uint64 omniChainId = 111;

    ProxyAdmin public proxyAdmin;

    MockPortal public portal;
    OmniAVSHarness public omniAVSImplementation;
    OmniAVSHarness public omniAVS;

    function setUp() public override {
        super.setUp();

        portal = new MockPortal();

        _deployProxyAdmin();
        _deployOmniAVS();
    }

    function _deployProxyAdmin() internal {
        vm.prank(proxyAdminOwner);
        proxyAdmin = new ProxyAdmin();
    }

    function _deployOmniAVS() internal {
        vm.startPrank(proxyAdminOwner);
        omniAVS =
            OmniAVSHarness(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));
        omniAVSImplementation = new OmniAVSHarness(delegation);
        proxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(omniAVS))), address(omniAVSImplementation));
        vm.stopPrank();

        IOmniAVS.StrategyParams[] memory strategyParams = _defaultStrategyParams();
        omniAVS.initialize(
            omniAVSOwner,
            IOmniPortal(address(portal)),
            omniChainId,
            minimumOperatorStake,
            maxOperatorCount,
            strategyParams
        );
    }

    function _defaultStrategyParams() internal view returns (IOmniAVS.StrategyParams[] memory params) {
        params = new IOmniAVS.StrategyParams[](1);
        params[0] = _wethStategyParams();
    }

    function _wethStategyParams() internal view returns (IOmniAVS.StrategyParams memory) {
        return IOmniAVS.StrategyParams({ strategy: wethStrat, multiplier: uint96(omniAVS.WEIGHTING_DIVISOR()) });
    }
}
