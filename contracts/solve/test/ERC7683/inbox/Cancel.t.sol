// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";

contract SolverNet_Inbox_Cancel_Test is TestBase {
    function setUp() public override {
        super.setUp();
        vm.chainId(srcChainId);
    }

    function test_cancel_succeeds() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        fundUser(resolvedOrder.minReceived);

        // Store initial state for comparison
        bytes32 expectedOrderId = inbox.getNextId();
        uint256 initialUserBalance = token1.balanceOf(user);
        uint256 initialInboxBalance = token1.balanceOf(address(inbox));

        // Open the order
        vm.prank(user);
        inbox.open(order);

        // Cancel the order as the user
        vm.expectEmit(true, true, true, true, address(inbox));
        emit ISolverNetInbox.Reverted(expectedOrderId);

        vm.prank(user);
        inbox.cancel(expectedOrderId);

        // Verify order state and history
        (
            IERC7683.ResolvedCrossChainOrder memory storedOrder,
            ISolverNetInbox.OrderState memory state,
            ISolverNetInbox.StatusUpdate[] memory history
        ) = inbox.getOrder(expectedOrderId);

        // Verify that stored resolved order aligns with the original order
        assertResolved(user, resolvedOrder.orderId, order, storedOrder);

        // Verify order state is now Reverted
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Reverted), "order state: status");
        assertEq(state.acceptedBy, address(0), "order state: accepted by should be zero");

        // Verify order history
        assertEq(history.length, 2, "order history: length"); // Should have Open and Cancel events
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "order history: initial status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "order history: initial timestamp");
        assertEq(uint8(history[1].status), uint8(ISolverNetInbox.Status.Reverted), "order history: reverted status");
        assertEq(history[1].timestamp, uint40(block.timestamp), "order history: reverted timestamp");

        // Verify latest order ID by status has been updated
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Reverted), expectedOrderId, "latest reverted order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending), expectedOrderId, "latest pending order id"
        );

        // Verify token balances after cancellation (deposits should be returned)
        assertEq(token1.balanceOf(user), initialUserBalance, "user balance after");
        assertEq(token1.balanceOf(address(inbox)), initialInboxBalance, "inbox balance after");
    }

    function test_cancel_empty_user_address_succeeds() public {
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        ISolverNet.OrderData memory orderData = abi.decode(order.orderData, (ISolverNet.OrderData));
        orderData.user = address(0);
        order.orderData = abi.encode(orderData);

        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        fundUser(resolvedOrder.minReceived);

        // Store initial state for comparison
        bytes32 expectedOrderId = inbox.getNextId();
        uint256 initialUserBalance = token1.balanceOf(user);
        uint256 initialInboxBalance = token1.balanceOf(address(inbox));

        // Expect the event with the correct namespace and data
        vm.expectEmit(true, true, true, true, address(inbox));
        emit IERC7683.Open(expectedOrderId, resolvedOrder);

        vm.prank(user);
        inbox.open(order);

        // Cancel the order as the user
        vm.expectEmit(true, true, true, true, address(inbox));
        emit ISolverNetInbox.Reverted(expectedOrderId);

        vm.prank(user);
        inbox.cancel(expectedOrderId);

        // Verify order state and history
        (
            IERC7683.ResolvedCrossChainOrder memory storedOrder,
            ISolverNetInbox.OrderState memory state,
            ISolverNetInbox.StatusUpdate[] memory history
        ) = inbox.getOrder(expectedOrderId);

        // Verify that stored resolved order aligns with the original order
        assertResolved(user, resolvedOrder.orderId, order, storedOrder);

        // Verify order state is now Reverted
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Reverted), "order state: status");
        assertEq(state.acceptedBy, address(0), "order state: accepted by should be zero");

        // Verify order history
        assertEq(history.length, 2, "order history: length"); // Should have Open and Cancel events
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "order history: initial status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "order history: initial timestamp");
        assertEq(uint8(history[1].status), uint8(ISolverNetInbox.Status.Reverted), "order history: reverted status");
        assertEq(history[1].timestamp, uint40(block.timestamp), "order history: reverted timestamp");

        // Verify latest order ID by status has been updated
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Reverted), expectedOrderId, "latest reverted order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending), expectedOrderId, "latest pending order id"
        );

        // Verify token balances after cancellation (deposits should be returned)
        assertEq(token1.balanceOf(user), initialUserBalance, "user balance after");
        assertEq(token1.balanceOf(address(inbox)), initialInboxBalance, "inbox balance after");
    }

    function test_cancel_reverts_not_user() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        fundUser(resolvedOrder.minReceived);

        // Open the order
        vm.prank(user);
        inbox.open(order);

        bytes32 orderId = inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending);

        // Try to cancel as non-user (solver)
        vm.prank(solver);
        vm.expectRevert(abi.encodeWithSelector(Ownable.Unauthorized.selector));
        inbox.cancel(orderId);
    }

    function test_cancel_reverts_order_not_pending_or_rejected() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        fundUser(resolvedOrder.minReceived);

        // Open the order
        vm.prank(user);
        inbox.open(order);

        bytes32 orderId = inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending);

        // Cancel the order first
        vm.prank(user);
        inbox.cancel(orderId);

        // Try to cancel it again
        vm.prank(user);
        vm.expectRevert(ISolverNetInbox.OrderNotPending.selector);
        inbox.cancel(orderId);
    }
}
