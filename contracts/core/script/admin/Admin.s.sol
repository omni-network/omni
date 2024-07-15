// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { InitializableHelper } from "script/utils/InitializableHelper.sol";
import { EIP1967Helper } from "script/utils/EIP1967Helper.sol";
import { OmniPortal } from "src/xchain/OmniPortal.sol";

/**
 * @title Admin
 * @notice A colleciton of admin scripts.
 */
contract Admin is Script {
    /**
     * @notice Pause a portal contract.
     * @param admin     The owner of the portal contract.
     * @param portal    The address of the portal contract.
     */
    function pausePortal(address admin, address portal) public {
        vm.startBroadcast(admin);
        OmniPortal(portal).pause();
        vm.stopBroadcast();
    }

    /**
     * @notice Unpause a portal contract.
     * @param admin     The owner of the portal contract.
     * @param portal    The address of the portal contract.
     */
    function unpausePortal(address admin, address portal) public {
        vm.startBroadcast(admin);
        OmniPortal(portal).unpause();
        vm.stopBroadcast();
    }

    /**
     * @notice Upgrade a portal contract.
     * @param admin     The address of the admin account, owner of the proxy admin
     * @param deployer  The address of the account that will deploy the new implementation.
     * @param portal    The address of the portal contract.
     * @param data      Calldata to execute after upgrading the contract.
     */
    function upgradePortal(address admin, address deployer, address portal, bytes calldata data) public {
        // deploy new implementation
        vm.startBroadcast(deployer);
        address impl = address(new OmniPortal());
        vm.stopBroadcast();

        // upgrade proxy
        vm.startBroadcast(admin);
        address proxyAdmin = EIP1967Helper.getAdmin(portal);
        ProxyAdmin(proxyAdmin).upgradeAndCall(ITransparentUpgradeableProxy(portal), impl, data);
        vm.stopBroadcast();

        // run tests
        // TODO: add more
        require(InitializableHelper.areInitializersDisabled(impl), "initializers not disabled");
    }
}
