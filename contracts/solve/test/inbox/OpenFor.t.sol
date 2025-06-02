// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_OpenFor_Test is TestBase {
    function test_openFor_reverts() public {
        inbox.pauseOpen(true);

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(erc20Vault), 0, user, defaultAmount);
        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](1);
        expenses[0] = getExpense(address(erc20Vault), address(token2), defaultAmount);

        (SolverNet.OmniOrderData memory orderData, IERC7683.GaslessCrossChainOrder memory order) = getGaslessOrder(
            user,
            user,
            1,
            destChainId,
            defaultOpenDeadline,
            defaultFillDeadline,
            address(1),
            defaultAmount,
            calls,
            expenses
        );
        assertTrue(inbox.validateFor(order), "order should be valid");

        bytes32 digest = HashLib.gaslessOrderDigest(order, orderData, address(inbox));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(userPk, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Should revert when `open` is paused
        vm.expectRevert(ISolverNetInbox.IsPaused.selector);
        vm.prank(user);
        inbox.openFor(order, signature, "");

        // Should revert if `open` and `close` are paused
        inbox.pauseAll(true);
        vm.expectRevert(ISolverNetInbox.AllPaused.selector);
        vm.prank(user);
        inbox.openFor(order, signature, "");

        inbox.pauseAll(false);

        // Should revert if a native deposit is specified for a gasless order
        orderData.deposit = SolverNet.Deposit({ token: address(0), amount: 1 ether });
        order.orderData = abi.encode(orderData);

        digest = HashLib.gaslessOrderDigest(order, orderData, address(inbox));
        (v, r, s) = vm.sign(userPk, digest);
        signature = abi.encodePacked(r, s, v);

        vm.expectRevert(ISolverNetInbox.InvalidNativeDeposit.selector);
        vm.prank(user);
        inbox.openFor(order, signature, "");

        // Should revert if less tokens are received than expected due to max transfer balance override
        orderData.deposit = SolverNet.Deposit({ token: address(maxTransferToken), amount: type(uint96).max });
        order.orderData = abi.encode(orderData);

        digest = HashLib.gaslessOrderDigest(order, orderData, address(inbox));
        (v, r, s) = vm.sign(userPk, digest);
        signature = abi.encodePacked(r, s, v);

        maxTransferToken.mint(user, defaultAmount);
        vm.expectRevert(ISolverNetInbox.InvalidERC20Deposit.selector);
        vm.prank(user);
        inbox.openFor(order, signature, "");

        // Should revert if less tokens are received than expected due to fee on transfer
        orderData.deposit = SolverNet.Deposit({ token: address(feeOnTransferToken), amount: 1 ether });
        order.orderData = abi.encode(orderData);

        digest = HashLib.gaslessOrderDigest(order, orderData, address(inbox));
        (v, r, s) = vm.sign(userPk, digest);
        signature = abi.encodePacked(r, s, v);

        fundUser(orderData, true);
        vm.expectRevert(ISolverNetInbox.InvalidERC20Deposit.selector);
        vm.prank(user);
        inbox.openFor(order, signature, "");
    }

    function test_openFor_succeeds() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(erc20Vault), 0, user, defaultAmount);
        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](1);
        expenses[0] = getExpense(address(erc20Vault), address(token2), defaultAmount);

        (SolverNet.OmniOrderData memory orderData, IERC7683.GaslessCrossChainOrder memory order) = getGaslessOrder(
            user,
            user,
            1,
            destChainId,
            defaultOpenDeadline,
            defaultFillDeadline,
            address(token1),
            defaultAmount,
            calls,
            expenses
        );
        assertTrue(inbox.validateFor(order), "order should be valid");

        bytes32 digest = HashLib.gaslessOrderDigest(order, orderData, address(inbox));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(userPk, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolveFor(order, "");
        bytes32 orderId = inbox.getOrderId(user, 1, true);
        assertEq(resolvedOrder.orderId, orderId, "order id should match");

        fundUser(orderData, true);
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit IERC7683.Open(resolvedOrder.orderId, resolvedOrder);
        inbox.openFor(order, signature, "");

        (IERC7683.ResolvedCrossChainOrder memory resolved2,, uint248 orderOffset) = inbox.getOrder(orderId);
        resolved2.openDeadline = defaultOpenDeadline; // This piece of data is not preserved by the inbox

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

        uint256 snapshot = vm.snapshotState();
        test_openFor_reverts();
        vm.revertToState(snapshot);

        test_openFor_succeeds();
    }
}
