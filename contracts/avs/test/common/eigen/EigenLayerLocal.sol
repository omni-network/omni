// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import { ERC20PresetFixedSupply } from "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { IBeacon } from "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import { UpgradeableBeacon } from "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import { IETHPOSDeposit } from "eigenlayer-contracts/src/contracts/interfaces/IETHPOSDeposit.sol";
import { IBeaconChainOracle } from "eigenlayer-contracts/src/contracts/interfaces/IBeaconChainOracle.sol";

import { DelegationManager } from "eigenlayer-contracts/src/contracts/core/DelegationManager.sol";
import { AVSDirectory } from "eigenlayer-contracts/src/contracts/core/AVSDirectory.sol";
import { StrategyManager } from "eigenlayer-contracts/src/contracts/core/StrategyManager.sol";
import { StrategyBase } from "eigenlayer-contracts/src/contracts/strategies/StrategyBase.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { Slasher } from "eigenlayer-contracts/src/contracts/core/Slasher.sol";
import { EigenPodManager } from "eigenlayer-contracts/src/contracts/pods/EigenPodManager.sol";
import { EigenPod, IEigenPod } from "eigenlayer-contracts/src/contracts/pods/EigenPod.sol";
import { DelayedWithdrawalRouter } from "eigenlayer-contracts/src/contracts/pods/DelayedWithdrawalRouter.sol";
import { IDelayedWithdrawalRouter } from "eigenlayer-contracts/src/contracts/pods/DelayedWithdrawalRouter.sol";
import { PauserRegistry } from "eigenlayer-contracts/src/contracts/permissions/PauserRegistry.sol";

import { EmptyContract } from "eigenlayer-contracts/src/test/mocks/EmptyContract.sol";
import { ETHPOSDepositMock } from "eigenlayer-contracts/src/test/mocks/ETHDepositMock.sol";

import { MockERC20 } from "test/common/MockERC20.sol";
import { EigenPodManagerHarness } from "./EigenPodManagerHarness.sol";
import { IEigenDeployer } from "./IEigenDeployer.sol";

import { CommonBase } from "forge-std/Base.sol";

/**
 * @title EigenLayerLocal
 * @dev A local IEigenDeployer. This contract is used when not running against
 *      a testnet or mainnet fork. It requires deploys all eigenlayer core
 *      contracts.
 */
contract EigenLayerLocal is IEigenDeployer, CommonBase {
    // EigenLayer contracts
    ProxyAdmin proxyAdmin;
    PauserRegistry pauserReg;
    Slasher slasher;
    DelegationManager delegation;
    AVSDirectory avsDirectory;
    StrategyManager strategyManager;
    EigenPodManager eigenPodManager;
    IEigenPod pod;
    IDelayedWithdrawalRouter delayedWithdrawalRouter;
    IETHPOSDeposit ethPOSDeposit;
    IBeacon eigenPodBeacon;

    // config
    IStrategy[] initializeStrategiesToSetDelayBlocks;
    uint256[] initializeWithdrawalDelayBlocks;
    uint256 minWithdrawalDelayBlocks = 0;
    uint32 PARTIAL_WITHDRAWAL_FRAUD_PROOF_PERIOD_BLOCKS = 7 days / 12 seconds;
    uint64 MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = 32e9;
    uint64 GOERLI_GENESIS_TIME = 1_616_508_000;

    // addrs
    address pauser;
    address unpauser;
    address beaconChainOracleAddress;

    // eignlayer multisig in prod, anvil account9 locally
    address constant proxyAdminOwner = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;

    /// Canonical, virtual beacon chain ETH strategy
    address constant beaconChainETHStrategy = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0;

    function deploy() public returns (Deployments memory) {
        pauser = address(69);
        unpauser = address(489);
        // deploy proxy admin for ability to upgrade proxy contracts
        proxyAdmin = new ProxyAdmin();

        //deploy pauser registry
        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserReg = new PauserRegistry(pausers, unpauser);

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        EmptyContract emptyContract = new EmptyContract();
        delegation =
            DelegationManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));
        avsDirectory =
            AVSDirectory(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));
        strategyManager =
            StrategyManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));
        slasher = Slasher(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));
        eigenPodManager =
            EigenPodManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));
        delayedWithdrawalRouter = DelayedWithdrawalRouter(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );

        ethPOSDeposit = new ETHPOSDepositMock();
        pod = new EigenPod(
            ethPOSDeposit,
            delayedWithdrawalRouter,
            eigenPodManager,
            MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            GOERLI_GENESIS_TIME
        );

        eigenPodBeacon = new UpgradeableBeacon(address(pod));

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        DelegationManager delegationImplementation = new DelegationManager(strategyManager, slasher, eigenPodManager);
        AVSDirectory avsDirectoryImplementation = new AVSDirectory(delegation);
        StrategyManager strategyManagerImplementation = new StrategyManager(delegation, eigenPodManager, slasher);
        Slasher slasherImplementation = new Slasher(strategyManager, delegation);
        EigenPodManager eigenPodManagerImplementation =
            new EigenPodManagerHarness(ethPOSDeposit, eigenPodBeacon, strategyManager, slasher, delegation);
        DelayedWithdrawalRouter delayedWithdrawalRouterImplementation = new DelayedWithdrawalRouter(eigenPodManager);

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(delegation))),
            address(delegationImplementation),
            abi.encodeWithSelector(
                DelegationManager.initialize.selector,
                proxyAdminOwner,
                pauserReg,
                0, /*initialPausedStatus*/
                minWithdrawalDelayBlocks,
                initializeStrategiesToSetDelayBlocks,
                initializeWithdrawalDelayBlocks
            )
        );
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(avsDirectory))),
            address(avsDirectoryImplementation),
            abi.encodeWithSelector(
                AVSDirectory.initialize.selector, proxyAdminOwner, pauserReg, 0 /*initialPausedStatus*/
            )
        );
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(strategyManager))),
            address(strategyManagerImplementation),
            abi.encodeWithSelector(
                StrategyManager.initialize.selector,
                proxyAdminOwner,
                proxyAdminOwner,
                pauserReg,
                0 /*initialPausedStatus*/
            )
        );
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(slasher))),
            address(slasherImplementation),
            abi.encodeWithSelector(Slasher.initialize.selector, proxyAdminOwner, pauserReg, 0 /*initialPausedStatus*/ )
        );
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation),
            abi.encodeWithSelector(
                EigenPodManager.initialize.selector,
                type(uint256).max, // maxPods
                beaconChainOracleAddress,
                proxyAdminOwner,
                pauserReg,
                0 /*initialPausedStatus*/
            )
        );
        uint256 initPausedStatus = 0;
        uint256 withdrawalDelayBlocks = PARTIAL_WITHDRAWAL_FRAUD_PROOF_PERIOD_BLOCKS;
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(delayedWithdrawalRouter))),
            address(delayedWithdrawalRouterImplementation),
            abi.encodeWithSelector(
                DelayedWithdrawalRouter.initialize.selector,
                proxyAdminOwner,
                pauserReg,
                initPausedStatus,
                withdrawalDelayBlocks
            )
        );

        address[] memory strategies = new address[](2);

        strategies[0] = _deployWETHStrategy();
        strategies[1] = beaconChainETHStrategy;

        return Deployments({
            proxyAdminOwner: proxyAdminOwner,
            proxyAdmin: address(proxyAdmin),
            pauserRegistry: address(pauserReg),
            delegationManager: address(delegation),
            eigenPodManager: address(eigenPodManager),
            strategyManager: address(strategyManager),
            slasher: address(slasher),
            avsDirectory: address(avsDirectory),
            strategies: strategies
        });
    }

    function _deployWETHStrategy() internal returns (address) {
        IERC20 weth = new MockERC20("weth", "WETH");
        StrategyBase impl = new StrategyBase(strategyManager);
        IStrategy wethStrat = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(impl),
                    address(proxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, weth, pauserReg)
                )
            )
        );
        return address(wethStrat);
    }
}
