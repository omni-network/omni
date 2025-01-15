// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_MarkFulfilled_Test is TestBase {
    function setUp() public override {
        super.setUp();
        vm.chainId(srcChainId);
    }

    function test_mark_fulfilled_succeeds() public {
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

        // Accept the order as solver
        vm.prank(solver);
        inbox.accept(expectedOrderId);

        // Mock xcall from outbox with fillHash
        bytes32 fillHash = fillHash(expectedOrderId, resolvedOrder.fillInstructions[0].originData);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(ISolverNetInbox.markFilled, (expectedOrderId, fillHash)),
            100_000
        );

        // Verify order state and history
        (
            IERC7683.ResolvedCrossChainOrder memory storedOrder,
            ISolverNetInbox.OrderState memory state,
            ISolverNetInbox.StatusUpdate[] memory history
        ) = inbox.getOrder(expectedOrderId);

        // Verify that stored resolved order aligns with the original order
        assertResolved(user, resolvedOrder.orderId, order, storedOrder);

        // Verify order state is now Filled
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Filled), "order state: status");
        assertEq(state.acceptedBy, solver, "order state: accepted by should be solver");

        // Verify order history
        assertEq(history.length, 3, "order history: length"); // Should have Open, Accept, and Fill events
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "order history: initial status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "order history: initial timestamp");
        assertEq(uint8(history[1].status), uint8(ISolverNetInbox.Status.Accepted), "order history: accepted status");
        assertEq(history[1].timestamp, uint40(block.timestamp), "order history: accepted timestamp");
        assertEq(uint8(history[2].status), uint8(ISolverNetInbox.Status.Filled), "order history: filled status");
        assertEq(history[2].timestamp, uint40(block.timestamp), "order history: filled timestamp");

        // Verify latest order ID by status has been updated
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Filled), expectedOrderId, "latest filled order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending), expectedOrderId, "latest pending order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Accepted), expectedOrderId, "latest accepted order id"
        );

        // Verify token balances haven't changed
        assertEq(token1.balanceOf(user), initialUserBalance - resolvedOrder.minReceived[0].amount, "user balance after");
        assertEq(
            token1.balanceOf(address(inbox)),
            initialInboxBalance + resolvedOrder.minReceived[0].amount,
            "inbox balance after"
        );
    }

    function test_mark_fulfilled_reverts_not_outbox() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Open the order
        vm.prank(user);
        inbox.open(order);

        bytes32 orderId = inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending);

        // Accept the order as solver
        vm.prank(solver);
        inbox.accept(orderId);

        // Mock xcall from a different address (not outbox)
        bytes32 fillHash = fillHash(orderId, resolvedOrder.fillInstructions[0].originData);
        vm.expectRevert(ISolverNetInbox.NotOutbox.selector);
        portal.mockXCall(
            destChainId,
            address(0x123), // Random address instead of outbox
            address(inbox),
            abi.encodeCall(ISolverNetInbox.markFilled, (orderId, fillHash)),
            100_000
        );
    }

    function test_mark_fulfilled_reverts_order_not_accepted() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Open the order
        vm.prank(user);
        inbox.open(order);

        bytes32 orderId = inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending);

        // Try to mark filled without accepting first
        bytes32 fillHash = fillHash(orderId, resolvedOrder.fillInstructions[0].originData);
        vm.expectRevert(ISolverNetInbox.OrderNotAccepted.selector);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(ISolverNetInbox.markFilled, (orderId, fillHash)),
            100_000
        );
    }

    function test_mark_fulfilled_reverts_wrong_fill_hash() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Open the order
        vm.prank(user);
        inbox.open(order);

        bytes32 orderId = inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending);

        // Accept the order as solver
        vm.prank(solver);
        inbox.accept(orderId);

        // Mock xcall with wrong fill hash
        bytes32 wrongFillHash = bytes32(uint256(1)); // Some random hash
        vm.expectRevert(ISolverNetInbox.WrongFillHash.selector);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(ISolverNetInbox.markFilled, (orderId, wrongFillHash)),
            100_000
        );
    }

    function test_mark_fulfilled_reverts_wrong_source_chain() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Open the order
        vm.prank(user);
        inbox.open(order);

        bytes32 orderId = inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending);

        // Accept the order as solver
        vm.prank(solver);
        inbox.accept(orderId);

        // Mock xcall from outbox with wrong source chain ID
        bytes32 fillHash = fillHash(orderId, resolvedOrder.fillInstructions[0].originData);
        vm.expectRevert(ISolverNetInbox.WrongSourceChain.selector);
        portal.mockXCall(
            destChainId + 1, // Wrong chain ID
            address(outbox),
            address(inbox),
            abi.encodeCall(ISolverNetInbox.markFilled, (orderId, fillHash)),
            100_000
        );
    }
}
