// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";

contract SolverNet_Inbox_Cancel_Test is TestBase {
    function setUp() public override {
        super.setUp();
        vm.chainId(srcChainId);
    }

    function test_cancel_pending_succeeds() public {
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

        // Verify order state is now Reverted
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Reverted), "order state: status");
        assertEq(state.acceptedBy, address(0), "order state: accepted by should be zero");

        // Verify order history
        assertEq(history.length, 2, "order history: length"); // Should have Open and Cancel events
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "order history: initial status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "order history: initial timestamp");
        assertEq(uint8(history[1].status), uint8(ISolverNetInbox.Status.Reverted), "order history: reverted status");
        assertEq(history[1].timestamp, uint40(block.timestamp), "order history: reverted timestamp");

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
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Reverted), expectedOrderId, "latest reverted order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending), expectedOrderId, "latest pending order id"
        );

        // Verify token balances after cancellation (deposits should be returned)
        assertEq(token1.balanceOf(user), initialUserBalance, "user balance after");
        assertEq(token1.balanceOf(address(inbox)), initialInboxBalance, "inbox balance after");
    }

    function test_cancel_rejected_succeeds() public {
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

        // Reject the order
        vm.prank(solver);
        inbox.reject(expectedOrderId, 1);

        // Cancel the rejected order
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

        // Verify order state is now Reverted
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Reverted), "order state: status");
        assertEq(state.acceptedBy, address(0), "order state: accepted by should be zero");

        // Verify order history
        assertEq(history.length, 3, "order history: length"); // Should have Open, Reject, and Cancel events
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "order history: initial status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "order history: initial timestamp");
        assertEq(uint8(history[1].status), uint8(ISolverNetInbox.Status.Rejected), "order history: rejected status");
        assertEq(history[1].timestamp, uint40(block.timestamp), "order history: rejected timestamp");
        assertEq(uint8(history[2].status), uint8(ISolverNetInbox.Status.Reverted), "order history: reverted status");
        assertEq(history[2].timestamp, uint40(block.timestamp), "order history: reverted timestamp");

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
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Reverted), expectedOrderId, "latest reverted order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending), expectedOrderId, "latest pending order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Rejected), expectedOrderId, "latest rejected order id"
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
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

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
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Open the order
        vm.prank(user);
        inbox.open(order);

        bytes32 orderId = inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending);

        // Cancel the order first
        vm.prank(user);
        inbox.cancel(orderId);

        // Try to cancel it again
        vm.prank(user);
        vm.expectRevert(ISolverNetInbox.OrderNotPendingOrRejected.selector);
        inbox.cancel(orderId);
    }
}
