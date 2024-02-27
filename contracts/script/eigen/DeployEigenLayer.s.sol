// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import { MockERC20 } from "test/common/MockERC20.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { EigenLayerLocal } from "test/avs/eigen/deploy/EigenLayerLocal.sol";
import { IEigenDeployer } from "test/avs/eigen/deploy/IEigenDeployer.sol";
import { Script } from "forge-std/Script.sol";

contract DeployLocalEigenLayer is Script, EigenLayerLocal {
    function run() external {
        vm.startBroadcast();
        Deployments memory deployments = EigenLayerLocal.deploy();
        _writeDeployments(deployments);
    }

    function _writeDeployments(IEigenDeployer.Deployments memory deps) private {
        string memory defaultOutputDir = "script/eigen/output";
        string memory outputDir = vm.envOr("OUTPUT_DIR", defaultOutputDir);
        string memory outputFile = string.concat(outputDir, "/deployments.json");

        string memory jsonId = "id";

        // seralize all contract addresses in base json
        vm.serializeAddress(jsonId, "proxyAdmin", deps.proxyAdmin);
        vm.serializeAddress(jsonId, "pauserRegistry", deps.pauserRegistry);
        vm.serializeAddress(jsonId, "avsDirectory", deps.avsDirectory);
        vm.serializeAddress(jsonId, "delegationManager", deps.delegationManager);
        vm.serializeAddress(jsonId, "slasher", deps.slasher);
        vm.serializeAddress(jsonId, "strategyManager", deps.strategyManager);
        vm.serializeAddress(jsonId, "eigenPodManager", deps.eigenPodManager);

        // serialize token symbol mapped to strategy address
        string memory strategies = "strategies";
        string memory strategiesJson;
        for (uint256 i = 0; i < deps.strategies.length; i++) {
            IStrategy strat = IStrategy(deps.strategies[i]);
            strategiesJson =
                vm.serializeAddress(strategies, MockERC20(address(strat.underlyingToken())).symbol(), address(strat));
        }

        // join stragies map with base json
        string memory json = vm.serializeString(jsonId, strategies, strategiesJson);

        vm.writeJson(json, outputFile);
    }
}
