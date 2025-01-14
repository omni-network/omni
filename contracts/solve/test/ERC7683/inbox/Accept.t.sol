// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";

contract SolverNet_Inbox_Accept_Test is TestBase {
    function setUp() public override {
        super.setUp();
        vm.chainId(srcChainId);
    }

    function test_accept_succeeds() public {
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
        vm.expectEmit(true, true, true, true, address(inbox));
        emit ISolverNetInbox.Accepted(expectedOrderId, solver);

        vm.prank(solver);
        inbox.accept(expectedOrderId);

        // Verify order state and history
        (
            IERC7683.ResolvedCrossChainOrder memory storedOrder,
            ISolverNetInbox.OrderState memory state,
            ISolverNetInbox.StatusUpdate[] memory history
        ) = inbox.getOrder(expectedOrderId);

        // Verify order state is now Accepted
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Accepted), "order state: status");
        assertEq(state.acceptedBy, solver, "order state: accepted by");

        // Verify order history
        assertEq(history.length, 2, "order history: length"); // Should have Open and Accept events
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "order history: initial status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "order history: initial timestamp");
        assertEq(uint8(history[1].status), uint8(ISolverNetInbox.Status.Accepted), "order history: accepted status");
        assertEq(history[1].timestamp, uint40(block.timestamp), "order history: accepted timestamp");

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
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Accepted), expectedOrderId, "latest accepted order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending), expectedOrderId, "latest pending order id"
        );

        // Verify token balances haven't changed after acceptance
        assertEq(token1.balanceOf(user), initialUserBalance - resolvedOrder.minReceived[0].amount, "user balance after");
        assertEq(
            token1.balanceOf(address(inbox)),
            initialInboxBalance + resolvedOrder.minReceived[0].amount,
            "inbox balance after"
        );
    }

    function test_accept_reverts_not_solver() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Open the order
        vm.prank(user);
        inbox.open(order);

        bytes32 orderId = inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending);

        // Try to accept as non-solver
        vm.prank(user);
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.accept(orderId);
    }

    function test_accept_reverts_order_not_pending() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Open the order
        vm.prank(user);
        inbox.open(order);

        bytes32 orderId = inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending);

        // Accept the order first
        vm.prank(solver);
        inbox.accept(orderId);

        // Try to accept it again
        vm.prank(solver);
        vm.expectRevert(ISolverNetInbox.OrderNotPending.selector);
        inbox.accept(orderId);
    }
}
