// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

// solhint-disable no-console
// solhint-disable var-name-mixedcase
// solhint-disable max-states-count
// solhint-disable state-visibility

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IDelegationManager } from "src/ext/IDelegationManager.sol";

import { OmniAVS } from "src/OmniAVS.sol";
import { EigenM2Deployments } from "./EigenM2Deployments.sol";
import { StrategyParams } from "./StrategyParams.sol";

import { Script } from "forge-std/Script.sol";
import { console } from "forge-std/console.sol";

/**
 * @title DeployAVS
 * @dev A script + utilities to deploy the mainnet OmniAVS contract
 */
contract DeployAVS is Script {
    uint32 _maxOperatorCount = 30;
    uint64 _omniChainId = 166;
    uint96 _minOperatorStake = 1 ether;
    address _proxyAdmin = 0x42A72499eDDB0374ebFba44Fc880F82CCe736614;
    address _deployer = 0x00000072e2740F8a9A4D20Ed05C1832d12498642;
    address _owner = 0xFf89C654846B2E4BC572cEABE77056daf7b299a3;
    address _portal = 0x000000000000000000000000000000000000dEaD;
    address _ethStakeInbox = 0x000000000000000000000000000000000000dEaD;
    string _metadataURI = "https://raw.githubusercontent.com/omni-network/omni/main/static/avs-metadata.json";
    bool _allowlistEnabled = false;

    // set by deploy(...)
    address _impl;
    bytes _implConstructorArgs;
    address _proxy;
    bytes _proxyConstructorArgs;

    function run() public {
        require(block.chainid == 1, "only mainnet deployment");
        require(_proxyAdmin != address(0), "proxyAdmin not set");
        require(_owner != address(0), "owner not set");

        uint256 deployerKey = vm.envUint("AVS_DEPLOYER_KEY");
        require(vm.addr(deployerKey) == _deployer, "wrong deployer key");

        vm.startBroadcast(deployerKey);
        OmniAVS avs =
            deploy(_owner, _proxyAdmin, _portal, _omniChainId, _ethStakeInbox, _minOperatorStake, _maxOperatorCount);
        vm.stopBroadcast();

        console.log("OmniAVS deployed at: ", address(avs));
        console.log("Implementation: ", _impl);
        console.log("Implementation Constructor Args: ");
        console.logBytes(_implConstructorArgs);
        console.log("Proxy: ", _proxy);
        console.log("Proxy Constructor Args: ");
        console.logBytes(_proxyConstructorArgs);
    }

    function deploy(
        address owner,
        address proxyAdmin,
        address portal,
        uint64 omniChainId,
        address ethStakeInbox,
        uint96 minOperatorStake,
        uint32 maxOperatorCount
    ) public returns (OmniAVS) {
        address impl = address(
            new OmniAVS(
                IDelegationManager(EigenM2Deployments.DelegationManager), IAVSDirectory(EigenM2Deployments.AVSDirectory)
            )
        );

        _impl = impl;
        _implConstructorArgs = abi.encode(
            IDelegationManager(EigenM2Deployments.DelegationManager), IAVSDirectory(EigenM2Deployments.AVSDirectory)
        );

        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            impl,
            proxyAdmin,
            abi.encodeWithSelector(
                OmniAVS.initialize.selector,
                owner,
                portal,
                omniChainId,
                ethStakeInbox,
                minOperatorStake,
                maxOperatorCount,
                StrategyParams.mainnet(),
                _metadataURI,
                _allowlistEnabled
            )
        );

        _proxy = address(proxy);
        _proxyConstructorArgs = abi.encode(
            impl,
            proxyAdmin,
            abi.encodeWithSelector(
                OmniAVS.initialize.selector,
                owner,
                portal,
                omniChainId,
                ethStakeInbox,
                minOperatorStake,
                maxOperatorCount,
                StrategyParams.mainnet(),
                _metadataURI,
                _allowlistEnabled
            )
        );

        return OmniAVS(address(proxy));
    }
}
