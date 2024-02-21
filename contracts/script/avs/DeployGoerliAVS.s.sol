// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { EigenM2GoerliDeployments } from "test/avs/eigen/EigenM2GoerliDeployments.sol";
import { AVSDeploy } from "./AVSDeploy.sol";
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
        address proxy = AVSDeploy.proxy(proxyAdmin);
        address impl = AVSDeploy.impl(EigenM2GoerliDeployments.DelegationManager, EigenM2GoerliDeployments.AVSDirectory);

        address[] memory allowlist = new address[](0);

        AVSDeploy.upgradeAndInit(
            proxyAdmin,
            proxy,
            impl,
            owner,
            portal,
            omniChainId,
            minimumOperatorStake,
            maxOperatorCount,
            allowlist,
            StrategyParams.goerli()
        );

        return proxy;
    }
}
