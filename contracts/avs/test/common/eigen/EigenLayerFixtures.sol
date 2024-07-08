// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ERC20PresetFixedSupply } from "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import { IStrategyManager } from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { StrategyManager } from "eigenlayer-contracts/src/contracts/core/StrategyManager.sol";
import { DelegationManager } from "eigenlayer-contracts/src/contracts/core/DelegationManager.sol";
import { AVSDirectory } from "eigenlayer-contracts/src/contracts/core/AVSDirectory.sol";
import { StrategyBase } from "eigenlayer-contracts/src/contracts/strategies/StrategyBase.sol";

import { EigenPodManagerHarness } from "./EigenPodManagerHarness.sol";
import { EigenLayerHolesky } from "./EigenLayerHolesky.sol";
import { EigenLayerMainnet } from "./EigenLayerMainnet.sol";
import { EigenLayerLocal } from "./EigenLayerLocal.sol";
import { IEigenDeployer } from "./IEigenDeployer.sol";

import { Test } from "forge-std/Test.sol";
import { MockERC20 } from "test/common/MockERC20.sol";
import { console } from "forge-std/console.sol";

/**
 * @title EigenLayerFixtures
 * @dev Deploys eigen layer test fixtures. Deploys local or goerli contracts, depending on the chain id.
 *      It also deploys an "unsupported strategy" that is always excluded from OmniAVS strategy params.
 */
contract EigenLayerFixtures is Test {
    // eigen deployments
    DelegationManager delegation;
    AVSDirectory avsDirectory;
    StrategyManager strategyManager;
    EigenPodManagerHarness eigenPodManager;

    // stragies
    address[] strategies;

    // unsupported strategy (always excluded from OmniAVS strategy params)
    StrategyBase unsupportedStrat;

    function isHolesky() public view returns (bool) {
        return block.chainid == 17_000;
    }

    function isMainnet() public view returns (bool) {
        return block.chainid == 1;
    }

    function setUp() public virtual {
        IEigenDeployer deployer;

        if (isMainnet()) {
            deployer = new EigenLayerMainnet();
        } else if (isHolesky()) {
            deployer = new EigenLayerHolesky();
        } else {
            deployer = new EigenLayerLocal();
        }

        IEigenDeployer.Deployments memory deployments = deployer.deploy();

        // we always deploy unsupported strategy
        unsupportedStrat = StrategyBase(
            _deployUnsupportedStrategy(deployments.strategyManager, deployments.proxyAdmin, deployments.pauserRegistry)
        );

        delegation = DelegationManager(deployments.delegationManager);
        avsDirectory = AVSDirectory(deployments.avsDirectory);
        strategyManager = StrategyManager(deployments.strategyManager);
        eigenPodManager = EigenPodManagerHarness(deployments.eigenPodManager);

        for (uint256 i = 0; i < deployments.strategies.length; i++) {
            strategies.push(deployments.strategies[i]);
        }

        _whitelistStrategies(strategies);
    }

    function _deployUnsupportedStrategy(address strategyManager_, address proxyAdmin_, address pauserRegistry_)
        internal
        returns (address)
    {
        IERC20 unsupportedToken = new MockERC20("unsupported", "UNSUPPORTED");
        StrategyBase impl = new StrategyBase(IStrategyManager(strategyManager_));
        return address(
            new TransparentUpgradeableProxy(
                address(impl),
                address(proxyAdmin_),
                abi.encodeWithSelector(StrategyBase.initialize.selector, unsupportedToken, pauserRegistry_)
            )
        );
    }

    function _whitelistStrategies(address[] memory strats) internal {
        vm.startPrank(strategyManager.strategyWhitelister());

        IStrategy[] memory _strategy = new IStrategy[](strats.length);
        bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](strats.length);

        for (uint256 i = 0; i < strats.length; i++) {
            _strategy[i] = IStrategy(strats[i]);
            _thirdPartyTransfersForbiddenValues[i] = false;
        }

        strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);

        vm.stopPrank();
    }
}
