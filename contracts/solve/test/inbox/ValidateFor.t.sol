// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_ValidateFor_Test is TestBase {
    function test_validateFor_reverts() public {
        IERC7683.GaslessCrossChainOrder memory order = IERC7683.GaslessCrossChainOrder({
            originSettler: address(0),
            user: address(0),
            nonce: 0,
            originChainId: 0,
            openDeadline: 0,
            fillDeadline: 0,
            orderDataType: bytes32(0),
            orderData: ""
        });

        // `originSettler` must be the inbox contract
        vm.expectRevert(ISolverNetInbox.InvalidOriginSettler.selector);
        inbox.validateFor(order);
        order.originSettler = address(inbox);

        // `user` must be specified
        vm.expectRevert(ISolverNetInbox.InvalidUser.selector);
        inbox.validateFor(order);
        order.user = user;

        // `nonce` must be non-zero
        vm.expectRevert(ISolverNetInbox.InvalidNonce.selector);
        inbox.validateFor(order);
        order.nonce = 1;

        // `originChainId` must be for the proper chain
        vm.expectRevert(ISolverNetInbox.InvalidOriginChainId.selector);
        inbox.validateFor(order);
        order.originChainId = srcChainId;

        // `openDeadline` must be in the future
        vm.expectRevert(ISolverNetInbox.InvalidOpenDeadline.selector);
        inbox.validateFor(order);
        order.openDeadline = defaultOpenDeadline;

        // `fillDeadline` must be after `openDeadline`
        vm.expectRevert(ISolverNetInbox.InvalidFillDeadline.selector);
        inbox.validateFor(order);
        order.fillDeadline = defaultFillDeadline;

        // `orderDataType` must be a valid typehash
        vm.expectRevert(ISolverNetInbox.InvalidOrderTypehash.selector);
        inbox.validateFor(order);
        // legacy typehash must work
        order.orderDataType = HashLib.OLD_ORDERDATA_TYPEHASH;
        vm.expectRevert(ISolverNetInbox.InvalidOrderData.selector); // error after typehash check
        inbox.validateFor(order);
        // new typehash must work
        order.orderDataType = HashLib.OMNIORDERDATA_TYPEHASH;
        vm.expectRevert(ISolverNetInbox.InvalidOrderData.selector); // error after typehash check
        inbox.validateFor(order);

        // `orderData` must be longer than 0 bytes
        vm.expectRevert(ISolverNetInbox.InvalidOrderData.selector);
        inbox.validateFor(order);
        order.orderData = hex"0420";

        // `orderData` must be valid SolverNet.OrderData
        vm.expectRevert();
        inbox.validateFor(order);
        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: address(0),
            destChainId: 0,
            deposit: SolverNet.Deposit({ token: address(0), amount: defaultAmount }),
            calls: new SolverNet.Call[](0),
            expenses: new SolverNet.TokenExpense[](0)
        });
        order.orderData = abi.encode(orderData);

        // `orderData.destChainId` must be non-zero
        vm.expectRevert(ISolverNetInbox.InvalidDestinationChainId.selector);
        inbox.validateFor(order);
        orderData.destChainId = destChainId;
        order.orderData = abi.encode(orderData);

        // `orderData.deposit.token` must be non-zero for gasless orders
        vm.expectRevert(ISolverNetInbox.InvalidNativeDeposit.selector);
        inbox.validateFor(order);
        orderData.deposit = SolverNet.Deposit({ token: address(token1), amount: defaultAmount });
        order.orderData = abi.encode(orderData);

        // `orderData.calls` must not be empty
        vm.expectRevert(ISolverNetInbox.InvalidMissingCalls.selector);
        inbox.validateFor(order);
        orderData.calls = new SolverNet.Call[](33);
        order.orderData = abi.encode(orderData);

        // `orderData.calls` must not exceed array length limit of 32
        vm.expectRevert(ISolverNetInbox.InvalidArrayLength.selector);
        inbox.validateFor(order);
        orderData.calls = new SolverNet.Call[](1);
        orderData.expenses = new SolverNet.TokenExpense[](33);
        order.orderData = abi.encode(orderData);

        // `orderData.expenses` must not exceed array length limit of 32
        vm.expectRevert(ISolverNetInbox.InvalidArrayLength.selector);
        inbox.validateFor(order);
        orderData.expenses = new SolverNet.TokenExpense[](1);
        order.orderData = abi.encode(orderData);

        // token address in `orderData.expenses` must be non-zero
        vm.expectRevert(ISolverNetInbox.InvalidExpenseToken.selector);
        inbox.validateFor(order);
        orderData.expenses[0].token = address(token2);
        order.orderData = abi.encode(orderData);

        // token amount in `orderData.expenses` must be non-zero
        vm.expectRevert(ISolverNetInbox.InvalidExpenseAmount.selector);
        inbox.validateFor(order);
        orderData.expenses[0].amount = defaultAmount;
        order.orderData = abi.encode(orderData);

        assertTrue(inbox.validateFor(order), "order should now be valid");
    }

    function test_validateFor_succeeds() public view {
        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(erc20Vault), 0, user, defaultAmount);

        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](1);
        expenses[0] = getExpense(address(erc20Vault), address(token2), defaultAmount);

        (, IERC7683.GaslessCrossChainOrder memory order) = getGaslessOrder(
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
    }

    function test_validateFor_hyperlane() public {
        address impl = address(new SolverNetInbox(address(0), address(mailboxes[uint32(srcChainId)])));
        inbox = SolverNetInbox(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
        inbox.initialize(address(this), solver);
        setRoutes(ISolverNetOutbox.Provider.Hyperlane);

        uint256 snapshot = vm.snapshotState();
        test_validateFor_reverts();
        vm.revertToState(snapshot);

        test_validateFor_succeeds();
    }
}
