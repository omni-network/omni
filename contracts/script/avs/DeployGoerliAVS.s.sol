// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

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
    uint32 internal _maxOperatorCount = 200;
    uint64 internal _omniChainId = 165;
    uint96 internal _minOperatorStake = 1 ether;
    address internal _create3Factory = address(0); // TODO
    address internal _proxyAdmin = address(0); // TODO
    address internal _owner = address(0); // TODO
    address internal _portal = address(0);
    address internal _ethStakeInbox = address(0);
    bytes32 internal _create3Salt = keccak256("omni-avs");

    /// @dev forge script entrypoint
    function run() public {
        require(_create3Factory != address(0), "create3Factory not set");

        uint256 deployerKey = vm.envUint("DEPLOYER_KEY");
        vm.startBroadcast(deployerKey);

        deploy(
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
    ) public returns (address) {
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

        address deployed = Create3(create3Factory).getDeployed(address(this), create3Salt);
        require(deployed.code.length == 0, "already deployed");

        return Create3(create3Factory).deploy(create3Salt, bytecode);
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
    ) public returns (address) {
        vm.startPrank(prank);
        address avs = deploy(
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
