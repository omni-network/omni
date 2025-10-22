// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_Handle_Test is TestBase {
    using AddrUtils for address;

    function setUp() public override {
        super.setUp();
        setRoutes(ISolverNetOutbox.Provider.Hyperlane);
    }

    function test_handle_reverts() public {
        // call must come from the Hyperlane mailbox
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.handle(destinationDomain, bytes32(0), bytes(""));

        // order must be pending
        bytes32 orderId = inbox.getNextOnchainOrderId(user);
        bytes memory message = mailboxes[destinationDomain]
        .buildMessage(originDomain, bytes32(0), address(inbox).toBytes32(), abi.encode(orderId, bytes32(0), address(0)));
        mailboxes[originDomain].addInboundMessage(message);
        vm.expectRevert(ISolverNetInbox.OrderNotPending.selector);
        mailboxes[originDomain].processNextInboundMessage();

        // prep: open a valid order
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");

        fundUser(orderData);
        vm.prank(user);
        inbox.open{ value: defaultAmount }(order);

        // order must be filled on the correct chain
        message = mailboxes[invalidDomain]
        .buildMessage(originDomain, bytes32(0), address(inbox).toBytes32(), abi.encode(orderId, bytes32(0), address(0)));
        mailboxes[originDomain].addInboundMessage(message);
        vm.expectRevert(ISolverNetInbox.WrongSourceChain.selector);
        mailboxes[originDomain].processInboundMessage(1);

        // order must be filled by the outbox
        message = mailboxes[destinationDomain]
        .buildMessage(originDomain, bytes32(0), address(inbox).toBytes32(), abi.encode(orderId, bytes32(0), address(0)));
        mailboxes[originDomain].addInboundMessage(message);
        vm.expectRevert(Ownable.Unauthorized.selector);
        mailboxes[originDomain].processInboundMessage(2);

        // order must have a matching fill hash
        message = mailboxes[destinationDomain]
        .buildMessage(
            originDomain,
            address(outbox).toBytes32(),
            address(inbox).toBytes32(),
            abi.encode(orderId, bytes32(0), address(0))
        );
        mailboxes[originDomain].addInboundMessage(message);
        vm.expectRevert(ISolverNetInbox.WrongFillHash.selector);
        mailboxes[originDomain].processInboundMessage(3);
    }

    function test_handle_succeeds() public {
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        fundUser(orderData);
        vm.prank(user);
        inbox.open(order);

        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        bytes memory message = mailboxes[destinationDomain]
        .buildMessage(
            originDomain,
            address(outbox).toBytes32(),
            address(inbox).toBytes32(),
            abi.encode(resolvedOrder.orderId, fillhash, solver)
        );
        mailboxes[originDomain].addInboundMessage(message);

        vm.expectEmit(true, true, true, true);
        emit ISolverNetInbox.Filled(resolvedOrder.orderId, fillhash, solver);
        mailboxes[originDomain].processNextInboundMessage();

        assertStatus(resolvedOrder.orderId, ISolverNetInbox.Status.Filled);
    }
}
