// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";

import { IEigenDeployer } from "./IEigenDeployer.sol";
import { EigenM2GoerliDeployments } from "./EigenM2GoerliDeployments.sol";

contract EigenLayerGoerli is IEigenDeployer {
    function deploy() external view returns (Deployments memory) {
        address proxyAdminAddr = _proxyAdmin(EigenM2GoerliDeployments.EigenPodManager);
        address proxyAdminOwner = ProxyAdmin(proxyAdminAddr).owner();

        address[] memory strategies = new address[](2);
        strategies[0] = EigenM2GoerliDeployments.stETHStrategy;
        strategies[1] = EigenM2GoerliDeployments.rETHStrategy;

        return Deployments({
            proxyAdminOwner: proxyAdminOwner,
            proxyAdmin: proxyAdminAddr,
            pauserRegistry: EigenM2GoerliDeployments.PauserRegistry,
            delegationManager: EigenM2GoerliDeployments.DelegationManager,
            eigenPodManager: EigenM2GoerliDeployments.EigenPodManager,
            strategyManager: EigenM2GoerliDeployments.StrategyManager,
            slasher: EigenM2GoerliDeployments.Slasher,
            avsDirectory: EigenM2GoerliDeployments.AVSDirectory,
            strategies: strategies
        });
    }

    function _proxyAdmin(address proxy) public view virtual returns (address) {
        // We need to manually run the static call since the getter cannot be flagged as view
        // bytes4(keccak256("admin()")) == 0xf851a440
        (bool success, bytes memory returndata) = address(proxy).staticcall(hex"f851a440");
        require(success);
        return abi.decode(returndata, (address));
    }
}
