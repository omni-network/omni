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

import { IEigenDeployer } from "./IEigenDeployer.sol";
import { EigenM2GoerliDeployments } from "./EigenM2GoerliDeployments.sol";
import { EigenPodManagerHarness } from "../EigenPodManagerHarness.sol";

import { Test } from "forge-std/Test.sol";

/**
 * @title EigenLayerGoerli
 * @dev A goerli IEigenDeployer. This contract is used when "deploying"
 *      EigenLayer on a goerli fork. It does not actually deploy anything, it just
 *      returns the addresses of the contracts that are already deployed on goerli.
 */
contract EigenLayerGoerli is IEigenDeployer, Test {
    function deploy() public returns (Deployments memory deps) {
        address proxyAdminAddr = _proxyAdmin(EigenM2GoerliDeployments.EigenPodManager);
        address proxyAdminOwner = ProxyAdmin(proxyAdminAddr).owner();

        address[] memory strategies = new address[](2);
        strategies[0] = EigenM2GoerliDeployments.stETHStrategy;
        strategies[1] = EigenM2GoerliDeployments.rETHStrategy;

        IERC20 stETH = IStrategy(EigenM2GoerliDeployments.stETHStrategy).underlyingToken();
        IERC20 rETH = IStrategy(EigenM2GoerliDeployments.rETHStrategy).underlyingToken();

        _replaceERC20(EigenM2GoerliDeployments.stETHStrategy, address(stETH));
        _replaceERC20(EigenM2GoerliDeployments.rETHStrategy, address(rETH));

        deps = Deployments({
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
