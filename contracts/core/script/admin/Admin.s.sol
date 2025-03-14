// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { InitializableHelper } from "script/utils/InitializableHelper.sol";
import { InitializableHelperSolady } from "script/utils/InitializableHelperSolady.sol";
import { EIP1967Helper } from "script/utils/EIP1967Helper.sol";
import { OmniPortal } from "src/xchain/OmniPortal.sol";
import { FeeOracleV1 } from "src/xchain/FeeOracleV1.sol";
import { FeeOracleV2 } from "src/xchain/FeeOracleV2.sol";
import { PortalRegistry } from "src/xchain/PortalRegistry.sol";
import { OmniGasPump } from "src/token/OmniGasPump.sol";
import { OmniGasStation } from "src/token/OmniGasStation.sol";
import { OmniBridgeCommon } from "src/token/OmniBridgeCommon.sol";
import { OmniBridgeNative } from "src/token/OmniBridgeNative.sol";
import { OmniBridgeL1 } from "src/token/OmniBridgeL1.sol";
import { Staking } from "src/octane/Staking.sol";
import { Slashing } from "src/octane/Slashing.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { SolverNetInbox } from "solve/src/SolverNetInbox.sol";
import { SolverNetOutbox } from "solve/src/SolverNetOutbox.sol";
import { SolverNetMiddleman } from "solve/src/SolverNetMiddleman.sol";
import { SolverNetExecutor } from "solve/src/SolverNetExecutor.sol";
import { Script } from "forge-std/Script.sol";

import { BridgeL1PostUpgradeTest } from "./BridgeL1PostUpgradeTest.sol";
import { BridgeNativePostUpgradeTest } from "./BridgeNativePostUpgradeTest.sol";
import { StakingPostUpgradeTest } from "./StakingPostUpgradeTest.sol";
import { FeeOracleV2PostUpdateTest } from "./FeeOracleV2PostUpdateTest.sol";
import { SolverNetPostUpgradeTest } from "./SolverNetPostUpgradeTest.sol";

/**
 * @title Admin
 * @notice A colleciton of admin scripts.
 */
contract Admin is Script {
    /// @dev Start broadcating from `sender`
    modifier withBroadcast(address sender) {
        vm.startBroadcast(sender);
        _;
        vm.stopBroadcast();
    }

    /**
     * @notice Pause a portal contract.
     * @param admin     The owner of the portal contract.
     * @param portal    The address of the portal contract.
     */
    function pausePortal(address admin, address portal) public withBroadcast(admin) {
        OmniPortal(portal).pause();
    }

    /**
     * @notice Unpause a portal contract.
     * @param admin     The owner of the portal contract.
     * @param portal    The address of the portal contract.
     */
    function unpausePortal(address admin, address portal) public withBroadcast(admin) {
        OmniPortal(portal).unpause();
    }

    /**
     * @notice Upgrade a portal contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param portal    The address of the portal contract.
     * @param data      Calldata to execute after upgrading the contract.
     */
    function upgradePortal(address admin, address deployer, address portal, bytes calldata data) public {
        vm.startBroadcast(deployer);
        address impl = address(new OmniPortal());
        vm.stopBroadcast();

        _upgradeProxy(admin, portal, impl, data);

        // TODO: add post upgrade tests
    }

    /**
     * @notice Pause all xcalls from a portal.
     * @param admin     The owner of the portal contract.
     * @param portal    The address of the portal contract.
     */
    function pauseXCall(address admin, address portal) public withBroadcast(admin) {
        OmniPortal(portal).pauseXCall();
    }

    /**
     * @notice Unpause all xcalls from a portal to a specific chain.
     * @param admin     The owner of the portal contract.
     * @param portal    The address of the portal contract.
     * @param to        The chain id to pause xcalls to
     */
    function pauseXCallTo(address admin, address portal, uint64 to) public withBroadcast(admin) {
        OmniPortal(portal).pauseXCallTo(to);
    }

    /**
     * @notice Unpause all xcalls from a portal.
     * @param admin     The owner of the portal contract.
     * @param portal    The address of the portal contract.
     */
    function unpauseXCall(address admin, address portal) public withBroadcast(admin) {
        OmniPortal(portal).unpauseXCall();
    }

    /**
     * @notice Unpause all xcalls from a portal to a specific chain.
     * @param admin     The owner of the portal contract.
     * @param portal    The address of the portal contract.
     * @param to        The chain id to unpause xcalls to
     */
    function unpauseXCallTo(address admin, address portal, uint64 to) public withBroadcast(admin) {
        OmniPortal(portal).unpauseXCallTo(to);
    }

    /**
     * @notice Pause all xsubmits from a portal.
     * @param admin     The owner of the portal contract.
     * @param portal    The address of the portal contract.
     */
    function pauseXSubmit(address admin, address portal) public withBroadcast(admin) {
        OmniPortal(portal).pauseXSubmit();
    }

    /**
     * @notice Unpause all xsubmits from a portal.
     * @param admin     The owner of the portal contract.
     * @param portal    The address of the portal contract.
     */
    function unpauseXSubmit(address admin, address portal) public withBroadcast(admin) {
        OmniPortal(portal).unpauseXSubmit();
    }

    /**
     * @notice Pause all xsubmits from a portal to a specific chain.
     * @param admin     The owner of the portal contract.
     * @param portal    The address of the portal contract.
     * @param from      The chain id to pause xsubmits from
     */
    function pauseXSubmitFrom(address admin, address portal, uint64 from) public withBroadcast(admin) {
        OmniPortal(portal).pauseXSubmitFrom(from);
    }

    /**
     * @notice Unpause all xsubmits from a portal to a specific chain.
     * @param admin     The owner of the portal contract.
     * @param portal    The address of the portal contract.
     * @param from      The chain id to unpause xsubmits from
     */
    function unpauseXSubmitFrom(address admin, address portal, uint64 from) public withBroadcast(admin) {
        OmniPortal(portal).unpauseXSubmitFrom(from);
    }

    /**
     * @notice Pause a bridge action.
     * @param admin     The owner of the bridge contract.
     * @param bridge    The address of the bridge contract.
     * @param action    The action to pause.
     */
    function pauseBridge(address admin, address bridge, bytes32 action) public {
        OmniBridgeCommon b = OmniBridgeCommon(bridge);

        require(
            action == b.ACTION_WITHDRAW() || action == b.ACTION_BRIDGE() || action == b.KeyPauseAll(), "invalid action"
        );

        vm.startBroadcast(admin);
        b.pause(action);
        vm.stopBroadcast();
    }

    /**
     * @notice Unpause a bridge action.
     * @param admin     The owner of the bridge contract.
     * @param bridge    The address of the bridge contract.
     * @param action    The action to unpause.
     */
    function unpauseBridge(address admin, address bridge, bytes32 action) public {
        OmniBridgeCommon b = OmniBridgeCommon(bridge);

        require(
            action == b.ACTION_WITHDRAW() || action == b.ACTION_BRIDGE() || action == b.KeyPauseAll(), "invalid action"
        );

        vm.startBroadcast(admin);
        b.unpause(action);
        vm.stopBroadcast();
    }

    /**
     * @notice Pause all functions on the SolverNetInbox.
     * @param admin          The owner of the SolverNetInbox.
     * @param solverNetInbox The address of the SolverNetInbox.
     * @param pause          Whether to pause or unpause the SolverNetInbox.
     */
    function pauseSolverNetAll(address admin, address solverNetInbox, bool pause) public withBroadcast(admin) {
        SolverNetInbox(solverNetInbox).pauseAll(pause);
    }

    /**
     * @notice Pause the open function on the SolverNetInbox.
     * @param admin          The owner of the SolverNetInbox.
     * @param solverNetInbox The address of the SolverNetInbox.
     * @param pause          Whether to pause or unpause the SolverNetInbox.
     */
    function pauseSolverNetOpen(address admin, address solverNetInbox, bool pause) public withBroadcast(admin) {
        SolverNetInbox(solverNetInbox).pauseOpen(pause);
    }

    /**
     * @notice Pause the close function on the SolverNetInbox.
     * @param admin          The owner of the SolverNetInbox.
     * @param solverNetInbox The address of the SolverNetInbox.
     * @param pause          Whether to pause or unpause the SolverNetInbox.
     */
    function pauseSolverNetClose(address admin, address solverNetInbox, bool pause) public withBroadcast(admin) {
        SolverNetInbox(solverNetInbox).pauseClose(pause);
    }

    /**
     * @notice Upgrade a FeeOracleV1 contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param proxy     The address of the proxy to upgrade.
     */
    function upgradeFeeOracleV1(address admin, address deployer, address proxy, bytes calldata data) public {
        vm.startBroadcast(deployer);
        address impl = address(new FeeOracleV1());
        vm.stopBroadcast();

        _upgradeProxy(admin, proxy, impl, data);

        // TODO: add post upgrade tests
    }

    /**
     * @notice Upgrade a FeeOracleV2 contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param proxy     The address of the proxy to upgrade.
     */
    function upgradeFeeOracleV2(address admin, address deployer, address proxy, bytes calldata data) public {
        vm.startBroadcast(deployer);
        address impl = address(new FeeOracleV2());
        vm.stopBroadcast();

        _upgradeProxy(admin, proxy, impl, data);

        // TODO: add post upgrade tests
    }

    /**
     * @notice Upgrade an OmniGasPump contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param proxy     The address of the proxy to upgrade.
     */
    function upgradeGasPump(address admin, address deployer, address proxy, bytes calldata data) public {
        vm.startBroadcast(deployer);
        address impl = address(new OmniGasPump());
        vm.stopBroadcast();

        _upgradeProxy(admin, proxy, impl, data);

        // TODO: add post upgrade tests
    }

    /**
     * @notice Upgrade an OmniGasStation contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param proxy     The address of the proxy to upgrade.
     */
    function upgradeGasStation(address admin, address deployer, address proxy, bytes calldata data) public {
        vm.startBroadcast(deployer);
        address impl = address(new OmniGasStation());
        vm.stopBroadcast();

        _upgradeProxy(admin, proxy, impl, data);

        // TODO: add post upgrade tests
    }

    /**
     * @notice Upgrade the Staking predeploy.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     */
    function upgradeStaking(address admin, address deployer, bytes calldata data) public {
        Staking staking = Staking(Predeploys.Staking);

        // read storage pre-upgrade
        address owner = staking.owner();
        bool isAllowlistEnabled = staking.isAllowlistEnabled();

        vm.startBroadcast(deployer);
        address impl = address(new Staking());
        vm.stopBroadcast();

        _upgradeProxy(admin, Predeploys.Staking, impl, data);

        // assert storage unchanged
        require(staking.owner() == owner, "owner changed");
        require(staking.isAllowlistEnabled() == isAllowlistEnabled, "isAllowlistEnabled changed");

        new StakingPostUpgradeTest().run();
    }

    /**
     * @notice Upgrade the Staking predeploy.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     */
    function upgradeSlashing(address admin, address deployer, bytes calldata data) public {
        vm.startBroadcast(deployer);
        address impl = address(new Slashing());
        vm.stopBroadcast();

        _upgradeProxy(admin, Predeploys.Slashing, impl, data, false, false); // Slashing has no initializers

        // TODO: add post upgrade tests
    }

    /**
     * @notice Upgrade the OmniBridgeNative predeploy.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     */
    function upgradeBridgeNative(address admin, address deployer, bytes calldata data) public {
        OmniBridgeNative b = OmniBridgeNative(Predeploys.OmniBridgeNative);

        // retrieve pause states
        bool allPaused = b.isPaused(b.KeyPauseAll());
        bool bridgePaused = b.isPaused(b.ACTION_BRIDGE());
        bool withdrawPaused = b.isPaused(b.ACTION_WITHDRAW());

        // bridge must be paused
        // require(bridgePaused, "bridge is not paused");

        // read storage pre-upgrade
        address owner = b.owner();
        address omni = address(b.omni());
        address l1Bridge = b.l1Bridge();
        uint64 l1ChainId = b.l1ChainId();
        uint256 l1Deposits = WithL1BridgeBalanceView(address(b)).l1BridgeBalance();

        vm.startBroadcast(deployer);
        address impl = address(new OmniBridgeNative());
        vm.stopBroadcast();

        _upgradeProxy(admin, Predeploys.OmniBridgeNative, impl, data);

        // assert storage unchanged
        require(b.owner() == owner, "owner changed");
        require(b.l1ChainId() == l1ChainId, "l1ChainId changed");
        require(address(b.omni()) == omni, "omni changed");
        require(b.l1Deposits() == l1Deposits, "l1Deposits changed");
        require(b.l1Bridge() == l1Bridge, "l1Bridge changed");
        require(b.isPaused(b.KeyPauseAll()) == allPaused, "all paused state changed");
        require(b.isPaused(b.ACTION_BRIDGE()) == bridgePaused, "bridge paused state changed");
        require(b.isPaused(b.ACTION_WITHDRAW()) == withdrawPaused, "withdraw paused state changed");

        new BridgeNativePostUpgradeTest().run();
    }

    /**
     * @notice Upgrade the OmniBridgeL1 contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param proxy     The address of the proxy to upgrade.
     */
    function upgradeBridgeL1(address admin, address deployer, address proxy, bytes calldata data) public {
        OmniBridgeL1 b = OmniBridgeL1(proxy);

        // retrieve pause states
        bool allPaused = b.isPaused(b.KeyPauseAll());
        bool bridgePaused = b.isPaused(b.ACTION_BRIDGE());
        bool withdrawPaused = b.isPaused(b.ACTION_WITHDRAW());

        // bridge must be paused
        // require(bridgePaused, "bridge is not paused");

        // read storage pre-upgrade
        address owner = b.owner();
        address token = address(b.token());
        address omni = address(b.omni());

        vm.startBroadcast(deployer);
        address impl = address(new OmniBridgeL1(token));
        vm.stopBroadcast();

        _upgradeProxy(admin, proxy, impl, data);

        // assert storage unchanged
        require(b.owner() == owner, "owner changed");
        require(address(b.token()) == token, "token changed");
        require(address(b.omni()) == omni, "omni changed");
        require(b.isPaused(b.KeyPauseAll()) == allPaused, "all paused state changed");
        require(b.isPaused(b.ACTION_BRIDGE()) == bridgePaused, "bridge paused state changed");
        require(b.isPaused(b.ACTION_WITHDRAW()) == withdrawPaused, "withdraw paused state changed");

        new BridgeL1PostUpgradeTest().run(proxy);
    }

    /**
     * @notice Upgrade the PortalRegistry predeploy.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     */
    function upgradePortalRegistry(address admin, address deployer, bytes calldata data) public {
        vm.startBroadcast(deployer);
        address impl = address(new PortalRegistry());
        vm.stopBroadcast();

        _upgradeProxy(admin, Predeploys.PortalRegistry, impl, data);

        // TODO: add post upgrade tests
    }

    /**
     * @notice Upgrade the SolverNetInbox contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param proxy     The address of the SolverNetInbox proxy to upgrade.
     */
    function upgradeSolverNetInbox(address admin, address deployer, address proxy, bytes calldata data) public {
        SolverNetInbox inbox = SolverNetInbox(proxy);

        address owner = inbox.owner();
        // uint256 deployedAt = inbox.deployedAt();
        address omni = address(inbox.omni());
        uint8 defaultConfLevel = inbox.defaultConfLevel();
        uint8 pauseState = inbox.pauseState();
        uint248 offset = inbox.getLatestOrderOffset();

        vm.startBroadcast(deployer);
        address impl = address(new SolverNetInbox());
        vm.stopBroadcast();

        _upgradeProxy(admin, proxy, impl, data, true, true);

        require(inbox.owner() == owner, "owner changed");
        // NOTE: This is disabled because ArbSys on Arbitrum chains doesn't work when forked by anvil
        // We will need to address this if we don't want the deployedAt value to change for upgrades
        // require(inbox.deployedAt() > deployedAt, "deployedAt didn't increase");
        require(address(inbox.omni()) == omni, "omni changed");
        require(inbox.defaultConfLevel() == defaultConfLevel, "defaultConfLevel changed");
        require(inbox.pauseState() == pauseState, "pauseState changed");
        require(inbox.getLatestOrderOffset() == offset, "offset changed");

        new SolverNetPostUpgradeTest().runInbox(proxy);
    }

    /**
     * @notice Upgrade the SolverNetOutbox contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param proxy     The address of the SolverNetOutbox proxy to upgrade.
     */
    function upgradeSolverNetOutbox(address admin, address deployer, address proxy, bytes calldata data) public {
        SolverNetOutbox outbox = SolverNetOutbox(proxy);

        address owner = outbox.owner();
        // uint256 deployedAt = outbox.deployedAt();
        address omni = address(outbox.omni());
        address executor = outbox.executor();

        vm.startBroadcast(deployer);
        address impl = address(new SolverNetOutbox());
        vm.stopBroadcast();

        _upgradeProxy(admin, proxy, impl, data, true, true);

        require(outbox.owner() == owner, "owner changed");
        // NOTE: This is disabled because ArbSys on Arbitrum chains doesn't work when forked by anvil
        // We will need to address this if we don't want the deployedAt value to change for upgrades
        // require(outbox.deployedAt() > deployedAt, "deployedAt didn't increase");
        require(address(outbox.omni()) == omni, "omni changed");
        require(outbox.executor() == executor, "executor changed");

        new SolverNetPostUpgradeTest().runOutbox(proxy);
    }

    /**
     * @notice Upgrade the SolverNetMiddleman contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param proxy     The address of the SolverNetMiddleman proxy to upgrade.
     */
    function upgradeSolverNetMiddleman(address admin, address deployer, address proxy, bytes calldata data) public {
        vm.startBroadcast(deployer);
        address impl = address(new SolverNetMiddleman());
        vm.stopBroadcast();

        _upgradeProxy(admin, proxy, impl, data, false, true); // Change the false parameter once its initializable

        new SolverNetPostUpgradeTest().runMiddleman(proxy);
    }

    /**
     * @notice Upgrade the SolverNetExecutor contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param proxy     The address of the SolverNetExecutor proxy to upgrade.
     */
    function upgradeSolverNetExecutor(
        address admin,
        address deployer,
        address proxy,
        address outbox,
        bytes calldata data
    ) public {
        SolverNetExecutor executor = SolverNetExecutor(payable(proxy));

        address _outbox = executor.outbox();

        vm.startBroadcast(deployer);
        address impl = address(new SolverNetExecutor(outbox));
        vm.stopBroadcast();

        _upgradeProxy(admin, proxy, impl, data, false, true); // Change the false parameter once its initializable

        require(executor.outbox() == _outbox, "outbox changed");

        new SolverNetPostUpgradeTest().runExecutor(proxy);
    }

    /**
     * @notice Sets the OmniPortal's fee oracle to the new FeeOracleV2 contract.
     * @param admin         The address of the admin account, owner of the OmniPortal contract.
     * @param portal        The address of the OmniPortal contract.
     * @param newFeeOracle  The address of the new FeeOracleV2 contract.
     */
    function setPortalFeeOracleV2(address admin, address portal, address newFeeOracle) public {
        address oldFeeOracle = OmniPortal(portal).feeOracle();
        require(oldFeeOracle != newFeeOracle, "new fee oracle required");

        vm.startBroadcast(admin);
        OmniPortal(portal).setFeeOracle(newFeeOracle);
        vm.stopBroadcast();

        require(OmniPortal(portal).feeOracle() == newFeeOracle, "portal assignment failed");
        require(FeeOracleV2(newFeeOracle).manager() != address(0), "fee oracle not initialized");
        require(FeeOracleV2(newFeeOracle).version() == 2, "fee oracle not FeeOracleV2");

        new FeeOracleV2PostUpdateTest().run(newFeeOracle);
    }

    /**
     * @notice Upgrade a proxy contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param proxy     The address of the proxy to upgrade.
     * @param impl      The address of the new implementation.
     * @param data      Calldata to execute after upgrading the contract.
     */
    function _upgradeProxy(address admin, address proxy, address impl, bytes calldata data) internal {
        _upgradeProxy(admin, proxy, impl, data, true, false);
    }

    /**
     * @notice Upgrade a proxy contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param proxy     The address of the proxy to upgrade.
     * @param impl      The address of the new implementation.
     * @param data      Calldata to execute after upgrading the contract.
     * @param initializable Whether the implementation is initializable.
     * @param soladyInitializable Whether the implementation uses Solady's Initializable library. (OZ is default)
     */
    function _upgradeProxy(
        address admin,
        address proxy,
        address impl,
        bytes calldata data,
        bool initializable,
        bool soladyInitializable
    ) internal {
        address proxyAdmin = EIP1967Helper.getAdmin(proxy);

        vm.startBroadcast(admin);
        ProxyAdmin(proxyAdmin).upgradeAndCall(ITransparentUpgradeableProxy(proxy), impl, data);
        vm.stopBroadcast();

        if (initializable) {
            if (!soladyInitializable) {
                require(InitializableHelper.areInitializersDisabled(impl), "initializers not disabled");
            } else {
                require(InitializableHelperSolady.areInitializersDisabled(impl), "initializers not disabled");
            }
        }

        require(EIP1967Helper.getImplementation(proxy) == impl, "upgrade failed");
    }
}

/// @dev Helper interface for native bridge before l1BridgeBalance -> l1Deposits rename.
interface WithL1BridgeBalanceView {
    function l1BridgeBalance() external view returns (uint256);
}
