// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { Bytes32AddressLib } from "solmate/src/utils/Bytes32AddressLib.sol";

import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";

import { IDelegationManager } from "src/interfaces/IDelegationManager.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { OmniAVS } from "src/protocol/OmniAVS.sol";

import { Create3 } from "src/deploy/Create3.sol";
import { DeployGoerliAVS } from "script/avs/DeployGoerliAVS.s.sol";
import { StrategyParams } from "script/avs/StrategyParams.sol";
import { MockOmniPredeploys } from "test/utils/MockOmniPredeploys.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { EigenLayerFixtures } from "./eigen/EigenLayerFixtures.sol";
import { Empty } from "./Empty.sol";

/**
 * @title Fixtures
 * @dev Common fixtures contract for all AVS tests.
 */
contract Fixtures is EigenLayerFixtures {
    using Bytes32AddressLib for bytes32;

    address multisig = makeAddr("multisig");
    address proxyAdminOwner = multisig;
    address omniAVSOwner = multisig;

    uint32 maxOperatorCount = 10;
    uint96 minOperatorStake = 1 ether;
    uint64 omniChainId = 111;
    address ethStakeInbox = MockOmniPredeploys.ETH_STAKE_INBOX;

    ProxyAdmin proxyAdmin;
    MockPortal portal;
    OmniAVS omniAVS;

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

        omniAVS = isGoerli() ? _deployGoerliAVS() : _deployLocalAVS();
    }

    function _deployGoerliAVS() internal returns (OmniAVS) {
        DeployGoerliAVS deployer = new DeployGoerliAVS();

        Create3 create3 = new Create3();
        bytes32 create3Salt = keccak256("avs-goerli-fork-test");

        return OmniAVS(
            deployer.prankDeploy(
                proxyAdminOwner,
                address(create3),
                create3Salt,
                omniAVSOwner,
                address(proxyAdmin),
                address(portal),
                omniChainId,
                ethStakeInbox,
                minOperatorStake,
                maxOperatorCount
            )
        );
    }

    function _deployLocalAVS() internal returns (OmniAVS) {
        Create3 create3 = new Create3();
        bytes32 create3Salt = keccak256("avs-local-test");

        address impl =
            address(new OmniAVS(IDelegationManager(address(delegation)), IAVSDirectory(address(avsDirectory))));

        bytes memory bytecode = abi.encodePacked(
            vm.getCode("TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy.0.8.12"),
            abi.encode(
                address(impl),
                proxyAdmin,
                abi.encodeWithSelector(
                    OmniAVS.initialize.selector,
                    omniAVSOwner,
                    portal,
                    omniChainId,
                    ethStakeInbox,
                    minOperatorStake,
                    maxOperatorCount,
                    _localStrategyParams()
                )
            )
        );

        address proxy = create3.deploy(create3Salt, bytecode);

        return OmniAVS(proxy);
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
