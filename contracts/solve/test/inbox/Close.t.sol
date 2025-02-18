// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_Close_Test is TestBase {
    function test_close_reverts() public {
        inbox.pauseClose(true);

        // should revert when `close` is paused
        vm.expectRevert(ISolverNetInbox.IsPaused.selector);
        vm.prank(user);
        inbox.close(bytes32(uint256(1)));

        // should revert if `open` and `close` are paused
        inbox.pauseAll(true);
        vm.expectRevert(ISolverNetInbox.AllPaused.selector);
        vm.prank(user);
        inbox.close(bytes32(uint256(1)));

        inbox.pauseAll(false);

        // order must be pending
        vm.expectRevert(ISolverNetInbox.OrderNotPending.selector);
        inbox.close(bytes32(uint256(1)));

        // prep: open a valid order to close
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        inbox.open{ value: defaultAmount }(order);

        // order must be closed by order owner
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.close(resolvedOrder.orderId);

        // order can only be closed after fill deadline has elapsed
        vm.prank(user);
        vm.expectRevert(ISolverNetInbox.OrderStillValid.selector);
        inbox.close(resolvedOrder.orderId);
    }

    function test_close_nativeDeposit_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open{ value: defaultAmount }(order);

        vm.prank(user);
        vm.warp(defaultFillDeadline + defaultFillBuffer + 1);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetInbox.Closed(resolvedOrder.orderId);
        inbox.close(resolvedOrder.orderId);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Closed),
            resolvedOrder.orderId,
            "order should be closed"
        );
        assertEq(user.balance, defaultAmount, "deposit should have been returned to the user");
    }

    function test_close_erc20Deposit_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open(order);

        vm.prank(user);
        vm.warp(defaultFillDeadline + defaultFillBuffer + 1);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetInbox.Closed(resolvedOrder.orderId);
        inbox.close(resolvedOrder.orderId);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Closed),
            resolvedOrder.orderId,
            "order should be closed"
        );
        assertEq(token1.balanceOf(user), defaultAmount, "deposit should have been returned to the user");
    }
}
