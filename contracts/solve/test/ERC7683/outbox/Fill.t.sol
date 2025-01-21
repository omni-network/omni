// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { Refunder } from "test/utils/Refunder.sol";

contract SolverNet_Outbox_Fill_Test is TestBase {
    using AddrUtils for address;
    using AddrUtils for bytes32;

    Refunder internal refunder;

    function setUp() public override {
        super.setUp();
        vm.chainId(destChainId);

        refunder = new Refunder();
    }

    function test_fill_succeeds() public {
        // Create and resolve an order on source chain
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        fundSolver(resolvedOrder.maxSpent);
        (resolvedOrder.minReceived, resolvedOrder.maxSpent);

        bytes32 orderId = inbox.getNextId();

        // Calculate the fill fee
        uint256 fillFee = outbox.fillFee(srcChainId);
        vm.deal(solver, fillFee);

        // Store initial balances
        uint256 initialOutboxNativeBalance = address(outbox).balance;
        uint256 initialExecutorNativeBalance = address(outbox.executor()).balance;

        // Store initial token balances
        address[] memory tokens = new address[](resolvedOrder.maxSpent.length);
        uint256[] memory solverBalances = new uint256[](resolvedOrder.maxSpent.length);
        uint256[] memory outboxBalances = new uint256[](resolvedOrder.maxSpent.length);
        uint256[] memory executorBalances = new uint256[](resolvedOrder.maxSpent.length);
        uint256[] memory vaultBalances = new uint256[](resolvedOrder.maxSpent.length);

        for (uint256 i = 0; i < resolvedOrder.maxSpent.length; i++) {
            IERC7683.Output memory expense = resolvedOrder.maxSpent[i];
            address token = expense.token.toAddress();

            tokens[i] = token;
            solverBalances[i] = MockToken(token).balanceOf(solver);
            outboxBalances[i] = MockToken(token).balanceOf(address(outbox));
            executorBalances[i] = MockToken(token).balanceOf(address(outbox.executor()));
            vaultBalances[i] = MockToken(token).balanceOf(address(vault));
        }

        // Get the fill data from the order
        ISolverNet.FillOriginData memory fillData =
            abi.decode(resolvedOrder.fillInstructions[0].originData, (ISolverNet.FillOriginData));

        // Expect the Filled event
        vm.expectEmit(true, true, true, true, address(outbox));
        emit ISolverNetOutbox.Filled(orderId, fillHash(orderId, resolvedOrder.fillInstructions[0].originData), solver);

        // Fill the order as solver
        vm.prank(solver);
        outbox.fill{ value: fillFee }(orderId, resolvedOrder.fillInstructions[0].originData, "");

        // Verify the order is marked as filled
        assertTrue(
            outbox.didFill(orderId, resolvedOrder.fillInstructions[0].originData), "order should be marked as filled"
        );

        // Verify native token balances
        assertEq(address(outbox).balance, initialOutboxNativeBalance, "outbox native balance after");
        assertEq(address(outbox.executor()).balance, initialExecutorNativeBalance, "executor native balance after");

        // Verify token balances after fill
        for (uint256 i = 0; i < resolvedOrder.maxSpent.length; i++) {
            IERC7683.Output memory expense = resolvedOrder.maxSpent[i];
            address token = tokens[i];

            // Verify token balances
            assertEq(
                MockToken(token).balanceOf(solver), solverBalances[i] - expense.amount, "solver token balance after"
            );
            assertEq(MockToken(token).balanceOf(address(outbox)), outboxBalances[i], "outbox token balance after");
            assertEq(
                MockToken(token).balanceOf(address(outbox.executor())),
                executorBalances[i],
                "executor token balance after"
            );
            assertEq(
                MockToken(token).balanceOf(address(vault)),
                vaultBalances[i] + expense.amount,
                "vault token balance after"
            );
            assertEq(vault.balances(user), expense.amount, "vault user balance after");

            // Verify allowances were reset for each spender in fillData
            for (uint256 j = 0; j < fillData.call.expenses.length; j++) {
                ISolverNet.TokenExpense memory tokenExpense = fillData.call.expenses[j];
                if (tokenExpense.token.toAddress() == token) {
                    assertEq(
                        MockToken(token).allowance(address(outbox.executor()), tokenExpense.spender.toAddress()),
                        0,
                        "executor allowance after"
                    );
                }
            }
        }
    }

    function test_fill_call_refund() public {
        // Create and resolve an order on source chain
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        // Tamper with order to send ETH to Refunder
        ISolverNet.FillOriginData memory fillData =
            abi.decode(resolvedOrder.fillInstructions[0].originData, (ISolverNet.FillOriginData));
        fillData.call.expenses = new ISolverNet.TokenExpense[](0);
        fillData.call.target = address(refunder).toBytes32();
        fillData.call.value = 1 wei;

        bytes32 orderId = inbox.getNextId();

        // Calculate the fill fee
        uint256 fillFee = outbox.fillFee(srcChainId);
        vm.deal(solver, fillFee + 1 wei);

        uint256 solverBalance = solver.balance;

        // Fill the order as solver and expect refund
        vm.prank(solver);
        outbox.fill{ value: fillFee + 1 wei }(orderId, abi.encode(fillData), "");

        // Verify the order is marked as filled
        assertTrue(outbox.didFill(orderId, abi.encode(fillData)), "order should be marked as filled");

        // Verify native token balances
        assertEq(address(outbox).balance, 0, "outbox native balance after");
        assertEq(address(outbox.executor()).balance, 0, "executor native balance after");
        assertEq(solver.balance, solverBalance - fillFee, "solver balance after");
    }

    function test_fill_overpayment_refund() public {
        // Create and resolve an order on source chain
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        fundSolver(resolvedOrder.maxSpent);
        (resolvedOrder.minReceived, resolvedOrder.maxSpent);

        bytes32 orderId = inbox.getNextId();

        // Calculate the fill fee and add extra ETH
        uint256 fillFee = outbox.fillFee(srcChainId);
        uint256 extraEth = 0.1 ether;
        vm.deal(solver, fillFee + extraEth);

        uint256 initialSolverBalance = solver.balance;

        // Fill with overpayment
        vm.prank(solver);
        outbox.fill{ value: fillFee + extraEth }(orderId, resolvedOrder.fillInstructions[0].originData, "");

        // Verify balances after fill
        assertEq(address(outbox).balance, 0, "outbox should have no ETH");
        assertEq(address(outbox.executor()).balance, 0, "executor should have no ETH");
        assertEq(solver.balance, initialSolverBalance - fillFee, "solver should be refunded excess ETH");
    }

    function test_fill_reverts_insufficient_fee() public {
        // Create and resolve an order on source chain
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        fundSolver(resolvedOrder.maxSpent);
        (resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Tamper with order to ensure insufficient ETH is provided
        ISolverNet.FillOriginData memory fillData =
            abi.decode(resolvedOrder.fillInstructions[0].originData, (ISolverNet.FillOriginData));
        fillData.call.target = bytes32(0);
        fillData.call.value = 1 wei;
        vm.deal(address(outbox), 1 wei);

        bytes32 orderId = inbox.getNextId();

        // Calculate the fill fee
        uint256 fillFee = outbox.fillFee(srcChainId);
        vm.deal(solver, fillFee);

        // Try to fill with insufficient fee
        vm.prank(solver);
        vm.expectRevert(ISolverNetOutbox.InsufficientFee.selector);
        outbox.fill{ value: fillFee }(orderId, abi.encode(fillData), "");
    }

    function test_fill_reverts_already_filled() public {
        // Create and resolve an order on source chain
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        resolvedOrder.maxSpent[0].amount = resolvedOrder.maxSpent[0].amount * 2;
        fundSolver(resolvedOrder.maxSpent);
        (resolvedOrder.minReceived, resolvedOrder.maxSpent);

        bytes32 orderId = inbox.getNextId();

        // Calculate the fill fee
        uint256 fillFee = outbox.fillFee(srcChainId);
        vm.deal(solver, fillFee * 2);

        // Fill the order first time
        vm.prank(solver);
        outbox.fill{ value: fillFee }(orderId, resolvedOrder.fillInstructions[0].originData, "");

        // Try to fill again
        vm.prank(solver);
        vm.expectRevert(ISolverNetOutbox.AlreadyFilled.selector);
        outbox.fill{ value: fillFee }(orderId, resolvedOrder.fillInstructions[0].originData, "");
    }

    function test_fill_reverts_invalid_origin_data() public {
        bytes32 orderId = inbox.getNextId();

        // Calculate the fill fee
        uint256 fillFee = outbox.fillFee(srcChainId);

        // Try to fill with invalid origin data
        vm.deal(solver, fillFee);
        vm.prank(solver);
        vm.expectRevert(); // Will revert on decoding invalid origin data
        outbox.fill{ value: fillFee }(orderId, "invalid origin data", "");
    }

    function test_fill_reverts_wrong_chain() public {
        // Create and resolve an order on source chain
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        fundSolver(resolvedOrder.maxSpent);
        (resolvedOrder.minReceived, resolvedOrder.maxSpent);

        bytes32 orderId = inbox.getNextId();

        // Switch to wrong destination chain
        vm.chainId(destChainId + 1);

        // Calculate the fill fee
        uint256 fillFee = outbox.fillFee(srcChainId);
        vm.deal(solver, fillFee);

        // Try to fill on wrong chain
        vm.prank(solver);
        vm.expectRevert(ISolverNetOutbox.WrongDestChain.selector);
        outbox.fill{ value: fillFee }(orderId, resolvedOrder.fillInstructions[0].originData, "");
    }
}
