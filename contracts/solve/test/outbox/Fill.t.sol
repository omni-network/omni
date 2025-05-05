// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { Refunder } from "test/utils/Refunder.sol";

contract SolverNet_Outbox_Fill_Test is TestBase {
    Refunder internal refunder;

    function setUp() public override {
        super.setUp();
        refunder = new Refunder();
    }

    function test_fill_reverts(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(2)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        vm.chainId(destChainId);
        bytes32 orderId = inbox.getNextOrderId(user);
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

        // `msg.value` must be at least `fillFee`
        vm.deal(address(outbox), fillFee);
        vm.expectRevert(ISolverNetOutbox.InsufficientFee.selector);
        outbox.fill(orderId, fillDataBytes, fillerData);

        // fill must be for a configured source chain
        fillData.srcChainId = destChainId + 1;
        fillDataBytes = abi.encode(fillData);
        vm.expectRevert(ISolverNetOutbox.InvalidConfig.selector);
        outbox.fill{ value: fillFee }(orderId, fillDataBytes, fillerData);

        // `fill` cannot be called twice for the same order
        fillData.srcChainId = srcChainId;
        fillDataBytes = abi.encode(fillData);
        outbox.fill{ value: fillFee }(orderId, fillDataBytes, fillerData);
        vm.expectRevert(ISolverNetOutbox.AlreadyFilled.selector);
        outbox.fill{ value: fillFee }(orderId, fillDataBytes, fillerData);

        vm.stopPrank();
    }

    function test_fill_nativeExpense_succeeds(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(2)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");

        vm.chainId(srcChainId);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
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
        provider = uint8(bound(provider, uint8(1), uint8(2)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getErc20ForErc20VaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");

        vm.chainId(srcChainId);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
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
        provider = uint8(bound(provider, uint8(1), uint8(2)));
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
        fundSolver(orderData, fillFee);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
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
        provider = uint8(bound(provider, uint8(1), uint8(2)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        assertTrue(inbox.validate(order), "order should be valid");

        vm.chainId(srcChainId);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        fundSolver(orderData, fillFee * 2);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
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
        provider = uint8(bound(provider, uint8(1), uint8(2)));
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
        fundSolver(orderData, fillFee);

        // simple check to make sure the correct event type is emitted
        if (provider == 1) {
            vm.expectEmit(false, false, false, false, address(portal));
            emit IOmniPortal.XMsg(0, 0, 0, address(0), address(0), bytes(""), 0, 0);
        } else {
            vm.expectEmit(false, false, false, false, address(mailboxes[uint32(destChainId)]));
            emit IMailbox.Dispatch(address(0), 0, bytes32(0), bytes(""));
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
}
