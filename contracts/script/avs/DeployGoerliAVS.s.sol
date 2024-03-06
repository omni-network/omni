// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

// solhint-disable no-console

import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IDelegationManager } from "src/interfaces/IDelegationManager.sol";

import { Create3 } from "src/deploy/Create3.sol";
import { OmniAVS } from "src/protocol/OmniAVS.sol";
import { EigenM2GoerliDeployments } from "test/avs/common/eigen/EigenM2GoerliDeployments.sol";
import { StrategyParams } from "./StrategyParams.sol";

import { Script } from "forge-std/Script.sol";
import { console } from "forge-std/console.sol";

/**
 * @title DeployGoerliAVS
 * @dev A script + utilites for deploying OmnIAVS to Goerli. It exposes a
 *      deploy function, so that fork tests can use the same deployment logic as the
 *      deploy script.
 */
contract DeployGoerliAVS is Script {
    uint32 internal _maxOperatorCount = 200;
    uint64 internal _omniChainId = 165;
    uint96 internal _minOperatorStake = 1 ether;
    address internal _create3Factory = 0x30C9242B7b5Ec92500a2752b651BEDE1a42dc12D;
    address internal _proxyAdmin = 0x6Ecc7b4588fa09d7aB468fE5D778c97dF28c021E;
    address internal _deployer = 0x28D1a0B2C6E11fAf43582054694587c762b03A38;
    address internal _owner = 0x46D005a3D18740Cd63e8072078796f043d040191;
    address internal _portal = address(0);
    address internal _ethStakeInbox = address(0);
    bytes32 internal _create3Salt = keccak256("omni-avs");
    string internal _metadataURI = "https://raw.githubusercontent.com/omni-network/omni/main/static/avs-metadata.json";

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

        console.log("OmniAVS deployed at: ", address(avs));
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

        address deployed = create3.getDeployed(address(this), create3Salt);
        require(deployed.code.length == 0, "already deployed");

        address impl = address(
            new OmniAVS(
                IDelegationManager(EigenM2GoerliDeployments.DelegationManager),
                IAVSDirectory(EigenM2GoerliDeployments.AVSDirectory)
            )
        );

        bytes memory bytecode = abi.encodePacked(
            vm.getCode("TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy.0.8.12"),
            abi.encode(
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
            )
        );

        return OmniAVS(create3.deploy(create3Salt, bytecode));
    }

    /// @dev deploy OmniAVS, but with a prank. necessary because we cannot
    //       vm.startPrank() outside of this contract
    function prankDeploy(
        address prank,
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
        vm.startPrank(prank);
        OmniAVS avs = deploy(
            create3Factory,
            create3Salt,
            owner,
            proxyAdmin,
            portal,
            omniChainId,
            ethStakeInbox,
            minOperatorStake,
            maxOperatorCount
        );
        vm.stopPrank();
        return avs;
    }
}
