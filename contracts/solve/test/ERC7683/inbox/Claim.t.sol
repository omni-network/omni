// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";

contract SolverNet_Inbox_Claim_Test is TestBase {
    function setUp() public override {
        super.setUp();
        vm.chainId(srcChainId);
    }

    function test_claim_succeeds() public {
        // Create and open an order
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Store initial state for comparison
        bytes32 expectedOrderId = inbox.getNextId();
        uint256 initialUserBalance = token1.balanceOf(user);
        uint256 initialInboxBalance = token1.balanceOf(address(inbox));
        uint256 initialSolverBalance = token1.balanceOf(solver);

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

        // Claim the order as solver
        vm.expectEmit(true, true, true, true, address(inbox));
        emit ISolverNetInbox.Claimed(expectedOrderId, solver, solver, resolvedOrder.minReceived);

        vm.prank(solver);
        inbox.claim(expectedOrderId, solver);

        // Verify order state and history
        (
            IERC7683.ResolvedCrossChainOrder memory storedOrder,
            ISolverNetInbox.OrderState memory state,
            ISolverNetInbox.StatusUpdate[] memory history
        ) = inbox.getOrder(expectedOrderId);

        // Verify order state is now Claimed
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Claimed), "order state: status");
        assertEq(state.acceptedBy, solver, "order state: accepted by should be solver");

        // Verify order history
        assertEq(history.length, 4, "order history: length"); // Should have Open, Accept, Fill, and Claim events
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "order history: initial status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "order history: initial timestamp");
        assertEq(uint8(history[1].status), uint8(ISolverNetInbox.Status.Accepted), "order history: accepted status");
        assertEq(history[1].timestamp, uint40(block.timestamp), "order history: accepted timestamp");
        assertEq(uint8(history[2].status), uint8(ISolverNetInbox.Status.Filled), "order history: filled status");
        assertEq(history[2].timestamp, uint40(block.timestamp), "order history: filled timestamp");
        assertEq(uint8(history[3].status), uint8(ISolverNetInbox.Status.Claimed), "order history: claimed status");
        assertEq(history[3].timestamp, uint40(block.timestamp), "order history: claimed timestamp");

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
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed), expectedOrderId, "latest claimed order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending), expectedOrderId, "latest pending order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Accepted), expectedOrderId, "latest accepted order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Filled), expectedOrderId, "latest filled order id"
        );

        // Verify token balances after claim (deposits should be transferred to solver)
        assertEq(token1.balanceOf(user), initialUserBalance - resolvedOrder.minReceived[0].amount, "user balance after");
        assertEq(token1.balanceOf(address(inbox)), initialInboxBalance, "inbox balance after");
        assertEq(
            token1.balanceOf(solver), initialSolverBalance + resolvedOrder.minReceived[0].amount, "solver balance after"
        );
    }

    function test_claim_native_succeeds() public {
        // Create and open an order with native token deposit
        IERC7683.OnchainCrossChainOrder memory order = randNativeOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        // Deal native tokens to user
        vm.deal(user, resolvedOrder.minReceived[0].amount);

        // Store initial state for comparison
        bytes32 expectedOrderId = inbox.getNextId();
        uint256 initialUserBalance = user.balance;
        uint256 initialInboxBalance = address(inbox).balance;
        uint256 initialSolverBalance = solver.balance;

        // Open the order with native token value
        vm.prank(user);
        inbox.open{ value: resolvedOrder.minReceived[0].amount }(order);

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

        // Claim the order as solver
        vm.expectEmit(true, true, true, true, address(inbox));
        emit ISolverNetInbox.Claimed(expectedOrderId, solver, solver, resolvedOrder.minReceived);

        vm.prank(solver);
        inbox.claim(expectedOrderId, solver);

        // Verify order state and history
        (
            IERC7683.ResolvedCrossChainOrder memory storedOrder,
            ISolverNetInbox.OrderState memory state,
            ISolverNetInbox.StatusUpdate[] memory history
        ) = inbox.getOrder(expectedOrderId);

        // Verify order state is now Claimed
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Claimed), "order state: status");
        assertEq(state.acceptedBy, solver, "order state: accepted by should be solver");

        // Verify order history
        assertEq(history.length, 4, "order history: length"); // Should have Open, Accept, Fill, and Claim events
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "order history: initial status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "order history: initial timestamp");
        assertEq(uint8(history[1].status), uint8(ISolverNetInbox.Status.Accepted), "order history: accepted status");
        assertEq(history[1].timestamp, uint40(block.timestamp), "order history: accepted timestamp");
        assertEq(uint8(history[2].status), uint8(ISolverNetInbox.Status.Filled), "order history: filled status");
        assertEq(history[2].timestamp, uint40(block.timestamp), "order history: filled timestamp");
        assertEq(uint8(history[3].status), uint8(ISolverNetInbox.Status.Claimed), "order history: claimed status");
        assertEq(history[3].timestamp, uint40(block.timestamp), "order history: claimed timestamp");

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
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed), expectedOrderId, "latest claimed order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending), expectedOrderId, "latest pending order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Accepted), expectedOrderId, "latest accepted order id"
        );
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Filled), expectedOrderId, "latest filled order id"
        );

        // Verify token balances after claim (native tokens should be transferred to solver)
        assertEq(user.balance, initialUserBalance - resolvedOrder.minReceived[0].amount, "user balance after");
        assertEq(address(inbox).balance, initialInboxBalance, "inbox balance after");
        assertEq(solver.balance, initialSolverBalance + resolvedOrder.minReceived[0].amount, "solver balance after");
    }

    function test_claim_reverts_order_not_filled() public {
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

        // Try to claim before order is filled
        vm.expectRevert(ISolverNetInbox.OrderNotFilled.selector);
        vm.prank(solver);
        inbox.claim(orderId, solver);
    }

    function test_claim_reverts_not_solver() public {
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

        // Mark order as filled
        bytes32 fillHash = fillHash(orderId, resolvedOrder.fillInstructions[0].originData);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(ISolverNetInbox.markFilled, (orderId, fillHash)),
            100_000
        );

        // Try to claim as someone other than the solver
        vm.expectRevert(Ownable.Unauthorized.selector);
        vm.prank(user);
        inbox.claim(orderId, user);
    }

    function test_claim_reverts_zero_recipient() public {
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

        // Mark order as filled
        bytes32 fillHash = fillHash(orderId, resolvedOrder.fillInstructions[0].originData);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(ISolverNetInbox.markFilled, (orderId, fillHash)),
            100_000
        );

        // Try to claim to address(0)
        vm.expectRevert(ISolverNetInbox.InvalidRecipient.selector);
        vm.prank(solver);
        inbox.claim(orderId, address(0));
    }
}
