// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_Open_Test is TestBase {
    function test_v2_open_reverts() public {
        inbox.pauseOpen(true);

        address[] memory expenseTokens = new address[](1);
        expenseTokens[0] = address(token2);
        uint96[] memory expenseAmounts = new uint96[](1);
        expenseAmounts[0] = defaultAmount;

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getArbitraryVaultOrder(address(0), defaultAmount, expenseTokens, expenseAmounts);
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

        // Should revert if less tokens are received than expected due to max transfer balance override
        orderData.deposit = SolverNet.Deposit({ token: address(maxTransferToken), amount: type(uint96).max });
        order.orderData = abi.encode(orderData);

        maxTransferToken.mint(user, 1 ether);
        vm.startPrank(user);
        maxTransferToken.approve(address(inbox), type(uint256).max);
        vm.expectRevert(ISolverNetInboxV2.InvalidERC20Deposit.selector);
        inbox.open(order);
        vm.stopPrank();

        // Should revert if less tokens are received than expected due to fee on transfer
        orderData.deposit = SolverNet.Deposit({ token: address(feeOnTransferToken), amount: 1 ether });
        order.orderData = abi.encode(orderData);

        feeOnTransferToken.mint(user, 1 ether);
        vm.startPrank(user);
        feeOnTransferToken.approve(address(inbox), type(uint256).max);
        vm.expectRevert(ISolverNetInboxV2.InvalidERC20Deposit.selector);
        inbox.open(order);
        vm.stopPrank();
    }

    function test_v2_open_nativeDeposit_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
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

    function test_v2_open_erc20Deposit_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
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

    function test_v2_open_hyperlane() public {
        address impl = address(new SolverNetInboxV2(address(0), address(mailboxes[uint32(srcChainId)])));
        inbox = SolverNetInboxV2(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
        inbox.initialize(address(this), solver);
        setRoutes(ISolverNetOutbox.Provider.Hyperlane);

        uint256 snapshot = vm.snapshotState();
        test_v2_open_reverts();
        vm.revertToState(snapshot);

        test_v2_open_nativeDeposit_succeeds();
        vm.revertToState(snapshot);

        test_v2_open_erc20Deposit_succeeds();
    }
}
