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

        // Verify order data matches resolved order
        assertEq(storedOrder.user, user, "stored order: user");
        assertEq(storedOrder.originChainId, resolvedOrder.originChainId, "stored order: origin chain id");
        assertEq(storedOrder.openDeadline, uint32(block.timestamp), "stored order: open deadline");
        assertEq(storedOrder.fillDeadline, resolvedOrder.fillDeadline, "stored order: fill deadline");
        assertEq(storedOrder.orderId, resolvedOrder.orderId, "stored order: order id");

        // Verify maxSpent (solver outputs)
        assertEq(storedOrder.maxSpent.length, resolvedOrder.maxSpent.length, "stored order: max spent length");
        for (uint256 i = 0; i < storedOrder.maxSpent.length; i++) {
            assertEq(storedOrder.maxSpent[i].token, resolvedOrder.maxSpent[i].token, "stored order: max spent token");
            assertEq(storedOrder.maxSpent[i].amount, resolvedOrder.maxSpent[i].amount, "stored order: max spent amount");
            assertEq(
                storedOrder.maxSpent[i].recipient,
                resolvedOrder.maxSpent[i].recipient,
                "stored order: max spent recipient"
            );
            assertEq(
                storedOrder.maxSpent[i].chainId, resolvedOrder.maxSpent[i].chainId, "stored order: max spent chain id"
            );
        }

        // Verify minReceived (user deposits)
        assertEq(storedOrder.minReceived.length, resolvedOrder.minReceived.length, "stored order: min received length");
        for (uint256 i = 0; i < storedOrder.minReceived.length; i++) {
            assertEq(
                storedOrder.minReceived[i].token, resolvedOrder.minReceived[i].token, "stored order: min received token"
            );
            assertEq(
                storedOrder.minReceived[i].amount,
                resolvedOrder.minReceived[i].amount,
                "stored order: min received amount"
            );
            assertEq(
                storedOrder.minReceived[i].recipient,
                resolvedOrder.minReceived[i].recipient,
                "stored order: min received recipient"
            );
            assertEq(
                storedOrder.minReceived[i].chainId,
                resolvedOrder.minReceived[i].chainId,
                "stored order: min received chain id"
            );
        }

        // Verify fill instructions
        assertEq(
            storedOrder.fillInstructions.length,
            resolvedOrder.fillInstructions.length,
            "stored order: fill instructions length"
        );
        for (uint256 i = 0; i < storedOrder.fillInstructions.length; i++) {
            assertEq(
                storedOrder.fillInstructions[i].destinationChainId,
                resolvedOrder.fillInstructions[i].destinationChainId,
                "stored order: fill instructions chain id"
            );
            assertEq(
                storedOrder.fillInstructions[i].destinationSettler,
                resolvedOrder.fillInstructions[i].destinationSettler,
                "stored order: fill instructions destination"
            );
            assertEq(
                keccak256(storedOrder.fillInstructions[i].originData),
                keccak256(resolvedOrder.fillInstructions[i].originData),
                "stored order: fill instructions origin data"
            );
        }

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

    function test_mark_fulfilled_reverts_wrong_call_hash() public {
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
        vm.expectRevert(ISolverNetInbox.WrongCallHash.selector);
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
