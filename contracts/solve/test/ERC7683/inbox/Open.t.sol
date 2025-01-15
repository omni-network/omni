// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_Open_Test is TestBase {
    using AddrUtils for address;

    function setUp() public override {
        super.setUp();
        vm.chainId(srcChainId);

        // Fund user with ETH for native token tests
        vm.deal(user, 100 ether);
    }

    function test_open_succeeds() public {
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        // Store initial state for comparison
        bytes32 expectedOrderId = inbox.getNextId();
        uint256 initialUserBalance = token1.balanceOf(user);
        uint256 initialInboxBalance = token1.balanceOf(address(inbox));

        // Expect the event with the correct namespace and data
        vm.expectEmit(true, true, true, true, address(inbox));
        emit IERC7683.Open(expectedOrderId, resolvedOrder);

        vm.prank(user);
        inbox.open(order);

        // Verify order state and history
        (
            IERC7683.ResolvedCrossChainOrder memory storedOrder,
            ISolverNetInbox.OrderState memory state,
            ISolverNetInbox.StatusUpdate[] memory history
        ) = inbox.getOrder(expectedOrderId);

        // Verify that stored resolved order aligns with the original order
        assertResolved(user, resolvedOrder.orderId, order, storedOrder);

        // Verify order state
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Pending), "order state: status");
        assertEq(state.acceptedBy, address(0), "order state: accepted by");

        // Verify order history
        assertEq(history.length, 1, "order history: length");
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "order history: status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "order history: timestamp");

        // Verify latest order ID by status
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending), expectedOrderId, "latest pending order id"
        );

        // Verify token transfers
        assertEq(token1.balanceOf(user), initialUserBalance - resolvedOrder.minReceived[0].amount, "user balance after");
        assertEq(
            token1.balanceOf(address(inbox)),
            initialInboxBalance + resolvedOrder.minReceived[0].amount,
            "inbox balance after"
        );
    }

    function test_open_reverts_native_deposit_without_value() public {
        // Create order with native deposit but don't send value
        ISolverNet.TokenExpense[] memory expenses = new ISolverNet.TokenExpense[](1);
        expenses[0] = ISolverNet.TokenExpense({
            token: address(token2).toBytes32(),
            spender: address(vault).toBytes32(),
            amount: 1 ether
        });

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: destChainId,
            target: address(vault).toBytes32(),
            value: 0,
            data: getVaultCalldata(user, 1 ether),
            expenses: expenses
        });

        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](1);
        deposits[0] = ISolverNet.Deposit({ token: bytes32(0), amount: 1 ether });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 days),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(ISolverNet.OrderData({ call: call, deposits: deposits }))
        });

        vm.prank(user);
        vm.expectRevert(ISolverNetInbox.InvalidNativeDeposit.selector);
        inbox.open(order);
    }

    function test_open_reverts_native_deposit_wrong_value() public {
        // Create order with native deposit but send wrong value
        ISolverNet.TokenExpense[] memory expenses = new ISolverNet.TokenExpense[](1);
        expenses[0] = ISolverNet.TokenExpense({
            token: address(token2).toBytes32(),
            spender: address(vault).toBytes32(),
            amount: 1 ether
        });

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: destChainId,
            target: address(vault).toBytes32(),
            value: 0,
            data: getVaultCalldata(user, 1 ether),
            expenses: expenses
        });

        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](1);
        deposits[0] = ISolverNet.Deposit({ token: bytes32(0), amount: 1 ether });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 days),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(ISolverNet.OrderData({ call: call, deposits: deposits }))
        });

        vm.prank(user);
        vm.expectRevert(ISolverNetInbox.InvalidNativeDeposit.selector);
        inbox.open{ value: 2 ether }(order);
    }

    function test_open_reverts_multiple_native_deposits() public {
        // Create order with multiple native deposits
        ISolverNet.TokenExpense[] memory expenses = new ISolverNet.TokenExpense[](1);
        expenses[0] = ISolverNet.TokenExpense({
            token: address(token2).toBytes32(),
            spender: address(vault).toBytes32(),
            amount: 1 ether
        });

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: destChainId,
            target: address(vault).toBytes32(),
            value: 0,
            data: getVaultCalldata(user, 1 ether),
            expenses: expenses
        });

        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](2);
        deposits[0] = ISolverNet.Deposit({ token: bytes32(0), amount: 1 ether });
        deposits[1] = ISolverNet.Deposit({ token: bytes32(0), amount: 1 ether });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 days),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(ISolverNet.OrderData({ call: call, deposits: deposits }))
        });

        vm.prank(user);
        vm.expectRevert(ISolverNetInbox.DuplicateNativeDeposit.selector);
        inbox.open{ value: 1 ether }(order);
    }

    function test_open_reverts_value_without_native_deposit() public {
        // Create order with no native deposit but send value
        ISolverNet.TokenExpense[] memory expenses = new ISolverNet.TokenExpense[](1);
        expenses[0] = ISolverNet.TokenExpense({
            token: address(token2).toBytes32(),
            spender: address(vault).toBytes32(),
            amount: 1 ether
        });

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: destChainId,
            target: address(vault).toBytes32(),
            value: 0,
            data: getVaultCalldata(user, 1 ether),
            expenses: expenses
        });

        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](1);
        deposits[0] = ISolverNet.Deposit({ token: address(token1).toBytes32(), amount: 1 ether });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 days),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(ISolverNet.OrderData({ call: call, deposits: deposits }))
        });

        // Mint and approve tokens for the ERC20 deposit
        token1.mint(user, 1 ether);
        vm.prank(user);
        token1.approve(address(inbox), 1 ether);

        vm.prank(user);
        vm.expectRevert(ISolverNetInbox.InvalidNativeDeposit.selector);
        inbox.open{ value: 1 ether }(order);
    }

    function test_open_reverts_zero_amount_erc20_deposit() public {
        // Create order with zero amount ERC20 deposit
        ISolverNet.TokenExpense[] memory expenses = new ISolverNet.TokenExpense[](1);
        expenses[0] = ISolverNet.TokenExpense({
            token: address(token2).toBytes32(),
            spender: address(vault).toBytes32(),
            amount: 1 ether
        });

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: destChainId,
            target: address(vault).toBytes32(),
            value: 0,
            data: getVaultCalldata(user, 1 ether),
            expenses: expenses
        });

        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](1);
        deposits[0] = ISolverNet.Deposit({ token: address(token1).toBytes32(), amount: 0 });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 days),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(ISolverNet.OrderData({ call: call, deposits: deposits }))
        });

        vm.prank(user);
        vm.expectRevert(ISolverNetInbox.NoDepositAmount.selector);
        inbox.open(order);
    }

    function test_open_reverts_insufficient_erc20_allowance() public {
        // Create order with ERC20 deposit but don't approve tokens
        ISolverNet.TokenExpense[] memory expenses = new ISolverNet.TokenExpense[](1);
        expenses[0] = ISolverNet.TokenExpense({
            token: address(token2).toBytes32(),
            spender: address(vault).toBytes32(),
            amount: 1 ether
        });

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: destChainId,
            target: address(vault).toBytes32(),
            value: 0,
            data: getVaultCalldata(user, 1 ether),
            expenses: expenses
        });

        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](1);
        deposits[0] = ISolverNet.Deposit({ token: address(token1).toBytes32(), amount: 1 ether });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 days),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(ISolverNet.OrderData({ call: call, deposits: deposits }))
        });

        // Mint tokens but don't approve
        token1.mint(user, 1 ether);

        vm.prank(user);
        vm.expectRevert(); // ERC20 transfer error
        inbox.open(order);
    }
}
