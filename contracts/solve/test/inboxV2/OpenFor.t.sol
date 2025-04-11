// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBaseV2.sol";

contract SolverNet_InboxV2_OpenFor_Test is TestBaseV2 {
    function test_openFor_reverts() public {
        inbox.pauseOpen(true);

        address[] memory expenseTokens = new address[](1);
        expenseTokens[0] = address(token2);
        uint96[] memory expenseAmounts = new uint96[](1);
        expenseAmounts[0] = defaultAmount;

        (SolverNet.OmniOrderData memory orderData, IERC7683.GaslessCrossChainOrder memory order) =
            getArbitraryVaultGaslessOrder(0, address(0), defaultAmount, expenseTokens, expenseAmounts);
        assertTrue(inbox.validateFor(order), "order should be valid");

        bytes memory signature = getPermit2Signature(order, orderData);

        // Should revert when `open` is paused
        vm.expectRevert(ISolverNetInboxV2.IsPaused.selector);
        vm.prank(user);
        inbox.openFor(order, signature, bytes(""));

        // Should revert if `open` and `close` are paused
        inbox.pauseAll(true);
        vm.expectRevert(ISolverNetInboxV2.AllPaused.selector);
        vm.prank(user);
        inbox.openFor(order, signature, bytes(""));

        inbox.pauseAll(false);

        // Should revert if msg.value doesn't match native deposit amount
        vm.expectRevert(ISolverNetInboxV2.InvalidNativeDeposit.selector);
        vm.prank(user);
        inbox.openFor(order, signature, bytes(""));

        // Should revert if order contains more than 32 calls
        SolverNet.Call[] memory originalCalls = orderData.calls;
        orderData.calls = new SolverNet.Call[](33);
        order.orderData = abi.encode(orderData);

        vm.deal(user, defaultAmount);
        vm.expectRevert(ISolverNetInboxV2.InvalidArrayLength.selector);
        vm.prank(user);
        inbox.openFor{ value: defaultAmount }(order, signature, bytes(""));

        // Should revert if order contains more than 32 expenses
        orderData.calls = originalCalls;
        orderData.expenses = new SolverNet.TokenExpense[](33);
        order.orderData = abi.encode(orderData);

        vm.expectRevert(ISolverNetInboxV2.InvalidArrayLength.selector);
        vm.prank(user);
        inbox.openFor{ value: defaultAmount }(order, signature, bytes(""));
    }

    function test_openFor_nativeDeposit_succeeds() public {
        (SolverNet.OmniOrderData memory orderData, IERC7683.GaslessCrossChainOrder memory order) =
            getNativeForNativeVaultGaslessOrder(0, defaultAmount, defaultAmount);
        assertTrue(inbox.validateFor(order), "order should be valid");

        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolveFor(order, bytes(""));
        bytes32 orderId = inbox.getNextGaslessOrderId(user, 0);

        fundUser(orderData.deposit);
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.openFor{ value: defaultAmount }(order, bytes(""), bytes(""));

        // Manually set the open deadline for comparison as we do not store it onchain, it is only validated at open
        (IERC7683.ResolvedCrossChainOrder memory resolved2,, uint248 orderOffset) = inbox.getOrder(orderId);
        resolved2.openDeadline = resolvedOrder.openDeadline;

        assertResolvedEq(resolvedOrder, resolved2);
        assertEq(orderOffset, inbox.getLatestOrderOffset(), "order offset should match contract state");
        assertStatus(orderId, ISolverNetInboxV2.Status.Pending);
        assertEq(address(inbox).balance, defaultAmount, "inbox should have received the deposit");
    }

    function test_openFor_erc20Deposit_succeeds() public {
        (SolverNet.OmniOrderData memory orderData, IERC7683.GaslessCrossChainOrder memory order) =
            getErc20ForErc20VaultGaslessOrder(0, defaultAmount, defaultAmount);
        assertTrue(inbox.validateFor(order), "order should be valid");

        bytes memory signature = getPermit2Signature(order, orderData);

        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolveFor(order, bytes(""));
        bytes32 orderId = inbox.getNextGaslessOrderId(user, 0);

        fundUser(orderData.deposit);
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.openFor(order, signature, bytes(""));

        // Manually set the open deadline for comparison as we do not store it onchain, it is only validated at open
        (IERC7683.ResolvedCrossChainOrder memory resolved2,, uint248 orderOffset) = inbox.getOrder(orderId);
        resolved2.openDeadline = resolvedOrder.openDeadline;

        assertResolvedEq(resolvedOrder, resolved2);
        assertEq(orderOffset, inbox.getLatestOrderOffset(), "order offset should match contract state");
        assertStatus(orderId, ISolverNetInboxV2.Status.Pending);
        assertEq(token1.balanceOf(address(inbox)), defaultAmount, "inbox should have received the deposit");
    }
}
