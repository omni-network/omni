// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_MarkFilled_Test is TestBase {
    function test_v2_markFilled_reverts() public {
        // order must be pending
        bytes32 orderId = inbox.getNextOnchainOrderId(user);
        vm.expectRevert(ISolverNetInboxV2.OrderNotPending.selector);
        inbox.markFilled(orderId, bytes32(0), address(0));

        // prep: open a valid order
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");

        fundUser(orderData);
        vm.prank(user);
        inbox.open{ value: defaultAmount }(order);

        // call must come from the Omni portal
        vm.expectRevert(ISolverNetInboxV2.WrongSourceChain.selector);
        inbox.markFilled(orderId, bytes32(0), address(0));

        // order must be filled on the correct chain
        vm.expectRevert(ISolverNetInboxV2.WrongSourceChain.selector);
        portal.mockXCall(
            destChainId + 1,
            address(outbox),
            address(inbox),
            abi.encodeWithSelector(ISolverNetInboxV2.markFilled.selector, orderId, bytes32(0), address(0))
        );

        // order must be filled by the outbox
        vm.expectRevert(Ownable.Unauthorized.selector);
        portal.mockXCall(
            destChainId,
            user,
            address(inbox),
            abi.encodeWithSelector(ISolverNetInboxV2.markFilled.selector, orderId, bytes32(0), address(0))
        );

        // order must have a matching fill hash
        vm.expectRevert(ISolverNetInboxV2.WrongFillHash.selector);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeWithSelector(ISolverNetInboxV2.markFilled.selector, orderId, bytes32(0), address(0))
        );

        // portal address must be set for remote orders
        address impl = address(new SolverNetInboxV2(address(0), address(mailboxes[uint32(srcChainId)])));
        inbox = SolverNetInboxV2(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
        inbox.initialize(address(this), solver);

        uint64[] memory chainIds = new uint64[](2);
        chainIds[0] = srcChainId;
        chainIds[1] = destChainId;
        address[] memory outboxes = new address[](2);
        outboxes[0] = address(outbox);
        outboxes[1] = address(outbox);
        inbox.setOutboxes(chainIds, outboxes);

        orderId = inbox.getNextOnchainOrderId(user);
        (orderData, order) = getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        bytes32 fillhash = fillHash(orderId, resolvedOrder.fillInstructions[0].originData);
        assertTrue(inbox.validate(order), "order should be valid");

        fundUser(orderData);
        vm.prank(user);
        inbox.open{ value: defaultAmount }(order);

        vm.expectRevert(ISolverNetInboxV2.WrongSourceChain.selector);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeWithSelector(ISolverNetInboxV2.markFilled.selector, orderId, fillhash, solver)
        );
    }

    function test_v2_markFilled_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        inbox.open(order);

        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetInboxV2.Filled(resolvedOrder.orderId, fillhash, solver);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeWithSelector(ISolverNetInboxV2.markFilled.selector, resolvedOrder.orderId, fillhash, solver)
        );

        assertStatus(resolvedOrder.orderId, ISolverNetInboxV2.Status.Filled);
    }

    function test_v2_markFilled_oldFillHash_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        inbox.open(order);

        bytes32 fillhash = keccak256(abi.encode(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData));
        vm.expectEmit(true, true, true, true);
        emit ISolverNetInboxV2.Filled(resolvedOrder.orderId, fillhash, solver);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeWithSelector(ISolverNetInboxV2.markFilled.selector, resolvedOrder.orderId, fillhash, solver)
        );

        assertStatus(resolvedOrder.orderId, ISolverNetInboxV2.Status.Filled);
    }
}
