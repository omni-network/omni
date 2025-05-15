// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_Open_Test is TestBase {
    function test_open_reverts() public {
        inbox.pauseOpen(true);

        address[] memory expenseTokens = new address[](1);
        expenseTokens[0] = address(token2);
        uint96[] memory expenseAmounts = new uint96[](1);
        expenseAmounts[0] = defaultAmount;

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getArbitraryVaultOrder(address(0), defaultAmount, expenseTokens, expenseAmounts);
        assertTrue(inbox.validate(order), "order should be valid");

        // Should revert when `open` is paused
        vm.expectRevert(ISolverNetInbox.IsPaused.selector);
        vm.prank(user);
        inbox.open(order);

        // Should revert if `open` and `close` are paused
        inbox.pauseAll(true);
        vm.expectRevert(ISolverNetInbox.AllPaused.selector);
        vm.prank(user);
        inbox.open(order);

        inbox.pauseAll(false);

        // Should revert if msg.value doesn't match native deposit amount
        vm.expectRevert(ISolverNetInbox.InvalidNativeDeposit.selector);
        vm.prank(user);
        inbox.open(order);

        // Should revert if order contains more than 32 calls
        SolverNet.Call[] memory originalCalls = orderData.calls;
        SolverNet.Call[] memory calls = new SolverNet.Call[](33);
        orderData.calls = calls;
        order.orderData = abi.encode(orderData);

        vm.deal(user, defaultAmount);
        vm.expectRevert(ISolverNetInbox.InvalidArrayLength.selector);
        vm.prank(user);
        inbox.open{ value: defaultAmount }(order);

        // Should revert if order contains more than 32 expenses
        orderData.calls = originalCalls;
        SolverNet.TokenExpense[] memory originalExpenses = orderData.expenses;
        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](33);
        orderData.expenses = expenses;
        order.orderData = abi.encode(orderData);

        vm.expectRevert(ISolverNetInbox.InvalidArrayLength.selector);
        vm.prank(user);
        inbox.open{ value: defaultAmount }(order);

        // Should revert if less tokens are received than expected due to max transfer balance override
        orderData.expenses = originalExpenses;
        orderData.deposit = SolverNet.Deposit({ token: address(maxTransferToken), amount: type(uint96).max });
        order.orderData = abi.encode(orderData);

        maxTransferToken.mint(user, 1 ether);
        vm.startPrank(user);
        maxTransferToken.approve(address(inbox), type(uint256).max);
        vm.expectRevert(ISolverNetInbox.InvalidERC20Deposit.selector);
        inbox.open(order);
        vm.stopPrank();

        // Should revert if less tokens are received than expected due to fee on transfer
        orderData.deposit = SolverNet.Deposit({ token: address(feeOnTransferToken), amount: 1 ether });
        order.orderData = abi.encode(orderData);

        feeOnTransferToken.mint(user, 1 ether);
        vm.startPrank(user);
        feeOnTransferToken.approve(address(inbox), type(uint256).max);
        vm.expectRevert(ISolverNetInbox.InvalidERC20Deposit.selector);
        inbox.open(order);
        vm.stopPrank();
    }

    function test_open_nativeDeposit_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        bytes32 orderId = inbox.getOrderId(user, inbox.getUserNonce(user));

        fundUser(orderData);
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open{ value: defaultAmount }(order);

        (IERC7683.ResolvedCrossChainOrder memory resolved2,, uint248 orderOffset) = inbox.getOrder(orderId);
        assertResolvedEq(resolvedOrder, resolved2);
        assertEq(orderOffset, inbox.getLatestOrderOffset(), "order offset should match contract state");
        assertStatus(orderId, ISolverNetInbox.Status.Pending);
        assertEq(address(inbox).balance, defaultAmount, "inbox should have received the deposit");
    }

    function test_open_erc20Deposit_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        bytes32 orderId = inbox.getOrderId(user, inbox.getUserNonce(user));

        fundUser(orderData);
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.open(order);

        (IERC7683.ResolvedCrossChainOrder memory resolved2,, uint248 orderOffset) = inbox.getOrder(orderId);
        assertResolvedEq(resolvedOrder, resolved2);
        assertEq(orderOffset, inbox.getLatestOrderOffset(), "order offset should match contract state");
        assertStatus(orderId, ISolverNetInbox.Status.Pending);
        assertEq(token1.balanceOf(address(inbox)), defaultAmount, "inbox should have received the deposit");
    }

    function test_open_hyperlane() public {
        address impl = address(new SolverNetInbox(address(0), address(mailboxes[uint32(srcChainId)])));
        inbox = SolverNetInbox(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
        inbox.initialize(address(this), solver);
        setRoutes(ISolverNetOutbox.Provider.Hyperlane);

        uint256 snapshot = vm.snapshot();
        test_open_reverts();
        vm.revertTo(snapshot);

        test_open_nativeDeposit_succeeds();
        vm.revertTo(snapshot);

        test_open_erc20Deposit_succeeds();
    }
}
