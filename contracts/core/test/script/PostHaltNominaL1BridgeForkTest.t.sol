// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Test } from "forge-std/Test.sol";
import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {
    ITransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { NominaBridgeL1 } from "src/token/nomina/NominaBridgeL1.sol";
import { EIP1967Helper } from "script/utils/EIP1967Helper.sol";
import { PostHaltNominaL1BridgeWithdrawals } from "script/admin/PostHaltNominaL1BridgeWithdrawals.s.sol";

/**
 * @title PostHaltNominaL1BridgeForkTest
 * @notice Mainnet fork test for NominaBridgeL1 upgrade and post-halt withdrawals.
 *         Skipped by default. Run with ETH_RPC_URL set to a mainnet fork URL.
 */
contract PostHaltNominaL1BridgeForkTest is Test {
    address constant upgrader = 0xF8740c09f25E2cbF5C9b34Ef142ED7B343B42360;
    address constant deployer = 0x9496Bf1Bd2Fa5BCba72062cC781cC97eA6930A13;
    address constant bridge = 0xBBB3f5BcB1c8B0Ee932EfAba2fDEE566b83053A5;
    bytes32 constant postHaltRoot = 0xd3a7b265fb589d5808e6d7b3f390af8d964c8af96fe7009f301e282366c5461a;

    function setUp() public {
        string memory rpcUrl = vm.envOr("ETH_RPC_URL", string(""));
        vm.skip(bytes(rpcUrl).length == 0);
        vm.createSelectFork(rpcUrl);
    }

    function test_upgradeBridgeL1_and_withdraw() public {
        NominaBridgeL1 b = NominaBridgeL1(bridge);

        // Read storage pre-upgrade
        address owner = b.owner();
        address omni = address(b.OMNI());
        address nomina = address(b.NOMINA());
        address portal = address(b.portal());

        // Deploy new implementation as deployer
        vm.prank(deployer);
        address impl = address(new NominaBridgeL1(omni, nomina));

        // Upgrade proxy with initializeV3
        bytes memory initData = abi.encodeCall(NominaBridgeL1.initializeV3, (postHaltRoot));
        address proxyAdmin = EIP1967Helper.getAdmin(bridge);
        vm.prank(upgrader);
        ProxyAdmin(proxyAdmin).upgradeAndCall(ITransparentUpgradeableProxy(bridge), impl, initData);

        // Assert storage preserved
        assertEq(b.owner(), owner, "owner changed");
        assertEq(address(b.OMNI()), omni, "omni changed");
        assertEq(address(b.NOMINA()), nomina, "nomina changed");
        assertEq(address(b.portal()), portal, "portal changed");
        assertTrue(b.isPaused(b.ACTION_BRIDGE()), "bridge should be paused");
        assertTrue(b.isPaused(b.ACTION_WITHDRAW()), "withdraw should be paused");
        assertEq(b.postHaltRoot(), postHaltRoot, "root mismatch");

        // Test withdrawals
        PostHaltNominaL1BridgeWithdrawals script = new PostHaltNominaL1BridgeWithdrawals();
        script.runNoBroadcast(bridge);

        // Verify accounts are claimed
        PostHaltNominaL1BridgeWithdrawals.Withdrawal[] memory withdrawals = script.getWithdrawals();
        for (uint256 i = 0; i < withdrawals.length; i++) {
            assertTrue(b.postHaltClaimed(withdrawals[i].account), "should be claimed");
        }
    }
}
