// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import { Script } from "forge-std/Script.sol";
import { EigenLayerDeployer } from "test/avs/eigen/EigenLayerDeployer.t.sol";

contract DeployEigenLayer is EigenLayerDeployer, Script {
    function run() external {
        vm.startBroadcast();
        _deployEigenLayerContractsLocal();
        _writeDeploymentsJson();
    }

    function _writeDeploymentsJson() private {
        string memory defaultOutputDir = "script/eigen/output";
        string memory outputDir = vm.envOr("OUTPUT_DIR", defaultOutputDir);
        string memory outputFile = string.concat(outputDir, "/deployments.json");

        string memory jsonId = "id";
        vm.serializeAddress(jsonId, "ProxyAdmin", address(eigenLayerProxyAdmin));
        vm.serializeAddress(jsonId, "PauserRegistry", address(eigenLayerPauserReg));
        vm.serializeAddress(jsonId, "AVSDirectory", address(avsDirectory));
        vm.serializeAddress(jsonId, "DelegationManager", address(delegation));
        vm.serializeAddress(jsonId, "Slasher", address(slasher));
        vm.serializeAddress(jsonId, "StrategyManager", address(strategyManager));
        vm.serializeAddress(jsonId, "EigenPodManager", address(eigenPodManager));
        vm.serializeAddress(jsonId, "EigenPod", address(pod));
        vm.serializeAddress(jsonId, "DelayedWithdrawalRouter", address(delayedWithdrawalRouter));
        vm.serializeAddress(jsonId, "ETHPOSDeposit", address(ethPOSDeposit));
        vm.serializeAddress(jsonId, "EigenPodBeacon", address(eigenPodBeacon));
        vm.serializeAddress(jsonId, "EigenToken", address(eigenToken));
        vm.serializeAddress(jsonId, "EigenStrategy", address(eigenStrat));
        vm.serializeAddress(jsonId, "WETH", address(weth));
        string memory json = vm.serializeAddress(jsonId, "WETHStrategy", address(wethStrat));
        vm.writeJson(json, outputFile);
    }
}
