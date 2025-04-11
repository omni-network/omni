// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBaseV2.sol";

contract SolverNet_InboxV2_Open_Test is TestBaseV2 {
    function test_open_reverts() public {
        inbox.pauseOpen(true);

        address[] memory expenseTokens = new address[](1);
        expenseTokens[0] = address(token2);
        uint96[] memory expenseAmounts = new uint96[](1);
        expenseAmounts[0] = defaultAmount;

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getArbitraryVaultOnchainOrder(address(0), defaultAmount, expenseTokens, expenseAmounts);
        assertTrue(inbox.validate(order), "order should be valid");

        // Should revert when `open` is paused
        vm.expectRevert(ISolverNetInboxV2.IsPaused.selector);
        vm.prank(user);
        inbox.open(order);

        // Should revert if `open` and `close` are paused
        inbox.pauseAll(true);
        vm.expectRevert(ISolverNetInboxV2.AllPaused.selector);
        vm.prank(user);
        inbox.open(order);

        inbox.pauseAll(false);

        // Should revert if msg.value doesn't match native deposit amount
        vm.expectRevert(ISolverNetInboxV2.InvalidNativeDeposit.selector);
        vm.prank(user);
        inbox.open(order);

        // Should revert if order contains more than 32 calls
        SolverNet.Call[] memory originalCalls = orderData.calls;
        orderData.calls = new SolverNet.Call[](33);
        order.orderData = abi.encode(orderData);

        vm.deal(user, defaultAmount);
        vm.expectRevert(ISolverNetInboxV2.InvalidArrayLength.selector);
        vm.prank(user);
        inbox.open{ value: defaultAmount }(order);

        // Should revert if order contains more than 32 expenses
        orderData.calls = originalCalls;
        orderData.expenses = new SolverNet.TokenExpense[](33);
        order.orderData = abi.encode(orderData);

        vm.expectRevert(ISolverNetInboxV2.InvalidArrayLength.selector);
        vm.prank(user);
        inbox.open{ value: defaultAmount }(order);
    }

    function test_open_nativeDeposit_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOnchainOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        bytes32 orderId = inbox.getNextOnchainOrderId(user);

        fundUser(orderData);
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open{ value: defaultAmount }(order);

        (IERC7683.ResolvedCrossChainOrder memory resolved2,, uint248 orderOffset) = inbox.getOrder(orderId);
        assertResolvedEq(resolvedOrder, resolved2);
        assertEq(orderOffset, inbox.getLatestOrderOffset(), "order offset should match contract state");
        assertStatus(orderId, ISolverNetInboxV2.Status.Pending);
        assertEq(address(inbox).balance, defaultAmount, "inbox should have received the deposit");
    }

    function test_open_erc20Deposit_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOnchainOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        bytes32 orderId = inbox.getNextOnchainOrderId(user);

        fundUser(orderData);
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open(order);

        (IERC7683.ResolvedCrossChainOrder memory resolved2,, uint248 orderOffset) = inbox.getOrder(orderId);
        assertResolvedEq(resolvedOrder, resolved2);
        assertEq(orderOffset, inbox.getLatestOrderOffset(), "order offset should match contract state");
        assertStatus(orderId, ISolverNetInboxV2.Status.Pending);
        assertEq(token1.balanceOf(address(inbox)), defaultAmount, "inbox should have received the deposit");
    }
}
