// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_Resolve_Test is TestBase {
    using AddrUtils for address;
    using AddrUtils for bytes32;

    function setUp() public override {
        super.setUp();
        vm.chainId(srcChainId);
    }

    function test_resolveOrder_succeeds() public {
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        ISolverNet.OrderData memory orderData = abi.decode(order.orderData, (ISolverNet.OrderData));
        ISolverNet.Call memory orderCall = orderData.call;
        ISolverNet.TokenExpense[] memory orderExpenses = orderCall.expenses;
        ISolverNet.Deposit[] memory orderDeposits = orderData.deposits;

        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        assertEq(resolvedOrder.user, user, "resolved order: user");
        assertEq(resolvedOrder.originChainId, srcChainId, "resolved order: origin chain id");
        assertEq(resolvedOrder.openDeadline, uint32(block.timestamp), "resolved order: open deadline");
        assertEq(resolvedOrder.fillDeadline, order.fillDeadline, "resolved order: fill deadline");
        assertEq(resolvedOrder.orderId, inbox.getNextId(), "resolved order: order id");

        assertEq(resolvedOrder.maxSpent.length, 1, "resolved order: max spent length");
        assertEq(resolvedOrder.maxSpent[0].token.toAddress(), address(token2), "resolved order: max spent token");
        assertEq(resolvedOrder.maxSpent[0].amount, orderExpenses[0].amount, "resolved order: max spent amount");
        assertEq(
            resolvedOrder.maxSpent[0].recipient, address(outbox).toBytes32(), "resolved order: max spent recipient"
        );
        assertEq(resolvedOrder.maxSpent[0].chainId, destChainId, "resolved order: max spent chain id");

        assertEq(resolvedOrder.minReceived.length, 1, "resolved order: min received length");
        assertEq(resolvedOrder.minReceived[0].token.toAddress(), address(token1), "resolved order: min received token");
        assertEq(resolvedOrder.minReceived[0].amount, orderDeposits[0].amount, "resolved order: min received amount");
        assertEq(resolvedOrder.minReceived[0].recipient, bytes32(0), "resolved order: min received recipient");
        assertEq(resolvedOrder.minReceived[0].chainId, srcChainId, "resolved order: min received chain id");

        assertEq(resolvedOrder.fillInstructions.length, 1, "resolved order: fill instructions length");
        assertEq(
            resolvedOrder.fillInstructions[0].destinationChainId,
            destChainId,
            "resolved order: fill instructions chain id"
        );
        assertEq(
            resolvedOrder.fillInstructions[0].destinationSettler,
            address(outbox).toBytes32(),
            "resolved order: fill instructions destination"
        );
        assertEq(
            keccak256(resolvedOrder.fillInstructions[0].originData),
            keccak256(abi.encode(ISolverNet.FillOriginData({ srcChainId: srcChainId, call: orderCall }))),
            "resolved order: fill instructions origin data"
        );
    }
}
