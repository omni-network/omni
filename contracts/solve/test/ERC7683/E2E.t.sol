// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "./TestBase.sol";

contract SolverNet_E2E_Test is TestBase {
    using AddrUtils for address;

    function test_e2e_complete_contract_order() public {
        // Prep: Set chainId to srcChainId
        vm.chainId(srcChainId);

        // 0. Generate order, validate it, resolve it, and prepare deposit tokens
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        assertTrue(inbox.validate(order), "order should be valid");
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        fundUser(resolvedOrder.minReceived);

        assertNullOrder(resolvedOrder.orderId);

        // 1. Open order on srcChain
        vm.prank(user);
        inbox.open(order);

        assertOpenedOrder(resolvedOrder.orderId);

        // 2. Accept order on srcChain
        vm.prank(solver);
        inbox.accept(resolvedOrder.orderId);

        assertAcceptedOrder(resolvedOrder.orderId);

        // Prep: Set chainId to destChainId and give solver some funds
        vm.chainId(destChainId);
        fundSolver(resolvedOrder.maxSpent);

        // 3. Fill order on destChain
        uint256 fillFee = outbox.fillFee(srcChainId);
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        vm.prank(solver);
        outbox.fill{ value: fillFee }(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, bytes(""));

        assertVaultDeposit(resolvedOrder.orderId);
        assertTrue(
            outbox.didFill(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData),
            "order should be filled"
        );

        // Prep: Set chainId back to srcChainId
        vm.chainId(srcChainId);

        // 4. Mock markFulfilled call from destChain to srcChain
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(ISolverNetInbox.markFilled, (resolvedOrder.orderId, fillHash)),
            100_000
        );

        assertFulfilledOrder(resolvedOrder.orderId);

        // 5. Claim order deposits on srcChain as solver
        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);

        assertClaimedOrder(resolvedOrder.orderId);
    }

    function test_e2e_complete_token_transfer_order() public {
        // Prep: Set chainId to srcChainId
        vm.chainId(srcChainId);

        // 0. Generate order, validate it, resolve it, and prepare deposit tokens
        uint256 rand = vm.randomUint(1, 1000);
        ISolverNet.Deposit[] memory deposit = new ISolverNet.Deposit[](1);
        deposit[0] = ISolverNet.Deposit({ token: address(token1).toBytes32(), amount: rand * 1 ether });

        ISolverNet.TokenExpense[] memory expense = new ISolverNet.TokenExpense[](1);
        expense[0] =
            ISolverNet.TokenExpense({ token: address(token2).toBytes32(), spender: bytes32(0), amount: rand * 1 ether });

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: destChainId,
            target: address(token2).toBytes32(),
            value: 0,
            data: abi.encodeCall(ERC20.transfer, (user, rand * 1 ether)),
            expenses: expense
        });

        ISolverNet.OrderData memory orderData = ISolverNet.OrderData({ user: user, call: call, deposits: deposit });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 minutes),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        assertTrue(inbox.validate(order), "order should be valid");
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(resolvedOrder.minReceived);

        // 1. Open order on srcChain
        vm.prank(user);
        inbox.open(order);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending),
            resolvedOrder.orderId,
            "order should be opened"
        );
        assertEq(
            token1.balanceOf(address(inbox)),
            resolvedOrder.minReceived[0].amount,
            "token1 should be deposited into inbox"
        );

        // 2. Accept order on srcChain
        vm.prank(solver);
        inbox.accept(resolvedOrder.orderId);

        // Prep: Set chainId to destChainId and give solver some funds
        vm.chainId(destChainId);
        fundSolver(resolvedOrder.maxSpent);

        // 3. Fill order on destChain
        uint256 fillFee = outbox.fillFee(srcChainId);
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        vm.prank(solver);
        outbox.fill{ value: fillFee }(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, bytes(""));

        assertTrue(
            outbox.didFill(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData),
            "order should be filled"
        );

        // Prep: Set chainId back to srcChainId
        vm.chainId(srcChainId);

        // 4. Mock markFulfilled call from destChain to srcChain
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(ISolverNetInbox.markFilled, (resolvedOrder.orderId, fillHash)),
            100_000
        );

        // 5. Claim order deposits on srcChain as solver
        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed),
            resolvedOrder.orderId,
            "order should be claimed"
        );
        assertEq(token1.balanceOf(address(inbox)), 0, "token1 should be claimed from inbox");
        assertEq(
            token1.balanceOf(address(solver)), resolvedOrder.minReceived[0].amount, "token1 should be claimed by solver"
        );
        assertEq(token2.balanceOf(address(user)), resolvedOrder.maxSpent[0].amount, "token2 should be received by user");
        assertEq(token2.balanceOf(address(solver)), 0, "token2 should be sent by solver");
    }

    function test_multi_token_deposits_and_expenses() public {
        // Prep: Set chainId to srcChainId
        vm.chainId(srcChainId);

        // 0. Generate order, validate it, resolve it, and prepare deposit tokens
        address[] memory srcDeposits = new address[](2);
        srcDeposits[0] = address(token1);
        srcDeposits[1] = address(token2);
        address[] memory destDeposits = new address[](2);
        destDeposits[0] = address(token3);
        destDeposits[1] = address(token4);

        IERC7683.OnchainCrossChainOrder memory order = randMultiTokenOrder(srcDeposits, destDeposits);
        assertTrue(inbox.validate(order));
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        fundUser(resolvedOrder.minReceived);

        // 1. Open order on srcChain
        vm.prank(user);
        inbox.open(order);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending),
            resolvedOrder.orderId,
            "order should be opened"
        );
        assertEq(
            token1.balanceOf(address(inbox)),
            resolvedOrder.minReceived[0].amount,
            "token1 should be deposited into inbox"
        );
        assertEq(
            token2.balanceOf(address(inbox)),
            resolvedOrder.minReceived[1].amount,
            "token2 should be deposited into inbox"
        );

        // 2. Accept order on srcChain
        vm.prank(solver);
        inbox.accept(resolvedOrder.orderId);

        // Prep: Set chainId to destChainId and give solver some funds
        vm.chainId(destChainId);
        fundSolver(resolvedOrder.maxSpent);

        // 3. Fill order on destChain
        uint256 fillFee = outbox.fillFee(srcChainId);
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        vm.prank(solver);
        outbox.fill{ value: fillFee }(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, bytes(""));

        assertTrue(
            outbox.didFill(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData),
            "order should be filled"
        );

        // Prep: Set chainId back to srcChainId
        vm.chainId(srcChainId);

        // 4. Mock markFulfilled call from destChain to srcChain
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(ISolverNetInbox.markFilled, (resolvedOrder.orderId, fillHash)),
            100_000
        );

        // 5. Claim order deposits on srcChain as solver
        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed),
            resolvedOrder.orderId,
            "order should be claimed"
        );
        assertEq(token1.balanceOf(address(inbox)), 0, "token1 should be claimed from inbox");
        assertEq(token2.balanceOf(address(inbox)), 0, "token2 should be claimed from inbox");
        assertEq(
            token1.balanceOf(address(solver)), resolvedOrder.minReceived[0].amount, "token1 should be claimed by solver"
        );
        assertEq(
            token2.balanceOf(address(solver)), resolvedOrder.minReceived[1].amount, "token2 should be claimed by solver"
        );
        assertEq(
            multiTokenVault.balances(address(user), address(token3)),
            resolvedOrder.maxSpent[0].amount,
            "token3 should be deposited into vault"
        );
        assertEq(
            multiTokenVault.balances(address(user), address(token4)),
            resolvedOrder.maxSpent[1].amount,
            "token4 should be deposited into vault"
        );
        assertEq(token3.balanceOf(address(solver)), 0, "token3 should be deposited by solver");
        assertEq(token4.balanceOf(address(solver)), 0, "token4 should be deposited by solver");
    }

    function test_mixed_deposits_and_expenses() public {
        // Prep: Set chainId to srcChainId
        vm.chainId(srcChainId);

        // 0. Generate order, validate it, resolve it, and prepare deposit tokens
        address[] memory srcDeposits = new address[](2);
        srcDeposits[0] = address(token1);
        srcDeposits[1] = address(0);
        address[] memory destDeposits = new address[](2);
        destDeposits[0] = address(0);
        destDeposits[1] = address(token4);

        IERC7683.OnchainCrossChainOrder memory order = randMultiTokenOrder(srcDeposits, destDeposits);
        assertTrue(inbox.validate(order));
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        fundUser(resolvedOrder.minReceived);

        // 1. Open order on srcChain
        vm.prank(user);
        inbox.open{ value: resolvedOrder.minReceived[1].amount }(order);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending),
            resolvedOrder.orderId,
            "open: order should be opened"
        );
        assertEq(
            token1.balanceOf(address(inbox)),
            resolvedOrder.minReceived[0].amount,
            "open: token1 should be deposited into inbox"
        );
        assertEq(
            address(inbox).balance,
            resolvedOrder.minReceived[1].amount,
            "open: native token should be deposited into inbox"
        );

        // 2. Accept order on srcChain
        vm.prank(solver);
        inbox.accept(resolvedOrder.orderId);

        // Prep: Set chainId to destChainId and give solver some funds
        vm.chainId(destChainId);
        fundSolver(resolvedOrder.maxSpent);

        // 3. Fill order on destChain
        uint256 fillFee = outbox.fillFee(srcChainId);
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        vm.prank(solver);
        outbox.fill{ value: fillFee + resolvedOrder.maxSpent[1].amount }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, bytes("")
        );

        assertTrue(
            outbox.didFill(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData),
            "order should be filled"
        );

        // Prep: Set chainId back to srcChainId
        vm.chainId(srcChainId);

        // 4. Mock markFulfilled call from destChain to srcChain
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(ISolverNetInbox.markFilled, (resolvedOrder.orderId, fillHash)),
            100_000
        );

        // 5. Claim order deposits on srcChain as solver
        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);

        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed),
            resolvedOrder.orderId,
            "claim: order should be claimed"
        );
        assertEq(token1.balanceOf(address(inbox)), 0, "claim: token1 should be claimed from inbox");
        assertEq(address(inbox).balance, 0, "claim: native token should be claimed from inbox");
        assertEq(
            token1.balanceOf(address(solver)),
            resolvedOrder.minReceived[0].amount,
            "claim: token1 should be claimed by solver"
        );
        assertEq(
            address(solver).balance,
            resolvedOrder.minReceived[1].amount,
            "claim: native token should be claimed by solver"
        );
        assertEq(
            multiTokenVault.balances(address(user), address(0)),
            resolvedOrder.maxSpent[1].amount,
            "claim: native token should be deposited into vault"
        );
        assertEq(
            multiTokenVault.balances(address(user), address(token4)),
            resolvedOrder.maxSpent[0].amount,
            "claim: token4 should be deposited into vault"
        );
        assertEq(token4.balanceOf(address(solver)), 0, "claim: token4 should be deposited by solver");
    }
}
