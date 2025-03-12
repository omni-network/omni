// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_Reject_Test is TestBase {
    function test_reject_reverts() public {
        // order must be rejected by a whitelisted solver
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.reject(bytes32(uint256(1)), 0);

        // order rejection reason must be non-zero
        vm.prank(solver);
        vm.expectRevert(ISolverNetInbox.InvalidReason.selector);
        inbox.reject(bytes32(uint256(1)), 0);

        // order must at least be in pending state
        vm.prank(solver);
        vm.expectRevert(ISolverNetInbox.OrderNotPending.selector);
        inbox.reject(bytes32(uint256(1)), 1);
    }

    function test_reject_nativeDeposit_succeeds() public {
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

        vm.prank(solver);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetInbox.Rejected(resolvedOrder.orderId, solver, 1);
        inbox.reject(resolvedOrder.orderId, 1);

        (, ISolverNetInbox.OrderState memory state,) = inbox.getOrder(resolvedOrder.orderId);

        assertEq(state.rejectReason, 1, "reject reason should be set");
        assertStatus(resolvedOrder.orderId, ISolverNetInbox.Status.Rejected);
        assertEq(user.balance, defaultAmount, "deposit should have been returned to the user");
    }

    function test_reject_erc20Deposit_succeeds() public {
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

        vm.prank(solver);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetInbox.Rejected(resolvedOrder.orderId, solver, 1);
        inbox.reject(resolvedOrder.orderId, 1);

        (, ISolverNetInbox.OrderState memory state,) = inbox.getOrder(resolvedOrder.orderId);

        assertEq(state.rejectReason, 1, "reject reason should be set");
        assertStatus(resolvedOrder.orderId, ISolverNetInbox.Status.Rejected);
        assertEq(token1.balanceOf(user), defaultAmount, "deposit should have been returned to the user");
    }
}
