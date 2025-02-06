// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "./TestBase.sol";

contract SolverNet_E2E_Test is TestBase {
    function test_e2e_nativeDeposit_nativeExpense() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(nativeVault), defaultAmount, user, defaultAmount);

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](0); // No expense for native call value

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getOrder(user, destChainId, defaultFillDeadline, address(0), defaultAmount, calls, expenses);

        assertTrue(inbox.validate(order), "order should be valid");

        fundUser(orderData);

        vm.startPrank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open{ value: defaultAmount }(order);
        vm.stopPrank();

        (IERC7683.ResolvedCrossChainOrder memory resolved2,) = inbox.getOrder(resolvedOrder.orderId);
        assertResolvedEq(resolvedOrder, resolved2);
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending),
            resolvedOrder.orderId,
            "order should be pending"
        );

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        vm.chainId(destChainId);
        vm.startPrank(solver);
        inbox.accept(resolvedOrder.orderId);
        outbox.fill{ value: defaultAmount + fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );
        vm.stopPrank();

        vm.chainId(srcChainId);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(SolverNetInbox.markFilled, (resolvedOrder.orderId, fillhash, solver)),
            getGasLimit(resolvedOrder.fillInstructions[0].originData)
        );

        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed),
            resolvedOrder.orderId,
            "order should be claimed"
        );
        assertEq(
            nativeVault.balances(user), defaultAmount, "user should have received the native expense as a vault deposit"
        );
        assertEq(address(nativeVault).balance, defaultAmount, "native vault should have received the native deposit");
        assertEq(solver.balance, defaultAmount, "solver should have received the native deposit as their reward");
    }

    function test_e2e_nativeDeposit_erc20Expense() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(erc20Vault), 0, user, defaultAmount);

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](1);
        expenses[0] = getExpense(address(erc20Vault), address(token2), defaultAmount);

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getOrder(user, destChainId, defaultFillDeadline, address(0), defaultAmount, calls, expenses);

        assertTrue(inbox.validate(order), "order should be valid");

        fundUser(orderData);

        vm.startPrank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open{ value: defaultAmount }(order);
        vm.stopPrank();

        (IERC7683.ResolvedCrossChainOrder memory resolved2,) = inbox.getOrder(resolvedOrder.orderId);
        assertResolvedEq(resolvedOrder, resolved2);
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending),
            resolvedOrder.orderId,
            "order should be pending"
        );

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        vm.chainId(destChainId);
        vm.startPrank(solver);
        inbox.accept(resolvedOrder.orderId);
        outbox.fill{ value: fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );
        vm.stopPrank();

        vm.chainId(srcChainId);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(SolverNetInbox.markFilled, (resolvedOrder.orderId, fillhash, solver)),
            getGasLimit(resolvedOrder.fillInstructions[0].originData)
        );

        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed),
            resolvedOrder.orderId,
            "order should be claimed"
        );
        assertEq(
            erc20Vault.balances(user), defaultAmount, "user should have received the erc20 expense as a vault deposit"
        );
        assertEq(
            token2.balanceOf(address(erc20Vault)), defaultAmount, "erc20 vault should have received the erc20 deposit"
        );
        assertEq(solver.balance, defaultAmount, "solver should have received the native deposit as their reward");
    }

    function test_e2e_erc20Deposit_nativeExpense() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(nativeVault), defaultAmount, user, defaultAmount);

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](0); // No expense for native call value

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getOrder(user, destChainId, defaultFillDeadline, address(token1), defaultAmount, calls, expenses);

        assertTrue(inbox.validate(order), "order should be valid");

        fundUser(orderData);

        vm.startPrank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open(order);
        vm.stopPrank();

        (IERC7683.ResolvedCrossChainOrder memory resolved2,) = inbox.getOrder(resolvedOrder.orderId);
        assertResolvedEq(resolvedOrder, resolved2);
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending),
            resolvedOrder.orderId,
            "order should be pending"
        );

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        vm.chainId(destChainId);
        vm.startPrank(solver);
        inbox.accept(resolvedOrder.orderId);
        outbox.fill{ value: defaultAmount + fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );
        vm.stopPrank();

        vm.chainId(srcChainId);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(SolverNetInbox.markFilled, (resolvedOrder.orderId, fillhash, solver)),
            getGasLimit(resolvedOrder.fillInstructions[0].originData)
        );

        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed),
            resolvedOrder.orderId,
            "order should be claimed"
        );
        assertEq(
            nativeVault.balances(user), defaultAmount, "user should have received the native expense as a vault deposit"
        );
        assertEq(address(nativeVault).balance, defaultAmount, "native vault should have received the native deposit");
        assertEq(
            token1.balanceOf(solver), defaultAmount, "solver should have received the erc20 deposit as their reward"
        );
    }

    function test_e2e_erc20Deposit_erc20Expense() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(erc20Vault), 0, user, defaultAmount);

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](1);
        expenses[0] = getExpense(address(erc20Vault), address(token2), defaultAmount);

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getOrder(user, destChainId, defaultFillDeadline, address(token1), defaultAmount, calls, expenses);

        assertTrue(inbox.validate(order), "order should be valid");

        fundUser(orderData);

        vm.startPrank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open(order);
        vm.stopPrank();

        (IERC7683.ResolvedCrossChainOrder memory resolved2,) = inbox.getOrder(resolvedOrder.orderId);
        assertResolvedEq(resolvedOrder, resolved2);
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending),
            resolvedOrder.orderId,
            "order should be pending"
        );

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        vm.chainId(destChainId);
        vm.startPrank(solver);
        inbox.accept(resolvedOrder.orderId);
        outbox.fill{ value: fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );
        vm.stopPrank();

        vm.chainId(srcChainId);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(SolverNetInbox.markFilled, (resolvedOrder.orderId, fillhash, solver)),
            getGasLimit(resolvedOrder.fillInstructions[0].originData)
        );

        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed),
            resolvedOrder.orderId,
            "order should be claimed"
        );
        assertEq(
            erc20Vault.balances(user), defaultAmount, "user should have received the erc20 expense as a vault deposit"
        );
        assertEq(
            token2.balanceOf(address(erc20Vault)), defaultAmount, "erc20 vault should have received the erc20 deposit"
        );
        assertEq(
            token1.balanceOf(solver), defaultAmount, "solver should have received the erc20 deposit as their reward"
        );
    }

    function test_e2e_nativeDeposit_mixedExpenses_multicall() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](2);
        calls[0] = getVaultCall(address(nativeVault), defaultAmount, user, defaultAmount);
        calls[1] = getVaultCall(address(erc20Vault), 0, user, defaultAmount);

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](1);
        expenses[0] = getExpense(address(erc20Vault), address(token2), defaultAmount);

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getOrder(user, destChainId, defaultFillDeadline, address(0), defaultAmount * 2, calls, expenses);

        assertTrue(inbox.validate(order), "order should be valid");

        fundUser(orderData);

        vm.startPrank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open{ value: defaultAmount * 2 }(order);
        vm.stopPrank();

        (IERC7683.ResolvedCrossChainOrder memory resolved2,) = inbox.getOrder(resolvedOrder.orderId);
        assertResolvedEq(resolvedOrder, resolved2);
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending),
            resolvedOrder.orderId,
            "order should be pending"
        );

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        vm.chainId(destChainId);
        vm.startPrank(solver);
        inbox.accept(resolvedOrder.orderId);
        outbox.fill{ value: defaultAmount + fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );
        vm.stopPrank();

        vm.chainId(srcChainId);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(SolverNetInbox.markFilled, (resolvedOrder.orderId, fillhash, solver)),
            getGasLimit(resolvedOrder.fillInstructions[0].originData)
        );

        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed),
            resolvedOrder.orderId,
            "order should be claimed"
        );
        assertEq(
            nativeVault.balances(user), defaultAmount, "user should have received the native expense as a vault deposit"
        );
        assertEq(
            erc20Vault.balances(user), defaultAmount, "user should have received the erc20 expense as a vault deposit"
        );
        assertEq(address(nativeVault).balance, defaultAmount, "native vault should have received the native deposit");
        assertEq(
            token2.balanceOf(address(erc20Vault)), defaultAmount, "erc20 vault should have received the erc20 deposit"
        );
        assertEq(solver.balance, defaultAmount * 2, "solver should have received the native deposit as their reward");
    }
}
