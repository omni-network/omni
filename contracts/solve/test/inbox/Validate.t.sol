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
            0, bytes32(0), 0, bytes(""), new ISolverNet.TokenExpense[](0), new ISolverNet.Deposit[](0)
        );

        // `call.chainId` cannot be 0
        vm.expectRevert(ISolverNetInbox.NoCallChainId.selector);
        inbox.validate(order);
        order.orderData = getOrderDataBytes(
            destChainId, bytes32(0), 0, bytes(""), new ISolverNet.TokenExpense[](0), new ISolverNet.Deposit[](0)
        );

        // `call.target` cannot be the zero address
        vm.expectRevert(ISolverNetInbox.NoCallTarget.selector);
        inbox.validate(order);
        order.orderData = getOrderDataBytes(
            destChainId,
            address(vault).toBytes32(),
            0,
            bytes(""),
            new ISolverNet.TokenExpense[](0),
            new ISolverNet.Deposit[](0)
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
            new ISolverNet.TokenExpense[](0),
            deposits
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
            new ISolverNet.TokenExpense[](1),
            deposits
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
            new ISolverNet.TokenExpense[](1),
            deposits
        );

        // `call.expenses[0].token` cannot be the zero address
        vm.expectRevert(ISolverNetInbox.NoExpenseToken.selector);
        inbox.validate(order);
        ISolverNet.TokenExpense[] memory expenses = new ISolverNet.TokenExpense[](1);
        expenses[0].token = address(token2).toBytes32();
        order.orderData = getOrderDataBytes(
            destChainId, address(vault).toBytes32(), 0, getVaultCalldata(user, rand * 1 ether), expenses, deposits
        );

        // `call.expenses[0].amount` cannot be 0
        vm.expectRevert(ISolverNetInbox.NoExpenseAmount.selector);
        inbox.validate(order);
        expenses[0].amount = rand * 1 ether;
        order.orderData = getOrderDataBytes(
            destChainId, address(vault).toBytes32(), 0, getVaultCalldata(user, rand * 1 ether), expenses, deposits
        );

        assertTrue(inbox.validate(order));
    }

    function test_validate_succeeds() public {
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        assertTrue(inbox.validate(order));
    }
}
