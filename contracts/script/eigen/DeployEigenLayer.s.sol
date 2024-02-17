// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import { Script } from "forge-std/Script.sol";
import { EigenLayerDeployer } from "test/avs/eigen/EigenLayerDeployer.t.sol";

contract DeployEigenLayer is EigenLayerDeployer, Script {
    function run() external {
        vm.startBroadcast();
        _deployEigenLayerContractsLocal(false /*mockAvsDirectory*/ );
        _writeDeploymentsJson();
    }

    function _writeDeploymentsJson() private {
        string memory defaultOutputDir = "script/eigen/output";
        string memory outputDir = vm.envOr("OUTPUT_DIR", defaultOutputDir);
        string memory outputFile = string.concat(outputDir, "/deployments.json");

        string memory jsonOId = "id";
        vm.serializeAddress(jsonOId, "ProxyAdmin", address(eigenLayerProxyAdmin));
        vm.serializeAddress(jsonOId, "PauserRegistry", address(eigenLayerPauserReg));
        vm.serializeAddress(jsonOId, "AVSDirectory", address(avsDirectory));
        vm.serializeAddress(jsonOId, "DelegationManager", address(delegation));
        vm.serializeAddress(jsonOId, "Slasher", address(slasher));
        vm.serializeAddress(jsonOId, "StrategyManager", address(strategyManager));
        vm.serializeAddress(jsonOId, "EigenPodManager", address(eigenPodManager));
        vm.serializeAddress(jsonOId, "EigenPod", address(pod));
        vm.serializeAddress(jsonOId, "DelayedWithdrawalRouter", address(delayedWithdrawalRouter));
        vm.serializeAddress(jsonOId, "ETHPOSDeposit", address(ethPOSDeposit));
        vm.serializeAddress(jsonOId, "EigenPodBeacon", address(eigenPodBeacon));
        vm.serializeAddress(jsonOId, "EigenToken", address(eigenToken));
        vm.serializeAddress(jsonOId, "EigenStrategy", address(eigenStrat));
        vm.serializeAddress(jsonOId, "WETH", address(weth));
        string memory json = vm.serializeAddress(jsonOId, "WETHStrategy", address(wethStrat));
        vm.writeJson(json, outputFile);
    }
}
