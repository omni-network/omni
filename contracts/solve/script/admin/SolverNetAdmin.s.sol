// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {
    ITransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { InitializableHelper } from "script/utils/InitializableHelper.sol";
import { InitializableHelperSolady } from "script/utils/InitializableHelperSolady.sol";
import { EIP1967Helper } from "script/utils/EIP1967Helper.sol";

import { SolverNetInbox } from "src/SolverNetInbox.sol";
import { SolverNetOutbox, ISolverNetOutbox } from "src/SolverNetOutbox.sol";
import { SolverNetExecutor } from "src/SolverNetExecutor.sol";

import { SolverNetPostUpgradeTest } from "./SolverNetPostUpgradeTest.sol";

import { Script } from "forge-std/Script.sol";

/**
 * @title SolverNetAdmin
 * @notice A colleciton of SolverNet admin scripts.
 */
contract SolverNetAdmin is Script {
    /// @dev Config struct for upgrading all SolverNet contracts.
    struct UpgradeAllConfig {
        address admin;
        address deployer;
        address inbox;
        address outbox;
        address executor;
        address omni;
        address mailbox;
    }

    /// @dev Start broadcating from `sender`
    modifier withBroadcast(address sender) {
        vm.startBroadcast(sender);
        _;
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
     * @notice Upgrade the SolverNetInbox contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param proxy     The address of the SolverNetInbox proxy to upgrade.
     * @param omni      The address of the OmniPortal.
     * @param mailbox   The address of the mailbox to use for the SolverNetInbox.
     * @param data      Calldata to execute after upgrading the contract.
     */
    function upgradeSolverNetInbox(
        address admin,
        address deployer,
        address proxy,
        address omni,
        address mailbox,
        bytes calldata data
    ) public {
        SolverNetInbox inbox = SolverNetInbox(proxy);

        address owner = inbox.owner();
        // uint256 deployedAt = inbox.deployedAt();
        address _omni = address(inbox.omni());
        if (_omni.code.length == 0) _omni = address(0); // Allow for replacing portal address with zero address if not deployed
        uint8 pauseState = inbox.pauseState();
        uint248 offset = inbox.getLatestOrderOffset();

        vm.startBroadcast(deployer);
        address impl = address(new SolverNetInbox(omni, mailbox));
        vm.stopBroadcast();

        _upgradeProxy(admin, proxy, impl, data, true, true);

        require(inbox.owner() == owner, "owner changed");
        // NOTE: This is disabled because ArbSys on Arbitrum chains doesn't work when forked by anvil
        // We will need to address this if we don't want the deployedAt value to change for upgrades
        // require(inbox.deployedAt() > deployedAt, "deployedAt didn't increase");
        require(address(inbox.omni()) == _omni, "omni changed");
        require(inbox.pauseState() == pauseState, "pauseState changed");
        require(inbox.getLatestOrderOffset() == offset, "offset changed");

        new SolverNetPostUpgradeTest().runInbox(proxy);
    }

    /**
     * @notice Upgrade the SolverNetOutbox contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param proxy     The address of the SolverNetOutbox proxy to upgrade.
     * @param mailbox   The address of the mailbox to use for the SolverNetOutbox.
     * @param data      Calldata to execute after upgrading the contract.
     * @param chainIds  The chain IDs of the chains to upgrade the SolverNetOutbox for.
     * @param configs   The inbox configs to use for the SolverNetOutbox.
     */
    function upgradeSolverNetOutbox(
        address admin,
        address deployer,
        address proxy,
        address executor,
        address omni,
        address mailbox,
        bytes calldata data,
        uint64[] calldata chainIds,
        ISolverNetOutbox.InboxConfig[] calldata configs
    ) public {
        SolverNetOutbox outbox = SolverNetOutbox(proxy);

        address owner = outbox.owner();
        // uint256 deployedAt = outbox.deployedAt();
        if (address(outbox.omni()).code.length != 0) {
            require(address(outbox.omni()) == omni, "omni changed");
        }
        require(outbox.executor() == executor, "executor changed");

        vm.startBroadcast(deployer);
        address impl = address(new SolverNetOutbox(executor, omni, mailbox));
        vm.stopBroadcast();

        _upgradeProxy(admin, proxy, impl, data, true, true);

        require(outbox.owner() == owner, "owner changed");
        // NOTE: This is disabled because ArbSys on Arbitrum chains doesn't work when forked by anvil
        // We will need to address this if we don't want the deployedAt value to change for upgrades
        // require(outbox.deployedAt() > deployedAt, "deployedAt didn't increase");
        require(address(outbox.omni()) == omni, "omni changed post upgrade");
        require(outbox.executor() == executor, "executor changed post upgrade");

        new SolverNetPostUpgradeTest().runOutbox(proxy, chainIds, configs);
    }

    /**
     * @notice Upgrade the SolverNetExecutor contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param proxy     The address of the SolverNetExecutor proxy to upgrade.
     * @param outbox    The address of the SolverNetOutbox to use for the SolverNetExecutor.
     * @param data      Calldata to execute after upgrading the contract.
     * @param chainIds  The chain IDs of the chains to upgrade the SolverNetExecutor for.
     */
    function upgradeSolverNetExecutor(
        address admin,
        address deployer,
        address proxy,
        address outbox,
        bytes calldata data,
        uint64[] calldata chainIds
    ) public {
        SolverNetExecutor executor = SolverNetExecutor(payable(proxy));

        address _outbox = executor.outbox();

        vm.startBroadcast(deployer);
        address impl = address(new SolverNetExecutor(outbox));
        vm.stopBroadcast();

        _upgradeProxy(admin, proxy, impl, data, false, true); // Change the false parameter once its initializable

        require(executor.outbox() == _outbox, "outbox changed");

        new SolverNetPostUpgradeTest().runExecutor(proxy, chainIds);
    }

    /**
     * @notice Upgrade all SolverNet contracts.
     * @param config       Config struct containing all relevant addresses.
     * @param data         The (re)initializer calldata to execute while upgrading the contracts.
     * @param chainIds     All chain IDs connected to these contracts (used for testing outbox and executor)
     * @param inboxConfigs The inbox configs to use for the SolverNetOutbox.
     */
    function upgradeAll(
        UpgradeAllConfig calldata config,
        bytes[] calldata data,
        uint64[] calldata chainIds,
        ISolverNetOutbox.InboxConfig[] calldata inboxConfigs
    ) public {
        require(data.length == 3, "data array length must be 3");

        upgradeSolverNetInbox(config.admin, config.deployer, config.inbox, config.omni, config.mailbox, data[0]);

        upgradeSolverNetOutbox(
            config.admin,
            config.deployer,
            config.outbox,
            config.executor,
            config.omni,
            config.mailbox,
            data[1],
            chainIds,
            inboxConfigs
        );

        upgradeSolverNetExecutor(config.admin, config.deployer, config.executor, config.outbox, data[2], chainIds);
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
