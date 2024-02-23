// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IDelegationManager } from "src/interfaces/IDelegationManager.sol";

import { Empty } from "test/common/Empty.sol";
import { OmniAVS } from "src/protocol/OmniAVS.sol";
import { EigenM2GoerliDeployments } from "test/avs/eigen/deploy/EigenM2GoerliDeployments.sol";
import { StrategyParams } from "./StrategyParams.sol";

import { Script } from "forge-std/Script.sol";

/**
 * @title DeployGoerliAVS
 * @dev A script + utilites for deploying OmnIAVS to Goerli. It exposes a
 *      deploy function, so that fork tests can use the same deployment logic as the
 *      deploy script.
 */
contract DeployGoerliAVS is Script {
    uint96 public constant MIN_OPERATOR_STAKE = 1 ether;
    uint32 public constant MAX_OPERATOR_COUNT = 10;

    /// @dev forge script entrypoint
    function run() public pure {
        revert("Not implemented");
    }

    /// @dev defines goerli deployment logic
    function deploy(address owner, address proxyAdmin, address portal, uint64 omniChainId) public returns (address) {
        address proxy = address(new TransparentUpgradeableProxy(address(new Empty()), proxyAdmin, ""));
        address impl = address(
            new OmniAVS(
                IDelegationManager(EigenM2GoerliDeployments.DelegationManager),
                IAVSDirectory(EigenM2GoerliDeployments.AVSDirectory)
            )
        );

        address[] memory allowlist = new address[](0);

        ProxyAdmin(proxyAdmin).upgradeAndCall(
            ITransparentUpgradeableProxy(proxy),
            impl,
            abi.encodeWithSelector(
                OmniAVS.initialize.selector,
                owner,
                portal,
                omniChainId,
                MIN_OPERATOR_STAKE,
                MAX_OPERATOR_COUNT,
                allowlist,
                StrategyParams.goerli()
            )
        );

        return proxy;
    }

    /// @dev deploy OmniAVS, but with a prank. necessary because we cannot
    //       vm.startPrank() outside of this contract
    function prankDeploy(address prank, address owner, address proxyAdmin, address portal, uint64 omniChainId)
        public
        returns (address)
    {
        vm.startPrank(prank);
        address avs = deploy(owner, proxyAdmin, portal, omniChainId);
        vm.stopPrank();
        return avs;
    }
}
