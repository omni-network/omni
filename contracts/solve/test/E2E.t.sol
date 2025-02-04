// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "./TestBase.sol";

contract SolverNet_E2E_Test is TestBase {
    function test_e2e_nativeDeposit_nativeExpense() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(nativeVault), defaultAmount, user, defaultAmount);

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](0); // No expense for native call value

        (, IERC7683.OnchainCrossChainOrder memory order) =
            getOrder(user, destChainId, defaultFillDeadline, address(0), defaultAmount, calls, expenses);

        assertTrue(inbox.validate(order), "order should be valid");

        vm.deal(user, defaultAmount);
        vm.startPrank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        vm.expectEmit();
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open{ value: defaultAmount }(order);
        vm.stopPrank();

        // check get order
        (IERC7683.ResolvedCrossChainOrder memory resolved2,) = inbox.getOrder(resolvedOrder.orderId);
        assertEq(resolvedOrder.orderId, resolved2.orderId);

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.chainId(destChainId);
        vm.deal(solver, defaultAmount + fillFee);
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
    }

    function test_e2e_nativeDeposit_erc20Expense() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(erc20Vault), 0, user, defaultAmount);

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](1);
        expenses[0] = getExpense(address(erc20Vault), address(token2), defaultAmount);

        (, IERC7683.OnchainCrossChainOrder memory order) =
            getOrder(user, destChainId, defaultFillDeadline, address(0), defaultAmount, calls, expenses);

        assertTrue(inbox.validate(order), "order should be valid");

        vm.deal(user, defaultAmount);
        vm.startPrank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        inbox.open{ value: defaultAmount }(order);
        vm.stopPrank();

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.chainId(destChainId);
        vm.deal(solver, fillFee);
        vm.startPrank(solver);
        token2.mint(solver, defaultAmount);
        token2.approve(address(outbox), defaultAmount);
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
    }

    function test_e2e_erc20Deposit_nativeExpense() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(nativeVault), defaultAmount, user, defaultAmount);

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](0); // No expense for native call value

        (, IERC7683.OnchainCrossChainOrder memory order) =
            getOrder(user, destChainId, defaultFillDeadline, address(token1), defaultAmount, calls, expenses);

        assertTrue(inbox.validate(order), "order should be valid");

        vm.startPrank(user);
        token1.mint(user, defaultAmount);
        token1.approve(address(inbox), defaultAmount);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        inbox.open(order);
        vm.stopPrank();

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.chainId(destChainId);
        vm.deal(solver, defaultAmount + fillFee);
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
    }

    function test_e2e_erc20Deposit_erc20Expense() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(erc20Vault), 0, user, defaultAmount);

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](1);
        expenses[0] = getExpense(address(erc20Vault), address(token2), defaultAmount);

        (, IERC7683.OnchainCrossChainOrder memory order) =
            getOrder(user, destChainId, defaultFillDeadline, address(token1), defaultAmount, calls, expenses);

        assertTrue(inbox.validate(order), "order should be valid");

        vm.startPrank(user);
        token1.mint(user, defaultAmount);
        token1.approve(address(inbox), defaultAmount);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        inbox.open(order);
        vm.stopPrank();

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.chainId(destChainId);
        vm.deal(solver, fillFee);
        vm.startPrank(solver);
        token2.mint(solver, defaultAmount);
        token2.approve(address(outbox), defaultAmount);
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
    }

    function test_e2e_nativeDeposit_mixedExpenses_multicall() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](2);
        calls[0] = getVaultCall(address(nativeVault), defaultAmount, user, defaultAmount);
        calls[1] = getVaultCall(address(erc20Vault), 0, user, defaultAmount);

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](1);
        expenses[0] = getExpense(address(erc20Vault), address(token2), defaultAmount);

        (, IERC7683.OnchainCrossChainOrder memory order) =
            getOrder(user, destChainId, defaultFillDeadline, address(0), defaultAmount * 2, calls, expenses);

        assertTrue(inbox.validate(order), "order should be valid");

        vm.deal(user, defaultAmount * 2);
        vm.startPrank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        inbox.open{ value: defaultAmount * 2 }(order);
        vm.stopPrank();

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.chainId(destChainId);
        vm.deal(solver, defaultAmount + fillFee);
        vm.startPrank(solver);
        token2.mint(solver, defaultAmount);
        token2.approve(address(outbox), defaultAmount);
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
    }
}
