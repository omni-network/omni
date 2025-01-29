// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_Validate_Test is TestBase {
    using AddrUtils for address;

    function test_validate_reverts() public {
        IERC7683.OnchainCrossChainOrder memory order;
        uint256 rand = vm.randomUint(1, 1000);

        // `fillDeadline` cannot be in the past unless it is 0
        uint256 timestamp = block.timestamp;
        vm.warp(timestamp + 1);
        order.fillDeadline = uint32(timestamp);
        vm.expectRevert(ISolverNetInbox.InvalidFillDeadline.selector);
        inbox.validate(order);
        order.fillDeadline = 0;
        vm.warp(timestamp);

        // `orderDataType` must be `ORDER_DATA_TYPEHASH`
        vm.expectRevert(ISolverNetInbox.InvalidOrderDataTypehash.selector);
        inbox.validate(order);
        order.orderDataType = ORDER_DATA_TYPEHASH;

        // `orderData` cannot be empty
        vm.expectRevert(ISolverNetInbox.InvalidOrderData.selector);
        inbox.validate(order);
        order.orderData = getOrderDataBytes(
            0, bytes32(0), 0, bytes(""), new ISolverNet.Deposit[](0), new ISolverNet.TokenExpense[](0)
        );

        // `orderData.destChainId` cannot be 0
        vm.expectRevert(ISolverNetInbox.InvalidDestChainId.selector);
        inbox.validate(order);
        ISolverNet.Call[] memory calls = new ISolverNet.Call[](0);
        order.orderData = abi.encode(
            ISolverNet.OrderData({
                owner: user,
                destChainId: destChainId,
                calls: calls,
                deposits: new ISolverNet.Deposit[](0),
                expenses: new ISolverNet.TokenExpense[](0)
            })
        );

        // `orderData.calls.length` must be greater than 0
        vm.expectRevert(ISolverNetInbox.NoCalls.selector);
        inbox.validate(order);
        order.orderData = getOrderDataBytes(
            destChainId, bytes32(0), 0, bytes(""), new ISolverNet.Deposit[](0), new ISolverNet.TokenExpense[](0)
        );

        // `orderData.calls[0].target` cannot be the zero address
        vm.expectRevert(ISolverNetInbox.NoCallTarget.selector);
        inbox.validate(order);
        order.orderData = getOrderDataBytes(
            destChainId,
            address(vault).toBytes32(),
            0,
            bytes(""),
            new ISolverNet.Deposit[](0),
            new ISolverNet.TokenExpense[](0)
        );

        // `orderData.calls[0].data` cannot be empty if value is zero
        vm.expectRevert(ISolverNetInbox.NoCalldata.selector);
        inbox.validate(order);
        order.orderData = getOrderDataBytes(
            destChainId,
            address(vault).toBytes32(),
            0,
            getVaultCalldata(user, rand * 1 ether),
            new ISolverNet.Deposit[](0),
            new ISolverNet.TokenExpense[](0)
        );

        // `deposits` cannot be empty
        vm.expectRevert(ISolverNetInbox.NoDeposits.selector);
        inbox.validate(order);
        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](1);
        deposits[0].token = address(token1).toBytes32();
        order.orderData = getOrderDataBytes(
            destChainId,
            address(vault).toBytes32(),
            0,
            getVaultCalldata(user, rand * 1 ether),
            deposits,
            new ISolverNet.TokenExpense[](0)
        );

        // `deposits[0].amount` cannot be 0
        vm.expectRevert(ISolverNetInbox.NoDepositAmount.selector);
        inbox.validate(order);
        deposits = new ISolverNet.Deposit[](2);
        deposits[0] = ISolverNet.Deposit({ token: address(0).toBytes32(), amount: rand * 1 ether });
        deposits[1] = ISolverNet.Deposit({ token: address(0).toBytes32(), amount: rand * 1 wei });
        order.orderData = getOrderDataBytes(
            destChainId,
            address(vault).toBytes32(),
            0,
            getVaultCalldata(user, rand * 1 ether),
            deposits,
            new ISolverNet.TokenExpense[](1)
        );

        // `deposits` cannot contain more than one native deposit
        vm.expectRevert(ISolverNetInbox.DuplicateNativeDeposit.selector);
        inbox.validate(order);
        deposits = new ISolverNet.Deposit[](1);
        deposits[0].amount = rand * 1 ether;
        order.orderData = getOrderDataBytes(
            destChainId,
            address(vault).toBytes32(),
            0,
            getVaultCalldata(user, rand * 1 ether),
            deposits,
            new ISolverNet.TokenExpense[](1)
        );

        // `expenses[0].amount` cannot be 0
        vm.expectRevert(ISolverNetInbox.NoExpenseAmount.selector);
        inbox.validate(order);
        ISolverNet.TokenExpense[] memory expenses = new ISolverNet.TokenExpense[](1);
        expenses[0].amount = rand * 1 ether;
        order.orderData = getOrderDataBytes(
            destChainId, address(vault).toBytes32(), 0, getVaultCalldata(user, rand * 1 ether), deposits, expenses
        );

        // If `expenses[0].token` is the zero address, `expenses[0].amount` must match total native value of all calls
        vm.expectRevert(ISolverNetInbox.InvalidNativeExpense.selector);
        inbox.validate(order);
        expenses = new ISolverNet.TokenExpense[](2);
        expenses[0].amount = rand * 1 ether;
        expenses[1].amount = rand * 1 ether;
        order.orderData = getOrderDataBytes(
            destChainId,
            address(vault).toBytes32(),
            rand * 1 ether,
            getVaultCalldata(user, rand * 1 ether),
            deposits,
            expenses
        );

        // `expenses` cannot contain more than one native expense
        vm.expectRevert(ISolverNetInbox.DuplicateNativeExpense.selector);
        inbox.validate(order);
        expenses = new ISolverNet.TokenExpense[](1);
        expenses[0].amount = rand * 1 ether;

        // Reset value param and properly set expense token address
        expenses[0].token = address(token2).toBytes32();
        order.orderData = getOrderDataBytes(
            destChainId, address(vault).toBytes32(), 0, getVaultCalldata(user, rand * 1 ether), deposits, expenses
        );

        assertTrue(inbox.validate(order));
    }

    function test_validate_succeeds() public {
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        assertTrue(inbox.validate(order));
    }
}
