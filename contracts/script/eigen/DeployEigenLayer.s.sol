// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import { Script } from "forge-std/Script.sol";
import { EigenLayerLocal } from "test/avs/eigen/deploy/EigenLayerLocal.sol";
import { IEigenDeployer } from "test/avs/eigen/deploy/IEigenDeployer.sol";

contract DeployLocalEigenLayer is Script {
    function run() external {
        vm.startBroadcast();
        IEigenDeployer deployer = new EigenLayerLocal();
        IEigenDeployer.Deployments memory deployments = deployer.deploy();
        _writeDeployments(deployments);
    }

    function _writeDeployments(IEigenDeployer.Deployments memory deps) private {
        string memory defaultOutputDir = "script/eigen/output";
        string memory outputDir = vm.envOr("OUTPUT_DIR", defaultOutputDir);
        string memory outputFile = string.concat(outputDir, "/deployments.json");

        string memory jsonId = "id";
        vm.serializeAddress(jsonId, "ProxyAdmin", deps.proxyAdmin);
        vm.serializeAddress(jsonId, "PauserRegistry", deps.pauserRegistry);
        vm.serializeAddress(jsonId, "AVSDirectory", deps.avsDirectory);
        vm.serializeAddress(jsonId, "DelegationManager", deps.delegationManager);
        vm.serializeAddress(jsonId, "Slasher", deps.slasher);
        vm.serializeAddress(jsonId, "StrategyManager", deps.strategyManager);
        string memory json = vm.serializeAddress(jsonId, "EigenPodManager", deps.eigenPodManager);
        // vm.serializeAddress(jsonId, "UnsupportedToken", address(unsupportedToken));
        // vm.serializeAddress(jsonId, "UnsupportedStrategy", address(unsupportedStrat));
        // vm.serializeAddress(jsonId, "WETH", address(weth));
        // string memory json = vm.serializeAddress(jsonId, "WETHStrategy", address(wethStrat));
        vm.writeJson(json, outputFile);
    }
}
