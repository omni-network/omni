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

    function test_resolveOrder_erc20_deposit_succeeds() public {
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        assertResolved(user, resolvedOrder.orderId, order, resolvedOrder);
    }

    function test_resolveOrder_native_deposit_succeeds() public {
        IERC7683.OnchainCrossChainOrder memory order = randNativeOrder();
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        assertResolved(user, resolvedOrder.orderId, order, resolvedOrder);
    }

    function test_resolveOrder_call_with_value_succeeds() public {
        IERC7683.OnchainCrossChainOrder memory order = randNativeOrder();
        ISolverNet.OrderData memory orderData = abi.decode(order.orderData, (ISolverNet.OrderData));
        orderData.call.value = 1 ether;
        order.orderData = abi.encode(orderData);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        assertResolved(user, resolvedOrder.orderId, order, resolvedOrder);
    }
}
