// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

// solhint-disable no-console
// solhint-disable var-name-mixedcase
// solhint-disable max-states-count
// solhint-disable state-visibility

import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IDelegationManager } from "src/interfaces/IDelegationManager.sol";

import { Create3 } from "src/deploy/Create3.sol";
import { OmniAVS } from "src/protocol/OmniAVS.sol";
import { EigenM2GoerliDeployments } from "test/avs/common/eigen/EigenM2GoerliDeployments.sol";
import { StrategyParams } from "./StrategyParams.sol";

import { Script } from "forge-std/Script.sol";

/**
 * @title DeployGoerliAVS
 * @dev A script + utilites for deploying OmnIAVS to Goerli. It exposes a
 *      deploy function, so that fork tests can use the same deployment logic as the
 *      deploy script.
 */
contract DeployGoerliAVS is Script {
    uint32 _maxOperatorCount = 200;
    uint64 _omniChainId = 165;
    uint96 _minOperatorStake = 1 ether;
    address _create3Factory = 0x30C9242B7b5Ec92500a2752b651BEDE1a42dc12D;
    address _proxyAdmin = 0x6Ecc7b4588fa09d7aB468fE5D778c97dF28c021E;
    address _deployer = 0x28D1a0B2C6E11fAf43582054694587c762b03A38;
    address _owner = 0x46D005a3D18740Cd63e8072078796f043d040191;
    address _portal = address(0);
    address _ethStakeInbox = address(0);
    bytes32 _create3Salt = keccak256("omni-avs");
    string _metadataURI = "https://raw.githubusercontent.com/omni-network/omni/main/static/avs-metadata.json";

    // set by deploy(...)
    address _impl;
    bytes _implConstructorArgs;
    address _proxy;
    bytes _proxyConstructorArgs;

    // where to write the output
    string _outputFile = "script/avs/output/deploy-goerli-avs.json";

    /// @dev forge script entrypoint
    function run() public {
        require(block.chainid == 5, "Don't deploy on this chain yet!");
        require(_create3Factory != address(0), "create3Factory not set");
        require(_proxyAdmin != address(0), "proxyAdmin not set");
        require(_owner != address(0), "owner not set");

        uint256 deployerKey = vm.envUint("AVS_DEPLOYER_KEY");
        uint256 ownerKey = vm.envUint("AVS_OWNER_KEY");

        require(vm.addr(deployerKey) == _deployer, "wrong deployer key");
        require(vm.addr(ownerKey) == _owner, "wrong owner key");

        vm.startBroadcast(deployerKey);
        OmniAVS avs = deploy(
            _create3Factory,
            _create3Salt,
            _owner,
            _proxyAdmin,
            _portal,
            _omniChainId,
            _ethStakeInbox,
            _minOperatorStake,
            _maxOperatorCount
        );

        vm.stopBroadcast();

        vm.startBroadcast(ownerKey);
        avs.disableAllowlist();
        avs.setMetadataURI(_metadataURI);
        vm.stopBroadcast();

        _writeOutput();
    }

    /// @dev defines goerli deployment logic
    function deploy(
        address create3Factory,
        bytes32 create3Salt,
        address owner,
        address proxyAdmin,
        address portal,
        uint64 omniChainId,
        address ethStakeInbox,
        uint96 minOperatorStake,
        uint32 maxOperatorCount
    ) public returns (OmniAVS) {
        Create3 create3 = Create3(create3Factory);

        address impl = address(
            new OmniAVS(
                IDelegationManager(EigenM2GoerliDeployments.DelegationManager),
                IAVSDirectory(EigenM2GoerliDeployments.AVSDirectory)
            )
        );

        _impl = impl;
        _implConstructorArgs = abi.encode(
            IDelegationManager(EigenM2GoerliDeployments.DelegationManager),
            IAVSDirectory(EigenM2GoerliDeployments.AVSDirectory)
        );

        bytes memory proxyConstructorArgs = abi.encode(
            address(impl),
            proxyAdmin,
            abi.encodeWithSelector(
                OmniAVS.initialize.selector,
                owner,
                portal,
                omniChainId,
                ethStakeInbox,
                minOperatorStake,
                maxOperatorCount,
                StrategyParams.goerli()
            )
        );

        bytes memory bytecode = abi.encodePacked(
            vm.getCode("TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy.0.8.12"), proxyConstructorArgs
        );

        address proxy = create3.deploy(create3Salt, bytecode);

        _proxy = proxy;
        _proxyConstructorArgs = proxyConstructorArgs;

        return OmniAVS(proxy);
    }

    function _writeOutput() internal {
        string memory root = "root";
        vm.serializeAddress(root, "deployer", _deployer);
        vm.serializeAddress(root, "proxyAdmin", _proxyAdmin);
        vm.serializeAddress(root, "owner", _owner);
        vm.serializeAddress(root, "create3Factory", _create3Factory);
        vm.serializeBytes32(root, "create3Salt", _create3Salt);

        string memory proxy = "proxy";
        vm.serializeAddress(proxy, "address", _proxy);
        string memory proxyJson = vm.serializeBytes(proxy, "constructorArgs", _proxyConstructorArgs);

        string memory impl = "impl";
        vm.serializeAddress(impl, "address", _impl);
        string memory implJson = vm.serializeBytes(impl, "constructorArgs", _implConstructorArgs);

        string memory contracts = "contracts";
        vm.serializeString(contracts, "TransparentUpgradeableProxy", proxyJson);
        string memory contractsJson = vm.serializeString(contracts, "OmniAVS", implJson);

        string memory finalJson = vm.serializeString(root, "contracts", contractsJson);
        vm.writeFile(_outputFile, finalJson);
    }
}
