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
import { Create3 } from "core/deploy/Create3.sol";

import { DeployAVS as MainnetAVS } from "script/avs/DeployAVS.s.sol";
import { HoleskyAVS } from "./HoleskyAVS.sol";
import { StrategyParams } from "./StrategyParams.sol";
import { MockPortal } from "core-test/utils/MockPortal.sol";
import { EigenLayerFixtures } from "./eigen/EigenLayerFixtures.sol";
import { Empty } from "./Empty.sol";

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
    address ethStakeInbox = address(1234); // replace with actual address when implemented in core

    ProxyAdmin proxyAdmin;
    MockPortal portal;
    OmniAVS omniAVS;

    bool allowlistEnabled = true;
    string metadataURI = "https://raw.githubusercontent.com/omni-network/omni/main/static/avs-metadata.json";

    /// Canonical, virtual beacon chain ETH strategy
    address constant beaconChainETHStrategy = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0;

    function setUp() public virtual override {
        super.setUp();

        portal = new MockPortal();

        vm.prank(proxyAdminOwner);
        proxyAdmin = new ProxyAdmin();

        // allow avs contract to be set by env for, for fork tests
        // proxyAdmin override not required
        address avsOverride = vm.envOr("AVS_OVERRIDE", address(0));
        if (avsOverride != address(0)) {
            omniAVS = OmniAVS(avsOverride);
            omniAVSOwner = omniAVS.owner();
            return;
        }

        if (isMainnet()) {
            omniAVS = _deployMainnet();
            return;
        }

        if (isHolesky()) {
            omniAVS = _deployHoleskyAVS();
            return;
        }

        omniAVS = _deployLocalAVS();
    }

    function _deployMainnet() internal returns (OmniAVS) {
        MainnetAVS deployer = new MainnetAVS();

        return deployer.deploy(
            omniAVSOwner,
            address(proxyAdmin),
            address(portal),
            omniChainId,
            ethStakeInbox,
            minOperatorStake,
            maxOperatorCount
        );
    }

    function _deployHoleskyAVS() internal returns (OmniAVS) {
        HoleskyAVS deployer = new HoleskyAVS();

        Create3 create3 = new Create3();
        bytes32 create3Salt = keccak256("avs-holesky-fork-test");

        return deployer.deploy(
            address(create3),
            create3Salt,
            omniAVSOwner,
            address(proxyAdmin),
            address(portal),
            omniChainId,
            ethStakeInbox,
            minOperatorStake,
            maxOperatorCount
        );
    }

    function _deployLocalAVS() internal returns (OmniAVS) {
        address impl =
            address(new OmniAVS(IDelegationManager(address(delegation)), IAVSDirectory(address(avsDirectory))));

        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            impl,
            address(proxyAdmin),
            abi.encodeWithSelector(
                OmniAVS.initialize.selector,
                omniAVSOwner,
                address(portal),
                omniChainId,
                ethStakeInbox,
                minOperatorStake,
                maxOperatorCount,
                _localStrategyParams(),
                metadataURI,
                allowlistEnabled
            )
        );

        return OmniAVS(address(proxy));

        // Below is the code to deploy OmniAVS contract using Create3. It requires vm.getCode(),
        // which seems to be broken in the most recent veresion of forge.
        //
        // Create3 create3 = new Create3();
        // bytes32 create3Salt = keccak256("avs-local-test");
        //
        // address impl =
        //     address(new OmniAVS(IDelegationManager(address(delegation)), IAVSDirectory(address(avsDirectory))));
        //
        // bytes memory bytecode = abi.encodePacked(
        //     vm.getCode("TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy.0.8.12"),
        //     abi.encode(
        //         address(impl),
        //         proxyAdmin,
        //         abi.encodeWithSelector(
        //             OmniAVS.initialize.selector,
        //             omniAVSOwner,
        //             portal,
        //             omniChainId,
        //             ethStakeInbox,
        //             minOperatorStake,
        //             maxOperatorCount,
        //             _localStrategyParams(),
        //             metadataURI,
        //             allowlistEnabled
        //         )
        //     )
        // );
        //
        // address proxy = create3.deploy(create3Salt, bytecode);
        //
        // return OmniAVS(proxy);
    }

    function _localStrategyParams() internal view returns (IOmniAVS.StrategyParam[] memory params) {
        // add all EigenLayerDeployer.strategies
        params = new IOmniAVS.StrategyParam[](strategies.length);

        for (uint256 i = 0; i < strategies.length; i++) {
            params[i] = IOmniAVS.StrategyParam({
                strategy: IStrategy(strategies[i]),
                multiplier: uint96(1e18) // OmniAVS.STRATEGY_WEIGHTING_DIVISOR
             });
        }

        return params;
    }
}
