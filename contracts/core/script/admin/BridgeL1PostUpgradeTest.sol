// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { NominaBridgeNative } from "src/token/nomina/NominaBridgeNative.sol";
import { NominaBridgeL1 } from "src/token/nomina/NominaBridgeL1.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { Test } from "forge-std/Test.sol";
import { VmSafe } from "forge-std/Vm.sol";
import { PostHaltNominaL1BridgeWithdrawals } from "./PostHaltNominaL1BridgeWithdrawals.s.sol";

// solhint-disable state-visibility

/**
 * @title BridgeL1PostUpgradeTest
 * @dev Test NominaBridgeL1 post-upgrade functionality
 */
contract BridgeL1PostUpgradeTest is Test {
    NominaBridgeL1 b;
    MockPortal portal;
    IERC20 nomina;
    address owner;

    function run(address addr) public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setup(addr);
        _testBridge();
        _testWithdraw();
        _testPauseUnpause();
        _testPostHaltWithdrawals();
    }

    function _setup(address addr) internal {
        b = NominaBridgeL1(addr);
        nomina = b.NOMINA();
        owner = b.owner();

        // NominaBridgeL1 portal at slot 0, no admin setters
        portal = new MockPortal();
        vm.store(addr, bytes32(0), bytes32(uint256(uint160(address(portal)))));

        // After initializeV3, bridge and withdraw are paused
        // Verify they are paused as expected from initializeV3
        assertTrue(b.isPaused(b.ACTION_BRIDGE()), "Bridge should be paused after initializeV3");
        assertTrue(b.isPaused(b.ACTION_WITHDRAW()), "Withdraw should be paused after initializeV3");

        // Unpause for testing bridge and withdraw functionality
        vm.startPrank(owner);
        b.unpause(b.ACTION_BRIDGE());
        b.unpause(b.ACTION_WITHDRAW());
        vm.stopPrank();
    }

    function _testBridge() internal {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        address payor = address(this);
        uint256 fee = b.bridgeFee(payor, to, amount);
        uint256 bridgeBalance = nomina.balanceOf(address(b));

        vm.expectCall(
            address(portal),
            fee,
            abi.encodeCall(
                IOmniPortal.xcall,
                (
                    portal.omniChainId(),
                    ConfLevel.Finalized,
                    Predeploys.OmniBridgeNative,
                    abi.encodeCall(NominaBridgeNative.withdraw, (payor, to, amount)),
                    b.XCALL_WITHDRAW_GAS_LIMIT()
                )
            )
        );

        deal(address(nomina), payor, amount);
        vm.deal(payor, fee);
        vm.startPrank(payor);
        nomina.approve(address(b), amount);
        b.bridge{ value: fee }(to, amount);
        vm.stopPrank();

        assertEq(nomina.balanceOf(address(b)), bridgeBalance + amount);
        assertEq(nomina.balanceOf(payor), 0);
    }

    function _testWithdraw() internal {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        uint256 bridgeBalance = nomina.balanceOf(address(b));

        vm.expectCall(address(nomina), abi.encodeCall(nomina.transfer, (to, amount)));
        portal.mockXCall({
            sourceChainId: portal.omniChainId(),
            sender: Predeploys.OmniBridgeNative,
            to: address(b),
            data: abi.encodeCall(NominaBridgeL1.withdraw, (to, amount)),
            gasLimit: 100_000
        });

        assertEq(nomina.balanceOf(to), amount);
        assertEq(nomina.balanceOf(address(b)), bridgeBalance - amount);
    }

    function _testPauseUnpause() internal {
        vm.prank(owner);
        b.pause();

        assertTrue(b.isPaused(b.ACTION_BRIDGE()));
        assertTrue(b.isPaused(b.ACTION_WITHDRAW()));

        vm.prank(owner);
        b.unpause();

        assertFalse(b.isPaused(b.ACTION_BRIDGE()));
        assertFalse(b.isPaused(b.ACTION_WITHDRAW()));
    }

    function _testPostHaltWithdrawals() internal {
        // Create the withdrawal script
        PostHaltNominaL1BridgeWithdrawals script = new PostHaltNominaL1BridgeWithdrawals();

        // Verify the merkle root was set correctly during initializeV3
        bytes32 expectedRoot = script.getWithdrawalRoot();
        assertEq(b.postHaltRoot(), expectedRoot, "Post halt root should match");

        // Execute all post-halt withdrawals using runNoBroadcast
        // runNoBroadcast performs the same logic as run(), but without startBroadcast/stopBroadcast
        // This allows us to test the withdrawal mechanism in a simulated environment
        // The function will:
        // - Process all withdrawals from the JSON file
        // - Execute them in batches of up to 100
        // - Verify each account's balance increases by the expected amount
        // - Revert if any verification fails
        script.runNoBroadcast(address(b));

        // Verify all accounts are marked as claimed after runNoBroadcast completes
        PostHaltNominaL1BridgeWithdrawals.Withdrawal[] memory withdrawals = script.getWithdrawals();
        for (uint256 i = 0; i < withdrawals.length; i++) {
            assertTrue(b.postHaltClaimed(withdrawals[i].account), "All accounts should be claimed");
        }
    }
}
