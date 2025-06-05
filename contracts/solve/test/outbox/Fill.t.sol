// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { Refunder } from "test/utils/Refunder.sol";
import { Receiver as ReceiverOnly } from "test/utils/Receiver.sol";

contract SolverNet_Outbox_Fill_Test is TestBase {
    Refunder internal refunder;

    function setUp() public override {
        super.setUp();
        refunder = new Refunder();
    }

    function test_fill_reverts(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(3)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        vm.chainId(destChainId);
        bytes32 orderId = inbox.getNextOnchainOrderId(user);
        SolverNet.FillOriginData memory fillData = SolverNet.FillOriginData({
            srcChainId: srcChainId,
            destChainId: srcChainId,
            fillDeadline: uint32(block.timestamp),
            calls: new SolverNet.Call[](0),
            expenses: new SolverNet.TokenExpense[](0)
        });
        bytes memory fillDataBytes = abi.encode(fillData);
        bytes memory fillerData = hex"0420";
        uint256 fillFee = outbox.fillFee(fillDataBytes);
        vm.deal(solver, fillFee * 2);

        // filler must be a whitelisted solver
        vm.expectRevert(Ownable.Unauthorized.selector);
        outbox.fill(orderId, "", "");
        vm.startPrank(solver);

        // `originData` must be encoded `FillOriginData`
        vm.expectRevert();
        outbox.fill(orderId, fillerData, "");

        // `destChainId` must match the local chain
        vm.expectRevert(ISolverNetOutbox.WrongDestChain.selector);
        outbox.fill(orderId, fillDataBytes, "");
        fillData.destChainId = destChainId;
        fillDataBytes = abi.encode(fillData);
        vm.warp(block.timestamp + 1);

        // `fillDeadline` cannot be in the past
        vm.expectRevert(ISolverNetOutbox.FillDeadlinePassed.selector);
        outbox.fill(orderId, fillDataBytes, "");
        vm.warp(block.timestamp - 1);

        // `fillerData` must be empty or 32 bytes
        vm.expectRevert(ISolverNetOutbox.BadFillerData.selector);
        outbox.fill(orderId, fillDataBytes, fillerData);
        fillerData = abi.encode(solver);

        // `msg.value` must be at least `fillFee`, skipped for trusted routes as fee is 0
        if (fillFee > 0) {
            vm.deal(address(outbox), fillFee);
            vm.expectRevert(ISolverNetOutbox.InsufficientFee.selector);
            outbox.fill(orderId, fillDataBytes, fillerData);
        }

        // fill must be for a configured source chain
        fillData.srcChainId = destChainId + 1;
        fillDataBytes = abi.encode(fillData);
        vm.expectRevert(ISolverNetOutbox.InvalidConfig.selector);
        outbox.fill{ value: fillFee }(orderId, fillDataBytes, fillerData);

        // `fill` cannot be called twice for the same order
        uint256 snapshot = vm.snapshotState();
        fillData.srcChainId = srcChainId;
        fillDataBytes = abi.encode(fillData);
        outbox.fill{ value: fillFee }(orderId, fillDataBytes, fillerData);
        vm.expectRevert(ISolverNetOutbox.AlreadyFilled.selector);
        outbox.fill{ value: fillFee }(orderId, fillDataBytes, fillerData);
        vm.revertToState(snapshot);

        // `fill` cannot be called twice if it had been filled using the old fill hash
        bytes32 oldFillHash = keccak256(abi.encode(orderId, fillDataBytes));
        bytes32 value = bytes32(type(uint256).max);
        bytes32 slot = keccak256(abi.encode(oldFillHash, uint256(4)));
        vm.store(address(outbox), slot, value);
        vm.expectRevert(ISolverNetOutbox.AlreadyFilled.selector);
        outbox.fill{ value: fillFee }(orderId, fillDataBytes, fillerData);
        vm.stopPrank();

        // `fill` cannot be called if the order uses less than half of what the solver was directed to spend
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        orderData.calls[0].params = abi.encode(user, (defaultAmount / 2) - 1);
        order.orderData = abi.encode(orderData);
        assertTrue(inbox.validate(order), "order should be valid");

        vm.chainId(srcChainId);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);
        vm.prank(solver);
        vm.expectRevert(ISolverNetOutbox.InsufficientSpend.selector);
        outbox.fill{ value: fillFee }(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, fillerData);
    }

    function test_fill_nativeExpense_succeeds(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(3)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");

        vm.chainId(srcChainId);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else if (provider == 2) {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
        } else {
            vm.expectEmit(true, true, true, true, address(outbox));
            emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        }

        vm.prank(solver);
        outbox.fill{ value: defaultAmount + fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );

        assertEq(address(outbox).balance, 0, "outbox native balance after");
        assertEq(address(outbox.executor()).balance, 0, "executor native balance after");
        assertEq(solver.balance, 0, "solver balance after");
        assertEq(nativeVault.balances(user), defaultAmount, "vault deposit balance after");
        assertEq(address(nativeVault).balance, defaultAmount, "vault native balance after");
    }

    function test_fill_erc20Expense_succeeds(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(3)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");

        vm.chainId(srcChainId);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else if (provider == 2) {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
        } else {
            vm.expectEmit(true, true, true, true, address(outbox));
            emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        }

        vm.prank(solver);
        outbox.fill{ value: fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );

        assertEq(token2.balanceOf(address(outbox)), 0, "outbox token2 balance after");
        assertEq(token2.balanceOf(address(outbox.executor())), 0, "executor token2 balance after");
        assertEq(token2.balanceOf(solver), 0, "solver token2 balance after");
        assertEq(erc20Vault.balances(user), defaultAmount, "vault deposit balance after");
        assertEq(token2.balanceOf(address(erc20Vault)), defaultAmount, "vault token2 balance after");
    }

    function test_fill_call_refund_succeeds(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(3)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        (SolverNet.OrderData memory orderData,) = getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        orderData.calls[0].target = address(refunder);
        IERC7683.OnchainCrossChainOrder memory order = getOrder(block.timestamp + 1, orderData);
        assertTrue(inbox.validate(order), "order should be valid");

        vm.chainId(srcChainId);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else if (provider == 2) {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
        } else {
            vm.expectEmit(true, true, true, true, address(outbox));
            emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        }

        vm.prank(solver);
        outbox.fill{ value: defaultAmount + fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );

        assertEq(address(outbox).balance, 0, "outbox native balance after");
        assertEq(address(outbox.executor()).balance, 0, "executor native balance after");
        assertEq(solver.balance, defaultAmount, "solver balance after");
        assertEq(address(refunder).balance, 0, "refunder balance after");
    }

    function test_fill_native_overpayment_refund_succeeds(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(3)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");

        vm.chainId(srcChainId);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee * 2);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else if (provider == 2) {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
        } else {
            vm.expectEmit(true, true, true, true, address(outbox));
            emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        }

        vm.prank(solver);
        outbox.fill{ value: defaultAmount + (fillFee * 2) }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );

        assertEq(address(outbox).balance, 0, "outbox native balance after");
        assertEq(address(outbox.executor()).balance, 0, "executor native balance after");
        assertEq(solver.balance, fillFee, "solver balance after");
        assertEq(nativeVault.balances(user), defaultAmount, "vault deposit balance after");
        assertEq(address(nativeVault).balance, defaultAmount, "vault native balance after");
    }

    function test_fill_erc20_overpayment_refund_succeeds(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(3)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        (SolverNet.OrderData memory orderData,) = getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        orderData.expenses[0].amount = defaultAmount * 2;
        IERC7683.OnchainCrossChainOrder memory order = getOrder(block.timestamp + 1, orderData);
        assertTrue(inbox.validate(order), "order should be valid");

        vm.chainId(srcChainId);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else if (provider == 2) {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
        } else {
            vm.expectEmit(true, true, true, true, address(outbox));
            emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        }

        vm.prank(solver);
        outbox.fill{ value: fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );

        assertEq(token2.balanceOf(address(outbox)), 0, "outbox token2 balance after");
        assertEq(token2.balanceOf(address(outbox.executor())), 0, "executor token2 balance after");
        assertEq(token2.balanceOf(solver), defaultAmount, "solver token2 balance after");
        assertEq(erc20Vault.balances(user), defaultAmount, "vault deposit balance after");
        assertEq(token2.balanceOf(address(erc20Vault)), defaultAmount, "vault token2 balance after");
        assertEq(token2.allowance(address(outbox), address(outbox.executor())), 0, "outbox token2 allowance after");
    }

    // @dev Test that a native transfer fill to a contract with only `receive()` succeeds.
    function test_fill_receive_succeeds() public {
        vm.chainId(destChainId);

        address receiver = address(new ReceiverOnly());

        SolverNet.FillOriginData memory fillData = SolverNet.FillOriginData({
            srcChainId: srcChainId,
            destChainId: destChainId,
            fillDeadline: uint32(block.timestamp + 1),
            calls: new SolverNet.Call[](1),
            expenses: new SolverNet.TokenExpense[](0)
        });

        // 1 ETH transfer to the receiver
        fillData.calls[0] = SolverNet.Call({ target: receiver, value: 1 ether, selector: bytes4(0), params: bytes("") });

        bytes32 orderId = bytes32(uint256(1234));
        uint256 fillFee = outbox.fillFee(abi.encode(fillData));

        // Fill order
        vm.deal(solver, 1 ether + fillFee);
        vm.prank(solver);
        outbox.fill{ value: 1 ether + fillFee }(orderId, abi.encode(fillData), abi.encode(solver));

        // Assert received
        assertEq(receiver.balance, 1 ether, "receiver should have received 1 ether");
    }

    function test_fill_erc20_partialExecutorBalance_succeeds(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(3)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");

        vm.chainId(srcChainId);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        // fund executor with defaultAmount - 1 of token2
        token2.mint(address(outbox.executor()), defaultAmount - 1);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else if (provider == 2) {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
        } else {
            vm.expectEmit(true, true, true, true, address(outbox));
            emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        }

        vm.prank(solver);
        outbox.fill{ value: fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );

        assertEq(token2.balanceOf(address(outbox)), 0, "outbox token2 balance after");
        assertEq(token2.balanceOf(address(outbox.executor())), 0, "executor token2 balance after");
        assertEq(token2.balanceOf(solver), defaultAmount - 1, "solver token2 balance after");
        assertEq(erc20Vault.balances(user), defaultAmount, "vault deposit balance after");
        assertEq(token2.balanceOf(address(erc20Vault)), defaultAmount, "vault token2 balance after");
    }

    function test_fill_erc20_fullExecutorBalance_succeeds(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(3)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");

        vm.chainId(srcChainId);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        // fund executor with defaultAmount of token2
        token2.mint(address(outbox.executor()), defaultAmount);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else if (provider == 2) {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
        } else {
            vm.expectEmit(true, true, true, true, address(outbox));
            emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        }

        vm.prank(solver);
        outbox.fill{ value: fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );

        assertEq(token2.balanceOf(address(outbox)), 0, "outbox token2 balance after");
        assertEq(token2.balanceOf(address(outbox.executor())), 0, "executor token2 balance after");
        assertEq(token2.balanceOf(solver), defaultAmount, "solver token2 balance after");
        assertEq(erc20Vault.balances(user), defaultAmount, "vault deposit balance after");
        assertEq(token2.balanceOf(address(erc20Vault)), defaultAmount, "vault token2 balance after");
    }

    function test_fill_erc20_excessExecutorBalance_succeeds(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(3)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");

        vm.chainId(srcChainId);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        // fund executor with defaultAmount * 2 of token2
        token2.mint(address(outbox.executor()), defaultAmount * 2);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else if (provider == 2) {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
        } else {
            vm.expectEmit(true, true, true, true, address(outbox));
            emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        }

        vm.prank(solver);
        outbox.fill{ value: fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );

        assertEq(token2.balanceOf(address(outbox)), 0, "outbox token2 balance after");
        assertEq(token2.balanceOf(address(outbox.executor())), 0, "executor token2 balance after");
        assertEq(token2.balanceOf(solver), defaultAmount * 2, "solver token2 balance after");
        assertEq(erc20Vault.balances(user), defaultAmount, "vault deposit balance after");
        assertEq(token2.balanceOf(address(erc20Vault)), defaultAmount, "vault token2 balance after");
    }
}
