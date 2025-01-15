// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";

contract SolverNet_Inbox_Reject_Test is TestBase {
    function setUp() public override {
        super.setUp();
        vm.chainId(srcChainId);
    }

    function test_reject_succeeds() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Store initial state for comparison
        bytes32 expectedOrderId = inbox.getNextId();
        uint256 initialUserBalance = token1.balanceOf(user);
        uint256 initialInboxBalance = token1.balanceOf(address(inbox));

        // Open the order
        vm.prank(user);
        inbox.open(order);

        // Reject the order as solver with reason code 1
        uint8 rejectReason = 1;
        vm.expectEmit(true, true, true, true, address(inbox));
        emit ISolverNetInbox.Rejected(expectedOrderId, solver, rejectReason);

        vm.prank(solver);
        inbox.reject(expectedOrderId, rejectReason);

        // Verify order state and history
        (
            IERC7683.ResolvedCrossChainOrder memory storedOrder,
            ISolverNetInbox.OrderState memory state,
            ISolverNetInbox.StatusUpdate[] memory history
        ) = inbox.getOrder(expectedOrderId);

        // Verify that stored resolved order aligns with the original order
        assertResolved(user, resolvedOrder.orderId, order, storedOrder);

        // Verify order state is now Rejected
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Rejected), "order state: status");
        assertEq(state.acceptedBy, address(0), "order state: accepted by should be zero");

        // Verify order history
        assertEq(history.length, 2, "order history: length"); // Should have Open and Reject events
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "order history: initial status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "order history: initial timestamp");
        assertEq(uint8(history[1].status), uint8(ISolverNetInbox.Status.Rejected), "order history: rejected status");
        assertEq(history[1].timestamp, uint40(block.timestamp), "order history: rejected timestamp");

        // Verify latest order ID by status has been updated
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Rejected), expectedOrderId, "latest rejected order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending), expectedOrderId, "latest pending order id"
        );

        // Verify token balances haven't changed after rejection
        assertEq(token1.balanceOf(user), initialUserBalance - resolvedOrder.minReceived[0].amount, "user balance after");
        assertEq(
            token1.balanceOf(address(inbox)),
            initialInboxBalance + resolvedOrder.minReceived[0].amount,
            "inbox balance after"
        );
    }

    function test_reject_reverts_not_solver() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Open the order
        vm.prank(user);
        inbox.open(order);

        bytes32 orderId = inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending);

        // Try to reject as non-solver
        vm.prank(user);
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.reject(orderId, 1);
    }

    function test_reject_reverts_order_not_pending() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Open the order
        vm.prank(user);
        inbox.open(order);

        bytes32 orderId = inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending);

        // Reject the order first
        vm.prank(solver);
        inbox.reject(orderId, 1);

        // Try to reject it again
        vm.prank(solver);
        vm.expectRevert(ISolverNetInbox.OrderNotPending.selector);
        inbox.reject(orderId, 1);
    }
}
