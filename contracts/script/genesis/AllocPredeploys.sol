// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { EIP1967Helper } from "./utils/EIP1967Helper.sol";

import { PortalRegistry } from "src/xchain/PortalRegistry.sol";
import { OmniBridgeNative } from "src/token/OmniBridgeNative.sol";
import { Staking } from "src/octane/Staking.sol";

/**
 * @title AllocPredeploys
 * @notice Generate predeploy genesis allocs
 */
contract AllocPredeploys is Script {
    struct Config {
        address admin;
        bool enableStakingAllowlist;
        string output;
    }

    Config internal cfg;
    address internal deployer;

    function setUp() public {
        deployer = makeAddr("deployer");
    }

    function runWithCfg(Config calldata config) public {
        cfg = config;

        vm.startPrank(deployer);
        setProxies();
        setProxyAdmin();
        setPortalRegistry();
        setOmniBridgeNative();
        setWOmni();
        setStaking();
        setSlashing();
        vm.stopPrank();

        vm.resetNonce(msg.sender);
        vm.deal(msg.sender, 0);

        vm.resetNonce(deployer);
        vm.deal(deployer, 0);

        vm.dumpState(config.output);
    }

    /**
     * @notice Predeploy transparent proxies for each namespace
     */
    function setProxies() public {
        address[] memory namespaces = Predeploys.namespaces();
        for (uint256 i = 0; i < namespaces.length; i++) {
            setNamespaceProxies(namespaces[i]);
        }
    }

    /**
     * @notice Predeploy transparent proxies for all predeploys in a namespace, setting implementation if active
     */
    function setNamespaceProxies(address ns) internal {
        require(uint32(uint160(ns)) == 0, "invalid namespace");

        bytes memory code = vm.getDeployedCode("TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy");

        for (uint160 i = 1; i <= Predeploys.NamespaceSize; i++) {
            address addr = address(uint160(ns) + i);
            if (Predeploys.notProxied(addr)) {
                continue;
            }

            vm.etch(addr, code);
            EIP1967Helper.setAdmin(addr, Predeploys.ProxyAdmin);

            if (Predeploys.isActivePredeploy(addr)) {
                address impl = Predeploys.impl(addr);
                EIP1967Helper.setImplementation(addr, impl);
            }
        }
    }

    /**
     * @notice Setup ProxyAdmin predeploy
     */
    function setProxyAdmin() internal {
        // proxy admin has no initializer, so we set slot manually

        bytes32 ownerSlot = bytes32(0);

        vm.store(Predeploys.ProxyAdmin, ownerSlot, bytes32(uint256(uint160(cfg.admin))));
        vm.etch(Predeploys.ProxyAdmin, vm.getDeployedCode("out/ProxyAdmin.sol/ProxyAdmin.0.8.24.json"));
    }

    /**
     * @notice Setup PortalRegistry predeploy
     */
    function setPortalRegistry() internal {
        address impl = Predeploys.impl(Predeploys.PortalRegistry);
        vm.etch(impl, vm.getDeployedCode("PortalRegistry.sol:PortalRegistry"));

        PortalRegistry(impl).disableInitializers();
        PortalRegistry(Predeploys.PortalRegistry).initialize(cfg.admin);
    }

    /**
     * @notice Setup OmniBridgeNative predeploy
     */
    function setOmniBridgeNative() internal {
        uint256 totalSupply = 1e6 * 1e18; // 100M total supply
        vm.deal(Predeploys.OmniBridgeNative, totalSupply);

        address impl = Predeploys.impl(Predeploys.OmniBridgeNative);

        vm.etch(impl, vm.getDeployedCode("OmniBridgeNative.sol:OmniBridgeNative"));

        OmniBridgeNative(impl).disableInitializers();
        OmniBridgeNative(Predeploys.OmniBridgeNative).initialize(cfg.admin);
    }

    /**
     * @notice Setup WOmni predeploy
     */
    function setWOmni() internal {
        // not proxied
        vm.etch(Predeploys.WOmni, vm.getDeployedCode("WOmni.sol:WOmni"));
    }

    /**
     * @notice Setup Staking predeploy
     */
    function setStaking() internal {
        address impl = Predeploys.impl(Predeploys.Staking);
        vm.etch(impl, vm.getDeployedCode("Staking.sol:Staking"));

        Staking(impl).disableInitializers();
        Staking(Predeploys.Staking).initialize(cfg.admin, cfg.enableStakingAllowlist);
    }

    /**
     * @notice Setup Slashing predeploy
     */
    function setSlashing() internal {
        address impl = Predeploys.impl(Predeploys.Slashing);
        vm.etch(impl, vm.getDeployedCode("Slashing.sol:Slashing"));
    }
}
