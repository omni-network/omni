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

import { EigenLayerTestHelper } from "./eigen/EigenLayerTestHelper.t.sol";
import { OmniAVS } from "src/protocol/OmniAVS.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";

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

    BN254.G1Point internal defaultPubKey = BN254.G1Point(
        18_260_007_818_883_133_054_078_754_218_619_977_578_772_505_796_600_400_998_181_738_095_793_040_006_897,
        3_432_351_341_799_135_763_167_709_827_653_955_074_218_841_517_684_851_694_584_291_831_827_675_065_899
    );

    ProxyAdmin public proxyAdmin;

    RegistryCoordinatorHarness public registryCoordinatorImplementation;
    StakeRegistryHarness public stakeRegistryImplementation;
    IBLSApkRegistry public blsApkRegistryImplementation;
    IIndexRegistry public indexRegistryImplementation;
    OmniAVS public omniAVSImplementation;

    RegistryCoordinatorHarness public registryCoordinator;
    StakeRegistryHarness public stakeRegistry;
    BLSApkRegistryHarness public blsApkRegistry;
    IIndexRegistry public indexRegistry;
    OmniAVS public omniAVS;

    /// @notice StakeRegistry, Constant used as a divisor in calculating weights.
    uint256 public constant WEIGHTING_DIVISOR = 1e18;

    // only one quorum
    bytes public QUORUM_NUMBERS = hex"00";
    // string public constant defaultSocket = "12.34.56.78";

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

    /// @dev Deploy the ProxyAdmin contract, to be admin of all test proxies
    function _deployProxyAdmin() internal {
        vm.prank(proxyAdminOwner);
        proxyAdmin = new ProxyAdmin();
    }

    /// @dev Deploy the RegistryCoordinator proxy
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

        omniAVS = OmniAVS(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));

        uint64 omniChainId = 111;
        address stubPortal = makeAddr("tempPortal"); // TODO: use a real portal deployment

        omniAVSImplementation =
            new OmniAVS(delegation, registryCoordinator, stakeRegistry, IOmniPortal(stubPortal), omniChainId);

        proxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(omniAVS))), address(omniAVSImplementation));

        vm.stopPrank();

        omniAVS.initialize(omniAVSOwner);
    }

    /// @dev Deploy the RegistryCoordinator implementation
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
        minimumStakeForQuorums[0] = uint96(1000);

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

    /// @dev Single strategy (WETH) for now
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
