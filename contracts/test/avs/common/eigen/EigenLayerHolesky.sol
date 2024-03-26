// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import { ERC20, IERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { IBeacon } from "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { IETHPOSDeposit } from "eigenlayer-contracts/src/contracts/interfaces/IETHPOSDeposit.sol";
import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import { IEigenPodManager } from "eigenlayer-contracts/src/contracts/interfaces/IEigenPodManager.sol";
import { IStrategyManager } from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";
import { ISlasher } from "eigenlayer-contracts/src/contracts/interfaces/ISlasher.sol";

import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { StrategyParams } from "../StrategyParams.sol";
import { IEigenDeployer } from "./IEigenDeployer.sol";
import { EigenM2HoleskyDeployments } from "./EigenM2HoleskyDeployments.sol";
import { EigenPodManagerHarness } from "./EigenPodManagerHarness.sol";

import { Test } from "forge-std/Test.sol";

/**
 * @title EigenLayerHolesky
 * @dev A holesky IEigenDeployer. This contract is used when "deploying"
 *      EigenLayer on a holesky fork. It does not actually deploy anything, it just
 *      returns the addresses of the contracts that are already deployed on holesky.
 */
contract EigenLayerHolesky is IEigenDeployer, Test {
    address beaconEthStrategy = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0;

    function deploy() public returns (Deployments memory deps) {
        address proxyAdminAddr = _proxyAdmin(EigenM2HoleskyDeployments.EigenPodManager);
        address proxyAdminOwner = ProxyAdmin(proxyAdminAddr).owner();

        IOmniAVS.StrategyParam[] memory stratParams = StrategyParams.holesky();

        address[] memory strategies = new address[](stratParams.length);
        for (uint256 i = 0; i < stratParams.length; i++) {
            address strat = address(stratParams[i].strategy);
            strategies[i] = strat;

            if (strat == beaconEthStrategy) continue;

            IERC20 underlying = IStrategy(strat).underlyingToken();
            _replaceERC20(strat, address(underlying));
        }

        deps = Deployments({
            proxyAdminOwner: proxyAdminOwner,
            proxyAdmin: proxyAdminAddr,
            pauserRegistry: EigenM2HoleskyDeployments.PauserRegistry,
            delegationManager: EigenM2HoleskyDeployments.DelegationManager,
            eigenPodManager: EigenM2HoleskyDeployments.EigenPodManager,
            strategyManager: EigenM2HoleskyDeployments.StrategyManager,
            slasher: EigenM2HoleskyDeployments.Slasher,
            avsDirectory: EigenM2HoleskyDeployments.AVSDirectory,
            strategies: strategies
        });

        _replaceEigenPodManager(deps);
    }

    /// @dev Storage slot with the admin of the contract.
    /// This is the keccak-256 hash of "eip1967.proxy.admin" subtracted by 1, and is
    bytes32 internal constant PROXY_ADMIN_SLOT = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;

    function _proxyAdmin(address proxy) public view virtual returns (address) {
        bytes32 adminSlot = vm.load(proxy, PROXY_ADMIN_SLOT);
        return address(uint160(uint256(adminSlot)));
    }

    /// @dev Replace code at `token` with a basic ERC20 implementation.
    ///      Some underlying strategy tokens are not basic ERC20s (like stETH), and deal() does not work with them.
    ///      Replacting them with a basic ERC20 allows deal() balances to test operators & delegators.
    function _replaceERC20(address strategy, address token) internal {
        string memory name = ERC20(token).name();
        string memory symbol = ERC20(token).symbol();

        uint256 stratBalance = ERC20(token).balanceOf(strategy);

        ERC20 underlying = new ERC20(name, symbol);
        vm.etch(token, address(underlying).code);

        deal(address(token), strategy, stratBalance);
        assertEq(ERC20(token).balanceOf(strategy), stratBalance);
    }

    /// @dev Replace the EigenPodManager with our harness that allows to updatePodOwnerShares.
    function _replaceEigenPodManager(Deployments memory deps) internal {
        IEigenPodManager current = IEigenPodManager(deps.eigenPodManager);
        IETHPOSDeposit ethPOS = current.ethPOS();
        IBeacon eigenPodBeacon = current.eigenPodBeacon();

        vm.prank(deps.proxyAdmin);
        address impl = ITransparentUpgradeableProxy(address(current)).implementation();

        EigenPodManagerHarness harness = new EigenPodManagerHarness(
            ethPOS,
            eigenPodBeacon,
            IStrategyManager(deps.strategyManager),
            ISlasher(deps.slasher),
            IDelegationManager(deps.delegationManager)
        );

        vm.etch(impl, address(harness).code);
    }
}
