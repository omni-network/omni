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

        assertStatus(resolvedOrder.orderId, ISolverNetInbox.Status.Closed);
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

        assertStatus(resolvedOrder.orderId, ISolverNetInbox.Status.Closed);
        assertEq(token1.balanceOf(user), defaultAmount, "deposit should have been returned to the user");
    }

    function test_close_localOrder_succeeds() public {
        vm.chainId(srcChainId);
        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(nativeVault), defaultAmount, user, defaultAmount);

        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](0);

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getOrder(user, srcChainId, uint32(block.timestamp + 1), address(0), defaultAmount, calls, expenses);

        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open{ value: defaultAmount }(order);

        vm.prank(user);
        vm.warp(block.timestamp + 2);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetInbox.Closed(resolvedOrder.orderId);
        inbox.close(resolvedOrder.orderId);

        assertStatus(resolvedOrder.orderId, ISolverNetInbox.Status.Closed);
        assertEq(user.balance, defaultAmount, "deposit should have been returned to the user");
    }

    function test_close_hyperlane() public {
        address impl = address(new SolverNetInbox(address(0), address(mailboxes[uint32(srcChainId)])));
        inbox = SolverNetInbox(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
        inbox.initialize(address(this), solver);
        setRoutes(ISolverNetOutbox.Provider.Hyperlane);

        uint256 snapshot = vm.snapshotState();
        test_close_reverts();
        vm.revertToState(snapshot);

        test_close_nativeDeposit_succeeds();
        vm.revertToState(snapshot);

        test_close_erc20Deposit_succeeds();
        vm.revertToState(snapshot);

        test_close_localOrder_succeeds();
    }
}
