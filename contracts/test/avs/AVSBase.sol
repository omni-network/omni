// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

import { IBLSApkRegistry } from "eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol";
import { BN254 } from "eigenlayer-middleware/src/libraries/BN254.sol";
import { RegistryCoordinatorHarness } from "eigenlayer-middleware/test/harnesses/RegistryCoordinatorHarness.t.sol";
import { StakeRegistryHarness } from "eigenlayer-middleware/test/harnesses/StakeRegistryHarness.sol";
import { BLSApkRegistryHarness } from "eigenlayer-middleware/test/harnesses/BLSApkRegistryHarness.sol";
import { RegistryCoordinator } from "eigenlayer-middleware/src/RegistryCoordinator.sol";
import { BLSApkRegistry } from "eigenlayer-middleware/src/BLSApkRegistry.sol";
import { StakeRegistry } from "eigenlayer-middleware/src/StakeRegistry.sol";
import { IRegistryCoordinator } from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import { IBLSApkRegistry } from "eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol";
import { IStakeRegistry } from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import { IndexRegistry } from "eigenlayer-middleware/src/IndexRegistry.sol";
import { IIndexRegistry } from "eigenlayer-middleware/src/interfaces/IIndexRegistry.sol";
import { BitmapUtils } from "eigenlayer-middleware/src/libraries/BitmapUtils.sol";

import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { EigenLayerTestHelper } from "./eigen/EigenLayerTestHelper.t.sol";
import { OmniAVSHarness } from "./OmniAVSHarness.sol";

contract AVSBase is EigenLayerTestHelper {
    using BN254 for BN254.G1Point;

    address public omniAVSOwner = makeAddr("omniAVSOwner");
    address public proxyAdminOwner = makeAddr("proxyAdminOwner");
    address public registryCoordinatorOwner = makeAddr("registryCoordinatorOwner");
    address public churnApprover = makeAddr("churnApprover");
    address public ejector = makeAddr("ejector");

    uint32 defaultMaxOperatorCount = 10;
    uint16 defaultKickBIPsOfOperatorStake = 15_000;
    uint16 defaultKickBIPsOfTotalStake = 150;
    uint96 minimumStakeForQuorum = 1 ether;

    ProxyAdmin public proxyAdmin;

    RegistryCoordinatorHarness public registryCoordinatorImplementation;
    StakeRegistryHarness public stakeRegistryImplementation;
    IBLSApkRegistry public blsApkRegistryImplementation;
    IIndexRegistry public indexRegistryImplementation;
    OmniAVSHarness public omniAVSImplementation;

    RegistryCoordinatorHarness public registryCoordinator;
    StakeRegistryHarness public stakeRegistry;
    BLSApkRegistryHarness public blsApkRegistry;
    IIndexRegistry public indexRegistry;
    OmniAVSHarness public omniAVS;

    /// @notice StakeRegistry, Constant used as a divisor in calculating weights.
    uint256 public constant WEIGHTING_DIVISOR = 1e18;

    // only one quorum
    bytes public QUORUM_NUMBERS = hex"00";

    IBLSApkRegistry.PubkeyRegistrationParams pubkeyRegistrationParams;

    function setUp() public override {
        super.setUp();
        _deployProxyAdmin();
        _deployRegistryCoordinatorProxy();
        _deployRegistries();
        _deployOmniAVS();
        _deployRegistryCoordinatorImpl();
        _initRegistryCoordinator();
    }

    function _deployProxyAdmin() internal {
        vm.prank(proxyAdminOwner);
        proxyAdmin = new ProxyAdmin();
    }

    function _deployRegistryCoordinatorProxy() internal {
        vm.prank(registryCoordinatorOwner);
        registryCoordinator = RegistryCoordinatorHarness(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );
    }

    function _deployRegistries() internal {
        vm.startPrank(registryCoordinatorOwner);

        stakeRegistry = StakeRegistryHarness(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );
        indexRegistry =
            IndexRegistry(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));
        blsApkRegistry = BLSApkRegistryHarness(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );

        vm.stopPrank();
        vm.startPrank(proxyAdminOwner);

        stakeRegistryImplementation = new StakeRegistryHarness(IRegistryCoordinator(registryCoordinator), delegation);
        blsApkRegistryImplementation = new BLSApkRegistryHarness(registryCoordinator);
        indexRegistryImplementation = new IndexRegistry(registryCoordinator);

        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(stakeRegistry))), address(stakeRegistryImplementation)
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(blsApkRegistry))), address(blsApkRegistryImplementation)
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(indexRegistry))), address(indexRegistryImplementation)
        );

        vm.stopPrank();
    }

    function _deployOmniAVS() internal {
        vm.startPrank(proxyAdminOwner);
        omniAVS =
            OmniAVSHarness(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));
        omniAVSImplementation = new OmniAVSHarness(delegation, registryCoordinator, stakeRegistry);
        proxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(omniAVS))), address(omniAVSImplementation));

        vm.stopPrank();

        uint64 omniChainId = 111;
        IOmniPortal stubPortal = IOmniPortal(makeAddr("tempPortal")); // TODO: use a real portal deployment

        omniAVS.initialize(omniAVSOwner, stubPortal, omniChainId);
    }

    function _deployRegistryCoordinatorImpl() internal {
        vm.startPrank(proxyAdminOwner);

        registryCoordinatorImplementation =
            new RegistryCoordinatorHarness(omniAVS, stakeRegistry, blsApkRegistry, indexRegistry);

        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(registryCoordinator))),
            address(registryCoordinatorImplementation)
        );

        vm.stopPrank();
    }

    function _initRegistryCoordinator() internal {
        IStakeRegistry.StrategyParams[][] memory strategyParams = _strategyParams();
        IRegistryCoordinator.OperatorSetParam[] memory operatorSetParams = _operatorSetParams();

        uint96[] memory minimumStakeForQuorums = new uint96[](1);
        minimumStakeForQuorums[0] = uint96(minimumStakeForQuorum);

        registryCoordinator.initialize(
            registryCoordinatorOwner,
            churnApprover,
            ejector,
            eigenLayerPauserReg, // TODO: do we use our own pauser?
            0, // initialPausedStatus
            operatorSetParams,
            minimumStakeForQuorums,
            strategyParams
        );
    }

    function _operatorSetParams() internal view returns (IRegistryCoordinator.OperatorSetParam[] memory params) {
        params = new IRegistryCoordinator.OperatorSetParam[](1);
        params[0] = IRegistryCoordinator.OperatorSetParam({
            maxOperatorCount: defaultMaxOperatorCount,
            kickBIPsOfOperatorStake: defaultKickBIPsOfOperatorStake,
            kickBIPsOfTotalStake: defaultKickBIPsOfTotalStake
        });
    }

    function _strategyParams() internal view returns (IStakeRegistry.StrategyParams[][] memory params) {
        params = new IStakeRegistry.StrategyParams[][](1);
        params[0] = _defaultQuorumStrategyParams();
    }

    function _defaultQuorumStrategyParams() internal view returns (IStakeRegistry.StrategyParams[] memory params) {
        params = new IStakeRegistry.StrategyParams[](1);
        params[0] = _wethStategyParams();
    }

    function _wethStategyParams() internal view returns (IStakeRegistry.StrategyParams memory) {
        return IStakeRegistry.StrategyParams({ strategy: wethStrat, multiplier: uint96(WEIGHTING_DIVISOR) });
    }
}
