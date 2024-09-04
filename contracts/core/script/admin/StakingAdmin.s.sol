// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { InitializableHelper } from "script/utils/InitializableHelper.sol";
import { EIP1967Helper } from "script/utils/EIP1967Helper.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { Staking } from "src/octane/Staking.sol";
import { Staking_Test } from "test/octane/Staking.t.sol";

/**
 * @title StakingAdmin
 * @notice A colleciton of admin scripts for the Staking contract.
 */
contract StakingAdmin is Script {
    /**
     * @notice Set the allowlist of the staking contract.
     * @param admin     The owner of the staking contract.
     * @param allowlist The new allowlist.
     */
    function setAllowlist(address admin, address[] calldata allowlist) public {
        vm.startBroadcast(admin);
        Staking(Predeploys.Staking).setAllowlist(allowlist);
        vm.stopBroadcast();
    }

    /**
     * @notice Enable or disable the allowlist of the staking contract.
     * @param admin     The owner of the staking contract.
     * @param enabled   The new allowlist status.
     */
    function setAllowlistEnabled(address admin, bool enabled) public {
        vm.startBroadcast(admin);
        Staking(Predeploys.Staking).setAllowlistEnabled(enabled);
        vm.stopBroadcast();
    }

    /**
     * @notice Upgrade a staking contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     */
    function upgrade(address admin, address deployer) public {
        // deploy new implementation
        vm.startBroadcast(deployer);
        address impl = address(new Staking());
        vm.stopBroadcast();

        // upgrade proxy
        vm.startBroadcast(admin);

        address proxyAdmin = EIP1967Helper.getAdmin(Predeploys.Staking);
        ProxyAdmin(proxyAdmin).upgradeAndCall(ITransparentUpgradeableProxy(Predeploys.Staking), impl, "");

        vm.stopBroadcast();

        // quick checks
        require(InitializableHelper.areInitializersDisabled(impl), "initializers not disabled");
        require(Staking(Predeploys.Staking).owner() == admin, "owner not set");

        // run tests
        Staking_Test tests = new Staking_Test();
        tests.run(Predeploys.Staking);
    }
}
