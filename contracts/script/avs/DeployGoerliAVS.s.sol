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

// solhint-disable const-name-snakecase

contract DeployGoerliAVS is Script {
    uint96 public constant minimumOperatorStake = 1 ether;
    uint32 public constant maxOperatorCount = 10;

    function run() public pure {
        revert("Not implemented");
    }

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
                minimumOperatorStake,
                maxOperatorCount,
                allowlist,
                StrategyParams.goerli()
            )
        );

        return proxy;
    }
}
