// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { PortalRegistry } from "src/xchain/PortalRegistry.sol";
import { OmniBridgeNative } from "src/token/OmniBridgeNative.sol";
import { Staking } from "src/octane/Staking.sol";
import { Preinstalls } from "src/octane/Preinstalls.sol";
import { InitializableHelper } from "script/utils/InitializableHelper.sol";
import { EIP1967Helper } from "script/utils/EIP1967Helper.sol";
import { Script } from "forge-std/Script.sol";

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

    /**
     * @notice Deployment config passed to run / runNoStateDump
     */
    Config internal cfg;

    /**
     * @notice Predeploy deployer address, used for each `new` call in this script
     */
    address internal deployer = 0xDDdDddDdDdddDDddDDddDDDDdDdDDdDDdDDDDDDd;

    function run(Config calldata config) public {
        _run(config);

        // Reset so its not included state dump
        vm.etch(msg.sender, "");
        vm.resetNonce(msg.sender);
        vm.deal(msg.sender, 0);

        // Note we do not reset deployer nonce, because the ProxyAdmin contracts deployed per each temp
        // TransparentUpgradeableProxy are downstream of the deployers CREATE address chain.
        //
        // Keeping the nonce ensures no new deployments conflict with genesis deployments - though this will likely
        // never happen, as the deployer address does not have a known private key
        //
        // We could reset the nonce. In this case, the first Predeploys.NamespaceSize * 2 deployments from the deployer
        // address would, that themselves deploy new contracts the constructor, would revert.

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

        for (uint160 i = 1; i <= Predeploys.NamespaceSize; i++) {
            address addr = address(uint160(ns) + i);
            if (Predeploys.notProxied(addr)) {
                continue;
            }

            address impl = Predeploys.isActivePredeploy(addr) ? Predeploys.impl(addr) : address(0);

            // set impl code to non-zero length, so it passes TransparentUpgradeableProxy constructor check
            // assert it is not already set
            require(impl.code.length == 0, "impl already set");
            vm.etch(impl, "00");

            // new use new, so that the immutable variable the holds the ProxyAdmin addr is set in properly in bytecode
            address tmp = address(new TransparentUpgradeableProxy(impl, cfg.admin, ""));
            vm.etch(addr, tmp.code);

            // set implempentation storage manually
            EIP1967Helper.setImplementation(addr, impl);

            // set admin storage, to follow EIP1967 standard
            EIP1967Helper.setAdmin(addr, EIP1967Helper.getAdmin(tmp));

            // reset impl
            vm.etch(impl, "");

            //
            vm.etch(tmp, "");

            // can we reset nonce here? we are using "deployer" addr
            vm.resetNonce(tmp);
        }

        vm.etch(address(0), "");
    }

    /**
     * @notice Setup PortalRegistry predeploy
     */
    function setPortalRegistry() internal {
        address impl = Predeploys.impl(Predeploys.PortalRegistry);
        vm.etch(impl, vm.getDeployedCode("PortalRegistry.sol:PortalRegistry"));

        InitializableHelper.disableInitializers(impl);
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

        InitializableHelper.disableInitializers(impl);
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

        InitializableHelper.disableInitializers(impl);
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
