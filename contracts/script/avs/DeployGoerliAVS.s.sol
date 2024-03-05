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
    uint64 internal _omniChainId = 165;
    address internal _create3Factory = address(uint160(uint256(keccak256("TODO"))));
    address internal _proxyAdmin = address(uint160(uint256(keccak256("TODO"))));
    address internal _owner = address(uint160(uint256(keccak256("TODO"))));
    address internal _portal = address(0);
    address internal _ethStakeInbox = address(0);
    bytes32 internal _create3Salt = keccak256("TODO");

    /// @dev forge script entrypoint
    function run() public {
        uint64 omniChainId = uint64(vm.envUint("OMNI_CHAIN_ID"));
        uint256 deployerKey = vm.envUint("DEPLOYER_KEY");
        address factory = vm.envOr("CREATE3_FACTORY_ADDRESS", _create3Factory);
        address proxyAdmin = vm.envOr("PROXY_ADMIN", _proxyAdmin);
        address owner = vm.envOr("AVS_OWNER", _owner);
        address portal = vm.envOr("PORTAL_ADDRESS", _portal);
        address ethStakeInbox = vm.envOr("ETH_STAKE_INBOX", _ethStakeInbox);
        bytes32 salt = vm.envOr("CREATE3_SALT", _create3Salt);

        vm.startBroadcast(deployerKey);
        deploy(factory, salt, owner, proxyAdmin, portal, omniChainId, ethStakeInbox);
    }

    /// @dev defines goerli deployment logic
    function deploy(
        address create3Factory,
        bytes32 create3Salt,
        address owner,
        address proxyAdmin,
        address portal,
        uint64 omniChainId,
        address ethStakeInbox
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
                    OmniAVS.initialize.selector, owner, portal, omniChainId, ethStakeInbox, StrategyParams.goerli()
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
        address ethStakeInbox
    ) public returns (address) {
        vm.startPrank(prank);
        address avs = deploy(create3Factory, create3Salt, owner, proxyAdmin, portal, omniChainId, ethStakeInbox);
        vm.stopPrank();
        return avs;
    }
}
