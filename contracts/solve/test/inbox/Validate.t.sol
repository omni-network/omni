// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_Validate_Test is TestBase {
    function test_validate_reverts() public {
        IERC7683.OnchainCrossChainOrder memory order =
            IERC7683.OnchainCrossChainOrder({ fillDeadline: 1, orderDataType: bytes32(0), orderData: "" });

        // `fillDeadline` must be in the future
        vm.expectRevert(ISolverNetInbox.InvalidFillDeadline.selector);
        inbox.validate(order);
        order.fillDeadline = defaultFillDeadline;

        // `orderDataType` must be correct
        vm.expectRevert(ISolverNetInbox.InvalidOrderTypehash.selector);
        inbox.validate(order);
        order.orderDataType = HashLib.OLD_ORDERDATA_TYPEHASH;

        // `orderData` must not be empty
        vm.expectRevert(ISolverNetInbox.InvalidOrderData.selector);
        inbox.validate(order);
        order.orderData = "milady";

        // `orderData` must be valid SolverNet.OrderData
        vm.expectRevert();
        inbox.validate(order);
        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: address(0),
            destChainId: 0,
            deposit: SolverNet.Deposit({ token: address(0), amount: 0 }),
            calls: new SolverNet.Call[](0),
            expenses: new SolverNet.TokenExpense[](1)
        });
        order.orderData = abi.encode(orderData);

        // `destChainId` must be non-zero
        vm.expectRevert(ISolverNetInbox.InvalidDestinationChainId.selector);
        inbox.validate(order);
        orderData.destChainId = destChainId;
        order.orderData = abi.encode(orderData);

        // `calls` must not be empty
        vm.expectRevert(ISolverNetInbox.InvalidMissingCalls.selector);
        inbox.validate(order);
        orderData.calls = new SolverNet.Call[](33);
        order.orderData = abi.encode(orderData);

        // `calls` must not exceed array length limit of 32
        vm.expectRevert(ISolverNetInbox.InvalidArrayLength.selector);
        inbox.validate(order);
        orderData.calls = new SolverNet.Call[](1);
        orderData.calls[0].target = address(erc20Vault);
        orderData.expenses = new SolverNet.TokenExpense[](33);
        order.orderData = abi.encode(orderData);

        // `expenses` must not exceed array length limit of 32
        vm.expectRevert(ISolverNetInbox.InvalidArrayLength.selector);
        inbox.validate(order);
        orderData.expenses = new SolverNet.TokenExpense[](1);
        order.orderData = abi.encode(orderData);

        // `expenses` must not have a zero token
        vm.expectRevert(ISolverNetInbox.InvalidExpenseToken.selector);
        inbox.validate(order);
        orderData.expenses[0].token = address(token2);
        order.orderData = abi.encode(orderData);

        // `expenses` must not have a zero amount
        vm.expectRevert(ISolverNetInbox.InvalidExpenseAmount.selector);
        inbox.validate(order);
        orderData.expenses[0].amount = defaultAmount;
        order.orderData = abi.encode(orderData);
    }

    function test_validate_succeeds() public view {
        (, IERC7683.OnchainCrossChainOrder memory order) = getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        inbox.validate(order);
    }

    function test_validate_hyperlane() public {
        address impl = address(new SolverNetInbox(address(0), address(mailboxes[uint32(srcChainId)])));
        inbox = SolverNetInbox(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
        inbox.initialize(address(this), solver);
        setRoutes(ISolverNetOutbox.Provider.Hyperlane);

        uint256 snapshot = vm.snapshotState();
        test_validate_reverts();
        vm.revertToState(snapshot);

        test_validate_succeeds();
    }
}
