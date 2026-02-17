// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
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
        assertTrue(b.isPaused(b.ACTION_BRIDGE()), "Bridge should be paused after initializeV3");
        assertTrue(b.isPaused(b.ACTION_WITHDRAW()), "Withdraw should be paused after initializeV3");
    }

    function _testBridge() internal {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        address payor = address(this);
        uint256 fee = b.bridgeFee(payor, to, amount);

        deal(address(nomina), payor, amount);
        vm.deal(payor, fee);
        vm.startPrank(payor);
        nomina.approve(address(b), amount);
        vm.expectRevert("NominaBridge: paused");
        b.bridge{ value: fee }(to, amount);
        vm.stopPrank();
    }

    function _testWithdraw() internal {
        address to = makeAddr("to");
        uint256 amount = 1e18;

        vm.expectRevert("NominaBridge: paused");
        portal.mockXCall({
            sourceChainId: portal.omniChainId(),
            sender: Predeploys.OmniBridgeNative,
            to: address(b),
            data: abi.encodeCall(NominaBridgeL1.withdraw, (to, amount)),
            gasLimit: 100_000
        });
    }

    function _testPostHaltWithdrawals() internal {
        // Create the withdrawal script
        PostHaltNominaL1BridgeWithdrawals script = new PostHaltNominaL1BridgeWithdrawals();

        // Verify the merkle root was set correctly during initializeV3
        bytes32 expectedRoot = script.getWithdrawalRoot();
        assertEq(b.postHaltRoot(), expectedRoot, "Post halt root should match");
        assertEq(expectedRoot, 0xd3a7b265fb589d5808e6d7b3f390af8d964c8af96fe7009f301e282366c5461a);

        // Run first 50 and last 50 withdrawals to test without processing all 7526
        uint256 n = 50;
        uint256 total = script.TOTAL_WITHDRAWALS();

        script.runNoBroadcastRange(address(b), 0, n);
        script.runNoBroadcastRange(address(b), total - n, n);

        // Verify tested accounts are marked as claimed
        PostHaltNominaL1BridgeWithdrawals.Withdrawal[] memory withdrawals = script.getWithdrawals();
        for (uint256 i = 0; i < n; i++) {
            assertTrue(b.postHaltClaimed(withdrawals[i].account), "First 50 should be claimed");
        }
        for (uint256 i = total - n; i < total; i++) {
            assertTrue(b.postHaltClaimed(withdrawals[i].account), "Last 50 should be claimed");
        }
    }
}
