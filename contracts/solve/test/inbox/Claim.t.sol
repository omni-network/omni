// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_Claim_Test is TestBase {
    function test_claim_reverts() public {
        // order must be filled
        vm.expectRevert(ISolverNetInbox.OrderNotFilled.selector);
        inbox.claim(bytes32(uint256(1)), address(0));

        // prep: open and fill order
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        inbox.open(order);

        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeWithSelector(ISolverNetInbox.markFilled.selector, resolvedOrder.orderId, fillhash, solver)
        );

        // order must be claimed by the claimant
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.claim(resolvedOrder.orderId, address(0));
    }

    function test_claim_nativeDeposit_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        inbox.open{ value: defaultAmount }(order);

        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeWithSelector(ISolverNetInbox.markFilled.selector, resolvedOrder.orderId, fillhash, solver)
        );

        vm.prank(solver);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetInbox.Claimed(resolvedOrder.orderId, solver, solver);
        inbox.claim(resolvedOrder.orderId, address(solver));

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed),
            resolvedOrder.orderId,
            "order should be claimed"
        );
        assertEq(solver.balance, defaultAmount, "deposit should have been claimed by the solver");
    }

    function test_claim_erc20Deposit_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        inbox.open(order);

        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeWithSelector(ISolverNetInbox.markFilled.selector, resolvedOrder.orderId, fillhash, solver)
        );

        vm.prank(solver);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetInbox.Claimed(resolvedOrder.orderId, solver, solver);
        inbox.claim(resolvedOrder.orderId, address(solver));

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed),
            resolvedOrder.orderId,
            "order should be claimed"
        );
        assertEq(token1.balanceOf(solver), defaultAmount, "deposit should have been claimed by the solver");
    }
}
