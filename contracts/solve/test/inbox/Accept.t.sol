// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_Accept_Test is TestBase {
    function test_accept_reverts() public {
        // order must be accepted by a whitelisted solver
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.accept(bytes32(uint256(1)));

        // order must be opened and pending
        vm.prank(solver);
        vm.expectRevert(ISolverNetInbox.OrderNotPending.selector);
        inbox.accept(bytes32(uint256(1)));

        // prep: open valid order, but warp to after fill deadline
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        order.fillDeadline = uint32(block.timestamp + 1);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        inbox.open{ value: defaultAmount }(order);
        vm.warp(block.timestamp + 2);

        // order must be accepted before the fill deadline
        vm.prank(solver);
        vm.expectRevert(ISolverNetInbox.FillDeadlinePassed.selector);
        inbox.accept(resolvedOrder.orderId);
    }

    function test_accept_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        inbox.open(order);

        vm.prank(solver);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetInbox.Accepted(resolvedOrder.orderId, solver);
        inbox.accept(resolvedOrder.orderId);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Accepted),
            resolvedOrder.orderId,
            "order should be accepted"
        );
    }
}
