// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { Preinstalls } from "src/octane/Preinstalls.sol";
import { EIP1967Helper } from "script/genesis/utils/EIP1967Helper.sol";
import { AllocPredeploys } from "script/genesis/AllocPredeploys.s.sol";
import { Staking } from "src/octane/Staking.sol";
import { Test } from "forge-std/Test.sol";
import { Process } from "./utils/Process.sol";

/**
 * @title AllocPredeploys_Test
 * @notice Test suite for AllocPredeploys script.
 * @dev We inherit from AllocPredeploys so that vm.stateDump() is called from this contract,
 *      which keeps this contract's state out of the state dump.
 */
contract AllocPredeploys_Test is Test, AllocPredeploys {
    /**
     * @notice Tests predeploy allocs, asserting the number of allocs is expected.
     */
    function test_num_allocs() public {
        address admin = makeAddr("admin");
        string memory output = tmpfile();

        this.run(AllocPredeploys.Config({ admin: admin, chainId: 165, enableStakingAllowlist: false, output: output }));

        uint256 expected = 0;
        expected += 1024 * 2; // namespace size * 2
        expected += 1; // ProxyAdmin
        expected += 4; // predeploy implementations (excl. not prodiex WOmni and ProxyAdmin)
        expected += 15; // preinstalls
        expected += 1; // 4788 deployer account (nonce set to 1)

        assertEq(expected, getJSONKeyCount(output), "key count check");

        deleteFile(output);
    }

    function test_genesis() public {
        this.runNoStateDump(
            AllocPredeploys.Config({ admin: makeAddr("admin"), chainId: 165, enableStakingAllowlist: false, output: "" })
        );

        _testPredeploys();
        _testPreinstalls();
    }

    function _testPreinstalls() internal view {
        _assertPreinstall(Preinstalls.MultiCall3, Preinstalls.MultiCall3Code);
        _assertPreinstall(Preinstalls.Create2Deployer, Preinstalls.Create2DeployerCode);
        _assertPreinstall(Preinstalls.Safe_v130, Preinstalls.Safe_v130Code);
        _assertPreinstall(Preinstalls.SafeL2_v130, Preinstalls.SafeL2_v130Code);
        _assertPreinstall(Preinstalls.MultiSendCallOnly_v130, Preinstalls.MultiSendCallOnly_v130Code);
        _assertPreinstall(Preinstalls.SafeSingletonFactory, Preinstalls.SafeSingletonFactoryCode);
        _assertPreinstall(Preinstalls.DeterministicDeploymentProxy, Preinstalls.DeterministicDeploymentProxyCode);
        _assertPreinstall(Preinstalls.MultiSend_v130, Preinstalls.MultiSend_v130Code);
        _assertPreinstall(Preinstalls.Permit2, Preinstalls.getPermit2Code(cfg.chainId));
        _assertPreinstall(Preinstalls.SenderCreator_v060, Preinstalls.SenderCreator_v060Code);
        _assertPreinstall(Preinstalls.EntryPoint_v060, Preinstalls.EntryPoint_v060Code);
        _assertPreinstall(Preinstalls.SenderCreator_v070, Preinstalls.SenderCreator_v070Code);
        _assertPreinstall(Preinstalls.EntryPoint_v070, Preinstalls.EntryPoint_v070Code);
        _assertPreinstall(Preinstalls.BeaconBlockRoots, Preinstalls.BeaconBlockRootsCode);
        _assertPreinstall(Preinstalls.ERC1820Registry, Preinstalls.ERC1820RegistryCode);

        // BeaconBlockRootsSender must have nonce 1
        assertEq(vm.getNonce(Preinstalls.BeaconBlockRootsSender), 1, "BeaconBlockRootsSender nonce check");
    }

    function _assertPreinstall(address addr, bytes memory code) internal view {
        assertNotEq(code.length, 0, "must have code");
        assertEq(addr.code, code, "equal code must be deployed");
        assertEq(vm.getNonce(addr), 1, "preinstall account must have 1 nonce");
    }

    function _testPredeploys() internal {
        _testProxies();

        // test owners
        assertEq(cfg.admin, OwnableUpgradeable(Predeploys.PortalRegistry).owner(), "PortalRegistry owner check");
        assertEq(cfg.admin, OwnableUpgradeable(Predeploys.OmniBridgeNative).owner(), "OmniBridgeNative owner check");
        assertEq(cfg.admin, OwnableUpgradeable(Predeploys.Staking).owner(), "Staking owner check");

        // test proxies initialized
        assertTrue(_isInitialized(Predeploys.PortalRegistry), "PortalRegistry initialized check");
        assertTrue(_isInitialized(Predeploys.OmniBridgeNative), "OmniBridgeNative initialized check");
        assertTrue(_isInitialized(Predeploys.Staking), "Staking initialized check");

        // test initializers disabled on implementations
        assertTrue(
            _areInitializersDisabled(Predeploys.impl(Predeploys.PortalRegistry)), "PortalRegistry initializer check"
        );
        assertTrue(
            _areInitializersDisabled(Predeploys.impl(Predeploys.OmniBridgeNative)), "OmniBridgeNative initializer check"
        );
        assertTrue(_areInitializersDisabled(Predeploys.impl(Predeploys.Staking)), "Staking initializer check");
    }

    /**
     * @notice Test that all proxies have the correct admin set and implementaion set.
     */
    function _testProxies() internal {
        _forAllProxies(_testProxy);
    }

    /**
     * Test that a give proxy has the correct admin and implementation.
     */
    function _testProxy(address addr) internal view {
        assertEq(Predeploys.ProxyAdmin, EIP1967Helper.getAdmin(addr), "admin check");

        address expectedImpl = Predeploys.isActivePredeploy(addr) ? Predeploys.impl(addr) : address(0);
        assertEq(expectedImpl, EIP1967Helper.getImplementation(addr), "implementation check");
    }

    /**
     * @notice Call f for all proxies in each namespace.
     */
    function _forAllProxies(function (address) f) internal {
        address[] memory namespaces = Predeploys.namespaces();
        for (uint256 i = 0; i < namespaces.length; i++) {
            address ns = namespaces[i];

            for (uint160 j = 1; i <= Predeploys.NamespaceSize; i++) {
                address addr = address(uint160(ns) + j);

                if (Predeploys.notProxied(addr)) {
                    continue;
                }

                f(addr);
            }
        }
    }

    /**
     * @notice Returns the Initializable._initialized value for a given address, at slot 0.
     */
    function _getInitialized(address addr) internal view returns (uint256) {
        return uint256(vm.load(addr, bytes32(0)));
    }

    /**
     * @notice Returns true if the address has been initialized.
     */
    function _isInitialized(address addr) internal view returns (bool) {
        return _getInitialized(addr) == uint8(1);
    }

    /**
     * @notice Returns true if the initializers are disabled for a given address.
     */
    function _areInitializersDisabled(address addr) internal view returns (bool) {
        return _getInitialized(addr) == type(uint8).max;
    }

    //////////////////////////////////////////////////////////////////////////////
    //                      FS / JSON Utils                                     //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Creates a temp file and returns the path to it.
     */
    function tmpfile() internal returns (string memory) {
        string[] memory commands = new string[](3);
        commands[0] = "bash";
        commands[1] = "-c";
        commands[2] = "mktemp";
        bytes memory result = Process.run(commands);
        return string(result);
    }

    /**
     * @notice Deletes a file at a given filesystem path.
     */
    function deleteFile(string memory path) internal {
        string[] memory commands = new string[](3);
        commands[0] = "bash";
        commands[1] = "-c";
        commands[2] = string.concat("rm ", path);
        Process.run({ _command: commands, _allowEmpty: true });
    }

    /**
     * @notice Returns the number of top level keys in a JSON object at a given file path.
     */
    function getJSONKeyCount(string memory path) internal returns (uint256) {
        string[] memory commands = new string[](3);
        commands[0] = "bash";
        commands[1] = "-c";
        commands[2] = string.concat("jq 'keys | length' < ", path, " | xargs cast abi-encode 'f(uint256)'");
        return abi.decode(Process.run(commands), (uint256));
    }
}
