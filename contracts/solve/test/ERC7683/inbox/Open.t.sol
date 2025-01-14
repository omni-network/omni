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

        // Verify order data matches resolved order
        assertEq(storedOrder.user, user, "stored order: user");
        assertEq(storedOrder.originChainId, resolvedOrder.originChainId, "stored order: origin chain id");
        assertEq(storedOrder.openDeadline, uint32(block.timestamp), "stored order: open deadline");
        assertEq(storedOrder.fillDeadline, resolvedOrder.fillDeadline, "stored order: fill deadline");
        assertEq(storedOrder.orderId, resolvedOrder.orderId, "stored order: order id");

        // Verify maxSpent (solver outputs)
        assertEq(storedOrder.maxSpent.length, resolvedOrder.maxSpent.length, "stored order: max spent length");
        for (uint256 i = 0; i < storedOrder.maxSpent.length; i++) {
            assertEq(storedOrder.maxSpent[i].token, resolvedOrder.maxSpent[i].token, "stored order: max spent token");
            assertEq(storedOrder.maxSpent[i].amount, resolvedOrder.maxSpent[i].amount, "stored order: max spent amount");
            assertEq(
                storedOrder.maxSpent[i].recipient,
                resolvedOrder.maxSpent[i].recipient,
                "stored order: max spent recipient"
            );
            assertEq(
                storedOrder.maxSpent[i].chainId, resolvedOrder.maxSpent[i].chainId, "stored order: max spent chain id"
            );
        }

        // Verify minReceived (user deposits)
        assertEq(storedOrder.minReceived.length, resolvedOrder.minReceived.length, "stored order: min received length");
        for (uint256 i = 0; i < storedOrder.minReceived.length; i++) {
            assertEq(
                storedOrder.minReceived[i].token, resolvedOrder.minReceived[i].token, "stored order: min received token"
            );
            assertEq(
                storedOrder.minReceived[i].amount,
                resolvedOrder.minReceived[i].amount,
                "stored order: min received amount"
            );
            assertEq(
                storedOrder.minReceived[i].recipient,
                resolvedOrder.minReceived[i].recipient,
                "stored order: min received recipient"
            );
            assertEq(
                storedOrder.minReceived[i].chainId,
                resolvedOrder.minReceived[i].chainId,
                "stored order: min received chain id"
            );
        }

        // Verify fill instructions
        assertEq(
            storedOrder.fillInstructions.length,
            resolvedOrder.fillInstructions.length,
            "stored order: fill instructions length"
        );
        for (uint256 i = 0; i < storedOrder.fillInstructions.length; i++) {
            assertEq(
                storedOrder.fillInstructions[i].destinationChainId,
                resolvedOrder.fillInstructions[i].destinationChainId,
                "stored order: fill instructions chain id"
            );
            assertEq(
                storedOrder.fillInstructions[i].destinationSettler,
                resolvedOrder.fillInstructions[i].destinationSettler,
                "stored order: fill instructions destination"
            );
            assertEq(
                keccak256(storedOrder.fillInstructions[i].originData),
                keccak256(resolvedOrder.fillInstructions[i].originData),
                "stored order: fill instructions origin data"
            );
        }

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
