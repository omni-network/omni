// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_Cancel_Test is TestBase {
    function test_cancel_reverts() public {
        // order must be pending
        vm.expectRevert(ISolverNetInbox.OrderNotPending.selector);
        inbox.cancel(bytes32(uint256(1)));

        // prep: open a valid order to cancel
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        inbox.open{ value: defaultAmount }(order);

        // order must be cancelled by order owner
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.cancel(resolvedOrder.orderId);
    }

    function test_cancel_nativeDeposit_succeeds() public {
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
        vm.expectEmit(true, true, true, true);
        emit ISolverNetInbox.Reverted(resolvedOrder.orderId);
        inbox.cancel(resolvedOrder.orderId);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Reverted),
            resolvedOrder.orderId,
            "order should be reverted"
        );
        assertEq(user.balance, defaultAmount, "deposit should have been returned to the user");
    }

    function test_cancel_erc20Deposit_succeeds() public {
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
        inbox.cancel(resolvedOrder.orderId);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Reverted),
            resolvedOrder.orderId,
            "order should be reverted"
        );
        assertEq(token1.balanceOf(user), defaultAmount, "deposit should have been returned to the user");
    }
}
