// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Test } from "forge-std/Test.sol";
import { console } from "forge-std/console.sol";
import { PostHaltNominaL1BridgeWithdrawals } from "script/admin/PostHaltNominaL1BridgeWithdrawals.s.sol";
import { NominaBridgeL1 } from "src/token/nomina/NominaBridgeL1.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { MockOmni } from "nomina/test/utils/MockOmni.sol";
import { Nomina } from "nomina/src/token/Nomina.sol";

/**
 * @title PostHaltNominaL1BridgeWithdrawals_Test
 * @notice Test suite for PostHaltNominaL1BridgeWithdrawals script.
 */
contract PostHaltNominaL1BridgeWithdrawals_Test is Test {
    PostHaltNominaL1BridgeWithdrawals script;
    NominaBridgeL1 bridge;
    MockPortal portal;
    MockOmni omni;
    Nomina nomina;

    address owner;
    address proxyAdmin;
    address mintAuthority;
    address minter;
    address initialSupplyRecipient;
    uint256 totalOmniSupply = 100_000_000 ether;
    uint256 constant BRIDGE_BALANCE = 1_468_214_187_959_326_110_293_249_525;

    function setUp() public {
        vm.skip(!vm.envOr("RUN_SLOW_TESTS", false));

        script = new PostHaltNominaL1BridgeWithdrawals();

        owner = makeAddr("owner");
        proxyAdmin = makeAddr("proxyAdmin");
        mintAuthority = makeAddr("mintAuthority");
        minter = makeAddr("minter");
        initialSupplyRecipient = makeAddr("initialSupplyRecipient");

        portal = new MockPortal();
        omni = new MockOmni(totalOmniSupply, initialSupplyRecipient);
        nomina = new Nomina(address(omni), mintAuthority);

        vm.prank(mintAuthority);
        nomina.setMinter(minter);

        // Deploy and initialize bridge with correct root
        bytes32 correctRoot = script.getWithdrawalRoot();
        bridge = _freshBridge(correctRoot);
    }

    function test_getWithdrawals() public {
        PostHaltNominaL1BridgeWithdrawals.Withdrawal[] memory withdrawals = script.getWithdrawals();

        // Verify length and some samples
        assertEq(withdrawals.length, 7526);
        assertEq(withdrawals[0].account, 0xe04D63E6E6209C5A64be616443C00B2b6E705d92);
        assertEq(withdrawals[0].balance, 112_527_512_025_000_000_000);
        assertEq(withdrawals[500].account, 0x34a2199d7e305F8cEe4e51F3A798719251A12FA2);
        assertEq(withdrawals[500].balance, 1_301_758_396_090_019_968_399);
        assertEq(withdrawals[3000].account, 0x1414D15a81cDb257ACE4701067D5A855f322F0e7);
        assertEq(withdrawals[3000].balance, 7_500_000_000_000_000_000);
        assertEq(withdrawals[7000].account, 0x88a112CfC5648853c6405EB0F4Ced188C82635A0);
        assertEq(withdrawals[7000].balance, 7_378_741_541_974_874_231_821);
        assertEq(withdrawals[7500].account, 0x47C2162AA3AbCEB10442D6B51F25bf846a23A920);
        assertEq(withdrawals[7500].balance, 1_635_116_010_019_723_934_360);
        assertEq(withdrawals[7525].account, 0x41EFfB8F5d8fd4e4826D19fd53Ae5c469AAf2894);
        assertEq(withdrawals[7525].balance, 96_843_139_813_607_464_500);

        // Verify total balance
        uint256 totalBalance = 0;
        for (uint256 i = 0; i < withdrawals.length; i++) {
            totalBalance += withdrawals[i].balance;
        }
        assertEq(totalBalance, BRIDGE_BALANCE, "Total balance should match expected sum");
    }

    function test_getWithdrawalRoot_matchesExpected() public {
        bytes32 root = script.getWithdrawalRoot();

        assertEq(
            root,
            0xd3a7b265fb589d5808e6d7b3f390af8d964c8af96fe7009f301e282366c5461a,
            "Root does not match expected value"
        );
    }

    function test_runNoBroadcast() public {
        PostHaltNominaL1BridgeWithdrawals.Withdrawal[] memory withdrawals = script.getWithdrawals();

        // Execute all withdrawals in batches of 100
        script.runNoBroadcast(address(bridge));

        // Verify bridge balance is now zero (all funds withdrawn)
        assertEq(nomina.balanceOf(address(bridge)), 0, "Bridge balance should be zero");

        // Verify each account received correct amount and is marked as claimed
        for (uint256 i = 0; i < withdrawals.length; i++) {
            assertEq(nomina.balanceOf(withdrawals[i].account), withdrawals[i].balance, "Incorrect balance");
            assertTrue(bridge.postHaltClaimed(withdrawals[i].account), "Should be claimed");
        }
    }

    function test_runNoBroadcast_revertsWhenInvalidBridgeAddress() public {
        vm.expectRevert("Invalid bridge address");
        script.runNoBroadcast(address(0));
    }

    function test_runNoBroadcast_revertsWhenNoRootSet() public {
        // Deploy a fresh bridge without V3 initialization
        NominaBridgeL1 b = _freshBridge(bytes32(0));

        // Should revert because V3 not initialized (no root set)
        vm.expectRevert("Post halt root mismatch");
        script.runNoBroadcast(address(b));
    }

    function test_runNoBroadcast_revertsWhenRootMismatch() public {
        // Deploy a fresh bridge with wrong root
        bytes32 wrongRoot = keccak256("wrong root");
        NominaBridgeL1 b = _freshBridge(wrongRoot);

        // The script checks if the generated root matches what's in the contract
        vm.expectRevert("Post halt root mismatch");
        script.runNoBroadcast(address(b));
    }

    function test_runNoBroadcast_revertsWhenAlreadyClaimed() public {
        // First run succeeds
        script.runNoBroadcast(address(bridge));

        // Second attempt should fail
        vm.expectRevert("NominaBridge: already claimed");
        script.runNoBroadcast(address(bridge));
    }

    function _freshBridge(bytes32 rootToInitialize) internal returns (NominaBridgeL1) {
        address impl = address(new NominaBridgeL1(address(omni), address(nomina)));
        NominaBridgeL1 b = NominaBridgeL1(
            address(
                new TransparentUpgradeableProxy(
                    impl, proxyAdmin, abi.encodeCall(NominaBridgeL1.initialize, (owner, address(portal)))
                )
            )
        );

        b.initializeV2();

        // Fund bridge with total withdrawal amount.
        // Convert OMNI to NOM (with 1 wei buffer). Then transfer BRIDGE_BALANCE to the bridge.
        vm.startPrank(initialSupplyRecipient);
        omni.approve(address(nomina), (BRIDGE_BALANCE / 75) + 1);
        nomina.convert(initialSupplyRecipient, (BRIDGE_BALANCE / 75) + 1);
        nomina.transfer(address(b), BRIDGE_BALANCE);
        vm.stopPrank();

        // Initialize V3 with provided root (use bytes32(0) to skip V3 initialization)
        if (rootToInitialize != bytes32(0)) {
            vm.prank(proxyAdmin);
            b.initializeV3(rootToInitialize);
        }

        return b;
    }
}
