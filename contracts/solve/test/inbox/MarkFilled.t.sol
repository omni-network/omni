// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_MarkFilled_Test is TestBase {
    function test_markFilled_reverts() public {
        // order must be pending
        bytes32 orderId = inbox.getOrderId(user, inbox.getUserNonce(user));
        vm.expectRevert(ISolverNetInbox.OrderNotPending.selector);
        inbox.markFilled(orderId, bytes32(0), address(0));

        // prep: open a valid order
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");

        fundUser(orderData);
        vm.prank(user);
        inbox.open{ value: defaultAmount }(order);

        // call must come from the Omni portal
        vm.expectRevert(ISolverNetInbox.WrongSourceChain.selector);
        inbox.markFilled(orderId, bytes32(0), address(0));

        // order must be filled on the correct chain
        vm.expectRevert(ISolverNetInbox.WrongSourceChain.selector);
        portal.mockXCall(
            destChainId + 1,
            address(outbox),
            address(inbox),
            abi.encodeWithSelector(ISolverNetInbox.markFilled.selector, orderId, bytes32(0), address(0))
        );

        // order must be filled by the outbox
        vm.expectRevert(Ownable.Unauthorized.selector);
        portal.mockXCall(
            destChainId,
            user,
            address(inbox),
            abi.encodeWithSelector(ISolverNetInbox.markFilled.selector, orderId, bytes32(0), address(0))
        );

        // order must have a matching fill hash
        vm.expectRevert(ISolverNetInbox.WrongFillHash.selector);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeWithSelector(ISolverNetInbox.markFilled.selector, orderId, bytes32(0), address(0))
        );
    }

    function test_markFilled_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        inbox.open(order);

        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetInbox.Filled(resolvedOrder.orderId, fillhash, solver);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeWithSelector(ISolverNetInbox.markFilled.selector, resolvedOrder.orderId, fillhash, solver)
        );

        assertStatus(resolvedOrder.orderId, ISolverNetInbox.Status.Filled);
    }
}
