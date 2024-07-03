// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { EIP1967Helper } from "./utils/EIP1967Helper.sol";

import { PortalRegistry } from "src/xchain/PortalRegistry.sol";
import { OmniBridgeNative } from "src/token/OmniBridgeNative.sol";
import { Staking } from "src/octane/Staking.sol";
import { Preinstalls } from "src/octane/Preinstalls.sol";

/**
 * @title AllocPredeploys
 * @notice Generate predeploy genesis allocs
 */
contract AllocPredeploys is Script {
    struct Config {
        address admin;
        uint256 chainId;
        bool enableStakingAllowlist;
        string output;
    }

    Config internal cfg;
    address internal deployer;

    function setUp() public {
        deployer = makeAddr("deployer");
    }

    function run(Config calldata config) public {
        _run(config);

        // Reset so its not included state dump
        vm.etch(msg.sender, "");
        vm.resetNonce(msg.sender);
        vm.deal(msg.sender, 0);

        vm.resetNonce(deployer);
        vm.deal(deployer, 0);

        vm.dumpState(cfg.output);
    }

    function runNoStateDump(Config calldata config) public {
        _run(config);
    }

    function _run(Config calldata config) internal {
        cfg = config;

        vm.chainId(cfg.chainId);

        vm.startPrank(deployer);
        setPredeploys();
        setPreinstalls();
        vm.stopPrank();
    }

    /**
     * @notice Predeploy transparent proxies for each namespace
     */
    function setProxies() internal {
        address[] memory namespaces = Predeploys.namespaces();
        for (uint256 i = 0; i < namespaces.length; i++) {
            setNamespaceProxies(namespaces[i]);
        }
    }

    /**
     * @notice Set all protocol predeploys
     */
    function setPredeploys() internal {
        setProxies();
        setProxyAdmin();
        setPortalRegistry();
        setOmniBridgeNative();
        setWOmni();
        setStaking();
        setSlashing();
    }

    /**
     * @notice Set all preinstalls (non protocol predeploys)
     */
    function setPreinstalls() internal {
        vm.etch(Preinstalls.MultiCall3, Preinstalls.MultiCall3Code);
        vm.etch(Preinstalls.Create2Deployer, Preinstalls.Create2DeployerCode);
        vm.etch(Preinstalls.Safe_v130, Preinstalls.Safe_v130Code);
        vm.etch(Preinstalls.SafeL2_v130, Preinstalls.SafeL2_v130Code);
        vm.etch(Preinstalls.MultiSendCallOnly_v130, Preinstalls.MultiSendCallOnly_v130Code);
        vm.etch(Preinstalls.SafeSingletonFactory, Preinstalls.SafeSingletonFactoryCode);
        vm.etch(Preinstalls.DeterministicDeploymentProxy, Preinstalls.DeterministicDeploymentProxyCode);
        vm.etch(Preinstalls.MultiSend_v130, Preinstalls.MultiSend_v130Code);
        vm.etch(Preinstalls.Permit2, Preinstalls.getPermit2Code(cfg.chainId));
        vm.etch(Preinstalls.SenderCreator_v060, Preinstalls.SenderCreator_v060Code);
        vm.etch(Preinstalls.EntryPoint_v060, Preinstalls.EntryPoint_v060Code);
        vm.etch(Preinstalls.SenderCreator_v070, Preinstalls.SenderCreator_v070Code);
        vm.etch(Preinstalls.EntryPoint_v070, Preinstalls.EntryPoint_v070Code);
        vm.etch(Preinstalls.ERC1820Registry, Preinstalls.ERC1820RegistryCode);
        vm.etch(Preinstalls.BeaconBlockRoots, Preinstalls.BeaconBlockRootsCode);

        // 4788 sender nonce must be incremented, since it's part of later upgrade-transactions.
        // For the upgrade-tx to not create a contract that conflicts with an already-existing copy,
        // the nonce must be bumped.
        vm.setNonce(Preinstalls.BeaconBlockRootsSender, 1);

        // we set nonce of all preinstalls to 1, reflecting regular user execution
        // as contracts are deployed with a nonce of 1 (post eip-161)
        vm.setNonce(Preinstalls.MultiCall3, 1);
        vm.setNonce(Preinstalls.Create2Deployer, 1);
        vm.setNonce(Preinstalls.Safe_v130, 1);
        vm.setNonce(Preinstalls.SafeL2_v130, 1);
        vm.setNonce(Preinstalls.MultiSendCallOnly_v130, 1);
        vm.setNonce(Preinstalls.SafeSingletonFactory, 1);
        vm.setNonce(Preinstalls.DeterministicDeploymentProxy, 1);
        vm.setNonce(Preinstalls.MultiSend_v130, 1);
        vm.setNonce(Preinstalls.Permit2, 1);
        vm.setNonce(Preinstalls.SenderCreator_v060, 1);
        vm.setNonce(Preinstalls.EntryPoint_v060, 1);
        vm.setNonce(Preinstalls.SenderCreator_v070, 1);
        vm.setNonce(Preinstalls.EntryPoint_v070, 1);
        vm.setNonce(Preinstalls.ERC1820Registry, 1);
        vm.setNonce(Preinstalls.BeaconBlockRoots, 1);
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
